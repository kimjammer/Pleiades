package main

type Projects struct {
}

type Project struct {
	id          string
	title       string
	description string
	users       []User
	tasks       []Task
	polls       []Poll
}

type User struct {
	id           string
	firstName    string
	lastName     string
	email        string
	password     string //TODO: Change me to correct type
	phoneNumber  string
	userPhoto    string //TODO: Change me to correct type
	availability []Availability
	projects     []string
}

type Availability struct {
	dayOfWeek   int
	startOffset int
	endOffset   int
}

type Poll struct {
	id          string
	title       string
	description string
	options     []Option
}

type Task struct {
	id           string
	title        string
	description  string
	dueDate      string //TODO: Change me to correct type
	kanbanColumn string
	timeEstimate int
	completed    bool
	sessions     []Session
	assignees    []string
}

type Option struct {
	id            string
	title         string
	likedUsers    []string
	neutralUsers  []string
	dislikedUsers []string
}

type Session struct {
	id        string
	startTime string //TODO: Change me to correct type
	endTime   string //TODO: Change me to correct type
	user      string
}
