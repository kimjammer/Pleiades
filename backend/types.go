package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAndLeft struct {
	User        string
	LeftProject bool
}

type Project struct {
	Id              string `bson:"_id"`
	Title           string
	Description     string
	Users           []UserAndLeft
	Tasks           []Task
	Polls           []Poll
	DemoButtonState string
}

type User struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	FirstName    string
	LastName     string
	Email        string
	Password     string //stored as bcrypt hash (string)
	PhoneNumber  string
	UserPhoto    string //TODO: Change me to correct type
	Availability []Availability
	Projects     []string
}

type Availability map[string]([]int)

type Poll struct {
	Id          string
	Title       string
	Description string
	Options     []Option
}

type Task struct {
	Id           string
	Title        string
	Description  string
	DueDate      string //TODO: Change me to correct type
	KanbanColumn string
	TimeEstimate int
	Completed    bool
	Sessions     []Session
	Assignees    []string
}

type Option struct {
	Id            string
	Title         string
	LikedUsers    []string
	NeutralUsers  []string
	DislikedUsers []string
}

type Session struct {
	Id        string
	StartTime string //TODO: Change me to correct type
	EndTime   string //TODO: Change me to correct type
	User      string
}

type Invitation struct {
	Id              string `bson:"_id"`
	CreatedAt       time.Time
	Project         string
}
