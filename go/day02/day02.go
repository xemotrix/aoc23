package day02

import (
	"aoc23/inputs"
	"fmt"
	"strconv"
	"strings"
)

func Run() [2]int {
	games := parse(inputs.Input02)
	p1 := part1(games)
	p2 := part2(games)
	return [2]int{p1, p2}
}

const (
	thresRed   = 12
	thresGreen = 13
	thresBlue  = 14
)

type game struct {
	num    int
	result [][3]int
}

func parse(input string) []game {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	games := make([]game, len(lines))

	for i, line := range lines {
		games[i] = parseLine(line)
	}

	return games
}

func parseNum(s string) (int, int) {
	acc := ""
	for _, c := range s {
		if c >= '0' || c <= '9' {
			acc += string(c)
		}
	}
	n, _ := strconv.Atoi(acc)
	return n, len(acc)
}

func parseLine(line string) game {
	var g game
	fmt.Sscanf(line, "Game %d:", &g.num)
	line = strings.Split(line, ":")[1]
	parts := strings.Split(line, ";")
	g.result = make([][3]int, len(parts))
	for i, p := range parts {
		for _, c := range strings.Split(p, ",") {
			var num int
			fmt.Sscanf(c, "%d", &num)
			if strings.HasSuffix(c, "red") {
				g.result[i][0] = num
			} else if strings.HasSuffix(c, "green") {
				g.result[i][1] = num
			} else if strings.HasSuffix(c, "blue") {
				g.result[i][2] = num
			} else {
				panic(fmt.Sprintf("unknown color: \"%s\"", c))
			}
		}
	}
	return g
}

func part2(games []game) int {
	sum := 0
	for _, g := range games {
		maxs := [3]int{0, 0, 0}
		for _, r := range g.result {
			if r[0] > maxs[0] {
				maxs[0] = r[0]
			}
			if r[1] > maxs[1] {
				maxs[1] = r[1]
			}
			if r[2] > maxs[2] {
				maxs[2] = r[2]
			}
		}
		sum += maxs[0] * maxs[1] * maxs[2]
	}
	return sum
}

func part1(games []game) int {
	sum := 0
	for _, g := range games {
		possible := true
		for _, r := range g.result {
			if r[0] > thresRed || r[1] > thresGreen || r[2] > thresBlue {
				possible = false
			}
		}
		if possible {
			sum += g.num
		}
	}
	return sum
}
