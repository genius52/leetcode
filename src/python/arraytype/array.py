import sys
from threading import Lock
from enum import Enum
from heapq import *
from functools import reduce

# if __name__ == "__main__":
#     s = Solution()
#     image = [[1,1,1],[1,1,0],[1,0,1]]
#     s.floodFill(image,1,1,2)

class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution_697(object):
    def findShortestSubArray(self, nums) -> int:
        dict = {}
        for i in nums:
            if i in dict:
                dict[i] += 1
            else:
                dict[i] = 1
        max_occurs_ele = -1
        max_occurs_count = 0
        for k,v in dict.items():
            if v > max_occurs_count:
                max_occurs_count = v
                max_occurs_ele = k
        res = []
        for k,v in dict.items():
            if v == max_occurs_count:
                res.append(k)
        degree = len(nums)
        for e in res:
            j = len(nums) - 1
            end = j
            start = nums.index(e)
            while j > 0:
                if nums[j] == e:
                    end = j
                    break
                else:
                    j -= 1
            if (end - start + 1) <= degree:
                degree = end - start + 1
        return degree


class Solution_860:
    def lemonadeChange(self, bills) -> bool:
        change = {}
        for pay in bills:
            if pay == 5:
                if pay in change:
                    change[pay] += 1
                else:
                    change[pay] = 1
            elif pay == 10:
                if 5 not in change or change[5] <= 0:
                    return False
                else:
                    change[5] -= 1
                if pay in change:
                    change[pay] += 1
                else:
                    change[pay] = 1
            elif pay == 20:
                if 5 not in change or change[5] <= 0:
                    return False
                if 10 in change and change[10] > 0:
                    change[10] -=1
                    change[5] -= 1
                elif change[5] >= 3:
                    change[5] -= 3
                else:
                    return False
        return True


class Solution_455:
    def findContentChildren(self, g, s) -> int:
        g = sorted(g)
        s = sorted(s)
        i = 0
        j = 0
        while i < len(g) and j < len(s):
            if s[j] >= g[i]:
                i += 1
            j += 1
        return i


class Solution_1071:
    def gcdOfStrings(self, str1: str, str2: str) -> str:
        res = ""
        if str1 == str2:
            return str1
        short_str = str1 if len(str1) < len(str2) else str2
        long_str = str1 if len(str1) >= len(str2) else str2
        pos = long_str.find(short_str,0)
        if pos == -1:
            return res
        rest = long_str[len(short_str):]
        return self.gcdOfStrings(rest,short_str)

class Employee:
    def __init__(self, id, importance, subordinates):
        # It's the unique id of each node.
        # unique id of this employee
        self.id = id
        # the importance value of this employee
        self.importance = importance
        # the id of direct subordinates
        self.subordinates = subordinates

"""
# Employee info
class Employee:
    def __init__(self, id, importance, subordinates):
        # It's the unique id of each node.
        # unique id of this employee
        self.id = id
        # the importance value of this employee
        self.importance = importance
        # the id of direct subordinates
        self.subordinates = subordinates
"""
#[[1,5,[2,3]],[2,3,[]],[3,3,[]]]
#[[1,5,[2,3]],[2,3,[4]],[3,4,[]],[4,1,[]]]
importance_map = {}
class Solution_690:
    def get_importance(self,employee):
        res = 0
        res += employee.importance
        for id in employee.subordinates:
            res += self.get_importance(importance_map[id])
        return res


    def getImportance(self, employees, id):
        """
        :type employees: Employee
        :type id: int
        :rtype: int
        """
        global importance_map
        for p in employees:
            importance_map[p.id] = p
            if p.id == id:
                root_people = p
        total = 0
        total += self.get_importance(root_people)
        return total


class Solution_1018:
    def prefixesDivBy5(self, A):
        res = []
        val = -1
        for i in range(len(A)):
            if val == -1:
                val = A[i]
            else:
                val = val * 2 + A[i]
            res.append(val%5==0)
        return res


