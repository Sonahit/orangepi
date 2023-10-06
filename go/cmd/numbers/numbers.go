package main

import "display/lib"

type NumberSegments struct {
	pins DataPins
}

type DataPins struct {
	a, b, c, d, e, f, g int
}

/*
_a_
f b
_g_
e c
_d_
*/

func NewNumber(segments ...int) NumberSegments {
	return NumberSegments{
		pins: DataPins{
			a: segments[0],
			b: segments[1],
			c: segments[2],
			d: segments[3],
			e: segments[4],
			f: segments[5],
			g: segments[6],
		},
	}
}

func (n NumberSegments) Render(pins DataPins) {
	n.writePin(n.pins.a, pins.a)
	n.writePin(n.pins.b, pins.b)
	n.writePin(n.pins.c, pins.c)
	n.writePin(n.pins.d, pins.d)
	n.writePin(n.pins.e, pins.e)
	n.writePin(n.pins.f, pins.f)
	n.writePin(n.pins.g, pins.g)
}

func (n NumberSegments) writePin(value, pin int) {
	if value == 0 {
		lib.DigitalWrite(pin, lib.DIGITAL_LOW)
	} else {
		lib.DigitalWrite(pin, lib.DIGITAL_HIGH)
	}
}

var One = NewNumber(0, 1, 1, 0, 0, 0, 0)
var Two = NewNumber(1, 1, 0, 1, 1, 0, 1)

// var Three = NewNumber(0, 1, 1)
// var Four = NewNumber(0, 1, 1)
// var Five = NewNumber(0, 1, 1)
// var Six = NewNumber(0, 1, 1)
// var Seven = NewNumber(0, 1, 1)
// var Eight = NewNumber(0, 1, 1)
// var Nine = NewNumber(0, 1, 1)
// var Zero = NewNumber(0, 1, 1)
