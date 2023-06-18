package main

func firstUniqChar(s string) int {

	hm := make(map[byte]int)
	lenOfs := len(s)

	for i := 0; i < lenOfs; i++ {
		hm[s[i]]++
	}

	for i := 0; i < lenOfs; i++ {
		if hm[s[i]] == 1 {
			return i
		}
	}
	return -1
}
