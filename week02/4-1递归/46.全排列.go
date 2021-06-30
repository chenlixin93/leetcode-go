/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) (ans [][]int) {
	// 来自笨猪爆破组，返回值稍加改动，有官方题解那味儿了
	visited := map[int]bool{}

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}

		for _, v := range nums {
			if visited[v] {
				continue
			}
			path = append(path, v)
			visited[v] = true
			dfs(path)
			path = path[:len(path) - 1]
			visited[v] = false
		}
	}
	dfs([]int{})
	return
}
// @lc code=end

