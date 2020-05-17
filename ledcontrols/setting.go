package ledcontrols

import (
	"fmt"
)

func SetLEDColours(colour string){
	fmt.Println("Colour from the API is: "+colour)
}

func SetLEDPreset(preset string){
	fmt.Println("Preset from the API is: "+preset)
}