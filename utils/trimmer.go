package utils

import "strings"

//Trimmer is a function that removes the rows and columns that do not contain 
//uppercase letters (A-Z) thus effectively trimming the empty rows and columns 
//from the shapes.
func Trimmer(tetrominoes [][]string) [][]string{
	var letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var trimmedTetrominoes [][]string

	for _, tetromino := range tetrominoes{
		filteredRows := FilterRows(tetromino, letters)
		if len(filteredRows) == 0{
			continue
		}

		columnWithLetters := GetColumsWithLetters(filteredRows, letters)
		trimmedTetrominoes = append(trimmedTetrominoes, TrimRowsByColumns(filteredRows,columnWithLetters))

	}
	return trimmedTetrominoes
}

// filterRows is a helper function that filters out empty rows
// i.e rows that don't contain letters
func FilterRows(tetromino []string, letters string) []string{
	var output[]string
	for _, row := range tetromino{
		if strings.IndexAny(row, letters) != -1{
			output = append(output, row)
		}
	}
	return output
}

//getColumnsWithLetters is a helper function used to identify columns
//that have letter(s)
func GetColumsWithLetters(rows []string, letters string) []bool{
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

//trimRowsByColumns is a helper function to trim the rows based on valid columns
//i.e columns that have letter(s)
func TrimRowsByColumns(rows []string, columnWithLetters []bool) []string{
	var trimmedRows []string
	for _, row := range rows{
		var newRow strings.Builder
		for col := 0; col < len(row); col++{
			if columnWithLetters[col]{
				newRow.WriteByte(row[col])
			}
		}
		trimmedRows = append(trimmedRows, newRow.String())
	}
	return trimmedRows
}