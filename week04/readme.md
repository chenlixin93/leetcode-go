- [week04](#week04)
  * [homework](#homework)
  * [二叉堆](#二叉堆)
  * [二叉搜索树](#二叉搜索树)
  * [二分查找](#二分查找)
  * [三分查找](#三分查找)

# week04

## homework

- [设计推特（Medium）](https://leetcode-cn.com/problems/design-twitter/)

```go

```

- [数据流的中位数（选做）（Hard）](https://leetcode-cn.com/problems/find-median-from-data-stream/)

```go

```

- [寻找旋转排序数组中的最小值 II （Hard）](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/)

```go

```

## 二叉堆

- [golang中container/heap包用法](https://studygolang.com/articles/5415) 

- [合并 K 个升序链表（Hard）](https://leetcode-cn.com/problems/merge-k-sorted-lists/)

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if lists == nil || len(lists) == 0 {
        return nil
    }

    var h IntHeap // 最小堆用于维护当前k个节点
    heap.Init(&h) // 用于节点间的比较
    for _, list := range lists {
        if list != nil {
            heap.Push(&h, list)
        }
    }

    // 定义节点用于统一处理
    dummy := &ListNode{}
    p := dummy // p初始指向dummy节点

    //  当最小堆不为空时
    for h.Len() > 0 {
        min := heap.Pop(&h).(*ListNode) // 取出堆顶元素，即取出最小值节点
        p.Next = min // 游标p指向最小值节点
        p = p.Next // p向后移动一个位置
        // 确定一个节点在合并链表中的位置
        if min.Next != nil {
            heap.Push(&h, min.Next) // 如果最小值节点后面的节点非空，则把该节点加入最小堆中
        }
    }
    return dummy.Next
}

type IntHeap []*ListNode

func (h IntHeap) Len() int {
    return len(h)
}

func (h IntHeap) Less(i,j int) bool {
    return h[i].Val < h[j].Val
}

func (h IntHeap) Swap(i,j int) {
    h[i],h[j] = h[j],h[i]
}

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(*ListNode)) // .(type)
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n - 1]
    *h = old[0:n-1] // 前闭后开
    return x 
}
```

- [滑动窗口最大值（Hard）](https://leetcode-cn.com/problems/sliding-window-maximum/)

```go
import (
    "sort"
)

func maxSlidingWindow(nums []int, k int) []int {
    a = nums
    q := &hp{make([]int, k)} // 设定窗口
    for i := 0; i < k; i++ {
        q.IntSlice[i] = i
    }
    heap.Init(q)

    n := len(nums)
    ans := make([]int, 1, n - k + 1)
    ans[0] = nums[q.IntSlice[0]]
    for i := k; i < n; i++ {
        heap.Push(q, i)
        for q.IntSlice[0] <= i - k {
            heap.Pop(q)
        }
        ans = append(ans, nums[q.IntSlice[0]])
    }
    return ans
}

var a []int
type hp struct{ sort.IntSlice }

func (h hp) Less(i,j int) bool { // 重载
    return a[h.IntSlice[i]] > a[h.IntSlice[j]]
}

func (h *hp) Push(v interface{}) {
    h.IntSlice = append(h.IntSlice, v.(int)) // 整型
}

func (h *hp) Pop() interface{} {
    a := h.IntSlice
    v := a[len(a) - 1]
    h.IntSlice = a[:len(a) - 1] // 前闭后开
    return v
}
```

## 二叉搜索树

- [二叉搜索树中的插入操作（Medium）](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/)

```go
// 模板
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// 返回插入以后的新树根
    // （如果root本来就存在，还是返回root；如果root是空，返回新建的点）
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}
```

- [后继者（Medium）](https://leetcode-cn.com/problems/successor-lcci/)

```go

```

- [删除二叉搜索树中的节点（Medium）](https://leetcode-cn.com/problems/delete-node-in-a-bst/)

```go
// 模板
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil { return nil }
	if root.Val == key {
		if root.Left == nil { return root.Right } // 没有左子树，让right代替root的位置
		if root.Right == nil { return root.Left } // 没有右子树

		next := root.Right
		for next.Left != nil { next = next.Left } // 左右子树都不为空，找后继，右子树一路向左。(二叉搜索树特性)
		root.Right = deleteNode(root.Right, next.Val) // 找到了，现在右子树里面删除
		root.Val = next.Val // 将后继结点的值替换到当前位置
		return root
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
```

- [把二叉搜索树转换为累加树（Medium）](https://leetcode-cn.com/problems/convert-bst-to-greater-tree/)

```go

```

## 二分查找

- [二分查找（Easy）](https://leetcode-cn.com/problems/binary-search/)

```go

```

- [在排序数组中查找元素的第一个和最后一个位置（Medium）](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/)

```go
// 模板？？
func searchRange(nums []int, target int) []int {
	var ans = make([]int, 2)
	
	left := 0
	right := len(nums) //  0...n-1; n表示不存在
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] >= target { 
			right = mid 
		} else {
			left = mid + 1
		}
	}
	ans[0] = right // 第一个 >=target 的数

	left = -1
	right = len(nums) - 1 // 0...n-1; -1表示不存在
	for left < right {
		mid := (left + right + 1) >> 1 //  -1 + 0
		if nums[mid] <= target { 
			left = mid 
		} else {
			right = mid - 1
		}
	}
	ans[1] = right // 最后一个 <=target 的数

	if ans[0] > ans [1] {
		ans = []int{-1, -1}
	}
	return ans
}
```

- [x 的平方根（Easy）](https://leetcode-cn.com/problems/sqrtx/)

```go
func mySqrt(x int) int {
	left := 0
	right := x
	for left < right {
		mid := (left + right + 1) >> 1
		if mid * mid <= x {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return right
}
```

- [搜索二维矩阵（Medium）](https://leetcode-cn.com/problems/search-a-2d-matrix/)

```go

```

- [寻找旋转排序数组中的最小值（Medium）](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)

```go
func findMin(nums []int) int {
	// 题目条件：找第一个<=末尾的数
	l := 0
	r := len(nums) - 1

	for l < r {
		mid := (l + r) >> 1
		if nums[mid] <= nums[r] { 
			r = mid 
		} else { 
			l = mid + 1 
		}
	}
	return nums[l]
}
```

## 三分查找

- [寻找峰值（Medium）](https://leetcode-cn.com/problems/find-peak-element/)
- [猜数字大小（Easy）](https://leetcode-cn.com/problems/guess-number-higher-or-lower/)
- [分割数组的最大值（Hard）](https://leetcode-cn.com/problems/split-array-largest-sum/)
- [制作 m 束花所需的最少天数（Medium）](https://leetcode-cn.com/problems/minimum-number-of-days-to-make-m-bouquets/)