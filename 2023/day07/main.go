package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/MueR/adventofcode.go/cast"
)

var (
	//go:embed input.txt
	input string
)

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 0, "part 1 or 2")
	flag.Parse()

	s := time.Now()
	if part != 2 {
		ans := part1(input)
		fmt.Printf("Part 1 output: %v  (%v)\n", ans, time.Since(s))
	}
	s = time.Now()
	if part != 1 {
		ans := part2(input)
		fmt.Printf("Part 2 output: %v  (%v)\n", ans, time.Since(s))
	}
}

func part1(input string) int {
	cardOrder := strings.Split("23456789TJQKA", "")
	hands := parseInput(input, false)
	slices.SortFunc(hands, func(a, b *Hand) int {
		return a.Compare(b, cardOrder)
	})
	score := 0
	for i := range hands {
		score += (i + 1) * hands[i].Bid
	}
	return score
}

func part2(input string) int {
	cardOrder := strings.Split("J23456789TQKA", "")
	hands := parseInput(input, true)
	slices.SortFunc(hands, func(a, b *Hand) int {
		return a.Compare(b, cardOrder)
	})
	score := 0
	for i := range hands {
		// fmt.Printf("[%-15s] %v (%v) %4d * %4d = %6d [Total: %9d]\n", handType(hands[i]), hands[i].String(), hands[i].Sorted, i+1, hands[i].Bid, (i+1)*hands[i].Bid, score+(i+1)*hands[i].Bid)
		score += (i + 1) * hands[i].Bid
	}
	return score
}

func parseInput(input string, joker bool) (ans []*Hand) {
	for _, line := range strings.Split(input, "\n") {
		cl := strings.Fields(line)
		hand := &Hand{
			Cards: strings.Split(cl[0], ""),
			Bid:   cast.ToInt(cl[1]),
		}
		hand.DetermineType(joker)
		ans = append(ans, hand)
	}
	return ans
}

type Hand struct {
	Cards  []string
	Sorted []string
	Bid    int
	Type   int
}

func (h *Hand) Compare(other *Hand, cardOrder []string) int {
	if h.Type > other.Type {
		return 1
	} else if h.Type < other.Type {
		return -1
	}
	for i := range h.Cards {
		hi := slices.Index(cardOrder, h.Cards[i])
		oi := slices.Index(cardOrder, other.Cards[i])
		if hi > oi {
			return 1
		}
		if hi < oi {
			return -1
		}
	}
	return 0
}

func (h *Hand) SortCards(cardOrder []string) string {
	sort.Slice(h.Cards, func(i, j int) bool {
		return slices.Index(cardOrder, h.Cards[i]) < slices.Index(cardOrder, h.Cards[j])
	})
	return strings.Join(h.Cards, "")
}

func (h *Hand) String() string {
	return strings.Join(h.Cards, "")
}

func (h *Hand) DetermineType(joker bool) *Hand {
	h.Type = 0
	cards := make(map[string]int)
	sorted := make([]string, 0)
	for _, c := range h.Cards {
		sorted = append(sorted, c)
		cards[string(c)]++
	}
	sort.Strings(sorted)
	h.Sorted = sorted
	if sorted[0] == sorted[4] {
		h.Type = 6 // Five of a kind
	} else if sorted[0] == sorted[3] || sorted[1] == sorted[4] {
		h.Type = 5 // Four of a kind
	} else if (sorted[0] == sorted[2] && sorted[3] == sorted[4]) || (sorted[0] == sorted[1] && sorted[2] == sorted[4]) {
		h.Type = 4 // Full house
	} else if sorted[0] == sorted[2] || sorted[1] == sorted[3] || sorted[2] == sorted[4] {
		h.Type = 3 // Three of a kind
	} else if len(cards) == 3 {
		h.Type = 2 // Two pairs
	} else if len(cards) == 4 {
		h.Type = 1 // One pair
	} else {
		h.Type = 0 // High card
	}

	if !joker || cards["J"] == 0 {
		return h
	}

	switch h.Type {
	case 6:
		return h // Five of a kind with joker is still a five of a kind
	case 5:
		h.Type = 6 // Four of a kind with joker is a five of a kind
	case 4:
		if cards["J"] == 2 {
			h.Type = 6 // Full house with two jokers is a five of a kind
		} else {
			h.Type = 5 // Full house with one joker is a four of a kind
		}
	case 3:
		h.Type = 5
	case 2:
		if cards["J"] == 2 {
			h.Type = 5 // Two pairs with two jokers is a four of a kind
		} else {
			h.Type = 4 // Two pairs with one joker is a full house
		}
	case 1:
		h.Type = 3 // One pair with joker is a three of a kind
	case 0:
		h.Type = 1 // High card with joker is a one pair
	}
	return h
}

func handType(h *Hand) string {
	switch h.Type {
	case 6:
		return "Five of a kind"
	case 5:
		return "Four of a kind"
	case 4:
		return "Full house"
	case 3:
		return "Three of a kind"
	case 2:
		return "Two pairs"
	case 1:
		return "One pair"
	default:
		return "High card"
	}
}
