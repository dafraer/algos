func countBits(n int) []int {
    ans := make([]int, n+1)
    for i := 0; i < n+1; i++ {
        num := i
        for num > 0 {
           ans[i]+=num%2
           num /= 2 
        }

    }
    return ans
}
