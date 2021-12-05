package common

import (
	"bufio"
	"os"
	"log"
)

// GetInputData gets data for day1 puzzle
func GetInputDataFromFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	input := []string{}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}
