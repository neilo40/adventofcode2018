package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	ids := strings.Split(string(input), "\n")
	var twos = 0
	var threes = 0
	for _, id := range ids {
		if uniqueCount(id, 2) {
			twos++
		}
		if uniqueCount(id, 3) {
			threes++
		}
	}
	fmt.Printf("checksum = %d\n", twos*threes)
}

func uniqueCount(s string, n int) bool {
	var freqCount = make(map[rune]int)
	for _, letter := range s {
		if _, ok := freqCount[letter]; ok {
			freqCount[letter]++
		} else {
			freqCount[letter] = 1
		}
	}
	for _, v := range freqCount {
		if v == n {
			return true
		}
	}
	return false
}
