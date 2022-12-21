package main

import (
	"2022-aoc/util"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	f := util.LoadFl("../input.txt")
	m := findMarker(f[0], 4)
	spew.Dump(m[0])
	m = findMarker(f[0], 14)
	spew.Dump(m[0])

}

func findMarker(s string, w int) []int {
	// walk chunks
	var reg []int
	for i := w; i < len(s)-w; i++ {
		// capture the previous w chars
		c := s[i-w : i]
		if allUnique(c) {
			// is a marker. note the pos
			reg = append(reg, i)
		}
	}
	return reg
}

func allUnique(s string) bool {
	// uninitialized rune
	var prev rune
	s = util.SortString(s)
	for _, c := range s {
		if prev == c {
			return false
		}
		prev = c
	}
	return true
}
