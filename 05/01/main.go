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

var containerPos = [9]int{1, 5, 9, 13, 17, 21, 25, 29, 33}
var contBlank = " "

// containermap!
type containers map[int][]string

func main() {

	f := loadfl("../input.txt")

	// first 8 lines are the stack
	st := f[:8]
	// next 11 is moves table
	mv := f[10:]

	// build container state
	// 2x2, row then column
	// populate container state
	containers := stackContainers(st)
	fmt.Println("Initial state:")
	containers.seeContainers()
	mvt := buildMoves(mv)
	fmt.Printf("Total moves: %d\n", len(mvt))
	containers.runSingleMoves(mvt)
	fmt.Println("Final state:")
	containers.seeContainers()

	// reset and run part 2
	containers = stackContainers(st)
	fmt.Println("Initial state:")
	containers.seeContainers()
	containers.runMultiMoves(mvt)
	fmt.Println("Final state:")
	containers.seeContainers()

}

// helper functions

func (c containers) seeContainers() {
	// do visual confirmation on container layout
	for i := 1; i <= len(c); i++ {
		fmt.Printf("column number: %d, containers: %v\n", i, c[i])
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

// container and move setup process

func stackContainers(st []string) containers {
	// invert so we go column -> row
	// init columsn
	c := make(map[int][]string)

	// take raw line and pluck out container ID
	// iterate slice backwards so containers are "stacked"
	// bottom to top in slice
	for i := len(st) - 1; i > -1; i-- {
		// extract columns
		for n, p := range containerPos {
			nc := string(st[i][p])
			if nc != contBlank {
				c[n+1] = append(c[n+1], nc)
			}
		}
	}
	return c
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

// single container movement code

func (c containers) singleMove(m move) {
	for i := 0; i < m.qty; i++ {
		// if any containers left:
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

func (c containers) runSingleMoves(mv []move) {
	for _, m := range mv {
		c.singleMove(m)
	}
}

// multi container movement code

func (c containers) multiMove(m move) {
	p := len(c[m.from]) - m.qty
	if p < -1 {
		panic("not enough containers to take")
	}
	c[m.to] = append(c[m.to], c[m.from][p:]...)
	c[m.from] = c[m.from][:p]
}

func (c containers) runMultiMoves(mv []move) {
	for _, m := range mv {
		c.multiMove(m)
	}
}
