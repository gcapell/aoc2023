package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	total := 0
	lines := strings.Split(string(data), "\n")
	cardCount := make([]int, len(lines)) // cardnum -> total number of that card
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		cardCount[i]++
		total += cardCount[i]
		won := score(line)
		for j := i + 1; j <= i+won; j++ {
			cardCount[j] += (cardCount[i])
		}
	}
	fmt.Println(total)
}

func score(line string) int {
	_, nums, _ := strings.Cut(line, ":")
	winners, guesses, _ := strings.Cut(nums, "|")
	return common(set(winners), set(guesses))
}

func set(nums string) map[string]bool {
	m := make(map[string]bool)
	for _, n := range strings.Fields(nums) {
		m[n] = true
	}
	return m
}

func common(winners, guesses map[string]bool) int {
	total := 0
	for n := range winners {
		if guesses[n] {
			total++
		}
	}
	return total
}
