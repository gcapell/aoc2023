package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type board struct {
	ports   []port
	symbols map[pos]string
	gears   map[pos]bool
}

func newBoard() *board {
	return &board{
		symbols: make(map[pos]string),
		gears:   make(map[pos]bool),
	}
}

type port struct {
	p pos
	s string
}

type pos struct {
	row, col int
}

func main() {
	board := loadBoard()
	//fmt.Println(board.sumPartNums())
	fmt.Println(board.sumGearRatios())
}

func loadBoard() *board {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")

	b := newBoard()

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		b.addLine(i, line)
	}
	return b
}

var pattern = regexp.MustCompile(`[0-9]+|[^.]`)

const allMatches = -1

func (b *board) addLine(row int, line string) {
	for _, m := range pattern.FindAllStringIndex(line, allMatches) {
		s := line[m[0]:m[1]]
		p := pos{row, m[0]}
		if unicode.IsDigit(rune(s[0])) {
			b.ports = append(b.ports, port{p, s})
		} else {
			b.symbols[p] = s
			if s == "*" {
				b.gears[p] = true
			}
		}
	}
}

func (b *board) sumPartNums() int {
	total := 0
	for _, port := range b.ports {
		if _, ok := port.isAdjacent(b.symbols); ok {
			total += atoi(port.s)
		}
	}
	return total
}

func (p port) isAdjacent(symbols map[pos]string) (string, bool) {
	for _, p2 := range p.neighbours() {
		if s, ok := symbols[p2]; ok {
			return s, true
		}
	}
	return "", false
}

func (p port) adjacentGears(gears map[pos]bool) []pos {
	var reply []pos
	for _, p2 := range p.neighbours() {
		if gears[p2] {
			reply = append(reply, p2)
		}
	}
	return reply
}

func (p port) neighbours() []pos {
	reply := make([]pos, 0, len(p.s)*2+3)
	for row := p.p.row - 1; row <= p.p.row+1; row++ {
		for col := p.p.col - 1; col <= p.p.col+len(p.s); col++ {
			reply = append(reply, pos{row, col})
		}
	}
	return reply
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func (b *board) sumGearRatios() int {
	gearAdjacencies := make(map[pos][]string)
	for _, port := range b.ports {
		for _, g := range port.adjacentGears(b.gears) {
			gearAdjacencies[g] = append(gearAdjacencies[g], port.s)
		}
	}
	total := 0
	for _, ports := range gearAdjacencies {
		if len(ports) != 2 {
			continue
		}
		total += atoi(ports[0]) * atoi(ports[1])
	}
	return total
}
