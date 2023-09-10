#[cxx::bridge]
#[allow(dead_code)]
mod ffi {
    #[repr(u32)]
    enum LCDMode {
        OUTPUT = 1,
    }

    #[repr(u32)]
    enum PinValue {
        LOW = 0,
        HIGH = 1,
    }

    struct Pin {
        pub index: u32,
        pub mode: LCDMode,
    }

    #[allow(non_snake_case)]
    struct Pins {
        pub is4PinMode: bool,
        pub RS: Pin,
        pub RW: Pin,
        pub E: Pin,
        pub D0: Pin,
        pub D1: Pin,
        pub D2: Pin,
        pub D3: Pin,
        pub D4: Pin,
        pub D5: Pin,
        pub D6: Pin,
        pub D7: Pin,
        pub LCD_DELAY_MS: u32,
    }

    unsafe extern "C++" {

        include!("rs/lib/lcd.h");

        type Pin;
        type Pins;

        #[rust_name = "switch_to_command"]
        fn switchToCommand();

        #[rust_name = "switch_to_char"]
        fn switchToChar();

        #[rust_name = "read_lcd"]
        fn readLcd();
        #[rust_name = "lcd_command"]
        fn lcdCommand(command: u32);
        #[rust_name = "lcd_string"]
        fn lcdString(text: &CxxString);
        #[rust_name = "read_mode_lcd"]
        fn readModeLcd();
        #[rust_name = "write_mode_lcd"]
        fn writeModeLcd();
        #[rust_name = "is_4bit_mode"]
        fn is4bitMode() -> bool;

        #[rust_name = "init_lcd"]
        fn initLcd(initPins: Pins) -> u32;

        #[rust_name = "digital_write"]
        fn lcdDigitalWrite(pin: u32, value: u32);
    }
}

pub use self::ffi::{
    digital_write, init_lcd, is_4bit_mode, lcd_command, read_lcd, read_mode_lcd, switch_to_char,
    switch_to_command, write_mode_lcd, LCDMode, Pin, PinValue, Pins,
};

pub fn lcd_string(str: &str) {
    cxx::let_cxx_string!(txt = str);

    ffi::lcd_string(&txt)
}

pub fn lcd_commands(commands: &[u32]) {
    for cmd in commands {
        lcd_command(*cmd);
    }
}
