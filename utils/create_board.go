package utils

// CreateBoard generated a 2D slice representing a board.
// Each cell is initialized with the string "."
func CreateBoard(size int) [][]string {
	board := make([][]string, size)
	for i := range board {
		board[i] = make([]string, size)
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	return board
}
