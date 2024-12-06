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
	// totalSum := getAnswer(blob)
	totalSum := getAnswerDo(blob)
	fmt.Println("total sum:", totalSum) // total sum must be : 174561379
}

func getAnswer(blob string) int {
	var totalSum int
	reMul := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	for i := 0; i < len(blob); i++ {
		if i+4 < len(blob) && string(blob[i:i+4]) == "mul(" {
			if points := reMul.FindIndex([]byte(blob[i:])); points != nil && points[0] == 0 {
				absPoints := []int{i + points[0], i + points[1]}
				left := absPoints[0]
				right := absPoints[1]
				line := blob[left:right]

				partSum := ParseMul(line)
				totalSum += partSum
				// i += points[1] - 1  // without this totat is: 186843163
			}
		}
	}
	return totalSum
}


func getAnswerDo(blob string) int {
	var totalSum int
	reMul := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	var do = true
	for i := 0; i < len(blob); i++ {
		if i+4 < len(blob) && string(blob[i:i+4]) == "do()" {
			do = true
			i += 3
			continue
		}
		if i+7 < len(blob) && string(blob[i:i+7]) == "don't()" {
			do = false
			i += 6
			continue
		}
		if do && i+4 < len(blob) && string(blob[i:i+4]) == "mul(" {
			if points := reMul.FindIndex([]byte(blob[i:])); points != nil && points[0] == 0 {
				absPoints := []int{i + points[0], i + points[1]}
				left := absPoints[0]
				right := absPoints[1]
				line := blob[left:right]

				partSum := ParseMul(line)
				totalSum += partSum
				i += points[1] - 1  
			}
		}
	}
	return totalSum
}

// parses top level
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
					// fmt.Println(fmt.Errorf("no nums in line [%v] mulcount [%-v]", line, mc))
					mc.Done()
					continue
				}
				sum += pkg.ToInt(mc.leftNum) * pkg.ToInt(mc.rightNum)
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
