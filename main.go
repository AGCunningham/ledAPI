package main

import (
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
	server "ledAPI/webserver"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", server.HomeLink)
	router.HandleFunc("/colour/{colour}", server.GetColour).Methods("GET")
	router.HandleFunc("/preset/{preset}", server.SetPreset).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", router))
}