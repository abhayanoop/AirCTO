package main

import (
	"net/http"
	"time"
)

var cronJobFrequency = 24 * time.Hour

func init() {

	go sendPeriodicEmailsForOpenIssues(cronJobFrequency)
}

func main() {

	http.HandleFunc("/issue/read", authMiddleware(getIssueHandler, http.MethodGet))
	http.HandleFunc("/issue/create", authMiddleware(createIssueHandler, http.MethodPost))
	http.HandleFunc("/issue/update", authMiddleware(updateIssueHandler, http.MethodPut))
	http.HandleFunc("/issue/delete", authMiddleware(deleteIssueHandler, http.MethodDelete))

	http.ListenAndServe(":8080", nil)
}

func authMiddleware(
	fn func(http.ResponseWriter, *http.Request), method string,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if method == r.Method {

			fn(w, r)

		} else {

			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