class Solution_720:
    def longestWord(self, words) -> str:
        if len(words) == 1:
            return words[0]
        new_words = sorted(words)
        dict = set()
        res = ""
        for w in new_words:
            if len(w) > len(res):
                exist = True
                i = 1
                while i < len(w):
                    if w[:i] not in dict:
                        exist = False
                        break
                    i += 1
                if exist:
                    res = w
            dict.add(w)
        return res

#[8,7,3,5,3,6,1,4])
class Solution_645:
    def findErrorNums(self, nums):
        res = []
        dict = {}
        for num in nums:
            if num not in dict:
                dict[num] = 1
            else:
                res.append(num)
        i = 1
        while i < len(nums):
            if i not in dict:
                res.append(i)
                break
            i += 1
        return res


def get_value(n):
    if n == 0:
        return 1
    elif n == 1:
        return n
    else:
        return n * get_value(n - 1)

def gen_last_value(n, m):
    first = get_value(n)
    second = get_value(m)
    third = get_value((n - m))
    return first / (second * third)

class Solution_1128:
    def numEquivDominoPairs(self, dominoes):
        res = 0
        dict = {}
        for ele in dominoes:
            min = ele[0] if ele[0] < ele[1] else ele[1]
            max = ele[0] if ele[0] > ele[1] else ele[1]
            if (min,max) in dict:
                #res += dict[(min,max)]
                dict[(min, max)] += 1
            else:
                dict[(min,max)] = 1

        for k,v in dict.items():
            if v >= 2:
                res += gen_last_value(v,2)

        return res


class Solution_438:
    def findAnagrams(self, s: str, p: str):
        res = []
        len_p = len(p)
        record = [-1] * 26
        for c in p:
            i = ord(c)-ord('a')
            if record[i] == -1:
                record[i] = 1
            else:
                record[i] += 1
        start = 0
        end = 0
        match_cnt = 0
        while end < len(s):
            if record[ord(s[end])-ord('a')] > 0:
                match_cnt += 1
            record[ord(s[end])-ord('a')] -= 1

            if end - start + 1 == len_p:
                if match_cnt == len_p:
                    res.append(start)
#开始新的搜索
                if record[ord(s[start])-ord('a')] >= 0:
                    match_cnt -= 1
#还原之前减去的
                record[ord(s[start])-ord('a')] += 1
                start += 1
            end += 1

        return res


class Solution_172:
    def trailingZeroes(self, n: int) -> int:
        res = 0
        sum = 1
        for i in range(1,n+1):
            sum *= i

        s = str(sum)[::-1]
        for c in s:
            if c == '0':
                res += 1
            else:
                break
        return res


class Solution_219:
    def containsNearbyDuplicate(self, nums, k) -> bool:
        record = {}
        for index in range(0,len(nums)):
            if nums[index] in record:
                if index - record[nums[index]] <= k:
                    return True
                else:
                    record[nums[index]] = index
            else:
                record[nums[index]] = index
        return False

#69. Sqrt(x)
class Solution_69:
    def mySqrt(self, x: int) -> int:
        if x == 0:
            return 0
        low = 1
        high = x
        while low <= high:
            mid = low + (high-low)//2
            val = mid * mid
            if val == x:
                return mid
            elif val < x:
                low = mid + 1
            else:
                high = mid - 1
        return high


class Solution_605:
    def canPlaceFlowers(self, flowerbed: [int], n: int) -> bool:
        if n == 0:
            return True
        flowerbed.insert(0,0)
        flowerbed.append(0)
        l = len(flowerbed)
        index = 0
        total = 0
        start_count = False
        temp_cnt = 0
        while index < l:
            if flowerbed[index] == 0:
                temp_cnt += 1
                start_count = True
            else:
                if start_count:
                    total += (temp_cnt - 1) // 2
                start_count = False
                temp_cnt = 0
            index += 1

        if start_count:
            total += (temp_cnt - 1) / 2
        return total >= n

import math


class Solution_686:
    def repeatedStringMatch(self, A: str, B: str) -> int:
        b_map = {}
        a_map = {}
        for b_e in B:
            if b_e in b_map:
                b_map[b_e] += 1
            else:
                b_map[b_e] = 1
        for a_e in A:
            if a_e in a_map:
                a_map[a_e] += 1
            else:
                a_map[a_e] = 1
        res = - 1
        prev = -1
        for k,v in b_map.items():
            if k not in a_map:
                return -1
            if prev == -1:
                prev = a_map[k]/v
                continue
            times = a_map[k]/v
            if times != prev:
                return -1
        if prev > 1:
            return 1
        else:
            res = 1/prev
            if res - (int(res)) == 0:
                return res + 1
            else:
                return math.ceil(1/prev)


