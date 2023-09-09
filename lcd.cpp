#include "lcd.h"

#include "iostream"
#include "wiringPi.h"

// RS, RW, E,  D4, D5, D6, D7, D1, D2, D3
// 0   1  2   3   4   5   6, 7, 8, 9, 10
uint pinSetup[11] = {};

static bool is4PinMode = false;

uint pin(uint index) { return pinSetup[index]; }
void setupPinGPIO(uint index, uint pin, int mode, std::string type) {
  pinSetup[index] = pin;
  std::cout << "Setup " << type << " pin " << pin << " with mode " << mode
            << std::endl;
  pinMode(pin, mode);
}

void switchToCommand() {
  digitalWrite(pin(PIN_RS), LOW);
  delay(LCD_DELAY_MS);
}

void switchToChar() {
  digitalWrite(pin(PIN_RS), HIGH);
  delay(LCD_DELAY_MS);
}

void readLcd() {
  digitalWrite(pin(PIN_E), HIGH);
  delay(LCD_DELAY_MS);
  digitalWrite(pin(PIN_E), LOW);
}

void writeDataPins4Pin(uint data) {
  digitalWrite(pin(PIN_D7), (data >> 7) & 1);
  digitalWrite(pin(PIN_D6), (data >> 6) & 1);
  digitalWrite(pin(PIN_D5), (data >> 5) & 1);
  digitalWrite(pin(PIN_D4), (data >> 4) & 1);

  // When writing to the display, data is transfered only
  // on the high to low transition of the E signal.
  readLcd();

  digitalWrite(pin(PIN_D7), (data >> 3) & 1);
  digitalWrite(pin(PIN_D6), (data >> 2) & 1);
  digitalWrite(pin(PIN_D5), (data >> 1) & 1);
  digitalWrite(pin(PIN_D4), (data >> 0) & 1);

  readLcd();
}

void writeDataPins8Pin(uint data) {
  digitalWrite(pin(PIN_D7), (data >> 7) & 1);
  digitalWrite(pin(PIN_D6), (data >> 6) & 1);
  digitalWrite(pin(PIN_D5), (data >> 5) & 1);
  digitalWrite(pin(PIN_D4), (data >> 4) & 1);

  digitalWrite(pin(PIN_D3), (data >> 3) & 1);
  digitalWrite(pin(PIN_D2), (data >> 2) & 1);
  digitalWrite(pin(PIN_D1), (data >> 1) & 1);
  digitalWrite(pin(PIN_D0), (data >> 0) & 1);

  readLcd();
}

void writeDataPins(uint data) {
  is4PinMode ? writeDataPins4Pin(data) : writeDataPins8Pin(data);
}

void lcdCommand(uint command) {
  switchToCommand();
  writeDataPins(command);
}

void setupPinMode(uint mode) { is4PinMode = mode == 0 ? true : false; }

void readModeLcd() { digitalWrite(pin(PIN_RW), HIGH); }

void writeModeLcd() { digitalWrite(pin(PIN_RW), LOW); }

void lcdChar(const char chr) { writeDataPins((uint)chr); }

void lcdString(std::string str) {
  switchToChar();
  for (std::string::iterator it = str.begin(); it != str.end(); ++it) {
    lcdChar(*it);
  }
}