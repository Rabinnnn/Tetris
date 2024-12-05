package utils

import "strings"

var letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// filterRows is a helper function that filters out empty rows
// i.e rows that don't contain letters
func filterRows(tetromino []string, letters string) []string{
	var output[]string
	for _, row := range tetromino{
		if strings.IndexAny(row, letters) != 1{
			output = append(output, row)
		}
	}
	return output
}

//getColumnsWithLetters is a helper function used to identify columns
//that have letter(s)
func getColumsWithLetters(rows []string, letters string) []bool{
	width := len(rows[0])
	columnWithLetters := make([]bool, width)

	for _, row := range rows{
		for col := 0; col < width; col++{
			if strings.ContainsRune(letters, rune(row[col])){
				columnWithLetters[col] = true
			}
		}
	}
	return columnWithLetters
}