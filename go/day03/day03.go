package day03

import (
	"aoc23/inputs"
)

func Run() [2]int {
	return [2]int{
		part1([]rune(inputs.Input03)),
		part2(),
	}
}

type symbol struct {
	sym    rune
	neighs []int
}

var partIdx = make([]symbol, 0, 1000)

func part2() int {
	score := 0
	for _, sym := range partIdx {
		if sym.sym != '*' || len(sym.neighs) != 2 {
			continue
		}

		score += sym.neighs[0] * sym.neighs[1]
	}
	return score
}

func part1(runes []rune) int {
	width := 0
	for runes[width] != '\n' {
		width++
	}
	for i, r := range runes {
		if r == '.' || (r >= '0' && r <= '9') || r == '\n' {
			continue
		}
		manageSymbol(runes, i, width)
	}

	score := calcScore()
	return score
}

func manageSymbol(runes []rune, idx, width int) {
	neighs := getNeighIdx(idx, width)
	sym := symbol{
		sym:    runes[idx],
		neighs: make([]int, 0, 4),
	}
	for _, i := range neighs {
		if i < 0 || i >= len(runes) {
			continue
		}
		if runes[i] >= '0' && runes[i] <= '9' {
			sym.neighs = append(sym.neighs, revealNum(runes, i))
		}
	}
	partIdx = append(partIdx, sym)
}

func calcScore() int {
	score := 0
	for _, sym := range partIdx {
		for _, neigh := range sym.neighs {
			score += neigh
		}
	}
	return score
}

func revealNum(runes []rune, idx int) int {
	offset := 0
	for i := idx + 1; i >= 0 && runes[i] != '\n'; i++ {
		if runes[i] >= '0' && runes[i] <= '9' {
			offset++
			continue
		}
		break
	}

	mul := 1
	num := 0
	for i := idx + offset; i >= 0 && runes[i] != '\n'; i-- {
		if runes[i] >= '0' && runes[i] <= '9' {
			num += int(runes[i]-'0') * mul
			mul *= 10
			runes[i] = '.'
		} else {
			break
		}
	}
	return num
}

func getNeighIdx(idx, width int) [8]int {
	return [8]int{
		idx - (width + 1) - 1,
		idx - (width + 1),
		idx - (width + 1) + 1,
		idx - 1,
		idx + 1,
		idx + (width + 1) - 1,
		idx + (width + 1),
		idx + (width + 1) + 1,
	}
}
