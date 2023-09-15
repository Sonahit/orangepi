fn get_pad_str(size: usize, pad: &[u8]) -> String {
    let mut padded_str = String::with_capacity(size);
    let mut pad_str = String::from_utf8(pad.into()).unwrap();
    for i in 0..size {
        if padded_str.len() + pad_str.len() > size {
            pad_str.shrink_to(size - padded_str.len());
            padded_str.push_str(&pad_str)
        } else {
            padded_str.push_str(&pad_str);
        }
    }

    padded_str
}

pub trait Padding<T: Sized> {
    fn left_pad(&self, pad: &str, pad_len: usize) -> String;
    fn right_pad(&self, pad: &str, pad_len: usize) -> String;
}

impl Padding<&[u8]> for &[u8] {
    fn left_pad(&self, pad: &str, pad_len: usize) -> String {
        let pad_str = get_pad_str(pad_len - self.len(), pad.as_bytes());

        String::from_utf8([self, pad_str.as_bytes()].concat()).unwrap()
    }

    fn right_pad(&self, pad: &str, pad_len: usize) -> String {
        let pad_str = get_pad_str(pad_len - self.len(), pad.as_bytes());

        String::from_utf8([pad_str.as_bytes(), self].concat()).unwrap()
    }
}

impl Padding<&[i32]> for &[i32] {
    fn left_pad(&self, pad: &str, pad_len: usize) -> String {
        let pad_str = get_pad_str(pad_len - self.len(), pad.as_bytes());
        let mut new_str = String::with_capacity(pad_len);

        for byte in *self {
            new_str.push(char::from_u32(byte.unsigned_abs()).unwrap())
        }

        [new_str, pad_str].concat()
    }

    fn right_pad(&self, pad: &str, pad_len: usize) -> String {
        let pad_str = get_pad_str(pad_len - self.len(), pad.as_bytes());

        let mut new_str = String::with_capacity(pad_len);

        for byte in *self {
            new_str.push(char::from_u32(byte.unsigned_abs()).unwrap())
        }

        [pad_str, new_str].concat()
    }
}

impl Padding<&str> for &str {
    fn left_pad(&self, pad: &str, pad_len: usize) -> String {
        let pad_str = get_pad_str(pad_len - self.len(), pad.as_bytes());

        [self, pad_str.as_str()].concat()
    }

    fn right_pad(&self, pad: &str, pad_len: usize) -> String {
        let pad_str = get_pad_str(pad_len - self.len(), pad.as_bytes());

        [pad_str.as_str(), self].concat()
    }
}
