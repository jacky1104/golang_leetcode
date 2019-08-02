package leetcode

import (
	"sort"
)

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

/**
49. Group Anagrams
https://leetcode.com/problems/group-anagrams/
*/
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 || len(strs) == 1 {
		return [][]string{strs}
	}
	mark := make(map[string][]string)
	for _, str := range strs {
		k := []byte(str)
		sort.Slice(k, func(i int, j int) bool { return k[i] < k[j] })
		if v, ok := mark[string(k)]; ok {
			mark[string(k)] = append(v, str)
		} else {
			mark[string(k)] = []string{str}
		}
	}

	result := make([][]string, 0)

	for _, v := range mark {
		result = append(result, v)
	}

	return result
}

/**
56. Merge Intervals
https://leetcode.com/problems/merge-intervals/
*/
func merge(intervals [][]int) [][]int {

	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//merge
	i := 0
	for i < len(intervals)-1 {
		if intervals[i][1] >= intervals[i+1][0] {
			if intervals[i][1] < intervals[i+1][1] {
				intervals[i][1] = intervals[i+1][1]
			}
			//delete intervals[i+1]
			if i+2 >= len(intervals) {
				intervals = intervals[:i+1]
			} else {
				intervals = append(intervals[:i+1], intervals[i+2:]...)
			}
		} else {
			i++
		}
	}
	return intervals
}

/**
75. Sort Colors
https://leetcode.com/problems/sort-colors/
*/
func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}

	zero, two := 0, len(nums)-1

	for i := zero; i <= two; i++ {
		if nums[i] == 2 {
			nums[i], nums[two] = nums[two], nums[i]
			two--
			i--
		} else if nums[i] == 0 {
			nums[i], nums[zero] = nums[zero], nums[i]
			zero++
		}
	}

}
