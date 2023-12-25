package day11

import (
	"aoc23/inputs"

	"golang.org/x/exp/slices"
)

func Run() [2]int {
	stars := parse(inputs.Input11)

	p1 := part1(stars)
	p2 := part2(stars)
	return [2]int{p1, p2}
}

func part2(stars [][2]int) int {
	return calcDists(stars, 1000000)
}

func part1(stars [][2]int) int {
	return calcDists(stars, 2)
}

func calcDists(stars [][2]int, factor int) int {
	copied := make([][2]int, len(stars))
	copy(copied, stars)
	stars = expandUniverse(copied, factor)
	sum := 0
	for i := range stars {
		for j := range stars {
			if i < j {
				continue
			}
			sum += abs(stars[i][0]-stars[j][0]) + abs(stars[i][1]-stars[j][1])
		}
	}
	return sum
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func parse(input string) [][2]int {
	var width int
	for i, c := range input {
		if c == '\n' {
			width = i
			break
		}
	}
	width++

	stars := make([][2]int, 0)
	for i, c := range input {
		if c == '#' {
			stars = append(stars, [2]int{i / width, i % width})
		}
	}
	return stars
}

func expandUniverse(stars [][2]int, factor int) [][2]int {
	stars = expandDimension(stars, factor, 0)
	slices.SortFunc(stars, func(i, j [2]int) int {
		return i[1] - j[1]
	})
	stars = expandDimension(stars, factor, 1)
	return stars
}

func expandDimension(stars [][2]int, factor, dim int) [][2]int {
	last := -1
	toExpand := make([]int, 0)
	for _, star := range stars {
		if star[dim] == last {
			continue
		}
		if !(star[dim] == last+1) {
			for i := last + 1; i < star[dim]; i++ {
				toExpand = append(toExpand, i)
			}
		}
		last = star[dim]
	}

	expandIdx := 0
	offset := 0
	for i, star := range stars {
		for expandIdx < len(toExpand) && star[dim] > toExpand[expandIdx] {
			offset += factor - 1
			expandIdx++
		}
		stars[i][dim] += offset
	}
	return stars
}
