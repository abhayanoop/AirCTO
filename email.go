package main

import (
	"fmt"
	"net/smtp"
	"strings"
	"time"
)

func sendUpdatedAssigneeEmail(recipient User, issue Issue) {

	// Channel waits for 12 mins after issue is updated, and then sends email
	timer := time.NewTimer(12 * time.Minute)
	<-timer.C

	body := "To: " + recipient.Email + "\r\n" +
		"Subject: New issue assigned to you\r\n" +
		"\r\n" +
		"Hi " + recipient.FirstName + ", \n\nIssue " +
		issue.ID + " - " + issue.Title + " has been assigned to you.\r\n"

	err := smtp.SendMail(
		smtpServerPort,
		smtpAuth,
		smtpFrom,
		[]string{recipient.Email},
		[]byte(body),
	)

	if err != nil {
		fmt.Errorf("Error occured while sending email - " + err.Error())
	}
}

func sendPeriodicEmailsForOpenIssues(timePeriod time.Duration) {

	ticker := time.NewTicker(timePeriod)

	// Start ticker channel
	for _ = range ticker.C {

		for _, user := range Users {

			usersOpenIssues := []string{}

			// Get all issues assigned to user that are open
			for _, issue := range Issues {

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
				smtpServerPort,
				smtpAuth,
				smtpFrom,
				[]string{user.Email},
				[]byte(body),
			)

			if err != nil {
				fmt.Errorf("Error occured while sending email - " + err.Error())
			}
		}
	}
}
