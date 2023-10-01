package main

import "log"

const (
	PIN_WPI = 2
)

func mainTone() {
	if err := WiringSetup(); err != 0 {
		log.Fatalf("Setup error %d", err)
	}

	buzzer := NewBuzzerPwm(PIN_WPI)
	buzzer.ChangeHz(1000)

	for {
	}
}
