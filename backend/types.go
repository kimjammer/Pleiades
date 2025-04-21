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
	Created         int
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
	UserPhoto    []byte
	Availability []Availability
	Projects     []string
}

type Availability struct {
	DayOfWeek   int
	StartOffset int
	EndOffset   int
}

type Poll struct {
	Id          string `bson:"_id"`
	Title       string
	Description string
	Options     []Option
	DueDate     int //num milliseconds since 1970
}

type Task struct {
	Id           string    `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	DueDate      int       `json:"dueDate"`
	KanbanColumn string    `json:"kanbanColumn"`
	TimeEstimate int       `json:"timeEstimate"`
	Completed    bool      `json:"completed"`
	Sessions     []Session `json:"sessions"`
	Assignees    []string  `json:"assignees"`
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
	StartTime int
	EndTime   int
	User      string
}

type Invitation struct {
	Id        string `bson:"_id"`
	CreatedAt time.Time
	ProjectId string
}

type PwdResetToken struct {
	CreatedAt time.Time
	Token     string
	User      primitive.ObjectID
}
