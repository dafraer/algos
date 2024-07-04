func findDifference(nums1 []int, nums2 []int) [][]int {
    ans := make([][]int, 2)
    m := make(map[int]int)
    for _, v := range nums1 {
        m[v] = 1
    }
    for _, v := range nums2 {
        if m[v] == 1 {
            m[v] = 4
        } else if m[v] == 0{
            m[v] = 2
        }
    }
    for k, v := range m {
        if v == 1 {
            ans[0] = append(ans[0], k)
        } else if v == 2 {
            ans[1] = append(ans[1], k)
        }
    }
    return ans
}
