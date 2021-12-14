package main

import (
	"fmt"
	"adventofcode2021/common"
	"strings"
	"strconv"
)

func main() {

	data := common.GetInputDataFromFile("../common/day8_data.txt")

	// Number of 1, 4, 7, and 8 in the output
	part1num := 0
	part2sum := 0
	
	for _, d := range data {
		out := strings.Split(d, " | ")[1]
		patterns := strings.Split(d, " | ")[0]
		
		f := make(map[int]string)

		for _, p := range strings.Fields(patterns) {
			if len(p) == 2 {
				f[1] = p
			}
			if len(p) == 4 {
				f[4] = p
			}
		}

		// Identify a 6
		for _, p := range strings.Fields(patterns) {
			if len(p) == 6 {
				if decodeSixes(p, f[1], f[4]) == 6 {
					f[6] = p
				}
			}
		}

		values := strings.Fields(out)
		digits := []int{}
		for _, v := range values {
			size := len(v)
			switch size {
			case 2:
				part1num++
				digits = append(digits, 1)
			case 5:
				digits = append(digits, decode(v, f[1], f[6]))
			case 4:
				part1num++
				digits = append(digits, 4)
			case 6:
				digits = append(digits, decodeSixes(v, f[1], f[4]))
			case 3:
				part1num++
				digits = append(digits, 7)
			case 7:
				part1num++
				digits = append(digits, 8)
			}
		}

		

		output := ""
		for _, d := range digits {
			output += fmt.Sprint(d)
		}

		result, _ := strconv.Atoi(output)
		part2sum += result

		// fmt.Println(values, "|", result)

	}
	fmt.Println("[Part 1] how many times do digits 1, 4, 7, or 8 appear? ", part1num)
	fmt.Println("[Part 2] What do you get if you add up all of the output values? ", part2sum)

}


func decodeSixes(pattern string, one string, four string) int {
	digit9 := true
	for _, s := range four {
		if !strings.ContainsRune(pattern, s) {
			digit9 = false
		}
	}

	if digit9 {
		return 9
	} else if strings.Contains(pattern, string(one[0])) && strings.Contains(pattern, string(one[1])) {
		return 0
	} else {
		return 6
	}
}


func decode(pattern string, one string, six string) int {
	if strings.Contains(pattern, string(one[0])) && strings.Contains(pattern, string(one[1])) {
		return 3
	}

	digit5 := true
	for _, s := range pattern {
		if !strings.ContainsRune(six, s) {
			digit5 = false
		}
	}

	if digit5 {
		return 5
	} else {
		return 2
	}
}