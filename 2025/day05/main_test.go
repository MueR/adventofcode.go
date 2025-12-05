package main

import (
	"testing"
)

var example = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int64
	}{
		{
			name:  "Example",
			input: example,
			want:  3,
		},
		//{
		//	name:  "Actual",
		//	input: input,
		//	want:  1502,
		//},
	}
	render = false
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve = parseInput(tt.input)
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
			want:  14,
		},
		//{
		//	name:  "Actual",
		//	input: input,
		//	want:  9083,
		//},
	}
	render = false
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve = parseInput(tt.input)
			if got := part2(); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
