from arraytype.array import *
from tree import *
from other import *
from Longest_Absolute_File_Path import *
from stringissue.Solution1147 import *

if __name__ == '__main__':
    s1147 = Solution_1147()
    text = "ghiabcdefhelloadamhelloabcdefghi"
    res = s1147.longestDecomposition(text)

    s388 = Solution_388()
    res = s388.lengthLongestPath("dir")

    s1342 = Solution_1342()
    res = s1342.numberOfSteps(99)

    s443 = Solution_443()
    # input =["a","b","b","b","b","b","b","b","b","b","b","b","b"]
    #input = ["a","a","b","b","c","c","c"]
    input = ["a","b","b","b","b","b","b","b","b","b","b","b","b","c","c"]
    res = s443.compress(input)

    s1331 = Solution_1331()
    res = s1331.arrayRankTransform([37,12,28,9,100,56,80,5,12])

    s54 = Solution_54()
    res = s54.spiralOrder([])

    s63 = Solution_63()
    s63.uniquePathsWithObstacles2([[0,0,0],[0,1,0],[0,0,0]])

    s59 = Solution_59()
    res = s59.generateMatrix(6)

    s60 = Solution_60()
    res = s60.getPermutation(9,331987)

    s77 = Solution_77()
    res = s77.combine(4,2)

    s80 = Solution_80()
    res = s80.removeDuplicates([0,0,0,0,0,1,1,1,2,2,3,3,5,5,6])

    s81 = Solution_81()
    res = s81.search([1,3,1,1,1],3)

    l1 = ListNode(1)
    l2 = ListNode(2)
    l31 = ListNode(3)
    l32 = ListNode(3)
    l41 = ListNode(4)
    l42 = ListNode(4)
    l5 = ListNode(5)
    l1.next = l2
    l2.next = l31
    l31.next = l32
    l32.next = l41
    l41.next = l42
    l42.next = l5

    s82 = Solution_82()
    res = s82.deleteDuplicates(l1)

    s89 = Solution_89()
    res = s89.grayCode(3)

    s91 = Solution_91()
    res = s91.numDecodings("1010")

    s90 = Solution_90()
    res = s90.subsetsWithDup([-1,1,2])

    s93 = Solution_93()
    res = s93.restoreIpAddresses("010010")

    l1 = ListNode(1)
    l2 = ListNode(2)
    l3 = ListNode(3)
    l4 = ListNode(4)
    l5 = ListNode(5)
    l1.next = l2
    l2.next = l3
    l3.next = l4
    l4.next = l5
    s92 = Solution_92()
    res = s92.reverseBetween(l1,2,5)

    s1184 = Solution_1184()
    res = s1184.distanceBetweenBusStops([1,2,3,4],0,1)

    s1232 = Solution_1232()
    res = s1232.checkStraightLine([[-1,0],[2,0],[3,0]])

    s40 = Solution_40()
    res = s40.combinationSum2([14,6,25,9,30,20,33,34,28,30,16,12,31,9,9,12,34,16,25,32,8,7,30,12,33,20,21,29,24,17,27,34,11,17,30,6,32,21,27,17,16,8,24,12,12,28,11,33,10,32,22,13,34,18,12],27)

    s1310 = Solution_1310()
    res = s1310.xorQueries([1,3,4,8],[[0,1],[1,2],[0,3],[3,3]])

    s33 = Solution_33()
    res = s33.search([3,1],1)

    l1 = ListNode(1)
    l2 = ListNode(2)
    l3 = ListNode(3)
    l4 = ListNode(4)
    l1.next = l2
    l2.next = l3
    l3.next = l4

    s25 = Solution_25()
    res = s25.reverseKGroup(l1,3)

    s24 = Solution_24()
    res = s24.swapPairs(l1)

    s16 = Solution_16()
    res = s16.threeSumClosest([1,2,4,8,16,32,64,128],82)

    s15 = Solution_15ex()
    res = s15.threeSum([6,-5,-6,-1,-2,8,-1,4,-10,-8,-10,-2,-4,-1,-8,-2,8,9,-5,-2,-8,-9,-3,-5])

    s15 = Solution_15()
    res = s15.threeSum([0,0,0,0,0,0,-1,-1,-1,1,1,1])

    s5 = Solution_5()
    res = s5.longestPalindrome("abcccddddcc")
    print(res)

    s315 = Solution_315()
    res = s315.countSmaller([2,0,1])

    s239 = Solution_239()
    s239.maxSlidingWindow([],3)
    l = [1,2,3]
    res = [[i*i,i*i*i] for i in l]

    #[2,3,4,6,6,9]
    s1309 =  Solution_1309()
    input = "10#11#12"
    res = s1309.freqAlphabets(input)

    input = [[9,9,4],[6,6,8],[2,1,1]]
    s329 = Solution_329()
    ret = s329.longestIncreasingPath(input)

    board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']]
    word = 'ABCCED'
    s79 = Solution_79()
    ret = s79.exist(board,word)

    s1190 = Solution_1190()
    ret = s1190.reverseParentheses("")

    s894 = Solution_894()
    ret = s894.allPossibleFBT(7)

    s605 = Solution_605()
    input = [0,1,0]
    #[0,0,0,0,1,0,0,0,1,0,0]
    ret = s605.canPlaceFlowers(input,4)
    s69 = Solution_69()
    s69.mySqrt(36)

    s219 = Solution_219()
    s219.containsNearbyDuplicate([1,0,1,1],1)

    s172 = Solution_172()
    s172.trailingZeroes(17)

    s438 = Solution_438()
    s438.findAnagrams("cbaebabacd","abc")

    s1128 = Solution_1128()
    s1128.numEquivDominoPairs([[1, 2], [2, 1], [3, 4], [5, 6]])

    s645 = Solution_645()
    s645.findErrorNums([8,7,3,5,3,6,1,4])
    #s645.findErrorNums([8, 7, 6,5,4,3,2,1])

    s720 = Solution_720()
    res = s720.longestWord(["yo","ew","fc","zrc","yodn","fcm","qm","qmo","fcmz","z","ewq","yod","ewqz","y"])

    s1018 = Solution_1018()
    s1018.prefixesDivBy5([1,1,0,1])

    s1071 = Solution_1071()
    res = s1071.gcdOfStrings("leet","code")

    [5, 5, 5, 10, 5, 5, 10, 20, 20, 20]

    l = [1,2,2,3,1,4,2]
    s697 = Solution_697()
    res = s697.findShortestSubArray(l)
    # s0 = Solution_array()
    # s0.floodFill()

    s = Solution()
    res = s.twoCitySchedCost([[518,518],[71,971],[121,862],[967,607],[138,754],[513,337],[499,873],[337,387],[647,917],[76,417]])

    data = [1, 2, 3, 4, 5, 6]
    for i in range (1,7):
        combine(data, 0, [], i)

    canThreePartsEqualSum([1,2,3,1,2])


