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
```

- [最长公共子序列（Medium）](https://leetcode-cn.com/problems/longest-common-subsequence/)

```go
```

- [最长递增子序列（Medium）](https://leetcode-cn.com/problems/longest-increasing-subsequence/)

```go
```

- [最大子序和（Easy）](https://leetcode-cn.com/problems/maximum-subarray/)

```go
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
