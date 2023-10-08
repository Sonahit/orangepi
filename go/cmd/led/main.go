package main

import (
	"display/lib"
	"display/pkg"
	"log"
	"os"
	"os/signal"
	"syscall"
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
	0b10001,
	0b10001,
	0b10001,
	0b10000,
	0b01000,
	0b00100,
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
	0b0000000000,
	0b0111100000,
	0b1000010000,
	0b1000001000,
	0b0000000100,
	0b0000000100,
	0b0000000100,
	0b0000111100,
	0b0000111100,
	0b0000000100,
	0b0000000100,
	0b0000000100,
	0b1000001000,
	0b1000010000,
	0b0111100000,
	0b0000000000,
})

func setupLeds(lcd I2CLed) {
	log.Println("Setup")
	defer log.Println("SetupDone")

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		lcd.Clear()
		os.Exit(1)
	}()

	lcd.SetMode(LCD_MODE_TWO_LINES, LCD_MODE_FOUR_BYTES)
	lcd.SetMode(LCD_MODE_TWO_LINES, LCD_MODE_FOUR_BYTES)
	lcd.DisplayOn()
	lcd.FirstLineSetup()
	lcd.SetMode(LCD_MODE_TWO_LINES, LCD_MODE_FOUR_BYTES)
	lcd.Clear()
}

func logicLed(lcd I2CLed) {
	lineOne := pkg.LeftPad("1", ' ', lcd.Width)
	lineTwo := pkg.RightPad("1", ' ', lcd.Width)
	sectionIacoNum, iacoSections := chIaco.SplitBySections()
	_, golovaSections := chGolova.SplitBySections()

	for i, section := range iacoSections {
		charLoc := i
		lcd.CreateCustomChar(charLoc, section)
		pkg.Sleep(100)
	}

	lcd.CreateCustomChar(sectionIacoNum, chProdol)

	for i, section := range golovaSections {
		charLoc := i + sectionIacoNum + 1
		lcd.CreateCustomChar(charLoc, section)
		pkg.Sleep(100)
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

		lcd.SetCursor(8, 0)
		lcd.WriteCustomChar(0 + sectionIacoNum + 1)
		lcd.SetCursor(7, 0)
		lcd.WriteCustomChar(1 + sectionIacoNum + 1)
		lcd.SetCursor(8, 1)
		lcd.WriteCustomChar(2 + sectionIacoNum + 1)
		lcd.SetCursor(7, 1)
		lcd.WriteCustomChar(3 + sectionIacoNum + 1)

		// pkg.sleep(1000)

		// lcd.TextString(lineTwo, LCD_LINE_ONE)
		// lcd.TextString(lineOne, LCD_LINE_TWO)
		// lcd.SetCursor(7, 1)
		// lcd.WriteCustomChar(1)

		log.Println("Logic End")
		pkg.Sleep(1000)
	}
}

func main() {
	fd, err := lib.I2cSetup(LCD_BUS, LCD_PORT)

	if err != 0 {
		log.Fatalf("Error on I2CInit %d", err)
	}

	lcd := NewI2CLed(fd)
	setupLeds(lcd)
	pkg.Sleep(2000)
	logicLed(lcd)
}