class Solution_754:
    def reachNumber(self, target: int) -> int:

        pass

cnt = 0
l = []
class Solution_894:
    def allPossibleFBT(self, N: int):
        t = TreeNode(0)
        global cnt
        cnt += 1
        return l
        pass

    def recrusive(self,parent:TreeNode,cur_num,N :int,record):
        if cur_num > N - 2:
            return

        left = TreeNode(0)
        right = TreeNode(0)
        parent.left = left
        parent.right = right
        self.recrusive(left,cur_num + 2,N)
        self.recrusive(right,cur_num + 2, N)
        pass



class Solution_1200:
    def minimumAbsDifference(self, arr):
        s = sorted(arr)
        min_distinct = s[-1] - s[0]
        res = []
        for i in range(0,len(s) - 1):
            if (s[i+1] - s[i]) < min_distinct:
                min_distinct = s[i+1] - s[i]

        for i in range(0, len(s) - 1):
            if (s[i+1] - s[i]) == min_distinct:
                couple = [s[i],s[i+1]]
                res.append(couple)

        return res


import time
class FizzBuzz_1195:
    def __init__(self, n: int):
        self.n = n
        self.cur = 1
        self.fizzlock = Lock()
        self.buzzlock = Lock()
        self.fizzbuzzlock = Lock()
        self.numberlock = Lock()

    # printFizz() outputs "fizz"
    def fizz(self, printFizz: 'Callable[[], None]') -> None:
        while True:
            self.fizzlock.acquire()
            self.fizzbuzzlock.acquire()
            if self.cur % 3 == 0 and self.cur % 5 != 0:
                printFizz()
                self.cur += 1
            self.fizzlock.release()
            self.fizzbuzzlock.release()
            if self.cur >= self.n:
                break
        pass

    # printBuzz() outputs "buzz"
    def buzz(self, printBuzz: 'Callable[[], None]') -> None:
        while True:
            self.buzzlock.acquire()
            self.fizzbuzzlock.acquire()
            if self.cur % 5 == 0 and self.cur % 3 != 0:
                print("buzz " + self.cur)
                printBuzz()
                self.cur += 1
            self.buzzlock.release()
            self.fizzbuzzlock.release()
            if self.cur >= self.n:
                break
        pass

    # printFizzBuzz() outputs "fizzbuzz"
    def fizzbuzz(self, printFizzBuzz: 'Callable[[], None]') -> None:
        while True:
            self.lock.acquire()
            if self.n % 5 == 0 and self.n % 3 == 0:
                print("fizzbuzz " + self.cur)
                printFizzBuzz()
                self.cur += 1
            self.lock.release()
            if self.cur >= self.n:
                break
        pass

    # printNumber(x) outputs "x", where x is an integer.
    def number(self, printNumber: 'Callable[[int], None]') -> None:
        while True:
            self.lock.acquire()
            if self.cur % 5 != 0 and self.cur % 3 != 0:
                print("number " + self.cur)
                printNumber(self.cur)
                self.cur += 1
            self.lock.release()
            if self.cur >= self.n:
                break
        pass


class Solution_1190:
    def reverseParentheses(self, s: str) -> str:

        pass


class Solution_79:
    def dfs(self,board,word,cur_pos,row,col):
        if cur_pos >= len(word):
            return True
        if row < 0 or row >= len(board) or col < 0 or col >= len(board[0]) or board[row][col] == "." :
            return False
        if board[row][col] != word[cur_pos]:
            return False
        tmp = board[row][col]
        board[row][col] = '.'
        cur_pos += 1
        ret = self.dfs(board, word, cur_pos, row - 1, col) or self.dfs(board, word, cur_pos, row + 1, col) or \
               self.dfs(board, word, cur_pos, row, col - 1) or self.dfs(board, word, cur_pos, row, col + 1)
        board[row][col] = tmp
        return ret

    def exist(self, board, word: str) -> bool:
        for i in range(len(board)):
            for j in range(len(board[i])):
                ret = self.dfs(board,word,0,i,j)
                if ret:
                    return ret
                pass
        return False

