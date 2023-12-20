package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	prog, inputs := cut(string(data), "\n\n")
	p := parseProg(prog)

	total := 0
	for _, i := range strings.Fields(inputs) {
		total += run(p, parseInput(i))
	}
	fmt.Println("total", total)
}

func run(src map[string][]instruction, p piece) int {
	//fmt.Println()
	//fmt.Println(p)
	s := "in"
outer:
	for {
		//fmt.Println("executing", s)
		for _, i := range src[s] {
			if next, ok := i(p); ok {
				//fmt.Println(next)
				switch next {
				case "R":
					return 0
				case "A":
					return p.x + p.m + p.a + p.s
				default:
					s = next
					continue outer
				}
			}
		}
	}
}

func parseInput(s string) piece {
	s = strings.Trim(s, "{}")
	chunks := strings.Split(s, ",")
	return piece{
		x: extractVal(chunks[0], "x"),
		m: extractVal(chunks[1], "m"),
		a: extractVal(chunks[2], "a"),
		s: extractVal(chunks[3], "s"),
	}
}

func extractVal(s string, key string) int {
	k, digits := cut(s, "=")
	if k != key {
		panic(fmt.Sprint(s, key))
	}
	return atoi(digits)
}

// ([a-z]+)([<>]):([a-z]+|[AR])
var conditionalRegexp = regexp.MustCompile(`([amsx])([<>])([0-9]+):([a-z]+|[AR]).*`)

type piece struct {
	x, m, a, s int
}

type instruction func(piece) (string, bool)

type prog map[string][]instruction

func parseProg(s string) prog {
	p := make(map[string][]instruction)
	for _, line := range strings.Fields(s) {
		name, content := cut(line, "{")
		content = strings.TrimSuffix(content, "}")
		var wf []instruction
		for _, section := range strings.Split(content, ",") {
			m := conditionalRegexp.FindStringSubmatch(section)
			if len(m) > 0 {
				wf = append(wf, conditional(m[1:]))
			} else {
				wf = append(wf, absolute(section))
			}
		}
		p[name] = wf
	}
	return p
}

func absolute(s string) instruction {
	return func(piece) (string, bool) {
		return s, true
	}
}

func conditional(m []string) instruction {

	key, cond, val, dst := m[0], m[1], m[2], m[3]
	v := atoi(val)
	return func(p piece) (string, bool) {
		//fmt.Println(m)
		var n int
		switch key {
		case "x":
			n = p.x
		case "m":
			n = p.m
		case "a":
			n = p.a
		case "s":
			n = p.s
		default:
			panic(key)
		}
		//fmt.Printf("%v(%s) %s %v?\n", n, key, cond, v)
		switch cond {
		case "<":
			if n < v {
				return dst, true
			}
		case ">":
			if n > v {
				return dst, true
			}
		default:
			panic(cond)
		}
		return "", false
	}
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func cut(s, tok string) (string, string) {
	a, b, ok := strings.Cut(s, tok)
	if !ok {
		panic(s)
	}
	return a, b
}
