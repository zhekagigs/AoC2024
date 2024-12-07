package main

import (
	"AoC2024/pkg"
	"fmt"
)

const XMAS = "XMAS"

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
	totalSum := getAnswer(lines)
	fmt.Println("total sum:", totalSum) // total sum must be : 174561379
}

func getAnswer(lines []string) int {
	var totalSum int
	b := NewBoard(lines)
	if b == nil {
		return 0
	}
	for i := range b.board {
		for j, value := range b.board[i] {
			b.cursor.x, b.cursor.y = j, i
			if value == 'X' {
				totalSum = tryDirection(b.moveCursorRight, totalSum, b, j, i)
				totalSum = tryDirection(b.moveCursorLeft, totalSum, b, j, i)
				totalSum = tryDirection(b.moveCursorUp, totalSum, b, j, i)
				totalSum = tryDirection(b.moveCursorDown, totalSum, b, j, i)
				totalSum = tryDirection(b.moveCursorUpLeft, totalSum, b, j, i)
				totalSum = tryDirection(b.moveCursorUpRight, totalSum, b, j, i)
				totalSum = tryDirection(b.moveCursorDownLeft, totalSum, b, j, i)
				totalSum = tryDirection(b.moveCursorDownRight, totalSum, b, j, i)
			}
		}
	}
	return totalSum
}

func tryDirection(move func(), totalSum int, b *Board, j, i int) int {
	b.scanXmas('X', move)
	if word.accumulator == XMAS {
		
		totalSum += 1
		fmt.Println(totalSum, word.accumulator, i, j)
	}
	resetCurWordState(b, j, i)
	return totalSum
}

func resetCurWordState(b *Board, j int, i int) {
	word.Reset()
	b.cursor.x, b.cursor.y = j, i
}

func (b *Board) scanXmas(nextRune rune, move func()) {
	move()
	if word.Continue(nextRune) {
		nextRune = b.getValue(b.cursor)
		if nextRune == -1 {
			return
		}
		b.scanXmas(nextRune, move)
	}
	b.cursor.x -= 1
}

func (b *Board) moveCursorRight() {
	b.cursor.x += 1
}
func (b *Board) moveCursorLeft() {
	b.cursor.x -= 1
}

func (b *Board) moveCursorUp() {
	b.cursor.y -= 1
}
func (b *Board) moveCursorDown() {
	b.cursor.y += 1
}
func (b *Board) moveCursorUpLeft() {
	b.cursor.x -= 1
	b.cursor.y -= 1
}

func (b *Board) moveCursorUpRight() {
	b.cursor.x += 1
	b.cursor.y -= 1
}

func (b *Board) moveCursorDownRight() {
	b.cursor.x += 1
	b.cursor.y += 1
}

func (b *Board) moveCursorDownLeft() {
	b.cursor.x -= 1
	b.cursor.y += 1
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
	if w.index >= len(w.origin) - 1 {
		fmt.Printf("error in word %-v\n", w)

		return false
	}
	if rune(w.origin[w.index]) == r {
		w.index += 1
		return true
	} else {
		return false
	}
}


