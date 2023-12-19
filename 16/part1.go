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
	lines := strings.Fields(string(data))

	fmt.Println(energised(lines))

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
	squares   []string
	photons   []photon
	energised map[pos]bool
}

func energised(lines []string) int {
	b := board{
		squares: lines,
		photons: []photon{
			{pos{0, 0}, right},
		},
		energised: make(map[pos]bool),
	}
	b.sim()
	return len(b.energised)
}

func (b *board) sim() {
	visited := make(map[photon]bool)
	for len(b.photons) != 0 {
		//fmt.Println(b.photons)
		var next []photon
		for _, p := range b.photons {
			b.energised[p.p] = true
			visited[p] = true
			n := p.next(b.char(p.p))
			n = slices.DeleteFunc(n, func(p photon) bool { return !b.valid(p.p) || visited[p] })
			next = append(next, n...)
		}
		b.photons = next
	}
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
