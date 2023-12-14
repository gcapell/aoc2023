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
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	paras := strings.Split(string(data), "\n\n")

	ranges := intervals(ints(strings.TrimPrefix(paras[0], "seeds: ")))
	//fmt.Println("ranges", ranges)

	for _, para := range paras[1:] {
		ranges = transform(ranges, parseMap(para))
	}
	fmt.Println(lowestStart(ranges))
}

func lowestStart(intervals []interval) int {
	var lowest int
	for pos, i := range intervals {
		if pos == 0 || i.start < lowest {
			lowest = i.start
		}
	}
	return lowest
}
func transform(ranges []interval, transforms []intervalTransform) []interval {
	sort.Slice(transforms, func(a, b int) bool { return transforms[a].start < transforms[b].start })
	var reply []interval
	for _, r := range ranges {
		transformed := transformOne(r, transforms)
		//fmt.Printf("tx(%v, %v)->%v\n", r, transforms, transformed)
		reply = append(reply, transformed...)
	}
	//fmt.Printf("transform %v with %v => %v\n", ranges, transforms, reply)
	return reply
}

func transformOne(r interval, transforms []intervalTransform) []interval {
	var reply []interval
	// skip transforms before us
	// ... yes we could binary search, but we won't today
	for len(transforms) > 0 && transforms[0].end <= r.start {
		transforms = transforms[1:]
	}
	for len(transforms) > 0 && transforms[0].start < r.end {
		t := transforms[0]
		transforms = transforms[1:]

		if t.start <= r.start {
			if t.end < r.end {
				// bite from the front
				reply = append(reply, interval{
					start: r.start + t.offset,
					end:   t.end + t.offset,
				})
				r = interval{t.end, r.end}
			} else {
				// complete overlap
				reply = append(reply, interval{
					start: r.start + t.offset,
					end:   r.end + t.offset,
				})
				return reply
			}
		} else {
			if t.end < r.end {
				// bite from the middle
				reply = append(reply, interval{
					start: r.start,
					end:   t.start,
				}, interval{
					start: t.start + t.offset,
					end:   t.end + t.offset,
				})
				r = interval{t.end, r.end}
			} else {
				// bite from the end
				reply = append(reply, interval{
					start: r.start,
					end:   t.start,
				}, interval{
					start: t.start + t.offset,
					end:   r.end + t.offset,
				})
				return reply
			}
		}
	}
	reply = append(reply, r)
	return reply
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

type interval struct {
	start, end int
}

func intervals(ns []int) []interval {
	var reply []interval
	for j := 0; j*2 < len(ns); j++ {
		reply = append(reply, interval{
			start: ns[j*2],
			end:   ns[j*2] + ns[j*2+1],
		})
	}
	return reply
}

type intervalTransform struct {
	start, end, offset int
}

func parseMap(para string) []intervalTransform {
	_, mapSrc, ok := strings.Cut(para, "map:")
	if !ok {
		panic(para)
	}
	var reply []intervalTransform
	for _, line := range strings.Split(mapSrc, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		reply = append(reply, parseInterval(line))
	}
	return reply
}

func parseInterval(line string) intervalTransform {
	nums := ints(line)
	if len(nums) != 3 {
		panic(line)
	}
	return intervalTransform{
		start:  nums[1],
		end:    nums[1] + nums[2],
		offset: nums[0] - nums[1],
	}
}
