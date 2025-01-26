package controllers

import (
	"codingChallenge/model"
	"codingChallenge/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func getNextActionsGlobal(w http.ResponseWriter, r *http.Request) {
	var percent map[string]map[string]float64
	var actionType model.ActionType
	actionType.Type = r.URL.Query().Get("type")
	if actionType == "" {
		http.Error(w, "Missing 'type' query parameter", http.StatusBadRequest)
		return
	}

	percent = repository.CalculatePercentages()
	actionType.Type = strings.ToUpper(actionType.Type)
	percentages, exists := percent[actionType.Type]
	if !exists {
		http.Error(w, fmt.Sprintf("No data available for action type: %s", actionType), http.StatusNotFound)
		return
	}

	// Respond with percentages in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(percentages)
}
