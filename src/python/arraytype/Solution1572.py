class Solution:
    def diagonalSum(self, mat: list[list[int]]) -> int:
        l = len(mat)
        res = 0
        for i in range(l):
            res += mat[i][i]
            res += mat[i][l - 1 - i]
        if l | 1 == l:
            res -= mat[l//2][l//2]
        return res