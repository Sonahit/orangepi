#include <string>

#pragma once

#define PIN_RS 0
#define PIN_E 1
#define PIN_D4 2
#define PIN_D5 3
#define PIN_D6 4
#define PIN_D7 5

const int LCD_DELAY_MS = 5;

typedef unsigned int uint;

void switchToCommand();
void switchToChar();
void readLcd();
void writeDataPins(uint data);
void lcdCommand(uint command);
void lcdChar(const char chr);
void lcdString(std::string str);

void setupPinGIOP(uint index, uint pin, int mode);