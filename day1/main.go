package main

import (
	"log"
)

func main() {
	input := GetInputData()
	log.Printf("Processing %d lines of input\n", len(input))

	// Number of increments counted so far
	var depthIncrements int
	depth := input[0]

	for _, i := range input {
		if i - depth > 0 {
			depthIncrements++
		}
		depth = i
	}

	log.Printf("%v measurements are larger than the previous measurement.\n", depthIncrements)
}