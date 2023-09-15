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
const PULSE_SLEEP_S: f32 = 0.0005;
const LCD_WIDTH: u32 = 16;

fn init_i2c() -> i2c::I2CPort {
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
        let bits_high = mode | (bits & 0xF0) | LCD_BACKLIGHT;
        i2c::i2c_write(self.fd(), bits_high);
        self.lcd_enable(bits_high);

        let bits_low = mode | ((bits << 4) & 0xF0) | LCD_BACKLIGHT;
        i2c::i2c_write(self.fd(), bits_low);
        self.lcd_enable(bits_low);
    }

    fn lcd_enable(&self, bits: i32) {
        i2c::i2c_write(self.fd(), bits | LCD_ENABLE);
        thread::sleep(time::Duration::from_secs_f32(PULSE_SLEEP_S));
        i2c::i2c_write(self.fd(), bits & !LCD_ENABLE);
    }

    fn lcd_clear(&self) {
        self.lcd_cmd(0x01);
    }

    fn lcd_set_4bytes(&self, num_lines: u8) {
        let f = 0b000; // 5x8 dots
        let num_lines = if num_lines == 1u8 { 0b0000 } else { 0b1000 };
        let dl = 0b00000; // 4 bit mode
        let cmd = 0b00100000 | dl | num_lines | f;
        self.lcd_cmd(cmd);
        self.lcd_cmd(cmd & 1);
    }

    fn lcd_display_on(&self) {
        self.lcd_cmd(0b00001110);
    }
    fn lcd_first_line_setup(&self) {
        self.lcd_cmd(0b00000110);
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
    // https://www.electronicsforu.com/technology-trends/learn-electronics/16x2-lcd-pinout-diagram
    let port = init_i2c();
    port.lcd_set_4bytes(1);
    port.lcd_display_on();
    port.lcd_first_line_setup();
    port.lcd_clear();
    println!("Setup done");

    thread::sleep(time::Duration::from_millis(1000));

    println!("Logic start");
    logic(port);
}
