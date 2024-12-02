package pkg

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Set map[int64]struct{}

func (s Set) Add(num int64) {
	s[num] = struct{}{}
}

func (s Set) Remove(item int64) {
	delete(s, item)
}

func (s Set) Has(item int64) bool {
	_, exists := s[item]
	return exists
}

func ReadFile(fp string) ([]string, error) {
	if fp == "" {
		return nil, fmt.Errorf("empty string as file path error")
	}
	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)

	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text, nil
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Toint64(num string) int64 {
	i, err := strconv.ParseInt(num, 10, 32)
	Check(err)
	return int64(i)
}

func RemoveIndex(s []int64, i int) []int64 {
    return append(s[:i], s[i+1:]...)
}

func ToIntArr(text string) []int64{
	var result []int64
	nums := strings.Split(text, " ")
	for _, num := range nums {
		result = append(result, Toint64(num))
	}
	return result

}

func SortAscending(nums []int64) []int64 {
	slices.Sort(nums)
	return nums
}

func ToSlice(set Set) []int64 {
	var slice []int64
	for k := range set {
		slice = append(slice, k)
	}
	return slice
}

func Abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func GetSmallest(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