# Input: nums =
# [
#   [9,9,4],
#   [6,6,8],
#   [2,1,1]
# ]
# Output: 4

class Num_State(Enum):
    Undefined = 0
    Increase = 1
    Decrease = 2

class Solution_329:
    def dfs(self,matrix,row,col,pre_num,is_increase,record,trace):
        k = str(row) + "," + str(col)
        if row < 0 or row >= len(matrix) or col < 0 or col >= len(matrix[0]) or trace.get(k):
            return 0

        if is_increase:
            if matrix[row][col] <= pre_num:
                return 0
        else:
            if matrix[row][col] >= pre_num:
                return 0
        if k in trace:
            return cnt

        trace[k] = 1
        if k in record:
             return record[k]
        ret = 1 + max(self.dfs(matrix,row - 1,col,matrix[row][col],is_increase,record,trace),self.dfs(matrix,row + 1,col,matrix[row][col],is_increase,record,trace),
        self.dfs(matrix, row, col - 1, matrix[row][col], is_increase,record,trace),self.dfs(matrix,row,col + 1,matrix[row][col],is_increase,record,trace))
        del trace[k]
        record[k] = ret
        return ret

    def longestIncreasingPath(self, matrix) -> int:
        max_increase_cnt = 0
        max_decrease_cnt = 0
        increase_record = {}
        decrease_record = {}
        for i in range(len(matrix)):
            for j in range(len(matrix[i])):
                trace1 = {}
                max_increase_cnt = max(self.dfs(matrix,i,j, sys.maxsize,False,increase_record,trace1),max_increase_cnt)
                trace2 = {}
                max_decrease_cnt = max(self.dfs(matrix, i, j,-sys.maxsize,True,decrease_record,trace2), max_decrease_cnt)
        return max(max_increase_cnt,max_decrease_cnt)

#Input: s = "10#11#12"
#Output: "jkab"
#Explanation: "j" -> "10#" , "k" -> "11#" , "a" -> "1" , "b" -> "2".
class Solution_1309:
    def freqAlphabets(self, s: str) -> str:
        i = len(s) - 1
        res = ""
        while i >= 0:
            if s[i] == "#":
                r = chr(int(s[i-2:i]) - 1 + ord("a"))
                res = r + res
                i = i - 3
            else:
                r = chr(ord(s[i]) - ord("1") + ord("a"))
                res = r + res
                i = i - 1
        return res

#Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
# Output: [3,3,5,5,6,7]
# Explanation:
#
# Window position                Max
# ---------------               -----
# [1  3  -1] -3  5  3  6  7       3
#  1 [3  -1  -3] 5  3  6  7       3
#  1  3 [-1  -3  5] 3  6  7       5
#  1  3  -1 [-3  5  3] 6  7       5
#  1  3  -1  -3 [5  3  6] 7       6
#  1  3  -1  -3  5 [3  6  7]      7
class Solution_239:
    def maxSlidingWindow(self, nums, k: int):
        heap = []

        heappush(heap, 2)
        heappush(heap, 3)
        heappush(heap, 3)
        heappush(heap,1)
        heappush(heap, 1)
        n = nlargest(1,heap)
        a = []
        heappop(heap)
        heapify(heap)
        return nums

# Input: "babad"
# Output: "bab"
# Note: "aba" is also a valid answer.
class Solution_5:
    def longestPalindrome(self, s: str) -> str:
        l = len(s)
        if l <= 1:
            return s
        dp = [] # dp[i][j] = 1 means from i to j is palindrome
        for i in range(l):
            dp.append([])
            for j in range(l):
                if i == j:
                    dp[i].append(1)
                else:
                    dp[i].append(0)
        end = 1
        max_len = 0
        res = ""
        while end < l:
            begin = end
            while begin >= 0:
                if s[begin] == s[end]:
                    if (end - begin) < 3 or dp[begin+1][end-1] == 1:
                        dp[begin][end] = 1
                        if (end - begin + 1) > max_len:
                            max_len = end - begin + 1
                            res = s[begin:end + 1]
                begin -= 1
            end += 1
        return res

