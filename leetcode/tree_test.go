package leetcode

import (
	"fmt"
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

/**
1022. Sum of Root To Leaf Binary Numbers
https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/
*/
func sumRootToLeaf(root *TreeNode) int {

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	res := findPath(root)
	var sum int64
	for i := 0; i < len(res); i++ {
		val, err := strconv.ParseInt(res[i], 2, 64)
		if err == nil {
			sum += val
		}
	}
	return int(sum)

}

func findPath(root *TreeNode) []string {

	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	left := findPath(root.Left)
	for i := 0; i < len(left); i++ {
		left[i] = strconv.Itoa(root.Val) + left[i]
	}

	right := findPath(root.Right)
	for i := 0; i < len(right); i++ {
		right[i] = strconv.Itoa(root.Val) + right[i]
	}

	return append(left, right...)
}

func sumRootToLeaf2(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, cur int) int {
	if root == nil {
		return 0
	}
	cur = cur<<1 | root.Val
	if root.Left == nil && root.Right == nil {
		return cur
	}
	return dfs(root.Left, cur) + dfs(root.Right, cur)
}

/**
95. Unique Binary Search Trees II
https://leetcode.com/problems/unique-binary-search-trees-ii/

Dynamic programming, Tree
*/

func generateTrees(n int) []*TreeNode {

	if n == 0 {
		return nil
	}

	return generateTreesWithStartEndValue(1, n)
}

/**
End value should be greater than start value.
*/
func generateTreesWithStartEndValue(start, end int) []*TreeNode {

	if start > end {
		return []*TreeNode{nil}
	}

	if start == end {
		return []*TreeNode{&TreeNode{Val: start}}
	}

	if start == end-1 {
		one := TreeNode{Val: start, Right: &TreeNode{Val: end}}
		two := TreeNode{Val: end, Left: &TreeNode{Val: start}}
		return []*TreeNode{&one, &two}
	}

	res := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {

		//left
		left := generateTreesWithStartEndValue(start, i-1)
		//right
		right := generateTreesWithStartEndValue(i+1, end)

		for j := 0; j < len(left); j++ {
			for k := 0; k < len(right); k++ {
				tempRoot := &TreeNode{Val: i}
				tempRoot.Left = left[j]
				tempRoot.Right = right[k]
				res = append(res, tempRoot)
			}
		}
	}

	return res
}

/**
105. Construct Binary Tree from Preorder and Inorder Traversal
https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
*/
func buildTree(preorder []int, inorder []int) *TreeNode {

	if preorder == nil || len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	root := preorder[0]
	rootNode := &TreeNode{Val: root}
	rootIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root {
			rootIndex = i
			break
		}
	}

	leftInorder := inorder[:rootIndex]
	leftPreorder := preorder[1 : len(leftInorder)+1]
	rightInorder := inorder[rootIndex+1:]
	rightPreorder := preorder[1+len(leftInorder):]

	rootNode.Left = buildTree(leftPreorder, leftInorder)
	rootNode.Right = buildTree(rightPreorder, rightInorder)

	return rootNode

}

/**
106. Construct Binary Tree from Inorder and Postorder Traversal
https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
*/

func buildTree2(inorder []int, postorder []int) *TreeNode {

	if inorder == nil || len(inorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}

	root := postorder[len(postorder)-1]
	rootNode := &TreeNode{Val: root}
	rootIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root {
			rootIndex = i
			break
		}
	}

	leftInorder := inorder[:rootIndex]
	leftPostOrder := postorder[:len(leftInorder)]
	rightInorder := inorder[rootIndex+1:]
	rightPostOrder := postorder[len(leftInorder) : len(leftInorder)+len(rightInorder)]
	rootNode.Left = buildTree(leftInorder, leftPostOrder)
	rootNode.Right = buildTree(rightInorder, rightPostOrder)

	return rootNode

}

/**
108. Convert Sorted Array to Binary Search Tree
https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
*/
func sortedArrayToBST(nums []int) *TreeNode {

	if nums == nil || len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	if len(nums) == 2 {
		return &TreeNode{Val: nums[0], Right: &TreeNode{Val: nums[1]}}
	}

	if len(nums) == 3 {
		return &TreeNode{Val: nums[1], Left: &TreeNode{Val: nums[0]}, Right: &TreeNode{Val: nums[2]}}
	}

	middle := len(nums) / 2

	root := &TreeNode{Val: nums[middle]}
	root.Left = sortedArrayToBST(nums[:middle])
	root.Right = sortedArrayToBST(nums[middle+1:])

	return root
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
109. Convert Sorted List to Binary Search Tree
https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/
*/
func sortedListToBST(head *ListNode) *TreeNode {
	return sortedListToBST2(head, nil)
}

func sortedListToBST2(head *ListNode, tail *ListNode) *TreeNode {

	if head == nil {
		return nil
	}
	// one node
	if head.Next == nil || head == tail {
		return &TreeNode{Val: head.Val}
	}

	// two node
	if head.Next == tail || head.Next.Next == nil {
		return &TreeNode{Val: head.Val, Right: &TreeNode{Val: head.Next.Val}}
	}

	//three node
	if head.Next.Next.Next == nil {
		return &TreeNode{Val: head.Next.Val, Left: &TreeNode{Val: head.Val}, Right: &TreeNode{Val: head.Next.Next.Val}}
	}
	count := 0
	fmt.Println("head", head.Val)
	for pos := head; pos != tail; pos = pos.Next {
		count++
	}
	if tail != nil {
		count++
	}

	middleIndex := count / 2
	midPtr := head
	for i := 0; i < middleIndex; i++ {
		midPtr = midPtr.Next
	}
	fmt.Println("middle", midPtr.Val)

	preMiddlePtr := head
	for preMiddlePtr.Next != midPtr {
		preMiddlePtr = preMiddlePtr.Next
	}
	left := sortedListToBST2(head, preMiddlePtr)
	right := sortedListToBST2(midPtr.Next, tail)
	root := &TreeNode{Val: midPtr.Val, Left: left, Right: right}

	return root

}

/**
114. Flatten Binary Tree to Linked List
https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
*/
func flatten(root *TreeNode) {

	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}

	flatten(root.Right)
	flatten(root.Left)

	left := root.Left
	for left != nil && left.Right != nil {
		left = left.Right
	}

	if left != nil {
		left.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}

}

/**
129. Sum Root to Leaf Numbers
https://leetcode.com/problems/sum-root-to-leaf-numbers/
*/
func sumNumbers(root *TreeNode) int {

	path := dfsPath(root)
	sum := 0

	for _, v := range path {
		i, _ := strconv.Atoi(v)
		sum += i
	}
	return sum

}

func dfsPath(root *TreeNode) []string {

	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	path := make([]string, 0)
	leftPath := dfsPath(root.Left)
	for i := 0; i < len(leftPath); i++ {
		leftPath[i] = strconv.Itoa(root.Val) + leftPath[i]
	}

	rightPath := dfsPath(root.Right)
	for i := 0; i < len(rightPath); i++ {
		rightPath[i] = strconv.Itoa(root.Val) + rightPath[i]
	}
	path = append(path, leftPath...)
	path = append(path, rightPath...)

	return path

}

/**
129. Sum Root to Leaf Numbers
https://leetcode.com/problems/sum-root-to-leaf-numbers/
*/
func sumNumbers2(root *TreeNode) int {

	return sumNumbers3(root, 0)

}

func sumNumbers3(root *TreeNode, high int) int {

	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val + high*10
	}

	return sumNumbers3(root.Left, high*10+root.Val) + sumNumbers3(root.Right, high*10+root.Val)

}
