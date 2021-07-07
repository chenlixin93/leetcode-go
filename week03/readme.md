- [week03](#week03)
  * [homework](#homework)
  * [树-二叉树-树的遍历](#树-二叉树-树的遍历)
  * [树的直径-最近公共祖先-树的变形](#树的直径-最近公共祖先-树的变形)
  * [图-图的遍历](#图-图的遍历)
  * [DFS-BFS](#dfs-bfs)

# week03

## homework

- [从中序与后序遍历序列构造二叉树（Medium）]()

```go
```

- [课程表 II （Medium）]()

```go
```

- [被围绕的区域（Medium）]()

```go
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
```

## DFS-BFS

- [电话号码的字母组合（Medium）](lems/letter-combinations-of-a-phone-number/)

```go
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

- [岛屿数量（Medium）]()

```go
```

- [最小基因变化（Medium）]()

```go
```

- [矩阵中的最长递增路径（Hard）]()

```go
```