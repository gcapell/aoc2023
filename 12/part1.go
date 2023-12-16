package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	total := 0
	for _, line := range lines() {
		total += arrangements(line)
	}
	fmt.Println("total", total)
}

func lines() []string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func match(pattern []byte, runs []int) bool {
	inRun := false
	runLength := 0
	for _, b := range pattern {
		if b == '#' {
			if !inRun {
				inRun = true
				runLength = 1
				if len(runs) == 0 {
					return false
				}
			} else {
				runLength++
				if runLength > runs[0] {
					return false
				}
			}
		} else {
			if inRun {
				if runLength != runs[0] {
					return false
				}
				inRun = false
				runLength = 0
				runs = runs[1:]
			}
		}
	}
	if inRun {
		return len(runs) == 1 && runLength == runs[0]
	}
	return len(runs) == 0

}

func inc(pattern []byte, unknownPos []int) bool {
	for _, pos := range unknownPos {
		if pattern[pos] == '.' {
			pattern[pos] = '#'
			return true
		}
		pattern[pos] = '.'
	}
	return false
}

func arrangements(line string) int {
	pattern, runs, ok := strings.Cut(line, " ")
	if !ok {
		panic(line)
	}
	runsF := strings.Split(runs, ",")
	runsN := make([]int, len(runsF))
	for i, r := range runsF {
		runsN[i] = atoi(r)
	}

	var unknownPos []int
	for i, b := range pattern {
		if b == '?' {
			unknownPos = append(unknownPos, i)
		}
	}

	patternB := []byte(pattern)
	for _, pos := range unknownPos {
		patternB[pos] = '.'
	}
	total := 0
	for {
		matched := match(patternB, runsN)
		//fmt.Println("match", string(patternB), runs, matched)
		if matched {
			total++
		}
		if !inc(patternB, unknownPos) {
			break
		}
	}
	fmt.Println("arrange", pattern, runsN, total)
	return total
}
