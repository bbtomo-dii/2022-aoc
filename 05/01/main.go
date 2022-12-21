package main

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"os"
	"sort"
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
	mv := f[11:]

	// build crate state
	// 2x2, row then column
	// populate crate state
	crates := stackCrates(st)

	// inspect
	crates.seeCrates()
	mvt := buildMoves(mv)
	crates.runMoves(mvt)
	crates.seeCrates()

}

func (c *crates) runMoves(mv []move) {
	for _, m := range mv {
		conts := c.getContainers(m.from, m.qty)
		c.putContainers(m.to, conts)
		fmt.Println("Crates update:")
		c.seeCrates()
	}
}

func (c crates) putContainers(col int, conts []string) {
	// reverse slice
	conts = sort.StringSlice(conts)
	c[col] = append(c[col], conts...)
}

func (c crates) getContainers(col, num int) []string {
	i := len(c[col]) - num

	// pop containers out
	conts := c[col][i:]
	// delete from slice
	c[col] = c[col][:i]
	return conts
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
	for col, cont := range c {
		var cr []string
		for _, val := range cont {
			cr = append(cr, val)
		}
		fmt.Printf("column number: %d, crates: %s\n", col, cr)
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
