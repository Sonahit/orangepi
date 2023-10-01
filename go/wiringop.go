package main

// #cgo LDFLAGS: -lwiringPi -lpthread
// #include "lib/wiringop.hpp"
import "C"

func WiringSetup() int {
	return int(C.setup())
}
