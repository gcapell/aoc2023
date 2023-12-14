package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	seq, graph := parse()
	steps := 0
	for !graph.atZZZ() {
		graph.step(seq.nextIsLeft())
		steps++
	}
	fmt.Println(steps, "steps")
}

type seq struct {
	instructions string
	pos          int
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
	nodes   map[string]*node
	current string
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
	return g.current == "ZZZ"
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
