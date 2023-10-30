func topKFrequent(nums []int, k int) []int {
	countM := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for _, v := range nums {
		countM[v]++
	}

	for k, v := range countM {
		freq[v] = append(freq[v], k)
	}

	result := make([]int, 0, k)
	for i := len(freq) - 1; i > 0; i-- {
		for _, v := range freq[i] {
			result = append(result, v)
			if len(result) == k {
				return result
			}
		}
	}

	return result
}