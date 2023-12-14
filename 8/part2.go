package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	seq, g := parse()
	var lengths []int
	for _, start := range g.startNodes {
		n := g.cycleLength(start, seq)
		lengths = append(lengths, n)
		fmt.Println(start, n)
	}
	fmt.Println(lengths, lcm(lengths))
}

var primes = []int{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43,
	47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101,
	103, 107, 109, 113, 127, 131, 137, 139, 149,
}

func factors(n int) map[int]int {
	counts := make(map[int]int)
	for _, p := range primes {
		for n%p == 0 {
			n = n / p
			counts[p]++
			if n == 1 {
				return counts
			}
		}
	}
	counts[n] = 1
	return counts
}

func lcm(ns []int) int {
	maxCounts := make(map[int]int)
	for _, n := range ns {
		f := factors(n)
		fmt.Println("factors", n, f)
		for n, c := range f {
			if c > maxCounts[n] {
				maxCounts[n] = c
			}
		}
	}
	fmt.Println("maxCounts", maxCounts)
	total := 1
	for n, c := range maxCounts {
		for j := 0; j < c; j++ {
			total *= n
		}
	}
	return total
}

func (g *graph) cycleLength(start string, seq *seq) int {
	steps := 0
	seq.reset()
	g.current = start
	for !g.atZZZ() {
		g.step(seq.nextIsLeft())
		steps++
	}
	return steps
}

type seq struct {
	instructions string
	pos          int
}

func (s *seq) reset() {
	s.pos = 0
}

func newSeq(s string) *seq {
	return &seq{
		instructions: strings.TrimSpace(s),
		pos:          0,
	}
}

func (s *seq) nextIsLeft() bool {
	i := s.instructions[s.pos]
	s.pos = (s.pos + 1) % len(s.instructions)
	return i == 'L'
}

type graph struct {
	nodes      map[string]*node
	current    string
	startNodes []string
}

type node struct {
	left, right string
}

var matcher = regexp.MustCompile(`(...) = \((...), (...)\)`)

func newGraph(s string) *graph {
	g := &graph{
		nodes:   make(map[string]*node),
		current: "AAA",
	}
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		ms := matcher.FindStringSubmatch(line)
		g.nodes[ms[1]] = &node{ms[2], ms[3]}
		if ms[1][2] == 'A' {
			g.startNodes = append(g.startNodes, ms[1])
		}
	}

	return g
}

func (g *graph) step(left bool) {
	n := g.nodes[g.current]
	if left {
		g.current = n.left
	} else {
		g.current = n.right
	}
}

func (g *graph) atZZZ() bool {
	return g.current[2] == 'Z'
}

func parse() (*seq, *graph) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	seqData, graphData, ok := strings.Cut(string(data), "\n\n")
	if !ok {
		log.Fatal("eek")
	}
	return newSeq(seqData), newGraph(graphData)
}
