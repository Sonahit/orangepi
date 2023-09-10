mod wiring_pi;

use wiring_pi as wp;

const PINS: wp::Pins = wp::Pins {
    D7: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 0,
    },
    D6: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 1,
    },

    D5: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 2,
    },
    D4: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 5,
    },
    D3: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 3,
    },
    D2: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 17,
    },
    D1: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 19,
    },
    D0: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 20,
    },
    E: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 22,
    },
    RW: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 23,
    },
    RS: wp::Pin {
        mode: wp::LCDMode::OUTPUT,
        index: 25,
    },
    is4PinMode: true,
};

fn setup() {
    println!("Setup");
    wp::digital_write(PINS.E.index, wp::PinValue::LOW.repr);
    wp::write_mode_lcd();
    println!("Setup done");
}

fn logic() {
    // https://www.sparkfun.com/datasheets/LCD/HD44780.pdf Table 12 4bit 8digit 1 line
    #[rustfmt::skip]
    wp::lcd_commands(
        &[
            0b0010, 
            0b0010, 
            0b1110,
            0b110
        ]
    );

    wp::lcd_string("hiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii");

    wp::lcd_command(0b10);

    // 2 lines, 5x7 format
    // wp::lcd_command(0b00111000);

    // display on
    // wp::lcd_command(0b00001110);

    // clear display (optional here)
    // wp::lcd_command(0x01);

    // first line
    // entrymode
    // wp::lcd_command(0b00000110);

    // wp::lcd_string("Hitachi");
    //nextline
    // wp::lcd_command(0b11000000);
    // wp::lcd_string("MICROCOM");

    // HOME
    // wp::lcd_command(0b10);
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
