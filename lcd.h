#include <string>

#include "wiringPi.h"

#ifndef PIN_RS
#define PIN_RS 7
#endif

#ifndef PIN_E
#define PIN_E 8
#endif

#ifndef PIN_D4
#define PIN_D4 9
#endif

#ifndef PIN_D5
#define PIN_D5 10
#endif

#ifndef PIN_D6
#define PIN_D6 11
#endif

#ifndef PIN_D7
#define PIN_D7 12
#endif

#pragma once

const int LCD_DELAY_MS = 5;

typedef unsigned int uint;

void switchToCommand() {
  digitalWrite(PIN_RS, LOW);
  delay(LCD_DELAY_MS);
}

void switchToChar() {
  digitalWrite(PIN_RS, HIGH);
  delay(LCD_DELAY_MS);
}

void writeToDisplay() {
  digitalWrite(PIN_E, HIGH);
  delay(LCD_DELAY_MS);
  digitalWrite(PIN_E, LOW);
}

void writeDataPins(uint data) {
  digitalWrite(PIN_D7, (data >> 7) & 1);
  digitalWrite(PIN_D6, (data >> 6) & 1);
  digitalWrite(PIN_D5, (data >> 5) & 1);
  digitalWrite(PIN_D4, (data >> 4) & 1);

  // When writing to the display, data is transfered only
  // on the high to low transition of the E signal.
  writeToDisplay();

  digitalWrite(PIN_D7, (data >> 3) & 1);
  digitalWrite(PIN_D6, (data >> 2) & 1);
  digitalWrite(PIN_D5, (data >> 1) & 1);
  digitalWrite(PIN_D4, (data >> 0) & 1);

  writeToDisplay();
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