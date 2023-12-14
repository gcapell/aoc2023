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
		total += predict(line)
	}
	fmt.Println("total", total)
}

func predict(line []int) int {
	lines := [][]int{line}
	for {
		d := diffs(line)
		lines = append(lines, d)
		if zeroes(d) {
			break
		}
		line = d
	}
	reverse(lines)
	return extrapolate(lines)
}

func reverse(lines [][]int) {
	for j, k := 0, len(lines)-1; j < k; j, k = j+1, k-1 {
		lines[j], lines[k] = lines[k], lines[j]
	}
}

func extrapolate(lines [][]int) int {
	//fmt.Println("extrapolate", lines)
	delta := 0
	for _, line := range lines {
		delta = line[0] - delta
	}
	return delta
}

func zeroes(line []int) bool {
	for _, n := range line {
		if n != 0 {
			return false
		}
	}
	return true
}

func diffs(line []int) []int {
	reply := make([]int, len(line)-1)
	for i := range reply {
		reply[i] = line[i+1] - line[i]
	}
	return reply
}

func lines() [][]int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	var reply [][]int
	for _, line := range lines {
		chunks := strings.Fields(line)
		if len(chunks) == 0 {
			continue
		}
		ints := make([]int, len(chunks))
		for i, chunk := range chunks {
			var err error
			ints[i], err = strconv.Atoi(chunk)
			if err != nil {
				panic(chunk)
			}
		}
		reply = append(reply, ints)
	}
	return reply
}
