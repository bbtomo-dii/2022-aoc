package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	f := loadfl("input.txt")
	var pack [][]string
	// get splits
	for _, l := range f {
		ll := len(l)
		if ll > 0 {
			c := ll / 2
			pack = append(pack, []string{l[c:], l[:(ll - c)]})
		}
	}
	fmt.Println(pack)

	// test outputs
	fmt.Println(toAlphaIndex([]rune("a")[0]))
	fmt.Println(toAlphaIndex([]rune("A")[0]))

	// run problem space
	var score int
	for _, p := range pack {
		// dummy map for unique matches
		found := make(map[string]interface{})
		for _, i := range p[0] {
			s := string(i)
			if strings.Contains(p[1], s) {
				if _, ok := found[s]; !ok {
					// register the string as found in the map
					found[s] = nil
				}
			}
		}
		// score our map
		for k := range found {
			score += toAlphaIndex([]rune(k)[0])
		}

	}
	fmt.Printf("Final score is: %d", score)
}

func loadfl(f string) []string {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(dat), "\n")
}

// convert to ascii number
func toAlphaIndex(r rune) int {
	i := int(r - '0')
	// for lowercase: offset to 1 index and return
	if !unicode.IsUpper(r) {
		return i - 48
	}
	// for uppercase: offset to 27 index and return
	return (i) + 10
}