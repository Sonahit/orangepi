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

	rng := 50
	speed := rng

	if err := lib.WiringSetup(); err != 0 {
		log.Fatalf("Setup error %d", err)
	}
	defer lib.PwmStop(pin)
	lib.PinMode(pin, lib.PIN_OUTPUT)
	lib.PwmCreate(pin, speed, rng)

	for {
		lib.PwmWrite(pin, speed)
		pkg.Sleep(2000)

		if speed < 0 {
			speed = 0
		} else {
			speed -= 10
		}
	}
}
