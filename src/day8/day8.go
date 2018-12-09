package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"strconv"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	atoms := strings.Split(string(input), " ")
	for _, atom := range atoms {
		a, _ := strconv.ParseInt(atom, 10, 64)
		fmt.Println(a)
	}
}
