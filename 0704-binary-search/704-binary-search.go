func search(nums []int, target int) int {
	left, right := 0, len(nums) - 1

	for right >= left {
		mid := (left + right) / 2

		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}