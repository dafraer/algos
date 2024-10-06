func longestOnes(nums []int, k int) int {
    l := 0
    mx := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            k--
        }
        if k < 0 {
            if nums[l] == 0 {
                k++
            }
            l++
        } else {
            mx = max(mx, i-l+1)
        }
        
    }
    return mx
}
