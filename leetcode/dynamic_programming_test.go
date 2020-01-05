package leetcode

import (
	"fmt"
	"testing"
)

/**
https://leetcode.com/problems/maximum-subarray/submissions/
53. Maximum Subarray
*/
func maxSubArray(nums []int) int {

	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	maxSoFar, maxEndingHere := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxEndingHere = max(maxEndingHere+nums[i], nums[i])
		maxSoFar = max(maxSoFar, maxEndingHere)

	}

	return maxSoFar
}

/**
https://leetcode.com/problems/unique-paths/
62. Unique Paths
*/
func uniquePaths(m int, n int) int {

	if m == 1 || n == 1 {
		return 1
	}
	record := make([][]int, m)

	for i := 0; i < m; i++ {
		record[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		record[i][n-1] = 1
	}

	for i := 0; i < n; i++ {
		record[m-1][i] = 1
	}

	for j := m - 2; j >= 0; j-- {
		for i := n - 2; i >= 0; i-- {
			record[j][i] = record[j+1][i] + record[j][i+1]
		}
	}

	return record[0][0]
}

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
			count++
		}
	}
	fmt.Println("Count:", count)
	fmt.Println("Count:", float64(count)/1000000)

}

//https://leetcode.com/problems/unique-paths-ii/
func TestUniquePath(t *testing.T) {

	obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}

	fmt.Println(uniquePathsWithObstacles(obstacleGrid))

	obstacleGrid = [][]int{{0, 0}}
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
	mark := make([][]int, m)
	for i := 0; i < m; i++ {
		mark[i] = make([]int, n)
	}

	if obstacleGrid[m-1][n-1] == 1 {
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
			} else {
				mark[i][j] = mark[i+1][j] + mark[i][j+1]
			}
		}
	}

	return mark[0][0]

}

//https://leetcode.com/problems/minimum-path-sum/

func TestMinPathSum(t *testing.T) {

	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}

	fmt.Println(minPathSum(grid))

	grid = [][]int{{0, 1}}
	fmt.Println(minPathSum(grid))

	grid = [][]int{{0}}
	fmt.Println(minPathSum(grid))

	grid = [][]int{{1, 0}}
	fmt.Println(minPathSum(grid))

	grid = [][]int{{1}, {1}}
	fmt.Println(minPathSum(grid))

}

func minPathSum(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])
	//init the matrix
	mark := make([][]int, m)
	for i := 0; i < m; i++ {
		mark[i] = make([]int, n)
	}
	mark[m-1][n-1] = grid[m-1][n-1]
	// the last column
	for i := m - 2; i >= 0; i-- {
		mark[i][n-1] = mark[i+1][n-1] + grid[i][n-1]
	}
	//the last row
	for i := n - 2; i >= 0; i-- {
		mark[m-1][i] = mark[m-1][i+1] + grid[m-1][i]
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			if mark[i][j+1] > mark[i+1][j] {
				mark[i][j] = grid[i][j] + mark[i+1][j]
			} else {
				mark[i][j] = grid[i][j] + mark[i][j+1]
			}
		}
	}

	return mark[0][0]

}

//https://leetcode.com/problems/unique-binary-search-trees/
//https://leetcode.com/problems/unique-binary-search-trees-ii/

func TestUniqueBinarySearchTree(t *testing.T) {

	fmt.Println(generateTreesOne(1))
	fmt.Println(generateTreesOne(2))
	fmt.Println(generateTreesOne(3))
	fmt.Println(generateTreesOne(4))
	fmt.Println(generateTreesOne(5))
}

func generateTreesOne(n int) int {

	if n == 1 || n == 2 {
		return n
	}

	if n == 3 {
		return 5
	}
	count := make([]int, n+1)
	count[0] = 1
	count[1] = 1
	count[2] = 2
	count[3] = 5

	for i := 4; i <= n; i++ {
		//root node value
		for j := 1; j <= i; j++ {
			//left tree * right tree
			count[i] += count[j-1] * count[i-j]
		}
	}
	return count[n]

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
