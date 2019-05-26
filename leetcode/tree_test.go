package leetcode

import (
	"strconv"
)

func inOrderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	result := make([]int, 0)
	if root.Left != nil {
		result = inOrderTraversal(root.Left)
	}

	result = append(result, root.Val)

	if root.Right != nil {
		result = append(result, inOrderTraversal(root.Right)...)
	}
	return result
}

func postOrderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	result := make([]int, 0)
	if root.Left != nil {
		result = postOrderTraversal(root.Left)
	}

	if root.Right != nil {
		result = append(result, postOrderTraversal(root.Right)...)
	}
	result = append(result, root.Val)

	return result
}

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}
	left, lban := GetMaxTreeDepthWhenBalance(root.Left)
	right, rban := GetMaxTreeDepthWhenBalance(root.Right)

	if !lban || !rban {
		return false
	}

	if abs(left-right) > 1 {
		return false
	}

	return true
}

func GetMaxTreeDepthWhenBalance(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	left, leftBan := GetMaxTreeDepthWhenBalance(root.Left)
	right, rightBan := GetMaxTreeDepthWhenBalance(root.Right)

	if !leftBan || !rightBan {
		return 0, false
	}

	if abs(left-right) > 1 {
		return 0, false
	}

	return 1 + maxValue(left, right), true
}

func abs(value int) int {
	if value >= 0 {
		return value
	}
	return value * -1
}

func maxValue(left, right int) int {
	if left-right > 0 {
		return left
	}
	return right
}

/**
111. Minimum Depth of Binary Tree
https://leetcode.com/problems/minimum-depth-of-binary-tree/
*/
func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if left == 0 {
		return 1 + right
	}
	if right == 0 {
		return 1 + left
	}
	if left > right {
		return right + 1
	}

	return left + 1

}

func hasPathSum(root *TreeNode, sum int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val == sum {
		return true
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)

}

/**
113. Path Sum II
https://leetcode.com/problems/path-sum-ii/
*/
func pathSum(root *TreeNode, sum int) [][]int {

	var res = make([][]int, 0)

	if root == nil {
		return res
	}

	if root.Left == nil && root.Right == nil && root.Val == sum {
		res = append(res, []int{root.Val})
		return res
	}

	left := pathSum(root.Left, sum-root.Val)
	right := pathSum(root.Right, sum-root.Val)

	if len(left) > 0 {
		for _, v := range left {
			res = append(res, append([]int{root.Val}, v...))
		}
	}

	if len(right) > 0 {
		for _, v := range right {
			res = append(res, append([]int{root.Val}, v...))
		}
	}

	return res

}

/**
257. Binary Tree Paths
https://leetcode.com/problems/binary-tree-paths/
*/

func binaryTreePaths(root *TreeNode) []string {

	if root == nil {
		return []string{}
	}

	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	left := binaryTreePaths(root.Left)
	for i := 0; i < len(left); i++ {
		left[i] = strconv.Itoa(root.Val) + "->" + left[i]
	}

	right := binaryTreePaths(root.Right)
	for i := 0; i < len(right); i++ {
		right[i] = strconv.Itoa(root.Val) + "->" + right[i]
	}

	return append(left, right...)

}

///**
//988. Smallest String Starting From Leaf
//https://leetcode.com/problems/smallest-string-starting-from-leaf/
// */
//func smallestFromLeaf(root *TreeNode) string {
//
//}

/**
100. Same Tree
https://leetcode.com/problems/same-tree/
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}

	if p != nil && q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) {
		return true
	}

	return false
}

/**
897. Increasing Order Search Tree
https://leetcode.com/problems/increasing-order-search-tree/
*/

func increasingBST(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return root
	}

	left := increasingBST(root.Left)
	right := increasingBST(root.Right)
	root.Right = right
	root.Left = nil

	if left == nil {
		return root
	}

	temp := left
	for temp != nil && temp.Right != nil {
		temp = temp.Right
	}
	temp.Right = root
	return left
}

/**
872. Leaf-Similar Trees
https://leetcode.com/problems/leaf-similar-trees/
*/

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	leafOne := getLeafFromTree(root1)
	leafTwo := getLeafFromTree(root2)

	if len(leafOne) != len(leafTwo) {
		return false
	}

	if len(leafOne) == 0 {
		return true
	}
	for i := 0; i < len(leafOne); i++ {
		if leafOne[i] != leafTwo[i] {
			return false
		}

	}
	return true

}

func getLeafFromTree(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)

	if root.Left == nil && root.Right == nil {
		res = append(res, root.Val)
		return res
	}

	res = append(res, getLeafFromTree(root.Left)...)
	res = append(res, getLeafFromTree(root.Right)...)

	return res

}

/**
98. Validate Binary Search Tree
https://leetcode.com/problems/validate-binary-search-tree/
*/
func isValidBST(root *TreeNode) bool {

	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}

	if root.Left != nil {
		if root.Val <= root.Left.Val {
			return false
		}

		leftRight := root.Left
		for leftRight != nil && leftRight.Right != nil {
			leftRight = leftRight.Right
		}
		if leftRight.Val >= root.Val {
			return false
		}
	}

	if root.Right != nil {
		if root.Val >= root.Right.Val {
			return false
		}

		rightLeft := root.Right
		for rightLeft != nil && rightLeft.Left != nil {
			rightLeft = rightLeft.Left
		}
		if rightLeft.Val <= root.Val {
			return false
		}
	}

	return isValidBST(root.Left) && isValidBST(root.Right)

}
