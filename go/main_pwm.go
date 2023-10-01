package main

import "log"

const (
	PIN_WPI  = 2
	RANGE_HZ = 200000
)

func mainPwm() {
	if err := WiringSetup(); err != 0 {
		log.Fatalf("Setup error %d", err)
	}

	PinMode(PIN_WPI, PIN_OUTPUT)

	for {
		DigitalWrite(PIN_WPI, DIGITAL_LOW)
		sleep(2)
		DigitalWrite(PIN_WPI, DIGITAL_HIGH)
		sleep(2)
	}
}
