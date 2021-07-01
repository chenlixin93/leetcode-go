/*
 * @lc app=leetcode.cn id=1074 lang=golang
 *
 * [1074] 元素和为目标值的子矩阵数量
 */

// @lc code=start
func numSubmatrixSumTarget(matrix [][]int, target int) (ans int) {
	var subarraySum func(nums []int, k int) int
	subarraySum = func(nums []int, k int) (count int) {
		mp := map[int]int{0: 1} // 和为0有一次
		for i, pre := 0, 0; i < len(nums); i++ {
			pre += nums[i]
			if _, ok := mp[pre - k]; ok {
				count += mp[pre - k]
			}

			mp[pre]++
		}
		return
	}

	// 上下边界构成矩形，即二维的数组
	// 二维数组可以通过列压缩成一维组，由子矩形转化为子数组的问题求解
	for i := range matrix {// 枚举上边界
		columnSum := make([]int, len(matrix[0]))
		for _, row := range matrix[i:] { // 枚举下边界
			for c, v := range row {
				columnSum[c] += v // 更新每列的元素和
			}
			ans += subarraySum(columnSum, target)
		}
	}
	return
}
// @lc code=end