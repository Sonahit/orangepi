#define PIN_RS 7
#define PIN_E 8
#define PIN_D4 9
#define PIN_D5 10
#define PIN_D6 11
#define PIN_D7 12

#include "lcd.h"
#include "wiringPi.h"

void setupPins() {
  pinMode(PIN_RS, OUTPUT);
  pinMode(PIN_E, OUTPUT);
  pinMode(PIN_D4, OUTPUT);
  pinMode(PIN_D5, OUTPUT);
  pinMode(PIN_D6, OUTPUT);
  pinMode(PIN_D7, OUTPUT);

  // 4-bit mode, 2 lines, 5x7 format
  lcdCommand(0b00110000);
  // display & cursor home (keep this!)
  lcdCommand(0b00000010);
  // display on, right shift, underline off, blink off
  lcdCommand(0b00001100);
  // clear display (optional here)
  lcdCommand(0b00000001);

  lcdCommand(0b10000000);  // set address to 0x00
  lcdString("Using HD44780");
  lcdCommand(0b11000000);  // set address to 0x40
  lcdString("LCD directly! :)");
}

int main() {
  setupPins();
  return 0;
}