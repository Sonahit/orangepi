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
	B_PIN  = PIN_0
	A_PIN  = PIN_1
	DP_PIN = PIN_2
	F_PIN  = PIN_3
	G_PIN  = PIN_4
	VCC_1  = PIN_5
	C_PIN  = PIN_6
	D_PIN  = PIN_7
	E_PIN  = PIN_8
	VCC_2  = PIN_9
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
	lib.PinMode(VCC_1, lib.PIN_OUTPUT)
	lib.PinMode(C_PIN, lib.PIN_OUTPUT)
	lib.PinMode(D_PIN, lib.PIN_OUTPUT)
	lib.PinMode(E_PIN, lib.PIN_OUTPUT)
	lib.PinMode(VCC_2, lib.PIN_OUTPUT)
}

func logicNumber() {
	lib.DigitalWrite(VCC_1, lib.DIGITAL_HIGH)
	lib.DigitalWrite(VCC_2, lib.DIGITAL_HIGH)
	lib.DigitalWrite(DP_PIN, lib.DIGITAL_HIGH)

	for {
		for _, pin := range DATA_PINS {
			toggle(pin)
			pkg.Sleep(1000)
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
