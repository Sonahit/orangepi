

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
  uint pinD0 = atoi(argv[PIN_D0 + 1]);
  uint pinD1 = atoi(argv[PIN_D1 + 1]);
  uint pinD2 = atoi(argv[PIN_D2 + 1]);
  uint pinD3 = atoi(argv[PIN_D3 + 1]);
  setupPinMode(atoi(argv[PIN_MODE + 1]));

  println("Setup pins");

  setupPinGPIO(PIN_RS, pinRs, OUTPUT, "RS");
  setupPinGPIO(PIN_RW, pinRW, OUTPUT, "RW");
  setupPinGPIO(PIN_E, pinE, OUTPUT, "E");
  setupPinGPIO(PIN_D4, pinD4, OUTPUT, "D4");
  setupPinGPIO(PIN_D5, pinD5, OUTPUT, "D5");
  setupPinGPIO(PIN_D6, pinD6, OUTPUT, "D6");
  setupPinGPIO(PIN_D7, pinD7, OUTPUT, "D7");
  setupPinGPIO(PIN_D0, pinD0, OUTPUT, "D0");
  setupPinGPIO(PIN_D1, pinD1, OUTPUT, "D1");
  setupPinGPIO(PIN_D2, pinD2, OUTPUT, "D2");
  setupPinGPIO(PIN_D3, pinD3, OUTPUT, "D3");

  digitalWrite(PIN_E, LOW);
  writeModeLcd();

  println("Setup pins done");
}

void logic() {
  if (is4bitMode()) {
    // 4-bit mode, 2 lines, 5x7 format
    lcdCommand(0x28);
  }

  // lcd on cursor blink
  lcdCommand(0x0f);

  // clear display (optional here)
  lcdCommand(0x01);

  // move to first line
  lcdCommand(0x80);
  lcdString("plz hlep");

  // move to second line
  // lcdCommand(0xC0);
  // lcdString("LCD directly! :)");
}

int main(int argc, char* argv[]) {
  if (wiringPiSetup() != 0) {
    return 1;
  }

  setupPins(argc, argv);

  logic();

  return 0;
}