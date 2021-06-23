<?php
/*
 * @lc app=leetcode.cn id=560 lang=php
 *
 * [560] 和为K的子数组
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return Integer
     */
    function subarraySum($nums, $k) {
        $map = [0 => 1]; // 
        $prefix_sum = 0;
        $ans = 0;

        foreach ($nums as $key => $num) {
            // 累加前缀和
            $prefix_sum += $num;

            // sum(l, r) = S[r] - S[l - 1]
            if (isset($map[$prefix_sum - $k])) {
                $ans += $map[$prefix_sum - $k];
            }

            if (isset($map[$prefix_sum])) {
                $map[$prefix_sum]++;
            } else {
                $map[$prefix_sum] = 1;
            }
        }
        return $ans;
    }
}
// @lc code=end