package main

import (
	"fmt"
	"os"
	"strings"
)

type pos struct {
	row, col int
}

var pipes = map[byte][]direction{
	'7': {down, left},
	'F': {down, right},
	'J': {up, left},
	'L': {up, right},
	'|': {down, up},
	'-': {left, right},
}

type direction func(p pos, board []string) (pos, bool)

func down(p pos, board []string) (pos, bool) {
	if p.row+1 < len(board) {
		return pos{p.row + 1, p.col}, true
	}
	return pos{}, false
}
func up(p pos, board []string) (pos, bool) {
	if p.row > 0 {
		return pos{p.row - 1, p.col}, true
	}
	return pos{}, false
}

func left(p pos, board []string) (pos, bool) {
	if p.col > 0 {
		return pos{p.row, p.col - 1}, true
	}
	return pos{}, false
}
func right(p pos, board []string) (pos, bool) {
	if p.col+1 < len(board[0]) {
		return pos{p.row, p.col + 1}, true
	}
	return pos{}, false
}

func main() {
	boardData, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	board := strings.Fields(string(boardData))
	s := findS(board)
	for _, n := range sNeighbours(s, board) {
		path, ok := findPath(n, s, board)
		if ok {
			fmt.Printf("circuit from %v leaves %v arrives %v\n", s, n, len(path))
			fmt.Println("contained", contained(path, board))
			break
		} else {
			fmt.Printf("circuit from %v leaves %v dies %v\n", s, n, len(path))
		}
	}

}

type hState int

const (
	hNone hState = iota
	hUp
	hDown
)

func contained(path map[pos]bool, board []string) int {
	total := 0
	for row := 0; row < len(board); row++ {
		in := false
		h := hNone
		for col := 0; col < len(board[0]); col++ {

			if path[pos{row, col}] {
				switch board[row][col] {
				case '-':
				case '|':
					in = !in
				case '7':
					if h == hDown {
						in = !in
					}
					h = hNone
				case 'F':
					h = hUp
				case 'J':
					if h == hUp {
						in = !in
					}
					h = hNone
				case 'L':
					h = hDown
				}
			} else if in {
				total++
			}
		}
	}
	return total
}

func sNeighbours(p pos, board []string) []pos {
	var reply []pos
	if p.row+1 < len(board) {
		reply = append(reply, pos{p.row + 1, p.col})
	}

	if p.row > 0 {
		reply = append(reply, pos{p.row - 1, p.col})
	}

	if p.col > 0 {
		reply = append(reply, pos{p.row, p.col - 1})
	}

	if p.col+1 < len(board[0]) {
		reply = append(reply, pos{p.row, p.col + 1})
	}
	return reply
}

// neighbour returns places we can get to from 's' (which is not 'prev')
func neighbour(s, prev pos, board []string) (pos, bool) {
	for _, d := range pipes[board[s.row][s.col]] {
		next, ok := d(s, board)
		if ok && next != prev {
			return next, true
		}
	}
	//fmt.Printf("neighbors(%v-%s)->%v\n", s, string(board[s.row][s.col]), reply)
	return pos{}, false
}

func replaceS(board []string, dst, first, this pos) {

	var c byte
	if first.row == this.row {
		c = '-'
	} else if first.col == this.col {
		c = '|'
	} else {
		l, r := first, this
		if l.row > r.row {
			l, r = r, l
		}
		if l.col < dst.col {
			if r.row < dst.row {
				c = 'J'
			} else {
				c = '7'
			}
		} else {
			if l.row < dst.row {
				c = 'L'
			} else {
				c = 'F'
			}
		}
	}
	b := []byte(board[dst.row])
	b[dst.col] = c
	board[dst.row] = string(b)
}

func findPath(first, dst pos, board []string) (map[pos]bool, bool) {
	path := make(map[pos]bool)
	prev := dst
	this := first
	for {
		path[this] = true
		//fmt.Printf("findPath this:%v, prev:%v, dst:%v\n", this, prev, dst)
		if this == dst {
			replaceS(board, dst, first, prev)
			return path, true
		}
		next, ok := neighbour(this, prev, board)

		if !ok {
			return path, false
		}
		this, prev = next, this
	}
}

func findS(lines []string) pos {
	for r, row := range lines {
		for c, ch := range row {
			if ch == 'S' {
				return pos{r, c}
			}
		}
	}
	panic(lines)
}
