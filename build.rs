use std::path::{Path, PathBuf};

fn main() {
    let bridge_files = Path::new("./src/wiring_pi/")
        .read_dir()
        .unwrap()
        .filter_map(|v| {
            if let Ok(entry) = v {
                Some(entry.path())
            } else {
                None
            }
        })
        .collect::<Vec<PathBuf>>();

    let cpp_files = Path::new("./lib/")
        .read_dir()
        .unwrap()
        .filter_map(|v| {
            if let Ok(entry) = v {
                Some(entry.path())
            } else {
                None
            }
        })
        .collect::<Vec<PathBuf>>();

    cxx_build::bridges(&bridge_files)
        .files(&cpp_files)
        .flag_if_supported("-std=c++14")
        .compile("rs");

    println!("cargo:rustc-link-lib=wiringPi");

    let rerun_files = [cpp_files, bridge_files].concat();
    for entry in rerun_files {
        println!("cargo:rerun-if-changed={}", entry.to_string_lossy());
    }
}
