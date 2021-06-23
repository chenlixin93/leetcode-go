<?php
/*
 * @lc app=leetcode.cn id=66 lang=php
 *
 * [66] 加一
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[] $digits
     * @return Integer[]
     */
    function plusOne($digits) {

        for ($i=count($digits) - 1; $i >= 0; $i--) { 
            $digits[$i] += 1; // 从末尾加1
            $digits[$i] = $digits[$i]%10;
            if ($digits[$i]%10 != 0) {
                return $digits;
            }
        }
        return array_merge([1], $digits);
    }
}
// @lc code=end

