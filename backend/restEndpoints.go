package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/mailjet/mailjet-apiv3-go/v4"
	"google.golang.org/api/idtoken"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

func projectsHandler(c *gin.Context) {
	//Get current user
	userId := c.GetString("userId")
	objId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: objId}}
	var crrUser User
	err := db.Collection("users").FindOne(c, filter).Decode(&crrUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			//TODO: Add error handling for user not found
		} else {
			panic(err)
		}
	}

	//Get projects
	var response ProjectsResponse
	for _, projectId := range crrUser.Projects {
		filter := bson.D{{Key: "_id", Value: projectId}}
		var project Project
		err := db.Collection("projects").FindOne(c, filter).Decode(&project)
		if err != nil {
			panic(err)
		}
		response.Projects = append(response.Projects, minimalProject{
			Id:          project.Id,
			Title:       project.Title,
			Description: project.Description,
		})
	}

	c.JSON(http.StatusOK, response)
}

func newProjectHandler(c *gin.Context) {
	log.Println("Creating New Project")
	//Get current user
	userId := c.GetString("userId")
	objId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: objId}}

	var crrUser User
	err := db.Collection("users").FindOne(c, filter).Decode(&crrUser)
	if err != nil {
		//TODO: Add error handling for user not found
		log.Println("User not found!")
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	var newProject NewProjectRequest
	if err := c.ShouldBindJSON(&newProject); err != nil {
		log.Println("Invalid JSON data")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	//Validation
	if newProject.Title == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//Create Project
	project := Project{
		Id:          uuid.New().String(),
		Title:       newProject.Title,
		Description: newProject.Description,
		Created:     int(time.Now().UnixMilli()),
		Users:       []UserAndLeft{{crrUser.Id.Hex(), false}},
		Tasks:       []Task{},
		Polls:       []Poll{},
	}
	_, err = db.Collection("projects").InsertOne(c, project)
	if err != nil {
		log.Println("Error inserting project: ", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	//Add project to user
	crrUser.Projects = append(crrUser.Projects, project.Id)
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "projects", Value: crrUser.Projects}}}}
	_, err = db.Collection("users").UpdateOne(c, filter, update)
	if err != nil {
		log.Println("Error updating user with new project: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

// this function checks to see if an email is already registered in the database
func checkEmail(c *gin.Context) {
	log.Println("checking email")
	email := c.Query("email")
	
	_, err := findUserByEmail(c, email)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"exists": true})
	}
	//c is sent at end of function
}

func mintToken(c *gin.Context, user *User) {
	//set cookie upon successful login (cookie is user id)
	log.Println("User id: ", user.Id, user.Id.Hex())
	token := makeToken(user.Id.Hex())
	c.SetSameSite(http.SameSiteNoneMode)
	log.Println("Token:", token)
	c.SetCookie("token", token, 3600, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{"exists": true, "userId": user.Id})
}

func registerUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	hashedPassword, err := encryptPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	} else {
		newUser.Password = hashedPassword
	}

	createUser(c, &newUser)
}

func login(c *gin.Context) {
	listUsers()
	email := c.Query("email")
	password := c.Query("password")

	user, err := findUserByEmail(c, email)

	if err == nil {
		//check if passwords match
		if checkPassword(user.Password, password) {
			mintToken(c, &user)
		} else {
			c.JSON(http.StatusOK, gin.H{"exists": false})
		}
	}
}

func verifyGoogle(c *gin.Context) (payload *idtoken.Payload, err error) {
	var req struct {
		Credential string `form:"credential" binding:"required"`
	}
	if err = c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	payload, err = idtoken.Validate(c, req.Credential, "")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
	}
	return
}

func googleLogin(c *gin.Context) {
	payload, err := verifyGoogle(c)
	if err != nil {
		return
	}

	email, _ := payload.Claims["email"].(string)

	user, err := findUserByEmail(c, email)
	if err == nil {
		mintToken(c, &user)
	} else {
		// TODO: autocreate?
		return
	}

	c.Redirect(http.StatusSeeOther, os.Getenv("PROTOCOL") + os.Getenv("HOST") + "/home")
}

func googleRegistration(c *gin.Context) {
	payload, err := verifyGoogle(c)
	if err != nil {
		return
	}

	email, _ := payload.Claims["email"].(string)
	name, _ := payload.Claims["name"].(string)

	newUser := User{
		FirstName: name,
		Email: email,
		LastName: "",
	}
	createUser(c, &newUser)
}

func invite(c *gin.Context) {
	invitationId := createInvite(c)
	if invitationId == "" {
		return
	}

	c.String(http.StatusOK, invitationId)
}

