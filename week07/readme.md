# week07

## homework

- [冗余连接（Medium）](https://leetcode-cn.com/problems/redundant-connection/)

```go
```

- [岛屿数量（Medium）](https://leetcode-cn.com/problems/number-of-islands/)

```go
```
要求：使用并查集而非 DFS/BFS 实现

## 动态规划的优化

- 引入

给定 n 个二元组(x1,y1),(x2,y2),...(xn,yn), 已经按照 x 从小到大排好序了

求yi + yj + |xi - xj|的最大值（i != j）

```go
// 朴素O(n^2)
for (int i = 1; i <= n; i++) {
    for int j = 1; j <= n; j++) {
        if (i != j) {
            ans = max(ans, y[i] + y[j] + abs(x[i] - x[j]))
        }
    }
}

// 第一步优化
// 式子的值与i，j的顺序无关，不妨设j < i
// 计算量少了一半，可惜还是O(n^2)
// xi,xj 大小关系已知，绝对值也可以拆了
for (int i = 2; i <= n; i++) {
    for int j = 1; j < i; j++) {
        ans = max(ans, y[i] + y[j] + x[i] - x[j])
    }
}

// 第二步优化
// y[i] + x[i] 并不随着 j 而变化，可以提出来在外边算
// 减少了一些加法的次数，还是O(n^2)
for (int i = 2; i <= n; i++) {
    int temp = -1000000000;
    for int j = 1; j < i; j++) {
        temp = max(temp, y[j] - x[j])
    }
    ans = max(ans, y[i] + x[i] + temp)
}

// 第三步优化
// i=4时，temp = max(y1-x1,y2-x2,y3-x3)
// i=5时，temp = max(temp, y4-x4)
int temp;
for (int i = 2; i <= n; i++) {
    temp = max(temp, y[i-1] - x[i-1])
    ans = max(ans, y[i] + x[i] + temp)
}
```

总结：
在动态规划中经常遇到类似的式子，i 是状态变量，j 是决策变量
分离状态变量和决策变量。当循环多余两重时，关注最里面的两重循环，把外层看作定值。
对于一个状态变量，决策变量的取值范围称为“决策候选集合”，观察这个集合随着状态变量的变化情况

一旦发现冗余，或者有更高效维护候选“候选集合”的数据结构，就可以省去一层循环扫描！

- [满足不等式的最大值（Hard）](https://leetcode-cn.com/problems/max-value-of-equation/)

```go
func findMaxValueOfEquation(points [][]int, k int) int {
    // 结合二元组不等式的推导优化
    // 还是设j < i,多了 x[i] - x[j] < k 的条件
    // 也就是j和i离得不能太远
    // 当i增大时，j的取值范围上下界同时增大，要维护y[j] - x[j]的max

    // 本质是，求滑动窗口最大值
    ans := -1000000000
    var q Deque
    for i := 0; i < len(points); i++ {
        // 求上界j <= i-1,下界x[j] >= x[i] - k
        // 在这个范围中 y[j]-x[j]的最大值
        // 考虑两个选项 j1 < j2
        // 写出j1 比 j2 优的条件
        // y[j1] - x[j1] > y[j2] - x[j2]
        // 1. 队头合法性
        // x[j]: points[q.Front()][0]
        for q.Len() != 0 && points[q.Front().(int)][0] < points[i][0] - k {
            q.PopFront()
        }
        // 2. 取队头为最优解
        // y[i]: points[i][1]
        // x[i]: points[i][0]
        // y[i] + x[i] + max{y[j] - x[j]}
        if q.Len() > 0 {
            ans = max(ans, points[i][1] + points[i][0] + points[q.Front().(int)][1] - points[q.Front().(int)][0])
        }
        // 3. 维护队列单调性，队尾插入新选项i
        // 【核心】j1 < j2 , 而y[j1] - x[j1] > y[j2] - x[j2]。即队尾比下一个进来的优，就可以保证单调性。也能认为当前队头是最优解。
        for q.Len() > 0 && points[q.Back().(int)][1] - points[q.Back().(int)][0] <= points[i][1] - points[i][0] {
            q.PopBack()
        }
        q.PushBack(i)
    }
    return ans
}

func max(a,b int) int {
    if a > b {return a}
    return b
}

// 引入deque实现 https://github.com/gammazero/deque/blob/master/deque.go
```

- [环形子数组的最大和（Medium）](https://leetcode-cn.com/problems/maximum-sum-circular-subarray/)

```go
```

## 区间动态规划

- 引入

```go
// 区间动态规划的子问题是基于一个区间的
// 区间长度作为DP的阶段，区间端点作为DP的状态
// 在计算区间长度为len的子问题是，要先计算好所有长度<len的子问题
```

- [戳气球（Hard）](https://leetcode-cn.com/problems/burst-balloons/)

```go
// 思路一：先戳哪个气球？
// 戳完p以后，子问题[l, p-1]和[p+1, r]两端
// 相邻的气球发生了变化！
// 它们和[l, r]不再是同类子问题！

// 思路二：最后一个戳的是哪个气球？
// 先戳完[l, p-1]和[p+1, r],最后戳p
// 子问题两端相邻的气球不变，只是区间点是变化信息
// 满足同类子问题！
func maxCoins(nums []int) int {
    // f[l, r]表示戳破区间l～r之间的所有气球，所获硬币的最大数量
    // 决策：最后一个戳的是p
    // f[l,r] = max(f[l, p-1] + f[p+1, r] + nums[p]*nums[l-1]*nums[r+1])
    // 初值，当l>r时，f[l][r] = 0(不合法)
    // 目标：f[1,n]
    n := len(nums)
    nums = append(append([]int{1}, nums...), 1) // 前后插入1、1
    f := make([][]int, n + 2)
    for i := range f {
        f[i] = make([]int, n + 2)
    }
    // 区间DP
    // 先枚举区间长度
    for len := 1; len <= n; len++ {
        // 在长度范围内
        for l := 1; l <= n - len + 1; l++ {
            r := l + len - 1 // 闭区间 len = r - l + 1
            // 搜索左端点到右端点的所有情况
            for p := l; p <= r; p++ {
                f[l][r] = max(f[l][r], f[l][p-1] + f[p+1][r] + nums[l-1] * nums[p] * nums[r+1])
            }
        }
    }
    return f[1][n]
}

func max(a,b int) int {
    if a > b {return a}
    return b
}
```

- [合并石头的最低成本（Hard](https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/)

```go
// 思路：
// f[l,r] 表示l～r合成一堆的最低成本？
// 不行，l~r不一定要合成一堆，可能会合成若干堆，然后跟其他部分一起凑齐k堆，再合成一堆
// 如何表示“l～r合成若干堆”这个子问题？信息不够，往状态里加

// f[l,r,i]表示把l～r合并成i堆的最低成本
// 决策一：恰好凑成k堆，合成一堆
// f[l,r,1] = f[l,r,k] + sum[l,r](sum[r]-sum[l-1])
// 决策二：分成两个子问题，l～p合成 j 堆，p+1～r合成 i-j 堆，一共 i 堆
// f[l,r,i] = min(f[l][p][j], f[p+1][r][i-j]) 其中 i > 1
// 时间复杂度n^3*k^2

// 决策二可以优化，不需要枚举j，考虑第一堆是哪一段就行了
// f[l,r,i] = min{f[l,p,1] + f[p+1,r,i-1]} 其中 i > 1
// 时间复杂度n^3*k
```

```go
func mergeStones(stones []int, k int) int {
    n := len(stones)
    sum := make([]int, n)
    sum[0] = stones[0]
    // 计算区间和
    for i := 1; i < n; i++ {
        // 0 ~ i 的和
        sum[i] = sum[i - 1] + stones[i]
    }

    f := make([][][]int, n+1)
    for i := range f {
        f[i] = make([][]int, n+1)
        for j := range f[i] {
            f[i][j] = make([]int, k+1)
            for l := range f[i][j] {
                f[i][j][l] = 1e9 // 初始为极大值，min时代表不合法的情况
            }
        }
    }
    for i := 0; i < n; i++ {
        f[i][i][1] = 0 //  同一堆合成1堆成本为0
    }
    cur_sum := 0
    // 枚举区间长度
    for len := 2; len <= n; len++ {
        // 枚举左端点
        for l := 0; l <= n - len; l++ {
            // 计算右端点
            r := l + len - 1
            // 先计算合成 i > 1堆的最优解，最终推导合成i=1堆的最优解
            for i := 2; i <= k; i++ {
                for p := l; p < r; p++ {
                    f[l][r][i] = min(f[l][r][i], f[l][p][1] + f[p+1][r][i-1])
                }
            }
            if l > 0 {
                cur_sum = sum[r] - sum[l-1]
            } else {
                cur_sum = sum[r] - 0
            }
            f[l][r][1] = min(f[l][r][1], f[l][r][k] + cur_sum)
        }
    } 
    if f[0][n-1][1] >= 1e9 {return -1}
    return f[0][n-1][1]
}

func min(a,b int) int {
    if a < b {return a}
    return b
}
```

## 树形动态规划

- 引入

```go
// 树形DP与线性DP没有本质区别
// 其实只是套在深度优先遍历里的动态规划（在DFS的过程中实现DP）
// 子问题就是一棵子树，状态一般表示为“以x为根的子树”，决策是“x的子结点”

// 复杂的题目可以在此基础上增加更多与题目相关的状态、决策
```

- [打家劫舍 III （Medium）](https://leetcode-cn.com/problems/house-robber-iii/)

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
    dp := make(map[*TreeNode][]int)
    return dfs(root, dp)
}

// 定义
// f[x,0]表示以x为根的子树，在不打劫x的情况下，能够盗取的最高金额
// f[x,1]表示以x为根的子树，在打劫x的情况下，能够盗取的最高金额
// 决策
// f[x,0] = max(f[y,0],f[y,1]) // 打劫或者不打劫子树能够获得的最优解
// f[x,1] = val(x) + max(f[y,0]) // 打劫x同时不打劫子树能够获得的最优解
// 目标 max(f[root,0],f[root,1])
func dfs(root *TreeNode, dp map[*TreeNode][]int) int {
    if root == nil {return 0}
    // 初始化
    dp[root] = []int{0, root.Val}
    // 子树不为空的，要累加其结果
    if root.Left != nil {
        dp[root][0] = dp[root][0] + dfs(root.Left, dp)
        dp[root][1] = dp[root][1] + dp[root.Left][0]
    }
    if root.Right != nil {
        dp[root][0] = dp[root][0] + dfs(root.Right, dp)
        dp[root][1] = dp[root][1] + dp[root.Right][0]
    }
    return max(dp[root][0], dp[root][1])
}

func max(a,b int) int {
    if a > b {return a}
    return b
}
```

## 字典树

- [实现 Trie (前缀树) （Medium）](https://leetcode-cn.com/problems/implement-trie-prefix-tree/)

```go
```

- [单词搜索 II （Hard）](https://leetcode-cn.com/problems/word-search-ii/)

```go
```

## 并查集

- [省份数量（Medium）](https://leetcode-cn.com/problems/number-of-provinces/)

```go
```

- [被围绕的区域（Medium）](https://leetcode-cn.com/problems/surrounded-regions/)

```go
```