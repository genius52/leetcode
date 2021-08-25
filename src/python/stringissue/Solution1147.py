
class Solution_1147:
    def longestDecomposition(self, text: str) -> int:
        l = len(text)
        left = 0
        right = l - 1
        head = 0
        tail = 0
        sub_len = 0
        cnt = 0
        last_left = 0
        last_right = l - 1
        find = False
        while left < right:
            head = head * 26 + (ord(text[left]) - ord('a'))
            tail = (ord(text[right]) - ord('a')) * pow(26,sub_len) + tail
            if tail == head:
                sub_len = 0
                head = 0
                tail = 0
                cnt += 1
                last_left = left
                last_right = right
                find = True
            else:
                sub_len += 1
            left += 1
            right -= 1
        if not find:
            return 1
        elif last_left + 1 == last_right:
            return cnt * 2
        else:
            return cnt * 2 + 1
