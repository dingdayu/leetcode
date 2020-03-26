package main

import (
	"fmt"
)

func main() {
	board := [][]byte{
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		[]byte{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		[]byte{'.', 'p', 'B', 'R', 'B', 'p', '.', '.'},
		[]byte{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		[]byte{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
	}
	fmt.Printf("num: %d\n", numRookCaptures(board))
}

func numRookCaptures(board [][]byte) int {
	cnt, st, ed := 0, 0, 0
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}

	for i, bor := range board {
		for j, v := range bor {
			if v == 'R' {
				st = i
				ed = j
			}
		}
	}

	for i := 0; i < 4; i++ {
		for step := 0; ; step++ {
			tx := st + step*dx[i]
			ty := ed + step*dy[i]
			if tx < 0 || tx >= 8 || ty < 0 || ty >= 8 || board[tx][ty] == 'B' {
				break
			}
			if board[tx][ty] == 'p' {
				cnt++
				break
			}
		}
	}
	return cnt
}
