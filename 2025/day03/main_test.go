package main

import (
	"testing"
)

var example = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int64
	}{
		{
			name:  "Example",
			input: example,
			want:  357,
		},
		{
			name:  "Actual",
			input: input,
			want:  17445,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed = parseInput(tt.input)
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
		want  int64
	}{
		{
			name:  "Example",
			input: example,
			want:  3121910778619,
		},
		{
			name:  "Actual",
			input: input,
			want:  173229689350551,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed = parseInput(tt.input)
			if got := part2(); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
