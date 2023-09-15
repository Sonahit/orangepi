pub struct Padding(pub usize);

impl Padding {
    pub fn width(&self) -> usize {
        self.0
    }

    fn get_pad_str(size: usize, pad: char) -> String {
        let mut padded_str = String::with_capacity(size);

        for _ in 0..size {
            padded_str.push(pad);
        }

        padded_str
    }

    pub fn left_pad(&self, str: &str, pad: char) -> String {
        let pad_str = Self::get_pad_str(self.width() - str.len(), pad);

        let mut new_str = String::with_capacity(self.width());

        new_str.push_str(str);
        new_str.push_str(&pad_str);
        new_str
    }

    pub fn right_pad(&self, str: &str, pad: char) -> String {
        let pad_str = Self::get_pad_str(self.width() - str.len(), pad);

        let mut new_str = String::with_capacity(self.width());

        new_str.push_str(&pad_str);
        new_str.push_str(str);
        new_str
    }

    pub fn left_pad_u8(&self, str: &[u8], pad: char) -> Vec<u8> {
        let pad_str = Self::get_pad_str(self.width() - str.len(), pad);

        [str, pad_str.as_bytes()].concat()
    }

    pub fn right_pad_u8(&self, str: &[u8], pad: char) -> Vec<u8> {
        let pad_str = Self::get_pad_str(self.width() - str.len(), pad);

        [str, pad_str.as_bytes()].concat()
    }
}
