package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/pathparams/{param1}/{param2}", handlerWithPathParams).Methods("GET")
	router.HandleFunc("/queryparams", handlerwithQueryParams).Methods("POST")
	router.HandleFunc("/requestbody", handlerwithRequestBody).Methods("POST")

	http.ListenAndServe(":8080", router)
}

func handlerWithPathParams(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resp := make([]string, 0)
	for i, j := range vars {
		fmt.Println(i, j)
		resp = append(resp, "Got Mapping for "+i+" with value "+j)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func handlerwithQueryParams(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	resp := make([]string, 0)
	for i, j := range params {
		fmt.Println(i, j)
		resp = append(resp, "Got Mapping for "+i+" with value "+j[0])

	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func handlerwithRequestBody(w http.ResponseWriter, r *http.Request) {
	var data any
	json.NewDecoder(r.Body).Decode(&data)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
