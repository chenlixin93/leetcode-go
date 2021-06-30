/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n > 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n) // 负负为正
}

func quickMul(x float64, n int) float64 {
	if n == 0 { return 1 }

	y := quickMul(x, n/2)
	if n%2 == 0 { return y * y }
	return y * y * x // n为奇数，n/2 向下取整，所以这里还得乘回1个x
}
// @lc code=end

