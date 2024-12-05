package utils
import "fmt"

// PrintBoard displays the contents of the board in a grid-like format.
func PrintBoard(board [][]string){
	for _, row := range board{
		for _, cell := range row{
			fmt.Print(cell)
		}
		fmt.Println()
	}
	fmt.Println()
}