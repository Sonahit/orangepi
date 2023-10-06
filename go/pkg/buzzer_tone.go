package pkg

import "display/lib"

type BuzzerTone struct {
	Pin       int
	CurrentHz int
}

func NewBuzzer(pin int) *BuzzerTone {
	lib.ToneCreate(pin)
	return &BuzzerTone{
		Pin: pin,
	}
}

func (buz *BuzzerTone) ChangeHz(hz int) {
	buz.CurrentHz = hz
	lib.ToneWrite(buz.Pin, buz.CurrentHz)
}

func (buz *BuzzerTone) Stop() {
	lib.ToneStop(buz.Pin)
}
