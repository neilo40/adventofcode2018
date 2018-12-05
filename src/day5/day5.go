package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	rawInput, _ := ioutil.ReadFile("input.txt")
	var originalInput = string(rawInput)
	var alphabet = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's',
		't', 'u', 'v', 'w', 'x', 'y', 'z'}
	var minLength = 50000
	for _, letter := range alphabet {
		var input = strings.Replace(originalInput, string(letter), "", -1)
		input = strings.Replace(input, strings.ToUpper(string(letter)), "", -1)
		var finalLength = checkInput(input)
		if finalLength < minLength {
			minLength = finalLength
		}
	}
	fmt.Println(minLength)
}

func checkInput(input string) int {
	var foundMatch = false
	for i := 0; i < len(input)-1; i++ {
		if isSameLetterDiffCase(string(input[i]), string(input[i+1])) {
			input = input[0:i] + input[i+2:]
			foundMatch = true
		}
	}
	if foundMatch {
		return checkInput(input)
	} else {
		return len(input)
	}
}

func isSameLetterDiffCase(first string, second string) bool {
	var firstIsUpper = false
	var secondIsUpper = false
	if strings.ToUpper(first) == first {
		firstIsUpper = true
	}
	if strings.ToUpper(second) == second {
		secondIsUpper = true
	}
	if firstIsUpper != secondIsUpper {
		if strings.ToUpper(first) == strings.ToUpper(second) {
			return true
		}
	}
	return false
}
