package controllers

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/get/", getUserID)
	mux.HandleFunc("/count/", getUserActionsTotal)
	mux.HandleFunc("/nextAction", getNextActionsGlobal)

	return mux
}
