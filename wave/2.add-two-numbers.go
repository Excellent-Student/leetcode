/*
 * @lc app=leetcode.cn id=2 lang=golang
 * @lcpr version=30109
 *
 * [2] 两数相加
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	temp := &ListNode{}
	cur := temp
	t := 0
	for l1 != nil || l2 != nil || t != 0 {
		if l1 != nil {
			t += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			t += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{t % 10, nil}
		cur = cur.Next
		t /= 10
	}
	return temp.Next
}
// @lc code=end



/*
// @lcpr case=start
// [2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n
// @lcpr case=end

// @lcpr case=start
// [9,9,9,9,9,9,9]\n[9,9,9,9]\n
// @lcpr case=end

 */

