func groupAnagrams(strs []string) [][]string {
	res := make(map[[26]int][]string)

	for _, v := range strs {
		var count [26]int

		for _, v := range v {
			count[int(v)-int('a')] += 1
		}

		res[count] = append(res[count], v)
	}

	var result [][]string

	for _, v := range res {
		result = append(result, v)
	}

	return result
}