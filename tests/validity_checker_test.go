package utils

import (
	"testing"
	"tetris/utils"
)

func TestHasSixContacts(t *testing.T) {
	type args struct {
		tetromino []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "6Contacts", args: args{tetromino: []string{"...#", "...#", "...#", "...#"}}, want: true},
		{name: "5Contacts", args: args{tetromino: []string{"...#", "...#", "...#", "..#."}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.HasSixContacts(tt.args.tetromino); got != tt.want {
				t.Errorf("HasSixContacts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	type args struct {
		tetrominoes [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "case1", args: args{tetrominoes: [][]string{{"...A", "...A", "...A", "...A"}, {"....", "....", "....", "BBBB"}}}, want: "OK"},
		{name: "case2", args: args{tetrominoes: [][]string{{"...A", "..A.", "...A", "...A"}, {"....", "....", "....", "BBBB"}}}, want: "ERROR"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.IsValid(tt.args.tetrominoes); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