# Given arraytype nums = [-1, 0, 1, 2, -1, -4],
# A solution set is:
# [
#   [-1, 0, 1],
#   [-1, -1, 2]
# ]
class Solution_15:
    def threeSum(self, nums):
        res = []
        l = len(nums)
        if l < 3:
            return res
        record = {}
        nums = sorted(nums)
        for i in range(len(nums)):
            if nums[i] in record:
                record[nums[i]].append(i)
            else:
                record[nums[i]] = [i]
        res_record = {}
        slow = 0
        while slow < l:
            fast = slow + 1
            if nums[slow] > 0:
                break
            while fast < l:
                #s = ''.join(str(i) for i in sorted([nums[slow], nums[fast], find]))
                if nums[slow] + nums[fast] > 0:
                    break
                k = str(nums[slow]) + str(nums[fast])
                if k in res_record:
                    fast += 1
                    continue
                find = -(nums[slow] + nums[fast])
                if find in record:
                    pos_list = record[find]
                    dup = True
                    find_pos = -1
                    for pos in pos_list:
                        if pos != slow and pos != fast:
                            dup = False
                            find_pos = pos
                            break
                    if dup:
                        fast += 1
                        continue
                    res.append([nums[slow],nums[fast],find])
                    res_record[str(nums[slow]) + str(nums[fast])] = True
                    res_record[str(nums[fast]) + str(nums[slow])] = True
                    res_record[str(nums[slow]) + str(nums[find])] = True
                    res_record[str(nums[find]) + str(nums[slow])] = True
                    res_record[str(nums[fast]) + str(nums[find])] = True
                    res_record[str(nums[find]) + str(nums[fast])] = True
                fast += 1
            slow += 1
        return res

class Solution_15ex:
    def threeSum(self, nums):
        res = []
        l = len(nums)
        if l < 3:
            return res
        nums = sorted(nums)
        head = 0
        end = l - 1
        last_head = 2e31
        while head <= end:
            mid = head + 1
            end = l - 1
            if head <= end and nums[head] == last_head:
                head += 1
                continue
            last_head = nums[head]
            while mid < end:
                last_end = nums[end]
                if nums[end] < 0:
                    break
                if nums[head] + nums[mid] + nums[end] < 0:
                    mid += 1
                elif nums[head] + nums[mid] + nums[end] > 0:
                    end -= 1
                    while end >= head and nums[end] == last_end:
                        end -= 1
                        continue
                else:
                    res.append([nums[head],nums[mid],nums[end]])
                    mid += 1
                    end -= 1
                    while end >= head and nums[end] == last_end:
                        end -= 1
                        continue
            head += 1
        return res

# Given arraytype nums = [-1, 2, 1, -4], and target = 1.
# The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
class Solution_16:
    def threeSumClosest(self, nums, target: int) -> int:
        diff = 2e31
        nums = sorted(nums)
        head = 0
        l = len(nums)
        end = l - 1
        sum = 0
        while head < end:
            last_head = nums[head]
            mid = head + 1
            end = l - 1
            while mid < end:
                last_end = nums[end]
                if abs(nums[head] + nums[mid] + nums[end] - target) < diff:
                    diff = abs(nums[head] + nums[mid] + nums[end] - target)
                    sum = nums[head] + nums[mid] + nums[end]
                if nums[head] + nums[mid] + nums[end] < target:
                    mid += 1
                elif nums[head] + nums[mid] + nums[end] > target:
                    end -= 1
                    while end >= 0 and last_end == nums[end]:
                        end -= 1
                        continue
                else:
                    return sum
            head += 1
            while head < l and last_head == nums[head]:
                head += 1
                continue
        return sum

