mod wiring_pi;
use wiring_pi::ffi as wp;
const PINS: wp::Pins = wp::Pins {
    D0: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 20,
    },
    D1: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 19,
    },
    D2: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 17,
    },
    D3: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 3,
    },
    D4: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 5,
    },
    D5: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 2,
    },
    D6: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 1,
    },
    D7: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 0,
    },
    E: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 22,
    },
    RS: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 25,
    },
    RW: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 23,
    },
    is4PinMode: true,
};

fn setup() {
    println!("Setup");
    wp::digital_write(PINS.E.index, wp::PinValue::LOW.repr);
    wp::write_mode_lcd();
    // 4-bit mode, 2 lines, 5x7 format
    wp::lcd_command(0x38);

    // lcd on cursor blink
    wp::lcd_command(0x0e);
    // clear display (optional here)
    wp::lcd_command(0x01);
    println!("Setup done");
}

fn logic() {
    // clear display (optional here)
    wp::lcd_command(0x01);

    // move to first line
    wp::lcd_command(0x83);

    cxx::let_cxx_string!(txt = "plz hlep");

    wp::lcd_string(&txt);

    // move to second line
    wp::lcd_command(0xC0);
    cxx::let_cxx_string!(txt = "LCD directly! :)");
    wp::lcd_string(&txt);
}

fn main() {
    let status = wp::init_lcd(PINS);

    if status != 0 {
        panic!("WiringOP didnt start");
    }

    println!("Hello, world! {}", status);

    setup();
    logic();
}
