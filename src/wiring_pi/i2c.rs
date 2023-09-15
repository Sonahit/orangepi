#[cxx::bridge]
#[allow(dead_code)]
mod ffi {
    unsafe extern "C++" {
        include!("rs/lib/i2c.h");

        #[rust_name = "i2c_setup_device"]
        unsafe fn i2cSetupDevice(devicePath: &CxxString, deviceId: i32) -> i32;

        #[rust_name = "i2c_setup"]
        fn i2cSetup(devId: i32) -> i32;

        #[rust_name = "i2c_get_error"]
        fn i2cGetError() -> i32;

        #[rust_name = "i2c_read"]
        fn i2cRead(fd: i32) -> i32;

        #[rust_name = "i2c_read_reg8"]
        fn i2cReadReg8(fd: i32, reg: i32) -> i32;

        #[rust_name = "i2c_read_reg16"]
        fn i2cReadReg16(fd: i32, reg: i32) -> i32;

        #[rust_name = "i2c_write"]
        fn i2cWrite(fd: i32, data: i32) -> i32;

        #[rust_name = "i2c_write_reg8"]
        fn i2cWriteReg8(fd: i32, reg: i32, data: i32) -> i32;

        #[rust_name = "i2c_write_reg16"]
        fn i2cWriteReg16(fd: i32, reg: i32, data: i32) -> i32;
    }
}

use cxx::let_cxx_string;

pub use self::ffi::{
    i2c_get_error, i2c_read, i2c_read_reg16, i2c_read_reg8, i2c_setup, i2c_setup_device, i2c_write,
    i2c_write_reg16, i2c_write_reg8,
};

pub struct I2CPort(i32, i32, u32);

#[allow(dead_code)]
impl I2CPort {
    pub fn fd(&self) -> i32 {
        self.0
    }

    pub fn device_id(&self) -> i32 {
        self.1
    }

    pub fn width(&self) -> u32 {
        self.2
    }
    pub fn widthu(&self) -> usize {
        self.2 as usize
    }
}

pub fn setup_i2c(device_id: i32, width: u32) -> Result<I2CPort, i32> {
    let fd = ffi::i2c_setup(device_id);

    if fd <= 0 {
        Err(ffi::i2c_get_error())
    } else {
        Ok(I2CPort(fd, device_id, width))
    }
}

pub fn setup_i2c_device(smbus: u8, device_id: i32, width: u32) -> Result<I2CPort, i32> {
    unsafe {
        let_cxx_string!(c = format!("/dev/i2c-{}", smbus));

        let fd = ffi::i2c_setup_device(&c, device_id);

        if fd <= 0 {
            Err(ffi::i2c_get_error())
        } else {
            Ok(I2CPort(fd, device_id, width))
        }
    }
}
