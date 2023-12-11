package main

import (
	"testing"
)

var example = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Example",
			input: example,
			want:  374,
		},
		{
			name:  "Actual",
			input: input,
			want:  9686930,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
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
		empty int
	}{
		{
			name:  "Example",
			input: example,
			want:  1030,
			empty: 10,
		},
		{
			name:  "Example",
			input: example,
			want:  8410,
			empty: 100,
		},
		{
			name:  "Actual",
			input: input,
			want:  630728425490,
			empty: 1000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input, tt.empty); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
