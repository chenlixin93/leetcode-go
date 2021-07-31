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
func maxProfit(prices []int) int {

	n := len(prices)
	if n == 0 { return 0 }

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0 // 没有股票，也没有利润
	dp[0][1] = -prices[0] // 买入股票，负利润

	for i := 1; i < n; i++ {
		// 当前没有持有股票的最大利润, 可能是前一天没有持有股票，或者前一天持有股票，今天卖出
		dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i])
		// 当前持有股票的最大利润, 可能是前一天持有股票的利润，或者前一天没有股票，今天买入
		// 注意题目只允许买入一次，也就是说之前没有买入也没有卖出过，那么持有的金额一定是初始状态，即买入后利润为 dp[0][0] - prices[i] = - prices[i]
		dp[i][1] = max(dp[i - 1][1], - prices[i])
	}

	return dp[n - 1][0] // 手里持有的现金（利润）达到最大，此时一定不是买入股票（只允许买入一次，还没卖出肯定达不到最大）
}

func max(a,b int) int {
	if a > b { return a }
	return b
}
```

- [买卖股票的最佳时机 II （Easy）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)

```go
func maxProfit(prices []int) int {
	// 注意题目没有限制买入卖出的次数，但是仍要遵守先买入才能卖出
	n := len(prices)
	if n == 0 { return 0 }

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0 // 没有股票，也没有利润
	dp[0][1] = -prices[0] // 买入股票，负利润

	for i := 1; i < n; i++ {
		// 当前没有持有股票的最大利润, 可能是前一天没有持有股票，或者前一天持有股票，今天卖出
		dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i])
		// 当前持有股票的最大利润, 可能是前一天持有股票的利润，或者前一天没有股票，今天买入
		dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] - prices[i])
	}

	return max(dp[n - 1][0], dp[n - 1][1])
}

func max(a,b int) int {
	if a > b { return a }
	return b
}
```

- [买卖股票的最佳时机 III （Easy）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/)

```go
func maxProfit(prices []int) int {
    // 注意题目限制完整交易的次数为2，买入前必须卖掉之前的股票
    n := len(prices)
    if n == 0 { return 0 }
    // 加一位，方便后续计算 i-1
    new_prices := make([]int, n + 1)
    new_prices[0] = 0
    for i := 1; i <= n; i++ {
        new_prices[i] = prices[i-1]
    }

    // 定义dp[i][j][k]
    dp := make([][][]int, n + 1)
    c := 2 // 最多两笔交易
    for i := range dp {
        dp[i] = make([][]int, 2)
        for j := range dp[i] {
            dp[i][j] = make([]int, c + 1)
            for k := range dp[i][j] {
                dp[i][j][k] = -1000000000 // 负无穷
            }
        }
    }

    dp[0][0][0] = 0 // 第0天的状态没有持有股票、也没有交易次数
    for i := 1; i <= n; i++ {
        for j := 0; j <= 1; j++ {
            for k := 0; k <= c; k++ {
                dp[i][j][k] = dp[i-1][j][k]
                if j == 0 { 
                    // 当前不持有股票 = 前一天没有股票或者前一天持有股票，今天卖出
                    dp[i][0][k] = max(dp[i][0][k], dp[i-1][1][k] + new_prices[i])
                }
                if j == 1 && k > 0 { 
                    // 当前持有股票 = 前一天已有股票或者前一天没有股票，今天买入
                    //（交易次数只在买入时发生变化，从 k-1 转移到当前 k 次）
                    dp[i][1][k] = max(dp[i][1][k], dp[i-1][0][k-1] - new_prices[i])
                }
            }
        }
    }

    ans := 0 
    for k := 0; k <= c; k++ {
        ans = max(ans, dp[n][0][k])
    }
    return ans
}

