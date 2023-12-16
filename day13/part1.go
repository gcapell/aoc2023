package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	total := 0
	for _, b := range boards() {
		total += score(b)
	}
	fmt.Println("total", total)
}

type board string

func boards() []board {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	var reply []board
	for _, s := range strings.Split(string(data), "\n\n") {
		reply = append(reply, parseBoard(s))
	}
	return reply
}

func parseBoard(s string) board {
	return board(s)
}

func score(b board) int {
	rows := strings.Fields(string(b))
	//fmt.Printf("score\n%v\n", rows)

	total := 0
	for r := 1; r < len(rows); r++ {
		if mirror(rows, r) {
			total += 100 * r
		}
	}
	cols := make([]string, len(rows[0]))
	for c := range rows[0] {
		letters := make([]byte, len(rows))
		for i, r := range rows {
			letters[i] = r[c]
		}
		cols[c] = string(letters)
	}

	for c := 1; c < len(cols); c++ {
		if mirror(cols, c) {
			total += c
		}
	}

	return total
}

func mirror(rows []string, r int) bool {
	//fmt.Println("mirror", r)
	for down, up := r, r-1; (down < len(rows)) && up >= 0; down, up = down+1, up-1 {
		//fmt.Println(down, up)
		if rows[down] != rows[up] {
			return false
		}
	}
	return true
}
