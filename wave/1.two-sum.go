/*
 * @lc app=leetcode.cn id=1 lang=golang
 * @lcpr version=30109
 *
 * [1] 两数之和
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func twoSum(nums []int, target int) []int {

	// 1
	// for i := 0; i < len(nums); i++ {
	// 	a := nums[i]
	// 	for j:=0; j < len(nums); j++ {
	// 		b := nums[j]
	// 		if a + b == target {
	// 			return []int{i, j}
	// 		}
	// 	}
	// }

	// 2
	keyMap := make(map[int]int)
	for i:=0; i<len(nums); i++ {
		num := nums[i]

		value, ok := keyMap[target - num]
		fmt.Println(i, value)
		if ok {
			return []int{value, i}
		}
		keyMap[num] = i
	}

	return nil
}
// @lc code=end



/*
// @lcpr case=start
// [2,7,11,15]\n9\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [3,3]\n6\n
// @lcpr case=end

 */

