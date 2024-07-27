func findCircleNum(isConnected [][]int) int {
    nodes := make([][]int, len(isConnected))
    for i := 0; i < len(isConnected); i++ {

        for j := 0; j < len(isConnected[i]); j++ {
            if isConnected[i][j] == 1 {
              nodes[i] = append(nodes[i], j)  
            }
        }
    }
    colors := make(map[int]int)
    for i := 0; i < len(nodes); i++ {
        dfs(i,nodes, i, colors)
    }
    provinces := make(map[int]bool)
    cnt := 0
    for i := 0; i < len(nodes); i++ {
        n, ok := colors[i]
        if !ok {
            cnt++
        } else {
            provinces[n] = true
        }
    }
    for _, v := range provinces {
        if v {
            cnt++
        } 
    }
    return cnt
}

func dfs(cur int, nodes [][]int, color int, colors map[int]int) {
    if colors[cur] == color {
        return 
    }
    colors[cur] = color
    for i := 0; i < len(nodes[cur]); i++ {
        dfs(nodes[cur][i], nodes, color, colors)
    }
}