func max(a,b int) int {
    if a > b {return a}
    return b
}
```

- [买卖股票的最佳时机 IV （Hard）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/)

```go
func maxProfit(k int, prices []int) int {
    // 注意题目限制完整交易的次数为k，买入前必须卖掉之前的股票
    n := len(prices)
    if n == 0 { return 0 }
    // 加一位，方便后续计算 i-1
    new_prices := make([]int, n + 1)
    new_prices[0] = 0
    for i := 1; i <= n; i++ {
        new_prices[i] = prices[i-1]
    }

    // 定义dp[i][j][k]
    dp := make([][][]int, n + 1)
    c := k // 最多k笔交易
    for i := range dp {
        dp[i] = make([][]int, 2)
        for j := range dp[i] {
            dp[i][j] = make([]int, c + 1)
            for k := range dp[i][j] {
                dp[i][j][k] = -1000000000 // 负无穷
            }
        }
    }

    dp[0][0][0] = 0 // 第0天的状态没有持有股票、也没有交易次数
    for i := 1; i <= n; i++ {
        for j := 0; j <= 1; j++ {
            for k := 0; k <= c; k++ {
                dp[i][j][k] = dp[i-1][j][k]
                if j == 0 { 
                    // 当前不持有股票 = 前一天没有股票或者前一天持有股票，今天卖出
                    dp[i][0][k] = max(dp[i][0][k], dp[i-1][1][k] + new_prices[i])
                }
                if j == 1 && k > 0 { 
                    // 当前持有股票 = 前一天已有股票或者前一天没有股票，今天买入
                    //（交易次数只在买入时发生变化，从 k-1 转移到当前 k 次）
                    dp[i][1][k] = max(dp[i][1][k], dp[i-1][0][k-1] - new_prices[i])
                }
            }
        }
    }

    ans := 0 
    for k := 0; k <= c; k++ {
        ans = max(ans, dp[n][0][k])
    }
    return ans
}

func max(a,b int) int {
    if a > b {return a}
    return b
}
```

- [买卖股票的最佳时机含手续费（Medium）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)

```go
func maxProfit(prices []int, fee int) int {
    // 注意题目不限制交易次数，套用买卖股票2的模版即可，买入时减去fee
	n := len(prices)
	if n == 0 { return 0 }

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0 // 没有股票，也没有利润
	dp[0][1] = - prices[0] - fee // 买入股票，负利润

	for i := 1; i < n; i++ {
		// 当前没有持有股票的最大利润, 可能是前一天没有持有股票，或者前一天持有股票，今天卖出
		dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i])
		// 当前持有股票的最大利润, 可能是前一天持有股票的利润，或者前一天没有股票，今天买入
		dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] - prices[i] - fee)
	}

	return max(dp[n - 1][0], dp[n - 1][1])
}

func max(a,b int) int {
	if a > b { return a }
	return b
}
```

- [最佳买卖股票时机含冷冻期（Medium）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)

```go
func maxProfit(prices []int) int {
    // 注意题目不限制交易次数，买入前必须卖掉之前的股票
	// 与股票三、四对比，同样是三个状态，但是决策上要根据冷冻期进行调整
    n := len(prices)
    if n == 0 { return 0 }
    // 加一位，方便后续计算 i-1
    new_prices := make([]int, n + 1)
    new_prices[0] = 0
    for i := 1; i <= n; i++ {
        new_prices[i] = prices[i-1]
    }

    // 定义dp[i][j][k]
    dp := make([][][]int, n + 1)
    for i := range dp {
        dp[i] = make([][]int, 2)
        for j := range dp[i] {
            dp[i][j] = make([]int, 2)
            for k := range dp[i][j] {
                dp[i][j][k] = -1000000000 // 负无穷
            }
        }
    }

    dp[0][0][0] = 0 // 边界是0，其他不合法或者影响决策的初始为负无穷
    for i := 1; i <= n; i++ { // 第几天
        for j := 0; j <= 1; j++ { // 是否持有
            for l := 0; l <= 1; l++ { // 冷冻期
				if j == 0 && l == 0 {
					// 当前没有股票同时没有冷冻期 = 前一天不操作或者前一天没有股票但处于冷冻期（前一天持有股票不可能处于冷冻期，不考虑）
					dp[i][0][0] = max(dp[i-1][0][0], dp[i-1][0][1])
				}
				if j == 1 && l == 0 {
					// 当前持有股票同时没有冷冻期 = 前一天不操作或者前一天没有股票时，买入股票，转移到当前状态。
					dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][0][0] - prices[i-1])
				}
				if j == 0 && l == 1 {
					// 当前处于冷冻期，说明前一天持有股票，并且卖出
					dp[i][0][1] = dp[i - 1][1][0] + prices[i - 1]
				}
				if j == 1 && l == 1 {
					// dp[i][1][1]; // 当前处于冷冻期，同时持有股票（不可能，不考虑）
				}
			}
		}
	}

    ans := 0 
    for k := 0; k <= 1; k++ {
        ans = max(ans, dp[n][0][k])
    }
    return ans
}

