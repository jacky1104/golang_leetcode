package leetcode

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	result := make([][]int, 0)

	tempQueue := []*TreeNode{root}

	for len(tempQueue) > 0 {
		nextLevel := make([]*TreeNode, 0)
		levelValue := make([]int, 0)
		for _, v := range tempQueue {
			levelValue = append(levelValue, v.Val)
			if v.Left != nil {
				nextLevel = append(nextLevel, v.Left)
			}
			if v.Right != nil {
				nextLevel = append(nextLevel, v.Right)
			}

		}
		result = append(result, levelValue)
		if len(nextLevel) > 0 {
			tempQueue = nextLevel
		} else {
			break
		}

	}

	return result

}
