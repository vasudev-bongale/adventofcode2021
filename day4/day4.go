package main

import (
	"adventofcode2021/common"
	"fmt"
	"strings"
	"strconv"
)

type Board struct {
	table       [][]int
	bingoTable  [][]int
	won         bool
}

func main() {
	draws, boards := bingoData()

	winners := 0
	wins := []int{}
	scores := []int{}

	gameFinished:
	for _, draw := range draws {
		// Update playTable of each board
		for i, b := range boards {
			b.recordDraw(draw)
			if b.isWinner() {
				var alreadyWon bool
				for _, winner := range wins {
					if i == winner {
						alreadyWon = true
					}
				}
				if !alreadyWon {
					winners += 1
					wins = append(wins, i)
					scores = append(scores, b.calculateScore(draw))
				}

				if winners == len(boards) {
					break gameFinished
				}
			}
		}
	}
	fmt.Printf("First Winner: Board: %d with score: %d\n", wins[0], scores[0])
	fmt.Printf("Last Winner: Board: %d with score: %d\n", wins[len(wins)-1], scores[len(scores)-1])

}

// Given the current draw value, update the current play table
func (b *Board) recordDraw(draw int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
		   if b.table[i][j] == draw {
			   b.bingoTable[i][j] = 1
		   }
		}
	 }
}

func (b *Board) calculateScore(draw int) int {
	score := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.bingoTable[i][j] == 0 {
				score += b.table[i][j]
			}
		}
	}
	return score * draw
}

func (b *Board) isWinner() (winner bool) {	
	
	// check if a col is full
	for col := 0; col < 5; col++ {
		var sum int
		for row := 0; row < 5; row++ {
			sum += b.bingoTable[row][col]
		}
		if sum == 5 {
			winner = true
		}
	}

	// check if a row is full
	for row := 0; row < 5; row++ {
		var sum int
		for col := 0; col < 5; col++ {
			sum += b.bingoTable[row][col]
		}
		if sum == 5 {
			winner = true
		}
	}
	return
}

func createBoard(boardStr []string) Board {
	b := Board{}
	board := make([][]int, 5)
	for i := range board {
		board[i] = make([]int, 5)
	}

	for row, line := range boardStr {
		nums := strings.Fields(line)
		for col, num := range nums {
			n, _ := strconv.Atoi(num)
			board[row][col] = n
		}
	}

	b.table = board
	b.bingoTable = make([][]int, 5)
	for i := range b.bingoTable {
		b.bingoTable[i] = make([]int, 5)
	}
	return b
}

func bingoData() ([]int, []Board) {
	input := common.GetInputDataFromFile("../common/day4_data.txt")
	draws := []int{}

	for _, value := range strings.Split(input[0], ",") {
		draw, _ := strconv.Atoi(value)
		draws = append(draws, draw)
	}

	boards := []Board{}

	l := 2
	for l < len(input) {
		board := createBoard(input[l:l+5])
		boards = append(boards, board)
		l = l + 6
	}
	return draws, boards
} 