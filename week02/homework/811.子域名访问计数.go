/*
 * @lc app=leetcode.cn id=811 lang=golang
 *
 * [811] 子域名访问计数
 */

// @lc code=start
func subdomainVisits(cpdomains []string) (ans []string) {
	count := make(map[string]int, 0)
	for _, v := range cpdomains {
		splits := strings.Split(v, " ")
		num, _ := strconv.Atoi(splits[0])
		for {
			count[splits[1]] += num
			// 找当前域名第一个点的位置
			dotIndex := strings.Index(splits[1], ".")
			if dotIndex < 0 { break }
			// 得到更高一级的域名
			splits[1] = splits[1][dotIndex+1:] 
		}
	}

	//ans := make([]string, 0, len(count))
	for k, v := range count {
		ans = append(ans, fmt.Sprintf("%d %s", v, k))
	}
	return
}
// @lc code=end