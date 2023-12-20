package main

import (
	"testing"
)

var example = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Example",
			input: example,
			want:  32000000,
		},
		{
			name: "Example2",
			input: `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`,
			want: 11687500,
		},
		// {
		// 	name:  "Actual",
		// 	input: input,
		// 	want:  867118762,
		// },
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
	}{
		// {
		// 	name:  "Example",
		// 	input: example,
		// 	want:  0,
		// },
		{
			name:  "Actual",
			input: input,
			want:  217317393039529,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
