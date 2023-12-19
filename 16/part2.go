package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	b := board{
		squares: strings.Fields(string(data)),
	}
	max := 0
	for i, p := range startPhotons(len(b.squares), len(b.squares[0])) {
		e := b.sim(p)
		//fmt.Println(p, "->", e)
		if i == 0 || e > max {
			max = e
		}
	}
	fmt.Println(max)
}

func startPhotons(rows, cols int) []photon {
	reply := make([]photon, 0, 2*rows+2*cols+4)

	for c := 0; c < cols; c++ {
		reply = append(reply, photon{pos{0, c}, down})
		reply = append(reply, photon{pos{rows - 1, c}, up})
	}
	for r := 0; r < rows; r++ {
		reply = append(reply, photon{pos{r, 0}, right})
		reply = append(reply, photon{pos{r, cols - 1}, left})
	}
	return reply
}

type direction int

const (
	right direction = iota
	left
	up
	down
)

type pos struct{ r, c int }

type photon struct {
	p pos
	d direction
}

func (p photon) String() string {
	return fmt.Sprintf("(%v %v)", p.p, p.d)
}

func (d direction) String() string {
	switch d {
	case left:
		return "L"
	case right:
		return "R"
	case up:
		return "U"
	case down:
		return "D"
	default:
		panic(d)

	}
}

type board struct {
	squares []string
}

func (b *board) sim(p photon) int {
	visited := make(map[photon]bool)
	energised := make(map[pos]bool)
	photons := []photon{p}
	for len(photons) != 0 {
		//fmt.Println(b.photons)
		var next []photon
		for _, p := range photons {
			energised[p.p] = true
			visited[p] = true
			n := p.next(b.char(p.p))
			n = slices.DeleteFunc(n, func(p photon) bool { return !b.valid(p.p) || visited[p] })
			next = append(next, n...)
		}
		photons = next
	}
	return len(energised)
}

func (b *board) char(p pos) byte {
	return b.squares[p.r][p.c]
}

func (b *board) valid(p pos) bool {
	return p.r >= 0 && p.r < len(b.squares) && p.c >= 0 && p.c < len(b.squares[0])
}

func (p photon) next(c byte) []photon {
	passthru := []photon{{p.p.next(p.d), p.d}}
	switch c {
	case '.':
		return passthru
	case '|':
		if p.d == left || p.d == right {
			return []photon{
				{p.p.next(up), up},
				{p.p.next(down), down},
			}
		} else {
			return passthru
		}
	case '-':
		if p.d == up || p.d == down {
			return []photon{
				{p.p.next(left), left},
				{p.p.next(right), right},
			}
		} else {
			return passthru
		}
	case '/':
		d := reflectSlash(p.d)
		return []photon{
			{p.p.next(d), d},
		}
	case '\\':
		d := reflectSlash(p.d).reverse()
		return []photon{
			{p.p.next(d), d},
		}
	default:
		panic(c)
	}
}

func (p pos) next(d direction) pos {
	switch d {
	case left:
		return pos{p.r, p.c - 1}
	case right:
		return pos{p.r, p.c + 1}
	case up:
		return pos{p.r - 1, p.c}
	case down:
		return pos{p.r + 1, p.c}
	default:
		panic(d)
	}
}

func (d direction) reverse() direction {
	switch d {
	case left:
		return right
	case right:
		return left
	case up:
		return down
	case down:
		return up
	default:
		panic(d)
	}
}

func reflectSlash(d direction) direction {
	switch d {
	case left:
		return down
	case right:
		return up
	case up:
		return right
	case down:
		return left
	default:
		panic(d)
	}
}
