package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rng struct {
	l, r int
}

func main() {

	olCount := 0
	tcCount := 0

	f := loadfl("../input.txt")
	for _, v := range f {
		if v != "" {
			p := strings.Split(v, ",")
			r1 := getNums(p[0])
			r2 := getNums(p[1])
			if r1.inside(r2) || r2.inside(r1) {
				olCount++
			}
			if r1.touches(r2) || r2.touches(r1) {
				tcCount++
			}
		}
	}

	fmt.Printf("Overlap Count: %d\n", olCount)
	fmt.Printf("Touch Count: %d\n", tcCount)
}

func (l rng) touches(r rng) bool {
	if ((l.l >= r.l) && (l.l <= r.r)) || ((l.r >= r.l) && (l.r <= r.r)) {
		return true
	}
	return false
}

func (l rng) inside(r rng) bool {
	if (l.l >= r.l) && (l.r <= r.r) {
		return true
	}
	return false
}

func getNums(s string) rng {
	p := strings.Split(s, "-")
	l, err := strconv.Atoi(p[0])
	explode(err)
	r, err := strconv.Atoi(p[1])
	explode(err)
	return rng{l, r}
}

func explode(err error) {
	if err != nil {
		panic(err)
	}
}
func loadfl(f string) []string {
	dat, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(dat), "\n")
}
