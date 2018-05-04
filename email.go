package main

import (
	"fmt"
	"net/smtp"
	"strings"
	"time"

	"AirCTO/dbfuncs"
	"AirCTO/structs"
)

func sendUpdatedAssigneeEmail(recipientUsername string, issue structs.Issue) {

	recipient, err := dbfuncs.GetUserFromUsername(recipientUsername)
	if err != nil {
		fmt.Errorf("Error occured while getting user for sending email - " + err.Error())
		return
	}

	// Channel waits for 12 mins after issue is updated, and then sends email
	timer := time.NewTimer(12 * time.Minute)
	<-timer.C

	body := "To: " + recipient.Email + "\r\n" +
		"Subject: New issue assigned to you\r\n" +
		"\r\n" +
		"Hi " + recipient.FirstName + ", \n\nIssue " +
		issue.ID + " - " + issue.Title + " has been assigned to you.\r\n"

	err = smtp.SendMail(
		structs.SMTPServerPort,
		structs.SMTPAuth,
		structs.SMTPFrom,
		[]string{recipient.Email},
		[]byte(body),
	)

	if err != nil {
		fmt.Errorf("Error occured while sending email - " + err.Error())
	}
}

func sendPeriodicEmailsForOpenIssues(timePeriod time.Duration) {

	ticker := time.NewTicker(timePeriod)

	users, err := dbfuncs.GetAllUsers()
	if err != nil {
		fmt.Errorf("Error occured while getting users for sending email - " + err.Error())
		return
	}

	issues, err := dbfuncs.GetAllIssues()
	if err != nil {
		fmt.Errorf("Error occured while getting issues for sending email - " + err.Error())
		return
	}

	// Start ticker channel
	for _ = range ticker.C {

		for _, user := range users {

			usersOpenIssues := []string{}

			// Get all issues assigned to user that are open
			for _, issue := range issues {

				if issue.AssignedTo == user.Username {

					if issue.Status == "Open" {
						usersOpenIssues = append(usersOpenIssues, issue.ID+" - "+issue.Title)
					}
				}
			}

			body := "To: " + user.Email + "\r\n" +
				"Subject: Open Issues assigned to you\r\n" +
				"\r\n" +
				"Hi " + user.FirstName +
				", \n\n Open Issues assigned to you: \n" +
				strings.Join(usersOpenIssues, ",") + "\r\n"

			err := smtp.SendMail(
				structs.SMTPServerPort,
				structs.SMTPAuth,
				structs.SMTPFrom,
				[]string{user.Email},
				[]byte(body),
			)

			if err != nil {
				fmt.Errorf("Error occured while sending email - " + err.Error())
			}
		}
	}
}
