package main

import (
	"AoC2024/pkg"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := pkg.GetLinesFromArgFile()
	totalSum1 := getAnswer1(lines)
	fmt.Println("total sum:", totalSum1)
}

type Rule struct {
	before, after string
}

type Update [][]string

type Network struct {
	rules map[string][]string // key - before, []string - after
}

func (n *Network) AddRule(rule Rule) {
	n.rules[rule.before] = append(n.rules[rule.before], rule.after)
}

func (n *Network) Check(u []string) bool {
	for i, num := range u {
		allNumsAfterNum := n.rules[num]
		leftovers := u[i+1:]
		for _, leftOverNum := range leftovers {
			if !slices.Contains(allNumsAfterNum, leftOverNum) {
				return false
			}
		}
	}
	return true
}

func (n *Network) isAfter(before, after string) bool {
	return slices.Contains(n.rules[before], after)
}

func (n *Network) SortByRules(u []string) []string {
	var result []string = u
	for i := 0; i < len(u) - 1; i++ {
		for j := 0; j < len(u) - i - 1; j++ {
			if !n.isAfter(u[j], u[j+1]) {
				u[j], u[j+1] = u[j+1], u[j]
			}
		}
	}
	return result
}

func getAnswer1(lines []string) int64 {
	var updates Update
	var net Network = Network{
		rules: make(map[string][]string),
	}
	for _, v := range lines {
		if strings.Contains(v, "|") {
			rule := Rule{
				before: strings.Split(v, "|")[0],
				after:  strings.Split(v, "|")[1],
			}
			net.AddRule(rule)
		} else if strings.Contains(v, ",") {
			vSlice := strings.Split(v, ",")
			updates = append(updates, vSlice)
		}
	}
	var sum int64 = 0
	for _, u := range updates {
		if net.Check(u) {
			continue
		} else {
			sortedU := net.SortByRules(u)
			num, err := strconv.ParseInt(sortedU[len(sortedU)/2], 10, 64)
			if err != nil {
				panic(err)
			}
			sum += num
		}
	}
	return sum
}
