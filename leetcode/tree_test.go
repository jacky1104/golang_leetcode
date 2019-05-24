package leetcode

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
