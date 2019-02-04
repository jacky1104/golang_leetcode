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
		result[i] = max(result[i-2]+nums[i], result[i-1])
	}

	return result[length-1]
}

//House Robber 2
//https://leetcode.com/problems/house-robber-ii/
func rob2(nums []int) int {
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

func TestYushu(t *testing.T) {
	count := 0
	for i := 0; i < 1000000; i++ {
		if i%500 > 400 && i%500 < 470 {
			count ++
		}
	}
	fmt.Println("Count:", count)
	fmt.Println("Count:", float64(count)/1000000)

}

//https://leetcode.com/problems/unique-paths-ii/
func TestUniquePath(t *testing.T) {

	obstacleGrid := [][]int {{0,0,0},{0,1,0},{0,0,0}}

	fmt.Println(uniquePathsWithObstacles(obstacleGrid))

	obstacleGrid = [][]int{{0,0}}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))

	obstacleGrid = [][]int{{0}}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))

	obstacleGrid = [][]int{{1}}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))

}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	//init the matrix
	mark := make([][]int,m)
	for i:= 0; i<m ; i++  {
		mark[i] = make([]int,n)
	}

	if obstacleGrid[m-1][n-1] == 1{
		return 0
	}

	// the last column
	for i := m - 1; i >= 0; i-- {
		if obstacleGrid[i][n-1] == 1 {
			mark[i][n-1] = 0
			break
		} else {
			mark[i][n-1] = 1
		}
	}

	//the last row
	for i := n - 2; i >= 0; i-- {
		if obstacleGrid[m-1][i] == 1 {
			mark[m-1][i] = 0
			break
		} else {
			mark[m-1][i] = 1
		}
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				mark[i][j] = 0
			}else{
				mark[i][j] = mark[i+1][j] + mark[i][j+1]
			}
		}
	}

	return mark[0][0]

}
