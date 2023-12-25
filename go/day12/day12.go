package day12

import (
	"aoc23/inputs"
	"fmt"
	"strconv"
	"strings"
)

var memo map[string]int

func Run() [2]int {
	puzzles := parse(inputs.Input12)
	memo = make(map[string]int)
	p1 := part1(puzzles)
	p2 := part2(parse(inputs.Input12))
	return [2]int{p1, p2}
}

func part2(puzzles []puzzle) int {
	for i, p := range puzzles {
		newPu := p
		for i := 0; i < 4; i++ {
			newPu.springs += "?" + p.springs
			newPu.pattern = append(newPu.pattern, p.pattern...)
		}
		puzzles[i] = newPu
	}

	sum := 0
	for _, p := range puzzles {
		combis := solve(p)
		sum += combis
	}
	return sum
}

func part1(puzzles []puzzle) int {
	sum := 0
	for _, p := range puzzles {
		combis := solve(p)
		sum += combis
	}
	return sum
}

func (p puzzle) key() string {
	return fmt.Sprintf("%v%v", string(p.springs), p.pattern)
}

func solve(p puzzle) int {
	key := p.key()
	if v, ok := memo[key]; ok {
		return v
	}
	res := ssolve(p)
	memo[key] = res
	return res
}

func ssolve(p puzzle) int {
	// check some end conditions
	if len(p.pattern) == 0 {
		if len(p.springs) == 0 {
			return 1
		}
		for _, c := range p.springs {
			if c == '#' {
				return 0
			}
		}
		return 1
	}
	if len(p.springs) == 0 {
		return 0
	}

	// save time if we can't fit the rest of the pattern
	minLen := 0
	for _, p := range p.pattern {
		minLen += p
	}
	if len(p.springs) < minLen {
		return 0
	}

	// check the first character
	switch p.springs[0] {
	case '.':
		p.springs = p.springs[1:]
		return solve(p)
	case '?':
		a := solve(puzzle{strings.Replace(p.springs, "?", "#", 1), p.pattern})
		b := solve(puzzle{strings.Replace(p.springs, "?", ".", 1), p.pattern})
		return a + b

	case '#':
		fits := true
		for i := 0; i < p.pattern[0]; i++ {
			if i >= len(p.springs) || p.springs[i] == '.' {
				fits = false
				break
			}
		}
		if fits {
			p.springs = p.springs[p.pattern[0]:]
			p.pattern = p.pattern[1:]
			if len(p.springs) > 0 && p.springs[0] == '?' {
				p.springs = p.springs[1:]
			} else if len(p.springs) > 0 && p.springs[0] == '#' {
				return 0
			}
			return solve(p)
		}
		return 0
	}
	panic("unreachable")
}

type puzzle struct {
	springs string
	pattern []int
}

func parse(input string) []puzzle {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	puzzles := make([]puzzle, len(lines))
	for i, l := range lines {
		parts := strings.Split(l, " ")
		pattern := parsePattern(parts[1])
		puzzles[i] = puzzle{parts[0], pattern}
	}
	return puzzles
}

func parsePattern(input string) []int {
	numsStr := strings.Split(input, ",")
	nums := make([]int, len(numsStr))
	for j, n := range numsStr {
		nums[j], _ = strconv.Atoi(n)
	}
	return nums
}
