package main

type BuzzerPwm struct {
	pin       int
	RangeHz   int
	CurrentHz int
}

func NewBuzzerPwm(pin, rangeHz int) *BuzzerPwm {
	PwmCreate(pin, 0, rangeHz)
	return &BuzzerPwm{
		pin:     pin,
		RangeHz: rangeHz,
	}
}

func (buz *BuzzerPwm) ChangeHz(hz int) {
	buz.CurrentHz = hz
	PwmWrite(buz.pin, buz.CurrentHz)
}

func (buz *BuzzerPwm) Stop() {
	PwmStop(buz.pin)
}
