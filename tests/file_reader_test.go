package utils

import (
	"reflect"
	"testing"
	"tetris/utils"
)

func TestReader(t *testing.T) {
	tests := []struct {
		name string
		want [][]string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.Reader(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reader() = %v, want %v", got, tt.want)
			}
		})
	}
}
