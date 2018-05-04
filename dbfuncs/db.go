package dbfuncs

import (
	"AirCTO/structs"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

var DB = structs.DB

func init() {

	var (
		pgInfo string
		err    error
	)

	if structs.DBPassword != "" {
		pgInfo = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			structs.DBHost, structs.DBPort, structs.DBUser,
			structs.DBPassword, structs.DBName, structs.DBSSLMode)
	} else {
		pgInfo = fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s sslmode=%s",
			structs.DBHost, structs.DBPort, structs.DBUser,
			structs.DBName, structs.DBSSLMode)
	}

	if DB, err = sql.Open("postgres", pgInfo); err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	if _, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS users(
		username TEXT PRIMARY KEY, 
		email TEXT NOT NULL, 
		first_name TEXT, 
		last_name TEXT,
		password TEXT, 
		access_token TEXT UNIQUE)
		`); err != nil {
		panic(err)
	}

	if _, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS issues(
		id SERIAL PRIMARY KEY, 
		title TEXT NOT NULL,
		description TEXT,
	    assigned_to TEXT REFERENCES users(username) ON DELETE CASCADE, 
		created_by TEXT REFERENCES users(username) ON DELETE CASCADE,
		 status TEXT)`); err != nil {
		panic(err)
	}

	users := []string{
		`('user1', 'user1@gmail.com', 'Jacob', 'Bent', 'pass1', 'tokenuser1')`,
		`('user2', 'user2@gmail.com', 'Lilly', 'Holland', 'pass2', 'tokenuser2')`,
		`('user3', 'user3@gmail.com', 'Sam', 'Rowley', 'pass3', 'tokenuser3')`,

		// add users here
	}

	addUsersQuery := `INSERT INTO users(username, email, first_name, last_name, 
		password, access_token) VALUES ` + strings.Join(users, ",") + ` 
		ON CONFLICT(username) DO NOTHING`

	if _, err = DB.Exec(addUsersQuery); err != nil {
		panic(err)
	}

}

func CreateIssue(issue structs.Issue) (err error) {

	query := `INSERT INTO issues(id, title, description, assigned_to, created_by, status)
		 	  VALUES(DEFAULT, $1, $2, $3, $4, $5)`

	_, err = DB.Exec(query, issue.Title, issue.Description, issue.AssignedTo,
		issue.CreatedBy, issue.Status)

	return
}

func GetAllIssues() (issues []structs.Issue, err error) {

	var rows *sql.Rows

	query := `SELECT id, title, description, assigned_to, created_by, status from issues`

	rows, err = DB.Query(query)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {

		var issue structs.Issue

		if err = rows.Scan(&issue.ID, &issue.Title, &issue.Description,
			&issue.AssignedTo, &issue.CreatedBy, &issue.Status); err != nil {
			return
		}

		issues = append(issues, issue)
	}

	return
}

func GetIssue(id string) (issue structs.Issue, err error) {

	query := `SELECT id, title, description, assigned_to, created_by, status 
 			  FROM issues WHERE id = $1`

	err = DB.QueryRow(query, id).Scan(&issue.ID, &issue.Title, &issue.Description,
		&issue.AssignedTo, &issue.CreatedBy, &issue.Status)

	return
}

func UpdateIssue(id string, updatedIssue structs.Issue) (err error) {

	query := `UPDATE issues SET title = $1, description = $2, assigned_to = $3,
			  created_by = $4, status = $5 WHERE id = $6`

	_, err = DB.Exec(query, updatedIssue.Title, updatedIssue.Description,
		updatedIssue.AssignedTo, updatedIssue.CreatedBy, updatedIssue.Status, id)

	return
}

func DeleteIssue(id string) (err error) {

	query := `DELETE FROM issues WHERE id = $1`

	_, err = DB.Exec(query, id)

	return
}

func CheckUserExists(token string) (exists bool, err error) {

	query := `SELECT EXISTS(SELECT 1 FROM users WHERE access_token = $1)`

	err = DB.QueryRow(query, token).Scan(&exists)

	return
}

func GetUserFromAccessToken(token string) (user structs.User, err error) {

	query := `SELECT username, email, first_name, last_name FROM users WHERE access_token=$1`

	err = DB.QueryRow(query, token).Scan(&user.Username, &user.Email,
		&user.FirstName, &user.LastName)

	return
}

func GetUserFromUsername(username string) (user structs.User, err error) {

	query := `SELECT username, email, first_name, last_name FROM users WHERE username=$1`

	err = DB.QueryRow(query, username).Scan(&user.Username, &user.Email,
		&user.FirstName, &user.LastName)

	return
}

func GetAllUsers() (users []structs.User, err error) {

	var rows *sql.Rows

	query := `SELECT username, email, first_name, last_name FROM users`

	rows, err = DB.Query(query)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {

		var user structs.User

		if err = rows.Scan(&user.Username, &user.Email, &user.FirstName,
			&user.LastName); err != nil {
			return
		}

		users = append(users, user)
	}

	return
}
