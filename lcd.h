#include <string>

#pragma once

#define PIN_RS 0
#define PIN_RW 1
#define PIN_E 2
#define PIN_D4 3
#define PIN_D5 4
#define PIN_D6 5
#define PIN_D7 6
#define PIN_D0 7
#define PIN_D1 8
#define PIN_D2 9
#define PIN_D3 10
#define PIN_MODE 11

const int LCD_DELAY_MS = 50;

typedef unsigned int uint;

void switchToCommand();
void switchToChar();
void readLcd();
void writeDataPins4Pin(uint data);
void lcdCommand(uint command);
void lcdString(std::string str);
void readModeLcd();
void writeModeLcd();
// 0 == 4pin, anything else is 8
void setupPinMode(uint mode);
bool is4bitMode();

void setupPinGPIO(uint index, uint pin, int mode, std::string type);