class Solution:
    def tictactoe(self, moves: list[list[int]]) -> str:
        l = len(moves)
        rows = [0] * 3
        columns = [0] * 3
        tile = 0
        reverse_tile = 0
        for i in range(l):
            x,y = moves[i]
            if (i | 1) == i:#Round B
                rows[x] -= 1
                columns[y] -= 1
                if x == y:
                    tile -= 1
                if (x + y) == 2:
                    reverse_tile -= 1
            else:           #Round A
                rows[x] += 1
                columns[y] += 1
                if x == y:
                    tile += 1
                if (x + y) == 2:
                    reverse_tile += 1
        for i in range(3):
            if rows[i] == 3 or columns[i] == 3:
                return "A"
            elif rows[i] == -3 or columns[i] == -3:
                return "B"
        if tile == 3 or reverse_tile == 3:
            return "A"
        elif tile == -3 or reverse_tile == -3:
            return "B"
        if l == 9:
            return "Draw"
        else:
            return "Pending"

if __name__ == '__main__':
    s1275 = Solution()
    moves = [[0, 0], [1, 1], [0, 1], [0, 2], [1, 0], [2, 0]]
    res = s1275.tictactoe(moves)
    print(res)