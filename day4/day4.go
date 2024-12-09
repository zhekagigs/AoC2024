package main

import (
	"AoC2024/pkg"
	"fmt"
)

var XMAS string = "MAS"

type Word struct {
	origin      string
	index       int
	accumulator string
}

var word *Word = &Word{
	origin:      XMAS,
	index:       0,
	accumulator: "",
}

type Position struct {
	x int
	y int
}

type Board struct {
	board  [][]rune
	cursor Position
	lenX   int
	lenY   int
}

func (w *Word) Reset() {
	w.index = 0
	w.accumulator = ""
}

func main() {
	lines := pkg.GetLinesFromArgFile()
	totalSum2 := getAnswer2(lines)
	fmt.Println("total sum:", totalSum2)
}

func getAnswer2(lines []string) int {
	var totalSum int
	b := NewBoard(lines)
	if b == nil {
		return 0
	}
	for y := range b.board {
		for x, value := range b.board[y] {
			b.cursor.x, b.cursor.y = x, y
			if value == 'A' {
				upLeft := b.getValue(Position{x: x - 1, y: y - 1})
				upRight := b.getValue(Position{x: x + 1, y: y - 1})
				downLeft := b.getValue(Position{x: x - 1, y: y + 1})
				downRight := b.getValue(Position{x: x + 1, y: y + 1})

				all := []rune{upLeft, upRight, downLeft, downRight}
				var m, s int
				for _, r := range all {
					if r == -1 {
						continue
					} else {
						if r == 'M' {
							m++
						}
						if r == 'S' {
							s++
						}
						if m == 2 && s == 2 {
							if upLeft != downRight && upRight != downLeft {
								totalSum++
							}
						}
					}
				}
			}
		}
	}
	return totalSum
}

func (b *Board) getValue(p Position) rune {
	if p.x < b.lenX && p.x > -1 && p.y > -1 && p.y < b.lenY {
		return b.board[p.y][p.x]
	} else {
		return -1
	}
}

// NewBoard creates a new Board instance from the given lines of input.
// The Board is initialized with the runes from the input lines, and the
// cursor is set to the top-left position (0, 0).
func NewBoard(lines []string) *Board {
	if len(lines) == 0 {
		return nil
	}
	var b *Board = &Board{
		board: make([][]rune, len(lines)),
		cursor: Position{
			x: 0,
			y: 0,
		},
	}
	for i, line := range lines {
		b.board[i] = make([]rune, len(line))
		for j, v := range line {
			b.board[i][j] = v
		}
	}
	b.lenY, b.lenX = len(b.board), len(b.board[0])
	return b
}

func (w *Word) Continue(r rune) bool {
	w.accumulator += string(r)
	if w.index >= len(w.origin)-1 {
		// fmt.Printf("error in word %-v\n", w)
		return false
	}
	if rune(w.origin[w.index]) == r {
		w.index += 1
		return true
	} else {
		return false
	}
}
