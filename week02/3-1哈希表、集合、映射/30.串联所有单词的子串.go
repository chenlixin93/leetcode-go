/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	// 没有单词时，不符合条件
	if (len(words) < 1) { return []int{} }

	word_nums := len(words)
	word_len := len(words[0])
	// 需要用上所有单词，如果字符串长度小于单词总长度，不符合条件
	if len(s) < word_nums * word_len { return []int{} }

	var map_1 = make(map[string]int)
	for _, v := range words {
		map_1[v]++ // 统计单词出现次数
	}

	var res []int
	// [当前下标+单词总长度]不会越界
	for i := 0; i + (word_nums * word_len) <= len(s); i++ {
		var count int
		var map_2 = make(map[string]int)
		for step := 0; step < word_nums; step++ {
			start := i + step * word_len
			word := s[start:start+word_len]
			if nums, found := map_1[word]; found && nums > map_2[word] {
				// 在map_1中能找到当前s截取到的单词，并且之前统计的次数大于map2中的次数
				// 也就是单词组的出现次数能覆盖字符串中的出现次数
				map_2[word]++
				count++
			} else {
				break;
			}
		}
		if count == word_nums {
			res = append(res, i)
		}
	}

	return res
}
// @lc code=end

