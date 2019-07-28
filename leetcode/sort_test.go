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
