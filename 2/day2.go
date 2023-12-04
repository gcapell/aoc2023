package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type reveal struct {
	red, green, blue int
}

type game struct {
	id      int
	reveals []reveal
}

func main() {
	games := loadGames()

	fmt.Println("sum: ", sumPossible(games, reveal{
		red:   12,
		green: 13,
		blue:  14,
	}))
	fmt.Println("sumPower: ", sumPower(games))
}

func sumPower(games []game) int64 {
	var total int64
	for _, g := range games {
		total += g.minSet().power()
	}
	return total
}

func (g *game) minSet() reveal {
	var reply reveal
	for _, r := range g.reveals {
		reply.incMax(r)
	}
	return reply
}

func (r *reveal) incMax(maxR reveal) {
	r.green = max(r.green, maxR.green)
	r.red = max(r.red, maxR.red)
	r.blue = max(r.blue, maxR.blue)
}

func (r reveal) power() int64 {
	return int64(r.red) * int64(r.green) * int64(r.blue)
}

func sumPossible(games []game, m reveal) int {
	total := 0
	for _, g := range games {
		if g.possible(m) {
			total += g.id
		}
	}
	return total
}

func (g *game) possible(maxR reveal) bool {
	for _, r := range g.reveals {
		if !r.possible(maxR) {
			return false
		}
	}
	return true
}

func (r *reveal) possible(maxR reveal) bool {
	return r.red <= maxR.red && r.blue <= maxR.blue && r.green <= maxR.green
}

func loadGames() []game {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	fmt.Println(len(lines), lines[0])
	games := make([]game, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		games = append(games, parseGame(line))
	}
	fmt.Println(len(games), games[0])
	return games
}

func parseGame(line string) game {
	title, details, ok := strings.Cut(line, ":")
	if !ok {
		log.Fatalf("no : in %s", line)
	}
	gameRecords := strings.Split(details, ";")
	reveals := make([]reveal, len(gameRecords))
	for i, gr := range gameRecords {
		reveals[i] = parseReveal(gr)
	}
	return game{
		id:      parseTitle(title),
		reveals: reveals,
	}
}

func parseTitle(s string) int {
	fields := strings.Fields(s)
	n, err := strconv.Atoi(fields[1])
	if err != nil {
		log.Fatalf("bad title %q", s)
	}
	return n
}

func parseReveal(s string) reveal {
	var r reveal
	for _, ch := range strings.Split(s, ",") {
		f := strings.Fields(ch)
		if len(f) != 2 {
			log.Fatalf("bad fields in %q in %q", ch, s)
		}
		n, err := strconv.Atoi(f[0])
		if err != nil {
			log.Fatalf("bad num %q in %q in %q", f[0], ch, s)
		}
		switch f[1] {
		case "blue":
			r.blue = n
		case "red":
			r.red = n
		case "green":
			r.green = n
		default:
			log.Fatalf("bad color in %q in %q in %q", f[1], ch, s)
		}
	}
	return r
}
