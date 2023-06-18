package main

import "fmt"

func missingNumber(nums []int) int {
	fmt.Println("nums: ", nums)
	hm := make(map[int]int)
	for _, element := range nums {
		hm[element] = 1
	}
	for i := 0; i <= len(nums); i++ {
		if _, ok := hm[i]; !ok {
			return i
		}
	}
	return -1
}
