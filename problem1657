func closeStrings(word1 string, word2 string) bool {
    chcnt1 := make(map[byte]int)
    chcnt2 := make(map[byte]int)
    ajaj1 := make(map[int]byte)
    ajaj2 := make(map[int]byte)
    govno1 := make(map[int]int)
    govno2 := make(map[int]int)
    if len(word1) != len(word2) {
        return false
    }
    for i:=0;i<len(word1);i++{
        chcnt1[word1[i]]++
        chcnt2[word2[i]]++
    }
    for k, v := range chcnt1{
        if _, ok := chcnt2[k]; !ok {
            return false
        }
        ajaj1[v] = k
        govno1[v]++
    }
    for k, v := range chcnt2{
        if _, ok := chcnt1[k]; !ok {
            return false
        }
        ajaj2[v] = k
        govno2[v]++
    }
    for k, v := range govno1{
        if v != govno2[k] {
            return false
        }
    }
    for k, v := range govno2{
        if v != govno1[k] {
            return false
        }
    }
    for i := 0; i < len(word1); i++ {
        if ajaj1[i] == 0 && ajaj2[i] == 0 {
            continue
        }
        if ajaj1[i] != 0 && ajaj2[i] != 0 {
            continue
        }
        return false
    }    
    return true
}
