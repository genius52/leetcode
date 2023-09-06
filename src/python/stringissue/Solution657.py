import collections

class Solution:
    def judgeCircle(self, moves: str) -> bool:
        y = 0
        x = 0
        for m in moves:
            if m == 'L':
                x += 1
            elif m == 'R':
                x -= 1
            elif m == 'U':
                y += 1
            elif m == 'D':
                y -= 1
        return x == 0 and y == 0

    def judgeCircle2(self, moves: str) -> bool:
        c = collections.Counter(moves)
        return (c['L'] - c['R']) == 0 and (c['U'] - c['D']) == 0