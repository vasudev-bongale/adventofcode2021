package main

import (
	"fmt"
	"adventofcode2021/common"
	"strings"
	"strconv"
)

func main() {
	input := common.GetInputDataFromFile("../common/day2_data.txt")

	var hPos, depth, aim int

	for _, way := range input {
		command := strings.Split(way, " ")

		direction := command[0]
		value, _ := strconv.Atoi(command[1])

		switch direction {
		case "forward":
			hPos += value
			depth += aim * value
		case "down":
			// depth += value
			aim += value
		case "up":
			// depth -= value
			aim -= value
		}
	}

	fmt.Printf("Horizontal Postion: %v, Depth: %v. Product: %v\n", hPos, depth, hPos * depth)
}