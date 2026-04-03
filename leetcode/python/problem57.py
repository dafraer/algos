class Solution:
    def insert(
        self, intervals: List[List[int]], newInterval: List[int]
    ) -> List[List[int]]:
        if len(intervals) == 0:
            return [newInterval]
        for i in range(len(intervals)):
            if intervals[i][1] < newInterval[0]:
                continue
            if intervals[i][0] <= newInterval[0] and newInterval[1] <= intervals[i][1]:
                return intervals
            if intervals[i][0] > newInterval[1]:
                return intervals[:i] + [newInterval] + intervals[i:]
            else:
                int_begin = min(intervals[i][0], newInterval[0])
                int_end = max(intervals[i][1], newInterval[1])
                res = intervals[:i]
                for j in range(i, len(intervals)):
                    if intervals[j][0] > int_end:
                        return res + [[int_begin, int_end]] + intervals[j:]
                    else:
                        int_end = max(int_end, intervals[j][1])
                res.append([int_begin, int_end])
                return res
        return intervals + [newInterval]
