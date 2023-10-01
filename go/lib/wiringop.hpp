
#include "softPwm.h"
#include "wiringPi.h"

int setup() { return wiringPiSetup(); }

void _digitalWrite(int pin, int value) { digitalWrite(pin, value); }
void _pinMode(int pin, int value) { pinMode(pin, value); }