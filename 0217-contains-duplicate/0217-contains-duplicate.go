func containsDuplicate(nums []int) bool {
    prev := make(map[int]bool, len(nums))

    for _, v := range nums {
        _, ok := prev[v]

        if ok {
            return true
        }

        prev[v] = true
    }

    return false
}