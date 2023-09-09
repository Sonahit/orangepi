

#include "lcd.h"
#include "wiringPi.h"

void setupPins(int argc, char* argv[]) {
  uint pinRs = (uint)argv[PIN_RS];
  uint pinE = (uint)argv[PIN_E];
  uint pinD4 = (uint)argv[PIN_D4];
  uint pinD5 = (uint)argv[PIN_D5];
  uint pinD6 = (uint)argv[PIN_D6];
  uint pinD7 = (uint)argv[PIN_D7];

  setupPinGIOP(PIN_RS, pinRs, OUTPUT);
  setupPinGIOP(PIN_E, pinE, OUTPUT);
  setupPinGIOP(PIN_D4, pinD4, OUTPUT);
  setupPinGIOP(PIN_D5, pinD5, OUTPUT);
  setupPinGIOP(PIN_D6, pinD6, OUTPUT);
  setupPinGIOP(PIN_D7, pinD7, OUTPUT);

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

int main(int argc, char* argv[]) {
  setupPins(argc, argv);
  return 0;
}