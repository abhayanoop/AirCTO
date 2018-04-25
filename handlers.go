package main

import (
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

	var issue Issue

	if err = json.Unmarshal(b, &issue); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := Issues[issue.ID]; !ok {
		Issues[issue.ID] = issue
		fmt.Fprintf(w, "%v", "Issue created successfully")
	} else {
		http.Error(w, "Issue ID already exists", http.StatusInternalServerError)
	}

	return
}

func getIssueHandler(w http.ResponseWriter, r *http.Request) {

	issueID := r.URL.Query().Get("id")

	if issue, ok := Issues[issueID]; ok {

		issueJSON, err := json.Marshal(issue)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%v", string(issueJSON))

	} else {

		http.Error(w, "Issue not found!", http.StatusNotFound)
	}

	return
}

func getAllIssuesHandler(w http.ResponseWriter, r *http.Request) {

	issuesJSON, err := json.Marshal(Issues)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%v", string(issuesJSON))

	return
}

func updateIssueHandler(w http.ResponseWriter, r *http.Request) {

	issueID := r.URL.Query().Get("id")
	username := Users[r.Header.Get("Authorization")].Username

	if currentIssue, ok := Issues[issueID]; ok {

		if currentIssue.CreatedBy != username {
			http.Error(w, "User is not authorized to perform this function", http.StatusUnauthorized)
			return
		}

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var updatedIssue Issue

		if err = json.Unmarshal(b, &updatedIssue); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if updatedIssue.AssignedTo != currentIssue.AssignedTo {

			go sendUpdatedAssigneeEmail(Users[updatedIssue.AssignedTo], updatedIssue)
		}

		Issues[currentIssue.ID] = updatedIssue

		fmt.Fprintf(w, "%v", "Issue updated successfully")

	} else {

		http.Error(w, "Issue not found!", http.StatusNotFound)
	}

	return
}

func deleteIssueHandler(w http.ResponseWriter, r *http.Request) {

	issueID := r.URL.Query().Get("id")
	username := Users[r.Header.Get("Authorization")].Username

	if issue, ok := Issues[issueID]; ok {

		if issue.CreatedBy != username {
			http.Error(w, "User is not authorized to perform this function", http.StatusUnauthorized)
			return
		}

		delete(Issues, issueID)
		fmt.Fprintf(w, "%v", "Issue deleted successfully")

	} else {

		http.Error(w, "Issue not found!", http.StatusNotFound)
	}

	return
}
