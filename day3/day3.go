package main

import (
	"AoC2024/pkg"
	// "bytes"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	lines := pkg.GetLinesFromArgFile()
	blob := strings.Join(lines, "")
	// reMul := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	// reDo := regexp.MustCompile(`do()`)
	// reDont := regexp.MustCompile(`don't()`)
	// matchDo := reDo.FindAllIndex([]byte(blob), -1)
	// matchDont := reDont.FindAllIndex([]byte(blob), -1)
	// matches := reMul.FindAllIndex([]byte(blob), -1)
	totalSum := 0

	reMul := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	counter := 0
	for i := 0; i < len(blob); i++ {
		if i+7 < len(blob) && string(blob[i:i+4]) == "mul(" {
			if points := reMul.FindIndex([]byte(blob[i:])); points != nil {
				// Adjust points for the offset
				absPoints := []int{points[0] + i, points[1] + i}
				// fmt.Printf("i: [%v], points: [%v]\n", i, absPoints)
				tempsum := addToTotal(absPoints, blob, totalSum)
				totalSum += tempsum 
				fmt.Println(tempsum, totalSum)
				counter++
				if totalSum > 174561379 {
					fmt.Println(i, len(blob), counter, "the end")
					break
				}
			}
		}
	}
		fmt.Println("total sum:", totalSum) // total sum: 174561379
	}

func addToTotal(points []int, blob string, totalSum int) int {
	left := points[0]
	right := points[1]
	line := blob[left:right]
	fmt.Println( string(line))
	partSum := ParseMul(line)
	totalSum += partSum
	// fmt.Println(line, partSum, totalSum)
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
