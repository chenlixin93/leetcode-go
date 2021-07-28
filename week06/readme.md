- [week06](#week06)
  * [homework](#homework)
  * [动态规划一](#动态规划一)
  * [动态规划二](#动态规划二)
    + [买卖股票系列问题](#买卖股票系列问题)
    + [线性DP问题](#线性dp问题)
    + [背包问题](#背包问题)

# week06

## homework

- [爬楼梯（Easy）](https://leetcode-cn.com/problems/climbing-stairs/description/)

```go
```

- [三角形最小路径和（Medium）](https://leetcode-cn.com/problems/triangle/description/)

```go
```

- [最长递增子序列的个数（Medium）](https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/)

```go
```

- [完全平方数（Medium）](https://leetcode-cn.com/problems/perfect-squares/)

```go
```

- [跳跃游戏（Medium）](https://leetcode-cn.com/problems/jump-game/)

```go
```
- [跳跃游戏 II （Medium）](https://leetcode-cn.com/problems/jump-game-ii/)

```go
```

## 动态规划（一）

- [零钱兑换（Medium）](https://leetcode-cn.com/problems/coin-change/)

```go
func coinChange(coins []int, amount int) int {
	INF := amount + 1 // 最小面值为1块
	opt := make([]int, amount + 1)

	opt[0] = 0
	for i := 1; i <= amount; i++ {
		opt[i] = INF
		for j := 0; j < len(coins); j++ {
			// 最优子结构推导出全局最优
			if i - coins[j] >= 0 {
				opt[i] = min(opt[i], opt[i - coins[j]] + 1)
			}
			
		}
	}
	if opt[amount] >= INF {
		opt[amount] = -1
	}
	return opt[amount]
}

func min(a,b int) int {
	if a < b { return a }
	return b
}
```

- [不同路径 II （Medium）](https://leetcode-cn.com/problems/unique-paths-ii/)

```go
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    n := len(obstacleGrid)
    m := len(obstacleGrid[0])

    f := make([][]int, n)
    for i := 0; i < n; i++ {
        f[i] = make([]int, m)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if obstacleGrid[i][j] == 1 {
                f[i][j] = 0
            } else if i == 0 && j == 0 {
                f[i][j] = 1
            } else if i == 0 {
                f[i][j] = f[i][j - 1]
            } else if j == 0 {
                f[i][j] = f[i - 1][j]
            } else {
                f[i][j] = f[i - 1][j] + f[i][j - 1] // 当前点的路径等于前一步的路径之和
            }
        }
    }
    return f[n - 1][m - 1]
}
```

- [最长公共子序列（Medium）](https://leetcode-cn.com/problems/longest-common-subsequence/)

```go
func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

	dp := make([][]int, m + 1)
	for i := range dp {
		dp[i] = make([]int, n + 1)
	}

	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 { // 两个字符相等时，公共长度+1
				dp[i + 1][j + 1] = dp[i][j] + 1
			} else { // 两个字符不等时，比较选其中一个字符能够达成的公共长度较大的
				dp[i + 1][j + 1] = max(dp[i][j + 1], dp[i + 1][j])
			}
		}
	}
	return dp[m][n]
}

func max(a,b int) int {
	if a > b { return a }
	return b
}
```

- [最长递增子序列（Medium）](https://leetcode-cn.com/problems/longest-increasing-subsequence/)

```go
func lengthOfLIS(nums []int) int {
	m := len(nums)
	// 定义dp[i]表示前i个数字的最大递增长度
	dp := make([]int, m)
	for i := 0; i < m; i++ {
		dp[i] = 1 // 初始长度都是1，子序列只有i本身
	}
	// 子序列不需要连续，在从前到后的搜索中，满足条件即可加1
	for i := 1; i < m; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
	}
	sort.Ints(dp)
	return dp[m - 1]
}

func max(a,b int) int {
	if a > b { return a }
	return b
}
```

- [最大子序和（Easy）](https://leetcode-cn.com/problems/maximum-subarray/)

```go
func maxSubArray(nums []int) int {
	// 定义dp[i]表示当前从下标0到i的最大和
	dp := make([]int, len(nums))
	for k,v := range nums {
		dp[k] = v
	}
	for i := 1; i < len(nums); i++ {
		dp[i] = max(0, dp[i - 1]) + nums[i] // 如果是负数，从0开始
	}
	sort.Ints(dp)
	return dp[len(nums) - 1]
}

func max(a,b int) int {
	if a > b { return a }
	return b
}
```

- [乘积最大子数组（Medium）](https://leetcode-cn.com/problems/maximum-product-subarray/)

```go
```

## 动态规划（二）

### 买卖股票系列问题

- [买卖股票的最佳时机（Easy）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)

```go
```

- [买卖股票的最佳时机 II （Easy）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)

```go
```

- [买卖股票的最佳时机 III （Easy）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/)

```go
```

- [买卖股票的最佳时机 IV （Hard）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/)

```go
```

- [买卖股票的最佳时机含手续费（Medium）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)

```go
```

- [最佳买卖股票时机含冷冻期（Medium）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)

```go
```

### 线性DP问题

- [打家劫舍（Medium）](https://leetcode-cn.com/problems/house-robber/)

```go
```

- [打家劫舍 II- 环形 DP （Medium）](https://leetcode-cn.com/problems/house-robber-ii/)

```go
```

- [编辑距离（重点题）（Hard）](https://leetcode-cn.com/problems/edit-distance/)

```go
```

### 背包问题

- [分割等和子集（Medium）](https://leetcode-cn.com/problems/partition-equal-subset-sum/)

```go
```

- [零钱兑换 II （Medium）](https://leetcode-cn.com/problems/coin-change-2/)

```go
```
