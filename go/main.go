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
	"fmt"
	"time"
)

func main() {
	// fmt.Println(fmtRes(8, measured(day08.Run)))
	start := time.Now()
	runAll()
	fmt.Printf("Total time: %s\n", time.Since(start))
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
	str := ""
	str += fmtRes(1, measured(day01.Run))
	str += fmtRes(2, measured(day02.Run))
	str += fmtRes(3, measured(day03.Run))
	str += fmtRes(4, measured(day04.Run))
	str += fmtRes(5, measured(day05.Run))
	str += fmtRes(6, measured(day06.Run))
	str += fmtRes(7, measured(day07.Run))
	str += fmtRes(8, measured(day08.Run))
	fmt.Println(str)
}
