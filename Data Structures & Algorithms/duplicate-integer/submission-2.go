func hasDuplicate(nums []int) bool {
    freq := map[int]int{}
    for _,n := range nums{
        freq[n] += 1
        if freq[n] > 1 {
            return true
        }
    }
    return false
}
