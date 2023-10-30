func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, v := range strs {
		var count [26]int

		for _, v := range v {
			count[v-'a']++
		}

		m[count] = append(m[count], v)
	}

	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}

	return result
}