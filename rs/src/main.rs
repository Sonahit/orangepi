mod wiring_pi;
use wiring_pi::ffi as wp;

fn main() {
    let status = wp::init_lcd(wp::Pins {
        D0: wp::Pin {
            mode: wp::LCDMode::OUTPUT,
            index: 0,
        },
        D1: wp::Pin {
            mode: wp::LCDMode::OUTPUT,
            index: 0,
        },
        D2: wp::Pin {
            mode: wp::LCDMode::OUTPUT,
            index: 0,
        },
        D3: wp::Pin {
            mode: wp::LCDMode::OUTPUT,
            index: 0,
        },
        D4: wp::Pin {
            mode: wp::LCDMode::OUTPUT,
            index: 0,
        },
        D5: wp::Pin {
            mode: wp::LCDMode::OUTPUT,
            index: 0,
        },
        D6: wp::Pin {
            mode: wp::LCDMode::OUTPUT,
            index: 0,
        },
        D7: wp::Pin {
            mode: wp::LCDMode::OUTPUT,
            index: 0,
        },
        E: wp::Pin {
            mode: wp::LCDMode::OUTPUT,
            index: 0,
        },
        RS: wp::Pin {
            mode: wp::LCDMode::OUTPUT,
            index: 0,
        },
        RW: wp::Pin {
            mode: wp::LCDMode::OUTPUT,
            index: 0,
        },
        is4PinMode: true,
    });
    println!("Hello, world! {}", status);
}
