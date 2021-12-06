package main

import (
	"adventofcode2021/common"
	"fmt"
	"strconv"
)

const (
	wordSize = 12
)

func main() {

	input := common.GetInputDataFromFile("../common/day3_data.txt")
	binaryValues := make([]int64, len(input))

	for j, d := range input {
		b, _ := strconv.ParseInt(d, 2, wordSize + 1)
		binaryValues[j] = b
	}

	var gammaRate, epsilonRate uint16

	ones := countBits(binaryValues, wordSize, 1)

	for i := 0; i < wordSize; i++ {
		if ones[i] > len(binaryValues) / 2 {
			gammaRate |= 1 << (wordSize - i - 1)
		}
	}
	epsilonRate = 0xFFF ^ gammaRate

	fmt.Printf("GammaRate %05b : %v, EpsilonRate %05b : %v\n", gammaRate, gammaRate, epsilonRate, epsilonRate)
	fmt.Printf("Part 1 Answer: %v\n", int(gammaRate) * int(epsilonRate))

	i := 0
	oxyRatings := binaryValues
	for len(oxyRatings) > 1 {
		numOnes := countBits(oxyRatings, wordSize, 1)
		numZeros := countBits(oxyRatings, wordSize, 0)
		if numOnes[i] > numZeros[i] || numOnes[i] == numZeros[i] {
			oxyRatings = filterOnBit(oxyRatings, i, 1)
		} else {
			oxyRatings = filterOnBit(oxyRatings, i, 0)
		}
		i++
	}

	i = 0
	co2Ratings := binaryValues
	for len(co2Ratings) > 1 {
		numOnes := countBits(co2Ratings, wordSize, 1)
		numZeros := countBits(co2Ratings, wordSize, 0)
		if numOnes[i] > numZeros[i] || numOnes[i] == numZeros[i] {
			co2Ratings = filterOnBit(co2Ratings, i, 0)
		} else {
			co2Ratings = filterOnBit(co2Ratings, i, 1)
		}
		i++
	}

	fmt.Printf("OxyRating: %d, Co2Rating: %d\n", oxyRatings[0], co2Ratings[0])
	fmt.Println("Part 2 Answer: ", co2Ratings[0] * oxyRatings[0])

}

func filterOnBit(data []int64 , position int, bitvalue int64) (output []int64) {
	for _, value := range data {
		if (value >> (wordSize - position - 1)) & 1 == bitvalue {
			output = append(output, value)
		} 
	}
	return
}

func countBits(data []int64, size int, bit int64) (numOnes []int) {
	numOnes = make([]int, size)
	for _, d := range data {
		for i := 0; i < size; i++ {
			if (d >> i) & 1 == bit {
				numOnes[wordSize - i - 1] += 1
			}
		}
	}
	return
}