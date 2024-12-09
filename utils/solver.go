package utils

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