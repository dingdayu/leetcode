package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return nil
	}
	var ret [][]int
	for k, v := range nums {
		if v > 0 {
			break
		}
		if k > 0 && v == nums[k-1] {
			continue
		}
		target := 0 - v
		i := k + 1
		j := len(nums) - 1
		for i < j {
			if nums[i]+nums[j] == target {
				ret = append(ret, []int{v, nums[i], nums[j]})
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				i++
				j--
			} else if nums[i]+nums[j] < target {
				i++
			} else {
				j--
			}
		}
	}
	return ret
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	nu := threeSum(nums)
	fmt.Println(nu)
}
