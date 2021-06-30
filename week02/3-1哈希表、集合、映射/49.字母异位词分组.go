/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
    mp := map[string][]string{}

    for _, str := range strs {
        s := []byte(str) // 转成字节数组
        //fmt.Println(s)
        sort.Slice(s, func(i, j int) bool {return s[i] < s[j]}) // map-reduce模型
        sortedStr := string(s)
        mp[sortedStr] = append(mp[sortedStr], str) // $mp['aet'][] = 'aet'
    }

    ans := make([][]string, 0, len(mp))
    for _, v := range mp {
        ans = append(ans, v)
    }
    return ans
}
// @lc code=end

