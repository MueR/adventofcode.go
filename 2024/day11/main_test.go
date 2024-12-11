package main

import (
	"testing"
)

var example = `125 17`

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		steps int
		want  int
	}{
		{
			name:  "Example",
			input: example,
			steps: 25,
			want:  55312,
		},
		{
			name:  "Actual",
			input: input,
			steps: 25,
			want:  199982,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = parseInput(tt.input)
			if got := part1(tt.input, tt.steps); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		steps int
		want  int
	}{
		// No sample result for part 2 today
		// {
		// 	name:  "Example",
		// 	input: example,
		// 	want:  0,
		// },
		{
			name:  "Actual",
			input: input,
			steps: 75,
			want:  237149922829154,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = parseInput(tt.input)
			if got := part2(tt.input, tt.steps); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
