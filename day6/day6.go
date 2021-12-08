package main

import (
	"fmt"
	"strings"
	"strconv"
)

const (
	days = 256
)

func main() {

	initStr := "5,1,1,3,1,1,5,1,2,1,5,2,5,1,1,1,4,1,1,5,1,1,4,1,1,1,3,5,1,1,1,1,1,1,1,1,1,4,4,4,1,1,1,1,1,4,1,1,1,1,1,5,1,1,1,4,1,1,1,1,1,3,1,1,4,1,4,1,1,2,3,1,1,1,1,4,1,2,2,1,1,1,1,1,1,3,1,1,1,1,1,2,1,1,1,1,1,1,1,4,4,1,4,2,1,1,1,1,1,4,3,1,1,1,1,2,1,1,1,2,1,1,3,1,1,1,2,1,1,1,3,1,3,1,1,1,1,1,1,1,1,1,3,1,1,1,1,3,1,1,1,1,1,1,2,1,1,2,3,1,2,1,1,4,1,1,5,3,1,1,1,2,4,1,1,2,4,2,1,1,1,1,1,1,1,2,1,1,1,1,1,1,1,1,4,3,1,2,1,2,1,5,1,2,1,1,5,1,1,1,1,1,1,2,2,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,4,1,1,1,1,1,3,1,1,5,1,1,1,1,5,1,4,1,1,1,4,1,3,4,1,4,1,1,1,1,1,1,1,1,1,3,5,1,3,1,1,1,1,4,1,5,3,1,1,1,1,1,5,1,1,1,2,2"
	
	initState := []int{}
	for _, s := range strings.Split(initStr, ",") {
		state, _ := strconv.Atoi(s)
		initState = append(initState, state)
	}

	fishStates := [9]int{}
	for _, s := range initState {
		fishStates[s]++
	}

	for d := 1; d <= days; d++ {
		mature := fishStates[0]
		for i := 0; i < 8; i++ {
			fishStates[i] = fishStates[i + 1]
		}
		fishStates[8] = mature
		fishStates[6] += mature
	}

	sum := 0
	for _, s := range(fishStates) {
		sum += s
	}

	fmt.Printf("After %d days, there are %d lanternfish\n", days, sum)

}