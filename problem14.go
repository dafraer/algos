func longestCommonPrefix(strs []string) string {
    ans := []byte(strs[0])
    for i := 0; i < len(strs); i++ {
        for j := 0; j < len(ans); j++ {
            if j >= len(strs[i]) {
                ans = ans[0:j]
                break
            }
            if strs[i][j] != ans[j] {
                ans = ans[0:j]
                break
            }
        } 
    }
    return string(ans)
}
