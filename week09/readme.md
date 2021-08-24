# week09

## homework

- [二进制矩阵中的最短路径（Medium）](https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/)

```go
func shortestPathBinaryMatrix(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	if grid[0][0] == 1 || grid[m - 1][n - 1] == 1 { // 特判
		return -1
	}
	var q [][2]int
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}

	q = append(q, [2]int{0, 0}) // 起点入队列
	visited[0][0] = 1 // 标记为已访问

	// 定义八个方向
	dx := [8]int{-1, 1, 0, 1, 0, -1, 1, -1}
	dy := [8]int{-1, 1, 1, 0, -1, 0, -1, 1}
	ans := 0

	// BFS模板
	for len(q) > 0 {
		size := len(q) // 当前层的长度
		for i := 0; i < size; i++ { // 只处理当前层的数据
			cur := q[0]
			q = q[1:] // 出队一个元素

			x := cur[0]
			y := cur[1]
			if x == m - 1 && y == n - 1 { // 终止条件
				return ans + 1
			}

			for j := 0; j < 8; j++ {
				nx := x + dx[j]
				ny := y + dy[j]
				// 检查边界、是否已访问过，不符合题意的情况
				if nx < 0 || nx >= m || ny < 0 || ny >= n || visited[nx][ny] == 1 || grid[nx][ny] == 1 { 
					continue
				}
				q = append(q, [2]int{nx, ny}) // 入队，下一层进行遍历
				visited[nx][ny] = 1
			}
		}
		ans++ // 当前层数
	}
	return -1
}
```

> 大多数情况下不会自己去实现平衡树或跳表
可以用语言内置的有序集合库
C++: set, multiset, map, multimap
Java: TreeSet, TreeMap
Python: OrderedDict

