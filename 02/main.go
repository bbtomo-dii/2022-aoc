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

func main() {
	s := score{0, 3, 6}
	t := tool{1, 2, 3}
	// outcome map - w/l/d for RHS (responding) player
	o := map[string]int{
		// losses
		"RS": s.l + t.s,
		"PR": s.l + t.r,
		"SP": s.l + t.p,
		// draws
		"RR": s.d + t.r,
		"PP": s.d + t.p,
		"SS": s.d + t.s,
		// wins
		"RP": s.w + t.p,
		"PS": s.w + t.s,
		"SR": s.w + t.r,
	}

	fmt.Println("Score table:")
	for k, v := range o {
		fmt.Printf("combo: %v, value: %v\n", k, v)
	}

	// look ma, boilerplate
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	r := string(dat)
	// go away stupid cipher
	r = strings.ReplaceAll(r, "X", "R")
	r = strings.ReplaceAll(r, "Y", "P")
	r = strings.ReplaceAll(r, "Z", "S")
	r = strings.ReplaceAll(r, "A", "R")
	r = strings.ReplaceAll(r, "B", "P")
	r = strings.ReplaceAll(r, "C", "S")

	ln := strings.Split(r, "\n")
	var moves []string
	for _, l := range ln {
		if l != "" {
			m := strings.Split(l, " ")
			moves = append(moves, fmt.Sprintf("%s%s", m[0], m[1]))
		}
	}

	// iterate comps
	tl := 0
	for _, r := range moves {
		// aggregate scores
		tl += o[r]

	}
	fmt.Printf("Final score: %v\n", tl)

}
