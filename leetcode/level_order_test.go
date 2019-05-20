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

func zigzagLevelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	var result = make([][]int, 0)
	level := 1

	parent := make([]*TreeNode, 0)
	parent = append(parent, root)

	for len(parent) > 0 {
		sub := make([]*TreeNode, 0)
		levelResult := make([]int, 0)
		for _, v := range parent {
			if level%2 == 0 {
				levelResult = append([]int{v.Val}, levelResult...)
			} else {
				levelResult = append(levelResult, v.Val)
			}
			if v.Left != nil {
				sub = append(sub, v.Left)
			}
			if v.Right != nil {
				sub = append(sub, v.Right)
			}
		}
		level++
		result = append(result, levelResult)
		if len(sub) > 0 {
			parent = sub
		} else {
			break
		}

	}

	return result
}

func levelOrderBottom(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	result := make([][]int, 0)
	currentLevel := make([]*TreeNode, 1)
	currentLevel[0] = root

	for len(currentLevel) > 0 {
		value := make([]int, 0)
		nextLevel := make([]*TreeNode, 0)
		for _, v := range currentLevel {
			value = append(value, v.Val)
			if v.Left != nil {
				nextLevel = append(nextLevel, v.Left)
			}
			if v.Right != nil {
				nextLevel = append(nextLevel, v.Right)
			}

		}
		currentLevel = nextLevel
		result = append(result, value)
	}

	//reverse result
	for i := len(result)/2 - 1; i >= 0; i-- {
		opp := len(result) - 1 - i
		result[i], result[opp] = result[opp], result[i]
	}

	return result

}
