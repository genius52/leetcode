class Solution:
    def maximumWealth(self, accounts: list[list[int]]) -> int:
        res = 0
        for account in accounts:
            total = 0
            for money in account:
                total += money
            if total > res:
                res = total
        return res
