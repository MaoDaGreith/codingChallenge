package repository

import (
	"codingChallenge/model"
	"encoding/json"
	"fmt"
	"math"
	"sort"
)

func SearchID(request model.UserIDGetter) (model.User, error) {
	var users map[int]model.User
	var user model.User

	content, err := dataReader("./repository/data/users.json")
	if err != nil {
		return user, err
	}

	var data []model.User
	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return user, err
	}

	var userRecords = make(map[int]model.User)
	for _, val := range data {
		userRecords[val.ID] = val
	}

	return users[request.ID], err
}

func ActionsCount(request model.ActionCountGetter) (model.Count, error) {
	var count model.Count

	content, err := dataReader("./repository/data/actions.json")
	if err != nil {
		return count, err
	}

	var data []model.Action
	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return count, err
	}

	for _, val := range data {
		if val.UserID == request.ID {
			count.Count++
		}
	}
	return count, nil
}

func CountTransitions() map[string]map[string]int {
	transitionMap := make(map[string]map[string]int)
	groupedActions, err := groupSortActions()
	if err != nil {
		fmt.Printf("Error grouping and sorting actions: %v\n", err)
		return transitionMap
	}

	for _, group := range groupedActions {
		actions := group.Actions
		for i := 1; i < len(actions); i++ {
			current := actions[i-1].Type
			next := actions[i].Type

			if _, exists := transitionMap[current]; !exists {
				transitionMap[current] = make(map[string]int)
			}
			transitionMap[current][next]++
		}
	}

	return transitionMap
}

func CalculatePercentages() map[string]map[string]float64 {
	percentages := make(map[string]map[string]float64)
	transitionMap := CountTransitions()

	for current, transitions := range transitionMap {
		total := 0
		for _, count := range transitions {
			total += count
		}

		percentages[current] = make(map[string]float64)
		for next, count := range transitions {
			percentages[current][next] = math.Round((float64(count)/float64(total))*100*100) / 100
		}
	}

	return percentages
}

func groupSortActions() ([]model.GroupedActions, error) {
	var sortedActions []model.GroupedActions
	content, err := dataReader("./repository/data/actions.json")
	if err != nil {
		return sortedActions, err
	}

	var actions []model.Action
	err = json.Unmarshal(content, &actions)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return sortedActions, err
	}

	// group by user id
	idGrouped := make(map[int][]model.Action)
	for _, action := range actions {
		idGrouped[action.UserID] = append(idGrouped[action.UserID], action)
	}

	// sort actions of each user by CreatedAt
	for userID, userActions := range idGrouped {
		sort.Slice(userActions, func(i, j int) bool {
			return userActions[i].CreatedAt.Before(userActions[j].CreatedAt)
		})
		sortedActions = append(sortedActions, model.GroupedActions{
			UserID:  userID,
			Actions: userActions,
		})
	}

	return sortedActions, nil
}
