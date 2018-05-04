package main

import (
	"AirCTO/dbfuncs"
	"AirCTO/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func createIssueHandler(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var issue structs.Issue

	if err = json.Unmarshal(b, &issue); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = dbfuncs.CreateIssue(issue); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%v", "Issue created successfully")

	return
}

func getIssueHandler(w http.ResponseWriter, r *http.Request) {

	issueID := r.URL.Query().Get("id")

	issue, err := dbfuncs.GetIssue(issueID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	issueJSON, err := json.Marshal(issue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%v", string(issueJSON))

	return
}

func getAllIssuesHandler(w http.ResponseWriter, r *http.Request) {

	issues, err := dbfuncs.GetAllIssues()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	issuesJSON, err := json.Marshal(issues)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%v", string(issuesJSON))

	return
}

func updateIssueHandler(w http.ResponseWriter, r *http.Request) {

	issueID := r.URL.Query().Get("id")
	userAccessToken := r.Header.Get("Authorization")

	user, err := dbfuncs.GetUserFromAccessToken(userAccessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username := user.Username

	currentIssue, err := dbfuncs.GetIssue(issueID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if currentIssue.CreatedBy != username {
		http.Error(w, "User is not authorized to perform this function", http.StatusUnauthorized)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var updatedIssue structs.Issue

	if err = json.Unmarshal(b, &updatedIssue); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if updatedIssue.AssignedTo != currentIssue.AssignedTo {

		go sendUpdatedAssigneeEmail(updatedIssue.AssignedTo, updatedIssue)
	}

	if err = dbfuncs.UpdateIssue(issueID, updatedIssue); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%v", "Issue updated successfully")

	return
}

func deleteIssueHandler(w http.ResponseWriter, r *http.Request) {

	issueID := r.URL.Query().Get("id")
	userAccessToken := r.Header.Get("Authorization")

	user, err := dbfuncs.GetUserFromAccessToken(userAccessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username := user.Username

	issue, err := dbfuncs.GetIssue(issueID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if issue.CreatedBy != username {
		http.Error(w, "User is not authorized to perform this function", http.StatusUnauthorized)
		return
	}

	if err = dbfuncs.DeleteIssue(issueID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%v", "Issue deleted successfully")

	return
}
