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

func NewNumber(segments DataPins) NumberSegments {
	return NumberSegments{
		pins: DataPins{
			a: segments.a,
			b: segments.b,
			c: segments.c,
			d: segments.d,
			e: segments.e,
			f: segments.f,
			g: segments.g,
		},
	}
}

func (n NumberSegments) Render(pins DataPins) {
	n.writePin(pins.a, n.pins.a)
	n.writePin(pins.b, n.pins.b)
	n.writePin(pins.c, n.pins.c)
	n.writePin(pins.d, n.pins.d)
	n.writePin(pins.e, n.pins.e)
	n.writePin(pins.f, n.pins.f)
	n.writePin(pins.g, n.pins.g)
}

func (n NumberSegments) writePin(pin, value int) {
	if value == 0 {
		lib.DigitalWrite(pin, lib.DIGITAL_HIGH)
	} else {
		lib.DigitalWrite(pin, lib.DIGITAL_LOW)
	}
}

var One = NewNumber(DataPins{b: 1, c: 1})
var Two = NewNumber(DataPins{a: 1, b: 1, d: 1, e: 1, g: 1})
var Three = NewNumber(DataPins{a: 1, b: 1, g: 1, c: 1, d: 1})
var Four = NewNumber(DataPins{f: 1, g: 1, b: 1, c: 1})
var Five = NewNumber(DataPins{a: 1, f: 1, g: 1, c: 1, d: 1})
var Six = NewNumber(DataPins{a: 1, f: 1, g: 1, e: 1, c: 1, d: 1})
var Seven = NewNumber(DataPins{a: 1, b: 1, c: 1})
var Eight = NewNumber(DataPins{a: 1, b: 1, c: 1, d: 1, e: 1, g: 1, f: 1})
var Nine = NewNumber(DataPins{a: 1, b: 1, c: 1, d: 1, g: 1, f: 1})
var Zero = NewNumber(DataPins{a: 1, b: 1, c: 1, d: 1, e: 1, f: 1})
