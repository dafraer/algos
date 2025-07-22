func maximumUniqueSubarray(nums []int) int {
    mx := 0
    sum := 0
    i, j := 0, 0
    m := make(map[int]struct{})
    for j < len(nums) {
        if _, ok := m[nums[j]]; ok {
            mx = max(sum, mx)
            for nums[i] != nums[j] {
                delete(m, nums[i])
                sum -= nums[i]
                i++
            }
            delete(m, nums[i])
            sum -= nums[i]
            i++
        } 
        sum += nums[j]
        m[nums[j]] = struct{}{}
        j++
    } 
    return max(sum, mx)
}
