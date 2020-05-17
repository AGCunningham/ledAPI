# ledAPI

This has been designed with the goal of running the webserver on a Raspberry Pi and using a RESTful API structure to control addressable LED strip(s).

## Setup

* A simple `go run ledAPI.go` is sufficient at the current time to run this locally or on a Pi
> In future this will be packaged into a binary that only need be run

## Endpoints

* `<staticIP>:10000/colour/{colour}` will be used to set all LEDs to a chosen colour - this will initially be from a preset list
* `<staticIP>:10000/preset/{preset}` will be used to chose a preset for the strip i.e. transition through all colours