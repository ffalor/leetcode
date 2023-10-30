func topKFrequent(nums []int, k int) (res []int) {
	countM := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for _, v := range nums {
		countM[v]++
	}

	for k, v := range countM {
		freq[v] = append(freq[v], k)
	}

	for i := len(freq) - 1; i > 0; i-- {
		res = append(res, freq[i]...)
		if len(res) == k {
			return
		}
	}

	return
}