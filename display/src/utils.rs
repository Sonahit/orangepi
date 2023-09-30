pub struct Padding(pub usize);

impl Padding {
    pub fn width(&self) -> usize {
        self.0
    }

    pub fn left_pad(&self, str: &str, pad: char) -> String {
        let pad_str = get_pad_str(self.width(), str.len(), pad);

        let mut new_str = String::with_capacity(self.width());

        new_str.push_str(str);
        new_str.push_str(&pad_str);
        new_str
    }

    pub fn right_pad(&self, str: &str, pad: char) -> String {
        let pad_str = get_pad_str(self.width(), str.len(), pad);

        let mut new_str = String::with_capacity(self.width());

        new_str.push_str(&pad_str);
        new_str.push_str(str);
        new_str
    }

    pub fn left_pad_u8(&self, str: &[u8], pad: char) -> Vec<u8> {
        let pad_str = get_pad_str(self.width(), str.len(), pad);

        [str, pad_str.as_bytes()].concat()
    }

    pub fn right_pad_u8(&self, str: &[u8], pad: char) -> Vec<u8> {
        let pad_str = get_pad_str(self.width(), str.len(), pad);

        [str, pad_str.as_bytes()].concat()
    }
}

fn get_pad_str(max_size: usize, str_len: usize, pad: char) -> String {
    let size = if str_len >= max_size {
        0
    } else {
        max_size - str_len
    };

    let mut padded_str = String::with_capacity(size);
    for _ in 0..size {
        padded_str.push(pad);
    }

    padded_str
}

pub fn left_pad(str: &str, pad: char, length: usize) -> String {
    let pad_str = get_pad_str(length, str.len(), pad);

    let mut new_str = String::with_capacity(length);

    new_str.push_str(str);
    new_str.push_str(&pad_str);
    new_str
}

pub fn right_pad(str: &str, pad: char, length: usize) -> String {
    let pad_str = get_pad_str(length, str.len(), pad);

    let mut new_str = String::with_capacity(length);

    new_str.push_str(&pad_str);
    new_str.push_str(str);
    new_str
}
