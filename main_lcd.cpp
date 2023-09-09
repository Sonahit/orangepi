

#include <iostream>

#include "lcd.h"
#include "wiringPi.h"

#define PIN_RS 0
#define PIN_RW 1
#define PIN_E 2
#define PIN_D4 3
#define PIN_D5 4
#define PIN_D6 5
#define PIN_D7 6

void println(std::string text) { std::cout << text << std::endl; }

void setupPins(int argc, char* argv[]) {
  println("Setup pins");
  uint pinRs = atoi(argv[PIN_RS + 1]);
  uint pinRW = atoi(argv[PIN_RW + 1]);
  uint pinE = atoi(argv[PIN_E + 1]);
  uint pinD4 = atoi(argv[PIN_D4 + 1]);
  uint pinD5 = atoi(argv[PIN_D5 + 1]);
  uint pinD6 = atoi(argv[PIN_D6 + 1]);
  uint pinD7 = atoi(argv[PIN_D7 + 1]);

  int fd =
      lcdInit(2, 16, 4, pinRs, pinE, pinD4, pinD5, pinD6, pinD7, 0, 0, 0, 0);

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