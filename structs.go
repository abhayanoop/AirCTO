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
	AssignedTo         string
	CreatedBy          string
	Status             string
}

// Temporary alternative for a database
var Issues map[string]Issue
var Users map[string]User

func init() {

	Issues = make(map[string]Issue)

	Users = make(map[string]User)

	// Hardcoded users

	Users["tokenuser1"] = User{
		Email:       "user1@gmail.com",
		Username:    "user1",
		FirstName:   "Jacob",
		LastName:    "Bent",
		Password:    "pass1",
		AccessToken: "tokenuser1",
	}

	Users["tokenuser2"] = User{
		Email:       "user2@gmail.com",
		Username:    "user2",
		FirstName:   "Lilly",
		LastName:    "Holland",
		Password:    "pass2",
		AccessToken: "tokenuser2",
	}

	Users["tokenabhay"] = User{
		Email:       "abhayanoop1994@gmail.com",
		Username:    "abhayanoop",
		FirstName:   "Abhay",
		LastName:    "Anoop",
		Password:    "myPassword",
		AccessToken: "tokenabhay",
	}

	// Hardcoded issues

	Issues["issue1"] = Issue{
		ID:          "issue1",
		Title:       "Issue 1",
		Description: "first issue",
		AssignedTo:  "abhayanoop",
		CreatedBy:   "abhayanoop",
		Status:      "Open",
	}

}
