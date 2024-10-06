func canVisitAllRooms(rooms [][]int) bool {
    visited := make(map[int]bool) 
    dfs(0, rooms, visited)
    for i := 0; i < len(rooms); i++ {
        if !visited[i] {
            return false
        }
    }
    return true
}

func dfs(cur int, rooms [][]int, visited map[int]bool) {
    if visited[cur] {
        return
    }
    visited[cur] = true
    for i := 0; i < len(rooms[cur]); i++ {
        dfs(rooms[cur][i], rooms, visited)
    }
}
