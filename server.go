package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func serve(port int) {
	http.HandleFunc("/services", getServicesHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func getServicesHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result.Endpoints)
}
