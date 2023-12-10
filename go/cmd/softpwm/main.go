package main

import (
	"display/lib"
	"display/pkg"
	"flag"
	"log"
)

func main() {
	var pin int
	flag.IntVar(&pin, "pin", 2, "pin pwm")
	flag.Parse()

	speed := 0
	rng := 50

	if err := lib.WiringSetup(); err != 0 {
		log.Fatalf("Setup error %d", err)
	}
	defer lib.PwmStop(pin)
	lib.PinMode(pin, lib.PIN_OUTPUT)
	lib.PwmCreate(pin, speed, rng)

	for {
		if speed >= rng {
			speed = 0
		} else {
			speed += 10
		}

		lib.PwmWrite(pin, speed)
		pkg.Sleep(5000)
	}
}
