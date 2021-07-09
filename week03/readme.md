- [week03](#week03)
  * [homework](#homework)
  * [树-二叉树-树的遍历](#树-二叉树-树的遍历)
  * [树的直径-最近公共祖先-树的变形](#树的直径-最近公共祖先-树的变形)
  * [图-图的遍历](#图-图的遍历)
  * [DFS-BFS](#dfs-bfs)

# week03

## homework

- [从中序与后序遍历序列构造二叉树（Medium）](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)

```go
// 递归
func buildTree(inorder []int, postorder []int) *TreeNode {
	idxMap := map[int]int{}
	for i, v := range inorder {
		idxMap[v] = i
	}

	var build func(int, int) *TreeNode
	build = func(inorderLeft, inorderRight int) *TreeNode {
		if inorderLeft > inorderRight {
			return nil // 左右索引相交了
		}

		val := postorder[len(postorder) - 1] // 后序遍历的末位元素是当前子树的根节点
		postorder = postorder[:len(postorder) - 1]
		root := &TreeNode{Val: val}

		inorderRootIndex := idxMap[val]
		root.Right = build(inorderRootIndex + 1, inorderRight)
		root.Left = build(inorderLeft, inorderRootIndex - 1)
		return root
	}
	return build(0, len(inorder) - 1) // 考虑子问题怎么处理即可
}
```

- [课程表 II （Medium）](https://leetcode-cn.com/problems/course-schedule-ii/)

```go
// BFS拓扑排序模板
func findOrder(numCourses int, prerequisites [][]int) []int {
	n := numCourses
	edges := make([][]int, n) // [[],[]]
	inDeg := make([]int, n) // [0,0]

	var addEdge func(x,y int)
	addEdge = func(x,y int) {
		edges[x] = append(edges[x], y)
		inDeg[y]++
	}
	// 建图，加边
	for _, pre := range prerequisites {
		ai := pre[0]
		bi := pre[1] // 先修课程
		addEdge(bi, ai) // 即ai的入度数+1
	}

	var topsort func(q []int) []int
	topsort = func(q []int) []int {
		res := []int{} // 存储已学习的课程
		m := 0
		for len(q) != 0 {
			x := q[0] // 取出可学习的课程
			q = q[1:] // 注意不能用 :=
			for _, y := range edges[x] {
				inDeg[y]-- // y课程入度数-1
				if inDeg[y] == 0 {
					q = append(q, y)
				}
			}
			res = append(res, x)
			m++
		}
		if m == n { return res } // 可以完成所有课程，返回学习顺序
		return []int{} // 返回空数组
	}

	q := []int{} // 存放入度数为0的课程编号，即可学习课程
	for i := 0; i < n; i++ {
		if inDeg[i] == 0 { 
			q = append(q, i)
		}
	}
	return topsort(q)
}
```

- [被围绕的区域（Medium）](https://leetcode-cn.com/problems/surrounded-regions/)

```go
// DFS
func solve(board [][]byte)  {
	/*
	* 解题思路
	* 假设有一场洪水，初始时从边界上的 O 开始所搜，相邻为 O 的，设为洪水不可到达 U
	* 最终遍历整个区域，遇到仍然为 O 的，将其淹没为 X；遇到 U 的，将其恢复为 O
	*/
	m := len(board)
	n := len(board[0])

	if m == 0 || n == 0 { return }

	var dfs = func(x,y int) {}
	dfs = func(x,y int) {
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] == 'U' || board[x][y] == 'X' { return }
		board[x][y] = 'U' // 设为不可达
		
		// 四个方向
		// {1, -1, 0, 0}
		// {0, 0, 1, -1}
		dfs(x + 1, y)
		dfs(x - 1, y)
		dfs(x, y + 1)
		dfs(x, y - 1)
	}

	// 处理第一行和最后一行
	for i := 0; i < n; i++ {
		if board[0][i] == 'O' { dfs(0, i) }
		if board[m - 1][i] == 'O' { dfs(m - 1, i) }
	}
	// 处理第一列和最后一列
	for j := 0; j < m; j++ {
		if board[j][0] == 'O' { dfs(j, 0) }
		if board[j][n - 1] == 'O' { dfs(j, n - 1) }
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' { board[i][j] = 'X'}
			if board[i][j] == 'U' { board[i][j] = 'O'}
		}
	}
}
```

## 树-二叉树-树的遍历

- [二叉树的中序遍历（Easy）](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) (res []int) {
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil { return }
        dfs(root.Left)
        res = append(res, root.Val)
        dfs(root.Right)
    }
    dfs(root)
    return
}
```

- [N 叉树的前序遍历（Easy）](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/)

```go
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) (res []int) {
    var dfs func(root *Node)
    dfs = func(root *Node) {
        if root == nil { return }
        res = append(res, root.Val)
        for _, v := range root.Children {
            if v != nil {
                dfs(v)
            }
        }
    }
    dfs(root)
    return
}
```

- [N 叉树的层序遍历（Medium）](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/)

```go
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) (res [][]int) {
    var dfs func(root *Node, level int) 
    dfs = func(root *Node, level int) {
        if root == nil { return }
        if len(res) == level {
            res = append(res, []int{})
        }
        res[level] = append(res[level], root.Val)
        for _,v := range root.Children {
            dfs(v, level + 1)
        }
    }
    dfs(root, 0)
    return
}
```

- [二叉树的序列化与反序列化（Hard）](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/)

```go
```

- [从前序与中序遍历序列构造二叉树（Medium）](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

```go
// 递归
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    var build func(preorder []int, l1 int, r1 int, 
                inorder []int, l2 int, r2 int) *TreeNode
    build = func(preorder []int, l1 int, r1 int, 
                inorder []int, l2 int, r2 int) *TreeNode { 
        if l1 > r1 { return nil }      
        root := &TreeNode{Val: preorder[l1]}

        inRootIdx := l2
        for inorder[inRootIdx] != preorder[l1] {
            inRootIdx++
        }
        leftSize := inRootIdx - l2
        root.Left = build(preorder, l1 + 1, l1 + leftSize,
                        inorder, l2, inRootIdx - 1)
        root.Right = build(preorder, l1 + leftSize + 1, r1,
                        inorder, inRootIdx + 1, r2)
        return root
    }
    return build(preorder, 0, len(preorder) - 1,
        inorder, 0, len(inorder) - 1)
}
```

## 树的直径-最近公共祖先-树的变形

- [树的直径（此题为 LeetCode 会员题选做）]()

```go

```

- [二叉树的最近公共祖先（Medium）](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    parent := map[int]*TreeNode{} // int => *TreeNode{}
    visited := map[int]bool{} // int => bool

    var dfs func(*TreeNode)
    dfs = func(r *TreeNode) {
        if r == nil {
            return
        }
        if r.Left != nil {
            parent[r.Left.Val] = r
            dfs(r.Left)
        }
        if r.Right != nil {
            parent[r.Right.Val] = r
            dfs(r.Right)
        }
    }
    dfs(root)

    for p != nil {
        visited[p.Val] = true
        p = parent[p.Val]// 从 p 走到根节点
    }
    for q != nil {
        if visited[q.Val] { // 如果q是访问过的，即是p、q的最小公共祖先
            return q
        }
        q = parent[q.Val] // 将q的父节点赋给q，继续下一次查找
    }
    return nil
}
```

## 图-图的遍历

- [冗余连接（Medium）](https://leetcode-cn.com/problems/redundant-connection/)

```go
// DFS找环法
func findRedundantConnection(input [][]int) (ans []int) {
    var max func(x,y int) int
    max = func(x,y int) int {
        if x > y { return x }
        return y
    }
    n := 0
    for _,edge := range input { // 出现过最大的点就是n
        u := edge[0]
        v := edge[1]
        n = max(u, n)
        n = max(v, n)
    }

    edges := map[int][]int{} // int => []int{}
    visited := map[int]bool{}
    for i := 0; i <= n; i++ {
        edges[i] = []int{}
        visited[i] = false
    }
    hasCycle := false

    var dfs func(x, father int)
    dfs = func(x, father int) {
        visited[x] = true
        for _,y := range edges[x] {
            if y == father { continue } // s => a, a => s // 避免无向边往回走
            if visited[y] { // 成环条件，再次走到之前访问过的点
                hasCycle = true 
            } else {
                dfs(y, x)
            }
        }
    }

    for _,edge := range input {
        u := edge[0]
        v := edge[1]

        edges[u] = append(edges[u], v)
        edges[v] = append(edges[v], u)

        // 每加一条边，检查是否出现环
        for i := 0; i <= n; i++ {
            visited[i] = false
        }
        dfs(u, -1)
        if hasCycle {return edge}
    }
    return nil
}
```

- [课程表（Medium）](https://leetcode-cn.com/problems/course-schedule/)

```go
// BFS拓扑排序模板
func canFinish(numCourses int, prerequisites [][]int) bool {
	n := numCourses
	edges := make([][]int, n) // [[],[]]
	inDeg := make([]int, n) // [0,0]

	var addEdge func(x,y int)
	addEdge = func(x,y int) {
		edges[x] = append(edges[x], y)
		inDeg[y]++
	}
	// 建图，加边
	for _, pre := range prerequisites {
		ai := pre[0]
		bi := pre[1] // 先修课程
		addEdge(bi, ai) // 即ai的入度数+1
	}

	var topsort func(q []int) bool
	topsort = func(q []int) bool {
		m := 0 // 统计已学习课程数 
		for len(q) != 0 {
			x := q[0] // 取出可学习的课程
			m++
			q = q[1:] // 注意不能用 :=
			for _, y := range edges[x] {
				inDeg[y]-- // y课程入度数-1
				if inDeg[y] == 0 {
					q = append(q, y)
				}
			}
			
		}
		return m == n
	}

	q := []int{} // 存放入度数为0的课程编号，即可学习课程
	for i := 0; i < n; i++ {
		if inDeg[i] == 0 { 
			q = append(q, i)
		}
	}
	return topsort(q)
}
```

## DFS-BFS

- [电话号码的字母组合（Medium）](lems/letter-combinations-of-a-phone-number/)

```go
// DFS
func letterCombinations(digits string) (ans []string) {
    if len(digits) == 0 { return }
    edges := map[string]string{}
    edges["2"] = "abc"
    edges["3"] = "def"
    edges["4"] = "ghi"
    edges["5"] = "jkl"
    edges["6"] = "mno"
    edges["7"] = "pqrs"
    edges["8"] = "tuv"
    edges["9"] = "wxyz"
    //fmt.Println(edges)

    var dfs func(digits string, index int, s string)
    dfs = func(digits string, index int, s string) {
        if index == len(digits) { 
            ans = append(ans, s)
            return
        }
        for _, ch := range edges[string(digits[index])] {
            s := fmt.Sprintf("%s%s", s, string(ch))
            dfs(digits, index + 1, s)
            s = s[0:len(s) - 1]
        }
    }
    dfs(digits, 0, "")
    return
}
```

- [N 皇后（Hard）](https://leetcode-cn.com/problems/n-queens/)

```go
// DFS，排列模板 + 合法性检查
func solveNQueens(n int) (res [][]string) {
    ans := [][]int{}

    var find func(row int, s []int, used,mapusedIplusJ,usedIminusJ map[int]bool)
    find = func(row int, s []int, used,usedIplusJ,usedIminusJ map[int]bool) {
        if row == n {
            tmp := make([]int, n)
            copy(tmp, s)
            ans = append(ans, tmp)
            return
        }
        for col := 0; col < n; col++ {
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
    //fmt.Println(ans)

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

- [岛屿数量（Medium）](https://leetcode-cn.com/problems/number-of-islands/)

```go
// DFS做法
func numIslands(grid [][]byte) (ans int) {
	m := len(grid)
	n := len(grid[0])

	visited := map[int][]bool{}
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	// 上-下-左-右四个方向
	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}

	var dfs func(grid [][]byte, x,y int)
	dfs = func(grid [][]byte, x,y int) {
		visited[x][y] = true
		for i := 0; i < 4; i++ {
			nx := x + dx[i]
			ny := y + dy[i]

			// 检查边界
			if nx < 0 || ny < 0 || nx >= m || ny >= n { continue }
			if grid[nx][ny] == '1' && !visited[nx][ny] {
				dfs(grid, nx, ny)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				dfs(grid, i, j)
				ans++
			}
		}
	}
	return
}
```

- [最小基因变化（Medium）](https://leetcode-cn.com/problems/minimum-genetic-mutation/)

```go
// BFS
func minMutation(start string, end string, bank []string) int {
	var mutationMap = map[uint8][3]string{
		'A': [...]string{"T", "G", "C"},
		'C': [...]string{"T", "G", "A"},
		'T': [...]string{"A", "G", "C"},
		'G': [...]string{"T", "A", "C"},
	}
	var idxOf func(string, []string) int
	idxOf = func(str string, bank []string) int {
		for i, s := range bank {
			if s == str {
				return i
			}
		}
		return -1
	}

	if (idxOf(end, bank) == -1) { return -1 } // 目标string不在bank里，不需要搜索

	isUsed := make([]bool, len(bank))

	queue := []string{start} // 放入队列
	count := 0
	for len(queue) > 0 { // 队列不为空时
		l := len(queue)
		for i := 0; i < l; i++ {
			curr := queue[i]
			if curr == end { return count }
			for j := 0; j < len(curr); j++ {
				for _, s := range mutationMap[curr[j]] {
					if idx := idxOf(curr[:j] + s + curr[j+1:], bank); // curr[:j] 前闭后开
                        idx != -1 && !isUsed[idx] {
						queue = append(queue, bank[idx])
						isUsed[idx] = true
					}
				}
			}
		}
		count++ // 没找到目标，但是找到一个存在于bank的结果
		queue = queue[l:] // 搜索过的部分截断，只取新进的部分作为队列
	}
	return -1
}
```

- [矩阵中的最长递增路径（Hard）]()

```go
```