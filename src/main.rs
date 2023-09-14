mod wiring_pi;

use std::{thread, time};

use wiring_pi::i2c::{self, I2CPort};

const LCD_CMD: i32 = 0;
const LCD_CHAR: i32 = 1;
const LCD_BACKLIGHT: i32 = 0x08;
const LCD_PORT: i32 = 0x27;
const LCD_BUS: u8 = 0;
const LINE_1: i32 = 0x80;
const LINE_2: i32 = 0xC0;
const LCD_ENABLE: i32 = 0b100;
const PULSE_SLEEP_MS: u64 = 50;
const LCD_WIDTH: u32 = 16;

fn setup() -> i2c::I2CPort {
    match i2c::setup_i2c_device(LCD_BUS, LCD_PORT, LCD_WIDTH) {
        Ok(port) => port,
        Err(err) => {
            panic!("Errored {}", err)
        }
    }
}

impl I2CPort {
    fn lcd_string(&self, str: &str, line: i32) {
        self.lcd_cmd(line);

        for char in str.chars() {
            self.lcd_char(char)
        }
    }

    fn lcd_char(&self, char: char) {
        self.lcd_bytes(char as i32, LCD_CHAR)
    }

    fn lcd_cmd(&self, bits: i32) {
        self.lcd_bytes(bits, LCD_CMD);
    }

    fn lcd_bytes(&self, bits: i32, mode: i32) {
        // bits & 0b11110000
        let bits = mode | (bits & 0xF0) | LCD_BACKLIGHT;
        i2c::i2c_write(self.fd(), bits);
        self.lcd_enable(bits);

        let bits = mode | ((bits << 4) & 0xF0) | LCD_BACKLIGHT;
        i2c::i2c_write(self.fd(), bits);
        self.lcd_enable(bits);
    }

    fn lcd_enable(&self, bits: i32) {
        i2c::i2c_write(self.fd(), bits | LCD_ENABLE);
        thread::sleep(time::Duration::from_millis(PULSE_SLEEP_MS));
        i2c::i2c_write(self.fd(), bits & !LCD_ENABLE);
    }

    fn lcd_clear(&self) {
        self.lcd_cmd(0x01);
    }
}

fn logic(port: I2CPort) {
    // https://www.sparkfun.com/datasheets/LCD/HD44780.pdf Table 12 4bit 8digit 1 line

    loop {
        port.lcd_string("Hello  <", LINE_1);
        port.lcd_string("World  <", LINE_2);

        thread::sleep(time::Duration::from_millis(1000));

        port.lcd_string(">  Hello", LINE_1);
        port.lcd_string(">  World", LINE_2);

        thread::sleep(time::Duration::from_millis(1000));
        println!("Loop done")
    }
}

fn main() {
    println!("Setup");
    let port = setup();
    port.lcd_cmd(0x33);
    port.lcd_cmd(0x32);
    port.lcd_cmd(0x06);
    port.lcd_cmd(0x0C);
    port.lcd_cmd(0x28);
    port.lcd_clear();
    println!("Setup done");

    thread::sleep(time::Duration::from_millis(1000));

    println!("Logic start");
    logic(port);
}
