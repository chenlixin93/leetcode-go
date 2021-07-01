/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) (ans []string) {
	// L 、R 指左右括号剩下的可用长度
	var dfs func(L,R int, path string)
	dfs = func(L,R int, path string)  {
		if 2*n == len(path) {
			ans = append(ans, path)
			return
		}

		if L > 0 {
			dfs(L - 1, R, path + "(")
		}
		if L < R {
			dfs(L, R - 1, path + ")")
		}
	}

	dfs(n, n, "")
	return
}
// @lc code=end

