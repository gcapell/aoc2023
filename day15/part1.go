package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	seq, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	s := &store{}
	for _, chunk := range strings.Split(strings.TrimSpace(string(seq)), ",") {
		s.apply(chunk)
		s.report()
		fmt.Println()
	}
	fmt.Println("totalPower", s.totalPower())
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

type box struct {
	lenses []*lens
}

func (b *box) report() string {
	var chunks []string
	for _, l := range b.lenses {
		chunks = append(chunks, fmt.Sprintf("[%s %d]", l.name, l.focalLength))
	}
	return strings.Join(chunks, " ")
}

type lens struct {
	name        string
	focalLength int
}

type store struct {
	boxes [256]box
}

func (s *store) report() {
	for i, b := range s.boxes {
		if len(b.lenses) > 0 {
			fmt.Printf("Box %d: %s\n", i, b.report())
		}
	}
}

func (b *box) update(i string, f int) {
	for _, l := range b.lenses {
		if l.name == i {
			l.focalLength = f
			return
		}
	}
	b.lenses = append(b.lenses, &lens{name: i, focalLength: f})
}

func (b *box) rm(name string) {
	for i, l := range b.lenses {
		if l.name == name {
			b.lenses = slices.Delete(b.lenses, i, i+1)
		}
	}
}

func (b *box) totalPower(n int) int {
	total := 0
	for slot, lens := range b.lenses {
		power := (n + 1) * (slot + 1) * lens.focalLength
		fmt.Printf("%s: %d (box %d) * %d (slot) * %d (focal) = %d\n",
			lens.name, n+1, n, slot+1, lens.focalLength, power)
		total += power
	}
	return total
}

func (s *store) apply(instruction string) {
	if strings.HasSuffix(instruction, "-") {
		s.minus(strings.TrimSuffix(instruction, "-"))
		return
	}
	label, focus, ok := strings.Cut(instruction, "=")
	if !ok {
		panic(instruction)
	}
	focusN, err := strconv.Atoi(focus)
	if err != nil {
		panic(instruction)
	}
	s.equals(label, focusN)
}

func (s *store) minus(i string) {
	fmt.Println("minus", i)
	s.boxes[hash(i)].rm(i)
}

func (s *store) equals(i string, f int) {
	fmt.Println("equals", i, f)
	s.boxes[hash(i)].update(i, f)
}

func (s *store) totalPower() int {
	total := 0
	for n, b := range s.boxes {
		total += b.totalPower(n)
	}
	return total
}
