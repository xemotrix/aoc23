package day09

import (
	"aoc23/inputs"
	"strconv"
	"strings"
)

func Run() [2]int {
	mea := parse()
	p1 := part1(mea)
	p2 := part2(mea)
	return [2]int{p1, p2}
}

func part1(mea [][]int) int {
	score := 0
	for i := 0; i < len(mea); i++ {
		score += predict(mea[i], func(a []int, b int) int { return a[len(a)-1] + b })
	}
	return score
}

func part2(mea [][]int) int {
	score := 0
	for i := 0; i < len(mea); i++ {
		score += predict(mea[i], func(a []int, b int) int { return a[0] - b })
	}
	return score
}

func predict(nums []int, f func([]int, int) int) int {
	diffs := make([]int, len(nums)-1)
	flag := 0
	for i := range nums[1:] {
		diff := nums[i+1] - nums[i]
		diffs[i] = diff
		flag |= diff
	}
	if flag == 0 {
		return nums[0]
	}
	return f(nums, predict(diffs, f))
}

func parse() [][]int {

	lines := strings.Split(inputs.Input09, "\n")
	lines = lines[:len(lines)-1]

	measurements := make([][]int, len(lines))

	for i := range lines {
		numStr := strings.Split(lines[i], " ")
		nums := make([]int, len(numStr))
		for j, ns := range numStr {
			nums[j], _ = strconv.Atoi(ns)
		}
		measurements[i] = nums
	}

	return measurements
}
