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
func climbStairs(n int) int {

    if n <= 2 {
        return n
    }

    f1 := 1
    f2 := 2
    f3 := 3
    for i := 3; i <= n; i++ {
        f3 = f2 + f1
        f1 = f2
        f2 = f3
        // 当n==4时，f4 = f3 + f2，所以提前把f3、f2填入f2、f1
    }
    return f3
}
```

- [三角形最小路径和（Medium）](https://leetcode-cn.com/problems/triangle/description/)

```go
func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    m := len(triangle[n-1]) // 取最后一行的长度
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, m+1)
    }
    // 自底向上
    for i := n-1; i >= 0; i-- {
        for j := 0; j <= i; j++ {
            dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
        }
    }
    return dp[0][0]
}

func min(a,b int) int {
    if a < b {return a}
    return b
}
```

- [最长递增子序列的个数（Medium）](https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/)

```go
```

- [完全平方数（Medium）](https://leetcode-cn.com/problems/perfect-squares/)

```go
func numSquares(n int) int {
    // 完全平方数就是物品（可以无限件使用），凑个正整数n就是背包，问凑满这个背包最少有多少物品？
    // dp[i]：和为i的完全平方数的最少数量为dp[i]
    // dp[i] 可以由dp[i - j * j]推出， dp[i - j * j] + 1 便可以凑成dp[i]
    f := make([]int, n + 1)
    for i := range f {
        f[i] = 1000000000
    }
    f[0] = 0
    for i := 0; i <= n; i++ { // 循环体积(背包)
        for j := 1; j * j <= i; j++ { // 循环物品
            f[i] = min(f[i], f[i - j * j] + 1)
        }
    }
    // 也可以先循环物品，再循环背包
    return f[n]
}

func min(a,b int) int {
    if a < b { return a }
    return b
}
```

- [跳跃游戏（Medium）](https://leetcode-cn.com/problems/jump-game/)

```go
func canJump(nums []int) bool {
    max_i := 0 // 能到达的最远距离
    idx := 0 // 记录最后的下标
    for i := range nums {
        // 如果当前下标 + 数字 > 最远下标，则更新最远下标
        if max_i >= i && i + nums[i] > max_i {
            max_i = i + nums[i]
        }
        idx = i
    }
    return max_i >= idx
}
```
- [跳跃游戏 II （Medium）](https://leetcode-cn.com/problems/jump-game-ii/)

```go
// 贪心
func jump(nums []int) int {
    // 决策包容性：同样是跳1步，从 a 跳到 “能跳得更远位置c”的 b，
    // 未来的可达集合包含了跳到其他b的可达集合，所以这个局部最优决策是正确的
    now := 0
    ans := 0
    right := 0
    //next := 0
    for now < len(nums) - 1 {
        if nums[now] == 0 { return -1 }

        right = now + nums[now]
        if right >= len(nums) - 1 { // 到达最后
            return ans + 1
        }
        // 从now出发可以到[now+1, right]
        next := now + 1
        for i := now + 2; i <= right; i++ { // 搜索[now+1, right]之间能跳到最远c的某个b点
            next_right := i + nums[i]
            if next_right > (next + nums[next]) {
                next = i
            }
        }
        now = next
        ans++
    }
    return ans
}

// DP

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

来自@**Johnny_牧云**的解析
https://leetcode-cn.com/problems/edit-distance/solution/zi-di-xiang-shang-he-zi-ding-xiang-xia-by-powcai-3/

