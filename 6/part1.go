package main

import (
	"fmt"
	"math"
)

type race struct {
	time, distance float64
}

func main() {
	races := []race{
		// {7, 9}, {15, 40}, {30, 200},
		{54, 446}, {81, 1292}, {70, 1035}, {88, 1007},
	}

	reply := 1
	for _, r := range races {
		reply *= nWins(r)
	}
	fmt.Println(reply)
}

func nWins(r race) int {
	// let t = time holding button
	// distance = t * (rtime-t)
	// t^2 -rtime.t + d = 0
	// quadratic formula -> rtime +/-
	d := math.Sqrt(r.time*r.time - 4*r.distance)
	a, b := math.Ceil((r.time+d)/2)-1, math.Floor((r.time-d)/2)+1
	fmt.Println(r, a, b)
	return int(a-b) + 1
}
