package main

import (
	"AirCTO/dbfuncs"
	"fmt"
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
	http.HandleFunc("/issue/list", authMiddleware(getAllIssuesHandler, http.MethodGet))

	fmt.Println("Listening to localhost:8080")

	http.ListenAndServe(":8080", nil)

}

func authMiddleware(
	fn func(http.ResponseWriter, *http.Request), method string,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		if exists, err := dbfuncs.CheckUserExists(token); err != nil || !exists {
			if err == nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			http.Error(w, err.Error(), http.StatusUnauthorized)

		} else {

			if method == r.Method {

				// Actual Handler Function
				fn(w, r)

			} else {

				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		}
	}
}
