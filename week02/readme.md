[TOC]

# week02

## homework

- [LRU 缓存机制（Medium）](https://leetcode-cn.com/problems/lru-cache/)

```go
/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start
type LRUCache struct {
	size int
	capacity int
	cache map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key: key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache: map[int]*DLinkedNode{},
		head: initDLinkedNode(0, 0),
		tail: initDLinkedNode(0, 0),
		capacity: capacity,
	}
	// 先把头结点和尾结点连接起来
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node) // 每次使用后将此结点移动至链表头部
	return node.value
}

func (this *LRUCache) Put(key int, value int)  {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++

		if this.size > this.capacity {
			remove := this.removeTail()
			delete(this.cache, remove.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head // 将当前节点的前驱指向头结点
	node.next = this.head.next // 将当前节点的后驱结点指向原来的next节点
	this.head.next.prev = node // 将原来的next节点的前驱结点指向插入的节点
	this.head.next = node // 将头结点的后驱结点指向当前节点
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev // 尾结点的前驱结点指向了最后一个有效节点
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
```

- [子域名访问计数（Easy）](https://leetcode-cn.com/problems/subdomain-visit-count/)

```go
/*
 * @lc app=leetcode.cn id=811 lang=golang
 *
 * [811] 子域名访问计数
 */

// @lc code=start
func subdomainVisits(cpdomains []string) (ans []string) {
	count := make(map[string]int, 0)
	for _, v := range cpdomains {
		splits := strings.Split(v, " ")
		num, _ := strconv.Atoi(splits[0])
		for {
			count[splits[1]] += num
			// 找当前域名第一个点的位置
			dotIndex := strings.Index(splits[1], ".")
			if dotIndex < 0 { break }
			// 得到更高一级的域名
			splits[1] = splits[1][dotIndex+1:] 
		}
	}

	//ans := make([]string, 0, len(count))
	for k, v := range count {
		ans = append(ans, fmt.Sprintf("%d %s", v, k))
	}
	return
}
// @lc code=end
```

- [数组的度（Easy）](https://leetcode-cn.com/problems/degree-of-an-array/)

```go
/*
 * @lc app=leetcode.cn id=697 lang=golang
 *
 * [697] 数组的度
 */

// @lc code=start
func findShortestSubArray(nums []int) (ans int) {
	mp := map[int]entry{}
	for i, v := range nums {
		if e, has := mp[v]; has {
			e.cnt++
			e.r = i
			mp[v] = e
		} else {
			mp[v] = entry{1, i, i} // 出现的次数，首次出现的位置，最后出现的位置
		}
	}
	maxCnt := 0
	for _, e := range mp {
		if e.cnt > maxCnt {
			maxCnt, ans = e.cnt, e.r - e.l + 1
		} else if e.cnt == maxCnt {
			ans = min(ans, e.r - e.l + 1)
		}
	}
	return
}

type entry struct {
	cnt, l, r int
}

func min(a,b int) int {
	if a < b { return a}
	return b
}
// @lc code=end
```

- [元素和为目标值的子矩阵数量（Hard）](https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target/)

```go
/*
 * @lc app=leetcode.cn id=1074 lang=golang
 *
 * [1074] 元素和为目标值的子矩阵数量
 */

// @lc code=start
func numSubmatrixSumTarget(matrix [][]int, target int) (ans int) {
	var subarraySum func(nums []int, k int) int
	subarraySum = func(nums []int, k int) (count int) {
		mp := map[int]int{0: 1} // 和为0有一次
		for i, pre := 0, 0; i < len(nums); i++ {
			pre += nums[i]
			if _, ok := mp[pre - k]; ok {
				count += mp[pre - k]
			}

			mp[pre]++
		}
		return
	}

	// 上下边界构成矩形，即二维的数组
	// 二维数组可以通过列压缩成一维组，由子矩形转化为子数组的问题求解
	for i := range matrix {// 枚举上边界
		columnSum := make([]int, len(matrix[0]))
		for _, row := range matrix[i:] { // 枚举下边界
			for c, v := range row {
				columnSum[c] += v // 更新每列的元素和
			}
			ans += subarraySum(columnSum, target)
		}
	}
	return
}
// @lc code=end
```

- [合并K 个升序链表（Hard） (要求：用分治实现，不要用堆)](https://leetcode-cn.com/problems/merge-k-sorted-lists/)

```go
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
```

## 哈希表、集合、映射

- [两数之和](https://leetcode-cn.com/problems/two-sum/description/)

```go
/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
    
    value_to_key := map[int]int{}

    for cur_key, v := range nums {
        key, ok := value_to_key[v]
        if ok {
            return []int{cur_key, key}
        } else {
            value_to_key[target - v] = cur_key
        }
    }

    return []int{}
}
// @lc code=end
```

- [字母异位词分组](https://leetcode-cn.com/problems/group-anagrams/)

```go
/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
    mp := map[string][]string{}

    for _, str := range strs {
        s := []byte(str) // 转成字节数组
        //fmt.Println(s)
        sort.Slice(s, func(i, j int) bool {return s[i] < s[j]}) // map-reduce模型
        sortedStr := string(s)
        mp[sortedStr] = append(mp[sortedStr], str) // $mp['aet'][] = 'aet'
    }

    ans := make([][]string, 0, len(mp))
    for _, v := range mp {
        ans = append(ans, v)
    }
    return ans
}
// @lc code=end
```

- [串联所有单词的子串](https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/)

```go
/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	// 没有单词时，不符合条件
	if (len(words) < 1) { return []int{} }

	word_nums := len(words)
	word_len := len(words[0])
	// 需要用上所有单词，如果字符串长度小于单词总长度，不符合条件
	if len(s) < word_nums * word_len { return []int{} }

	var map_1 = make(map[string]int)
	for _, v := range words {
		map_1[v]++ // 统计单词出现次数
	}

	var res []int
	// [当前下标+单词总长度]不会越界
	for i := 0; i + (word_nums * word_len) <= len(s); i++ {
		var count int
		var map_2 = make(map[string]int)
		for step := 0; step < word_nums; step++ {
			start := i + step * word_len
			word := s[start:start+word_len]
			if nums, found := map_1[word]; found && nums > map_2[word] {
				// 在map_1中能找到当前s截取到的单词，并且之前统计的次数大于map2中的次数
				// 也就是单词组的出现次数能覆盖字符串中的出现次数
				map_2[word]++
				count++
			} else {
				break;
			}
		}
		if count == word_nums {
			res = append(res, i)
		}
	}

	return res
}
// @lc code=end
```

## LRU

- [LRU 缓存机制](https://leetcode-cn.com/problems/lru-cache/)

```go
同作业题
```

## 递归

- [子集](https://leetcode-cn.com/problems/subsets/)

```go
/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) (ans [][]int) {
	// 官方题解，简洁
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		// 考虑选择当前位置
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1] // 恢复现场

		// 考虑不选择当前位置
		dfs(cur + 1)
	}
	dfs(0)
	return
}
// @lc code=end
```

- [组合](https://leetcode-cn.com/problems/combinations/)

```go
/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) (ans [][]int) {
	// 官方题解，简洁
	temp := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		// 剪枝：temp长度加上区间 [cur, n]的长度小于k，不可能构造出长度为 k 的 temp
		if len(temp) + (n - cur + 1) < k {
			return
		}

		if len(temp) == k {
			comb := make([]int, k)
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}
		// 考虑选择当前位置
		temp = append(temp, cur)
		dfs(cur + 1)
		temp = temp[:len(temp) - 1]

		// 考虑不选择当前位置
		dfs(cur + 1)
	}
	dfs(1)
	return
}
// @lc code=end
```

- [全排列](https://leetcode-cn.com/problems/permutations/)

```go
/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) (ans [][]int) {
	// 来自笨猪爆破组，返回值稍加改动，有官方题解那味儿了
	visited := map[int]bool{}

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}

		for _, v := range nums {
			if visited[v] {
				continue
			}
			path = append(path, v)
			visited[v] = true
			dfs(path)
			path = path[:len(path) - 1]
			visited[v] = false
		}
	}
	dfs([]int{})
	return
}
// @lc code=end
```

## 树

- [翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/description/)

```go
/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
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
 func invertTree(root *TreeNode) *TreeNode {
    if root == nil { return nil }

    left := invertTree(root.Left)
    right := invertTree(root.Right)

    root.Left = right
    root.Right = left

    return root
}
// @lc code=end
```

- [验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree/)

```go
/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower,upper int)  bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
// @lc code=end
```

- [二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)

```go
/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a int, b int) int {
	if a > b { return a }
	return b
}
// @lc code=end
```

- [二叉树的最小深度](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/)

```go
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
```

## 分治

- [Pow(x, n)](https://leetcode-cn.com/problems/powx-n/)

```go
/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n > 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n) // 负负为正
}

func quickMul(x float64, n int) float64 {
	if n == 0 { return 1 }

	y := quickMul(x, n/2)
	if n%2 == 0 { return y * y }
	return y * y * x // n为奇数，n/2 向下取整，所以这里还得乘回1个x
}
// @lc code=end
```

- [括号生成](https://leetcode-cn.com/problems/powx-n/)

```go
/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) (ans []string) {
	// L 、R 指左右括号剩下的可用长度
	var dfs func(L,R int, path string)
	dfs = func(L,R int, path string)  {
		if 2*n == len(path) {
			ans = append(ans, path)
			return
		}

		if L > 0 {
			dfs(L - 1, R, path + "(")
		}
		if L < R {
			dfs(L, R - 1, path + ")")
		}
	}

	dfs(n, n, "")
	return
}
// @lc code=end
```