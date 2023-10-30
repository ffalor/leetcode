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