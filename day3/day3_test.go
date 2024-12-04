package main

import "testing"

func TestParseMul(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "double call",
			args: args{
				line: "mul(11,8)mul(8,5))",
			},
			want: 19 + 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseMul(tt.args.line); got != tt.want {
				t.Errorf("ParseMul() = %v, want %v", got, tt.want)
			}
		})
	}
}
