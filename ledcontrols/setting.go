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
	// Pin out -> https://github.com/stianeikeland/go-rpio/blob/master/rpio.go
	dataPin := rpio.Pin(26)
	dataPin.Output()
	dataPin.High()

	fmt.Println("Pin has been set: "+set)

	if set == "high" {
		dataPin.High()
	} else if set == "low" {
		dataPin.Low()
	}
}