# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Codec_297:

    def serialize(self, root):
        """Encodes a tree to a single string.

        :type root: TreeNode
        :rtype: str
        """

    def deserialize(self, data):
        """Decodes your encoded data to tree.

        :type data: str
        :rtype: TreeNode
        """


class Solution_236:
    def __init__(self):
        self.res = None
        self.memo_p = {}
        self.memo_q = {}

    def is_child(self,node,sub_node,memo)->bool:
        if node is None:
            return False
        if node.val in memo:
            return True
        res = node.val == sub_node.val or self.is_child(node.left,sub_node,memo) or self.is_child(node.right, sub_node,memo)
        if res:
            if node.val not in memo:
                memo[node.val] = True
        return res

    def dfs(self,node,p,q):
        p_is_child = self.is_child(node,p,self.memo_p)
        q_is_child = self.is_child(node,q,self.memo_q)
        if not p_is_child or not q_is_child:
            return
        if p_is_child and q_is_child:
            self.res = node

        self.dfs(node.left,p,q)
        self.dfs(node.right,p,q)

    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        self.res = root
        self.dfs(root,p,q)
        return self.res

def red(l,r):
    if l is not None:
        l.next = r
    return r
from functools import reduce

class Solution_116:
    def connect(self, root: 'Node') -> 'Node':
        cur = [root]
        while len(cur) > 0:
            reduce(red, cur)
            cur = [child for parent in cur if parent for child in [parent.left,parent.right]]

            # last = None
            # for ele in cur:
            #     if last is not None:
            #         last.next = ele
            #     last = ele
        return root

# Input: [5,2,6,1]
# Output: [2,1,1,0]
# Explanation:
# To the right of 5 there are 2 smaller elements (2 and 1).
# To the right of 2 there is only 1 smaller element (1).
# To the right of 6 there is 1 smaller element (1).
# To the right of 1 there is 0 smaller element.
class Solution_315:
    def bst(self,nums,val,node,step):
        if val > node.val:
            step += 1
            if None == node.right:
                node.right = TreeNode(val)
                return step
            else:
                return self.bst(nums,val,node.right,step)
        else:
            if None == node.left:
                node.left = TreeNode(val)
                return step
            else:
                return self.bst(nums,val,node.left,step)

    def countSmaller(self, nums):
        l = len(nums)
        if 0 == l:
            return []
        root = TreeNode(nums[l-1])
        i = l - 2
        res = []
        res.append(0)
        while i >= 0:
            s = self.bst(nums, nums[i],root,0)
            print(s)
            res.append(s)
            i -= 1
        return res[::-1]

# Input: 3
# Output:
# [
#   [1,null,3,2],
#   [3,2,null,1],
#   [3,1,null,null,2],
#   [2,1,3],
#   [1,null,2,null,3]
# ]
# Explanation:
# The above output corresponds to the 5 unique BST's shown below:
#
#    1         3     3      2      1
#     \       /     /      / \      \
#      3     2     1      1   3      2
#     /     /       \                 \
#    2     1         2                 3
class Solution_95:
    def dfs_pre(self,parent,root,nums,res):
        if len(nums) == 0:
            res.append(root)
        copy_nums = nums.copy()
        for n in copy_nums:
            if n < parent.val:
                self.dfs_pre()

    def generateTrees(self, n: int):
        res = []
        i = 1
        while i <= n:
            nums = [x for x in range(1,n+1)]
            nums.remove(i)
            root = TreeNode(i)
            self.dfs_pre(root,root,nums,res)
            i += 1
        return res
