package main

import (
	"fmt"
)

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(Find(arr, 0, len(arr)-1, 6))
}

func Find(arr []int, leftIndex int, rightIndex int, val int) int {
	if leftIndex > rightIndex {
		return -1
	}
	middle := (leftIndex + rightIndex) / 2
	if arr[middle] > val {
		return Find(arr, leftIndex, middle-1, val)
	} else if arr[middle] < val {
		return Find(arr, middle+1, rightIndex, val)
	} else {
		return middle
	}
}
