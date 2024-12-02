package pkg

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	type args struct {
		fp string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				fp: "test1.txt",
			},
			want:    []string{"carl", "clary", "coraly"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFile(tt.args.fp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Add(t *testing.T) {
	type args struct {
		nums []int64
	}
	tests := []struct {
		name     string
		set      Set
		args     args
		expected []int64
	}{
		// TODO: Add test cases.
		{
			name:     "Test1",
			set:      make(Set),
			args:     args{nums: []int64{1, 2, 3}},
			expected: []int64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.args.nums {
				tt.set.Add(v)
			}
			keys := make([]int64, 0, len(tt.set))
			for k := range tt.set {
				keys = append(keys, k)
			}
			if !reflect.DeepEqual(tt.expected, keys) {
				t.Errorf("Set.Add() got [%v], want [%v]", keys, tt.expected)
			}

		})
	}
}
