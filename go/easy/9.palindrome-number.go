/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	r, t := 0, x

	for t != 0 {
		pop := t % 10
		t /= 10
		r = r*10 + pop
	}

	return r == x
}

// @lc code=end

