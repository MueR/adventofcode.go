package main

import (
	"testing"
)

var example = `2333133121414131402`

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Example",
			input: example,
			want:  1928,
		},
		{
			name:  "Actual",
			input: input,
			want:  6283404590840,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			disk := parseInput(tt.input)
			//printDisk(disk)
			if got := part1(disk); got != tt.want {
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
			want:  2858,
		},
		{
			name:  "Actual",
			input: input,
			want:  6304576012713,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			disk := parseInput(tt.input)
			if got := part2(disk); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
