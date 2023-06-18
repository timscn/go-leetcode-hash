package main

import (
	"fmt"
	"sort"
)

func findSumOfThree(nums []int, target int) bool {
	fmt.Println("nums: ", nums)
	sort.Ints(nums)
	low, high, TotalSum := 0, 0, 0
	lenOfArray := len(nums)

	for i := 0; i < lenOfArray-2; i++ {
		low = i + 1
		high = lenOfArray - 1
		for low <= high {
			TotalSum = nums[low] + nums[high] + nums[i]
			if TotalSum == target {
				return true
			} else if TotalSum < target {
				low += 1
			} else {
				high -= 1
			}
		}
	}

	return false
}
