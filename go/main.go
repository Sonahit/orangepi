package main

import (
	"log"
	"time"
)

const (
	LCD_BUS  = 0
	LCD_PORT = 0x27
)

func main() {
	fd, err := I2cSetup(LCD_BUS, LCD_PORT)

	if err != 0 {
		log.Fatalf("Error on I2CInit %d", err)
	}

	lcd := NewI2CLed(fd)
	setup(lcd)
	logic(lcd)
}

func setup(lcd I2CLed) {
	println("Setup")
	defer println("SetupDone")

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

	for {

		lcd.TextString(lineOne, LCD_LINE_ONE)
		lcd.TextString(lineTwo, LCD_LINE_TWO)

		sleep(1000)

		lcd.TextString(lineTwo, LCD_LINE_ONE)
		lcd.TextString(lineOne, LCD_LINE_TWO)

		sleep(1000)
	}
}

func sleep(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
