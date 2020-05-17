package webserver

import (
	"fmt"
	"net/http"
	
	"github.com/gorilla/mux"

	led "ledAPI/ledcontrols"
)

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hit /colour/{colour} to set an LED colour\n")
	fmt.Fprintf(w, "Hit /preset/{preset} to set an LED preset\n")
}

func GetColour(w http.ResponseWriter, r *http.Request){
	colourReq := mux.Vars(r)["colour"]

	led.SetLEDColours(colourReq)

	fmt.Fprintf(w, "Colour is: "+colourReq)
}

func SetPreset(w http.ResponseWriter, r *http.Request){
	presetReq := mux.Vars(r)["preset"]

	led.SetLEDPreset(presetReq)

	fmt.Fprintf(w, "Preset is: "+presetReq)
}

func SetPin(w http.ResponseWriter, r *http.Request){
	pinSet := mux.Vars(r)["set"]

	led.SetPinOut(pinSet)

	fmt.Fprintf(w, "Pin set to: "+pinSet)
}