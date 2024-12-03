package util

import (
	"bufio"
	"io"
	"iter"
	"strconv"
	"strings"
)

func LineToInts(s string) (line []int) {
	for _, i := range strings.Fields(s) {
		n, _ := strconv.Atoi(i)
		line = append(line, n)
	}
	return line
}

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

func SectionsOf(input io.Reader, delim string) iter.Seq[string] {
	scan := bufio.NewScanner(input)
	scan.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := strings.Index(string(data), delim); i >= 0 {
			return i + len(delim), data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	})

	return func(yield func(string) bool) {
		for scan.Scan() {
			if err := scan.Err(); err != nil && err != io.EOF {
				panic(err)
			}
			token := scan.Text()
			if token == "" {
				continue
			}
			if !yield(token) {
				break
			}
		}
	}
}
