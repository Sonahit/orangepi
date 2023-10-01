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
	scale := [8]int{262, 294, 330, 349, 392, 440, 494, 525}
	for {
		for _, s := range scale {
			log.Printf("%d\n", s)
			buzzer.ChangeHz(s)
			sleep(500)
		}
	}
}
