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
type Twitter struct {
	user map[int][]int
	tweet [][2]int
}


/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{user: make(map[int][]int)}
}


/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
	this.tweet = append(this.tweet, [2]int{userId, tweetId}) // 存放用户id - 推特id
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	userIds := this.user[userId]
	userIds = append(userIds, userId) // 把当前用户id推进数组
	res := []int{}
	for i := len(this.tweet) - 1; len(res) < 10 && i >= 0; i-- { // 遍历推特集合
		for j := 0; j < len(userIds); j++ { // 检查推特是否属于关注的用户或者自己
			if this.tweet[i][0] == userIds[j] {
				res = append(res, this.tweet[i][1]) // 存入结果集
				break
			}
		}
	}
	return res
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
	for i := 0; i < len(this.user[followerId]); i++ {
		if this.user[followerId][i] == followeeId {
			return // 重复则退出
		}
	}
	this.user[followerId] = append(this.user[followerId], followeeId)
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	for i := 0; i < len(this.user[followerId]); i++ {
		if this.user[followerId][i] == followeeId {
			// 判断该位置是否在末尾
			if i == len(this.user[followerId]) - 1 {
				this.user[followerId] = this.user[followerId][:i] // 前闭后开
			} else {
				this.user[followerId]  = append(this.user[followerId][:i], this.user[followerId][i+1:]...) // 前闭后开 + 前闭至结束
			}
			return // 找到，处理完则退出
		}
	}
}


/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
```

- [数据流的中位数（选做）（Hard）](https://leetcode-cn.com/problems/find-median-from-data-stream/)

**解法1-简单排序**

```go
// 解题思路
// 简单插入数组
// 找中位数前排好序，长度为奇数取a[mid]，偶数取a[mid] + a[mid - 1]的和 * 0.5
```

**解法2-大顶堆小顶堆**

```go
// 解题思路
// 维护大顶堆和小顶堆，大顶堆的数均小于小顶堆，保持大小顶堆之间相差一个数
// 数据流长度为奇数，那么最终大顶堆会多一个数，则是中位数；否则取两个堆的堆顶元素之和 * 0.5

// 示例
// 数据流[1, 2, 3, 4, 5]依次加入堆
// 大顶堆lo、小顶堆hi; 大顶堆的数均小于小顶堆;
// => 1
// init:  lo[]  hi[]
// step1: lo[1] hi[]
// step2: lo[]  hi[1]
// step3: len(lo) < len(hi) => lo[1] hi[]
// => 2
// init:  lo[1]   hi[]
// step1: lo[2,1] hi[]
// step2: lo[1]   hi[2]
// step3: len(lo) == len(hi), 不操作
// => 3
// init:  lo[1]   hi[2]
// step1: lo[3,1] hi[2]
// step2: lo[1]   hi[2,3]
// step3: len(lo) < len(hi) => lo[2,1] hi[3]
// => 4
// init:  lo[2,1]   hi[3]
// step1: lo[4,2,1] hi[3]
// step2: lo[2,1]   hi[3,4]
// step3: len(lo) == len(hi), 不操作
// => 5
// init:  lo[2,1]   hi[3,4]
// step1: lo[5,2,1] hi[3,4]
// step2: lo[2,1]   hi[3,4,5]
// step3: len(lo) < len(hi) => lo[3,2,1] hi[4,5]

// 最终运行效率，364ms & 23.4MB 有点慢
import "container/heap"

type MedianFinder struct {
    maxHeap *MaxIntHeap
    minHeap *MinIntHeap
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
    res := MedianFinder{maxHeap: &MaxIntHeap{},minHeap: &MinIntHeap{}}
    heap.Init(res.maxHeap)
    heap.Init(res.minHeap)
    return res
}


func (this *MedianFinder) AddNum(num int)  {
    heap.Push(this.maxHeap, num) // 先入大顶堆

    heap.Push(this.minHeap, heap.Pop(this.maxHeap)) // 从大顶堆Pop出堆顶元素（即当前堆最大数）

    for this.maxHeap.Len() < this.minHeap.Len() { // 保持大小顶堆之间相差一个数
        heap.Push(this.maxHeap, heap.Pop(this.minHeap))
    }
}


func (this *MedianFinder) FindMedian() float64 {
    if this.maxHeap.Len() > this.minHeap.Len() {
        return float64(this.maxHeap.Top())
    }

    return float64(this.maxHeap.Top() + this.minHeap.Top()) * 0.5
}

// 实现小顶堆
type MinIntHeap []int

func (h MinIntHeap) Len() int {
    return len(h)
}

func (h MinIntHeap) Less(i,j int) bool {
    return h[i] < h[j]
}

func (h MinIntHeap) Swap(i,j int) {
    h[i],h[j] = h[j],h[i]
}

func (h *MinIntHeap) Push(x interface{}) {
    *h = append(*h, x.(int)) // .(type)
}

