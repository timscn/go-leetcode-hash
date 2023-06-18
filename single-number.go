package main

import "fmt"

func SingleNumber(nums []int) int {
	fmt.Println("SingleNumber: ", nums)
	result := -1
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}
	for key, valueMap := range m {
		if valueMap == 1 {
			result = key
		}
	}
	return result
}
