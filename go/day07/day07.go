package day07

import (
	"aoc23/inputs"
	"math"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type hand struct {
	v        [13]byte
	original [5]byte
	sc       int
	bid      int
}

func (h *hand) score() {
	score := 0
	for i := 0; i < 13; i++ {
		num := h.v[i]
		if num != 0 {
			score += int(math.Pow(10, float64(num)))
		}
	}
	h.sc = score

}
func (h *hand) score2() {
	score := 0

	jokers := h.v[9]
	h.v[9] = 0

	if jokers == 5 {
		h.sc = int(math.Pow(10, float64(5)))
		h.v[9] = jokers
		return
	}

	max := byte(0)
	mostFreq := 0

	for i := 0; i < 13; i++ {
		if h.v[i] > max {
			max = byte(h.v[i])
			mostFreq = i
		}
	}

	for i := 12; i >= 0; i-- {
		num := h.v[i]
		if num != 0 {
			if i == mostFreq {
				num += jokers
			}
			score += int(math.Pow(10, float64(num)))
		}
	}
	h.v[9] = jokers
	h.sc = score
}

func (h *hand) wins(h2 *hand) int {
	if h.sc > h2.sc {
		return 1
	} else if h.sc < h2.sc {
		return -1
	}
	for i := 0; i < 5; i++ {
		if h.original[i] > h2.original[i] {
			return 1
		} else if h.original[i] < h2.original[i] {
			return -1
		}
	}
	return 0
}

func (h *hand) wins2(h2 *hand) int {
	if h.sc > h2.sc {
		return 1
	} else if h.sc < h2.sc {
		return -1
	}
	for i := 0; i < 5; i++ {
		o1 := h.original[i] + 1
		o2 := h2.original[i] + 1
		if o1 == 10 {
			o1 = 0
		}
		if o2 == 10 {
			o2 = 0
		}
		if o1 > o2 {
			return 1
		} else if o1 < o2 {
			return -1
		}
	}
	return 0
}

func Run() [2]int {
	hands := parse()
	p1 := part1(hands)
	p2 := part2(hands)
	return [2]int{p1, p2}
}

func part1(hands []hand) int {
	slices.SortFunc(hands, func(h1, h2 hand) int {
		return h1.wins(&h2)
	})
	finalScore := 0
	for i, h := range hands {
		finalScore += ((i + 1) * h.bid)
	}

	return finalScore
}

func part2(hands []hand) int {
	for i := range hands {
		hands[i].score2()
	}
	slices.SortFunc(hands, func(h1, h2 hand) int {
		return h1.wins2(&h2)
	})
	finalScore := 0
	for i, h := range hands {
		finalScore += ((i + 1) * h.bid)
	}

	return finalScore
}

func parse() []hand {
	lines := strings.Split(inputs.Input07, "\n")
	lines = lines[:len(lines)-1]
	hands := make([]hand, len(lines))
	for i, l := range lines {
		hands[i] = parseHand(l)
	}
	return hands
}

func parseHand(s string) hand {
	parts := strings.Split(s, " ")
	bid, _ := strconv.Atoi(parts[1])
	runes := []rune(parts[0])

	var values [13]byte
	orig := [5]byte{}
	for i, r := range runes {
		v := cardValueMap[r]
		orig[i] = v
		values[v]++
	}
	h := hand{v: values, bid: bid, original: orig}
	h.score()
	return h
}

var cardValueMap = map[rune]byte{
	'2': 0, '3': 1, '4': 2, '5': 3, '6': 4, '7': 5, '8': 6, '9': 7,
	'T': 8, 'J': 9, 'Q': 10, 'K': 11, 'A': 12,
}
