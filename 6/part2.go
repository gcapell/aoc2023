package main

import (
	"fmt"
	"math"
)

type race struct {
	time, distance float64
}

func main() {
	fmt.Println(nWins(race{54817088, 446129210351007}))
}

func nWins(r race) int {
	// let t = time holding button
	// distance = t * (rtime-t)
	// t^2 -rtime.t + d = 0
	// quadratic formula -> rtime +/-
	d := math.Sqrt(r.time*r.time - 4*r.distance)
	fmt.Println(d)
	a, b := math.Ceil((r.time+d)/2)-1, math.Floor((r.time-d)/2)+1
	fmt.Println(r, a, b)
	return int(a-b) + 1
}
