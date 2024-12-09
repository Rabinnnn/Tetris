package utils

// SolveBoard returns the solved board if it is solveable otherwise it returns an empty 2D slice
func SolveBoard(board [][]string, tetrominoes [][]string) [][]string{
	if solved(board, tetrominoes, 0){
		return board
	}
	return [][]string{}
}

// isPlaceable checks if a give tetromino can be placed in certain coordinates within the board.
// It iterates through the characters of the tetromino to determine if:
// -the character is not a dot(representing empty space)
// - it is not in a position that overlaps the boundaries of the board.
// If the tetromino can be place it returns 'true' otherwise it returns 'false'
func isPlaceable(board [][]string, tetromino []string, x, y int)bool{
	for i, row := range tetromino{
		for j, char := range row{
			if char != '.'{
				if x + j >= len(board[0]) || y + i >= len(board) || board[y+i][x+j] != "."{
					return false
				}
			}
		}
	}
	return true
}

//placeTetromino inserts a tetromino at a certain position on the board using the coordinates provided.
func placeTetromino(board [][]string, tetromino []string, x, y int){
	for i, row := range tetromino{
		for j, char := range row{
			if char != '.'{
				board[x+j][y+i] = string(char)
			}
		}
	}
}

//removeTetromino takes a tetromino at the given coordinates off the board.
func removeTetromino(board [][]string, tetromino []string, x, y int){
	for i, row := range tetromino{
		for j, char := range row{
			if char != '.'{
				board[x+j][y+i] = "."
			}
		}
	}
}

//solved is a function that tries to solve the puzzle through recurssion.
//It checks if a tetromino can be placed in a certain position first. If it's
//possible it places it there, otherwise it tries the next position. After placing
//it the function calls itself with the index incremented(to get the next tetromino).
//If the state of the board is !solved the tetromino is removed from the position.
func solved(board [][]string, tetrominoes [][]string, index int)bool{
	tetromino := tetrominoes[index]

	if index == len(tetrominoes){
		return true
	}

	for i := range board{
		for j := range board[i]{
			if isPlaceable(board, tetromino,i,j){
				placeTetromino(board,tetromino,i,j)
				if solved(board, tetrominoes, index+1){
					return true
				}

				removeTetromino(board,tetromino, i, j)
			}
		}
	}
return false
}