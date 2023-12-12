package util

import (
	"strconv"
	"strings"
)

func ParseIntList(s, sep string) (list []int) {
	for _, line := range strings.Split(s, sep) {
		i, _ := strconv.Atoi(line)
		list = append(list, i)
	}
	return list
}

func ParseIntMap(s, sep string) (m map[int]int) {
	for i, line := range strings.Split(s, sep) {
		v, _ := strconv.Atoi(line)
		m[i] = v
	}
	return m
}
