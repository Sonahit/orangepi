package lib

// #cgo LDFLAGS: -lwiringPi -lpthread
// #include "pwm.hpp"
import "C"

func PwmCreate(pin, value, rang int) {
	C.pwmCreate(C.int(pin), C.int(value), C.int(rang))
}

func PwmWrite(pin, value int) {
	C.pwmWrite(C.int(pin), C.int(value))
}

func PwmStop(pin int) {
	C.pwmStop(C.int(pin))
}
