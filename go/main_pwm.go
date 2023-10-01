package main

import "log"

const (
	PIN_WPI  = 2
	RANGE_HZ = 20000
)

func mainPwm() {
	if err := WiringSetup(); err != 0 {
		log.Fatalf("Setup error %d", err)
	}

	buzzer := NewBuzzerPwm(PIN_WPI, RANGE_HZ)

	buzzer.ChangeHz(20)

	for {
	}
}
