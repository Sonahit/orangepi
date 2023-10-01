package main

// #cgo LDFLAGS: -lwiringPi -lpthread
// #include "lib/wiringop.hpp"
import "C"

const (
	DIGITAL_HIGH = 1
	DIGITAL_LOW  = 0
	PIN_INPUT    = 0
	PIN_OUTPUT   = 1
)

func WiringSetup() int {
	return int(C.setup())
}

func DigitalWrite(pin, value int) {
	C._digitalWrite(C.int(pin), C.int(value))
}

func PinMode(pin, mode int) {
	C._pinMode(C.int(pin), C.int(mode))
}
