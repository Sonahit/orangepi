#include "lcd.h"

#include "iostream"
#include "wiringPi.h"

// RS, E, D4, D5, D6, D7
// 0   1  2   3   4   5
uint pinSetup[6] = {};

uint pin(uint index) { return pinSetup[index]; }
void setupPinGIOP(uint index, uint pin, int mode) {
  pinSetup[index] = pin;
  std::cout << "Setup pin " << pin << " with mode " << mode << std::endl;
  pinMode(pin, mode);
}

void switchToCommand() {
  digitalWrite(pin(PIN_RS), LOW);
  delay(LCD_DELAY_MS);
}

void switchToChar() {
  digitalWrite(pin(PIN_RS), HIGH);
  delay(LCD_DELAY_MS);
}

void readLcd() {
  digitalWrite(pin(PIN_E), HIGH);
  delay(LCD_DELAY_MS);
  digitalWrite(pin(PIN_E), LOW);
}

void writeDataPins(uint data) {
  digitalWrite(pin(PIN_D7), (data >> 7) & 1);
  digitalWrite(pin(PIN_D6), (data >> 6) & 1);
  digitalWrite(pin(PIN_D5), (data >> 5) & 1);
  digitalWrite(pin(PIN_D4), (data >> 4) & 1);

  // When writing to the display, data is transfered only
  // on the high to low transition of the E signal.
  readLcd();

  digitalWrite(pin(PIN_D7), (data >> 3) & 1);
  digitalWrite(pin(PIN_D6), (data >> 2) & 1);
  digitalWrite(pin(PIN_D5), (data >> 1) & 1);
  digitalWrite(pin(PIN_D4), (data >> 0) & 1);

  readLcd();
}

void lcdCommand(uint command) {
  switchToCommand();
  writeDataPins(command);
}

void lcdChar(const char chr) {
  switchToChar();
  writeDataPins((uint)chr);
}

void lcdString(std::string str) {
  for (std::string::iterator it = str.begin(); it != str.end(); ++it) {
    lcdChar(*it);
  }
}