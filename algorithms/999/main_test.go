package main

import (
	"testing"
)

func Test_numRookCaptures(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{board: [][]byte{
			[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', 'R', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRookCaptures(tt.args.board); got != tt.want {
				t.Errorf("numRookCaptures() = %v, want %v", got, tt.want)
			}
		})
	}
}
