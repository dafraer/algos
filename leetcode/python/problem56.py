 class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals.sort()
        int_begin =  intervals[0][0]
        int_end = intervals[0][1]
        res = []
        for interval in intervals:
            if interval[0] > int_end:
                res.append([int_begin, int_end])
                int_begin, int_end = interval[0], interval[1]
            else:
                int_end = max(int_end, interval[1])
        res.append([int_begin, int_end])
        return res
