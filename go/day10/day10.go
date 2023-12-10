package day10

import (
	"aoc23/inputs"
	"errors"
	"fmt"
)

func Run() [2]int {
	pmap := parse()
	p1 := part1(&pmap)
	p2 := part2(&pmap)
	// pmap.debug(pmap.start)
	return [2]int{p1, p2}
}

func part2(pmap *PipeMap) int {
	markOuter(pmap, 0)
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
			pmap.enclosed[i] = in
			c++
		}
	}
	return c
}

func markOuter(pmap *PipeMap, current int) {
	pmap.enclosed[current] = true
	for dir := NOR; dir <= WES; dir <<= 1 {
		newPos := pmap.move(current, dir)
		if newPos < 0 || newPos >= len(pmap.runes) || pmap.runes[newPos] == '\n' || pmap.visited[newPos] || pmap.enclosed[newPos] {
			continue
		}
		pmap.enclosed[newPos] = true
		markOuter(pmap, newPos)
	}
}

func part1(pmap *PipeMap) int {
	current := pmap.start
	pmap.visited[current] = true

	for dir := NOR; dir <= WES; dir <<= 1 {
		moved := pmap.move(current, dir)
		err := pmap.legal(current, moved, dir)
		if err != nil {
			continue
		}
		current = moved
		break
	}

	count := 1

outer:
	for {
		for dir := NOR; dir <= WES; dir <<= 1 {
			newPos := pmap.move(current, dir)
			err := pmap.legal(current, newPos, dir)
			if err != nil || pmap.visited[newPos] {
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
	height := 0

	graph := make([]byte, len(s))
	start := 0

	for i := range s {
		if s[i] == '\n' {
			if width == 0 {
				width = i
				height = len(s)/i - 1
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
	enclosed := make([]bool, len(graph))
	return PipeMap{
		runes:    s,
		graph:    graph,
		start:    start,
		visited:  visited,
		enclosed: enclosed,
		width:    width,
		height:   height,
	}
}

const (
	NOR byte = 0b0001 // North
	EAS byte = 0b0010 // East
	SOU byte = 0b0100 // South
	WES byte = 0b1000 // West
)

type PipeMap struct {
	runes    []rune
	graph    []byte
	visited  []bool
	enclosed []bool
	start    int
	width    int
	height   int
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

func (m *PipeMap) legal(idx, newPos int, dir byte) error {
	if m.graph[idx]&dir == 0 {
		return errors.New("illegal move")
	}

	if newPos < 0 || newPos >= len(m.graph) {
		return errors.New(fmt.Sprintf("index %d out of bounds", newPos))
	}

	if m.graph[newPos] == 0 {
		return errors.New("empty space")
	}

	newDirs := m.graph[newPos]
	var connected bool
	if dir == NOR {
		connected = (newDirs & SOU) != 0
	} else if dir == EAS {
		connected = (newDirs & WES) != 0
	} else if dir == SOU {
		connected = (newDirs & NOR) != 0
	} else if dir == WES {
		connected = (newDirs & EAS) != 0
	}

	if !connected {
		return errors.New(fmt.Sprintf(
			"not connected %d [%c] -> %d [%c] with direction %04b (%04b)",
			idx, m.runes[idx], newPos, m.runes[newPos], dir, newDirs))
	}

	return nil
}

func (pmap *PipeMap) debug(idx int) {
	pre := "\033[35;40m"
	post := "\033[0m"
	enclosed := "\033[92;40m"
	for i, r := range pmap.runes {
		if pmap.enclosed[i] {
			fmt.Print(enclosed)
		} else if pmap.visited[i] {
			fmt.Print(pre)
		}
		if i == idx {
			fmt.Print("@")
		} else if r == '-' {
			fmt.Printf("%c", '━')
		} else if r == '|' {
			fmt.Printf("%c", '┃')
		} else if r == 'L' {
			fmt.Printf("%c", '┗')
		} else if r == '7' {
			fmt.Printf("%c", '┓')
		} else if r == 'J' {
			fmt.Printf("%c", '┛')
		} else if r == 'F' {
			fmt.Printf("%c", '┏')
		} else if r == 'S' {
			fmt.Print("S")
		} else if r == '.' {
			fmt.Print("X")
		} else if r == '\n' {
			fmt.Println()
		}

		if pmap.visited[i] || pmap.enclosed[i] {
			fmt.Print(post)
		}
	}
	fmt.Println()
}
