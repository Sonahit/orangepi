
#include "softTone.h"
#include "wiringPi.h"

void toneCreate(int pin) { softToneCreate(pin); }
void toneWrite(int pin, int value) { softToneWrite(pin, value); }
void toneStop(int pin) { softToneStop(pin); }