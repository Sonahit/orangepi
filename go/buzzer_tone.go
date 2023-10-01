package main

type BuzzerPwm struct {
	pin       int
	CurrentHz int
}

func NewBuzzerPwm(pin int) *BuzzerPwm {
	ToneCreate(pin)
	return &BuzzerPwm{
		pin: pin,
	}
}

func (buz *BuzzerPwm) ChangeHz(hz int) {
	buz.CurrentHz = hz
	ToneWrite(buz.pin, buz.CurrentHz)
}

func (buz *BuzzerPwm) Stop() {
	ToneStop(buz.pin)
}
