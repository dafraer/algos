func tribonacci(n int) int {
    s := make([]int, max(n+1, 3))
    s[0] = 0
    s[1] = 1
    s[2] = 1
    for i := 3; i < n+1; i++ {
        s[i] = s[i-1] + s[i-2] + s[i-3]
    }
    return s[n]
}
