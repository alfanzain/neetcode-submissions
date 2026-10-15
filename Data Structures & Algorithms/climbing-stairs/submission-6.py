class Solution:
    def climbStairs(self, n: int) -> int:
        cache = [-1] * (n+1)
        def dfs(step: int) -> int:
            if step >= n:
                return step == n
            if cache[step] != -1:
                return cache[step]
                
            cache[step] = dfs(step + 1) + dfs(step + 2)
            return cache[step]

        return dfs(0)