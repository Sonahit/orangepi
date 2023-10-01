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

var chAll = NewCustomChar([]int{
	0b00000000000000000000000000000000000000000000,
	0b00011110000000000000000000000000000000000000,
	0b00100001000000000000000000000000000000000000,
	0b01000000100000000000000000000000000000000000,
	0b01000000011111111111111111111111111111000000,
	0b01000000000000000000000000000000010011100000,
	0b00100000000000000000000000000000010000110000,
	0b00100000000000000000000000000000010000010000,
	0b00100000000000000000000000000000011111111000,
	0b00100000000000000000000000000000010000010000,
	0b01000000000000000000000000000000010000110000,
	0b01000000011111111111111111111111111111100000,
	0b01000000100000000000000000000000000000000000,
	0b01000001000000000000000000000000000000000000,
	0b00011110000000000000000000000000000000000000,
	0b00000000000000000000000000000000000000000000,
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
	defer func() {
		lcd.Clear()
	}()
}

func setup(lcd I2CLed) {
	log.Println("Setup")
	defer log.Println("SetupDone")

	lcd.SetMode(LCD_MODE_TWO_LINES, LCD_MODE_FOUR_BYTES)
	lcd.SetMode(LCD_MODE_TWO_LINES, LCD_MODE_FOUR_BYTES)
	lcd.DisplayOn()
	lcd.FirstLineSetup()
	lcd.SetMode(LCD_MODE_TWO_LINES, LCD_MODE_FOUR_BYTES)
	lcd.Clear()
}

func logic(lcd I2CLed) {
	lineOne := LeftPad("1", ' ', lcd.Width)
	lineTwo := RightPad("1", ' ', lcd.Width)
	sectionsNum, sections := chAll.SplitBySections()
	for i := 0; i < sectionsNum; i++ {
		lcd.CreateCustomChar(i, sections[i])
	}
	for {
		log.Println("Logic start")
		lcd.TextString(lineOne, LCD_LINE_ONE)
		lcd.TextString(lineTwo, LCD_LINE_TWO)

		lcd.SetCursor(4, 0)
		lcd.WriteCustomChar(0)

		lcd.SetCursor(3, 0)
		lcd.WriteCustomChar(1)

		lcd.SetCursor(2, 0)
		lcd.WriteCustomChar(2)

		lcd.SetCursor(4, 1)
		lcd.WriteCustomChar(3)

		lcd.SetCursor(3, 1)
		lcd.WriteCustomChar(4)

		lcd.SetCursor(2, 1)
		lcd.WriteCustomChar(5)

		// sleep(1000)

		// lcd.TextString(lineTwo, LCD_LINE_ONE)
		// lcd.TextString(lineOne, LCD_LINE_TWO)
		// lcd.SetCursor(7, 1)
		// lcd.WriteCustomChar(1)

		log.Println("Logic End")
		sleep(1000)
	}
}

func sleep(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
