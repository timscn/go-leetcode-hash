package main

import "fmt"

func IsIsomorphic(s string, t string) bool {
	fmt.Println("s: ", s, "and t: ", t)
	hmS, hmT := make(map[uint8]int), make(map[uint8]int)
	for element := range s {
		if hmS[s[element]] != hmT[t[element]] {
			return false
		} else {
			hmS[s[element]] = element + 1
			hmT[t[element]] = element + 1
		}
	}
	return true
}
