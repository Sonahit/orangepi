package lib

// #cgo LDFLAGS: -lwiringPi -lpthread
// #include "tone.hpp"
import "C"

func ToneCreate(pin int) {
	C.toneCreate(C.int(pin))
}

func ToneWrite(pin, value int) {
	C.toneWrite(C.int(pin), C.int(value))
}

func ToneStop(pin int) {
	C.toneStop(C.int(pin))
}
