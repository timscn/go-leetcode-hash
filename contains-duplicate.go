package main

import (
	"fmt"
)

func ContainsDuplicateO2(nums []int) bool {
	fmt.Println("nums: ", nums)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func containsDuplicate3(nums []int) bool {
	hm := make(map[int]int)
	for _, value := range nums {
		hm[value] += 1
		if hm[value] > 1 {
			return true
		}
	}
	return false
}

func ContainsDuplicateMap(nums []int) bool {
	fmt.Println("nums: ", nums)
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = 1
	}
	return false
}

func containsNearbyDuplicate3(nums []int, k int) bool {
	hm := make(map[int]int)
	for ind, val := range nums {
		if v, ok := hm[val]; ok {
			if abs(ind, v) <= k {
				return true
			}
		}
		hm[val] = ind
	}
	return false
}

func abs(x, y int) int {
	result := x - y
	if result < 0 {
		return -result
	}
	return result
}
