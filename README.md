# ledAPI

This has been designed with the goal of running the webserver on a Raspberry Pi and using a RESTful API structure to control addressable LED strip(s).

## Setup

* A simple `go run ledAPI.go` is sufficient at the current time to run this locally or on a Pi

### Binary Run

* If you need to rebuild the binary `go build` will build and package the binary under the name `ledAPI`
* This can then be run with `./ledAPI` - or using `nohup ./ledAPI &` if you want to run it in the background on a Pi
    * To find the PID you can run `ps -ef | grep led` and then `kill -9 <pid>` to stop the process

## Endpoints

* `<staticIP>:10000/colour/{colour}` will be used to set all LEDs to a chosen colour - this will initially be from a preset list
* `<staticIP>:10000/preset/{preset}` will be used to chose a preset for the strip i.e. transition through all colours
* `<staticIP>:10000/pcon` will turn on a PC that has the MAC address listed in to the code providing it's on the same network

## Useful links
* https://www.unix.com/shell-programming-and-scripting/195303-how-stop-nohup-working-background.html
* https://sabhiram.com/development/2015/02/16/sending_wol_packets_with_golang.html
* https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da
