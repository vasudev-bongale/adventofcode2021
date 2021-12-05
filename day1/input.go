package main

import (
	"strconv"
	"bufio"
	"os"
	"log"
)

// GetInputData gets data for day1 puzzle
func GetInputData() []int {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := []int{}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		input = append(input, num)
	}
	return input
}
