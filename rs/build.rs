fn main() {
    cxx_build::bridge("src/wiring_pi.rs") // returns a cc::Build
        .file("lib/lcd.h")
        .file("lib/lcd.cpp")
        .flag_if_supported("-std=c++14")
        .compile("rs");

    println!("cargo:rustc-link-lib=wiringPi");
    println!("cargo:rerun-if-changed=src/wiring_pi.rs");
    println!("cargo:rerun-if-changed=lib/lcd.cpp");
    println!("cargo:rerun-if-changed=lib/lcd.h");
}
