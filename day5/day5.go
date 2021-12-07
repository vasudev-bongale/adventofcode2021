package main

import (
	"adventofcode2021/common"
	"fmt"
	"regexp"
	"strconv"
)

func main() {

	data := common.GetInputDataFromFile("../common/day5_data.txt")

	vents := make(map[[2]int]int)

	for _, line := range data {
		re := regexp.MustCompile(`(\d+),(\d+)`)
		values := re.FindAllStringSubmatch(line, 2)

		x1, _ := strconv.Atoi(values[0][1])
		y1, _ := strconv.Atoi(values[0][2])
		x2, _ := strconv.Atoi(values[1][1])
		y2, _ := strconv.Atoi(values[1][2])

		// Sort the order of points
		if x1 > x2 {
			temp := x2
			x2 = x1
			x1 = temp
			temp = y2
			y2 = y1
			y1 = temp
		} else if x1 == x2 {
			// swap the ys
			if y1 > y2 {
				temp := y2
				y2 = y1
				y1 = temp
			}
		}

		if y1 == y2  {
			for i := x1; i <= x2; i++ {
					vents[[2]int{i, y1}] += 1
			}
		} else if x1 == x2 {
			for i := y1; i <= y2; i++ {
				vents[[2]int{x1, i}] += 1
			}
		}

		slope := 0
		if x2 - x1 != 0 {
			slope = (y2 - y1)/(x2 - x1)
		}
		
		if slope == 1 {
			for i := 0; i <= (x2 - x1); i++ {
				vents[[2]int{x1+i, y1+i}] += 1
			}
		}

		if slope == -1 {
			for i := 0; i <= (x2 - x1); i++ {
				vents[[2]int{x1+i, y1-i}] += 1
			}
		}
	}
	fmt.Println("Part 2 Answer: ", overlaps(vents))

}

func overlaps(vents map[[2]int]int) int {
	dangerousVents := 0
	for _, v := range vents {
		if v > 1 {
			dangerousVents++
		}
	}
	return dangerousVents
}

func printUtil(vents map[[2]int]int) {

	for j := 0; j <=9; j++ {
		row := ""
		for i:=0; i<=9; i++ {
			v := vents[[2]int{i, j}]
			if v == 0 {
				row += "."
			} else {
				row += strconv.Itoa(v)
			}
		}
		fmt.Println(row)
	}
}