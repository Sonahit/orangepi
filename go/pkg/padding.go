package pkg

import (
	"strings"
)

func LeftPad(str string, pad rune, length int) string {
	if length-len(str) <= 0 {
		return str
	}
	return str + strings.Repeat(string(pad), length-len(str))
}

func RightPad(str string, pad rune, length int) string {
	if length-len(str) <= 0 {
		return str
	}
	return strings.Repeat(string(pad), length-len(str)) + str
}
