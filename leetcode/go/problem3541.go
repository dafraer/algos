func maxFreqSum(s string) int {
    m := make(map[byte]int)
    for i:=0;i<len(s);i++ {
        m[s[i]]++
    }
    mxC := 0
    mxV := 0
    for k, v := range(m) {
        if k == 'a' || k == 'i' || k == 'e' || k == 'o' || k == 'u' {
            mxV = max(mxV, v)
        } else {
            mxC = max(mxC, v)
        }
    }
    return mxC + mxV
}
