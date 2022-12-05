package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// speed code, commit horror

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var elfStack []int
	var accum int

	ln := strings.Split(string(dat), "\n")
	for _, l := range ln {
		if l == "" {
			// new elf. accumulate total, reset and increment
			elfStack = append(elfStack, accum)
			accum = 0
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		accum += n
	}

	// sort it
	sort.Ints(elfStack)
	fmt.Printf("Max Calorie Elf: %v\n", elfStack[len(elfStack)-1])

	// sum t3 elves
	t3e := (elfStack[len(elfStack)-1] + elfStack[len(elfStack)-2] + elfStack[len(elfStack)-3])
	fmt.Printf("Calorie sum T3 elves: %v", t3e)
}
 