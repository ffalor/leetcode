func finalValueAfterOperations(operations []string) int {

	ans := 0

	for _, v := range operations {
		switch v {
			case "X--", "--X": 
				ans--
			case "X++", "++X":
				ans++
		}
	}

	return ans
}
