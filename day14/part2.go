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

func (b *board) score() int {
	total := 0
	for r, row := range b.b {
		for _, c := range row {
			if c == 'O' {
				total += len(b.b) - r
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
	return &board{b: b, alt: alt}
}

func show(b *board) {
	for _, r := range b.b {
		fmt.Println(string(r))
	}
}

func slide(b *board) {
	for c := 0; c < len(b.b[0]); c++ {
		slideCol(b.b, c)
	}
}

func slideCol(b [][]byte, c int) {
	dst := 0
outer:
	for {
		var ok bool
		dst, ok = nextAvailable(b, c, dst)
		if !ok {
			return
		}
		for src := dst + 1; src < len(b); src++ {
			switch b[src][c] {
			case 'O':
				b[dst][c] = 'O'
				b[src][c] = '.'
				dst++
			case '#':
				dst = src + 1
				continue outer
			}
		}
		return
	}
}

func nextAvailable(b [][]byte, c, r int) (int, bool) {
	for ; r < len(b); r++ {
		if b[r][c] == '.' {
			return r, true
		}
	}
	return -1, false
}
