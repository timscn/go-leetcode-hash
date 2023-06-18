package main

import "fmt"

// O(n2) Complexity
func TwoSum(nums []int, target int) []int {
	fmt.Println("nums: ", nums, " and target is: ", target)
	for i, left := range nums {
		for j, right := range nums {
			if left+right == target && i != j {
				return []int{i, j}
			}
		}
	}
	return nil
}

// O(n)
func TwoSumHashMap(nums []int, target int) []int {
	fmt.Println("nums: ", nums, " and target is: ", target)
	hm := make(map[int]int)
	for i, num := range nums {
		if _, ok := hm[num]; ok {
			return []int{i, hm[num]}
		}
		hm[target-num] = i
	}
	return nil
}

func twoSum2(numbers []int, target int) []int {
	fmt.Println("nums: ", numbers, " and target is: ", target)
	left, right, sum := 0, len(numbers)-1, 0
	for left < right {
		sum = numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else {
			if sum > target {
				right--
			}
		}
	}
	return nil
}