func (h *MinIntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n - 1]
    *h = old[0:n-1] // 前闭后开
    return x 
}

func (h *MinIntHeap) Top() int {
    old := *h
    x := old[0]
    return x
}

// 实现大顶堆
type MaxIntHeap []int

func (h MaxIntHeap) Len() int {
    return len(h)
}

func (h MaxIntHeap) Less(i,j int) bool {
    return h[i] > h[j]
}

func (h MaxIntHeap) Swap(i,j int) {
    h[i],h[j] = h[j],h[i]
}

func (h *MaxIntHeap) Push(x interface{}) {
    *h = append(*h, x.(int)) // .(type)
}

func (h *MaxIntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n - 1]
    *h = old[0:n-1] // 前闭后开
    return x 
}

func (h *MaxIntHeap) Top() int {
    old := *h
    x := old[0]
    return x
}
/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
```

- [寻找旋转排序数组中的最小值 II （Hard）](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/)

```go
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	l := 0
	r := len(nums) - 1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] > nums[r] {
			l = mid + 1
		} else if nums[mid] < nums[r] {
			r = mid
		} else {
			r--
		}
	}
	return nums[l]
}
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    var find func(root *TreeNode, val int) *TreeNode
    find = func(root *TreeNode, val int) *TreeNode {
        var ans *TreeNode
        cur := root

        for cur != nil {
            if cur.Val > val { //case2: 当后继存在于经过的点(找到一个>val的最小节点)
                if ans == nil || ans.Val > cur.Val {
                    ans = cur
                }
            }
            if cur.Val == val {
                if cur.Right != nil { //case1: 检索到val且右子树存在，右子树一路向左
                    cur = cur.Right
                    for cur.Left != nil { cur = cur.Left }
                    return cur
                }
                break
            }
            if val < cur.Val { 
                cur = cur.Left
            } else {
                cur = cur.Right
            }
        }
        return ans
    }
    return find(root, p.Val)
}
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
// 模版2
func search(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
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
// 简单易懂，双百解法
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }

    row := len(matrix) - 1 // 从最后一行搜起
    col := 0
    for row >= 0 && col < len(matrix[0]) {
        if matrix[row][col] == target {
            return true
        }

        if matrix[row][col] < target { 
            col++ // 大于左下角，则向右移动一位
        } else {
            row-- // 小于左下角，则退后一行
        }
    }
    return false
}

// 二分，后续补上
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

```go

```
- [猜数字大小（Easy）](https://leetcode-cn.com/problems/guess-number-higher-or-lower/)

```go

```

- [分割数组的最大值（Hard）](https://leetcode-cn.com/problems/split-array-largest-sum/)

```go
func splitArray(nums []int, m int) int {
	// 求分割数量达到 m 个的时候，子数组和的最大值最小
	// 也就是说存在一个区间，最大值从左边界到右边界移动时，
	// 第一个达到的值即为最小(符合条件 m 个非空的连续子数组)

	left := 0 // 下界：每个分一组
	right := 0 // 上界：全部放一组
	for i := 0; i < len(nums); i++ {
		left = max(left, nums[i])
		right += nums[i]
	}
	for left < right {
		mid := (left + right) >> 1
		// 第一个使得判定问题得到true的位置
		if isValid(nums, m, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right 
	// 二分模板！！！
}

// 判定把nums分成 <= m 组，每组和 <= T
func isValid(nums []int, m int, T int) bool {
	groupSum := 0
	groupCount := 1
	for i := 0; i < len(nums); i++ {
		if groupSum + nums[i] <= T {
			groupSum += nums[i] // 放进当前数组，不超
		} else {
			groupCount++ // 超了，新开一组
			groupSum = nums[i]
		}
	}
	return groupCount <= m
}

func max(a,b int) int {
	if a > b { return a }
	return b
}
```

- [制作 m 束花所需的最少天数（Medium）](https://leetcode-cn.com/problems/minimum-number-of-days-to-make-m-bouquets/)

```go
func minDays(bloomDay []int, m int, k int) int {
    MAX := 1000000001
    left := 0
    right := MAX
    for left < right {
        mid := (left + right) >> 1
        if bouquetsOnDay(mid, bloomDay, k) >= m { // ?
            right = mid
        } else {
            left = mid + 1
        }
    }
    if right == MAX {
        right = -1
    }
    return right
}

func bouquetsOnDay(day int, bloomDay []int, k int) int {
    bouquets := 0
    consecutive := 0
    for _, bloom := range bloomDay {
        if bloom <= day {
            consecutive++
            if consecutive == k {
                bouquets++ // 成一束, 另起一束
                consecutive = 0
            }
        } else {
            consecutive = 0 // 当前数字大于假定的开花时间，不能开花
        }
    }
    return bouquets // 返回花束数量
}
```