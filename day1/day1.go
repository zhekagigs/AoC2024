package main

import (
	"AoC2024/pkg"
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()
	fp := flag.Arg(0)
	lines, err := pkg.ReadFile(fp)
	pkg.Check(err)

	// leftSide, rightSide := splitSides(lines)
	leftSide, rightSide := splitArrs(lines)
	a := pkg.SortAscending(leftSide)
	b := pkg.SortAscending(rightSide)

	fmt.Println("total_distance: ", calcTotalDistance(a, b))
	fmt.Println("total_similiarity: ", calcSimiliarity(a, b))
}

func calcSimiliarity(a, b []int64) int64 {
	var total_score int64
	for i := 0; i < len(a); i++ {
		var count int64 = 0
		for j := 0; j < len(b); j++ {
			if b[j] == a[i] {
				count += 1
			}
		}
		score := a[i] * count
		total_score += score
	}
	return total_score
}

func splitArrs(lines []string) ([]int64, []int64) {
	var leftSide, rightSide []int64
	for _, line := range lines {
		splits := strings.Split(line, "   ")
		a, b := splits[0], splits[1]
		leftSide = append(leftSide, pkg.Toint64(a))
		rightSide = append(rightSide, pkg.Toint64(b))
	}
	return leftSide, rightSide
}

func calcTotalDistance(a []int64, b []int64) int64 {
	var total_dist int64
	if len(a) != len(b) {
		panic(fmt.Errorf("inequal lenghts, a - %v, b - %v", a, b))
	}
	for i := 0; i < len(a); i++ {
		distance := pkg.Abs(a[i] - b[i])
		total_dist += distance
	}
	return total_dist
}
