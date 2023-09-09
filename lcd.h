#include <string>

#pragma once

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

const int LCD_DELAY_MS = 5;

typedef unsigned int uint;

void switchToCommand();
void switchToChar();
void writeToDisplay();
void writeDataPins(uint data);
void lcdCommand(uint command);
void lcdChar(const char chr);
void lcdString(std::string str);