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

var chIaco = NewCustomChar([]int{
	0b00011111000,
	0b00100000100,
	0b01000000010,
	0b10000000001,
	0b01000000000,
	0b00100000000,
	0b00100000000,
	0b00100000000,
	0b00100000000,
	0b00100000000,
	0b01000000000,
	0b10000000000,
	0b01000000001,
	0b00100000010,
	0b00010000100,
	0b00001111000,
})

var chProdol = NewCustomChar([]int{
	0b00000000000,
	0b00000000000,
	0b00000000000,
	0b11111111111,
	0b00000000000,
	0b00000000000,
	0b00000000000,
	0b00000000000,
	0b00000000000,
	0b00000000000,
	0b00000000000,
	0b00000000000,
	0b11111111111,
	0b00000000000,
	0b00000000000,
	0b00000000000,
})

var chGolova = NewCustomChar([]int{
	0b00000000000,
	0b00011111000,
	0b01000000100,
	0b10000000010,
	0b00000000001,
	0b00000000001,
	0b00000000001,
	0b11111111111,
	0b11111111111,
	0b00000000001,
	0b00000000001,
	0b00000000001,
	0b10000000010,
	0b00100000100,
	0b00011111000,
	0b00000000000,
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
	sectionProdolNum, prodolSections := chProdol.SplitBySections()
	sectionGolovaNum, golovaSections := chGolova.SplitBySections()

	locatedChars := make([]LocatedChar, 0, sectionGolovaNum+sectionIacoNum+sectionProdolNum)

	for i := 0; i < sectionIacoNum; i++ {
		locatedChars = append(locatedChars, lcd.CreateCustomChar(i, iacoSections[i]))
	}

	for i := 0; i < sectionProdolNum; i++ {
		locatedChars = append(locatedChars, lcd.CreateCustomChar(i+sectionProdolNum, prodolSections[i]))
	}

	for i := 0; i < sectionGolovaNum; i++ {
		locatedChars = append(locatedChars, lcd.CreateCustomChar(i+sectionProdolNum+sectionIacoNum, golovaSections[i]))
	}

	iaceVerh,
		iaceNiz,
		prodolVerh,
		prodolNiz,
		golovaVerh,
		golovaNiz := locatedChars[0],
		locatedChars[1],
		locatedChars[2],
		locatedChars[3],
		locatedChars[4],
		locatedChars[5]

	for {
		log.Println("Logic start")
		lcd.TextString(lineOne, LCD_LINE_ONE)
		lcd.TextString(lineTwo, LCD_LINE_TWO)

		lcd.SetCursor(4, 0)
		lcd.WriteCustomChar(iaceVerh.Location)

		lcd.SetCursor(4, 1)
		lcd.WriteCustomChar(iaceNiz.Location)

		lcd.SetCursor(5, 0)
		lcd.WriteCustomChar(prodolVerh.Location)

		lcd.SetCursor(5, 1)
		lcd.WriteCustomChar(prodolNiz.Location)

		lcd.SetCursor(6, 0)
		lcd.WriteCustomChar(prodolVerh.Location)

		lcd.SetCursor(6, 1)
		lcd.WriteCustomChar(prodolNiz.Location)

		lcd.SetCursor(7, 0)
		lcd.WriteCustomChar(prodolVerh.Location)

		lcd.SetCursor(7, 1)
		lcd.WriteCustomChar(prodolNiz.Location)

		lcd.SetCursor(8, 0)
		lcd.WriteCustomChar(golovaVerh.Location)

		lcd.SetCursor(8, 1)
		lcd.WriteCustomChar(golovaNiz.Location)

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
