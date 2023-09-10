mod wiring_pi;

use wiring_pi as wp;

fn setup(pins: &wp::Pins) {
    println!("Setup");
    wp::digital_write(pins.E.index, wp::PinValue::LOW.repr);
    wp::write_mode_lcd();
    println!("Setup done");
}

fn logic() {
    // https://www.sparkfun.com/datasheets/LCD/HD44780.pdf Table 12
    #[rustfmt::skip]
    wp::lcd_commands(
        &[
            0b00110000, 
            0b00001110, 
            0b00000110
        ]
    );

    wp::lcd_string("Hitachi");

    wp::lcd_command(0b00000111);

    wp::lcd_string(" Mikrocom");

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
    let pins = wp::Pins::default();
    let status = wp::init_lcd(pins.clone());

    if status != 0 {
        panic!("WiringOP didnt start");
    }

    println!("Hello, world! {}", status);

    setup(&pins);
    logic();
}
