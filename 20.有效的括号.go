/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	var h map[string]string{
		"{" : "}",
		"(" : ")",
		"[" : "]",
		"?" : "?"
	}

	for _, v := range h {
		fmt.print(v)
	}

	return nil
}
// @lc code=end

