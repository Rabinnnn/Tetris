package main

import (
	"fmt"
	"math"
	"tetris/utils"
)

func main() {
	tetrominoes := utils.Reader()
	output := utils.IsValid(tetrominoes)
	if output == "ERROR" {
		fmt.Println("ERROR")
		return
	}

	tetrominoes = utils.Trimmer(tetrominoes)
	size := int(math.Ceil(math.Sqrt(float64(len(tetrominoes) * 4))))

	var solvedBoard [][]string
	for {
		board := utils.CreateBoard(size)
		solvedBoard = utils.SolveBoard(board, tetrominoes)
		if solvedBoard != nil {
			break
		}
		size++
	}
	utils.PrintBoard(solvedBoard)
}
