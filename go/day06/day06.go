package day06

import (
	"aoc23/inputs"
	"math"
	"strconv"
	"strings"
)

func Run() [2]int {
	races := parse()
	p1 := part1(races)
	p2 := part2(races)
	return [2]int{p1, p2}
}

func part2(races []Race) int {
	timeStr := ""
	distanceStr := ""
	for _, r := range races {
		timeStr = timeStr + strconv.Itoa(r.Time)
		distanceStr = distanceStr + strconv.Itoa(r.Distance)
	}
	time, _ := strconv.Atoi(timeStr)
	distance, _ := strconv.Atoi(distanceStr)
	return calcStuff(distance, time)

}
func part1(races []Race) int {
	res := 1
	for _, r := range races {
		res *= calcStuff(r.Distance, r.Time)
	}
	return res
}

func calcStuff(dist int, time int) int {
	factor := math.Sqrt(float64(time*time - 4*dist))
	delta := 0.00000001
	r1 := int(math.Ceil((float64(time)-factor)/2.0 + delta))
	r2 := int(math.Floor((float64(time)+factor)/2.0 - delta))
	diff := r2 - r1 + 1
	return diff

}

type Race struct {
	Time     int
	Distance int
}

func parse() []Race {
	lines := strings.Split(inputs.Input06, "\n")
	times := []int{}
	for _, nstr := range strings.Split(lines[0], " ") {
		n, err := strconv.Atoi(nstr)
		if err == nil {
			times = append(times, n)
		}
	}
	distances := []int{}
	for _, nstr := range strings.Split(lines[1], " ") {
		n, err := strconv.Atoi(nstr)
		if err == nil {
			distances = append(distances, n)
		}
	}
	races := []Race{}
	for i := range times {
		races = append(races, Race{times[i], distances[i]})
	}

	return races
}
