package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	numsLen := len(nums)
	sort.Ints(nums)
	if numsLen%2 == 0 {
		i := numsLen / 2
		return (float64(nums[i-1]) + float64(nums[i])) / 2
	} else {
		return float64(nums[(numsLen-1)/2])
	}
}

func main() {
	a := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Println(a)
}
