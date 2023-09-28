class Solution:
    def isRobotBounded(self, instructions: str) -> bool:
        dirs = [[0,1],[1,0],[0,-1],[-1,0]]
        idx = 0
        x,y = 0,0
        for s in instructions:
            if s == "G":
                x += dirs[idx][0]
                y += dirs[idx][1]
                pass
            elif s == "L":
                idx = (idx + 4 - 1)% 4
                pass
            elif s == "R":
                idx = (idx + 1)%4
                pass
        if x == 0 and y == 0:
            return True
        if idx == 0:
            return False
        else:
            return True