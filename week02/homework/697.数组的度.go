/*
 * @lc app=leetcode.cn id=697 lang=golang
 *
 * [697] 数组的度
 */

// @lc code=start
func findShortestSubArray(nums []int) (ans int) {
	mp := map[int]entry{}
	for i, v := range nums {
		if e, has := mp[v]; has {
			e.cnt++
			e.r = i
			mp[v] = e
		} else {
			mp[v] = entry{1, i, i} // 出现的次数，首次出现的位置，最后出现的位置
		}
	}
	maxCnt := 0
	for _, e := range mp {
		if e.cnt > maxCnt {
			maxCnt, ans = e.cnt, e.r - e.l + 1
		} else if e.cnt == maxCnt {
			ans = min(ans, e.r - e.l + 1)
		}
	}
	return
}

type entry struct {
	cnt, l, r int
}

func min(a,b int) int {
	if a < b { return a}
	return b
}
// @lc code=end

