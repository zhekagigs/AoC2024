package main

import (
	"AoC2024/pkg"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	lines := pkg.GetLinesFromArgFile()
	blob := strings.Join(lines, "")
	re := regexp.MustCompile(`mul\([mul\(0-9\,\)]+\,[mul\(0-9\,\)]+\)`)
	match := re.FindAllIndex([]byte(blob), -1)
	totalSum := 0
	for _, points := range match {
		// var mulCount MulCount
		left := points[0]
		right := points[1]
		line := blob[left:right]
		totalSum += ParseMul(line)
		fmt.Println(line, totalSum)

	}
	fmt.Println("total sum:", totalSum)
}

func ParseMul(line string) int {
	var mc = &MulCount{
		opened:   0,
		closed:   0,
		pairs:    0,
		digit:    0,
		comma:    0,
		broken:   false,
		cutoff:   0,
		leftNum:  "",
		rightNum: "",
	}
	var sum int
	for _, v := range line {
		if v == '(' {
			mc.opened += 1
			continue
		}
		if unicode.IsDigit(v) && mc.comma == 0 && mc.opened > 0 && mc.opened > mc.closed {
			mc.leftNum += string(v)
		}
		if v == ',' {
			mc.comma += 1
		}
		if unicode.IsDigit(v) && mc.comma == 1 && mc.opened > 0 && mc.opened > mc.closed {
			mc.rightNum += string(v)
		}
		if v == ')' {
			mc.closed += 1
			if mc.opened == mc.closed {
				if mc.leftNum == "" || mc.rightNum == "" {
					fmt.Println(fmt.Errorf("no nums in line [%v] mulcount [%-v]", line, mc))
					mc.Done()
					continue
				}
				sum += pkg.ToInt(mc.leftNum) + pkg.ToInt(mc.rightNum)
				mc.Done()
			}
		}
	}
	return sum
}

type MulCount struct {
	opened   int
	closed   int
	pairs    int
	digit    int
	comma    int
	broken   bool
	cutoff   int
	leftNum  string
	rightNum string
}

func (m *MulCount) Done() {
	m.opened = 0
	m.closed = 0
	m.comma = 0
	m.leftNum = ""
	m.rightNum = ""
}
