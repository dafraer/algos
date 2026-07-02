#didnt solve
class Solution:
    def findMinHeightTrees(self, n: int, edges: List[List[int]]) -> List[int]:
        if len(edges) == 0: return [n-1]

        tree = {i:[] for i in range(n)}
        for a, b in edges:
            tree[a].append(b)
            tree[b].append(a)

        leaves = deque()
        edge_cnt = {}
        for k, v in tree.items():
            if len(v) == 1: 
                leaves.append(k)
            edge_cnt[k] = len(v)

        while leaves:
            if n <= 2: return list(leaves)
            for i in range(len(leaves)):
                node = leaves.popleft()
                n -=1 
                for nei in tree[node]:
                    edge_cnt[nei] -= 1
                    if edge_cnt[nei] == 1:
                        leaves.append(nei)



