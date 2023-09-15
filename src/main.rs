mod utils;
mod wiring_pi;

use std::{
    sync::{mpsc, Arc},
    thread::{self, JoinHandle},
    time,
};

use wiring_pi::i2c::{self, I2CPort};

use utils::Padding;

const LCD_CMD: i32 = 0;
const LCD_CHAR: i32 = 1;
const LCD_BACKLIGHT: i32 = 0x08;
const LCD_PORT: i32 = 0x27;
const LCD_BUS: u8 = 0;
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
enum ModBytes {
    Eight = 0b10000,
    Four = 0b00000,
}
enum ModLines {
    One = 0b0000,
    Two = 0b1000,
}

#[derive(Clone, Copy)]
enum LinePlace {
    One = 0x80,
    Two = 0xC0,
    None = 0x00,
}

impl I2CPort {
    fn lcd_text_str(&self, str: &str, line: LinePlace) {
        match line {
            LinePlace::None => (),
            _ => self.lcd_cmd(line as i32),
        };

        for char in str.chars() {
            self.lcd_char(char as u8)
        }
    }

    fn lcd_text_string(&self, str: String, line: LinePlace) {
        match line {
            LinePlace::None => (),
            _ => self.lcd_cmd(line as i32),
        };

        for char in str.chars() {
            self.lcd_char(char as u8)
        }
    }

    fn lcd_text_u8(&self, str: &[u8], line: LinePlace) {
        match line {
            LinePlace::None => (),
            _ => self.lcd_cmd(line as i32),
        };

        for char in str {
            self.lcd_char(*char)
        }
    }

    fn lcd_char(&self, char: u8) {
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

    fn lcd_set_mode_bytes(&self, num_lines: ModLines, mod_bytes: ModBytes) {
        let f = 0 << 2; // 5x8 dots
        let cmd = 0b00100000 | mod_bytes as i32 | num_lines as i32 | f;
        self.lcd_cmd(cmd);
    }

    fn lcd_clear_line(&self, line: LinePlace) {
        self.lcd_text_string(" ".repeat(self.width() as usize), line);
    }

    fn lcd_display(
        &self,
        d: wiring_pi::DigitalByte,
        c: wiring_pi::DigitalByte,
        b: wiring_pi::DigitalByte,
    ) {
        self.lcd_cmd(0b00001000 | ((d as i32) << 2) | ((c as i32) << 1) | b as i32);
    }

    fn lcd_display_on(&self) {
        self.lcd_display(
            wiring_pi::DigitalByte::High,
            wiring_pi::DigitalByte::Low,
            wiring_pi::DigitalByte::Low,
        );
    }

    fn lcd_display_on_with_cursor(&self) {
        self.lcd_display(
            wiring_pi::DigitalByte::High,
            wiring_pi::DigitalByte::High,
            wiring_pi::DigitalByte::High,
        );
    }

    fn lcd_first_line_setup(&self) {
        self.lcd_cmd(0b00000110);
    }

    fn lcd_sleep(&self, ms: u64) {
        thread::sleep(time::Duration::from_millis(ms));
    }
    fn lcd_sleep_init(&self) {
        thread::sleep(time::Duration::from_millis(2000));
    }
}

#[derive(Clone)]
struct MovingText {
    text: String,
    fill_with: char,
    index: u8,
    line: LinePlace,
}

impl MovingText {
    fn new(text: String, line: LinePlace) -> Self {
        Self {
            text,
            fill_with: ' ',
            line,
            index: 0,
        }
    }

    fn init(&self, port: &I2CPort) {
        let padding = Padding(port.width() as usize);
        let text = padding.left_pad(&self.text, self.fill_with);
        port.lcd_text_string(text, self.line);
    }

    fn move_one(&mut self, port: &I2CPort) {
        let padding = Padding(port.width() as usize);

        if self.index + 1 < port.width() as u8 {
            self.index += 1;
        } else {
            self.index = 0
        }
        let padding_index = Padding(self.index as usize + self.text.len());

        let text = padding.left_pad(
            &padding_index.right_pad(&self.text, self.fill_with),
            self.fill_with,
        );
        port.lcd_text_string(text, self.line);

        let overflow = port.width() as i8 - (self.index as usize + self.text.len()) as i8;
        if overflow < 0 {
            port.lcd_text_string(
                self.text
                    .chars()
                    .rev()
                    .take(overflow.unsigned_abs() as usize)
                    .collect::<String>()
                    .chars()
                    .rev()
                    .collect::<String>(),
                self.line,
            );
        }
    }
}

fn logic(port: I2CPort) {
    println!("Logic start");
    // https://www.sparkfun.com/datasheets/LCD/HD44780.pdf
    // (ROM Code: A00)
    let padding = Padding(port.width() as usize);

    let mut moving_text = MovingText::new("help".to_string(), LinePlace::One);
    moving_text.init(&port);

    loop {
        // port.lcd_string_u8(
        //     padding.left_pad_u8(&[0b11110100], "<").as_slice(),
        //     LinePlace::One,
        // );

        port.lcd_text_string(padding.right_pad("1", '0'), LinePlace::Two);
        moving_text.move_one(&port);
        port.lcd_sleep(1000);

        port.lcd_text_string(padding.left_pad("1", '2'), LinePlace::Two);
        moving_text.move_one(&port);
        // port.lcd_string_u8(
        //     padding.right_pad_u8(&[0b11110100], "<").as_slice(),
        //     LinePlace::Two,
        // );
        port.lcd_sleep(1000);
        println!("Loop done")
    }
}

fn setup(port: &I2CPort) {
    println!("Setup");
    port.lcd_set_mode_bytes(ModLines::Two, ModBytes::Four);
    port.lcd_set_mode_bytes(ModLines::Two, ModBytes::Four);
    port.lcd_display_on();
    port.lcd_first_line_setup();
    port.lcd_set_mode_bytes(ModLines::Two, ModBytes::Four);
    port.lcd_clear();
    println!("Setup done");
}

fn main() {
    // https://www.electronicsforu.com/technology-trends/learn-electronics/16x2-lcd-pinout-diagram
    let port = init_i2c();
    setup(&port);
    port.lcd_sleep_init();
    logic(port);
}
