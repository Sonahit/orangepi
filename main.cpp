

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
  std::cout << "Got args " << argc << std::endl;

  if (wiringPiSetup() != 0) {
    return 1;
  }

  setupPins(argc, argv);

  for (;;) {
  }
  return 0;
}