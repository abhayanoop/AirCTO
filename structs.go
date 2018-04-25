package main

import "net/smtp"

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

// Dummy smtp server credentials
var (
	smtpServerPort = "mail.example.com:25" // smtp host with port
	smtpFrom       = "abhay@example.org"   // From address for email
	smtpAuth       = smtp.PlainAuth(
		"",
		"abhay@example.com", // smtp username
		"password",          //smtp password
		"mail.example.com",  // smtp host
	)
)

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

	Users["tokenuser3"] = User{
		Email:       "user3@gmail.com",
		Username:    "user3",
		FirstName:   "Sam",
		LastName:    "Rowley",
		Password:    "pass3",
		AccessToken: "tokenuser3",
	}

}
