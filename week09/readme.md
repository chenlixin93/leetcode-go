# week09

## homework

- [二进制矩阵中的最短路径（Medium）](https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/)

```go
```

> 大多数情况下不会自己去实现平衡树或跳表
可以用语言内置的有序集合库
C++: set, multiset, map, multimap
Java: TreeSet, TreeMap
Python: OrderedDict

- 尝试用语言内置的有序集合库，或写一棵平衡树，来解决 [滑动窗口最大值（Hard）](https://leetcode-cn.com/problems/sliding-window-maximum/)

```go
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
```

- [解数独（Hard）](https://leetcode-cn.com/problems/sudoku-solver/)

```go
```

### 迭代加深、折半搜索与双向搜索

- 引入

```go

```

- [单词接龙（Hard）](https://leetcode-cn.com/problems/word-ladder/)

**解法1: 单向BFS**

```go
```

**解法2: 双向BFS**

```go
```

### 启发式搜索：A* 算法

- [滑动谜题（Hard）](https://leetcode-cn.com/problems/sliding-puzzle/)

```go
```

- [八数码（Medium）AcWing](https://www.acwing.com/problem/content/181/)

```go
```

- [八数码（打印方案）（Medium）AcWing](https://www.acwing.com/problem/content/847/)

```go
```