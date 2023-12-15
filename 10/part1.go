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
		prev, ok := findPath(n, s, board)
		if ok {
			fmt.Printf("circuit from %v leaves %v arrives %v\n", s, n, prev)
			break
		} else {
			fmt.Printf("circuit from %v leaves %v dies %v\n", s, n, prev)
		}
	}
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

func findPath(this, dst pos, board []string) (int, bool) {
	steps := 1
	prev := dst
	for {
		//fmt.Printf("findPath this:%v, prev:%v, dst:%v\n", this, prev, dst)
		if this == dst {
			return steps, true
		}
		next, ok := neighbour(this, prev, board)

		if !ok {
			return steps, false
		}
		this, prev = next, this
		steps++
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
