package main

import "fmt"

func Intersection(nums1 []int, nums2 []int) []int {
	fmt.Println("nums1: ", nums1, "and nums2: ", nums2)
	m1 := make(map[int]bool)
	for _, v := range nums1 {
		m1[v] = true
	}
	returnArray := []int{}
	for _, v := range nums2 {
		if m1[v] == true {
			returnArray = append(returnArray, v)
		}
	}
	return returnArray
}

func Intersection2(nums1 []int, nums2 []int) []int {
	fmt.Println("nums1: ", nums1, " and nums2: ", nums2)
	hm := make(map[int]int)

	for _, value := range nums1 {
		hm[value]++
	}

	returnArray := []int{}
	for _, value := range nums2 {
		if _, ok := hm[value]; ok {
			fmt.Println("hm[value]: ", hm)
			if hm[value] == 0 {
				continue
			}
			hm[value]--
			returnArray = append(returnArray, value)
		}
	}
	return returnArray
}
