class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


def canThreePartsEqualSum(A) -> bool:
    total = sum(A)
    if total % 3 != 0:
        return False
    target = total / 3
    count = 0
    s = 0
    for i in A:
        s += i
        if s == target:
            count += 1
            if count == 3:
                return True
            s = 0
    return False

def combine(data:list, step:int, select_data:list, target_num:int):
  if len(select_data) == target_num:   #选择的元素已经够了，就输出并返回
      print(select_data)
      return
  if step >= len(data):               #没有元素选了而且还没够，也是直接返回
      return
  select_data.append(data[step])             #选择当前元素
  combine(data, step + 1, select_data, target_num)
  select_data.pop()                         #别忘了从已选择元素中先删除
  combine(data, step + 1, select_data, target_num) #不选择当前元素


def mycmp(x, y):
    if (x[0] - x[1]) < (y[0] - y[1]):
        return -1
    elif (x[0] - x[1]) == (y[0] - y[1]):
        return 0
    else:
        return 1


class Solution(object):
    def twoCitySchedCost(self, costs) -> int:
        new_costs = sorted(costs,key= cmp_to_key(mycmp))
        res = 0
        for index, item in enumerate(new_costs):
            if index < len(costs)/2:
                res += item[0]
            else:
                res += item[1]
        return res

# Given 1->2->3->4, you should return the list as 2->1->4->3.
class Solution_24:
    def swapPairs(self, head: ListNode) -> ListNode:
        if head == None or head.next == None:
            return head
        s = []
        move = head
        cnt = 0
        while move != None:
            if cnt % 2 == 0:
                s.insert(cnt,move)
            else:
                s.insert(cnt - 1,move)
            cnt += 1
            move = move.next
        i = 0
        res = None
        while i < cnt:
            if None == res:
                res = s[i]
                move = s[i]
            else:
                move.next = s[i]
                move = move.next
            i += 1
        move.next = None
        return res


class Solution_25:
    def reverseKGroup(self, head, k: int) -> ListNode:
        if head == None or head.next == None or k == 0:
            return head
        s = []
        move = head
        cnt = 0
        i = 0
        while move != None:
            s.insert(cnt,move)
            cnt += 1
            if cnt % k == 0:
                l = s[cnt - k:cnt]
                l.reverse()
                s[cnt - k:cnt] = l
            move = move.next
        i = 0
        res = None
        while i < cnt:
            if None == res:
                res = s[i]
                move = s[i]
            else:
                move.next = s[i]
                move = move.next
            i += 1
        move.next = None
        return res

# class Solution_23:
#     def mergelist(self,list1,list2):
#         pass
#     def mergeKLists(self, lists):
#         first = 0
#         pass

# Input: nums = [4,5,6,7,0,1,2], target = 0
# Output: 4
# Input: nums = [4,5,6,7,0,1,2], target = 3
# Output: -1
class Solution_33:
    def search(self, nums, target: int) -> int:
        l = len(nums)
        low = 0
        high = l - 1
        while low <= high:
            mid = (low + high) // 2
            if nums[mid] == target:
                return mid

            if nums[mid] <= nums[high]:
                if nums[mid] < target <= nums[high]:
                    low = mid + 1
                else:
                    high = mid - 1
            else:
                if nums[low] <= target < nums[mid]:
                    high = mid - 1
                else:
                    low = mid + 1
        return -1

# You are given an arraytype coordinates, coordinates[i] = [x, y], where [x, y] represents the coordinate of a point.
# Check if these points make a straight line in the XY plane.
class Solution_1232:
    #def checkStraightLine(self, coordinates: List[List[int]]) -> bool:
    def checkStraightLine(self, coordinates) -> bool:
        l = len(coordinates)
        if l <= 2:
            return True
        # y = ax + b
        if coordinates[1][0] != coordinates[0][0]:
            a = (coordinates[1][1] - coordinates[0][1])/(coordinates[1][0] - coordinates[0][0])
            b = 1/2 * (coordinates[1][1] + coordinates[0][1] - \
                (coordinates[1][1] - coordinates[0][1])*(coordinates[1][0] + coordinates[0][0])/(coordinates[1][0] - coordinates[0][0]))
            i = 2
            while i < l:
                if a * (coordinates[i][0] - coordinates[i-1][0]) != (coordinates[i][1] - coordinates[i-1][1]):
                    return False
                i += 1
            return True
        else:
            i = 2
            while i < l:
                if coordinates[i][0] != coordinates[i-1][0]:
                    return False
                i += 1
            return True


class Solution_92:
    def reverseBetween(self, head: ListNode, m: int, n: int) -> ListNode:
        if m == n:
            return head
        m_node = head
        i = 1
        while i != m:
            m_node = m_node.next
            i += 1
        n_node = m_node
        val_list = [m_node.val]
        while i != n:
            n_node = n_node.next
            val_list.append(n_node.val)
            i += 1
        j = n - m + 1
        i = m
        while i != n:
            m_node.val = val_list[j-1]
            m_node = m_node.next
            i += 1
            j -= 1
        m_node.val = val_list[0]
        return head


class Solution_93:
    def dfs(self,s,cur_pos,step,sub_list,res):
        if step == 4:
            if cur_pos >= len(s):
                res['.'.join(sub_list)] = True
            return
        l = len(s)
        rest_len = l - cur_pos
        if cur_pos >= l or rest_len /(4-step) < 1 or rest_len /(4 - step) > 3:
            return

        for i in range(1,4):
            if i >= 2:
                if s[cur_pos] == '0':
                    return
            if int(s[cur_pos:cur_pos + i]) <= 255:
                list = sub_list.copy()
                list.append(s[cur_pos:cur_pos + i])
                self.dfs(s,cur_pos + i,step + 1,list,res)

    def restoreIpAddresses(self, s: str):
        l = len(s)
        res = {}
        self.dfs(s,0,0,[],res)
        return [ele for ele in res.keys()]


