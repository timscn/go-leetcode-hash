package main

import (
	"fmt"
)

func IsHappy(n int) bool {
	fmt.Println("n: ", n)
	hm := make(map[int]bool)
	for n != 1 {
		if _, ok := hm[n]; ok {
			return false
		}
		hm[n] = true
		newNum := 0
		for n > 0 {
			v := n % 10
			newNum += v * v
			n = n / 10
		}
		n = newNum
	}
	return true
}
