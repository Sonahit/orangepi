package main

import (
	"display/lib"
	"display/pkg"
	"log"
)

// |   19 |  17 |    SDA.1 | ALT3 | 0 | 27 || 28 | 1 | OUT  | SCL.1    | 18  | 18   |
// |    7 |  19 |     PA07 |  OFF | 0 | 29 || 30 |   |      | GND      |     |      |
// |    8 |  20 |     PA08 |  OFF | 0 | 31 || 32 | 1 | OUT  | RTS.1    | 21  | 200  |
// |    9 |  22 |     PA09 |  OFF | 0 | 33 || 34 |   |      | GND      |     |      |
// |   10 |  23 |     PA10 |  OFF | 0 | 35 || 36 | 1 | OUT  | CTS.1    | 24  | 201  |
// |   20 |  25 |     PA20 |  OFF | 0 | 37 || 38 | 1 | OUT  | TXD.1    | 26  | 198  |
// |      |     |      GND |      |   | 39 || 40 | 0 | OUT  | RXD.1    | 27  | 199  |
// +------+-----+----------+------+---+----++----+---+------+----------+-----+------+

const (
	PIN_0 = 25 // 6 b
	PIN_1 = 23 // 7 a
	PIN_2 = 22 // 8 DP
	PIN_3 = 20 // 9 f
	PIN_4 = 19 // 10 g
	PIN_5 = 27 // 5 GREEN
	PIN_6 = 26 // 4 c
	PIN_7 = 24 // 3 d
	PIN_8 = 21 // 2 e
	PIN_9 = 18 // 1 RED
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
	lib.DigitalWrite(GREEN_PIN, lib.DIGITAL_HIGH)
	lib.DigitalWrite(DP_PIN, lib.DIGITAL_HIGH)

	lib.DigitalWrite(F_PIN, lib.DIGITAL_HIGH)

	pin(RED_PIN, "RED")
	pin(GREEN_PIN, "GREEN")
	pin(DP_PIN, "DP")

	lib.DigitalWrite(RED_PIN, lib.DIGITAL_HIGH)
	lib.DigitalWrite(GREEN_PIN, lib.DIGITAL_HIGH)
	lib.DigitalWrite(DP_PIN, lib.DIGITAL_HIGH)

	pin(B_PIN, "B_PIN")
	pin(A_PIN, "A_PIN")
	pin(F_PIN, "F_PIN")
	pin(G_PIN, "G_PIN")
	pin(C_PIN, "C_PIN")
	pin(D_PIN, "D_PIN")
	pin(E_PIN, "E_PIN")
}

func pin(pin int, msg string) {
	log.Printf("%s", msg)
	lib.DigitalWrite(pin, lib.DIGITAL_HIGH)

	pkg.Sleep(2000)

	lib.DigitalWrite(pin, lib.DIGITAL_LOW)
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