#1310
# Input: arr = [1,3,4,8], queries = [[0,1],[1,2],[0,3],[3,3]]
# Output: [2,7,14,8]
class Solution_1310:
    #def xorQueries(self, arr: List[int], queries: List[List[int]]) -> List[int]:
    def xorQueries(self, arr, queries):
        # res = []
        # for ele in [arr[q[0]:q[1]+1] for q in queries]:
        #     res.append(reduce(lambda x,y:x^y,ele))
        # return res

        l = len(arr)
        i = 1
        record = [0] * l
        record[0] = arr[0]
        while i < l:
            record[i] = record[i-1] ^ arr[i]
            i += 1

        res = []
        for q in queries:
            if q[0] - 1 < 0:
                res.append(record[q[1]] ^ 0)
            else:
                res.append(record[q[1]] ^ record[q[0] - 1])
        return res

# Input: ")()())"
# Output: 4
# Explanation: The longest valid parentheses substring is "()()"
class Solution_32:
    def longestValidParentheses(self, s: str) -> int:
        pass

# Input: candidates = [10,1,2,7,6,1,5], target = 8,
# A solution set is:
# [
#   [1, 7],
#   [1, 2, 5],
#   [2, 6],
#   [1, 1, 6]
# ]
class Solution_40:
    def combine(self,candidates,cur_pos,cur_sum,target_num,cur_list,res,dup_record):
        if cur_sum > target_num:
            return
        if cur_sum == target_num:
            k = ''.join(str(i) for i in cur_list)
            if k not in dup_record:
                dup_record[k] = True
                res.append(cur_list.copy())
            return
        else:
            i = cur_pos
            while i < len(candidates):
                if (cur_sum + candidates[i]) > target_num:
                    break
                if i > cur_pos and candidates[i] == candidates[i-1]:
                    i += 1
                    continue
                cur_list.append(candidates[i])
                self.combine(candidates,i+1,cur_sum + candidates[i],target_num,cur_list,res,dup_record)
                del (cur_list[-1])
                #self.combine(candidates, i+1,cur_sum, target_num, cur_list, res,dup_record)
                i += 1


    def combinationSum2(self, candidates, target: int):
        candidates = sorted(candidates)
        dup_record = {}
        res = []
        self.combine(candidates,0,0,target,[],res,dup_record)
        return res


class Solution_1184:
    #def distanceBetweenBusStops(self, distance: List[int], start: int, destination: int) -> int:
    def distanceBetweenBusStops(self, distance, start: int, destination: int) -> int:
        l = len(distance)
        clockwise_disatance = 0
        i = start
        while i != destination:
            clockwise_disatance += distance[i]
            i = (i + 1) % l
        counterclockwise_distance = 0
        j = destination
        while j != start:
            counterclockwise_distance += distance[j]
            j = (j + 1) % l
        return min(clockwise_disatance,counterclockwise_distance)

class Solution_90:
    def dfs(self,cur_pos,target_len,nums,cur_list,res):
        if len(cur_list) > target_len or cur_pos > len(nums):
            return
        if len(cur_list) == target_len:
            res.append(cur_list)
            return
        for i in range(cur_pos,len(nums)):
            if (len(cur_list) + len(nums) - i) < target_len:
                break
            if i > cur_pos:
                if nums[i] == nums[i-1]:
                    continue
            self.dfs(i + 1,target_len,nums,cur_list + [nums[i]] ,res)
            i += 1

    def subsetsWithDup(self, nums):
        res = []
        nums = sorted(nums)
        for i in range(len(nums)+1):
            self.dfs(0,i,nums,[],res)
        return res


# input
# [0,0,0,0,0,1,1,1,2,2,3,3,5,5,6]
# Expected
# [0,0,1,1,2,2,3,3,5,5,6]
class Solution_80:
    def removeDuplicates(self, nums) -> int:
        l = len(nums)
        if l <= 2:
            return l
        tag = nums[0] - 1
        dup_cnt = 0
        pos = 1
        while pos < l - 1:
            if nums[pos-1] == nums[pos] and nums[pos] == nums[pos+1]:
                dup_cnt += 1
                nums[pos-1] = tag
            pos += 1
        res = l - dup_cnt
        slow = 0
        fast = 1
        while slow < res and fast < l:
            if nums[slow] != tag:
                slow += 1
            else:
                if fast <= slow:
                    fast = slow + 1
                while fast < l and nums[fast] == tag:
                    fast += 1
                if fast < l:
                    nums[slow],nums[fast] = nums[fast],nums[slow]
                    fast += 1
                slow += 1
        return res


