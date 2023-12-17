package main

import (
	"fmt"
)

func main() {
	fmt.Println(hash("HASH"))
}

func hash(s string) int {
	current := 0
	for _, r := range s {
		current += int(r)
		current *= 17
		current = current % 256
	}
	return current

}
