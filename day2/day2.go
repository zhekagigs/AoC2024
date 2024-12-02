package main

import (
	"AoC2024/pkg"
	"flag"
	"fmt"
)

const MAX_DIFF = 3

func main() {
	flag.Parse()
	fp := flag.Arg(0)
	lines, err := pkg.ReadFile(fp)
	pkg.Check(err)
	total_safe_reports := calcSafeReports(lines)
	fmt.Println("total safe reports: ", total_safe_reports)
}

func calcSafeReports(lines []string) int {
	total_safe_reports := 0
	for _, line := range lines {
		report := pkg.ToIntArr(line)
		if isSafe(report) || problemDampener(report) {
			total_safe_reports += 1
		}
	}
	return total_safe_reports
}

func problemDampener(a []int64) bool {
	for i := 0; i < len(a); i++ {
		temp := make([]int64, len(a))
		copy(temp, a)
		if isSafe(pkg.RemoveIndex(temp, i)) {
			return true
		}
	}
	return false
}

func isSafe(a []int64) bool {
	var isInOrder func(a, b int64) bool
	if isAsc(a[0], a[1]) {
		isInOrder = isAsc
	} else {
		isInOrder = isDesc
	}
	for i := 0; i < len(a)-1; i++ {
		if isInLimit(a[i], a[i+1]) && isInOrder(a[i], a[i+1]) {
			continue
		} else {
			return false
		}
	}
	return true
}

func isInLimit(a, b int64) bool {
	return pkg.Abs(a-b) <= 3
}

func isAsc(a, b int64) bool {
	return a < b
}

func isDesc(a, b int64) bool {
	return a > b
}
