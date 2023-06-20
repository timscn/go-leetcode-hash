package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	fmt.Println("jewels: ", jewels, " and stones: ", stones)
	countJewels := len(jewels)
	countStones := len(stones)
	count := 0
	for i := 0; i < countStones; i++ {
		fmt.Println("i: ", i, " and value is: ", string(stones[i]))
		for j := 0; j < countJewels; j++ {
			fmt.Println("j: ", j, " and value is: ", string(jewels[j]))
			if jewels[j] == stones[i] {
				count++
			}
		}
	}
	return count
}

func numJewelsInStonesHM(jewels string, stones string) int {
	fmt.Println("jewels: ", jewels, " and stones: ", stones)
	hmJ := make(map[rune]int)
	for _, value := range jewels {
		hmJ[value] += 1
	}
	count := 0
	for _, stone := range stones {
		if _, ok := hmJ[stone]; ok {
			count += 1
		}
	}
	return count
}
