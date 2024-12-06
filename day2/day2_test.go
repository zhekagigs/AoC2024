package main

import (
	"AoC2024/pkg"
	"reflect"
	"testing"
)

func Test_calcSafeReports(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  int
	}{
		{
			name: "All Safe",
			lines: []string{
				"1 2 3 4",
				"4 3 2 1",
				"1 1 1 1",
			},
			want: 3,
		},
		{
			name: "Some Safe, Some Not",
			lines: []string{
				"1 2 3 4",
				"4 8 2 1", // Unsafe
				"1 1 1 1",
				"1 2 5 4", // Unsafe
			},
			want: 2,
		},
		{
			name:  "Empty Input",
			lines: []string{},
			want:  0,
		},
		{
			name: "Safe after dampener",
			lines: []string{
				"1 2 7 4",
			},
			want: 1,
		},
		{
			name: "Unsafe after dampener",
			lines: []string{
				"1 7 2 4",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcSafeReports(tt.lines); got != tt.want {
				t.Errorf("calcSafeReports() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_is_safe(t *testing.T) {
	tests := []struct {
		name string
		a    []int64
		want bool
	}{
		{
			name: "Ascending Safe",
			a:    []int64{1, 2, 3, 4},
			want: true,
		},
		{
			name: "Descending Safe",
			a:    []int64{4, 3, 2, 1},
			want: true,
		},
		{
			name: "Constant Safe",
			a:    []int64{1, 1, 1, 1},
			want: false,
		},

		{
			name: "Unsafe Ascending",
			a:    []int64{1, 2, 7, 4},
			want: false,
		},
		{
			name: "Unsafe Descending",
			a:    []int64{7, 4, 1, 0},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSafe(tt.a); got != tt.want {
				t.Errorf("is_safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_problem_dampener(t *testing.T) {
	tests := []struct {
		name string
		a    []int64
		want bool
	}{
		{"safe_after_remove", []int64{1, 2, 7, 4}, true},
		{"unsafe_after_remove", []int64{1, 7, 2, 4}, false},
		{"empty", []int64{}, true},           // Empty should probably be considered safe
		{"single_element", []int64{1}, true}, // Single element should be safe
		{"all_safe", []int64{1, 2, 3}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problemDampener(tt.a); got != tt.want {
				t.Errorf("problem_dampener() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveIndex(t *testing.T) {
	tests := []struct {
		name string
		s    []int64
		i    int
		want []int64
	}{
		{"remove_middle", []int64{1, 2, 3, 4}, 2, []int64{1, 2, 4}},
		{"remove_first", []int64{1, 2, 3, 4}, 0, []int64{2, 3, 4}},
		{"remove_last", []int64{1, 2, 3, 4}, 3, []int64{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkg.RemoveIndex(tt.s, tt.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
