package main

import "fmt"

/*
 * @lc app=leetcode.cn id=4 lang=golang
 * @lcpr version=30111
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	// 数组合并
	result := append(nums1, nums2...)

	// 排序
	result = sorted(result)
	fmt.Println(result)

	index := len(result) / 2
	// 说明列表中是奇数
	if len(result)%2 == 1 {
		return float64(result[index])
	}

	// 说明是偶数
	one := result[index-1]
	two := result[index]
	return (float64(one) + float64(two)) / 2
}

func sorted(items []int) []int {
	if len(items) < 1 {
		return items
	}

	for i := 0; i < len(items); i++ {
		for j := 1; j < len(items)-i; j++ {
			if items[j] < items[j-1] {
				items[j], items[j-1] = items[j-1], items[j]
			}
		}
	}
	return items
}

// @lc code=end

/*
// @lcpr case=start
// [1,3]\n[2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n[3,4]\n
// @lcpr case=end

*/
