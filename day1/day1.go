package main

import (
	"log"
	"strconv"
	common "adventofcode2021/common"
)

func main() {
	input := common.GetInputDataFromFile("../common/day1_data.txt")

	// day1 data is all ints
	data := []int{}
	for _, i := range input {
		num, _ := strconv.Atoi(i)
		data = append(data, num)
	}

	log.Printf("Processing %d lines of input\n", len(input))

	// puzzle1
	log.Printf("%v measurements are larger than the previous measurement.\n", numIncrements(data))

	// puzzle2
	slidingWindow := 3
	slidingSumData := slidingSum(slidingWindow, data)
	log.Printf("%v sums are larger than the previous sum.\n", numIncrements(slidingSumData))
}

// numIncrements traverses an array of ints and counts if the next value is greater than the previous value
func numIncrements(data []int) (increments int) {
	// Number of increments counted so far
	depth := data[0]
	for _, i := range data {
		if i - depth > 0 {
			increments++
		}
		depth = i
	}
	return
}

// slidingSum generates a slice with sliding sum of a given window
func slidingSum(window int, data []int) []int {
	sum := make([]int, len(data) - window + 1)
	for d := 0; d < len(sum); d++ {
		currentSum := 0
		for w := 0; w < window; w++ {
			currentSum += data[d + w]
		}
		sum[d] = currentSum
	}
	return sum
}