/*
 * @lc app=leetcode.cn id=3 lang=golang
 * @lcpr version=30110
 *
 * [3] 无重复字符的最长子串
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func lengthOfLongestSubstring(s string) int {

	result := make([]string, 0)

	start := 0

	tempStr := ""

	// 先遍历出所有可能出现的结果
	for i := 0; i < len(s); i++ {
		resultStr := string(s[i])

		startStr := string(s[start])

		for j := i + 1; j < len(s); j++ {
			tempStr = string(s[j])
			if startStr == tempStr {
				break
			}
			resultStr += tempStr
		}

		result = append(result, resultStr)

		start += 1
	}

	// 计算未重复的最常子串
	resultStr := ""
	resultStrLen := 0
	for _, s2 := range result {
		isAppend := true
		for i := 0; i < len(s2); i++ {
			strOne := string(s2[i])
			for j := i + 1; j < len(s2); j++ {
				strTwo := string(s2[j])
				if strOne == strTwo {
					isAppend = false
				}
			}
		}

		if isAppend {

			if resultStrLen < len(s2) {
				resultStr = s2
				resultStrLen = len(s2)
			}
		}
	}

	fmt.Println(resultStr)

	return len(resultStr)
}
// @lc code=end



/*
// @lcpr case=start
// "abcabcbb"\n
// @lcpr case=end

// @lcpr case=start
// "bbbbb"\n
// @lcpr case=end

// @lcpr case=start
// "pwwkew"\n
// @lcpr case=end

 */

