package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"simple-backend/models"
	"time"
)

var dataFile = "repository/data.json"

func Login(customerID string) error {
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return err
	}

	var database map[string][]models.Customer
	err = json.Unmarshal(data, &database)
	if err != nil {
		return err
	}

	customers := database["customers"]
	for i, customer := range customers {
		if customer.ID == customerID {
			if customer.LoggedIn {
				return errors.New("customer already logged in")
			}

			customers[i].LoggedIn = true
			logAction(customerID, "login")
			return saveData(database)
		}
	}

	return errors.New("customer not found")
}

func Logout(customerID string) error {
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return err
	}

	var database map[string][]models.Customer
	err = json.Unmarshal(data, &database)
	if err != nil {
		return err
	}

	customers := database["customers"]
	for i, customer := range customers {
		if customer.ID == customerID {
			if !customer.LoggedIn {
				return errors.New("customer is not logged in")
			}

			customers[i].LoggedIn = false
			logAction(customerID, "logout")
			return saveData(database)
		}
	}

	return errors.New("customer not found")
}

func saveData(data map[string][]models.Customer) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dataFile, jsonData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func logAction(customerID, action string) {
	history := models.History{
		CustomerID: customerID,
		Action:     action,
		Timestamp:  time.Now().Format(time.RFC3339),
	}

	logData, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return
	}

	var database map[string]interface{}
	json.Unmarshal(logData, &database)

	historyList := database["history"].([]interface{})
	historyList = append(historyList, history)
	database["history"] = historyList

	newData, _ := json.MarshalIndent(database, "", "  ")
	ioutil.WriteFile(dataFile, newData, os.ModePerm)
}
