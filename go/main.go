package main

import (
	"log"
	"time"
)

const (
	LCD_BUS  = 0
	LCD_PORT = 0x27
)

var smileChar = NewCustomChar([]int{
	0b00000,
	0b00000,
	0b01010,
	0b00000,
	0b00000,
	0b10001,
	0b01110,
	0b00000,
})

var ch1 = NewCustomChar([]int{
	0b11111,
	0b00100,
	0b00100,
	0b00100,
	0b00100,
	0b00100,
	0b10100,
	0b01100,
})

var ch2 = NewCustomChar([]int{
	0b00000,
	0b00001,
	0b00110,
	0b01001,
	0b01000,
	0b01111,
	0b01000,
	0b01000,
})

var ch3 = NewCustomChar([]int{
	0b01000,
	0b01000,
	0b01000,
	0b01000,
	0b01001,
	0b00110,
	0b00010,
	0b00000,
})

var ch4 = NewCustomChar([]int{
	0b00101,
	0b01001,
	0b10001,
	0b00001,
	0b00001,
	0b00001,
	0b00001,
	0b11111,
})

var ch5 = NewCustomChar([]int{
	0b11110,
	0b00010,
	0b00010,
	0b00010,
	0b10010,
	0b01100,
	0b10000,
	0b00000,
})

var ch6 = NewCustomChar([]int{
	0b11110,
	0b00010,
	0b00010,
	0b00010,
	0b10010,
	0b01100,
	0b10000,
	0b00000,
})

var chAll = NewCustomChar([]int{
	0b000111111110000,
	0b001000010001000,
	0b011000010000100,
	0b010100010001010,
	0b010010010010010,
	0b010001010100010,
	0b010000111000010,
	0b011111111111110,
	0b010000111000010,
	0b010001010100010,
	0b010010010010010,
	0b010010010010010,
	0b010010010001010,
	0b010100010000100,
	0b001000010001000,
	0b000111111110000,
})

func main() {
	fd, err := I2cSetup(LCD_BUS, LCD_PORT)

	if err != 0 {
		log.Fatalf("Error on I2CInit %d", err)
	}

	lcd := NewI2CLed(fd)
	setup(lcd)
	sleep(2000)
	logic(lcd)
}

func setup(lcd I2CLed) {
	log.Println("Setup")
	defer log.Println("SetupDone")

	lcd.SetMode(LCD_MODE_TWO_LINES, LCD_MODE_FOUR_BYTES)
	lcd.SetMode(LCD_MODE_TWO_LINES, LCD_MODE_FOUR_BYTES)
	lcd.DisplayCursorOn()
	lcd.FirstLineSetup()
	lcd.SetMode(LCD_MODE_TWO_LINES, LCD_MODE_FOUR_BYTES)
	lcd.Clear()
}

func logic(lcd I2CLed) {
	lineOne := LeftPad("1", ' ', lcd.Width)
	lineTwo := RightPad("1", ' ', lcd.Width)
	lcd.CreateCustomChar(1, smileChar)
	for {
		log.Println("Logic start")
		lcd.TextString(lineOne, LCD_LINE_ONE)
		lcd.TextString(lineTwo, LCD_LINE_TWO)
		lcd.SetCursor(3, 0)
		lcd.WriteCustomChar(1)

		sleep(1000)

		lcd.TextString(lineTwo, LCD_LINE_ONE)
		lcd.TextString(lineOne, LCD_LINE_TWO)
		lcd.SetCursor(7, 1)
		lcd.WriteCustomChar(1)

		sleep(1000)
		log.Println("Logic End")
	}
}

func sleep(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
