package main

import (
    "fmt"
    "log"
	"net/http"
	
	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hit /colour/{colour} to set an LED colour\n")
	fmt.Fprintf(w, "Hit /preset/{preset} to set an LED preset\n")
}

func getColour(w http.ResponseWriter, r *http.Request){
	colourReq := mux.Vars(r)["colour"]

	setLEDColours(colourReq)

	fmt.Fprintf(w, "Colour is: "+colourReq)
}

func setPreset(w http.ResponseWriter, r *http.Request){
	presetReq := mux.Vars(r)["preset"]

	setLEDPreset(presetReq)

	fmt.Fprintf(w, "Preset is: "+presetReq)
}

func setLEDColours(colour string){
	fmt.Println("Colour from the API is: "+colour)
}

func setLEDPreset(preset string){
	fmt.Println("Preset from the API is: "+preset)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/colour/{colour}", getColour).Methods("GET")
	router.HandleFunc("/preset/{preset}", setPreset).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", router))
}