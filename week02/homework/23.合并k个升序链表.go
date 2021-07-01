/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	// 官方题解转化为 golang
	// 分治，链表数组区间一分为二，最终到达两个链表，
	// 开始合并两个链表，合并之后返还给上一层继续合并，最终合并为一个有序链表
	var mergeTwoLists func(a, b *ListNode) *ListNode
	mergeTwoLists = func(a, b *ListNode) *ListNode {
		if a == nil || b == nil {
			if a != nil {
				return a
			}
			return b
		}

		head := &ListNode{}
		tail := head
		for a != nil && b != nil {
			if a.Val < b.Val {
				tail.Next = a
				a = a.Next
			} else {
				tail.Next = b
				b = b.Next
			}
			tail = tail.Next
		}
		if a != nil {
			tail.Next = a
		} else {
			tail.Next = b
		}
		return head.Next
	}

	var merge func(lists []*ListNode, l, r int) *ListNode
	merge = func(lists []*ListNode, l, r int) *ListNode {
		if l == r { return lists[l] }
		if l > r { return nil }

		mid := (l + r) >> 1
		return mergeTwoLists(merge(lists, l, mid), merge(lists, mid + 1, r))
	}

	return merge(lists, 0, len(lists) - 1)
}
// @lc code=end

