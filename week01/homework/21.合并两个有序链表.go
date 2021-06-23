/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    // 定义一个结果节点
    var res *ListNode
    // 当l1节点的值大于l2节点的值，那么res指向l2的节点，从l2开始遍历，反之从l1开始
    if l1.Val >= l2.Val {
        res = l2
        res.Next = mergeTwoLists(l1, l2.Next)
    } else {
        res = l1
        res.Next = mergeTwoLists(l1.Next, l2)
    }
    return res
}
// @lc code=end

