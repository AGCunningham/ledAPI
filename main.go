package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stianeikeland/go-rpio/v4"

	server "ledAPI/webserver"
)

func main() {
	err := rpio.Open()
	defer rpio.Close()
	if err != nil {
		log.Fatal("GPIO Pins could not be initialised")
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", server.HomeLink)
	router.HandleFunc("/colour/{colour}", server.GetColour).Methods("GET")
	router.HandleFunc("/preset/{preset}", server.SetPreset).Methods("GET")
	router.HandleFunc("/pin/{set}", server.SetPin).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", router))
}