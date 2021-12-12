/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {

	l := len(needle)

	for i := 0; i < len(haystack)-l+1; i++ {
		if haystack[i:i+l] == needle {
			return i
		}
	}

	return -1
}

// @lc code=end

