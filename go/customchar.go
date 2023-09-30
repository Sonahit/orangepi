package main

// byte customChar[8] = {
// 	0b11111,
// 	0b11111,
// 	0b11111,
// 	0b11111,
// 	0b11111,
// 	0b11111,
// 	0b11111,
// 	0b11111
// };

// Чтобы сделать custom-char
// то нужно заполнить матрицу из 5x8 из 1 для активации

type CustomChar struct {
	Rows []int
}

func NewCustomChar(rows []int) CustomChar {
	return CustomChar{Rows: rows}
}