# Input: n = 4, k = 2
# Output:
# [
#   [2,4],
#   [3,4],
#   [2,3],
#   [1,2],
#   [1,3],
#   [1,4],
# ]
class Solution_77:
    def dfs(self,start,k,max_num,sub,res):
        cur_len = len(sub)
        if cur_len == k:
            res.append(sub)
            return

        for i in range(start,max_num):
            self.dfs(i + 1,k,max_num,sub + [i],res)

    def combine(self, n: int, k: int):
        res = []
        self.dfs(1,k,n+1,[],res)
        return res

class Solution_74:
    def dfs(self,matrix,cur_row,cur_col,rows,columns,target):
        if cur_row >= rows or cur_col >= columns:
            return False
        if target == matrix[cur_row][cur_col]:
            return True
        if target < matrix[cur_row][cur_col]:
            return False
        if (cur_row + 1) < rows:
            if target >= matrix[cur_row+1][cur_col]:
                return self.dfs(matrix,cur_row + 1,cur_col,rows,columns,target)
            else:
                return self.dfs(matrix, cur_row,cur_col + 1, rows, columns,target)
        else:
            return self.dfs(matrix, cur_row,cur_col + 1, rows, columns,target)


    def searchMatrix(self, matrix, target: int) -> bool:
        rows = len(matrix)
        if rows == 0:
            return False
        columns = len(matrix[0])
        if columns == 0:
            return False
        if target > matrix[rows-1][columns-1]:
            return False
        return self.dfs(matrix,0,0,rows,columns,target)


# Input: n = 3, k = 3
# Output: "213"
# "123"
# "132"
# "213"
# "231"
# "312"
# "321"
class Solution_60:
    def dfs(self,begin,end,nums,k,res):
        if len(res) >= k:
            return True
        if begin == end:
            res.append(nums.copy())
            return False
        for i in range(begin,end):
            #nums[begin],nums[i] = nums[i],nums[begin]
            tmp = nums[i]
            nums[begin+1:i+1] = nums[begin:i]
            nums[begin] = tmp
            if self.dfs(begin+1,end,nums,k,res):
                return True
            tmp = nums[begin]
            nums[begin:i] = nums[begin+1:i+1]
            nums[i] = tmp
            #nums[begin], nums[i] = nums[i], nums[begin]
        return False

    def getPermutation(self, n: int, k: int) -> str:
        res = []
        nums = []
        for i in range(1,n+1):
            nums.append(i)

        self.dfs(0,len(nums),nums,k,res)
        if len(res) == k:
            return ''.join(str(i) for i in res[k-1])
        return ""
#[[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 2, 1], [3, 1, 2]]

# Input: 3
# Output:
# [
#  [ 1, 2, 3 ],
#  [ 8, 9, 4 ],
#  [ 7, 6, 5 ]
# ]
class Solution_59:
    def generateMatrix(self, n: int):
        res = []
        for i in range(n):
            res.append([0] * n)
        cur_num = 1
        init_pos = 0
        for edge_len in range(n,0,-2):
            row = init_pos
            col = init_pos
            res[row][col] = cur_num
            cur_num += 1
            col += 1
            while col < init_pos + edge_len:
                res[row][col] = cur_num
                col += 1
                cur_num += 1
            col -= 1
            row += 1
            while row < init_pos + edge_len:
                res[row][col] = cur_num
                row += 1
                cur_num += 1
            row -= 1
            col -= 1
            while col >= init_pos:
                res[row][col] = cur_num
                col -= 1
                cur_num += 1
            row -= 1
            col += 1
            while row > init_pos:
                res[row][col] = cur_num
                row -= 1
                cur_num += 1
            init_pos += 1
        return res


