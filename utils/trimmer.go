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