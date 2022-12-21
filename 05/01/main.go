package main

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"os"
	"strconv"
	"strings"
)

type move struct {
	qty, from, to int
}

var cratePos = [9]int{1, 5, 9, 13, 17, 21, 25, 29, 33}
var contBlank = " "

// cratemap!
type crates map[int][]string

func main() {

	f := loadfl("../input.txt")

	// first 8 lines are the stack
	st := f[:8]
	// next 11 is moves table
	mv := f[10:]

	// build crate state
	// 2x2, row then column
	// populate crate state
	crates := stackCrates(st)
	fmt.Println("Initial state:")
	crates.seeCrates()
	mvt := buildMoves(mv)
	fmt.Printf("Total moves: %d\n", len(mvt))
	crates.runSingleMoves(mvt)
	fmt.Println("Final state:")
	crates.seeCrates()

	// reset and run part 2
	crates = stackCrates(st)
	fmt.Println("Initial state:")
	crates.seeCrates()
	crates.runMultiMoves(mvt)
	fmt.Println("Final state:")
	crates.seeCrates()

}

func (c crates) singleMove(m move) {
	for i := 0; i < m.qty; i++ {
		// if any crates left:
		l := len(c[m.from]) - 1
		if l > -1 {
			// move across
			c[m.to] = append(c[m.to], c[m.from][l])
			// delete
			c[m.from] = c[m.from][:l]
		} else {
			panic("attempt to pull from empty stack")
		}
	}
}

// remember: lower inclusive, higher exclusive
func (c crates) multiMove(m move) {
	p := len(c[m.from]) - m.qty
	if p < -1 {
		panic("not enough crates to take")
	}
	c[m.to] = append(c[m.to], c[m.from][p:]...)
	c[m.from] = c[m.from][:p]
}

func (c crates) runSingleMoves(mv []move) {
	for _, m := range mv {
		fmt.Printf("Exec move: %v\n", m)
		c.singleMove(m)
	}
}

func (c crates) runMultiMoves(mv []move) {
	for _, m := range mv {
		fmt.Printf("Exec move: %v\n", m)
		c.multiMove(m)
		fmt.Println("peek:")
		c.seeCrates()
	}
}

func buildMoves(st []string) []move {
	var mv []move
	for _, m := range st {
		if m != "" {
			a := strings.Split(m, " ")
			c, err := strconv.Atoi(a[1])
			explode(err)
			f, err := strconv.Atoi(a[3])
			explode(err)
			t, err := strconv.Atoi(a[5])
			explode(err)
			mv = append(mv, move{
				qty:  c,
				from: f,
				to:   t,
			})
		}
	}
	return mv
}

func stackCrates(st []string) crates {
	// invert so we go column -> row
	// init columsn
	c := make(map[int][]string)

	// take raw line and pluck out container ID
	// iterate slice backwards so crates are "stacked"
	// bottom to top in slice
	for i := len(st) - 1; i > -1; i-- {
		// extract columns
		for n, p := range cratePos {
			nc := string(st[i][p])
			if nc != contBlank {
				c[n+1] = append(c[n+1], nc)
			}
		}
	}
	return c
}

func (c crates) seeCrates() {
	// do visual confirmation on crate layout
	for i := 1; i <= len(c); i++ {
		fmt.Printf("column number: %d, crates: %v\n", i, c[i])
	}
}

func loadfl(f string) []string {
	dat, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(dat), "\n")
}

func explode(err error) {
	if err != nil {
		panic(err)
	}
}
