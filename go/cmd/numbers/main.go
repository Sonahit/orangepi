package main

import (
	"display/lib"
	"display/pkg"
	"log"
)

const (
	PIN_0 = 37
	PIN_1 = 35
	PIN_2 = 33
	PIN_3 = 31
	PIN_4 = 29
	PIN_5 = 27
	PIN_6 = 26
	PIN_7 = 24
	PIN_8 = 21
	PIN_9 = 18
)

var B_PIN,
	A_PIN,
	DP_PIN,
	F_PIN,
	G_PIN,
	GREEN_PIN,
	C_PIN,
	D_PIN,
	E_PIN,
	RED_PIN = PIN_0,
	PIN_1,
	PIN_2,
	PIN_3,
	PIN_4,
	PIN_5,
	PIN_6,
	PIN_7,
	PIN_8,
	PIN_9

var DATA_PINS = []int{
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
	for {
		for _, pin := range DATA_PINS {
			toggle(pin)
			pkg.Sleep(200)
		}
	}
}

var toggled map[int]bool

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
