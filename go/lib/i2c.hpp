#include "errno.h"
#include "wiringPi.h"
#include "wiringPiI2C.h"

#pragma once

/**
 * @param interface  /dev/i2c-0
 * @returns fd
 */
int i2cSetupDevice(const char* devicePath, int deviceId) {
  int fd = wiringPiI2CSetupInterface(devicePath, deviceId);

  return fd;
}

/**
 * @returns fd
 */
int i2cSetup(const int devId) {
  int fd = wiringPiI2CSetup(devId);

  return fd;
}

int i2cGetError() { return errno; }

int i2cRead(int fd) { return wiringPiI2CRead(fd); }
int i2cReadReg8(int fd, int reg) { return wiringPiI2CReadReg8(fd, reg); }
int i2cReadReg16(int fd, int reg) { return wiringPiI2CReadReg16(fd, reg); }

int i2cWrite(int fd, int data) { return wiringPiI2CWrite(fd, data); }

int i2cWriteReg8(int fd, int reg, int data) {
  return wiringPiI2CWriteReg8(fd, reg, data);
}

int i2cWriteReg16(int fd, int reg, int data) {
  return wiringPiI2CWriteReg16(fd, reg, data);
}
