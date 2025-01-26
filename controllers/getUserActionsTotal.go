package controllers

import (
	"codingChallenge/model"
	"codingChallenge/repository"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func getUserActionsTotal(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	segments := strings.Split(path, "/")
	var userID model.ActionCountGetter

	var err error

	// getting the id parameter from the URL
	if len(segments) >= 3 && segments[1] == "count" {
		userID.ID, err = strconv.Atoi(segments[2])
		if err != nil {
			http.Error(w, "Unable to convert ID from string to integer", http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "Invalid URL path", http.StatusNotFound)
		return
	}

	// calling the Search function for an id from the file
	data, err := repository.ActionsCount(userID)
	if err != nil {
		http.Error(w, "Error while searching user ID: "+err.Error(), http.StatusInternalServerError)
	}

	// encoding data to send to writer
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to encode data to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, "Failed to write: "+err.Error(), http.StatusInternalServerError)
	}
}
