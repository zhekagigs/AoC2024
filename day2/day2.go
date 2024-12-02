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
		if is_safe(report) || problem_dampener(report) {
			total_safe_reports += 1
		} 
	}
	return total_safe_reports
}

func problem_dampener(a []int64) bool {
    for i := 0; i < len(a); i++ {
        temp := make([]int64, len(a))
        copy(temp, a)
        if is_safe(pkg.RemoveIndex(temp, i)) {
            return true
        }
    }
    return false
}

func is_safe(a []int64) bool {
	var order_func func(a, b int64) bool
	if a[0] < a[1] {
		order_func = is_asc
	} else {
		order_func = is_desc
	}
	for i := 0; i < len(a)-1; i++ {
		if is_in_limit(a[i], a[i+1]) && order_func(a[i], a[i+1]) {
			continue
		} else {
			return false
		}
	}
	return true
}

func is_in_limit(a, b int64) bool {
	return pkg.Abs(a-b) <= 3
}

func is_asc(a, b int64) bool {
	return a < b
}

func is_desc(a, b int64) bool {
	return a > b
}
