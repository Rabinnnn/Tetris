package utils

import(
	"fmt"
	"os"
	"strings"
)

// Reader is a function designed to extract the tetrominoeses from a given file.
// It does the following:
// - Ensures that other than the program name there is only 1 other argument passed 
//   by the user, which is the name of the file containing the tetrominoeses.
// - Processes the data by replacing the '#' character with a letter and '.' remains
//   the same.
// - Returns a 2D array.
func Reader()[][]string{
	var tetromino []string
	var tetrominoes [][]string
	var str strings.Builder
	letter := 'A'

	if len(os.Args) != 2{
		fmt.Println("ERROR")
		os.Exit(0)
	}

	fileOutput, err := os.ReadFile(os.Args[1])
	if err != nil{
		fmt.Println("ERROR")
		os.Exit(0)
	}

	fileOutputLines := strings.Split(string(fileOutput), "\n")


	for i, line := range fileOutputLines{
		if line != ""{
			for _, char := range line{
				if char == '#'{
					str.WriteRune(letter)
				}else if char == '.'{
					str.WriteRune(char)
				}else{
					fmt.Println("ERROR")
					os.Exit(0)
				}
			}
			tetromino = append(tetromino, str.String())
		}
		if line == "" || i == len(fileOutputLines)-1{
			tetrominoes = append(tetrominoes, tetromino)
			tetromino = []string{}
			letter++
		}
	}

return tetrominoes
}