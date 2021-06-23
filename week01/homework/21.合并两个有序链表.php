<?
/*
 * @lc app=leetcode.cn id=21 lang=php
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val = 0, $next = null) {
 *         $this->val = $val;
 *         $this->next = $next;
 *     }
 * }
 */
class Solution {

    /**
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function mergeTwoLists($l1, $l2) {
        if ($l1 == null) return $l2;
        if ($l2 == null) return $l1;
        
        $res = new ListNode();
        if ($l1->val >= $l2->val) {
            $res->val = $l2->val;
            $res->next = $this->mergeTwoLists($l1, $l2->next);
        } else {
            $res->val = $l1->val;
            $res->next = $this->mergeTwoLists($l1->next, $l2);
        }

        return $res;
    }
}
// @lc code=end

