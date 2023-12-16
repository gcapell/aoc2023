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

func arrangements(line string) int {
	pattern, runs, ok := strings.Cut(line, " ")
	pattern = strings.Join([]string{pattern, pattern, pattern, pattern, pattern}, "?")
	runs = strings.Join([]string{runs, runs, runs, runs, runs}, ",")
	if !ok {
		panic(line)
	}
	runsF := strings.Split(runs, ",")
	runsN := make([]int, len(runsF))
	for i, r := range runsF {
		runsN[i] = atoi(r)
	}
	reply := dp(pattern, runsN, 0)
	return reply
}

var cache = make(map[string]int)

func intOf(b bool) int {
	if b {
		return 1
	}
	return 0
}

func dp2(pattern string, runs []int, runLength int) int {
	if len(pattern) == 0 {
		if runLength != 0 {
			if len(runs) != 1 {
				return 0
			}
			return intOf(runs[0] == runLength)
		}
		return intOf(len(runs) == 0)
	}

	switch pattern[0] {
	case '.':
		if runLength != 0 {
			if runs[0] != runLength {
				return 0
			}
			return dp(pattern[1:], runs[1:], 0)
		}
		return dp(pattern[1:], runs, 0)
	case '#':
		if len(runs) == 0 || runLength == runs[0] {
			return 0
		}
		return dp(pattern[1:], runs, runLength+1)

	default:
		return dp("."+pattern[1:], runs, runLength) + dp("#"+pattern[1:], runs, runLength)
	}

}

func dp(pattern string, runs []int, runLength int) int {
	key := fmt.Sprint(pattern, runs, runLength)
	if n, ok := cache[key]; ok {
		//fmt.Println("cached", key, "->", n)
		return n
	}
	reply := dp2(pattern, runs, runLength)
	cache[key] = reply
	return reply
}
