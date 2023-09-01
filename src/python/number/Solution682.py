class Solution682:
    def calPoints(self, operations: list[str]) -> int:
        q = []
        for opt in operations:
            if opt == "C":
                q.pop(len(q) - 1)
            elif opt == "D":
                q.append(q[len(q) - 1] * 2)
            elif opt == "+":
                sum = q[len(q) - 1] + q[len(q) - 2]
                q.append(sum)
                pass
            else:
                q.append(int(opt))
        total = 0
        for val in q:
            total += val
        return total