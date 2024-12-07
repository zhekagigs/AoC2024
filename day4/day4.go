package main

import (
	"AoC2024/pkg"
	"fmt"
)

const XMAS = "XMAS"

const X = "X"
const M = "M"
const A = "A"
const S = "S"

var word *Word = &Word{
	origin: XMAS,
	index: 1,
	accumulator: X,
}

var totalSum int


func main() {
	lines := pkg.GetLinesFromArgFile()
	// blob := strings.Join(lines, "")
	totalSum := getAnswer(lines)
	fmt.Println("total sum:", totalSum) // total sum must be : 174561379
}

func getAnswer(lines []string) int {
	b := NewBoard(lines)

	for i, _ := range b.board {
		for j, value := range b.board[i] {
			b.cursor.x, b.cursor.y = j, i
			if value == 'X'{
				b.scanXmas('M')
				if word.accumulator == XMAS {
					totalSum += 1
				}
				word.Reset()
			}
		}
	}
	return -1
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

func (b *Board) getValue(p Position) rune {
	if p.x < b.lenX && p.y < b.lenY {
		return b.board[p.x][p.y]
	} else {
		return -1
	}
}

func NewBoard(lines []string) *Board {
	var b *Board
	b.cursor.x = 0
	b.cursor.y = 0
	var i, j int
	for i, line := range lines {
		for j, v := range line {
			b.board[j][i] = v
		}
	}
	b.lenX, b.lenY = j, i
	return b
}

type Word struct {
	origin string
	index int
	accumulator string
}

func (w *Word) Continue (r rune) bool {
	w.accumulator += string(r)
	if rune(w.origin[w.index]) == r {
		return true
	} else {
		return false
	}
}

func (w *Word) Reset() {
	w.index = 1
	w.accumulator = ""
}


func(b *Board) scanXmas(nextRune rune) {
	b.peekLeft()
	b.cursor.x += 1
	nextRune = b.getValue(b.cursor)
	if word.Continue(nextRune) {

		b.scanXmas(nextRune)
	} 
	b.cursor.x -= 1
}

type Direction struct {
    dx, dy int
}

func (b *Board) peek(d Direction) rune {
    return b.getValue(Position{
        x: b.cursor.x + d.dx,
        y: b.cursor.y + d.dy,
    })
}

func (b *Board) peekLeft() rune   { return b.peek(Direction{-1, 0}) }
func (b *Board) peekRight() rune  { return b.peek(Direction{1, 0}) }
func (b *Board) peekUp() rune     { return b.peek(Direction{0, 1}) }
func (b *Board) peekDown() rune   { return b.peek(Direction{0, -1}) }
func (b *Board) peekUpLeft() rune { return b.peek(Direction{-1, 1}) }
func (b *Board) peekUpRight() rune { return b.peek(Direction{1, 1}) }
func (b *Board) peekDownLeft() rune { return b.peek(Direction{-1, -1}) }
func (b *Board) peekDownRight() rune { return b.peek(Direction{1, -1}) }
