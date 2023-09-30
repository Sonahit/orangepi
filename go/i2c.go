package main

// #cgo LDFLAGS: -lwiringPi
// #include "lib/i2c.hpp"
import "C"
import (
	"fmt"
	"log"
)

type I2CFd struct {
	fd C.int
}

func I2cSetup(bus int, deviceId int) (fd I2CFd, err int) {
	fdId := C.i2cSetupDevice(C.CString(fmt.Sprintf("/dev/i2c-%d", bus)), C.int(deviceId))

	if fdId < 0 {
		return I2CFd{}, I2cError()
	}

	return I2CFd{fd: fdId}, 0
}

func I2cError() int {
	return int(C.i2cGetError())
}

func I2cWrite(fd I2CFd, data int) {
	log.Printf("Writing %d\n", data)
	C.i2cWrite(fd.fd, C.int(data))
}

func I2cWriteReg8(fd I2CFd, reg uint, data int) {
	C.i2cWriteReg8(fd.fd, C.int(reg), C.int(data))
}

func I2cWriteReg16(fd *I2CFd, reg uint, data int) {
	C.i2cWriteReg16(fd.fd, C.int(reg), C.int(data))
}
