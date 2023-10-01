package main

import "time"

type I2CLed struct {
	fd          I2CFd
	Width       int
	row_offsets []int
}

const (
	LCD_CMD        = 0
	LCD_CHAR       = 1
	PULSE_SLEEP_NS = 500
)

const (
	LCD_ENABLE       = 0b100
	LCD_BACKLIGHT    = 0x08
	LCD_SETCGRAMADDR = 0x40
	LCD_SETDDRAMADDR = 0x80
)

const (
	LCD_MODE_EIGHT_BYTES = 0b10000
	LCD_MODE_FOUR_BYTES  = 0b00000

	LCD_MODE_ONE_LINES = 0b0000
	LCD_MODE_TWO_LINES = 0b1000
)

const (
	LCD_LINE_ONE = 0x80
	LCD_LINE_TWO = 0xc0
	NONE         = 0x00
)

func NewI2CLed(fd I2CFd) I2CLed {
	width := 16
	return I2CLed{
		fd:          fd,
		Width:       width,
		row_offsets: []int{0x00, 0x40, 0x00 + width, 0x40 + width},
	}
}

func (led I2CLed) lcdBytes(bits, mode int) {
	bits_high := mode | (bits & 0xf0) | LCD_BACKLIGHT
	bits_low := mode | ((bits << 4) & 0xf0) | LCD_BACKLIGHT

	led.fd.I2cWrite(bits_high)
	led.Enable(bits_high)

	led.fd.I2cWrite(bits_low)
	led.Enable(bits_low)

}

func (led I2CLed) Enable(bytes int) {
	waitPulse()
	led.fd.I2cWrite(bytes | LCD_ENABLE)
	waitPulse()
	led.fd.I2cWrite(bytes & (^LCD_ENABLE))
	waitPulse()
}

func (led I2CLed) Command(bytes int) {
	led.lcdBytes(bytes, LCD_CMD)
}
func (led I2CLed) Char(bytes int) {
	led.lcdBytes(bytes, LCD_CHAR)
}

func (led I2CLed) TextString(text string, line int) {
	lcdText[string](led, text, line)

}

func (led I2CLed) TextBytes(text, line int) {
	lcdText[int](led, text, line)

}

func (led I2CLed) Display(d, c, b int) {
	led.Command(0b00001000 | (d << 2) | (c << 1) | b)
}

func (led I2CLed) DisplayCursorOn() {
	led.Display(DIGITAL_HIGH, DIGITAL_HIGH, DIGITAL_HIGH)
}

func (led I2CLed) DisplayOn() {
	led.Display(DIGITAL_HIGH, DIGITAL_LOW, DIGITAL_LOW)
}

func (led I2CLed) SetMode(modLines, modBytes int) {
	f := 0 << 2 // 5x8 dots

	led.Command(0b00100000 | modBytes | modLines | f)
}

func (led I2CLed) Clear() {
	led.Command(0x01)
}

func (led I2CLed) FirstLineSetup() {
	led.Command(0b00000110)
}

func (led I2CLed) CreateCustomChar(inLocation int, char CustomChar) LocatedChar {
	location := inLocation & 7 // only 7 spots
	led.Command(LCD_SETCGRAMADDR | (location << 3))
	for _, ch := range char.Rows {
		led.Char(ch)
	}

	return NewLocatedChar(inLocation)
}

func (led I2CLed) SetCursor(x, y int) {
	led.Command(LCD_SETDDRAMADDR | (x + led.row_offsets[y]))
}

func (led I2CLed) WriteCustomChar(location int) {
	led.Char(location)
}

func lcdText[T string | int](led I2CLed, text T, line int) {
	led.Command(line)

	switch v := any(text).(type) {
	case string:
		for _, chr := range v {
			led.Char(int(chr))
		}
	case int:
		led.Char(v)
	}
}

func waitPulse() {
	time.Sleep(time.Duration(PULSE_SLEEP_NS) * time.Nanosecond)
}
