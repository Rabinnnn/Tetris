package utils

import "testing"


func TestHasSixContacts(t *testing.T) {
	type args struct {
		tetromino []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "6Contacts", args: args{tetromino: []string{"...#", "...#", "...#", "...#"}}, want:true},
		{name: "5Contacts", args: args{tetromino: []string{"...#", "...#", "...#", "..#."}}, want:false},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasSixContacts(tt.args.tetromino); got != tt.want {
				t.Errorf("HasSixContacts() = %v, want %v", got, tt.want)
			}
		})
	}
}
