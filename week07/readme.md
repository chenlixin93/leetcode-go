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

- [满足不等式的最大值（Hard）](https://leetcode-cn.com/problems/max-value-of-equation/)

```go
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