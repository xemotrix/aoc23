package day08

import (
	"aoc23/inputs"
	"strings"
	"unsafe"
)

const (
	AAA = 0x414141
	ZZZ = 0x5a5a5a
)

func Run() [2]int {
	dMap := parse()
	p1 := part1(dMap)
	p2 := part2(dMap)
	return [2]int{p1, p2}
}

func part1(dMap DesertMap) int {
	return calcPeriod(dMap.instructions, dMap.imap, dMap.begin, []uint16{dMap.end})
}

func part2(dMap DesertMap) int {
	res := 1
	for _, begin := range dMap.begins {
		res = lcm(res, calcPeriod(dMap.instructions, dMap.imap, begin, dMap.ends))
	}
	return res
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func calcPeriod(instructions []byte, imap [][2]uint16, start uint16, ends []uint16) int {
	var steps int = 0
	for {
		for _, end := range ends {
			if start == end {
				return steps
			}
		}
		start = imap[start][instructions[steps%len(instructions)]]
		steps++
	}
}

type DesertMap struct {
	imap         [][2]uint16
	instructions []byte
	begin        uint16
	end          uint16
	begins       []uint16
	ends         []uint16
}

func parse() DesertMap {
	lines := strings.Split(inputs.Input08, "\n")
	steps := []rune(lines[0])
	instr := make([]byte, len(steps))

	for i, r := range steps {
		if r == 'L' {
			instr[i] = 0
		} else if r == 'R' {
			instr[i] = 1
		}
	}

	directions := make([][2]uint32, len(lines)-3)
	trans := make(map[uint32]uint16)

	myMap := DesertMap{
		imap:         make([][2]uint16, len(lines)-3),
		instructions: instr,
		begins:       make([]uint16, 0),
		ends:         make([]uint16, 0),
	}

	for i, line := range lines[2 : len(lines)-1] {
		from := *(*uint32)(unsafe.Pointer(&([]byte(line[:3]))[0]))
		left := *(*uint32)(unsafe.Pointer(&([]byte(line[7:10]))[0]))
		right := *(*uint32)(unsafe.Pointer(&([]byte(line[12:15]))[0]))

		if from == AAA {
			myMap.begin = uint16(i)
		} else if from == ZZZ {
			myMap.end = uint16(i)
		}
		if from>>16 == 'A' {
			myMap.begins = append(myMap.begins, uint16(i))
		} else if from>>16 == 'Z' {
			myMap.ends = append(myMap.ends, uint16(i))
		}

		trans[from] = uint16(i)
		directions[i] = [2]uint32{left, right}
	}

	for i := uint16(0); i < uint16(len(directions)); i++ {
		myMap.imap[i] = [2]uint16{trans[directions[i][0]], trans[directions[i][1]]}
	}

	return myMap

}
