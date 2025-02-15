func pivotIndex(nums []int) int {
    sum := make([]int, len(nums))
    sum[0] = nums[0]
    for i:=1; i < len(nums); i++ {
        sum[i] = sum[i-1] + nums[i]
    }
    if sum[len(sum)-1]-sum[0] == 0 {
        return 0
    }
    for i:=1; i < len(nums); i++ {
        if sum[len(sum)-1]-sum[i] == sum[i-1] {
            return i
        }
    }
    return -1
}
