package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	IP      string `json:"ip"`
	Message string `json:"message"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleGet).Methods("GET")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	srv.ListenAndServe()
}

func handleGet(w http.ResponseWriter, req *http.Request) {
	res := Response{IP: req.RemoteAddr, Message: "Hello From Go!"}
	json.NewEncoder(w).Encode(res)
}
