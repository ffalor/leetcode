func isAnagram(s string, t string) bool {

    o := make(map[rune]int)

    for _, r := range s {
        o[r]++
    }

    for _, r := range t {
        o[r]--
    }

    for _, occurance := range o {
        if occurance != 0 {
            return false
        }
    }

    return true
}