package main

type User struct {
	Email               string
	Username            string
	FirstName, LastName string
	Password            string
	AccessToken         string
}

type Issue struct {
	ID                 string
	Title, Description string
	AssignedTo         User
	CreatedBy          User
	Status             string
}

// Temporary alternative for a database
var Issues map[string]Issue
var Users map[string]User

func init() {

	Issues = make(map[string]Issue)
	Users = make(map[string]User)
}