func join(c *gin.Context) {
	// get invite entry
	invitation, err := getInvite(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invite not found"})
		return
	}

	userId := c.GetString("userId")

	// Add user to project
	err = updateProject(invitation.ProjectId, func(project *Project) error {
		for i, user := range project.Users {
			if user.User == userId {
				project.Users[i].LeftProject = false
				return nil
			}
		}

		project.Users = append(project.Users, UserAndLeft{User: userId, LeftProject: false})

		return nil
	})
	if err != nil {
		panic("The updater cannot return an error")
	}

	requeryUsersForProject(invitation.ProjectId)

	// Add project to user
	objId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: objId}, {Key: "projects", Value: bson.M{"$ne": invitation.ProjectId}}}
	update := bson.M{"$push": bson.M{"projects": invitation.ProjectId}}
	_, err = db.Collection("users").UpdateOne(c, filter, update)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	user, err := getUserById(c, userId)
	if err != nil {
		panic(err)
	}

	err = applyCommandToProject(invitation.ProjectId, Notify{
		Who:      "",
		Category: "users",
		Title:    user.FirstName + " joined the project!",
		Message:  "Go say hi!",
	})

	c.String(http.StatusOK, "success")
}

func joinInfo(c *gin.Context) {
	invitation, err := getInvite(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"exists": false})
		return
	}

	// Get addl project info TODO: abstract
	filter := bson.D{{Key: "_id", Value: invitation.ProjectId}}
	var project Project
	err = db.Collection("projects").FindOne(c, filter).Decode(&project)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"exists": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"exists": true, "id": invitation.ProjectId, "title": project.Title, "description": project.Description})
}

func setAvailability(c *gin.Context) {
	// Decode new availability
	var newAvailability []Availability
	if err := c.ShouldBindJSON(&newAvailability); err != nil {
		log.Println("Invalid JSON data")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// Write new availability to db
	userId := c.GetString("userId")
	objId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: objId}}
	update := bson.M{"$set": bson.M{"availability": newAvailability}}
	_, err := db.Collection("users").UpdateOne(c, filter, update)
	if err != nil {
		panic(err)
	}

	user, _ := getUser(c) // error not possible b/c already validated

	requeryUser(user)
}

func encryptPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // Use a fixed cost
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func checkPassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

// Reset token cookie
func logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", "", -1, "/", "", true, true)
	c.Status(http.StatusOK)
}

// Used by frontend to check if user has a valid session, since this is an authRequired() endpoint
func verifySession(c *gin.Context) {
	c.Status(http.StatusOK)
}

// used to get basic user info such as name
func getUserInfo(c *gin.Context) {
	//Get current user
	userId := c.GetString("userId")
	objId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: objId}}
	var crrUser User
	err := db.Collection("users").FindOne(c, filter).Decode(&crrUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			//TODO: Add error handling for user not found
		} else {
			panic(err)
		}
	}

	log.Println("User name:", crrUser.FirstName)
	c.JSON(http.StatusOK, gin.H{"firstname": crrUser.FirstName})

}

