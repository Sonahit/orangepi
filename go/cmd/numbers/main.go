package main

import (
	"display/lib"
	"display/pkg"
	"log"
)

const (
	PIN_0 = 25
	PIN_1 = 23
	PIN_2 = 22
	PIN_3 = 20
	PIN_4 = 19
	PIN_5 = 27
	PIN_6 = 26
	PIN_7 = 24
	PIN_8 = 21
	PIN_9 = 18
)

const (
	B_PIN     = PIN_0
	A_PIN     = PIN_1
	DP_PIN    = PIN_2
	F_PIN     = PIN_3
	G_PIN     = PIN_4
	GREEN_PIN = PIN_5
	C_PIN     = PIN_6
	D_PIN     = PIN_7
	E_PIN     = PIN_8
	RED_PIN   = PIN_9
)

var DATA_PINS = [8]int{
	B_PIN,
	A_PIN,
	DP_PIN,
	F_PIN,
	G_PIN,
	C_PIN,
	D_PIN,
	E_PIN,
}

func setupNumber() {

	if err := lib.WiringSetup(); err != 0 {
		log.Fatalf("Setup error %d", err)
	}

	lib.PinMode(B_PIN, lib.PIN_OUTPUT)
	lib.PinMode(A_PIN, lib.PIN_OUTPUT)
	lib.PinMode(DP_PIN, lib.PIN_OUTPUT)
	lib.PinMode(F_PIN, lib.PIN_OUTPUT)
	lib.PinMode(G_PIN, lib.PIN_OUTPUT)
	lib.PinMode(GREEN_PIN, lib.PIN_OUTPUT)
	lib.PinMode(C_PIN, lib.PIN_OUTPUT)
	lib.PinMode(D_PIN, lib.PIN_OUTPUT)
	lib.PinMode(E_PIN, lib.PIN_OUTPUT)
	lib.PinMode(RED_PIN, lib.PIN_OUTPUT)
}

func logicNumber() {
	lib.DigitalWrite(RED_PIN, lib.DIGITAL_HIGH)
	for {
		for _, pin := range DATA_PINS {
			toggle(pin)
			pkg.Sleep(200)
		}
	}
}

var toggled map[int]bool = make(map[int]bool)

func toggle(pin int) {
	if toggled[pin] {
		lib.DigitalWrite(pin, lib.DIGITAL_LOW)
	} else {
		lib.DigitalWrite(pin, lib.DIGITAL_HIGH)
	}
	toggled[pin] = !toggled[pin]
}

func main() {
	setupNumber()
	logicNumber()
}
