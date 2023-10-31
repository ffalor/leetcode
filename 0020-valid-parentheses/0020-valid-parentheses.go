func isValid(s string) bool {
	stack := make([]byte, 0)
	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, c := range []byte(s) {
		pair, ok := pairs[c]
		if !ok {
			stack = append(stack, c)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if stack[len(stack)-1] != pair {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

