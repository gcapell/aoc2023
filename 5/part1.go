package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	paras := strings.Split(string(data), "\n\n")

	seeds := ints(strings.TrimPrefix(paras[0], "seeds: "))
	var transformers []transformer
	for _, para := range paras[1:] {
		transformers = append(transformers, parseMap(para))
	}
	var lowest int
	for i, s := range seeds {
		loc := transform(transformers, s)
		if i == 0 || loc < lowest {
			lowest = loc
		}
	}
	fmt.Println("lowest loc", lowest)

}

func transform(ts []transformer, n int) int {

	for _, t := range ts {
		m := t.transform(n)

		n = m
	}
	return n
}

func ints(line string) []int {
	ss := strings.Fields(line)
	reply := make([]int, len(ss))
	for i, s := range ss {
		var err error
		reply[i], err = strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
	}
	return reply
}

func (t transformer) transform(n int) int {
	for _, i := range t {
		if m, ok := i.transform(n); ok {
			return m
		}
	}
	return n
}

type interval struct {
	start, end, offset int
}

func (i interval) transform(n int) (int, bool) {
	if n >= i.start && n < i.end {
		return n + i.offset, true
	}
	return 0, false
}

type transformer []interval

func parseMap(para string) transformer {
	_, mapSrc, ok := strings.Cut(para, "map:")
	if !ok {
		panic(para)
	}
	var reply []interval
	for _, line := range strings.Split(mapSrc, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		reply = append(reply, parseInterval(line))
	}
	return reply
}

func parseInterval(line string) interval {
	nums := ints(line)
	if len(nums) != 3 {
		panic(line)
	}
	return interval{
		start:  nums[1],
		end:    nums[1] + nums[2],
		offset: nums[0] - nums[1],
	}
}