func uploadProfilePic(c *gin.Context) {
	log.Println("uploading profile pic")
	//get user
	crrUser, err := getUser(c)
	if err != nil {
		// TODO: convert to middleware
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	//log.Println("crrUser: ", crrUser)

	//store profile pic
	var reqBody struct {
		Image string `json:"image"`
	}
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Decode the base64 string to binary data
	imgData, err := base64.StdEncoding.DecodeString(reqBody.Image)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode base64"})
		return
	}
	crrUser.UserPhoto = imgData
	filter := bson.M{"_id": crrUser.Id}
	update := bson.M{"$set": bson.M{"userPhoto": imgData}}
	_, err = db.Collection("users").UpdateOne(c, filter, update)
	log.Println("profile pic uploaded successfully")
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func getProfilePic(c *gin.Context) {
	//Get user
	crrUser, err := getUserById(c, c.Query("id"))
	if err != nil {
		// TODO: convert to middleware
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err != nil { //if no pic
		log.Println("error encountered")
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile picture not found"})
		return
	} else if crrUser.UserPhoto == nil {
		log.Println("no profile pic in database")
		c.JSON(http.StatusOK, gin.H{"found": false})
	}
	log.Println("returning profile pic")
	c.Header("Content-Type", "image/png")
	c.Writer.Write(crrUser.UserPhoto)
}

func fetchName(c *gin.Context) {
	//Get user
	crrUser, err := getUserById(c, c.Query("id"))
	if err != nil {
		// TODO: convert to middleware
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"firstName": crrUser.FirstName, "lastName": crrUser.LastName})
}

func handleEvent(c *gin.Context) {
	// TODO: there should probably be auth for events that need login or spam prevention, but an attacker could spam create tasks anyways
	name := c.Query("name")
	valueStr := c.Query("value")
	if valueStr == "" {
		valueStr = "1"
	}

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing name or value"})
		return
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid value"})
		return
	}

	update := bson.M{"$inc": bson.M{name: value}}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = db.Collection("events").UpdateOne(ctx,
		bson.M{"_id": "events_document"},
		update,
		options.UpdateOne().SetUpsert(true),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func reportStats(c *gin.Context) {
	var result bson.M
	err := db.Collection("events").FindOne(c, bson.M{"_id": "events_document"}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusOK, bson.M{})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func forgotPasswordHandler(c *gin.Context) {
	//Read request
	var forgotRequest ForgotPasswordRequest
	if err := c.ShouldBindJSON(&forgotRequest); err != nil {
		log.Println("Invalid JSON data")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	//Check if email exists
	var user User
	err := db.Collection("users").FindOne(c, bson.M{"email": forgotRequest.Email}).Decode(&user)

	if err == nil {
		//Generate magic link
		protocol := os.Getenv("PROTOCOL")
		host := os.Getenv("HOST")
		token := uuid.New().String()
		magicLink := protocol + host + "/resetPassword?token=" + token

		pwdResetTokens := db.Collection("pwdResetTokens")

		ttl := int32(86400) // 1 day, in seconds
		if TEST {
			ttl = 30 // seconds
		}
		indexModel := mongo.IndexModel{
			Keys:    bson.D{{Key: "createdat", Value: 1}},
			Options: options.Index().SetExpireAfterSeconds(ttl).SetName("ttl_index"),
		}
		_, err := pwdResetTokens.Indexes().CreateOne(c, indexModel)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		// Create record
		resetToken := PwdResetToken{
			CreatedAt: time.Now(),
			Token:     token,
			User:      user.Id,
		}
		_, err = pwdResetTokens.InsertOne(c, resetToken)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		//DEBUG
		log.Println("Magic link: ", magicLink)

		//TODO: Send email with reset link
		messagesInfo := []mailjet.InfoMessagesV31{
			{
				From: &mailjet.RecipientV31{
					Email: "support@kimjammer.com",
					Name:  "Pleiades Support",
				},
				To: &mailjet.RecipientsV31{
					mailjet.RecipientV31{
						Email: user.Email,
						Name:  user.FirstName + " " + user.LastName,
					},
				},
				Subject:  "Password Reset",
				TextPart: "Forgot your password?\nClick the link below to reset it:\n" + magicLink,
				HTMLPart: "<h1 style=\"font-family:sans-serif\">Forgot your password?</h1><p style=\"font-family:sans-serif\">Click the link below to reset it:</p><div style=\"display:flex; justify-content:center;\"><a style=\"font-family:sans-serif; text-decoration:none; background-color:#272e3f; padding:0.8em; border-radius:0.5em; color:#ffffff;\" href=\"" + magicLink + "\">Reset Password<a></div>",
			},
		}
		messages := mailjet.MessagesV31{Info: messagesInfo}
		if mailjetClient != nil {
			_, err = mailjetClient.SendMailV31(&messages)
			if err != nil {
				log.Println(err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		} else {
			log.Println("No Mailjet client, email not sent!")
		}
	}

	c.Status(http.StatusOK)
}

func resetPasswordHandler(c *gin.Context) {
	//Read token and new password
	var resetRequest PasswordResetRequest
	if err := c.ShouldBindJSON(&resetRequest); err != nil {
		log.Println("Invalid JSON data")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	//Check if token is valid
	var result PwdResetToken
	collection := db.Collection("pwdResetTokens")
	filter := bson.M{"token": resetRequest.Token}
	err := collection.FindOne(c, filter).Decode(&result)
	if err != nil {
		log.Println(err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.AbortWithStatus(http.StatusBadRequest)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	//Find user to update
	var user User
	err = db.Collection("users").FindOne(c, bson.M{"_id": result.User}).Decode(&user)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	//Update password
	hashedPassword, err := encryptPassword(resetRequest.NewPassword)
	update := bson.D{{"$set", bson.D{{"password", hashedPassword}}}}
	userFiler := bson.M{"_id": user.Id}
	_, err = db.Collection("users").UpdateOne(c, userFiler, update)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	//Delete token
	_, err = collection.DeleteOne(c, filter)

	c.Status(http.StatusOK)
}

func sendInviteEmail(c *gin.Context) {
	// Get name and email from query parameters
	name := c.Query("name")
	email := c.Query("email")
	if name == "" || email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing name or email"})
		return
	}

	// Generate magic link
	protocol := os.Getenv("PROTOCOL")
	host := os.Getenv("HOST")
	inviteId := createInvite(c)
	if inviteId == "" {
		return
	}
	magicLink := protocol + host + "/join?id=" + inviteId
	message := "Your teammates are inviting you to join their Pleaides project: "

	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "support@kimjammer.com",
				Name:  "Pleiades Support",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: email,
					Name:  name,
				},
			},
			Subject:  "Pleiades Invitaiton",
			TextPart: message + magicLink,
			HTMLPart: fmt.Sprintf("%s <a href=\"%s\">%s</a>", message, magicLink, magicLink),
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	if mailjetClient != nil {
		_, err := mailjetClient.SendMailV31(&messages)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	} else {
		log.Println("No Mailjet client, email not sent!")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func purdueDirectory(c *gin.Context) {
	name := c.Query("name")
	// Make a POST request to https://www.purdue.edu/directory/ with form encoded body `SearchString: <name>`
	form := url.Values{"SearchString": {name}}

	resp, err := http.Post(
		"https://www.purdue.edu/directory/",
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Find the emails
	emails := doc.Find(".icon-key").Map(func(i int, s *goquery.Selection) string {
		// Search for alias because not all users have public emails
		// TODO: some users have a different email from their alias.
		return s.Next().Text() + "@purdue.edu"
	})

	// Find the names
	names := doc.Find(".cn-name").Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	})

	// Create a 2d array of names and emails as [name, email][]
	nameEmailMap := make([][2]string, len(names))
	for i := range names {
		nameEmailMap[i][0] = names[i]
		nameEmailMap[i][1] = emails[i]
	}

	c.JSON(http.StatusOK, nameEmailMap)
}

func getUserTasks(c *gin.Context) {
	crrUser, err := getUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	allProjects := db.Collection("projects")
	var userTasks []Task
	var projectNames []string
	//var projectIDs []string

	for _, projectId := range crrUser.Projects {
		var project Project
		err := allProjects.FindOne(c, bson.M{"_id": projectId}).Decode(&project)
		if err == nil {
			for _, task := range project.Tasks {
				for _, assignee := range task.Assignees {
					if assignee == crrUser.Id.Hex() {
						userTasks = append(userTasks, task)
						projectNames = append(projectNames, project.Title)
						//projectIDs = append(projectIDs, project.Id)
						break
					}
				}
			}
		}
	}
	log.Println(userTasks)
	log.Println(projectNames)
	c.JSON(http.StatusOK, gin.H{"success": true, "tasks": userTasks,
		"projectNames": projectNames})
}

func flipNotif(c *gin.Context) {
	crrUser, err := getUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	var req FlipRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	log.Println("before flip: ", crrUser.NotifSettings)

	//flip
	crrUser.NotifSettings[req.NotifIndex] = !crrUser.NotifSettings[req.NotifIndex]
	log.Println("after flip: ", crrUser.NotifSettings)

	//update in database
	_, err = db.Collection("users").UpdateByID(c, crrUser.Id,
		bson.M{"$set": bson.M{"notifSettings": crrUser.NotifSettings}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}
	log.Println("after flip: ", crrUser.NotifSettings)

	c.JSON(http.StatusOK, gin.H{"success": true})

}

func getNotifSettings(c *gin.Context) {
	crrUser, err := getUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	//TODO: make sure backend sends back correct boolean array
	log.Println("NotifSettings: ", crrUser.NotifSettings)
	c.JSON(http.StatusOK, gin.H{"success": true, "notifSettings": crrUser.NotifSettings})
}

// TEMPORARY
func listUsers() {
	collection := db.Collection("users") // Reference the "users" collection

	// MongoDB filter (empty filter `{}` to get all documents)
	filter := bson.M{}

	// Cursor to iterate over all documents
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error finding users:", err)
	}
	defer cursor.Close(context.TODO()) // Close cursor when function ends

	fmt.Println("Current Users in Database:")

	for cursor.Next(context.TODO()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			log.Println("Error decoding user:", err)
			continue
		}
		fmt.Printf("ID: %s, Name: %s %s, Email: %s, Password: %s\n", user.Id, user.FirstName, user.LastName, user.Email, user.Password)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal("Cursor error:", err)
	}
}
func deleteAllUsers() {
	collection := db.Collection("users") // Reference the "users" collection

	result, err := collection.DeleteMany(context.TODO(), bson.M{}) // Empty filter {} means delete all
	if err != nil {
		log.Fatal("Error deleting users:", err)
	}

	log.Printf("Deleted %d users from the database\n", result.DeletedCount)
}
func deleteAllProjects() {
	collection := db.Collection("projects") // Reference the "users" collection

	result, err := collection.DeleteMany(context.TODO(), bson.M{}) // Empty filter {} means delete all
	if err != nil {
		log.Fatal("Error deleting projects:", err)
	}

	log.Printf("Deleted %d users from the database\n", result.DeletedCount)
}
