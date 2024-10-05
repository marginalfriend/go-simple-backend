package models

type Customer struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	MerchantID string `json:"merchant_id"`
	LoggedIn   bool   `json:"logged_in"`
}
