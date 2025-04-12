package main

/*
Fields must be public (capitalized) to be serialized to JSON
*/

type ProjectsResponse struct {
	Projects []minimalProject `json:"projects"`
}

type minimalProject struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type NewProjectRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type NewPollRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Options     string `json:"options"`
	DueDate     string `json:"dueDate"`
}

type EmailCheckResponse struct {
	Exists bool `json:"exists"`
}

type LoginResponse struct {
	Exists bool `json:"exists"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type PasswordResetRequest struct {
	Token       string
	NewPassword string
}
