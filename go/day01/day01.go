package day01

import (
	"aoc23/inputs"
	"strings"
)

func parse(input string) []string {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	return lines
}
func Run() [2]int {
	lines := parse(inputs.Input01)
	p1 := part1(lines)
	p2 := part2(lines)

	return [2]int{p1, p2}
}

var numbers = [10]string{
	"zero", "one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine",
}

func part2(lines []string) int {
	var sum int
	for _, line := range lines {
		var num [2]rune
	first:
		for i, r := range line {
			if r >= '0' && r <= '9' {
				num[0] = r - '0'
				break
			}
			for j, n := range numbers {
				if strings.HasSuffix(line[:i+1], n) {
					num[0] = rune(j)
					break first
				}
			}
		}
	last:
		for i := len(line) - 1; i >= 0; i-- {
			r := rune(line[i])
			if r >= '0' && r <= '9' {
				num[1] = r - '0'
				break
			}
			for j, n := range numbers {
				if strings.HasPrefix(line[i:], n) {
					num[1] = rune(j)
					break last
				}
			}
		}
		sum += int(num[0])*10 + int(num[1])
	}
	return sum
}

func part1(lines []string) int {
	var sum int
	for _, line := range lines {
		var num [2]rune
		for _, r := range line {
			if r >= '0' && r <= '9' {
				num[0] = r - '0'
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			r := rune(line[i])
			if r >= '0' && r <= '9' {
				num[1] = r - '0'
				break
			}
		}
		sum += int(num[0])*10 + int(num[1])
	}
	return sum
}
