package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Util to reset the DB to a known state for testing
func resetDB() {
	_ = db.Drop(context.TODO())

	userId, _ := primitive.ObjectIDFromHex("67c7b20021675682e4395270")

	//Create a user
	user := User{
		Id:           userId,
		FirstName:    "Joe",
		LastName:     "Smith",
		Email:        "example@example.com",
		Password:     "notthepassword",
		PhoneNumber:  "1234567890",
		Projects:     []string{"53ed4d28-9279-4b4e-9256-b1e693332625"},
		Availability: []Availability{},
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err == nil {
		user.Password = string(hashedPassword)
	}

	project := Project{
		Id:          "53ed4d28-9279-4b4e-9256-b1e693332625",
		Title:       "Test Project",
		Description: "Test Description",
		Users:       []UserAndLeft{{"67c7b20021675682e4395270", false}},
		Tasks:       []Task{},
		Polls:       []Poll{},
	}
	_, _ = db.Collection("projects").InsertOne(context.TODO(), project)
	_, _ = db.Collection("users").InsertOne(context.TODO(), user)
}

func secondPerson() {
	userId, _ := primitive.ObjectIDFromHex("67e5b7abef0709e1426eed50")

	//Create a user
	user := User{
		Id:           userId,
		FirstName:    "Joe",
		LastName:     "Biden",
		Email:        "example@us.gov",
		Password:     "isthepasswordforsure",
		PhoneNumber:  "9234567890",
		Projects:     []string{},
		Availability: []Availability{},
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err == nil {
		user.Password = string(hashedPassword)
	}

	_, _ = db.Collection("users").InsertOne(context.TODO(), user)
}