func max(a,b int) int {
    if a > b {return a}
    return b
}
```

### 线性DP问题

- [打家劫舍（Medium）](https://leetcode-cn.com/problems/house-robber/)

```go
func rob(nums []int) int {

	n := len(nums)
	new_nums := make([]int, n + 1)
	new_nums[0] = 0
	for i := range nums {
		new_nums[i+1] = nums[i]
	}

	dp := make([][2]int, n + 1)
	for i := range dp {
		dp[i] = [2]int{0, 0}
	}

	// dp[i][j] j 表示第i天是否偷盗
	dp[0][0] = 0
	dp[0][1] = math.MinInt32  // -2147483648 // 一般写-1e9 即-1000000000
	for i := 1; i <= n; i++ {
		// for j := 0; j <= 1; j++ { // j == 1 处于冷冻期
			// 当前可偷 = 前一天未偷或者前一天偷盗
			dp[i][0] = max(dp[i - 1][0], dp[i - 1][1])
			// 当前不可偷 = 前一天未偷今天偷盗，转移状态到当前
			dp[i][1] = dp[i-1][0] + new_nums[i]
		// }
	}
	return max(dp[n][0], dp[n][1])
}

func max(a,b int) int {
	if a > b { return a }
	return b
}
```

- [打家劫舍 II- 环形 DP （Medium）](https://leetcode-cn.com/problems/house-robber-ii/)

```go
func rob(nums []int) int {
	// 解题思路：环形DP
	// 对比打家劫舍1，第一间屋子可偷可不偷，搜两遍比较大小
	n := len(nums)
	if n == 1 { return nums[0] }

	new_nums := make([]int, n + 1)
	new_nums[0] = 0 // 前置0方便计算
	for i := range nums {
		new_nums[i+1] = nums[i]
	}

	dp := make([][2]int, n + 1)
	for i := range dp {
		dp[i] = [2]int{0, 0}
	}

	// dp[i][j] j 表示第i天是否偷了
	// 第一种情况：不偷1可偷n
	dp[1][0] = 0 
	dp[1][1] = math.MinInt32  // 第一间偷了没用，不合法（不可达状态）
	for i := 2; i <= n; i++ {
		// 当前可偷 = 前一天未偷或者前一天偷盗
		dp[i][0] = max(dp[i - 1][0], dp[i - 1][1])
		// 当前不可偷 = 前一天未偷今天偷盗，转移状态到当前
		dp[i][1] = dp[i-1][0] + new_nums[i]
	}
	ans := max(dp[n][0], dp[n][1]) // 最后一间可偷可不偷
	// 第二种情况：不偷n可偷1
	dp[1][0] = 0 
	dp[1][1] = new_nums[1]  // 偷第一间的收益，合法（可达状态）
	for i := 2; i <= n; i++ {
		// 当前可偷 = 前一天未偷或者前一天偷盗
		dp[i][0] = max(dp[i - 1][0], dp[i - 1][1])
		// 当前不可偷 = 前一天未偷今天偷盗，转移状态到当前
		dp[i][1] = dp[i-1][0] + new_nums[i]
	}
	return max(ans, dp[n][0])
}

func max(a,b int) int {
	if a > b { return a }
	return b
}
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