# Input:
# [
#  [ 1, 2, 3 ],
#  [ 4, 5, 6 ],
#  [ 7, 8, 9 ]
# ]
# Output: [1,2,3,6,9,8,7,4,5]
class Solution_54:
    def spiralOrder(self, matrix):
        res = []
        h = len(matrix)
        if h == 0:
            return res
        w = len(matrix[0])
        min_len = min(w,h)
        start_pos = 0
        while w > 0 and h > 0:
            if start_pos > min_len // 2:
                break
            row = start_pos
            col = start_pos
            if w == 1:
                while row < start_pos + h:
                    res.append(matrix[row][col])
                    row += 1
                break
            if h == 1:
                while col < start_pos + w:
                    res.append(matrix[row][col])
                    col += 1
                break
            while col < start_pos + w - 1:
                res.append(matrix[row][col])
                col += 1
            while row < start_pos + h - 1:
                res.append(matrix[row][col])
                row += 1
            while col > start_pos:
                res.append(matrix[row][col])
                col -= 1
            while row > start_pos:
                res.append(matrix[row][col])
                row -= 1
            start_pos += 1
            w -= 2
            h -= 2
        return res


class Solution_63:
    def dfs(self,matrix,row,column,target_row,target_col,dp):
        if row == target_row and column == target_col:
            return 1
        if row > target_row or column > target_col:
            return 0
        if matrix[row][column] == 1:
            return 0
        k = str(row)+","+ str(column)
        if k in dp:
            return dp[k]
        dp[k] = self.dfs(matrix,row + 1,column,target_row,target_col,dp) + self.dfs(matrix,row,column + 1,target_row,target_col,dp)
        return dp[k]

    def uniquePathsWithObstacles(self, obstacleGrid) -> int:
        if len(obstacleGrid[0]) == 0:
            return 1
        if obstacleGrid[0][0] == 1 or obstacleGrid[len(obstacleGrid) - 1][len(obstacleGrid[0]) - 1]:
            return 0
        dp = {}
        return self.dfs(obstacleGrid,0,0,len(obstacleGrid) - 1,len(obstacleGrid[0]) - 1,dp)

    def uniquePathsWithObstacles2(self, obstacleGrid) -> int:
        if len(obstacleGrid[0]) == 0:
            return 1
        if obstacleGrid[0][0] == 1 or obstacleGrid[len(obstacleGrid) - 1][len(obstacleGrid[0]) - 1]:
            return 0
        dp = []
        for i in range(len(obstacleGrid)):
            dp.append([0] * len(obstacleGrid[0]))

        dp[len(obstacleGrid) - 1][len(obstacleGrid[0])-1] = 1
        for row in range(len(obstacleGrid) - 1, - 1,-1):
            for col in range(len(obstacleGrid[0]) - 1, - 1,-1):
                if obstacleGrid[row][col] == 1:
                    dp[row][col] = 0
                else:
                    if col + 1 < len(obstacleGrid[0]):
                        dp[row][col] += dp[row][col + 1]
                    if row + 1 < len(obstacleGrid):
                        dp[row][col] +=  dp[row+1][col]
        return dp[0][0]

class Solution_1331:
    def arrayRankTransform(self, arr):
        record = {}
        for i in range(len(arr)):
            if arr[i] not in record:
                record[arr[i]] = [i]
            else:
                record[arr[i]].append(i)
        keys = record.keys()
        keys = sorted(keys)
        res = [0] * len(arr)
        i = 1
        for key in keys:
            for pos in record[key]:
                res[pos] = i
            i += 1
        return res
        #return [dict[key] for key in keys]


class Solution_443:
    def compress(self, chars) -> int:
        l = len(chars)
        if l == 1:
            return 1
        begin = 0
        end = begin + 1
        input_pos = begin
        while end <= l:
            if end == l or chars[begin] != chars[end]:
                gap = end - begin
                if gap >= 2:
                    input = []
                    while gap > 0:
                        input.append( gap % 10)
                        gap = gap // 10
                    input = input[::-1]
                    chars[input_pos] = chars[begin]
                    input_pos += 1
                    for v in input:
                        chars[input_pos] = str(v)
                        input_pos += 1
                else:
                    chars[input_pos] = chars[begin]
                    input_pos += 1
                begin = end
            end += 1
        chars[0:] = chars[0:input_pos]
        return len(chars)

class Solution_1342:
    def numberOfSteps (self, num: int) -> int:
        digits = f'{num:b}'
        return digits.count('1') - 1 + len(digits)
