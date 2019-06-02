package leetcode

/**
26. Remove Duplicates from Sorted Array
https://leetcode.com/problems/remove-duplicates-from-sorted-array/

This method will change the input nums slice
*/

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}
	curIndex := 0
	last := nums[len(nums)-1]

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[curIndex] {
			curIndex++
			nums[curIndex] = nums[i]
		}

		if nums[curIndex] == last {
			nums = nums[:curIndex+1]
			break
		}
	}

	return len(nums)
}

func removeDuplicates2(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	cur := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != cur {
			count++
			cur = nums[i]
		}
	}

	return count
}

/**
80. Remove Duplicates from Sorted Array II
https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
*/
func removeDuplicates3(nums []int) int {

	if nums == nil || len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 || len(nums) == 2 {
		return len(nums)
	}
	curIndex := 0
	duplicateCount := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[curIndex] {
			curIndex++
			nums[curIndex] = nums[i]
			duplicateCount = 1
		} else {
			duplicateCount++
		}

		if duplicateCount == 2 {
			curIndex++
			nums[curIndex] = nums[i]
		}
	}
	nums = nums[:curIndex+1]

	return len(nums)
}
