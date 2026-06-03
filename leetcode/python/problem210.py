# Didnt solve
class Solution:
    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        graph = [[] for _ in range(numCourses)]
        indegree = [0] * numCourses
        for course, prereq in prerequisites:
            graph[prereq].append(course)
            indegree[course] += 1

        q = deque(i for i in range(numCourses) if indegree[i] == 0)
        res = []
        while q:
            node = q.popleft()
            res.append(node)
            for x in graph[node]:
                indegree[x] -= 1
                if indegree[x] == 0:
                    q.append(x)

        return res if len(res) == numCourses else []
