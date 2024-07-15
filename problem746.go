func minCostClimbingStairs(cost []int) int {
    ans := make([]int, len(cost))
    ans[0] = cost[0]
    ans[1] = cost[1]
    for i := 2; i < len(ans); i++ {
        ans[i] = cost[i] + min(ans[i-1], ans[i-2])
    }
    return min(ans[len(ans)-1], ans[len(ans)-2])
}
