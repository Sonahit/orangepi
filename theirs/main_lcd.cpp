

#include <iostream>

#include "lcd.h"
#include "wiringPi.h"

#define PIN_RS 23
#define PIN_RW 23
#define PIN_E 22
#define PIN_D4 5
#define PIN_D5 2
#define PIN_D6 1
#define PIN_D7 0
#define PIN_D0 20
#define PIN_D1 19
#define PIN_D2 17
#define PIN_D3 3

void println(std::string text) { std::cout << text << std::endl; }

void setupPins(int argc, char* argv[]) {
  println("Setup pins");

  int fd = lcdInit(2, 16, 8, PIN_RS, PIN_E, PIN_D0, PIN_D1, PIN_D2, PIN_D3,
                   PIN_D4, PIN_D5, PIN_D6, PIN_D7);

  println("Setup pins done");
  lcdCursor(fd, 1);
  lcdCursorBlink(fd, 1);

  lcdClear(fd);
  lcdPosition(fd, 0, 0);
  lcdPuts(fd, "Hello, world!");
  delay(1000);
}

int main(int argc, char* argv[]) {
  if (wiringPiSetup() != 0) {
    return 1;
  }

  setupPins(argc, argv);

  return 0;
}