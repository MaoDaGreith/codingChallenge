package controllers

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/getUser/", getUserID)
	mux.HandleFunc("/countActions/", getUserActionsTotal)
	mux.HandleFunc("/nextAction", getNextActionsGlobal)
	mux.HandleFunc("/referralIndex", getReferralIndex)

	return mux
}
