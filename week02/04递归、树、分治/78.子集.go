/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) (ans [][]int) {
	// 官方题解，简洁
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		// 考虑选择当前位置
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1] // 恢复现场

		// 考虑不选择当前位置
		dfs(cur + 1)
	}
	dfs(0)
	return
}
// @lc code=end

