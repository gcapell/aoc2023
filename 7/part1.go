package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var hands []hand
	for _, line := range strings.Split(string(input), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		hands = append(hands, parse(line))
	}
	sort.Slice(hands, func(a, b int) bool { return !hands[a].beats(hands[b]) })
	score := 0
	for i, h := range hands {
		///fmt.Println(i+1, h.score, h)
		score += (i + 1) * h.score
	}
	fmt.Println(score)
}

type hand struct {
	cards  string
	counts []int
	score  int
}

func cmp(a, b int) (bool, bool) {
	if a == b {
		return false, true
	}
	return a < b, false
}

var cardScore = map[byte]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func cmpCard(a, b byte) (bool, bool) {
	if a == b {
		return false, true
	}
	return cardScore[a] > cardScore[b], false
}

func (h hand) beats(o hand) bool {
	if smaller, same := cmp(len(h.counts), len(o.counts)); !same {
		return smaller
	}
	if smaller, same := cmp(h.counts[0], o.counts[0]); !same {
		return !smaller
	}

	for i := range h.cards {
		if better, same := cmpCard(h.cards[i], o.cards[i]); !same {
			return better
		}
	}
	panic(fmt.Sprintf("%v, %v", h, o))
}

func parse(line string) hand {
	cards, score, ok := strings.Cut(line, " ")
	if !ok {
		panic(line)
	}
	scoreVal, err := strconv.Atoi(score)
	if err != nil {
		panic(score)
	}
	counts := make(map[rune]int)
	for _, r := range cards {
		counts[r]++
	}
	var countVals []int
	for _, v := range counts {
		countVals = append(countVals, v)
	}
	sort.Slice(countVals, func(a, b int) bool { return countVals[a] > countVals[b] })

	return hand{
		cards:  cards,
		counts: countVals,
		score:  scoreVal,
	}
}
