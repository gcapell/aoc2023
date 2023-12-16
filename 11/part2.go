package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	stars := parse()
	expand(stars)
	fmt.Println(sumDistances(stars))
}

type star struct {
	r, c int
}

func (s *star) String() string {
	return fmt.Sprintf("%v", *s)
}

func parse() []*star {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	var reply []*star
	lines := strings.Fields(string(data))
	for r, line := range lines {
		for c, b := range line {
			if b == '#' {
				reply = append(reply, &star{r, c})
			}
		}
	}
	return reply
}

func expand(stars []*star) {
	expandX(stars)
	expandY(stars)
}

const expansionRate = 1_000_000

func delta(this, prev int) int {
	gap := this - (prev + 1)
	if gap == 0 {
		return 0
	}
	return gap * (expansionRate - 1)
}

func expandX(stars []*star) {
	sort.Slice(stars, func(a, b int) bool { return stars[a].r < stars[b].r })

	// empty[n] counts how many of the columns before  n are empty
	empty := make(map[int]int)
	total := 0
	prev := &star{-1, 0}
	for _, s := range stars {
		if s.r == prev.r {
			continue
		}
		d := delta(s.r, prev.r)
		total += d
		empty[s.r] = total
		prev = s
	}
	for _, s := range stars {
		s.r += empty[s.r]
	}
}

func expandY(stars []*star) {
	sort.Slice(stars, func(a, b int) bool { return stars[a].c < stars[b].c })

	// empty[n] counts how many of the rows before  n are empty
	empty := make(map[int]int)
	total := 0
	prev := &star{0, -1}
	for _, s := range stars {
		if s.c == prev.c {
			continue
		}
		total += delta(s.c, prev.c)
		empty[s.c] = total
		prev = s
	}
	for _, s := range stars {
		s.c += empty[s.c]
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func distance(a, b *star) int {
	return abs(a.r-b.r) + abs(a.c-b.c)
}

func sumDistances(stars []*star) int {
	total := 0
	for j, s := range stars {
		for _, s2 := range stars[j+1:] {
			total += distance(s, s2)
		}
	}
	return total
}
