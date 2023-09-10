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
        pub LCD_DISPLAY_MS: u32,
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

impl Clone for Pin {
    fn clone(&self) -> Self {
        Self {
            index: self.index,
            mode: self.mode,
        }
    }
}

impl Clone for Pins {
    fn clone(&self) -> Self {
        Self {
            is4PinMode: self.is4PinMode,
            RS: self.RS.clone(),
            RW: self.RW.clone(),
            E: self.E.clone(),
            D0: self.D0.clone(),
            D1: self.D1.clone(),
            D2: self.D2.clone(),
            D3: self.D3.clone(),
            D4: self.D4.clone(),
            D5: self.D5.clone(),
            D6: self.D6.clone(),
            D7: self.D7.clone(),
            LCD_DISPLAY_MS: self.LCD_DISPLAY_MS,
        }
    }
}

impl Default for Pins {
    fn default() -> Self {
        Self {
            D7: self::Pin {
                mode: self::LCDMode::OUTPUT,
                index: 0,
            },
            D6: self::Pin {
                mode: self::LCDMode::OUTPUT,
                index: 1,
            },

            D5: self::Pin {
                mode: self::LCDMode::OUTPUT,
                index: 2,
            },
            D4: self::Pin {
                mode: self::LCDMode::OUTPUT,
                index: 5,
            },
            D3: self::Pin {
                mode: self::LCDMode::OUTPUT,
                index: 3,
            },
            D2: self::Pin {
                mode: self::LCDMode::OUTPUT,
                index: 17,
            },
            D1: self::Pin {
                mode: self::LCDMode::OUTPUT,
                index: 19,
            },
            D0: self::Pin {
                mode: self::LCDMode::OUTPUT,
                index: 20,
            },
            E: self::Pin {
                mode: self::LCDMode::OUTPUT,
                index: 22,
            },
            RW: self::Pin {
                mode: self::LCDMode::OUTPUT,
                index: 23,
            },
            RS: self::Pin {
                mode: self::LCDMode::OUTPUT,
                index: 25,
            },
            is4PinMode: false,
            LCD_DISPLAY_MS: std::option_env!("DELAY")
                .unwrap_or("100")
                .parse::<u32>()
                .unwrap(),
        }
    }
}
