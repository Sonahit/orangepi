package main

import (
	"display/lib"
	"display/pkg"
	"flag"
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

var DATA_PINS = DataPins{
	a: A_PIN,
	b: B_PIN,
	c: C_PIN,
	d: D_PIN,
	e: E_PIN,
	f: F_PIN,
	g: G_PIN,
}

func setupNumber() {

	if err := lib.WiringSetup(); err != 0 {
		log.Fatalf("Setup error %d", err)
	}
	lib.PinModeDefault(VCC_1, lib.PIN_OUTPUT, lib.DIGITAL_HIGH)
	lib.PinModeDefault(VCC_2, lib.PIN_OUTPUT, lib.DIGITAL_HIGH)
	lib.PinModeDefault(DP_PIN, lib.PIN_OUTPUT, lib.DIGITAL_HIGH)

	lib.PinModeDefault(DATA_PINS.a, lib.PIN_OUTPUT, lib.DIGITAL_LOW)
	lib.PinModeDefault(DATA_PINS.b, lib.PIN_OUTPUT, lib.DIGITAL_LOW)
	lib.PinModeDefault(DATA_PINS.f, lib.PIN_OUTPUT, lib.DIGITAL_LOW)
	lib.PinModeDefault(DATA_PINS.g, lib.PIN_OUTPUT, lib.DIGITAL_LOW)
	lib.PinModeDefault(DATA_PINS.c, lib.PIN_OUTPUT, lib.DIGITAL_LOW)
	lib.PinModeDefault(DATA_PINS.d, lib.PIN_OUTPUT, lib.DIGITAL_LOW)
	lib.PinModeDefault(DATA_PINS.e, lib.PIN_OUTPUT, lib.DIGITAL_LOW)
}

func logicNumber() {
	var sleepMs int

	flag.IntVar(&sleepMs, "sleep", 500, "")
	flag.Parse()

	if sleepMs <= 0 {
		sleepMs = 500
	}

	for {
		println("LOGIC START")
		lib.DigitalWrite(DP_PIN, lib.DIGITAL_HIGH)

		One.Render(DATA_PINS)
		pkg.Sleep(sleepMs)

		Two.Render(DATA_PINS)
		pkg.Sleep(sleepMs)

		Three.Render(DATA_PINS)
		pkg.Sleep(sleepMs)

		Four.Render(DATA_PINS)
		pkg.Sleep(sleepMs)

		Five.Render(DATA_PINS)
		pkg.Sleep(sleepMs)

		Six.Render(DATA_PINS)
		pkg.Sleep(sleepMs)

		Seven.Render(DATA_PINS)
		pkg.Sleep(sleepMs)

		Eight.Render(DATA_PINS)
		pkg.Sleep(sleepMs)

		Nine.Render(DATA_PINS)
		pkg.Sleep(sleepMs)

		Zero.Render(DATA_PINS)
		pkg.Sleep(sleepMs)

		lib.DigitalWrite(DP_PIN, lib.DIGITAL_LOW)
		pkg.Sleep(500)

		println("LOGIC END")

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
