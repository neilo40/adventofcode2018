package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var frequency int64
	var seenFrequencies = make(map[int64]bool)
	seenFrequencies[0] = true
	input, _ := ioutil.ReadFile("input.txt")
	instructions := strings.Split(string(input), "\n")
	var foundDouble = false
	for foundDouble == false {
		instr := instructions
		for _, valueString := range instr {
			valueNumber, _ := strconv.ParseInt(valueString, 10, 64)
			frequency += valueNumber
			if seenFrequencies[frequency] {
				fmt.Printf("Seen frequency %d twice\n", frequency)
				foundDouble = true
				break
			} else {
				seenFrequencies[frequency] = true
			}
		}
	}
	fmt.Printf("Final Frequency is %d\n", frequency)
}
