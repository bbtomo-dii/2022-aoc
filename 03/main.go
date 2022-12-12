package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	f := loadfl("input.txt")
	var pack [][]string
	// get splits
	for _, l := range f {
		ll := len(l)
		c := ll / 2
		pack = append(pack, []string{l[c:], l[:(ll - c)]})
	}
	fmt.Println(pack)
	fmt.Println(toAlphaIndex([]rune("a")[0]))
	fmt.Println(toAlphaIndex([]rune("b")[0]))

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
	return int(r-'0') - 48
}