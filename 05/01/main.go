package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	f := loadfl("../input.txt")

	// first 9 lines are the stack
	st := f[:9]
	// discard 10, take moves
	mv := f[11:]

	// build crate state
	// 2x2, row then column
	crates := [8][9]*rune
	// populate crate state
	for _, v := range st {
		// string to runes
		r := []rune(v)
		// extract at column pos

		crates = append(crates,[9]rune{r[2],r[6],r[10],r[14],r[18],r[22],r[]})
		
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
