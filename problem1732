func largestAltitude(gain []int) int {
    mx := 0
    past := 0
    cur := 0
    for i := 0; i < len(gain); i++ {
        cur = past+ gain[i] 
        mx = max(mx, cur)
        past = cur
    }
    return mx
}
