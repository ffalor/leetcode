func isValidSudoku(board [][]byte) bool {
	
	var rows, columns, squares [9][9]bool

	for i, row := range board {
		for j, v := range row {
			if v != '.' {
				// 1 = 49
				// 2 = 50... etc
				k := int(v)-49
				
				// square index is i/3*3 + j/3 to allow us to map to the actual square in the 3x3 grid 
				if rows[i][k] || columns[j][k] || squares[i/3*3 + j/3][k] {
					return false
				}
				rows[i][k], columns[j][k], squares[i/3*3 + j/3][k] = true, true, true
			}

		}
	}
	return true
}
