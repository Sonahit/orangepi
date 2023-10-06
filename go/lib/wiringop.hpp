
#include "softPwm.h"
#include "wiringPi.h"

int setup() { return wiringPiSetup(); }

void _digitalWrite(int pin, int value) { digitalWrite(pin, value); }
int _digitalRead(int pin) { return digitalRead(pin); }
void _pinMode(int pin, int value) { pinMode(pin, value); }