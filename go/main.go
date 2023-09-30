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
	lcd.CreateCustomChar(0, smileChar)
	for {
		log.Println("Logic start")
		lcd.TextString(lineOne, LCD_LINE_ONE)
		lcd.TextString(lineTwo, LCD_LINE_TWO)
		lcd.SetCursor(3, 0)
		lcd.WriteCustomChar(0)

		sleep(1000)

		lcd.TextString(lineTwo, LCD_LINE_ONE)
		lcd.TextString(lineOne, LCD_LINE_TWO)
		lcd.SetCursor(7, 1)
		lcd.WriteCustomChar(0)

		sleep(1000)
		log.Println("Logic End")
	}
}

func sleep(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
