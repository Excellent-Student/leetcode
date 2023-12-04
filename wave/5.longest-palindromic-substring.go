/*
 * @lc app=leetcode.cn id=5 lang=golang
 * @lcpr version=30111
 *
 * [5] 最长回文子串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

func longestPalindrome(s string) string {

	var (
		start int
		end   int
	)

	for i := 0; i < len(s); i++ {
		strI := string(s[i])
		for j := i; j < len(s); j++ {
			strJ := string(s[j])
			if strI == strJ {
				if j-i > end-start {
					start = i
					end = j
				}
			}
		}
	}

	fmt.Println(start, end)
	if start == 0 && end == 0 {
		return ""
	}
	return s[start : end+1]
}

// @lc code=end

/*
// @lcpr case=start
// "babad"\n
// @lcpr case=end

// @lcpr case=start
// "cbbd"\n
// @lcpr case=end

*/

