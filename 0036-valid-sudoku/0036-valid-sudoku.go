func isValidSudoku(board [][]byte) bool {
	hashMap := map[string]bool{}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}

			rKey := fmt.Sprintf("%vr%d", board[r][c], r)
			cKey := fmt.Sprintf("%v%d", board[r][c], c)
			sqreKey := fmt.Sprintf("%vs%d-%d", board[r][c], r/3, c/3)

			_, rOk := hashMap[rKey]
			_, cOk := hashMap[cKey]
			_, sOk := hashMap[sqreKey]

			if rOk || cOk || sOk {
				return false
			}

			hashMap[rKey] = true
			hashMap[cKey] = true
			hashMap[sqreKey] = true
		}
	}

	return true
}
