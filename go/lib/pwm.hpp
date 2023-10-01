
#include "softPwm.h"
#include "wiringPi.h"

void pwmCreate(int pin, int value, int range) {
  softPwmCreate(pin, value, range);
}
void pwmWrite(int pin, int value) { softPwmWrite(pin, value); }
void pwmStop(int pin) { softPwmStop(pin); }
