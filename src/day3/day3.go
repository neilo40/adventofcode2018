package main

import (
	"io/ioutil"
	"strings"
)

// Vertex for holding coords
type Vertex struct {
	X int
	Y int
}

// Claim as per input
type Claim struct {
	ID          int
	TopLeft     Vertex
	BottomRight Vertex
}

func main() {
	var fabricClaims = make(map[Vertex][]int)
	input, _ := ioutil.ReadFile("input.txt")
	claims := strings.Split(string(input), "\n")
	for _, claim := range claims {

	}
}
