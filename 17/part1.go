package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Fields(string(data))
	losses := make([][]int, len(lines))
	for i, line := range lines {
		losses[i] = ints(line)
	}

	//fmt.Println(losses)

	dst := pos{
		len(losses) - 1,
		len(losses[0]) - 1,
	}
	fmt.Println(shortestPath(losses, dst))

}

func ints(line string) []int {
	reply := make([]int, len(line))
	for i, c := range line {
		reply[i] = int(c - '0')
	}
	return reply
}

type direction int

const (
	north direction = iota
	east
	west
	south
)

type pos struct{ r, c int }

// arrived at position p with move in direction d, after lastMoveCount steps in that direction
type fullPos struct {
	p             pos
	d             direction
	lastMoveCount int
}

func shortestPath(board [][]int, dst pos) int {
	border := make(map[fullPos]int) // shortest path to get to each fullPos
	border[fullPos{pos{0, 1}, east, 1}] = board[0][1]
	border[fullPos{pos{1, 0}, south, 1}] = board[1][0]
	visited := make(map[fullPos]bool)

	for {
		fp, cost := popShortest(border)
		visited[fp] = true
		if fp.p == dst {
			return cost
		}
		for _, p := range []fullPos{fp.left(), fp.right(), fp.forward()} {
			if !p.valid(board) || visited[p] {
				continue
			}
			newCost := cost + board[p.p.r][p.p.c]
			if n, ok := border[p]; !ok || newCost < n {
				border[p] = newCost
			}
		}
	}
}

func (fp fullPos) left() fullPos {
	d := fp.d.left()
	return fullPos{move(fp.p, d), d, 1}
}
func (fp fullPos) right() fullPos {
	d := fp.d.right()
	return fullPos{move(fp.p, d), d, 1}
}
func (fp fullPos) forward() fullPos {
	return fullPos{move(fp.p, fp.d), fp.d, fp.lastMoveCount + 1}
}

func move(p pos, d direction) pos {
	switch d {
	case north:
		return pos{p.r - 1, p.c}
	case south:
		return pos{p.r + 1, p.c}
	case east:
		return pos{p.r, p.c + 1}
	case west:
		return pos{p.r, p.c - 1}
	default:
		panic(fmt.Sprint(p, d))
	}
}

func (d direction) left() direction {
	switch d {
	case north:
		return west
	case south:
		return east
	case west:
		return south
	case east:
		return north
	default:
		panic(d)
	}
}

func (d direction) right() direction {
	switch d {
	case north:
		return east
	case south:
		return west
	case west:
		return north
	case east:
		return south
	default:
		panic(d)
	}
}

func (fp fullPos) valid(board [][]int) bool {
	return fp.lastMoveCount <= 3 && fp.p.r >= 0 && fp.p.r < len(board) && fp.p.c >= 0 && fp.p.c < len(board[0])
}

func popShortest(border map[fullPos]int) (fullPos, int) {
	if len(border) == 0 {
		panic("empty")
	}
	first := true
	var min int
	var minK fullPos
	for k, v := range border {
		if first || v < min {
			first = false
			min = v
			minK = k
		}
	}
	delete(border, minK)
	return minK, min
}
