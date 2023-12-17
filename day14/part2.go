package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	b := readBoard()

	b.rotate()
	b.rotate()
	b.rotate()

	for j := 0; j < 1000; j++ {
		b.cycle()
		fmt.Println(j+1, b.score())
	}
}

func (b *board) cycle() {
	// N facing west
	b.slide() // slide north

	b.rotate()
	// N facing north
	b.slide() // slide west

	b.rotate()
	// N facing east
	b.slide() // slide south

	b.rotate()
	// N facing south (east facing west)
	b.slide() // slide east

	b.rotate()
	// N facing west
}

type board struct {
	b   [][]byte
	alt [][]byte
}

// rotate quarter turn clockwise
func (b *board) rotate() {
	for r, row := range b.b {
		for c, elem := range row {
			b.alt[c][len(b.b)-1-r] = elem
		}
	}
	b.b, b.alt = b.alt, b.b
}

func (b *board) score() int {
	total := 0
	for _, row := range b.b {
		for c, elem := range row {
			if elem == 'O' {
				total += len(row) - c
			}
		}
	}
	return total
}

func readBoard() *board {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	b := bytes.Fields(data)
	alt := make([][]byte, len(b[0]))
	for j := range alt {
		alt[j] = make([]byte, len(b))
	}
	reply := &board{b: b, alt: alt}
	return reply
}

func (b *board) show() {
	for _, r := range b.b {
		fmt.Println(string(r))
	}
	fmt.Println()
}

func (b *board) slide() {
	for _, row := range b.b {
		slideRow(row)
	}
}

func slideRow(row []byte) {
	dst := 0
outer:
	for {
		n := bytes.IndexByte(row[dst:], '.')
		if n == -1 {
			return
		}
		dst += n
		for src := dst + 1; src < len(row); src++ {
			switch row[src] {
			case 'O':
				row[dst] = 'O'
				row[src] = '.'
				dst++
			case '#':
				dst = src + 1
				continue outer
			}
		}
		return
	}
}
