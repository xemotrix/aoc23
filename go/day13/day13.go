package day13

import (
	"aoc23/inputs"
	"strings"
)

func Run() [2]int {
	input := parse()
	p1, res := part1(input)
	p2 := part2(res)
	return [2]int{p1, p2}
}

func part1(input [][][]rune) (int, []p1Res) {
	sum := 0
	results := make([]p1Res, len(input))
	for i, part := range input {
		n, res := calc(part)
		sum += n
		results[i] = res
	}
	return sum, results
}

func part2(input []p1Res) int {
	result := 0
	for _, res := range input {
		for i := range res.h {
			if res.h[i] == res.nh-1 {
				result += i
				break
			}
		}
		for i := range res.v {
			if res.v[i] == res.nv-1 {
				result += i * 100
				break
			}
		}
	}
	return result
}

type p1Res struct {
	h  []int
	v  []int
	nh int
	nv int
}

func calc(part [][]rune) (int, p1Res) {
	hSymSet := make([]int, len(part[0]))
	vSymSet := make([]int, len(part))
	trans := transpose(part)
	for i := 0; i < max(len(part), len(trans)); i++ {
		if len(hSymSet) > 0 && i < len(part) {
			addFreqMaps(hSymSet, findSymmetry(part[i]))
		}
		if len(vSymSet) > 0 && i < len(trans) {
			addFreqMaps(vSymSet, findSymmetry(trans[i]))
		}
	}
	maxH, qH := findMaxIdx(hSymSet)
	res := p1Res{hSymSet, vSymSet, len(part), len(part[0])}
	if qH == len(part) {
		return maxH, res
	}
	maxV, qV := findMaxIdx(vSymSet)
	if qV == len(part[0]) {
		return maxV * 100, res
	}
	panic("not found")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxIdx(freq []int) (int, int) {
	maxN := 0
	maxIdx := -1
	for i, n := range freq {
		if n > maxN {
			maxN = n
			maxIdx = i
		}
	}
	return maxIdx, maxN
}

func addFreqMaps(a, b []int) {
	for i := range a {
		a[i] += b[i]
	}
}

func transpose(part [][]rune) [][]rune {
	res := make([][]rune, len(part[0]))
	for i := range res {
		res[i] = make([]rune, len(part))
	}
	for i := 0; i < len(part); i++ {
		for j := 0; j < len(part[0]); j++ {
			res[j][i] = part[i][j]
		}
	}
	return res
}

func findSymmetry(line []rune) []int {
	from := 0
	to := 1
	axis := make([]int, len(line))
	for {
		if from >= to {
			break
		}
		sym := isSymetric(line[from : to+1])
		if sym {
			axis[(from+to)/2+1]++
		}
		if to == len(line)-1 {
			from++
		} else {
			to++
		}
	}
	return axis
}

func isSymetric(line []rune) bool {
	if len(line)%2 != 0 {
		return false
	}
	for i, j := 0, len(line)-1; i < j; i, j = i+1, j-1 {
		if line[i] != line[j] {
			return false
		}
	}
	return true
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func parse() [][][]rune {
	parts := strings.Split(inputs.Input13, "\n\n")
	res := make([][][]rune, len(parts))
	for i, part := range parts {
		lines := strings.Split(part, "\n")
		if lines[len(lines)-1] == "" {
			lines = lines[:len(lines)-1]
		}
		linesRune := make([][]rune, len(lines))
		for j, line := range lines {
			linesRune[j] = []rune(line)
		}
		res[i] = linesRune
	}
	return res
}
