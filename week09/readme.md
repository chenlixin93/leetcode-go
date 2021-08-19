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