#include "lcd.h"

#include "iostream"
#include "wiringPi.h"

static Pins PINS;

uint initLcd(Pins intPins) {
  if (wiringPiSetup() != 0) {
    return 1;
  }

  PINS = intPins;

  pinMode(PINS.D0.index, PINS.D0.mode);
  pinMode(PINS.D1.index, PINS.D1.mode);
  pinMode(PINS.D2.index, PINS.D2.mode);
  pinMode(PINS.D3.index, PINS.D3.mode);
  pinMode(PINS.D4.index, PINS.D4.mode);
  pinMode(PINS.D5.index, PINS.D5.mode);
  pinMode(PINS.D6.index, PINS.D6.mode);
  pinMode(PINS.D7.index, PINS.D7.mode);
  pinMode(PINS.E.index, PINS.E.mode);
  pinMode(PINS.RS.index, PINS.RS.mode);
  pinMode(PINS.RW.index, PINS.RW.mode);

  return 0;
}

void switchToCommand() {
  digitalWrite(PINS.RS.index, LOW);
  delay(PINS.LCD_DELAY_MS);
}

void switchToChar() {
  digitalWrite(PINS.RS.index, HIGH);
  delay(PINS.LCD_DELAY_MS);
}

void readLcd() {
  digitalWrite(PINS.E.index, HIGH);
  delay(PINS.LCD_DELAY_MS);
  digitalWrite(PINS.E.index, LOW);
}

void writeDataPins4Pin(uint data) {
  digitalWrite(PINS.D7.index, (data >> 7) & 1);
  digitalWrite(PINS.D6.index, (data >> 6) & 1);
  digitalWrite(PINS.D5.index, (data >> 5) & 1);
  digitalWrite(PINS.D4.index, (data >> 4) & 1);

  // When writing to the display, data is transfered only
  // on the high to low transition of the E signal.
  readLcd();

  digitalWrite(PINS.D7.index, (data >> 3) & 1);
  digitalWrite(PINS.D6.index, (data >> 2) & 1);
  digitalWrite(PINS.D5.index, (data >> 1) & 1);
  digitalWrite(PINS.D4.index, (data >> 0) & 1);

  readLcd();
}

void writeDataPins8Pin(uint data) {
  digitalWrite(PINS.D7.index, (data >> 7) & 1);
  digitalWrite(PINS.D6.index, (data >> 6) & 1);
  digitalWrite(PINS.D5.index, (data >> 5) & 1);
  digitalWrite(PINS.D4.index, (data >> 4) & 1);

  digitalWrite(PINS.D3.index, (data >> 3) & 1);
  digitalWrite(PINS.D2.index, (data >> 2) & 1);
  digitalWrite(PINS.D1.index, (data >> 1) & 1);
  digitalWrite(PINS.D0.index, (data >> 0) & 1);

  readLcd();
}

void writeDataPins(uint data) {
  PINS.is4PinMode ? writeDataPins4Pin(data) : writeDataPins8Pin(data);
}

bool is4bitMode() { return PINS.is4PinMode; }

void lcdCommand(uint command) {
  switchToCommand();
  writeDataPins(command);
}

void readModeLcd() { digitalWrite(PINS.RW.index, HIGH); }

void writeModeLcd() { digitalWrite(PINS.RW.index, LOW); }

void lcdChar(const char chr) { writeDataPins((uint)chr); }

void lcdString(const std::string &str) {
  switchToChar();
  for (auto it = str.begin(); it != str.end(); ++it) {
    lcdChar(*it);
  }
}

void lcdDigitalWrite(uint pin, uint value) { digitalWrite(pin, value); }