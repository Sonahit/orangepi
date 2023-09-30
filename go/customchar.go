package main

import (
	"log"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

// byte customChar[8] = {
// 	0b100000000001,
// 	0b111111111111,
// 	0b000000000000,
// 	0b111111111111,
// 	0b111111111111,
// 	0b111111111111,
// 	0b000000000000,
// 	0b111111111111
// };

// Чтобы сделать custom-char
// то нужно заполнить матрицу из 5x8 из 1 для активации

type CustomChar struct {
	Rows []int
}

const (
	CHAR_HEIGHT = 8
	CHAR_LENGTH = 5
)

func NewCustomChar(rows []int) CustomChar {
	return CustomChar{
		Rows: rows,
	}
}

func printRows(rows []int) {
	for height := 0; height < len(rows); height += 1 {
		log.Println(strconv.FormatInt(int64(rows[height]), 2))
	}
}

func maxRowsLength(rows []int) int {
	return slices.MaxFunc[[]int](rows, func(a, b int) int {
		return strings.Compare(strconv.FormatInt(int64(a), 2), strconv.FormatInt(int64(b), 2))
	})
}

func rowsZero(rows []int) bool {
	return slices.Max[[]int](rows) == 0
}

func (ch CustomChar) SplitBySections() (int, [][]int) {
	copyMatrix := make([][]int, 0, len(ch.Rows))
	sections := make([][]int, 0, len(ch.Rows))
	if len(ch.Rows) > CHAR_HEIGHT {
		copyMatrix = append(copyMatrix, ch.Rows[:CHAR_HEIGHT], ch.Rows[CHAR_HEIGHT:len(ch.Rows)])
	} else {
		copyMatrix = append(copyMatrix, ch.Rows[:len(ch.Rows)])
	}

	for _, rows := range copyMatrix {
		for {
			if rowsZero(rows) {
				break
			}
			newRows := make([]int, len(rows))
			// firstChar from left
			// LE
			for height := 0; height < len(rows); height += 1 {
				row := 0
				if rows[height] != 0 {
					for length := 0; length < CHAR_LENGTH; length += 1 {
						bit := rows[height] >> length % 2
						row |= bit << length
					}
					rows[height] >>= CHAR_LENGTH
				}
				newRows[height] = int(row)
			}
			sections = append(sections, newRows)
		}
	}
	sectionsNumber := len(sections)
	printRows(sections[0])
	return sectionsNumber, sections
}
