package main

import (
	"testing"
)

func TestWithin8Chars(t *testing.T) {
	ch := NewCustomChar([]int{
		0b00001,
		0b00011,
		0b00111,
		0b01111,
		0b11111,
	})

	_, splitted := ch.SplitBySections()

	if splitted[0][4] != 0b11111 {
		t.Fatalf("Expected %d, got %d", 0b11111, splitted[0][4])
	}
}

func TestWithin10Chars(t *testing.T) {
	ch := NewCustomChar([]int{
		0b1001100001,
		0b1011000011,
		0b1011000111,
		0b1111001111,
		0b1000011111,
	})
	_, splitted := ch.SplitBySections()

	if splitted[0][4] != 0b11111 {
		t.Fatalf("Expected %d, got %d", 0b11111, splitted[0][4])
	}
	if splitted[1][0] != 0b10011 {
		t.Fatalf("Expected %d, got %d", 0b10011, splitted[1][0])
	}
}

func TestWithin11Chars(t *testing.T) {
	ch := NewCustomChar([]int{
		0b11001100001,
		0b11011000011,
		0b11011000111,
		0b11111001111,
		0b11000011111,
	})

	_, splitted := ch.SplitBySections()
	if splitted[2][0] != 0b1 {
		t.Fatalf("Expected %d, got %d", 0b1, splitted[2][0])
	}
}

func TestWithin11Chars2Height(t *testing.T) {
	ch := NewCustomChar([]int{
		0b00001,
		0b00011,
		0b00111,
		0b01111,
		0b11111,
		0b00001,
		0b00011,
		0b00111,
		0b11111,
	})

	_, splitted := ch.SplitBySections()
	if splitted[1][0] != 0b11111 {
		t.Fatalf("Expected %d, got %d", 0b11111, splitted[0][7])
	}
}

func TestSymbol(t *testing.T) {
	ch := NewCustomChar([]int{
		0b000111111110000,
		0b001000010001000,
		0b011000010000100,
		0b010100010001010,
		0b010010010010010,
		0b010001010100010,
		0b010000111000010,
		0b011111111111110,
		0b010000111000010,
		0b010001010100010,
		0b010010010010010,
		0b010010010010010,
		0b010010010001010,
		0b010100010000100,
		0b001000010001000,
		0b000111111110000,
	})

	_, splitted := ch.SplitBySections()

	if splitted[0][0] != 0b10000 {
		t.Fatalf("Expected %d, got %d", 0b10000, splitted[0][0])
	}
	if splitted[2][0] != 0b11 {
		t.Fatalf("Expected %d, got %d", 0b11, splitted[2][0])
	}
	if splitted[3][7] != 0b10000 {
		t.Fatalf("Expected %d, got %d", 0b10000, splitted[3][7])
	}
}

func TestSymbol100000Times(t *testing.T) {
	ch := NewCustomChar([]int{
		0b000111111110000,
		0b001000010001000,
		0b011000010000100,
		0b010100010001010,
		0b010010010010010,
		0b010001010100010,
		0b010000111000010,
		0b011111111111110,
		0b010000111000010,
		0b010001010100010,
		0b010010010010010,
		0b010010010010010,
		0b010010010001010,
		0b010100010000100,
		0b001000010001000,
		0b000111111110000,
	})
	for i := 0; i < 100000; i++ {
		slicesNum, splitted := ch.SplitBySections()

		if slicesNum == 0 {
			t.Fatal("Expected more than 0 slices")
		}
		if splitted[0][0] != 0b10000 {
			t.Fatalf("Expected %d, got %d", 0b10000, splitted[0][0])
		}
		if splitted[2][0] != 0b11 {
			t.Fatalf("Expected %d, got %d", 0b11, splitted[2][0])
		}
		if splitted[3][7] != 0b10000 {
			t.Fatalf("Expected %d, got %d", 0b10000, splitted[3][7])
		}
	}
}
