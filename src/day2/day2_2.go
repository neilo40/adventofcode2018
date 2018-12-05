package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	ids := strings.Split(string(input), "\n")
	for i, id1 := range ids {
		for _, id2 := range ids[i+1:] {
			if d, c := diffByOne(id1, id2); d {
				fmt.Printf("%s and %s diff by one char.  Common string: %s\n", id1, id2, c)
			}
		}
	}
}

func diffByOne(s1 string, s2 string) (bool, string) {
	var diffs = 0
	var sameChars = ""
	for i, c := range s1 {
		if string(c) != string(s2[i]) {
			diffs++
		} else {
			sameChars += string(c)
		}
	}
	return diffs == 1, sameChars
}
