package leetcode

import (
	"fmt"
	"testing"
)

func TestHouseRobber(t *testing.T) {

	fmt.Println("House Robber")
	//rob([]int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240})
	rob1([]int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240})

}

//recursive
func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	if length == 2 {
		return max(nums[0], nums[1])
	}

	return max(rob(nums[:length-2])+nums[length-1], rob(nums[:length-1]))
}

//bottom up
func rob1(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	if length == 2 {
		return max(nums[0], nums[1])
	}

	result := make([]int, length)
	result[0] = nums[0]
	result[1] = max(nums[0], nums[1])
	for i := 2; i < length; i++ {
		result[i] = max(result[i-2] + nums[i], result[i-1])
	}

	return  result[length-1]
}


//House Robber 2
//https://leetcode.com/problems/house-robber-ii/
func rob2(nums []int) int{
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	if length == 2 {
		return max(nums[0], nums[1])
	}

	return max(rob1(nums[:length-1]), rob1(nums[1:]))

}


func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