```
(一)、当word1[i]==word2[j]时,由于遍历到了i和j,说明word1的0~i-1和word2的0~j-1的匹配结果已经生成,
由于当前两个字符相同,因此无需做任何操作,dp[i][j]=dp[i-1][j-1]
(二)、当word1[i]!=word2[j]时,可以进行的操作有3个:
      ① 替换操作:可能word1的0~i-1位置与word2的0~j-1位置的字符都相同,
           只是当前位置的字符不匹配,进行替换操作后两者变得相同,
           所以此时dp[i][j]=dp[i-1][j-1]+1(这个加1代表执行替换操作)
      ②删除操作:若此时word1的0~i-1位置与word2的0~j位置已经匹配了,
         此时多出了word1的i位置字符,应把它删除掉,才能使此时word1的0~i(这个i是执行了删除操作后新的i)
         和word2的0~j位置匹配,因此此时dp[i][j]=dp[i-1][j]+1(这个加1代表执行删除操作)
      ③插入操作:若此时word1的0~i位置只是和word2的0~j-1位置匹配,
          此时只需要在原来的i位置后面插入一个和word2的j位置相同的字符使得
          此时的word1的0~i(这个i是执行了插入操作后新的i)和word2的0~j匹配得上,
          所以此时dp[i][j]=dp[i][j-1]+1(这个加1代表执行插入操作)
      ④由于题目所要求的是要最少的操作数:所以当word1[i] != word2[j] 时,
          需要在这三个操作中选取一个最小的值赋格当前的dp[i][j]
(三)总结:状态方程为:
if(word1[i] == word2[j]):
      dp[i][j] = dp[i-1][j-1]
else:
       min(dp[i-1][j-1],dp[i-1][j],dp[i][j-1])+1
```

**代码**
```go
func minDistance(word1 string, word2 string) int {
	// dp[i][j] word1前i个字符转换成word2前j个字符所使用的的最少步数（已达成时）
	n := len(word1)
	m := len(word2)
	dp := make([][]int, n + 1)
	for i := range dp {
		dp[i] = make([]int, m + 1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = i // abc->[]
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j // []->abc
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// i = 1 取word的第1个字符，即word[0] = word[i - 1]
			if word1[i - 1] == word2[j - 1] {
				// 第 i 个 == 第 j 个时，dp[i - 1][j - 1]无需操作
				dp[i][j] = dp[i - 1][j - 1]
			} else {
				// 第 i 个 != 第 j 个时
				insert := dp[i][j - 1] + 1 // 0 ~ i 和 0 ~ j-1 已经相同，要在i后面插入 j 的字符。比如 ab -> abc
				delete := dp[i - 1][j] + 1 // 0 ~ i-1 和 j 已经相同，多余的i位置删掉。比如 abcd -> abc
				replace := dp[i - 1][j - 1] + 1 // 前面相同，替换第i个或者替换第j个即可。比如 abcd -> abce
				dp[i][j] = min(insert, min(delete, replace))
			}
		}
	}
	return dp[n][m]
}

func min(a,b int) int {
	if a < b {return a}
	return b
}
```

### 背包问题

- [分割等和子集（Medium）](https://leetcode-cn.com/problems/partition-equal-subset-sum/)

```go
func canPartition(nums []int) bool {
	// 0/1背包
	// int数组，sum是奇数肯定不行；
	// sum是偶数的话，如果在数据范围内能够达到sum/2，那剩下的肯定也是sum/2
	n := len(nums)
	new_nums := make([]int, n + 1)
	for i := 1; i <= n; i++ { // 挪一位方便计算
		new_nums[i] = nums[i - 1]
	}
	sum := 0
	for i := 1; i <= n; i++ {
		sum += new_nums[i]
	}
	// 奇数不可达
	if sum % 2 == 1 { return false }
	sum = sum >> 1
	f := make([]bool, sum+1)
	f[0] = true // 0体积true
	// 先循环nums(物品)
	// 再循环sum(体积)
	for i := 1; i <= n; i++ {
		for j := sum; j - new_nums[i] >= 0; j-- {
			f[j] = f[j] || f[j - new_nums[i]]
		}
	}
	return f[sum]
}
```

- [零钱兑换 II （Medium）](https://leetcode-cn.com/problems/coin-change-2/)

```go
func change(amount int, coins []int) int {
	// 完全背包模型 + 计数
	f := make([]int, amount + 1) // 0 ~ amount 的方案数
	for i := range f {
		f[i] = 0 // 初始为0种方案	
	}
	f[0] = 1 // 什么都不给也是一种方案
	// 先循环coins(物品)
	// 再循环amount(体积)
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ { // 正序
			f[j] = f[j] + f[j - coins[i]]
		}
	}
	return f[amount]
}
```
