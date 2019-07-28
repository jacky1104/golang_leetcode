package leetcode

/**
242. Valid Anagram
https://leetcode.com/problems/valid-anagram/
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if len(s) == 0 {
		return true
	}

	sMap := make(map[byte]int)
	input := []byte(s)
	for i := 0; i < len(s); i++ {
		if v, ok := sMap[input[i]]; ok {
			v += 1
			sMap[input[i]] = v
		} else {
			sMap[input[i]] = 1
		}
	}

	outPut := []byte(t)
	for i := 0; i < len(s); i++ {
		if v, ok := sMap[outPut[i]]; ok {
			v -= 1
			if v < 0 {
				return false
			}
			sMap[outPut[i]] = v
		} else {
			return false
		}
	}

	return true

}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if len(s) == 0 {
		return true
	}

	array := [26]byte{}
	input := []byte(s)
	for i := 0; i < len(s); i++ {
		array[input[i]-97]++
	}

	outPut := []byte(t)
	for i := 0; i < len(s); i++ {
		array[outPut[i]-97]--
	}

	for i := 0; i < len(array); i++ {
		if array[i] != 0 {
			return false
		}
	}

	return true

}
