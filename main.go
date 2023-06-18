package main

import "fmt"

func main() {
	fmt.Println("Hey Hash")
	emptyMapPrint()
	fmt.Println("")
	fmt.Println("ContainsDuplicateMap...")
	nums := []int{1, 2, 3, 1}
	fmt.Println(ContainsDuplicateMap(nums))

	fmt.Println("SingleNumber...")
	nums2 := []int{4, 1, 2, 1, 2}
	fmt.Println(SingleNumber(nums2))

	fmt.Println("SingleNumber...")
	nums3 := []int{1, 2, 2, 1}
	nums4 := []int{2, 2}
	fmt.Println(Intersection(nums3, nums4))

	fmt.Println("TwoSum...")
	nums5 := []int{2, 7, 11, 15}
	target5 := 9
	fmt.Println("TwoSum: ", TwoSumHashMap(nums5, target5))

	fmt.Println("IsIsomorphic...")
	str1 := "egg"
	str2 := "add"
	fmt.Println("IsIsomorphic: ", IsIsomorphic(str1, str2))

	fmt.Println("majorityElement...")
	nums6 := []int{3, 2, 3}
	fmt.Println("majorityElement: ", majorityElement(nums6))

	fmt.Println("missingNumber...")
	nums7 := []int{3, 0, 1}
	fmt.Println("missingNumber: ", missingNumber(nums7))

	fmt.Println("findSumOfThree...")
	nums8 := []int{-1, 2, 1, -4}
	target2 := 2
	fmt.Println("findSumOfThree: ", findSumOfThree(nums8, target2))

	fmt.Println("twoSum2...")
	nums9 := []int{2, 7, 11, 15}
	target9 := 9
	fmt.Println("twoSum2: ", twoSum2(nums9, target9))

	fmt.Println("IsHappy...")
	n := 19
	fmt.Println("IsHappy: ", IsHappy(n))

	fmt.Println("firstUniqChar...")
	str := "leetcode"
	fmt.Println("firstUniqChar: ", firstUniqChar(str))

	fmt.Println("Intersection2...")
	nums10 := []int{4, 9, 5}
	nums11 := []int{9, 4, 9, 8, 4}
	fmt.Println("Intersection2: ", Intersection2(nums10, nums11))
}
