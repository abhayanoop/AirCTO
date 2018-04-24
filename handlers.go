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

	Issues[issue.ID] = issue

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

		fmt.Fprintf(w, "%v", issueJSON)

	} else {

		http.Error(w, "Issue not found!", http.StatusNotFound)
	}

	return
}

func updateIssueHandler(w http.ResponseWriter, r *http.Request) {

	issueID := r.URL.Query().Get("id")

	if currentIssue, ok := Issues[issueID]; ok {

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

		if updatedIssue.AssignedTo.Username != currentIssue.AssignedTo.Username {

			go sendUpdatedAssigneeEmail(updatedIssue.AssignedTo, updatedIssue)
		}

		Issues[currentIssue.ID] = updatedIssue

	} else {

		http.Error(w, "Issue not found!", http.StatusNotFound)
	}

	return
}

func deleteIssueHandler(w http.ResponseWriter, r *http.Request) {

	issueID := r.URL.Query().Get("id")

	if _, ok := Issues[issueID]; ok {

		delete(Issues, issueID)

	} else {

		http.Error(w, "Issue not found!", http.StatusNotFound)
	}

	return
}
