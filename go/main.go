package main

import (
	"aoc23/day01"
	"aoc23/day02"
	"aoc23/day03"
	"aoc23/day04"
	"aoc23/day05"
	"aoc23/day06"
	"aoc23/day07"
	"aoc23/day08"
	"aoc23/day09"
	"aoc23/day10"
	"aoc23/day11"
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println(fmtRes(11, measured(day11.Run)))
	// start := time.Now()
	// runAll()
	// fmt.Printf("Total time: %s\n", time.Since(start))
}

func fmtRes(day int, result measuredRes[[2]int]) string {
	return fmt.Sprintf(
		"\nDAY %d: %v\n- part1: %d\n- part2: %d\n",
		day,
		result.t,
		result.res[0],
		result.res[1],
	)
}

type measuredRes[T any] struct {
	res T
	t   time.Duration
}

func measured[T any](f func() T) measuredRes[T] {
	start := time.Now()
	res := f()
	return measuredRes[T]{res: res, t: time.Since(start)}
}

func runAll() {
	b := strings.Builder{}
	b.WriteString(fmtRes(1, measured(day01.Run)))
	b.WriteString(fmtRes(2, measured(day02.Run)))
	b.WriteString(fmtRes(3, measured(day03.Run)))
	b.WriteString(fmtRes(4, measured(day04.Run)))
	b.WriteString(fmtRes(5, measured(day05.Run)))
	b.WriteString(fmtRes(6, measured(day06.Run)))
	b.WriteString(fmtRes(7, measured(day07.Run)))
	b.WriteString(fmtRes(8, measured(day08.Run)))
	b.WriteString(fmtRes(9, measured(day09.Run)))
	b.WriteString(fmtRes(10, measured(day10.Run)))
	fmt.Println(b.String())
}