class Solution_91:
    def dfs(self,s,cur_pos,dp):
        if cur_pos < len(s) and s[cur_pos] == "0":
            return -2e31
        if cur_pos >= (len(s) - 1):
            return 1
        if cur_pos in dp:
            return dp[cur_pos]
        cnt_steptwo = 0
        if (cur_pos + 1) < len(s):
            val = int(s[cur_pos:cur_pos+2])
            if val <= 26:
                cnt_steptwo = self.dfs(s,cur_pos + 2,dp)
        cnt_stepone = self.dfs(s,cur_pos + 1,dp)
        if cnt_steptwo < 0:
            cnt_steptwo = 0
        if cnt_stepone < 0:
            cnt_stepone = 0
        dp[cur_pos] = cnt_stepone + cnt_steptwo
        return cnt_stepone + cnt_steptwo

    def is_valid(self,s):
        l = len(s)
        if l == 0 or s[0] == "0":
            return False
        i = 1
        while i < l:
            if s[i-1] == "0" and s[i] == "0":
                return False
            i += 1
        return True

    def numDecodings(self, s: str) -> int:
        if not self.is_valid(s):
            return 0
        dp = {}
        return self.dfs(s,0,dp)

# Input: head = 1->4->3->2->5->2, x = 3
# Output: 1->2->2->4->3->5
class Solution_86:
    def partition(self, head: ListNode, x: int) -> ListNode:
        less_tree_head = None
        less_visit = None
        bigger_tree_head = None
        bigger_visit = None
        visit = head
        while visit is not None:
            if visit.val < x:
                if less_tree_head is None:
                    less_tree_head = ListNode(visit.val)
                    less_visit = less_tree_head
                else:
                    node = ListNode(visit.val)
                    less_visit.next = node
                    less_visit = node
            else:
                if bigger_tree_head is None:
                    bigger_tree_head = ListNode(visit.val)
                    bigger_visit = bigger_tree_head
                else:
                    node = ListNode(visit.val)
                    bigger_visit.next = node
                    bigger_visit = node
            visit = visit.next

        if less_visit is not None:
            less_visit.next = bigger_tree_head
            return less_tree_head
        else:
            return bigger_tree_head

# Input: 2
# Output: [0,1,3,2]
# Explanation:
# 00 - 0
# 01 - 1
# 11 - 3
# 10 - 2
class Solution_89:
    def grayCode(self, n: int):
        res = [0]
        for i in range(n):
            l = len(res)
            bit = 1 << i
            for pos in range(l):
                res.append(res[l - 1 -  pos] | bit)
        return res


class Solution_82:
    def deleteDuplicates(self, head: ListNode):
        if head is None:
            return head
        record = {}
        cur = head
        last = None
        while cur:
            if cur.val in record:
                record[cur.val] += 1
            else:
                record[cur.val] = 1
            last = cur
            cur = cur.next
        nums = [key for key in record.keys() if record[key] == 1]
        is_reverse = False
        if len(nums) >= 2:
            is_reverse = head.val > last.val
        nums = sorted(nums,reverse = is_reverse)
        root = None
        move = None
        for i in range(len(nums)):
            n = ListNode(nums[i])
            if root is None:
                root = n
                move = n
            else:
                move.next = n
                move = move.next
        return root

# Input: nums = [2,5,6,0,0,1,2], target = 0
# Output: true
# Input: nums = [2,5,6,0,0,1,2], target = 3
# Output: false
class Solution_81:
    def search(self, nums, target: int) -> bool:
        low = 0
        high = len(nums) - 1
        while low <= high:
            mid = (low + high)//2
            if nums[mid] == target:
                return True

            if nums[low] == nums[mid] and nums[high] == nums[mid]:
                while low <= high:
                    if nums[low] == target:
                        return True
                    low += 1
                return False

            if nums[low] <= nums[mid]:
                if nums[low] <= target <= nums[mid]:
                    high = mid - 1
                else:
                    low = mid + 1
            else:
                if nums[mid] <= target <= nums[high]:
                    low = mid + 1
                else:
                    high = mid - 1
        return False


# Input: 0->1->2->NULL, k = 4
# Output: 2->0->1->NULL
# Explanation:
# rotate 1 steps to the right: 2->0->1->NULL
# rotate 2 steps to the right: 1->2->0->NULL
# rotate 3 steps to the right: 0->1->2->NULL
# rotate 4 steps to the right: 2->0->1->NULL
class Solution_61:
    def rotateRight(self, head: ListNode, k: int) -> ListNode:
        l = 0
        v = head
        while v:
            v = v.next
            l += 1
        if l == 0:
            return head
        if k % l == 0:
            return head
        cnt = k % l
        v = head
        while v and l - cnt - 1:
            v = v.next
            cnt += 1
        new_head = v.next
        v.next = None
        move = new_head
        while move.next:
            move = move.next
        move.next = head
        return new_head

class Solution_1969:
    def minNonZeroProduct(self, p: int) -> int:
        biggest = 2 ** p - 1
        pair = 2 ** p - 2
        times = 2 ** (p - 1) - 1
        return (biggest * pow(pair, times, 10 ** 9 + 7)) % (10 ** 9 + 7)