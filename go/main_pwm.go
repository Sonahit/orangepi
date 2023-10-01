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

	PinMode(PIN_WPI, 5)

	buzzer := NewBuzzerPwm(PIN_WPI, RANGE_HZ)
	buzzer.ChangeHz(1000)

	for {
	}
}
