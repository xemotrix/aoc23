package day10

import (
	"aoc23/inputs"
)

func Run() [2]int {
	pmap := parse()
	p1 := part1(&pmap)
	p2 := part2(&pmap)
	return [2]int{p1, p2}
}

func part2(pmap *PipeMap) int {
	c := 0
	var buffer rune
	in := false

	for i, r := range pmap.runes {
		if pmap.runes[i] == '\n' {
			in = false
			buffer = 0
			continue
		}
		if pmap.visited[i] {
			switch r {
			case '|':
				in = !in
			case 'L':
				buffer = 'L'
			case '7':
				if buffer == 'L' {
					buffer = 0
					in = !in
				} else {
					buffer = '7'
				}
			case 'J':
				if buffer == 'F' {
					buffer = 0
					in = !in
				} else {
					buffer = 'J'
				}
			case 'F':
				buffer = 'F'
			}
		} else if in {
			c++
		}
	}
	return c
}

func part1(pmap *PipeMap) int {
	current := pmap.start
	count := 1

outer:
	for {
		for dir := NOR; dir <= WES; dir <<= 1 {
			newPos := pmap.move(current, dir)
			if pmap.visited[newPos] || !pmap.legal(current, newPos, dir) {
				continue
			}
			pmap.visited[current] = true
			current = newPos
			count++
			continue outer
		}
		break
	}
	pmap.visited[current] = true
	return (count + 1) / 2
}

func parse() PipeMap {
	s := []rune(inputs.Input10)
	width := 0
	graph := make([]byte, len(s))
	start := 0

	for i := range s {
		if s[i] == '\n' {
			if width == 0 {
				width = i
			}
			continue
		}
		switch s[i] {
		case '|':
			graph[i] = NOR | SOU
		case '-':
			graph[i] = EAS | WES
		case 'L':
			graph[i] = NOR | EAS
		case '7':
			graph[i] = WES | SOU
		case 'J':
			graph[i] = WES | NOR
		case 'F':
			graph[i] = SOU | EAS
		case '.':
			continue
		case 'S':
			graph[i] = NOR | EAS | SOU | WES
			start = i
		default:
			panic("unknown rune")
		}
	}

	visited := make([]bool, len(graph))
	return PipeMap{
		runes:   s,
		graph:   graph,
		start:   start,
		visited: visited,
		width:   width,
	}
}

const (
	NOR byte = 0b0001 // North
	EAS byte = 0b0010 // East
	SOU byte = 0b0100 // South
	WES byte = 0b1000 // West
)

type PipeMap struct {
	runes   []rune
	graph   []byte
	visited []bool
	start   int
	width   int
}

func (m *PipeMap) move(idx int, dir byte) int {
	newPos := 0
	switch dir {
	case NOR:
		newPos = idx - m.width - 1
	case EAS:
		newPos = idx + 1
	case SOU:
		newPos = idx + m.width + 1
	case WES:
		newPos = idx - 1
	default:
		panic("unknown direction")
	}

	return newPos
}

func (m *PipeMap) legal(idx, newPos int, dir byte) bool {
	if m.graph[idx]&dir == 0 || newPos < 0 || newPos >= len(m.graph) || m.graph[newPos] == 0 {
		return false
	}

	newDirs := m.graph[newPos]
	if dir == NOR && (newDirs&SOU) != 0 {
		return true
	} else if dir == EAS && (newDirs&WES) != 0 {
		return true
	} else if dir == SOU && (newDirs&NOR) != 0 {
		return true
	} else if dir == WES && (newDirs&EAS) != 0 {
		return true
	}

	return false
}
