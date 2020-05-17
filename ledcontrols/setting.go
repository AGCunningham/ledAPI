package ledcontrols

import (
	"fmt"

	"github.com/stianeikeland/go-rpio/v4"
)

func SetLEDColours(colour string){
	fmt.Println("Colour from the API is: "+colour)
}

func SetLEDPreset(preset string){
	fmt.Println("Preset from the API is: "+preset)
}

func SetPinOut(set string){
	// GPIO Pin 19 - https://github.com/stianeikeland/go-rpio#usage
	dataPin := rpio.Pin(10)

	if set == "high" {
		dataPin.High()
	} else {
		dataPin.Low()
	}
}