pub struct Padding(pub usize);

impl Padding {
    pub fn width(&self) -> usize {
        self.0
    }

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

    pub fn left_pad(&self, str: &str, pad: &str) -> String {
        let pad_str = Self::get_pad_str(self.width() - str.len(), pad.as_bytes());

        let mut new_str = String::with_capacity(self.width());

        new_str.push_str(str);
        new_str.push_str(&pad_str);
        new_str
    }

    pub fn right_pad(&self, str: &str, pad: &str) -> String {
        let pad_str = Self::get_pad_str(self.width() - str.len(), pad.as_bytes());

        let mut new_str = String::with_capacity(self.width());

        new_str.push_str(&pad_str);
        new_str.push_str(str);
        new_str
    }

    pub fn left_pad_u8(&self, str: &[u8], pad: &str) -> String {
        let pad_str = Self::get_pad_str(self.width() - str.len(), pad.as_bytes());

        String::from_utf8([str, pad_str.as_bytes()].concat()).unwrap()
    }

    pub fn right_pad_u8(&self, str: &[u8], pad: &str) -> String {
        let pad_str = Self::get_pad_str(self.width() - str.len(), pad.as_bytes());

        String::from_utf8([pad_str.as_bytes(), str].concat()).unwrap()
    }
}
