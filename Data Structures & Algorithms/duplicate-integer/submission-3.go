func hasDuplicate(nums []int) bool {
    seen := map[int]bool{}
    for _,n := range nums{
        if seen[n] == true {
            return true
        }
        seen[n] = true
    }
    return false
}
