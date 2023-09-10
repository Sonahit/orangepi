#include <string>

#pragma once

typedef unsigned int uint;

struct Pin {
  uint index;
  uint mode;
};

struct Pins {
  bool is4PinMode = false;
  Pin RS;
  Pin RW;
  Pin E;
  Pin D0;
  Pin D1;
  Pin D2;
  Pin D3;
  Pin D4;
  Pin D5;
  Pin D6;
  Pin D7;
};

const int LCD_DELAY_MS = 50;

void switchToCommand();
void switchToChar();
void readLcd();
void writeDataPins4Pin(uint data);
void lcdCommand(uint command);
void lcdString(const std::string &str);
void readModeLcd();
void writeModeLcd();
bool is4bitMode();
uint initLcd(Pins iniPins);
void lcdDigitalWrite(uint pin, uint value);