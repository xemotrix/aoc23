package day04

import (
	"aoc23/inputs"
	"strconv"
	"strings"
)

func Run() [2]int {
	games := parse()
	p1, scores := part1(games)
	return [2]int{p1, part2(scores)}
}

func part2(wins []int) int {
	numCards := make([]int, len(wins))
	for i, s := range wins {
		numCards[i]++
		for j := i + 1; j <= i+s; j++ {
			numCards[j] += numCards[i]
		}
	}
	score := 0
	for _, c := range numCards {
		score += c
	}
	return score
}

func part1(games []Game) (int, []int) {
	totScore := 0
	wins := make([]int, len(games))
	for i, g := range games {
		table := make([]bool, 100)
		for _, w := range g.winning {
			table[w] = true
		}
		nWins := 0
		for _, p := range g.played {
			if table[p] {
				nWins++
			}
		}
		if nWins == 0 {
			continue
		}
		wins[i] = nWins
		totScore += 1 << nWins
	}
	return totScore, wins
}

type Game struct {
	winning []int
	played  []int
}

func parse() []Game {
	games := []Game{}
	lines := strings.Split(inputs.Input04, "\n")
	for _, line := range lines[:len(lines)-1] {
		games = append(games, parseLine(line))
	}
	return games
}

func parseLine(line string) Game {
	line = strings.Split(line, ":")[1]
	parts := strings.Split(line, "|")
	winning := []int{}
	for _, nstr := range strings.Split(strings.TrimSpace(parts[0]), " ") {
		n, err := strconv.Atoi(nstr)
		if err != nil {
			continue
		}
		winning = append(winning, n)
	}
	played := []int{}
	for _, nstr := range strings.Split(strings.TrimSpace(parts[1]), " ") {
		n, err := strconv.Atoi(nstr)
		if err != nil {
			continue
		}
		played = append(played, n)
	}
	return Game{
		winning: winning,
		played:  played,
	}
}
