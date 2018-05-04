package structs

import (
	"database/sql"
	"net/smtp"
)

// database credentials
var (
	DB *sql.DB

	DBHost     = "localhost"
	DBPort     = 5432
	DBName     = "postgres"
	DBUser     = "postgres"
	DBPassword = ""
	DBSSLMode  = "disable"
)

// Dummy smtp server credentials
var (
	SMTPServerPort = "mail.example.com:25" // SMTP host with port
	SMTPFrom       = "abhay@example.org"   // From address for email
	SMTPAuth       = smtp.PlainAuth(
		"",
		"abhay@example.com", // smtp username
		"password",          //smtp password
		"mail.example.com",  // smtp host
	)
)

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
// var Issues map[string]Issue
// var Users map[string]User

// func init() {

// Issues = make(map[string]Issue)

// Users = make(map[string]User)

// Hardcoded users

// Users["tokenuser1"] = User{
// 	Email:       "user1@gmail.com",
// 	Username:    "user1",
// 	FirstName:   "Jacob",
// 	LastName:    "Bent",
// 	Password:    "pass1",
// 	AccessToken: "tokenuser1",
// }

// Users["tokenuser2"] = User{
// 	Email:       "user2@gmail.com",
// 	Username:    "user2",
// 	FirstName:   "Lilly",
// 	LastName:    "Holland",
// 	Password:    "pass2",
// 	AccessToken: "tokenuser2",
// }

// Users["tokenuser3"] = User{
// 	Email:       "user3@gmail.com",
// 	Username:    "user3",
// 	FirstName:   "Sam",
// 	LastName:    "Rowley",
// 	Password:    "pass3",
// 	AccessToken: "tokenuser3",
// }

// }
