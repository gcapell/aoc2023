package main

import (
	"fmt"
	"strings"
)

func main() {
	for _, chunk := range strings.Split("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7", ",") {
		fmt.Println(chunk, hash(chunk))
	}
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
