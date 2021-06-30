/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil { return 0 }
	if root.Left == nil && root.Right == nil { return 1 }

	MinD := math.MaxInt32
	if root.Left != nil { MinD = min(minDepth(root.Left), MinD) }
	if root.Right != nil { MinD = min(minDepth(root.Right), MinD) }

	return MinD + 1
}

func min(a int, b int) int {
	if a < b { return a}
	return b
}
// @lc code=end
