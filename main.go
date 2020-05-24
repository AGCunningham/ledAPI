package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/stianeikeland/go-rpio/v4"

	// Initial Testing of LEDs
	github.com/mcuadros/go-rpi-ws281x

	server "ledAPI/webserver"
)

func main() {
	// err := rpio.Open()
	// defer rpio.Close()
	// if err != nil {
	// 	log.Fatal("GPIO Pins could not be initialised")
	// }

	// create a new canvas with the given width and height, and the config, in this
	// case the configuration is for a Unicorn pHAT (8x4 pixels matrix) with the default config
	
	// default configuration
	// Pin:        18,
	// Frequency:  800000, // 800khz
	// DMA:        10,
	// Brightness: 30,
	// StripType:  StripGRB,

	c, _ := ws281x.NewCanvas(150, 1, &ws281x.DefaultConfig)

	// initialize the canvas and the matrix
	c.Initialize()
	defer c.Close()

	// now we copy a white image into the ws281x.Canvas, this turn on all the leds to white
	draw.Draw(c, c.Bounds(), image.NewUniform(color.White), image.ZP, draw.Over)
	
	// Render required to turn the LEDs on
	c.Render()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", server.HomeLink)
	router.HandleFunc("/colour/{colour}", server.GetColour).Methods("GET")
	router.HandleFunc("/preset/{preset}", server.SetPreset).Methods("GET")
	router.HandleFunc("/pin/{set}", server.SetPin).Methods("GET")
	router.HandleFunc("/pcon", server.PowerOn).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", router))
}