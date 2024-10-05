package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"simple-backend/models"
)

func Payment(fromCustomerID, toCustomerID string, amount float64) error {
	if amount <= 0 {
		return errors.New("invalid amount")
	}

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
	var fromCustomer *models.Customer
	var toCustomer *models.Customer

	for i, customer := range customers {
		if customer.ID == fromCustomerID {
			fromCustomer = &customers[i]
		}
		if customer.ID == toCustomerID {
			toCustomer = &customers[i]
		}
	}

	if fromCustomer == nil || toCustomer == nil {
		return errors.New("customer not found")
	}

	if !fromCustomer.LoggedIn {
		return errors.New("customer not logged in")
	}

	logAction(fromCustomerID, "payment to "+toCustomerID)
	return nil
}
