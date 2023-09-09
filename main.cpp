

#include <iostream>

#include "lcd.h"
#include "wiringPi.h"

void println(std::string text) { std::cout << text << std::endl; }

void setupPins(int argc, char* argv[]) {
  uint pinRs = atoi(argv[PIN_RS + 1]);
  uint pinRW = atoi(argv[PIN_RW + 1]);
  uint pinE = atoi(argv[PIN_E + 1]);
  uint pinD4 = atoi(argv[PIN_D4 + 1]);
  uint pinD5 = atoi(argv[PIN_D5 + 1]);
  uint pinD6 = atoi(argv[PIN_D6 + 1]);
  uint pinD7 = atoi(argv[PIN_D7 + 1]);

  println("Setup pins");

  setupPinGPIO(PIN_RS, pinRs, OUTPUT, "RS");
  setupPinGPIO(PIN_RW, pinRW, OUTPUT, "RW");
  setupPinGPIO(PIN_E, pinE, OUTPUT, "E");
  setupPinGPIO(PIN_D4, pinD4, OUTPUT, "D4");
  setupPinGPIO(PIN_D5, pinD5, OUTPUT, "D5");
  setupPinGPIO(PIN_D6, pinD6, OUTPUT, "D6");
  setupPinGPIO(PIN_D7, pinD7, OUTPUT, "D7");

  writeModeLcd();

  println("Setup pins done");

  // 4-bit mode, 2 lines, 5x7 format
  lcdCommand(0x38);

  // LCD ON
  lcdCommand(0x0f);

  // RETURN HOME
  // lcdCommand(0x02);

  // clear display (optional here)
  lcdCommand(0x01);

  // move to first line
  lcdCommand(0x80);
  lcdString("Using HD44780");

  // move to second line
  lcdCommand(0xC0);

  lcdString("LCD directly! :)");
}

int main(int argc, char* argv[]) {
  if (wiringPiSetup() != 0) {
    return 1;
  }

  setupPins(argc, argv);

  for (;;) {
  }
  return 0;
}