package main

import "fmt"

func majorityElement(nums []int) int {
	fmt.Println("nums: ", nums)
	hm := make(map[int]int)
	lenOfNums := len(nums)
	majFormula := lenOfNums / 2
	fmt.Println("majFormula: ", majFormula)
	for _, value := range nums {
		fmt.Println("value: ", value)
		hm[value] += 1
		if hm[value] > majFormula {
			return value
		}
		fmt.Println("hm[value] :", hm[value])
	}
	return -1
}
