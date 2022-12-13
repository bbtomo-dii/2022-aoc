package main

// problem part 2

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	f := loadfl("../input.txt")
	var group [][]string

	// repack into groups
	for i := 0; i < len(f)-3; i += 3 {
		group = append(group, f[i:i+3])
	}

	// run problem space
	var score int
loop:
	for _, p := range group {
		unique := uniqueList(p)
		for _, u := range unique {
			if strings.Contains(p[0], u) && strings.Contains(p[1], u) && strings.Contains(p[2], u) {
				score += toAlphaIndex([]rune(u)[0])
				// jump to next group
				continue loop
			}
		}
	}
	fmt.Printf("Final score is: %d", score)
}

func loadfl(f string) []string {
	dat, err := os.ReadFile(f)
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

func uniqueList(s []string) []string {
	unique := make(map[string]interface{})
	for _, l := range s {
		for _, r := range l {
			rs := string(r)
			if _, ok := unique[string(rs)]; !ok {
				unique[rs] = nil
			}
		}
	}
	var list []string
	for k := range unique {
		list = append(list, k)
	}
	return list
}
