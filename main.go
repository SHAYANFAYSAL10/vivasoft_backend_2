package main

import (
	"VIVASOFT2/src/apis/user_api"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/user", user_api.FindAll).Methods("GET")
	router.HandleFunc("/api/user/create", user_api.Create).Methods("POST")
	router.HandleFunc("/api/user/update", user_api.Update).Methods("PATCH")
	router.HandleFunc("/api/user/delete/{id}", user_api.Delete).Methods("DELETE")

	err := http.ListenAndServe(":4000", router)
	if err != nil {
		fmt.Println((err))
	}
}
