package main

import (
	"strings"
)

func LeftPad(str string, pad rune, length int) string {
	return str + strings.Repeat(string(pad), length-len(str))
}

func RightPad(str string, pad rune, length int) string {
	return strings.Repeat(string(pad), length-len(str)) + str
}
