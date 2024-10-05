package controllers

import (
	"encoding/json"
	"net/http"
	"simple-backend/services"
	"strconv"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	fromCustomerID := r.URL.Query().Get("from_customer_id")
	toCustomerID := r.URL.Query().Get("to_customer_id")
	amount, _ := strconv.ParseFloat(r.URL.Query().Get("amount"), 64)

	err := services.Payment(fromCustomerID, toCustomerID, amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Payment successful")
}
