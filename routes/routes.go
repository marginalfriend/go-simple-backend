package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"simple-backend/controllers"
)

func RegisterRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
	router.HandleFunc("/logout", controllers.LogoutHandler).Methods("POST")
	router.HandleFunc("/payment", controllers.PaymentHandler).Methods("POST")

	http.ListenAndServe(":8080", router)
}
