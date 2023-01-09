package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Println("Part One:", partOne())
	fmt.Println("Part Two:", partTwo())
}

func score(item rune) int {
	r := int(unicode.ToLower(item) - 96)
	if unicode.IsUpper(item) {
		// Because of course they can't just follow the order of the ascii table...
		r += 26
	}
	return r
}

func partOne() int {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var prioritySum int
	for sc.Scan() {
		items := make(map[rune]bool)
		text := sc.Text()
		// Get the contents by compartment, left first
		for _, left := range text[:len(text)/2] {
			items[left] = true
		}

		for _, right := range text[len(text)/2:] {
			if items[right] {
				prioritySum += score(right)
				break
			}
		}
	}

	return prioritySum
}

func partTwo() int {
	// This can probably be done smarter without reading the file twice
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var prioritySum int
	for sc.Scan() {
		firstElf := itemSet(sc.Text())
		sc.Scan()
		secondElf := itemSet(sc.Text())
		sc.Scan()
		thirdElf := itemSet(sc.Text())

		for item := range firstElf {
			if secondElf[item] && thirdElf[item] {
				prioritySum += score(item)
				// Every elf has item, abort.
				break
			}
		}
	}
	return prioritySum
}

func itemSet(items string) (set map[rune]bool) {
	set = make(map[rune]bool)
	for _, item := range items {
		set[item] = true
	}
	return
}
