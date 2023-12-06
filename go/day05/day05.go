package day05

func Run() [2]int {
	al := parse()
	p1 := part1(al)
	p2 := part2(al)
	return [2]int{p1, p2}
}
