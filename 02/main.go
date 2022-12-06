package main

import (
	"fmt"
	"os"
	"strings"
)

// speed code, commit horror

type score struct {
	l, d, w int
}
type tool struct {
	r, p, s int
}

// remap map
var r = map[string]string{
	"A": "R",
	"B": "P",
	"C": "S",
	"X": "R",
	"Y": "P",
	"Z": "S",
}

var (
	s = score{0, 3, 6}
	t = tool{1, 2, 3}
	// outcome map - w/l/d for RHS (responding) player
	loss = map[string]int{
		// losses
		"RS": s.l + t.s,
		"PR": s.l + t.r,
		"SP": s.l + t.p,
	}
	draw = map[string]int{
		// draws
		"RR": s.d + t.r,
		"PP": s.d + t.p,
		"SS": s.d + t.s,
	}
	win = map[string]int{
		// wins
		"RP": s.w + t.p,
		"PS": s.w + t.s,
		"SR": s.w + t.r,
	}
)

func main() {
	// look ma, boilerplate
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	ln := strings.Split(string(dat), "\n")
	var moves []string
	for _, l := range ln {
		if l != "" {
			m := strings.Split(l, " ")
			// remap a/b/c x/y/z to their "intuitive" names, but store
			// the x/y value for the second part of the challenge
			moves = append(moves, fmt.Sprintf("%s%s%s", r[m[0]], r[m[1]], m[1]))
		}
	}

	// iterate comps for total score
	win := 0
	rig := 0

	for _, r := range moves {
		// aggregate winning scores
		fmt.Println(r)
		win += getOutcome(r)

		// aggregate rigged scores
		rig += getRigged(r)

	}

	fmt.Printf("Stage 1 - winning score: %v\n", win)
	fmt.Printf("Stage 2 - rigged score: %v\n", rig)

}

func getOutcome(s string) int {
	o := s[:2]
	if val, ok := win[o]; ok {
		return val
	}
	if val, ok := loss[o]; ok {
		return val
	}
	if val, ok := draw[o]; ok {
		return val
	}
	panic(fmt.Sprintf("%s is not a valid competition outcome", s))
}

func getRigged(s string) int {
	c := string(s[2])
	i := string(s[0])
	switch c {
	case "X":
		return loopMap(i, loss)
	case "Y":
		return loopMap(i, draw)
	case "Z":
		return loopMap(i, win)
	}
	panic(fmt.Sprintf("%s is not a valid match", c))
}

func loopMap(i string, m map[string]int) int {
	for k, v := range m {
		if string(k[0]) == i {
			return v
		}
	}
	panic(fmt.Sprintf("%s is not a valid input", i))
}