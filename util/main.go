package util

import (
	"os"
	"sort"
	"strings"
)

func LoadFl(f string) []string {
	dat, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(dat), "\n")
}

func Die(err error) {
	if err != nil {
		panic(err)
	}
}

// simple but expensive sort, should use runes
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}