package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	b := readBoard()
	show(b)
	slide(b)
	fmt.Println()
	show(b)
	fmt.Println(b.score())
}

type board struct {
	b   [][]byte
	alt [][]byte
}

func (b *board) rotate() {
	for r, row := range b.b {
		for c, elem := range row {
			b.alt[c][r] = elem
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
	reply.rotate()
	return reply
}

func show(b *board) {
	for _, r := range b.b {
		fmt.Println(string(r))
	}
}

func slide(b *board) {
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
