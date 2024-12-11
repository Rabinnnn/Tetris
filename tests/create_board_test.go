package utils

import (
	"reflect"
	"testing"
	"tetris/utils"
)

// Test CreateBoard function with 2 different input cases i.e size = 2 and size = 3
func TestCreateBoard(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{name: "Two", args: args{size: 2}, want: [][]string{{".", "."}, {".", "."}}},
		{name: "Three", args: args{size: 3}, want: [][]string{{".", ".", "."}, {".", ".", "."}, {".", ".", "."}}},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.CreateBoard(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
