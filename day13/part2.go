package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	total := 0
	for _, b := range boards() {
		total += fixSmudge(b)
	}
	fmt.Println("total", total)
}

type board [][]byte

func boards() []board {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	var reply []board
	for _, s := range bytes.Split(data, []byte("\n\n")) {
		reply = append(reply, board(bytes.Fields(s)))
	}
	return reply
}

func fixSmudge(b board) int {
	for r := 1; r < len(b); r++ {
		if fixSmudge1(b, r, true, len(b)) {
			return r * 100
		}
	}
	for c := 1; c < len(b[0]); c++ {
		if fixSmudge1(b, c, false, len(b[0])) {
			return c
		}
	}
	panic(fmt.Sprint(b))
}

func fixSmudge1(b board, r int, useRow bool, max int) bool {
	//fmt.Println("mirror", r)
	allowDiff := true

	for down, up := r, r-1; (down < max) && up >= 0; down, up = down+1, up-1 {
		sr := same(b, down, up, useRow, allowDiff)
		if !sr.ok {
			return false
		}
		if sr.foundDiff {
			allowDiff = false
		}
	}
	return !allowDiff
}

type sameResult struct {
	foundDiff bool
	ok        bool
}

func same(b board, m, n int, useRow, allowDiff bool) sameResult {
	result := sameResult{ok: true}
	if useRow {
		for j := 0; j < len(b[0]); j++ {
			if b[m][j] != b[n][j] {
				if allowDiff {
					result.foundDiff = true
					allowDiff = false
				} else {
					return sameResult{ok: false}
				}
			}
		}
		return result
	}
	for j := 0; j < len(b); j++ {
		if b[j][m] != b[j][n] {
			if allowDiff {
				result.foundDiff = true
				allowDiff = false
			} else {
				return sameResult{ok: false}
			}
		}
	}
	return result
}
