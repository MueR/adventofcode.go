package main

import (
	"testing"
)

var example = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Example",
			input: example,
			want:  2,
		},
		{
			name:  "Actual",
			input: input,
			want:  524,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list = parseInput(tt.input)
			if got := part1(); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Example",
			input: example,
			want:  4,
		},
		{
			name:  "Actual",
			input: input,
			want:  569,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list = parseInput(tt.input)
			if got := part2(); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
