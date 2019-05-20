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
