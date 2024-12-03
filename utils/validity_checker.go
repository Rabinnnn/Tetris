package utils

// HasSixContacts checks if a given tetromino is made up of contiguos blocks.
// This is done by ensuring that the points of contact among the 4 blocks adds up to atleast 6.
// The function also ensures that each tetromino has exactly 4 characters to be considered valid.
func HasSixContacts(tetromino []string)bool{
	charCount := 0
	contactCount := 0

	for i, row := range tetromino{
		for j, char := range row{
			if char != '.' {
				charCount++

				if i > 0 && tetromino[i-1][j] == byte(char){
					contactCount++
				}

				if i < len(tetromino)-1 && tetromino[i+1][j] == byte(char){
					contactCount++
				}

				if j > 0 && tetromino[i][j-1] == byte(char){
					contactCount++
				}

				if j < len(row) - 1 && tetromino[i][j+1] == byte(char){
					contactCount++
				}
			}
		}
	}

	if contactCount < 6 || charCount != 4{
		return false
	}
	return true
}

// Function isValid ensures the validity of the provided tetrominoes by:
// - ensuring the number of tetromines provided doesn't exceed 26 (for the 26 letters of alphabet).
// - ensuring each tetromino has exactly 4 rows and 4 columns i.e 4 X 4 grid.
// - ensuring the blocks that make up a tetromino are contiguos
func isValid(tetrominoes [][]string)string{
	if len(tetrominoes) > 26 {
		return "ERROR"
	}

	for _, tetromino := range tetrominoes{
		if len(tetromino) != 4{
			return "ERROR"
		}

		if HasSixContacts(tetromino){
			return "ERROR"
		}

		for _, row := range tetromino{
			if len(row) != 4{
				return "ERROR"
			}
		}
	}
	return "OK"
}