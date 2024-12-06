package main

import (
	"testing"
)

func TestGetAnswer(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// // TODO: Add test cases.
		// {
		// 	name: "simple",
		// 	args: args{
		// 		line: "mul(11,8)",
		// 	},
		// 	want: 88,
		// },
		// {
		// 	name: "double call",
		// 	args: args{
		// 		line: "mul(11,8)mul(8,5))",
		// 	},
		// 	want: 88 + 40,
		// },
		{
			name: "double with noise",
			args: args{
				line: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			},
			want: 161,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAnswer(tt.args.line); got != tt.want {
				t.Errorf("ParseMul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
