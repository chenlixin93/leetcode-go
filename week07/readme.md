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

- [戳气球（Hard）](https://leetcode-cn.com/problems/burst-balloons/)

```go
```

- [合并石头的最低成本（Hard](https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/)

```go
```

## 树形动态规划

- [打家劫舍 III （Medium）](https://leetcode-cn.com/problems/house-robber-iii/)

```go
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