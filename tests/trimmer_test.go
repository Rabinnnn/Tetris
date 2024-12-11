package utils

import (
	"reflect"
	"testing"
	"tetris/utils"
)

func TestFilterRows(t *testing.T) {
	type args struct {
		tetromino []string
		letters   string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "case1", args: args{tetromino: []string{"...A", "....", "...A", "...A"}, letters: "ABCD"}, want: []string{"...A", "...A", "...A"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.FilterRows(tt.args.tetromino, tt.args.letters); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterRows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetColumsWithLetters(t *testing.T) {
	type args struct {
		rows    []string
		letters string
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{name: "case1", args: args{rows: []string{"...A", "...A", "...A", "...A"}, letters: "ABCD"}, want: []bool{false, false, false, true}},
		{name: "case2", args: args{rows: []string{"B..B", "...B", "...B", "...B"}, letters: "ABCD"}, want: []bool{true, false, false, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetColumsWithLetters(tt.args.rows, tt.args.letters); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetColumsWithLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimRowsByColumns(t *testing.T) {
	type args struct {
		rows              []string
		columnWithLetters []bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "case1", args: args{rows: []string{"...A", "...A", "...A", "...A"}, columnWithLetters: []bool{false, false, false, true}}, want: []string{"A", "A", "A", "A"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.TrimRowsByColumns(tt.args.rows, tt.args.columnWithLetters); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrimRowsByColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimmer(t *testing.T) {
	type args struct {
		tetrominoes [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{name: "case1", args: args{tetrominoes: [][]string{{"...A", "...A", "...A", "...A"}, {"....", "....", "....", "BBBB"}}}, want: [][]string{{"A", "A", "A", "A"}, {"BBBB"}}},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.Trimmer(tt.args.tetrominoes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trimmer() = %v, want %v", got, tt.want)
			}
		})
	}
}
