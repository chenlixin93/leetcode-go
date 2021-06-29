/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
    
    value_to_key := map[int]int{}

    for cur_key, v := range nums {
        key, ok := value_to_key[v]
        if ok {
            return []int{cur_key, key}
        } else {
            value_to_key[target - v] = cur_key
        }
    }

    return []int{}
}
// @lc code=end

