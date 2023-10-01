package main

import (
	"log"
	"time"
)

// https://www.sparkfun.com/datasheets/LCD/HD44780.pdf
// (ROM Code: A00)

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

var chIaco = NewCustomChar([]int{
	0b00000,
	0b01110,
	0b01001,
	0b10001,
	0b10001,
	0b10000,
	0b01000,
	0b01000,
	0b00100,
	0b01000,
	0b01000,
	0b10001,
	0b10001,
	0b10001,
	0b01110,
	0b00000,
})

var chProdol = NewCustomChar([]int{
	0b00000,
	0b00000,
	0b00000,
	0b11111,
	0b11111,
	0b00000,
	0b00000,
	0b00000,
})

var chGolova = NewCustomChar([]int{
	0b00000,
	0b01111,
	0b10011,
	0b10011,
	0b10011,
	0b00011,
	0b00011,
	0b11111,
	0b11111,
	0b00011,
	0b00011,
	0b10011,
	0b10011,
	0b10011,
	0b01111,
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
	sectionIacoNum, iacoSections := chIaco.SplitBySections()
	_, golovaSections := chGolova.SplitBySections()

	for i, section := range iacoSections {
		charLoc := i
		lcd.CreateCustomChar(charLoc, section)
		sleep(100)
	}

	lcd.CreateCustomChar(sectionIacoNum, chProdol)

	for i, section := range golovaSections {
		charLoc := i + sectionIacoNum + 1
		lcd.CreateCustomChar(charLoc, section)
		sleep(100)
	}

	charLoc := sectionIacoNum

	for {
		log.Println("Logic start")
		lcd.TextString(lineOne, LCD_LINE_ONE)
		lcd.TextString(lineTwo, LCD_LINE_TWO)

		for i := range iacoSections {
			charLoc := i
			lcd.SetCursor(2, i)
			lcd.WriteCustomChar(charLoc)
		}
		for i := 0; i < 2; i++ {
			lcd.SetCursor(3, i)
			lcd.WriteCustomChar(charLoc)
			lcd.SetCursor(4, i)
			lcd.WriteCustomChar(charLoc)
			lcd.SetCursor(5, i)
			lcd.WriteCustomChar(charLoc)
			lcd.SetCursor(6, i)
			lcd.WriteCustomChar(charLoc)
		}

		for i := range golovaSections {
			charLoc := i + sectionIacoNum + 1
			lcd.SetCursor(7, i)
			lcd.WriteCustomChar(charLoc)
		}

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