- 尝试用语言内置的有序集合库，或写一棵平衡树，来解决 [滑动窗口最大值（Hard）](https://leetcode-cn.com/problems/sliding-window-maximum/)

**解法1：手动快排keys【失败】**

```go
// go 实现 java treeMap 的效果
// 47 / 61 个通过测试用例 状态：超出时间限制【数据量大过不了。。。】
func maxSlidingWindow(nums []int, k int) []int {
    // 值 =》出现的次数
    m := map[int]int{}
    // 先设定第一个窗口
    for i := 0; i < k; i++ {
        if _, ok := m[nums[i]]; ok {
            m[nums[i]] += 1
        } else {
            m[nums[i]] = 1
        }
    }
    // 每移动一步，每个窗口都需要取出一个最大值，那么结果集的长度=1 + (len - k)
    n := len(nums)
    var ans []int
    m, last_key := sortMapByKey(m)
    ans = append(ans, last_key)
    // 窗口开始滑动
    for i := k; i < n; i++ {
        // 插入新元素
        if _, ok := m[nums[i]]; ok {
            m[nums[i]] += 1
        } else {
            m[nums[i]] = 1
        }
        // 有元素滑出窗口，那么判断它出现次数，如果当前次数只有1，那么可以将其从窗口删除
        if out_of_window_count, ok := m[nums[i - k]]; ok {
            if out_of_window_count == 1 {
                delete(m, nums[i - k]) // 该元素移出窗口
            } else {
                m[nums[i - k]] = out_of_window_count - 1  // 出现次数减一
            }
        }
        // 取出此时窗口的最大值
        m, last_key = sortMapByKey(m)
        ans = append(ans, last_key)
    }
    return ans
}

// 返回排好序的map和最大key（窗口最大值）
func sortMapByKey(m map[int]int) (map[int]int, int) {
    var keys []int
    for k := range m {
        keys = append(keys, k)
    }
    sort.Ints(keys) // key排序

    sorted_m := map[int]int{}
    var last_key int
    for _, k := range keys {
        sorted_m[k] = m[k]
        last_key = k // 记录最后的key，即最大key
    }
    return sorted_m, last_key
}
```

**解法2：使用内置treeMap【通过，效率较低】**

```go
import "github.com/emirpasic/gods/maps/treemap"

func maxSlidingWindow(nums []int, k int) []int {
    // 值 =》出现的次数
    m := treemap.NewWithIntComparator()
    // 先设定第一个窗口
    for i := 0; i < k; i++ {
        if count, ok := m.Get(nums[i]); ok {
            m.Put(nums[i], count.(int) + 1) // need type assertion
        } else {
           m.Put(nums[i], 1) 
        }
    }
    // 每移动一步，每个窗口都需要取出一个最大值，那么结果集的长度=1 + (len - k)
    n := len(nums)
    var ans []int
    last_key, _ := m.Max()
    ans = append(ans, last_key.(int)) // need type assertion
    // 窗口开始滑动
    for i := k; i < n; i++ {
        // 插入新元素
        if count, ok := m.Get(nums[i]); ok {
            m.Put(nums[i], count.(int) + 1) // need type assertion
        } else {
           m.Put(nums[i], 1) 
        }
        // 有元素滑出窗口，那么判断它出现次数，如果当前次数只有1，那么可以将其从窗口删除
        if out_of_window_count, ok :=  m.Get(nums[i - k]); ok {
            if out_of_window_count == 1 {
                m.Remove(nums[i - k])  // 该元素移出窗口
            } else {  
                m.Put(nums[i - k], out_of_window_count.(int) - 1) // 出现次数减一
            }
        }
        // 取出此时窗口的最大值
       last_key, _ := m.Max()
        ans = append(ans, last_key.(int))
    }
    return ans
}
```

**解法3：优先队列（懒惰删除）**
```go
// 实现优先队列
import (
    "sort"
)

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

func maxSlidingWindow(nums []int, k int) []int {
    // 思路：懒惰删除
    // 延迟到 当未删除的值 会影响答案时 再进行
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
        for q.IntSlice[0] <= i - k { // 懒惰删除，检查【堆顶下标】是否在窗口内
            heap.Pop(q)
        }
        ans = append(ans, nums[q.IntSlice[0]])
    }
    return ans
}
```

- 尝试用语言内置的有序集合库，或写一棵平衡树，来解决 [邻值查找（Medium）AcWing](https://www.acwing.com/problem/content/138/)

```go
```

- [设计跳表（选做）（Hard）]()

```go
```

- [普通平衡树（选做）（Medium）AcWing](https://www.acwing.com/problem/content/255/s)

```go
```

## Part 1

### 搜索剪枝

- 引入

```go
// 复习：蛮力搜索的三种基本类型
递归形式        时间复杂度      问题举例
指数型          k^n           子集、大体积背包（每个数选或者不选）
排列型          n!            全排列、旅行商、N皇后（每个数放之前有多少种摆法）
组合型          n!/(m!(n-m)!) 组合选数（数学公式）

// 复习：初级搜索
搜索方向：
深度优先（DFS）
广度优先（BFS）

简单的优化：
判重（避免重复搜索某一状态）

状态空间、搜索树或图等概念
DFS、BFS对状态空间的遍历形成一棵树（或一张图）

// 剪枝
具有多项式时间算法的问题只是很一部分，更多的问题只能采取搜索的方式求解
蛮力搜索（或者说回溯）作为最原始的遍历状态空间的方法，本质上是试错
一个分支不可行，就需要取消上一步甚至前几步的计算，换个分支重新来过
在分支较多、较深的问题中，很容易导致复杂度为指数时间的运算

剪枝，就是通过已有的信息，提前判定某些分支不可行或者一定不由，从而减少需要访问的状态数量，
形象地说就像剪去“搜索树”的枝条，所以叫剪枝
```

- [括号生成（Medium）](https://leetcode-cn.com/problems/generate-parentheses/)

蛮力搜索：指数型，每个位置放两种括号之一，最后验证，O(n*2^n)

剪枝：实时维护左右括号的数量，不合法及时停止

```go
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
```

- [N 皇后（Hard）](https://leetcode-cn.com/problems/n-queens/)

蛮力搜索：排列型，每行放的皇后的列号是一个排列，最后验证斜线，O(n*n!)

剪枝：维护两种斜线（行号+列号、行号-列号）的已用值集合，排除造成重复的分支

```go
func solveNQueens(n int) (res [][]string) {
    ans := [][]int{}

    var find func(row int, s []int, used,mapusedIplusJ,usedIminusJ map[int]bool)
    find = func(row int, s []int, used,usedIplusJ,usedIminusJ map[int]bool) {
        if row == n {
            //fmt.Println(s)
            tmp := make([]int, n)
            copy(tmp, s)
            ans = append(ans, tmp)
            //fmt.Println(ans)
            return
        }
        for col := 0; col < n; col++ {
            // 剪枝
            if !used[col] && !usedIplusJ[row + col] && !usedIminusJ[row - col] {// 当前列、左右对角线不能存在值，否则会被攻击
                used[col] = true
                usedIplusJ[row + col] = true
                usedIminusJ[row - col] = true
                s = append(s, col)
                find(row + 1, s, used, usedIplusJ, usedIminusJ)
                s = s[0:len(s) - 1]
                usedIminusJ[row - col] = false
                usedIplusJ[row + col] = false
                used[col] = false
            }
        }

    }

    s := []int{}
    used := map[int]bool{}
    usedIplusJ := map[int]bool{}
    usedIminusJ := map[int]bool{}
    find(0, s, used, usedIplusJ, usedIminusJ)
    // fmt.Println(ans)

    // 打印答案部分
    for _,v := range ans {
        str_arr := []string{}
        for _,pos := range v {
            str := ""
            for i := 0; i < n; i++ {
                if i != pos {
                    str = fmt.Sprintf("%s%s", str, ".")
                } else {
                    str = fmt.Sprintf("%s%s", str, "Q")
                }
            }
            str_arr = append(str_arr, str)
        }
        res = append(res, str_arr)
    }
    return 
}
```

- [有效的数独（Medium）](https://leetcode-cn.com/problems/valid-sudoku/)

```go
func isValidSudoku(board [][]byte) bool {
	row := make(map[int]map[byte]bool) // 9 行
	col := make(map[int]map[byte]bool) // 9 列
	box := make(map[int]map[byte]bool) // 9 个 box
	for i := 0; i < 9; i++ { // 需要初始化，否则会报 nil map
		row[i] = map[byte]bool{}
		col[i] = map[byte]bool{}
		box[i] = map[byte]bool{}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if string(board[i][j]) != "." {
				digit := board[i][j]
				box_id := i/3 * 3 + j/3 // 切割成3*3=9个盒子，再转化成一维数字，编号0~8
				if row[i][digit] { return false } // 如果当前行已经出现该数字，返回false
				row[i][digit] = true
				if col[j][digit] { return false } // 如果当前列已经出现该数字，返回false
				col[j][digit] = true
				if box[box_id][digit] { return false } // 如果当前box已经出现该数字，返回false
				box[box_id][digit] = true
			}
		}
	}
	return true
}
```

- [解数独（Hard）](https://leetcode-cn.com/problems/sudoku-solver/)

> 优先选择“能填的合法数字最少的位置”，而不是第一个位置

如果是人类玩数独，策略一定是先填上“已经能够唯一确定的位置”，然后从那些填的比较满、选项比较少的位置突破

> 快速判断数独有效性

对于每行、每列、每个九宫格，分别用一个9位bool数组保存那些数字可填

对于一个位置，合并它所在的行、列、九宫格的3个bool数组，就可以得到能填的数字

当一个位置填上某个数后，更新对应的行、列、九宫格bool数组，回溯时还原现场

```go
/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte)  {
	row := make([][10]bool, 9) // [行号][1-9是否可用]
	col := make([][10]bool, 9) // [列号][1-9是否可用]
	box := make([][10]bool, 9)
	for i := 0; i < 9; i++ {
		for j := 0; j < 10; j++ {
			row[i][j] = true
			col[i][j] = true
			box[i][j] = true
		}
	}
	// 预处理数字的可用情况
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
                // 处理为不可用
				digit := board[i][j] - byte('0')
				row[i][digit] = false
				col[j][digit] = false
				box[i/3*3+j/3][digit] = false
			}
		}
	}
	dfs(board, row, col, box)
}

func dfs(board [][]byte, row,col,box [][10]bool) bool {
	// 找到第一个可填的数
	location := getLeastPossibleLocation(board, row, col, box)
	x := location[0]
	y := location[1]
	if x == -1 { return true } // 填满了，有解
	// 尝试填入1-9
	for digit := 1; digit <= 9; digit++ {
		boxId := x/3*3+y/3
        // 都可用时，才进入dfs（剪枝1）
		if row[x][digit] && col[y][digit] && box[boxId][digit] {
			// 回溯模板
			row[x][digit] = false
			col[y][digit] = false
			box[boxId][digit] = false
			board[x][y] = byte('0') + byte(digit)
			if dfs(board, row, col, box) { return true }
			board[x][y] = '.'
			row[x][digit] = true
			col[y][digit] = true
			box[boxId][digit] = true
		}
	}
	return false
}

// 遍历每次找第一个空位置

// VS

// 每次找分支较少的一个空（剪枝2）
func getLeastPossibleLocation(board [][]byte, row,col,box [][10]bool) [2]int {
	ansCnt := 10
	ans := [2]int{-1, -1}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// 如果是小数点
			if board[i][j] == '.' {
				cnt := 0
				// 统计当前位置还有几个数可填，取（分支，可能性）最小的坐标
				for digit := 1; digit <= 9; digit++ {
					if row[i][digit] && col[j][digit] && box[i/3*3+j/3][digit] {
						cnt++
					}
				}
				if cnt < ansCnt {
					ansCnt = cnt
					ans = [2]int{i, j}
				}
			}
		}
	}
	return ans
}
```

### 迭代加深、折半搜索与双向搜索

- 引入

```go
// 普通BFS：利用队列特性实现层序遍历
// slice 当作队列
var queue []int // 初始化
queue = append(queue, 1) // 入队一个元素
if len(queue) > 1 {
  queue = queue[1:] // 出队一个元素
}

// queue 当作队列 
// https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter03/03.3.html
import "container/list"
queue := list.New()
queue.PushBack(1)
if queue.Len() > 1 {
  queue.Remove(queue.Front())
}

// 双向BFS：当起点和终点是唯一时。
适用于层数不太深，但每层分支数量大的问题
从初态和终态出发各搜索一半状态，产生两棵深度减半的搜索树
在中间交会，用适当方法合并成最终的答案

// 迭代加深
对DFS的优化
适用于搜索树最大深度很大，但答案可能并不太深的问题
可以防止DFS一开始选错了分支，在不包含答案的深层子树上浪费许多时间

// 折半搜索
与双向BFS非常类似
双向BFS：有确定的起点、终点，同时开始，每边搜一步
折半搜索：用于集合类的问题，把集合分成两半，分别搜出所有方案，再合起来

//例如大体积背包（子集和问题）？
N个物品选出一些尽量填满体积为M的背包（N个数构成的集合，选出一些，和最接近M）
M可能很大，O(NM)的动态规划实际上是伪多项式算法
可以把集合分成两组，每组搜出所有可能的和
枚举其中一组的和sum，另一组提前拍好序，二分查找M-sum的前驱
O(2^n) 优化到 O(N*2^(n/2))
```

- [单词接龙（Hard）](https://leetcode-cn.com/problems/word-ladder/)

**解法1: 单向BFS**

```go
// slice当作队列
func ladderLength(beginWord string, endWord string, wordList []string) int {
    // BFS
    dist := make(map[string]int)
    for i := 0; i < len(wordList); i++ {
        dist[wordList[i]] = 1e9
    }
    dist[beginWord] = 1 // 初始序列长度

    var depth int
    var queue []string // 初始化队列
    queue = append(queue, beginWord) // 入队第一个元素
    for len(queue) > 0 {
        s := queue[0]
        queue = queue[1:] // 出队第一个元素
        depth = dist[s] // 当前序列长度
        for i := 0; i < len(s); i++ {
            // 当前位置尝试填入a-z的字符
            for ch := 'a'; ch <= 'z'; ch ++ {
                backup := s[i]
                s = s[0:i] + string(ch) + s[i+1:]
                //fmt.Println(s)
                if _, ok := dist[s]; ok {
                    if dist[s] > depth + 1 { // 找到当前字符，序列长度加1
                        dist[s] = depth + 1
                        queue = append(queue, s) // 入队列等待转换
                        if s == endWord { return dist[s] }
                    }
                }
                // 恢复现场
                s = s[0:i] + string(backup) + s[i+1:]
            }
        }
    }
    return 0

} // 212ms

// list当作队列
...
```

**解法2: 双向BFS**

```go
// Go 语言函数引用传递值，解决超时问题
func ladderLength(beginWord string, endWord string, wordList []string) int {
    flag := false
    for i := 0; i < len(wordList); i++ {
        if endWord == wordList[i] {
            flag = true
            break 
        }
    }
    if !flag {return 0}

    dist := make(map[string]int)
    distEnd := make(map[string]int)
    for i := 0; i < len(wordList); i++ {
        dist[wordList[i]] = 1e9
        distEnd[wordList[i]] = 1e9
    }
    dist[beginWord] = 1 // 初始序列长度
    distEnd[endWord] = 1 // 初始序列长度(从endWord往回走)

    var queue []string
    var queueEnd []string
    queue = append(queue, beginWord)
    queueEnd = append(queueEnd, endWord)
    res := -1
    for len(queue) > 0 || len(queueEnd) > 0 {
        res = expand(&queue, &dist, &distEnd)
        if res != -1 {return res}
        res = expand(&queueEnd, &distEnd, &dist)
        if res != -1 {return res}
    }
    return 0
}

// 向前走一步
func expand(queue *[]string, dist *map[string]int, distOther *map[string]int) int {
    var depth int
    for len((*queue)) > 0 {
        s := (*queue)[0]
        (*queue) = (*queue)[1:] // 前闭
        depth = (*dist)[s] // 当前序列长度
        for i := 0; i < len(s); i++ {
            // 当前位置尝试填入a-z的字符
            for ch := 'a'; ch <= 'z'; ch ++ {
                backup := s[i]
                s = s[0:i] + string(ch) + s[i+1:]
                if _, ok := (*dist)[s]; ok {
                    if (*dist)[s] > depth + 1 { // 找到当前字符，序列长度加1
                        (*dist)[s] = depth + 1
                        (*queue) = append((*queue), s) // 入队列等待转换
                        if (*distOther)[s] != 1e9 { // 到了对面搜过的一个状态，相遇了
                            return depth + (*distOther)[s]
                        }
                    }
                }
                // 恢复现场
                s = s[0:i] + string(backup) + s[i+1:]
            }
        }
    }
    return -1
}
```

### 启发式搜索：A* 算法

- 引入

```go
// 普通BFS ：按层扩展

// 优先队列BFS：每次从队列中取出当前代价最小的状态进行扩展
// 局限性：
一个状态的当前代价最小，只能说明从起始状态到该状态的代价很小，
而在未来的搜索中，从该状态到目标状态可能会花费很大的代价。

反之亦然，当前代价较大，也许未来代价较小，总代价反而更优。
优先队列BFS`缺少对未来的预估`

// A*算法 - 估价函数
A*算法是一种启发式搜索算法

A*算法的关键是设计一个`估价函数`：
以任意“状态”为输入，计算出从该状态到目标状态所需代价的`估计值`
在搜索中，维护一个堆（优先队列），优先选择“当前代价+`未来估价`”最小的状态进行扩展

估价函数的设计原则：估值必须比实际更优（估价代价 <= 未来实际代价）
只要保证以下原则，当目标状态第一次从堆中被取出时，就得到了最优解
```

- [滑动谜题（Hard）](https://leetcode-cn.com/problems/sliding-puzzle/)

**解法1：普通BFS**

```go
func slidingPuzzle(board [][]int) int {
    // BFS解法
    // 2*3 => 1*6
    list := make([]int, 6)
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            list = append(list, board[i][j]) // 先转化为一维
        }
    }
    start := zip(list) // 当前状态压缩为一个数
    target := 123450 // 目标状态

    var q []int // 初始化队列
    q = append(q, start) // 入队
    dist := make(map[int]int)
    dist[start] = 0 // 起点不需要移动
    for len(q) > 0 {
        //fmt.Println(q)
        now := q[0]
        q = q[1:] // 出队一个元素
        a := unzip(now) // 解压出一维数组
        pos := getZeroIndex(a)
        // 非两个最左侧，可尝试向左走一步
        if pos != 0 && pos != 3 {   insert(pos, pos - 1, a, now, &q, &dist)    }
        // 非两个最右侧，可尝试向右走一步
        if pos != 2 && pos != 5 {   insert(pos, pos + 1, a, now, &q, &dist)    }
        // 处于第二行，允许往上走一步
        if pos >= 3 {   insert(pos, pos - 3, a, now, &q, &dist)    }
        // 处于第一行，允许往下走一步
        if pos < 3 {    insert(pos, pos + 3, a, now, &q, &dist)    }
        if _,ok := dist[target]; ok {
            return dist[target]
        }
    }
    return -1
}
// 插入新位置
func insert(pos int, newPos int, a []int, now int, q *[]int, dist *map[int]int) {
    //fmt.Println("交换前", a)
    a[pos], a[newPos] = a[newPos], a[pos]
    //fmt.Println("交换后", a)
    next := zip(a)
    if _,ok := (*dist)[next]; !ok || (*dist)[next] > (*dist)[now] + 1 {
        // 如果【没找到next】的状态，或者next的已移动次数 > 当前移动次数，则将当前能到达next的移动次数更新并入队
        (*dist)[next] = (*dist)[now] + 1
        *q = append(*q, next)
    }
    // 恢复现场
    a[pos], a[newPos] = a[newPos], a[pos]
}

// [1,2,3,4,5,0] => 123450
func zip(a []int) int {
    res := 0
    for i := 0; i < len(a); i++ {
        res = res * 10 + (a)[i]
    }
    return res
}
// 123450 => [1,2,3,4,5,0]
func unzip(state int) []int {
    a := make([]int, 6)
    for i := 5; i >= 0; i-- {
        a[i] = state % 10
        state = state / 10 // 整除
    }
    return a
}
// 找到0的位置
func getZeroIndex(a []int) int {
    for i := 0; i < len(a); i++ {
        if a[i] == 0 {return i}
    }
    return -1 //不合法
}
```

**解法2：A*算法**

```go

```

- [八数码（Medium）AcWing](https://www.acwing.com/problem/content/181/)

```go
```

- [八数码（打印方案）（Medium）AcWing](https://www.acwing.com/problem/content/847/)

```go
```