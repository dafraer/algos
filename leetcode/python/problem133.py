"""
# Definition for a Node.
class Node:
    def __init__(self, val = 0, neighbors = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []
"""

from typing import Optional


class Solution:
    def cloneGraph(self, node: Optional["Node"]) -> Optional["Node"]:
        visited = {}

        def dfs(node):
            if not node:
                return None
            if node in visited:
                return visited[node]
            new_node = Node(node.val, [])
            visited[node] = new_node
            new_node.neighbors = [dfs(n) for n in node.neighbors]
            return new_node

        return dfs(node)
