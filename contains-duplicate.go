package main

import "fmt"

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
