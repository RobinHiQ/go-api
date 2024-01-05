package controller

import (
	"github.com/RobinHiQ/go-api/functions"
    "net/http"
)

// @Summary Get job description
// @Description Generates a job description for the given job title
// @Param jobTitle query string true "The title of the job"
// @Success 200 {string} string "Successful operation"
// @Router /job-description [get]

func GetJobDescription(w http.ResponseWriter, r *http.Request) {
	jobTitle := r.URL.Query().Get("jobTitle")
	if jobTitle == "" {
		http.Error(w, "Missing jobTitle parameter", http.StatusBadRequest)
		return
	}

	description, err := functions
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(description))
}