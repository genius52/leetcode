package main

import (
	"bytes"
	"container/heap"
	"container/list"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main(){

	{
		input := []int{-2,-3,-1}
		res := maxSubarraySumCircular(input)
		fmt.Println(res)
	}
	{
		input := []int{-2,1,-3,4,-1,2,1,-5,4}
		res := maxSubArray(input)
		fmt.Println(res)
	}
	{
		input := []int{4,4,4,4}
		k := 4
		threshold := 1
		res := numOfSubarrays(input,k,threshold)
		fmt.Println(res)
	}
	{
		pre := []int{1,2,4,5,3,6,7}
		post := []int{4,5,2,6,7,3,1}
		res := constructFromPrePost(pre,post)
		fmt.Println(res)
	}
	{
		input := []string{ "//", "/ "}
		res := regionsBySlashes(input)
		fmt.Println(res)
	}
	{
		input := []int{4,6,7,7}
		res := findSubsequences(input)
		fmt.Println(res)
	}
	{
		input := []int{3,2,4,1}
		res := pancakeSort(input)
		fmt.Println(res)
	}
	{
		var t1 TreeNode
		t1.Val = 1
		var t2 TreeNode
		t2.Val = 2
		var t3 TreeNode
		t3.Val = 3
		var t4 TreeNode
		t4.Val = 4
		//var t5 TreeNode
		//t5.Val = 5
		t1.Left = &t2
		t1.Right = &t3
		t2.Left = &t4
		//t2.Right = &t5
		res := lcaDeepestLeaves(&t1)
		fmt.Println(res)
	}
	{
		input := [][]int{{5,6},{1,0}}
		res := minFlips(input)
		fmt.Println(res)
	}
	{
		input := "1-2--3---4-5--6---7"
		res := recoverFromPreorder(input)
		fmt.Println(res)
	}
	{
		res := fractionToDecimal(2,4)
		fmt.Println(res)
	}
	{
		res := myPow(3,4)
		fmt.Println(res)
	}
	{
		//input := []int{1, 3, 2, 1, 3, 2}
		//wiggleSort(input)
		//fmt.Println("over")
	}
	{
		s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		wordDict := []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}
		res := wordBreak2(s,wordDict)
		fmt.Println(res)
	}
	{
		input := "cbcccbbccbbacbcaccbacabacacbcccbbccbbacb"
		res := longestPrefix(input)
		fmt.Println(res)
	}
	{

		board := [][]byte{{'a'}}
		words := []string{"a"}
		res := findWords2(board,words)
		fmt.Println(res)
	}
	{
		var input [][]int = [][]int{{1,0,0,0},{0,0,0,0},{0,0,0,0},{0,0,0,0},{0,0,0,2}}
		res := uniquePathsIII(input)
		fmt.Println(res)
	}
	{
		s := "ab"
		p := ".*"
		res := isMatch(s,p)
		fmt.Println(res)
	}
	{
		input := ")()())"
		res := longestValidParentheses(input)
		res = dp_longestValidParentheses(input)
		fmt.Println(res)
	}
	{
		input := []int{3,4,-1,1}
		res := firstMissingPositive(input)
		fmt.Println(res)
	}
	{
		var t1 TreeNode
		t1.Val = -2
		var t2 TreeNode
		t2.Val = 6
		var t3 TreeNode
		t3.Val = 0
		var t4 TreeNode
		t4.Val = -6
		t1.Left = &t2
		t2.Left = &t3
		t2.Right = &t4
		res := maxPathSum(&t1)
		fmt.Println(res)
	}
	{
		lo := 1
		hi := 1000
		k := 777
		res := getKth(lo,hi,k)
		fmt.Println(res)
	}
	{
		arr1 := []int{4,5,8}
		arr2 := []int{10,9,1,8}
		d := 2
		res := findTheDistanceValue(arr1,arr2,d)
		fmt.Println(res)
	}
	{
		s := "a"
		t := "aa"
		res := minWindow(s,t)
		fmt.Println(res)
	}
	{
		input := []int{4,2,0,3,2,5}
		res := largestRectangleArea(input)
		fmt.Println(res)
	}
	{
		input := []int{73, 74, 75, 71, 69, 72, 76, 73}
		res := dailyTemperatures(input)
		fmt.Println(res)
	}
	{
		input := []int{2,0,1}
		res := countSmaller(input)
		fmt.Println(res)
	}
	{
		input := "ababacb"
		k := 3
		res := longestSubstring(input,k)
		fmt.Println(res)
	}
	{
		input := []string{"2", "1", "+", "3", "*"}
		res := evalRPN(input)
		fmt.Println(res)
	}
	{
		input := "ac"
		res := longestPalindrome(input)
		fmt.Println(res)
	}
	{
		input := "abcabcabc"
		res := lengthOfLongestSubstring(input)
		fmt.Println(res)
	}
	{
		//tree_nums :=0
		//var trees []int
		//fmt.Scan(&tree_nums)
		//for i := 0; i < tree_nums; i++ {
		//	x:=0
		//	fmt.Scan(&x)
		//	trees = append(trees,x)
		//}
		//var hours_limit int
		//fmt.Scan(&hours_limit)
		//
		//var total int
		//for _,n := range trees{
		//	total += n
		//}
		//var least_eat_speed int = int(math.Ceil(float64(total) / float64(hours_limit)))
		//var cost_times int = hours_limit + 1
		//for cost_times > hours_limit{
		//	cost_times = 0
		//	for i := 0;i < len(trees);i++{
		//		var peachs int = trees[i]
		//		for peachs > 0{
		//			cost_times++
		//			peachs -= least_eat_speed
		//		}
		//	}
		//	if cost_times <= hours_limit{
		//		break
		//	}
		//	least_eat_speed++
		//}
		//fmt.Println(least_eat_speed)
	}
	{
		input := [][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}
		res := maximalRectangle(input)
		fmt.Println(res)
	}
	{
		s := "()())()"
		res := removeInvalidParentheses(s)
		fmt.Println(res)
	}
	{
		//word1 := "horse"
		//word2 := "ros"
		word1 := "ab"
		word2 := "c"
		res := minDistance(word1,word2)
		fmt.Println(res)
	}
	{
		var t1 TreeNode
		t1.Val = 1
		var t2 TreeNode
		t2.Val = 2
		var t3 TreeNode
		t3.Val = 3
		var t4 TreeNode
		t4.Val = 4
		var t5 TreeNode
		t5.Val = 5
		t1.Right = &t2
		t2.Right = &t3
		t3.Right = &t4
		t4.Right = &t5
		res := balanceBST(&t1)
		fmt.Println(res.Val)
	}
	{
		var customStack CustomStack = Constructor(3) // Stack is Empty []
		customStack.Push(1)                          // stack becomes [1]
		customStack.Push(2)                          // stack becomes [1, 2]
		res := customStack.Pop()                            // return 2 --> Return top of the stack 2, stack becomes [1]
		customStack.Push(2)                         // stack becomes [1, 2]
		customStack.Push(3)                          // stack becomes [1, 2, 3]
		customStack.Push(4)                          // stack still [1, 2, 3], Don't add another elements as size is 4
		customStack.Increment(5, 100)               // stack becomes [101, 102, 103]
		customStack.Increment(2, 100)                // stack becomes [201, 202, 103]
		res = customStack.Pop()                            // return 103 --> Return top of the stack 103, stack becomes [201, 202]
		res = customStack.Pop()                            // return 202 --> Return top of the stack 102, stack becomes [201]
		res = customStack.Pop()                            // return 201 --> Return top of the stack 101, stack becomes []
		res = customStack.Pop()                            // return -1 --> Stack is empty return -1.
		fmt.Println(res)
	}
	{
		input := [][]int{{7,8},{1,2}}
		res := luckyNumbers(input)
		fmt.Println(res)
	}
	{
		nums := 5
		input := [][]int{{1,0},{2,0},{3,1},{3,2}}
		res := findOrder(nums,input)
		fmt.Println(res)
	}
	{
		courses := 2
		input := [][]int{{1,0}}
		res := canFinish(courses,input)
		fmt.Println(res)
	}
	{
		input := [][]int{{1,2,3},{4,5,6},{7,8,9}}
		res := matrixBlockSum(input,1)
		fmt.Println(res)
	}
	{
		res := allPossibleFBT(7)
		fmt.Println(len(res))
	}
	{
		input := [][]int{{3,3,1,1},{2,2,1,2},{1,1,1,2}}
		res := diagonalSort(input)
		fmt.Println(res)
	}
	{
		input := []int{3,3,3,3,3,1,3}
		res := groupThePeople(input)
		fmt.Println(res)
	}
	{
		//[
		//	1->4->5,
		//	1->3->4,
		//	2->6
		//]
		var lists []*ListNode
		var l1 ListNode
		l1.Val = 1
		var l11 ListNode
		l11.Val = 4
		var l12 ListNode
		l12.Val = 5
		l1.Next = &l11
		l11.Next = &l12
		lists = append(lists, &l1)

		var l2 ListNode
		l2.Val = 1
		var l21 ListNode
		l21.Val = 3
		var l22 ListNode
		l22.Val = 4
		l2.Next = &l21
		l21.Next = &l22
		lists = append(lists,&l2)

		var l3 ListNode
		l3.Val = 2
		var l31 ListNode
		l31.Val = 6
		l3.Next = &l31
		lists = append(lists,&l3)
		res := mergeKLists(lists)
		fmt.Println(res)
	}
	{
		res := generateTheString(4)
		fmt.Println(res)
	}
	{
		s := ""
		res := sortString(s)
		fmt.Println(res)
	}
	{
		s := "(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())"
		res := checkValidString(s)
		fmt.Println(res)
	}
	{
		obj := Constructor933();
		res := obj.Ping(1);
		fmt.Println(res)
		res = obj.Ping(100);
		fmt.Println(res)
		res = obj.Ping(3001);
		fmt.Println(res)
		res = obj.Ping(3002);
		fmt.Println(res)
		res = obj.Ping(3007);
		fmt.Println(res)
		res = obj.Ping(3120);
		fmt.Println(res)
	}
	{
		input := []int{1,6,5,4,3,2,1}
		res := findPeakElement(input)
		fmt.Println(res)
	}
	{
		input := " 3+5 / 2 - 3"
		res := calculate(input)
		fmt.Println(res)
	}
	{
		nums := []int{30,1,20,21,15,19,2,-1,3}
		res := increasingTriplet(nums)
		fmt.Println(res)
	}
	{
		input := []int{3,1,5,8}
		res := maxCoins(input)
		fmt.Println(res)
	}
	{
		//input := []int{3,3,3,3,5,5,5,2,2,7}
		//input := []int{7,7,7,7,7,7}
		input := []int{1,2,3,4,5,6,7,8,9,10}
		res := minSetSize(input)
		fmt.Println(res)
	}
	{
		res := closestDivisors(208656121)
		fmt.Println(res)
	}
	{
		n := 4
		left := []int{1,-1,3,-1}
		right := []int{2,-1,-1,-1}
		res := validateBinaryTreeNodes(n,left,right)
		fmt.Println(res)
	}
	{
		input := []int{10,100,1000,10000}
		res := sortByBits(input)
		fmt.Println(res)
	}
	{
		num := "112"
		k := 1
		res := removeKdigits(num,k)
		fmt.Println(res)
	}
	{
		//["LRUCache","put","put","put","put","get","get"]
		//[[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
		obj := Constructor146(2);
		obj.Put(2, 1);
		obj.Put(2, 2);
		res := obj.Get(2);
		obj.Put(1, 1);
		obj.Put(4, 1);
		res = obj.Get(2);

		//obj.Put(2,1)
		//obj.Put(1,1)
		//obj.Put(2,3)
		//obj.Put(4,1)
		//res := obj.Get(1);
		//res = obj.Get(2);
		fmt.Println(res)
	}
	{
		//s := "catsandog"
		//wordDict := []string{"cats", "dog", "sand", "and", "cat"}
		s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
		wordDict := []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}
		res := wordBreak(s,wordDict)
		fmt.Println(res)
	}
	{
		//-1->5->3->4->0
		var l1 ListNode
		l1.Val = 1
		var l2 ListNode
		l2.Val = 5
		var l3 ListNode
		l3.Val = 3
		var l4 ListNode
		l4.Val = 4
		var l5 ListNode
		l5.Val = 0
		l1.Next = &l2
		l2.Next = &l3
		l3.Next = &l4
		l4.Next = &l5
		res := sortList(&l1)
		fmt.Println(res)
	}
	{
		input := []int{1,2,3,4}
		res := decompressRLElist(input)
		fmt.Println(res)
	}
	{
		input := []int{70,4,83,56,94,72,78,43,2,86,65,100,94,56,41,66,3,33,10,3,45,94,15,12,78,60,58,0,58,15,21,7,11,41,12,96,83,77,47,62,27,19,40,63,30,4,77,52,17,57,21,66,63,29,51,40,37,6,44,42,92,16,64,33,31,51,36,0,29,95,92,35,66,91,19,21,100,95,40,61,15,83,31,55,59,84,21,99,45,64,90,25,40,6,41,5,25,52,59,61,51,37,92,90,20,20,96,66,79,28,83,60,91,30,52,55,1,99,8,68,14,84,59,5,34,93,25,10,93,21,35,66,88,20,97,25,63,80,20,86,33,53,43,86,53,55,61,77,9,2,56,78,43,19,68,69,49,1,6,5,82,46,24,33,85,24,56,51,45,100,94,26,15,33,35,59,25,65,32,26,93,73,0,40,92,56,76,18,2,45,64,66,64,39,77,1,55,90,10,27,85,40,95,78,39,40,62,30,12,57,84,95,86,57,41,52,77,17,9,15,33,17,68,63,59,40,5,63,30,86,57,5,55,47,0,92,95,100,25,79,84,93,83,93,18,20,32,63,65,56,68,7,31,100,88,93,11,43,20,13,54,34,29,90,50,24,13,44,89,57,65,95,58,32,67,38,2,41,4,63,56,88,39,57,10,1,97,98,25,45,96,35,22,0,37,74,98,14,37,77,54,40,17,9,28,83,13,92,3,8,60,52,64,8,87,77,96,70,61,3,96,83,56,5,99,81,94,3,38,91,55,83,15,30,39,54,79,55,86,85,32,27,20,74,91,99,100,46,69,77,34,97,0,50,51,21,12,3,84,84,48,69,94,28,64,36,70,34,70,11,89,58,6,90,86,4,97,63,10,37,48,68,30,29,53,4,91,7,56,63,22,93,69,93,1,85,11,20,41,36,66,67,57,76,85,37,80,99,63,23,71,11,73,41,48,54,61,49,91,97,60,38,99,8,17,2,5,56,3,69,90,62,75,76,55,71,83,34,2,36,56,40,15,62,39,78,7,37,58,22,64,59,80,16,2,34,83,43,40,39,38,35,89,72,56,77,78,14,45,0,57,32,82,93,96,3,51,27,36,38,1,19,66,98,93,91,18,95,93,39,12,40,73,100,17,72,93,25,35,45,91,78,13,97,56,40,69,86,69,99,4,36,36,82,35,52,12,46,74,57,65,91,51,41,42,17,78,49,75,9,23,65,44,47,93,84,70,19,22,57,27,84,57,85,2,61,17,90,34,49,74,64,46,61,0,28,57,78,75,31,27,24,10,93,34,19,75,53,17,26,2,41,89,79,37,14,93,55,74,11,77,60,61,2,68,0,15,12,47,12,48,57,73,17,18,11,83,38,5,36,53,94,40,48,81,53,32,53,12,21,90,100,32,29,94,92,83,80,36,73,59,61,43,100,36,71,89,9,24,56,7,48,34,58,0,43,34,18,1,29,97,70,92,88,0,48,51,53,0,50,21,91,23,34,49,19,17,9,23,43,87,72,39,17,17,97,14,29,4,10,84,10,33,100,86,43,20,22,58,90,70,48,23,75,4,66,97,95,1,80,24,43,97,15,38,53,55,86,63,40,7,26,60,95,12,98,15,95,71,86,46,33,68,32,86,89,18,88,97,32,42,5,57,13,1,23,34,37,13,65,13,47,55,85,37,57,14,89,94,57,13,6,98,47,52,51,19,99,42,1,19,74,60,8,48,28,65,6,12,57,49,27,95,1,2,10,25,49,68,57,32,99,24,19,25,32,89,88,73,96,57,14,65,34,8,82,9,94,91,19,53,61,70,54,4,66,26,8,63,62,9,20,42,17,52,97,51,53,19,48,76,40,80,6,1,89,52,70,38,95,62,24,88,64,42,61,6,50,91,87,69,13,58,43,98,19,94,65,56,72,20,72,92,85,58,46,67,2,23,88,58,25,88,18,92,46,15,18,37,9,90,2,38,0,16,86,44,69,71,70,30,38,17,69,69,80,73,79,56,17,95,12,37,43,5,5,6,42,16,44,22,62,37,86,8,51,73,46,44,15,98,54,22,47,28,11,75,52,49,38,84,55,3,69,100,54,66,6,23,98,22,99,21,74,75,33,67,8,80,90,23,46,93,69,85,46,87,76,93,38,77,37,72,35,3,82,11,67,46,53,29,60,33,12,62,23,27,72,35,63,68,14,35,27,98,94,65,3,13,48,83,27,84,86,49,31,63,40,12,34,79,61,47,29,33,52,100,85,38,24,1,16,62,89,36,74,9,49,62,89}
		res := maxProfit3(input)
		fmt.Println(res)
	}
	{
		//input := []byte{'A','A','A','A','A','A','B','C','D','E','F','G'}
		//intervel := 2
		//res := leastInterval(input,intervel)
		//fmt.Println(res)
	}
	{
		input := []int{0,0,0,0,0,0,0,0,0,0}
		target := 0
		res := subarraySum(input,target)
		fmt.Println(res)
	}
	{
		input := 3
		res := numberOfSteps(input)
		fmt.Println(res)
	}
	{
		input := []int{2,3,-2,4,-2,4,6,-9,3}
		res := maxProduct(input)
		fmt.Println(res)
	}
	{
		input := "3[a]2[bc]"
		//res := decodeString(input)
		res := decodeString2(input)
		fmt.Println(res)
	}
	{
		input := [][]int{{1,1,0,0,0},{1,1,1,1,0},{1,0,0,0,0},{1,1,0,0,0},{1,1,1,1,1}}
		k := 3
		res := kWeakestRows(input,k)
		fmt.Println(res)
	}
	{
		input := []int{1,1,1,1,1}
		target := 3
		res := findTargetSumWays(input,target)
		fmt.Println(res)
	}
	{
		//    1
		//   / \
		//  2   5
		// / \   \
		//3   4   6
		var t1 TreeNode
		t1.Val = 1
		var t2 TreeNode
		t2.Val = 2
		var t3 TreeNode
		t3.Val = 3
		var t4 TreeNode
		t4.Val = 4
		var t5 TreeNode
		t5.Val = 5
		var t6 TreeNode
		t6.Val = 6
		t1.Left = &t2
		t1.Right = &t5
		t2.Left = &t3
		t2.Right = &t4
		t5.Right = &t6
		flatten(&t1)
		fmt.Println(t1)
	}
	{
		inorder := []int{9,3,15,20,7}
		postorder := []int{9,15,7,20,3}
		res := buildTree(inorder,postorder)
		fmt.Println(res)
	}
	{
		//[-10,-3,0,5,9]
		var l1 ListNode
		l1.Val = -10
		var l2 ListNode
		l2.Val = -3
		var l3 ListNode
		l3.Val = 0
		var l4 ListNode
		l4.Val = 5
		var l5 ListNode
		l5.Val = 9
		l1.Next = &l2
		l2.Next = &l3
		l3.Next = &l4
		l4.Next = &l5
		res := sortedListToBST(&l1)
		fmt.Println(res)
	}
	{
		input := []byte{'a','a','b','b','c','c','c'}
		res := compress(input)
		fmt.Println(res)
		fmt.Println(input)
	}
	{
		input := [][]int{{0,0},{1,0},{2,0}}
		res := numberOfBoomerangs(input)
		fmt.Println(res)
	}
	{
		secret := "1123"
		guess := "0111"
		res := getHint(secret,guess)
		fmt.Println(res)
	}
	{
		//res := countAndSay(90)
		//fmt.Println(res)
	}
	{
		input := []int{37,12,28,56,12}
		res := arrayRankTransform(input)
		fmt.Println(res)
	}
	{
		//[186,419,83,408]
		//6249
		coins := []int{186,419,83,408}
		target := 6249
		//coins :=  []int{3,7,405,436}
		//target := 8839
		res := coinChange(coins,target)
		//res = coinChange2(coins,target)
		fmt.Println(res)
	}
	{
		res := getRow(5)
		fmt.Println(res)
	}
	{
		A := "aa"
		B := "aaaaa"
		res := repeatedStringMatch(A,B)
		fmt.Println(res)
	}
	{
		res := countBinarySubstrings("00110011")
		fmt.Println(res)
	}
	{
		var input [][]int = [][]int{{2,1,1},{1,1,0},{0,1,1}}
		res := orangesRotting(input)
		fmt.Println(res)
	}
	{
		res := tribonacci(25)
		fmt.Println(res)
	}
	{
		res := maximum69Number(996)
		fmt.Println(res)
	}
	{
		beginWord := "nanny"
		endWord :=  "aloud"
		wordList := []string{"ricky","grind","cubic","panic","lover","farce","gofer","sales","flint","omens","lipid","briny","cloth","anted","slime","oaten","harsh","touts","stoop","cabal","lazed","elton","skunk","nicer","pesky","kusch","bused","kinda","tunis","enjoy","aches","prowl","babar","rooms","burst","slush","pines","urine","pinky","bayed","mania","light","flare","wares","women","verne","moron","shine","bluer","zeros","bleak","brief","tamra","vasts","jamie","lairs","penal","worst","yowls","pills","taros","addle","alyce","creep","saber","floyd","cures","soggy","vexed","vilma","cabby","verde","euler","cling","wanna","jenny","donor","stole","sakha","blake","sanes","riffs","forge","horus","sered","piked","prosy","wases","glove","onset","spake","benin","talks","sites","biers","wendy","dante","allan","haven","nears","shaka","sloth","perky","spear","spend","clint","dears","sadly","units","vista","hinds","marat","natal","least","bough","pales","boole","ditch","greys","slunk","bitch","belts","sense","skits","monty","yawns","music","hails","alien","gibes","lille","spacy","argot","wasps","drubs","poops","bella","clone","beast","emend","iring","start","darla","bells","cults","dhaka","sniff","seers","bantu","pages","fever","tacky","hoses","strop","climb","pairs","later","grant","raven","stael","drips","lucid","awing","dines","balms","della","galen","toned","snips","shady","chili","fears","nurse","joint","plump","micky","lions","jamal","queer","ruins","frats","spoof","semen","pulps","oldie","coors","rhone","papal","seals","spans","scaly","sieve","klaus","drums","tided","needs","rider","lures","treks","hares","liner","hokey","boots","primp","laval","limes","putts","fonda","damon","pikes","hobbs","specs","greet","ketch","braid","purer","tsars","berne","tarts","clean","grate","trips","chefs","timex","vicky","pares","price","every","beret","vices","jodie","fanny","mails","built","bossy","farms","pubic","gongs","magma","quads","shell","jocks","woods","waded","parka","jells","worse","diner","risks","bliss","bryan","terse","crier","incur","murky","gamed","edges","keens","bread","raced","vetch","glint","zions","porno","sizes","mends","ached","allie","bands","plank","forth","fuels","rhyme","wimpy","peels","foggy","wings","frill","edgar","slave","lotus","point","hints","germs","clung","limed","loafs","realm","myron","loopy","plush","volts","bimbo","smash","windy","sours","choke","karin","boast","whirr","tiber","dimes","basel","cutes","pinto","troll","thumb","decor","craft","tared","split","josue","tramp","screw","label","lenny","apses","slept","sikhs","child","bouts","cites","swipe","lurks","seeds","fists","hoard","steed","reams","spoil","diego","peale","bevel","flags","mazes","quart","snipe","latch","lards","acted","falls","busby","holed","mummy","wrong","wipes","carlo","leers","wails","night","pasty","eater","flunk","vedas","curse","tyros","mirth","jacky","butte","wired","fixes","tares","vague","roved","stove","swoon","scour","coked","marge","cants","comic","corns","zilch","typos","lives","truer","comma","gaily","teals","witty","hyper","croat","sways","tills","hones","dowel","llano","clefs","fores","cinch","brock","vichy","bleed","nuder","hoyle","slams","macro","arabs","tauts","eager","croak","scoop","crime","lurch","weals","fates","clipt","teens","bulls","domed","ghana","culls","frame","hanky","jared","swain","truss","drank","lobby","lumps","pansy","whews","saris","trite","weeps","dozes","jeans","flood","chimu","foxes","gelds","sects","scoff","poses","mares","famed","peers","hells","laked","zests","wring","steal","snoot","yodel","scamp","ellis","bandy","marry","jives","vises","blurb","relay","patch","haley","cubit","heine","place","touch","grain","gerry","badly","hooke","fuchs","savor","apron","judge","loren","britt","smith","tammy","altar","duels","huber","baton","dived","apace","sedan","basts","clark","mired","perch","hulks","jolly","welts","quack","spore","alums","shave","singe","lanny","dread","profs","skeet","flout","darin","newed","steer","taine","salvo","mites","rules","crash","thorn","olive","saves","yawed","pique","salon","ovens","dusty","janie","elise","carve","winds","abash","cheep","strap","fared","discs","poxed","hoots","catch","combo","maize","repay","mario","snuff","delve","cored","bards","sudan","shuns","yukon","jowls","wayne","torus","gales","creek","prove","needy","wisps","terri","ranks","books","dicky","tapes","aping","padre","roads","nines","seats","flats","rains","moira","basic","loves","pulls","tough","gills","codes","chest","teeny","jolts","woody","flame","asked","dulls","hotly","glare","mucky","spite","flake","vines","lindy","butts","froth","beeps","sills","bunny","flied","shaun","mawed","velds","voled","doily","patel","snake","thigh","adler","calks","desks","janus","spunk","baled","match","strip","hosed","nippy","wrest","whams","calfs","sleet","wives","boars","chain","table","duked","riped","edens","galas","huffs","biddy","claps","aleut","yucks","bangs","quids","glenn","evert","drunk","lusts","senna","slate","manet","roted","sleep","loxes","fluky","fence","clamp","doted","broad","sager","spark","belch","mandy","deana","beyer","hoist","leafy","levee","libel","tonic","aloes","steam","skews","tides","stall","rifts","saxon","mavis","asama","might","dotes","tangs","wroth","kited","salad","liens","clink","glows","balky","taffy","sided","sworn","oasis","tenth","blurt","tower","often","walsh","sonny","andes","slump","scans","boded","chive","finer","ponce","prune","sloes","dined","chums","dingo","harte","ahead","event","freer","heart","fetch","sated","soapy","skins","royal","cuter","loire","minot","aisle","horny","slued","panel","eight","snoop","pries","clive","pored","wrist","piped","daren","cells","parks","slugs","cubed","highs","booze","weary","stain","hoped","finny","weeds","fetid","racer","tasks","right","saint","shahs","basis","refer","chart","seize","lulls","slant","belay","clots","jinny","tours","modes","gloat","dunks","flute","conch","marts","aglow","gayer","lazes","dicks","chime","bears","sharp","hatch","forms","terry","gouda","thins","janet","tonya","axons","sewed","danny","rowdy","dolts","hurry","opine","fifty","noisy","spiky","humid","verna","poles","jayne","pecos","hooky","haney","shams","snots","sally","ruder","tempe","plunk","shaft","scows","essie","dated","fleet","spate","bunin","hikes","sodas","filly","thyme","fiefs","perks","chary","kiths","lidia","lefty","wolff","withe","three","crawl","wotan","brown","japed","tolls","taken","threw","crave","clash","layer","tends","notes","fudge","musky","bawdy","aline","matts","shirr","balks","stash","wicks","crepe","foods","fares","rotes","party","petty","press","dolly","mangy","leeks","silly","leant","nooks","chapt","loose","caged","wages","grist","alert","sheri","moody","tamps","hefts","souls","rubes","rolex","skulk","veeps","nonce","state","level","whirl","bight","grits","reset","faked","spiny","mixes","hunks","major","missy","arius","damns","fitly","caped","mucus","trace","surat","lloyd","furry","colin","texts","livia","reply","twill","ships","peons","shear","norms","jumbo","bring","masks","zippy","brine","dorks","roded","sinks","river","wolfs","strew","myths","pulpy","prank","veins","flues","minus","phone","banns","spell","burro","brags","boyle","lambs","sides","knees","clews","aired","skirt","heavy","dimer","bombs","scums","hayes","chaps","snugs","dusky","loxed","ellen","while","swank","track","minim","wiled","hazed","roofs","cantu","sorry","roach","loser","brass","stint","jerks","dirks","emory","campy","poise","sexed","gamer","catty","comte","bilbo","fasts","ledge","drier","idles","doors","waged","rizal","pured","weirs","crisp","tasty","sored","palmy","parts","ethel","unify","crows","crest","udder","delis","punks","dowse","totes","emile","coded","shops","poppa","pours","gushy","tiffs","shads","birds","coils","areas","boons","hulls","alter","lobes","pleat","depth","fires","pones","serra","sweat","kline","malay","ruled","calve","tired","drabs","tubed","wryer","slung","union","sonya","aided","hewed","dicey","grids","nixed","whits","mills","buffs","yucky","drops","ready","yuppy","tweet","napes","cadre","teach","rasps","dowdy","hoary","canto","posed","dumbo","kooks","reese","snaky","binge","byron","phony","safer","friar","novel","scale","huron","adorn","carla","fauna","myers","hobby","purse","flesh","smock","along","boils","pails","times","panza","lodge","clubs","colby","great","thing","peaks","diana","vance","whets","bergs","sling","spade","soaks","beach","traps","aspen","romps","boxed","fakir","weave","nerds","swazi","dotty","curls","diver","jonas","waite","verbs","yeast","lapel","barth","soars","hooks","taxed","slews","gouge","slags","chang","chafe","saved","josie","syncs","fonds","anion","actor","seems","pyrex","isiah","glued","groin","goren","waxes","tonia","whine","scads","knelt","teaks","satan","tromp","spats","merry","wordy","stake","gland","canal","donna","lends","filed","sacks","shied","moors","paths","older","pooch","balsa","riced","facet","decaf","attic","elder","akron","chomp","chump","picky","money","sheer","bolls","crabs","dorms","water","veers","tease","dummy","dumbs","lethe","halls","rifer","demon","fucks","whips","plops","fuses","focal","taces","snout","edict","flush","burps","dawes","lorry","spews","sprat","click","deann","sited","aunts","quips","godly","pupil","nanny","funks","shoon","aimed","stacy","helms","mints","banks","pinch","local","twine","pacts","deers","halos","slink","preys","potty","ruffs","pusan","suits","finks","slash","prods","dense","edsel","heeds","palls","slats","snits","mower","rares","ailed","rouge","ellie","gated","lyons","duded","links","oaths","letha","kicks","firms","gravy","month","kongo","mused","ducal","toted","vocal","disks","spied","studs","macao","erick","coupe","starr","reaps","decoy","rayon","nicks","breed","cosby","haunt","typed","plain","trays","muled","saith","drano","cower","snows","buses","jewry","argus","doers","flays","swish","resin","boobs","sicks","spies","bails","wowed","mabel","check","vapid","bacon","wilda","ollie","loony","irked","fraud","doles","facts","lists","gazed","furls","sunks","stows","wilde","brick","bowed","guise","suing","gates","niter","heros","hyped","clomp","never","lolls","rangy","paddy","chant","casts","terns","tunas","poker","scary","maims","saran","devon","tripe","lingo","paler","coped","bride","voted","dodge","gross","curds","sames","those","tithe","steep","flaks","close","swops","stare","notch","prays","roles","crush","feuds","nudge","baned","brake","plans","weepy","dazed","jenna","weiss","tomes","stews","whist","gibed","death","clank","cover","peeks","quick","abler","daddy","calls","scald","lilia","flask","cheer","grabs","megan","canes","jules","blots","mossy","begun","freak","caved","hello","hades","theed","wards","darcy","malta","peter","whorl","break","downs","odder","hoofs","kiddo","macho","fords","liked","flees","swing","elect","hoods","pluck","brook","astir","bland","sward","modal","flown","ahmad","waled","craps","cools","roods","hided","plath","kings","grips","gives","gnats","tabby","gauls","think","bully","fogey","sawed","lints","pushy","banes","drake","trail","moral","daley","balds","chugs","geeky","darts","soddy","haves","opens","rends","buggy","moles","freud","gored","shock","angus","puree","raves","johns","armed","packs","minis","reich","slots","totem","clown","popes","brute","hedge","latin","stoke","blend","pease","rubik","greer","hindi","betsy","flows","funky","kelli","humps","chewy","welds","scowl","yells","cough","sasha","sheaf","jokes","coast","words","irate","hales","camry","spits","burma","rhine","bends","spill","stubs","power","voles","learn","knoll","style","twila","drove","dacca","sheen","papas","shale","jones","duped","tunny","mouse","floss","corks","skims","swaps","inned","boxer","synch","skies","strep","bucks","belau","lower","flaky","quill","aural","rufus","floes","pokes","sends","sates","dally","boyer","hurts","foyer","gowns","torch","luria","fangs","moats","heinz","bolts","filet","firth","begot","argue","youth","chimp","frogs","kraft","smite","loges","loons","spine","domes","pokey","timur","noddy","doggy","wades","lanes","hence","louts","turks","lurid","goths","moist","bated","giles","stood","winos","shins","potts","brant","vised","alice","rosie","dents","babes","softy","decay","meats","tanya","rusks","pasts","karat","nuked","gorge","kinks","skull","noyce","aimee","watch","cleat","stuck","china","testy","doses","safes","stage","bayes","twins","limps","denis","chars","flaps","paces","abase","grays","deans","maria","asset","smuts","serbs","whigs","vases","robyn","girls","pents","alike","nodal","molly","swigs","swill","slums","rajah","bleep","beget","thanh","finns","clock","wafts","wafer","spicy","sorer","reach","beats","baker","crown","drugs","daisy","mocks","scots","fests","newer","agate","drift","marta","chino","flirt","homed","bribe","scram","bulks","servo","vesta","divas","preps","naval","tally","shove","ragas","blown","droll","tryst","lucky","leech","lines","sires","pyxed","taper","trump","payee","midge","paris","bored","loads","shuts","lived","swath","snare","boned","scars","aeons","grime","writs","paige","rungs","blent","signs","davis","dials","daubs","rainy","fawns","wrier","golds","wrath","ducks","allow","hosea","spike","meals","haber","muses","timed","broom","burks","louis","gangs","pools","vales","altai","elope","plied","slain","chasm","entry","slide","bawls","title","sings","grief","viola","doyle","peach","davit","bench","devil","latex","miles","pasha","tokes","coves","wheel","tried","verdi","wanda","sivan","prior","fryer","plots","kicky","porch","shill","coats","borne","brink","pawed","erwin","tense","stirs","wends","waxen","carts","smear","rival","scare","phase","bragg","crane","hocks","conan","bests","dares","molls","roots","dunes","slips","waked","fours","bolds","slosh","yemen","poole","solid","ports","fades","legal","cedes","green","curie","seedy","riper","poled","glade","hosts","tools","razes","tarry","muddy","shims","sword","thine","lasts","bloat","soled","tardy","foots","skiff","volta","murks","croci","gooks","gamey","pyxes","poems","kayla","larva","slaps","abuse","pings","plows","geese","minks","derby","super","inked","manic","leaks","flops","lajos","fuzes","swabs","twigs","gummy","pyres","shrew","islet","doled","wooly","lefts","hunts","toast","faith","macaw","sonia","leafs","colas","conks","altos","wiped","scene","boors","patsy","meany","chung","wakes","clear","ropes","tahoe","zones","crate","tombs","nouns","garth","puked","chats","hanks","baked","binds","fully","soaps","newel","yarns","puers","carps","spelt","lully","towed","scabs","prime","blest","patty","silky","abner","temps","lakes","tests","alias","mines","chips","funds","caret","splat","perry","turds","junks","cramp","saned","peary","snarl","fired","stung","nancy","bulge","styli","seams","hived","feast","triad","jaded","elvin","canny","birth","routs","rimed","pusey","laces","taste","basie","malls","shout","prier","prone","finis","claus","loops","heron","frump","spare","menus","ariel","crams","bloom","foxed","moons","mince","mixed","piers","deres","tempt","dryer","atone","heats","dario","hawed","swims","sheet","tasha","dings","clare","aging","daffy","wried","foals","lunar","havel","irony","ronny","naves","selma","gurus","crust","percy","murat","mauro","cowed","clang","biker","harms","barry","thump","crude","ulnae","thong","pager","oases","mered","locke","merle","soave","petal","poser","store","winch","wedge","inlet","nerdy","utter","filth","spray","drape","pukes","ewers","kinds","dates","meier","tammi","spoor","curly","chill","loped","gooey","boles","genet","boost","beets","heath","feeds","growl","livid","midst","rinds","fresh","waxed","yearn","keeps","rimes","naked","flick","plies","deeps","dirty","hefty","messy","hairy","walks","leper","sykes","nerve","rover","jived","brisk","lenin","viper","chuck","sinus","luger","ricks","hying","rusty","kathy","herds","wider","getty","roman","sandy","pends","fezes","trios","bites","pants","bless","diced","earth","shack","hinge","melds","jonah","chose","liver","salts","ratty","ashed","wacky","yokes","wanly","bruce","vowel","black","grail","lungs","arise","gluts","gluey","navel","coyer","ramps","miter","aldan","booth","musty","rills","darns","tined","straw","kerri","hared","lucks","metes","penny","radon","palms","deeds","earls","shard","pried","tampa","blank","gybes","vicki","drool","groom","curer","cubes","riggs","lanky","tuber","caves","acing","golly","hodge","beard","ginny","jibed","fumes","astor","quito","cargo","randi","gawky","zings","blind","dhoti","sneak","fatah","fixer","lapps","cline","grimm","fakes","maine","erika","dealt","mitch","olden","joist","gents","likes","shelf","silts","goats","leads","marin","spire","louie","evans","amuse","belly","nails","snead","model","whats","shari","quote","tacks","nutty","lames","caste","hexes","cooks","miner","shawn","anise","drama","trike","prate","ayers","loans","botch","vests","cilia","ridge","thugs","outed","jails","moped","plead","tunes","nosed","wills","lager","lacks","cried","wince","berle","flaws","boise","tibet","bided","shred","cocky","brice","delta","congo","holly","hicks","wraps","cocks","aisha","heard","cured","sades","horsy","umped","trice","dorky","curve","ferry","haler","ninth","pasta","jason","honer","kevin","males","fowls","awake","pores","meter","skate","drink","pussy","soups","bases","noyes","torts","bogus","still","soupy","dance","worry","eldon","stern","menes","dolls","dumpy","gaunt","grove","coops","mules","berry","sower","roams","brawl","greed","stags","blurs","swift","treed","taney","shame","easel","moves","leger","ville","order","spock","nifty","brian","elias","idler","serve","ashen","bizet","gilts","spook","eaten","pumas","cotes","broke","toxin","groan","laths","joins","spots","hated","tokay","elite","rawer","fiats","cards","sassy","milks","roost","glean","lutes","chins","drown","marks","pined","grace","fifth","lodes","rusts","terms","maxes","savvy","choir","savoy","spoon","halve","chord","hulas","sarah","celia","deems","ninny","wines","boggy","birch","raved","wales","beams","vibes","riots","warty","nigel","askew","faxes","sedge","sheol","pucks","cynic","relax","boers","whims","bents","candy","luann","slogs","bonny","barns","iambs","fused","duffy","guilt","bruin","pawls","penis","poppy","owing","tribe","tuner","moray","timid","ceded","geeks","kites","curio","puffy","perot","caddy","peeve","cause","dills","gavel","manse","joker","lynch","crank","golda","waits","wises","hasty","paves","grown","reedy","crypt","tonne","jerky","axing","swept","posse","rings","staff","tansy","pared","glaze","grebe","gonna","shark","jumps","vials","unset","hires","tying","lured","motes","linen","locks","mamas","nasty","mamie","clout","nader","velma","abate","tight","dales","serer","rives","bales","loamy","warps","plato","hooch","togae","damps","ofter","plumb","fifes","filmy","wiper","chess","lousy","sails","brahe","ounce","flits","hindu","manly","beaux","mimed","liken","forts","jambs","peeps","lelia","brews","handy","lusty","brads","marne","pesos","earle","arson","scout","showy","chile","sumps","hiked","crook","herbs","silks","alamo","mores","dunce","blaze","stank","haste","howls","trots","creon","lisle","pause","hates","mulch","mined","moder","devin","types","cindy","beech","tuned","mowed","pitts","chaos","colds","bidet","tines","sighs","slimy","brain","belle","leery","morse","ruben","prows","frown","disco","regal","oaken","sheds","hives","corny","baser","fated","throe","revel","bores","waved","shits","elvia","ferns","maids","color","coifs","cohan","draft","hmong","alton","stine","cluck","nodes","emily","brave","blair","blued","dress","bunts","holst","clogs","rally","knack","demos","brady","blues","flash","goofy","blocs","diane","colic","smile","yules","foamy","splay","bilge","faker","foils","condo","knell","crack","gallo","purls","auras","cakes","doves","joust","aides","lades","muggy","tanks","middy","tarps","slack","capet","frays","donny","venal","yeats","misty","denim","glass","nudes","seeps","gibbs","blows","bobbi","shane","yards","pimps","clued","quiet","witch","boxes","prawn","kerry","torah","kinko","dingy","emote","honor","jelly","grins","trope","vined","bagel","arden","rapid","paged","loved","agape","mural","budge","ticks","suers","wendi","slice","salve","robin","bleat","batik","myles","teddy","flatt","puppy","gelid","largo","attar","polls","glide","serum","fundy","sucks","shalt","sewer","wreak","dames","fonts","toxic","hines","wormy","grass","louse","bowls","crass","benny","moire","margo","golfs","smart","roxie","wight","reign","dairy","clops","paled","oddly","sappy","flair","shown","bulgy","benet","larch","curry","gulfs","fends","lunch","dukes","doris","spoke","coins","manna","conga","jinns","eases","dunno","tisha","swore","rhino","calms","irvin","clans","gully","liege","mains","besot","serge","being","welch","wombs","draco","lynda","forty","mumps","bloch","ogden","knits","fussy","alder","danes","loyal","valet","wooer","quire","liefs","shana","toyed","forks","gages","slims","cloys","yates","rails","sheep","nacho","divan","honks","stone","snack","added","basal","hasps","focus","alone","laxes","arose","lamed","wrapt","frail","clams","plait","hover","tacos","mooch","fault","teeth","marva","mucks","tread","waves","purim","boron","horde","smack","bongo","monte","swirl","deals","mikes","scold","muter","sties","lawns","fluke","jilts","meuse","fives","sulky","molds","snore","timmy","ditty","gasps","kills","carey","jawed","byers","tommy","homer","hexed","dumas","given","mewls","smelt","weird","speck","merck","keats","draws","trent","agave","wells","chews","blabs","roves","grieg","evens","alive","mulls","cared","garbo","fined","happy","trued","rodes","thurs","cadet","alvin","busch","moths","guild","staci","lever","widen","props","hussy","lamer","riley","bauer","chirp","rants","poxes","shyer","pelts","funny","slits","tinge","ramos","shift","caper","credo","renal","veils","covey","elmer","mated","tykes","wooed","briar","gears","foley","shoes","decry","hypes","dells","wilds","runts","wilts","white","easts","comer","sammy","lochs","favor","lance","dawns","bushy","muted","elsie","creel","pocks","tenet","cagey","rides","socks","ogled","soils","sofas","janna","exile","barks","frank","takes","zooms","hakes","sagan","scull","heaps","augur","pouch","blare","bulbs","wryly","homey","tubas","limbo","hardy","hoagy","minds","bared","gabby","bilks","float","limns","clasp","laura","range","brush","tummy","kilts","cooed","worms","leary","feats","robes","suite","veals","bosch","moans","dozen","rarer","slyer","cabin","craze","sweet","talon","treat","yanks","react","creed","eliza","sluts","cruet","hafts","noise","seder","flies","weeks","venus","backs","eider","uriel","vouch","robed","hacks","perth","shiny","stilt","torte","throb","merer","twits","reeds","shawl","clara","slurs","mixer","newts","fried","woolf","swoop","kaaba","oozed","mayer","caned","laius","lunge","chits","kenny","lifts","mafia","sowed","piled","stein","whack","colts","warms","cleft","girds","seeks","poets","angel","trade","parsi","tiers","rojas","vexes","bryce","moots","grunt","drain","lumpy","stabs","poohs","leapt","polly","cuffs","giddy","towns","dacha","quoth","provo","dilly","carly","mewed","tzars","crock","toked","speak","mayas","pssts","ocher","motel","vogue","camps","tharp","taunt","drone","taint","badge","scott","scats","bakes","antes","gruel","snort","capes","plate","folly","adobe","yours","papaw","hench","moods","clunk","chevy","tomas","narcs","vonda","wiles","prigs","chock","laser","viced","stiff","rouse","helps","knead","gazer","blade","tumid","avail","anger","egged","guide","goads","rabin","toddy","gulps","flank","brats","pedal","junky","marco","tinny","tires","flier","satin","darth","paley","gumbo","rared","muffs","rower","prude","frees","quays","homes","munch","beefs","leash","aston","colon","finch","bogey","leaps","tempo","posts","lined","gapes","locus","maori","nixes","liven","songs","opted","babel","wader","barer","farts","lisps","koran","lathe","trill","smirk","mamma","viler","scurf","ravel","brigs","cooky","sachs","fulls","goals","turfs","norse","hauls","cores","fairy","pluto","kneed","cheek","pangs","risen","czars","milne","cribs","genes","wefts","vents","sages","seres","owens","wiley","flume","haded","auger","tatty","onion","cater","wolfe","magic","bodes","gulls","gazes","dandy","snags","rowed","quell","spurn","shore","veldt","turns","slavs","coach","stalk","snuck","piles","orate","joyed","daily","crone","wager","solos","earns","stark","lauds","kasey","villa","gnaws","scent","wears","fains","laced","tamer","pipes","plant","lorie","rivet","tamed","cozen","theme","lifer","sunny","shags","flack","gassy","eased","jeeps","shire","fargo","timer","brash","behan","basin","volga","krone","swiss","docks","booed","ebert","gusty","delay","oared","grady","buick","curbs","crete","lucas","strum","besom","gorse","troth","donne","chink","faced","ahmed","texas","longs","aloud","bethe","cacao","hilda","eagle","karyn","harks","adder","verse","drays","cello","taped","snide","taxis","kinky","penes","wicca","sonja","aways","dyers","bolas","elfin","slope","lamps","hutch","lobed","baaed","masts","ashes","ionic","joyce","payed","brays","malts","dregs","leaky","runny","fecal","woven","hurls","jorge","henna","dolby","booty","brett","dykes","rural","fight","feels","flogs","brunt","preen","elvis","dopey","gripe","garry","gamma","fling","space","mange","storm","arron","hairs","rogue","repel","elgar","ruddy","cross","medan","loses","howdy","foams","piker","halts","jewel","avery","stool","cruel","cases","ruses","cathy","harem","flour","meted","faces","hobos","charm","jamar","cameo","crape","hooey","reefs","denny","mitts","sores","smoky","nopes","sooty","twirl","toads","vader","julep","licks","arias","wrote","north","bunks","heady","batch","snaps","claws","fouls","faded","beans","wimps","idled","pulse","goons","noose","vowed","ronda","rajas","roast","allah","punic","slows","hours","metal","slier","meaty","hanna","curvy","mussy","truth","troys","block","reels","print","miffs","busts","bytes","cream","otter","grads","siren","kilos","dross","batty","debts","sully","bares","baggy","hippy","berth","gorky","argon","wacko","harry","smoke","fails","perms","score","steps","unity","couch","kelly","rumps","fines","mouth","broth","knows","becky","quits","lauri","trust","grows","logos","apter","burrs","zincs","buyer","bayer","moose","overt","croon","ousts","lands","lithe","poach","jamel","waive","wiser","surly","works","paine","medal","glads","gybed","paint","lorre","meant","smugs","bryon","jinni","sever","viols","flubs","melts","heads","peals","aiken","named","teary","yalta","styes","heist","bongs","slops","pouts","grape","belie","cloak","rocks","scone","lydia","goofs","rents","drive","crony","orlon","narks","plays","blips","pence","march","alger","baste","acorn","billy","croce","boone","aaron","slobs","idyls","irwin","elves","stoat","doing","globe","verve","icons","trial","olsen","pecks","there","blame","tilde","milky","sells","tangy","wrack","fills","lofty","truce","quark","delia","stowe","marty","overs","putty","coral","swine","stats","swags","weans","spout","bulky","farsi","brest","gleam","beaks","coons","hater","peony","huffy","exert","clips","riven","payer","doped","salas","meyer","dryad","thuds","tilts","quilt","jetty","brood","gulch","corps","tunic","hubby","slang","wreck","purrs","punch","drags","chide","sulks","tints","huger","roped","dopes","booby","rosin","outer","gusto","tents","elude","brows","lease","ceres","laxer","worth","necks","races","corey","trait","stuns","soles","teems","scrip","privy","sight","minor","alisa","stray","spank","cress","nukes","rises","gusts","aurae","karma","icing","prose","biked","grand","grasp","skein","shaky","clump","rummy","stock","twain","zoned","offed","ghats","mover","randy","vault","craws","thees","salem","downy","sangs","chore","cited","grave","spinx","erica","raspy","dying","skips","clerk","paste","moved","rooks","intel","moses","avers","staid","yawls","blast","lyres","monks","gaits","floor","saner","waver","assam","infer","wands","bunch","dryly","weedy","honey","baths","leach","shorn","shows","dream","value","dooms","spiro","raped","shook","stead","moran","ditto","loots","tapir","looms","clove","stops","pinks","soppy","ripen","wench","shone","bauds","doric","leans","nadia","cries","camus","boozy","maris","fools","morns","bides","greek","gauss","roget","lamar","hazes","beefy","dupes","refed","felts","larry","guile","ables","wants","warns","toils","bathe","edger","paced","rinks","shoos","erich","whore","tiger","jumpy","lamas","stack","among","punts","scalp","alloy","solon","quite","comas","whole","parse","tries","reeve","tiled","deena","roomy","rodin","aster","twice","musts","globs","parch","drawn","filch","bonds","tells","droop","janis","holds","scant","lopes","based","keven","whiny","aspic","gains","franz","jerri","steel","rowel","vends","yelps","begin","logic","tress","sunni","going","barge","blood","burns","basks","waifs","bones","skill","hewer","burly","clime","eking","withs","capek","berta","cheap","films","scoot","tweed","sizer","wheat","acton","flung","ponds","tracy","fiver","berra","roger","mutes","burke","miked","valve","whisk","runes","parry","toots","japes","roars","rough","irons","romeo","cages","reeks","cigar","saiph","dully","hangs","chops","rolls","prick","acuff","spent","sulla","train","swell","frets","names","anita","crazy","sixth","blunt","fewer","large","brand","slick","spitz","rears","ogres","toffy","yolks","flock","gnawn","eries","blink","skier","feted","tones","snail","ether","barbs","noses","hears","upset","awash","cloud","trunk","degas","dungs","rated","shall","yeahs","coven","sands","susan","fable","gunny","began","serfs","balls","dinky","madge","prong","spilt","lilly","brawn","comet","spins","raids","dries","sorts","makes","mason","mayra","royce","stout","mealy","pagan","nasal","folds","libby","coups","photo","mosey","amens","speed","lords","board","fetal","lagos","scope","raked","bonus","mutts","willy","sport","bingo","thant","araby","bette","rebel","gases","small","humus","grosz","beset","slays","steve","scrap","blahs","south","pride","heels","tubes","beady","lacey","genus","mauls","vying","spice","sexes","ester","drams","today","comae","under","jests","direr","yoked","tempi","early","boats","jesus","warts","guppy","gilda","quota","token","edwin","ringo","gaped","lemon","hurst","manor","arrow","mists","prize","silas","blobs","diets","ervin","stony","buddy","bates","rabid","ducat","ewing","jaunt","beads","doyen","blush","thoth","tiles","piper","short","peron","alley","decks","shunt","whirs","cushy","roils","betty","plugs","woken","jibes","foray","merak","ruing","becks","whale","shoot","dwelt","spawn","fairs","dozed","celts","blond","tikes","sabin","feint","vamps","cokes","willa","slues","bills","force","curst","yokel","surer","miler","fices","arced","douse","hilly","lucio","tongs","togas","minty","sagas","pates","welsh","bruno","decal","elate","linux","gyros","pryor","mousy","pains","shake","spica","pupal","probe","mount","shirk","purus","kilns","rests","graze","hague","spuds","sweep","momma","burch","maces","samar","brace","riser","booms","build","camel","flyer","synge","sauna","tonga","tings","promo","hides","clair","elisa","bower","reins","diann","lubed","nulls","picks","laban","milch","buber","stomp","bosom","lying","haled","avert","wries","macon","skids","fumed","ogles","clods","antic","nosey","crimp","purge","mommy","cased","taxes","covet","clack","butch","panty","lents","machs","exude","tooth","adore","shuck","asses","after","terra","dices","aryan","regor","romes","stile","cairo","maura","flail","eaves","estes","sousa","visas","baron","civet","kitty","freed","ralph","tango","gawks","cheat","study","fancy","fiber","musks","souse","brims","claim","bikes","venue","sired","thymi","rivas","skimp","pleas","woman","gimpy","cawed","minos","pints","knock","poked","bowen","risky","towel","oinks","linus","heals","pears","codas","inner","pitch","harpy","niger","madly","bumpy","stair","files","nobel","celli","spars","jades","balmy","kooky","plums","trues","gloss","trims","daunt","tubby","dared","wadis","smell","darby","stink","drill","dover","ruler","laden","dikes","layla","fells","maker","joked","horns","these","baize","spahn","whens","edged","mushy","plume","tucks","spurs","husky","dried","bigot","pupas","drily","aware","hagar","newly","knots","pratt","feces","sabik","watts","cooke","riles","seamy","fleas","dusts","barfs","roans","pawns","vivid","kirks","tania","feral","tubae","horne","aries","brits","combs","chunk","stork","waned","texan","elide","glens","emery","autos","trams","dosed","cheri","baits","jacks","whose","fazed","matte","swans","maxed","write","spays","orion","traci","horse","stars","strut","goods","verge","scuff","award","dives","wires","burnt","dimly","sleds","mayan","biped","quirk","sofia","slabs","waste","robby","mayor","fatty","items","bowel","mires","swarm","route","swash","sooth","paved","steak","upend","sough","throw","perts","stave","carry","burgs","hilts","plane","toady","nadir","stick","foist","gnarl","spain","enter","sises","story","scarf","ryder","glums","nappy","sixes","honed","marcy","offer","kneel","leeds","lites","voter","vince","bursa","heave","roses","trees","argos","leann","grimy","zelma","crick","tract","flips","folks","smote","brier","moore","goose","baden","riled","looks","sober","tusks","house","acmes","lubes","chows","neath","vivas","defer","allay","casey","kmart","pests","proms","eying","cider","leave","shush","shots","karla","scorn","gifts","sneer","mercy","copes","faxed","spurt","monet","awoke","rocky","share","gores","drawl","tears","mooed","nones","wined","wrens","modem","beria","hovel","retch","mates","hands","stymy","peace","carat","coots","hotel","karen","hayed","mamet","cuing","paper","rages","suave","reuse","auden","costs","loner","rapes","hazel","rites","brent","pumps","dutch","puffs","noons","grams","teats","cease","honda","pricy","forgo","fleck","hired","silos","merge","rafts","halon","larks","deere","jello","cunts","sifts","boner","morin","mimes","bungs","marie","harts","snobs","sonic","hippo","comes","crops","mango","wrung","garbs","natty","cents","fitch","moldy","adams","sorta","coeds","gilds","kiddy","nervy","slurp","ramon","fuzed","hiker","winks","vanes","goody","hawks","crowd","bract","marla","limbs","solve","gloom","sloop","eaton","memos","tames","heirs","berms","wanes","faint","numbs","holes","grubs","rakes","waist","miser","stays","antis","marsh","skyed","payne","champ","jimmy","clues","fatal","shoed","freon","lopez","snowy","loins","stale","thank","reads","isles","grill","align","saxes","rubin","rigel","walls","beers","wispy","topic","alden","anton","ducts","david","duets","fries","oiled","waken","allot","swats","woozy","tuxes","inter","dunne","known","axles","graph","bumps","jerry","hitch","crews","lucia","banal","grope","valid","meres","thick","lofts","chaff","taker","glues","snubs","trawl","keels","liker","stand","harps","casks","nelly","debby","panes","dumps","norma","racks","scams","forte","dwell","dudes","hypos","sissy","swamp","faust","slake","maven","lowed","lilts","bobby","gorey","swear","nests","marci","palsy","siege","oozes","rates","stunt","herod","wilma","other","girts","conic","goner","peppy","class","sized","games","snell","newsy","amend","solis","duane","troop","linda","tails","woofs","scuds","shies","patti","stunk","acres","tevet","allen","carpi","meets","trend","salty","galls","crept","toner","panda","cohen","chase","james","bravo","styed","coals","oates","swami","staph","frisk","cares","cords","stems","razed","since","mopes","rices","junes","raged","liter","manes","rearm","naive","tyree","medic","laded","pearl","inset","graft","chair","votes","saver","cains","knobs","gamay","hunch","crags","olson","teams","surge","wests","boney","limos","ploys","algae","gaols","caked","molts","glops","tarot","wheal","cysts","husks","vaunt","beaus","fauns","jeers","mitty","stuff","shape","sears","buffy","maced","fazes","vegas","stamp","borer","gaged","shade","finds","frock","plods","skied","stump","ripes","chick","cones","fixed","coled","rodeo","basil","dazes","sting","surfs","mindy","creak","swung","cadge","franc","seven","sices","weest","unite","codex","trick","fusty","plaid","hills","truck","spiel","sleek","anons","pupae","chiba","hoops","trash","noted","boris","dough","shirt","cowls","seine","spool","miens","yummy","grade","proxy","hopes","girth","deter","dowry","aorta","paean","corms","giant","shank","where","means","years","vegan","derek","tales"}
		//beginWord := "qa"
		//endWord := "sq"
		//wordList := []string{"si","go","se","cm","so","ph","mt","db","mb","sb","kr","ln","tm","le","av","sm","ar","ci","ca","br","ti","ba","to","ra","fa","yo","ow","sn","ya","cr","po","fe","ho","ma","re","or","rn","au","ur","rh","sr","tc","lt","lo","as","fr","nb","yb","if","pb","ge","th","pm","rb","sh","co","ga","li","ha","hz","no","bi","di","hi","qa","pi","os","uh","wm","an","me","mo","na","la","st","er","sc","ne","mn","mi","am","ex","pt","io","be","fm","ta","tb","ni","mr","pa","he","lr","sq","ye"}
		//beginWord := "hot"
		//endWord := "dog"
		//wordList := []string{"hot","cog","dog","tot","hog","hop","pot","dot"}
		//beginWord := "hot"
		//endWord := "dog"
		//wordList := []string{"hot","dog","dot"}
		res := ladderLength(beginWord,endWord,wordList)
		fmt.Println(res)
	}
	{
		input := [][]byte{{'O','X','X','O','X'},{'X','O','O','X','O'},{'X','O','X','O','X'},{'O','X','O','O','O'},{'X','X','O','X','O'}}
		solve(input)
	}
	{
		res := multiply("111","222")
		fmt.Println(res)
	}
	{
		input := []int{1,1,2,5,2,3,4}
		res := permuteUnique(input)
		fmt.Println(res)
	}
	{
		input := [][]int{{1,4},{1,5}}
		res := merge(input)
		fmt.Println(res)
	}
	{
		input := []int{1,3,1,2,0,5}
		res := maxSlidingWindow(input,3)
		fmt.Println(res)
	}
	{
		board := [][]byte{{'A','B','C','E'}, {'S','F','C','S'}, {'A','D','E','E'}}
		word := "SEE"
		res := exist(board,word)
		fmt.Println(res)
	}
	{
		input := []int{8,2,4,4,4,9,5,2,5,8,8,0,8,6,9,1,1,6,3,5,1,2,6,6,0,4,8,6,0,3,2,8,7,6,5,1,7,0,3,4,8,3,5,9,0,4,0,1,0,5,9,2,0,7,0,2,1,0,8,2,5,1,2,3,9,7,4,7,0,0,1,8,5,6,7,5,1,9,9,3,5,0,7,5}
		res := jump(input)
		fmt.Println(res)
	}
	{
		input := []int{3,0,2,1,2}
		res := canReach(input,2)
		fmt.Println(res)
	}
	{
		input := []int{25000,24999,24998,6,5,4,3,2,1,1,0,0}
		res := canJump(input)
		fmt.Println(res)
	}
	{
		input := []int{9,1,4,7,3,-1,0,5,8,-1,6}
		res := longestConsecutive(input)
		fmt.Println(res)
	}
	{
		input := []int{10,9,2,5,3,4}
		res := lengthOfLIS(input)
		fmt.Println(res)
	}
	{
		preorder := []int{3,9,20,15,7}
		inorder := []int{9,3,15,20,7}
		res := buildTree2(preorder,inorder)
		fmt.Println(res)
	}
	{
		res := numSquares(12)
		fmt.Println(res)
	}
	{
		res := numTrees(4)
		fmt.Println(res)
	}
	{
		input := [][]byte{{'1','1','0','0','0'},{'1','1','0','0','0'},{'0','0','1','0','0'},{'0','0','0','1','1'}}
		res := numIslands(input)
		fmt.Println(res)
	}
	{
		input := "23"
		res := letterCombinations(input)
		fmt.Println(res)
	}
	{
		input := "aa"
		res := partition(input)
		fmt.Println(res)
	}
	{
		input := [][]int{{1,3,1},{1,5,1},{4,2,1}}
		res := minPathSum(input)
		fmt.Println(res)
	}
	{
		input := []int{2,3,6,7}
		target := 7
		res := combinationSum(input,target)
		fmt.Println(res)
	}
	{
		//input := [][]int{{0,1,0}, {0,0,1}, {1,1,1}, {0,0,0}}
		input := [][]int{{0,0}}
		gameOfLife(input)
	}
	{
		input := [][]int{{5,1,9,11},
			{2,4,8,10},
			{13,3,6,7},
			{15,14,12,16}}
			rotate(input)
		fmt.Println(input)
	}
	{
		input := []int{1,3,4,2,2}
		res := findDuplicate(input)
		fmt.Println(res)
	}
	{
		input := "MCMXCIV"
		res := romanToInt(input)
		fmt.Println(res)
	}
	{
		input := "A man, a plan, a canal: Panama"
		res := isPalindrome(input)
		fmt.Println(res)
	}
	{
		ss := strings.TrimLeft("00010","0")
		fmt.Println(ss)
		res := largestNumber([]int{3,32,34,5,9})
		fmt.Println(res)
	}
	{
		res := longestSubstring("aaabb", 3)
		fmt.Println(res)
	}
	{
		res := subtractProductAndSum(234)
		fmt.Println(res)
	}
	{
		res := reverseParentheses("(u(love)i)")
		fmt.Println(res)
	}
	{
		res := numberOfArithmeticSlices([]int{1,2,3,4,8,9,10})
		fmt.Println(res)
	}
	{
		//input := [][]int{{-19,-1,-96,48,-94,36,16,55,-42,37,-59,6,-32,96,95,-58,13,-34,94,85},{17,44,36,-29,84,80,-34,50,-99,64,13,91,-27,25,-36,57,20,98,-100,-72},{-92,-75,86,90,-4,90,64,56,50,-63,10,-15,90,-66,-66,32,-69,-78,1,60},{21,51,-47,-43,-14,99,44,90,8,11,99,-62,57,59,69,50,-69,32,85,13},{-28,90,12,-18,23,61,-55,-97,6,89,36,26,26,-1,46,-50,79,-45,89,86},{-85,-10,49,-10,2,62,41,92,-67,85,86,27,89,-50,77,55,22,-82,-94,-98},{-50,53,-23,55,25,-22,76,-93,-7,66,-75,42,-35,-96,-5,4,-92,13,-31,-100},{-62,-78,8,-92,86,69,90,-37,81,97,53,-45,34,19,-19,-39,-88,-75,-74,-4},{29,53,-91,65,-92,11,49,26,90,-31,17,-84,12,63,-60,-48,40,-49,-48,88},{100,-69,80,11,-93,17,28,-94,52,64,-86,30,-9,-53,-8,-68,-33,31,-5,11},{9,64,-31,63,-84,-15,-30,-10,67,2,98,73,-77,-37,-96,47,-97,78,-62,-17},{-88,-38,-22,-90,54,42,-29,67,-85,-90,-29,81,52,35,13,61,-18,-94,61,-62},{-23,-29,-76,-30,-65,23,31,-98,-9,11,75,-1,-84,-90,73,58,72,-48,30,-81},{66,-33,91,-6,-94,82,25,-43,-93,-25,-69,10,-71,-65,85,28,-52,76,25,90},{-3,78,36,-92,-52,-44,-66,-53,-55,76,-7,76,-73,13,-98,86,-99,-22,61,100},{-97,65,2,-93,56,-78,22,56,35,-24,-95,-13,83,-34,-51,-73,2,7,-86,-19},{32,94,-14,-13,-6,-55,-21,29,-21,16,67,100,77,-26,-96,22,-5,-53,-92,-36},{60,93,-79,76,-91,43,-95,-16,74,-21,85,43,21,-68,-32,-18,18,100,-43,1},{87,-31,26,53,26,51,-61,92,-65,17,-41,27,-42,-14,37,-46,46,-31,-74,23},{-67,-14,-20,-85,42,36,56,9,11,-66,-59,-55,5,64,-29,77,47,44,-33,-77}}
		input := [][]int{{1,2,3},{4,5,6},{7,8,9}}
		res := minFallingPathSum(input)
		fmt.Println(res)
	}
	{
		//input := [][]byte{{"1","0","1","0","0"},{"1","0","1","1","1"},{"1","1","1","1","1"},{"1","0","0","1","0"}}
	}
	{
		input := [][]int{{0,1,1,1}, {1,1,1,1}, {0,1,1,1}}
		res := countSquares(input)
		fmt.Println(res)
	}
	{
		//input := [][]int{{3,3},{5,-1},{-2,4}}
		//input := [][]int{{1,3},{-2,2},{2,-2}}
		input := [][]int{{8,6},{9,7},{2,8},{5,7},{10,1},{4,1},{2,7},{9,8},{-2,4},{3,3},{5,-1},{-2,4}}
		res := kClosest(input,4)
		fmt.Println(res)
	}
	{
		input := [][]int{{0,0,0,22,0,24},{34,23,18,0,23,2},{11,39,20,12,0,0},{39,8,0,2,0,1},{19,32,26,20,20,30},{0,38,26,0,29,31}}
		//input := [][]int{{1,0,7},{2,0,6},{3,4,5},{0,3,0},{9,0,20}}
		//input := [][]int{{0,6,0},{5,8,7},{0,9,0}}
		res := getMaximumGold(input)
		fmt.Println(res)
	}
	{
		//input := [][]int{{1},{2},{3},{}}
		//input := [][]int{{1,3,2},{2,3},{2,1,3,1},{}}
		input := [][]int{{1,3},{3,0,1},{2},{0}}
		res := canVisitAllRooms(input)
		fmt.Println(res)
	}
	{
		//gas := []int{1,2,3,4,5}
		//cost := []int{3,4,5,1,2}
		gas := []int{2,3,4}
		cost := []int{3,4,3}
		res := canCompleteCircuit(gas,cost)
		fmt.Println(res)
	}
	{
		//pre := []int{1,2,4,5,3,6,7}
		//post := []int{4,5,2,6,7,3,1}
		pre := []int{1,2,3}
		post := []int{2,3,1}
		res := constructFromPrePost(pre,post)
		fmt.Println(res)
	}
	{
		res := mctFromLeafValues([]int{6,2,4})
		fmt.Println(res)
	}
	{
		res := maxSumAfterPartitioning([]int{1,15,7,9,2,5,10},3)
		fmt.Println(res)
	}
	{
		res := customSortString("cba","abcd")
		fmt.Println(res)
	}
	{
		input := [][]int{{0,1},{1,1}}
		res := oddCells(2,3,input)
		fmt.Println(res)
	}
	{
		res := pathInZigZagTree(14)
		fmt.Println(res)
	}
	{
		input := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
		res := subdomainVisits(input)
		fmt.Println(res)
	}
	{
		queens := [][]int{{5,6},{7,7},{2,1},{0,7},{1,6},{5,1},{3,7},{0,3},{4,0},{1,2},{6,3},{5,0},{0,4},{2,2},{1,1},{6,4},{5,4},{0,0},{2,6},{4,5},{5,2},{1,4},{7,5},{2,3},{0,5},{4,2},{1,0},{2,7},{0,1},{4,6},{6,1},{0,6},{4,3},{1,7}}
		king := []int{3,4}
		res := queensAttacktheKing(queens,king)
		fmt.Println(res)
	}
	{
		var input [][]int = [][]int{{0,0,1,1},{1,0,1,0},{1,1,0,0}}
		matrixScore(input)
	}
	{
		input := "RLRRLLRLRL"
		res := balancedStringSplit(input)
		fmt.Println(res)
	}
	{
		input := "ababcbacadefegdehijhklij"
		res := partitionLabels(input)
		fmt.Println(res)
	}
	{
		s1 := "ab"
		s2 := "eidboaoo"
		res := checkInclusion(s1,s2)
		fmt.Println(res)
	}
	{
		houses := []int{1,2,3,4,9,7,6,98}
		heaters := []int{30,65}
		res := findRadius(houses,heaters)
		fmt.Println(res)
	}
	{
		arr1 := []int{2,3,1,3,2,4,6,7,9,2,7,19,19}
		arr2 := []int{2,1,4,3,9,6}
		res := relativeSortArray(arr1,arr2)
		fmt.Println(res)
	}
	{
		var input []int = []int{1,1,1,1,2,2,2,2,2,2}
		res := hasGroupsSizeX(input)
		fmt.Println(res)
	}
	{
		var input []byte = []byte{'c', 'f', 'j'}
		res := nextGreatestLetter(input,'c')
		fmt.Println(res)
	}
	{
		var input []int = []int{30,20,150,100,40}
		res := numPairsDivisibleBy60(input)
		fmt.Println(res)
	}
	{
		nums := []int{3,3,2,6,7,8}
		largest := make([]int, 4, 4)
		copy(largest,nums[0:4])
		fmt.Println(largest)
	}
	{
		res := countBinarySubstrings(  "00110011")
		fmt.Println(res)
	}
	{
		var input [][]int = [][]int{{1,3},{2,3},{3,1}}
		res := findJudge(3,input)
		fmt.Println(res)
	}
	{
		var input []string = []string{"Hello","Alaska", "Dad", "Peace"}
		ret := findWords(input)
		fmt.Println(ret)
	}
	{
		floor := []int{2,3,4,5,18,17,6}
		var res int = maxArea(floor)
		fmt.Println(res)
	}
	{
		generate(5)
	}
	{
		var arr []int = []int{4,3,2,7,8,2,3,1}
		var res = findDuplicates(arr)
		fmt.Println(res)
	}

	var arr []int = []int{1,2,2,3,3,3,3,4,4,5,6,6,7,7,8}
	len := remove_duplicated_sorted_array(arr)
	fmt.Println(len)
	fmt.Println("start")


	ll := tree_node{
		4,
		nil,
		nil,
	}
	lr := tree_node{
		8,
		nil,
		nil,
	}
	l :=tree_node{
		6,
		&ll,
		&lr,
	}

	rr := tree_node{
		16,
		nil,
		nil,
	}
	rl := tree_node{
		12,
		nil,
		nil,
	}
	r := tree_node{
		14,
		&rl,
		&rr,

	}
	root := tree_node{
		11,
		&l,
		&r,
	}
	inorder_tree(&root)
	fmt.Println(root)
}

//func remove_duplicated_sorted_array(arr []int)  []int{
//	result := make([]int,0)
//	len := len(arr)
//	cur := arr[0]
//	result = append(result, cur)
//	for i := 1;i<len;i++{
//		if arr[i] == cur{
//			i++
//		}else{
//			result = append(result, arr[i])
//			cur = arr[i]
//		}
//	}
//	return result
//}

func remove_duplicated_sorted_array(arr []int)(length int){
	len := len(arr)
	if len <= 1{
		return
	}
	//cur_val := arr[0]
	cur_pos := 0
	for i := 1 ;i<len;i++{
		if arr[cur_pos] != arr[i]{
			cur_pos++
			arr[cur_pos] = arr[i]
		}
	}
	return cur_pos + 1
}

func sum_exist(arr []int){

}

func findDuplicates(nums []int) []int {
	var res []int
	for i := 0; i < len(nums);{
		if(nums[nums[i] - 1] != nums[i]){
			nums[nums[i] - 1],nums[i] = nums[i],nums[nums[i] - 1]
		} else{
			i++
		}
	}
	for i := 0;i < len(nums);i++{
		if(nums[i]-1 != i){
			res = append(res, nums[i])
		}
	}
	return res
}

//Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
//Output:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	var m map[string][]string
	for i := 0;i < len(strs);i++{
		data := []rune(strs[i])
		sort.Sort(sortRunes(data))
		var s string = string(data[:])
		m[s] = append(m[s],strs[i])
	}
	var res [][]string
	for _,val := range m{
		res = append(res,val)
	}
	return res
}

//a[ i ][ j ] = a[i - 1] [j - 1] + a[i - 1][ j ]
func generate(numRows int) [][]int {
	var res [][]int
	for i := 0;i < numRows;i++{
		var tmp []int = make([]int,i+1)
		tmp[0] = 1
		res = append(res, tmp)
		for j := 1;j <= i;j++{
			if i == j{
				res[i][j] = 1
			} else{
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			}
		}
	}
	return res
}

func maxArea(height []int) int {
	var length int = len(height)
	var low int = 0
	var high int = length - 1
	var max_cap int = 0
	for low < high  {
		var cap int = int(math.Abs(float64(high - low)) * math.Min(float64(height[high]),float64(height[low])))
		if cap > max_cap{
			max_cap = cap
		}

		//var left_cap int = int(math.Abs(float64(high - low - 1)) * math.Min(float64(height[high]),float64(height[low+1])))
		//var right_cap int = int(math.Abs(float64(high - low - 1)) * math.Min(float64(height[high - 1]),float64(height[low])))
		if (height[high - 1] >= height[low+1]){
			high--
		} else{
			low++
		}
	}
	return max_cap
}

//Here is my idea: instead of calculating area by height*width, we can think it in a cumulative way. In other words, sum water amount of each bin(width=1).
//Search from left to right and maintain a max height of left and right separately, which is like a one-side wall of partial container.
// Fix the higher one and flow water from the lower part. For example, if current height of left is lower, we fill water in the left bin.
// Until left meets right, we filled the whole container
func trap(height []int) int {
	var cap int
	var left_max int = 0
	var right_max int = 0
	var low int = 0
	var high int = len(height) - 1
	for low <= high{
		if(height[low] < height[high]){
			if(height[low] > left_max){
				left_max = height[low]
			} else{
				cap += left_max - height[low]
			}
			low++
		}else{
			if(height[high] > right_max){
				right_max = height[high]
			}else{
				cap += right_max - height[high]
			}
			high--
		}
	}
	return cap
}

func distributeCandies(candies int, num_people int) []int {
	var res []int = make([]int,num_people,num_people)
	var cnt int = 1
	for candies > cnt{
		res[(cnt - 1) % num_people] += cnt
		candies -= cnt
		cnt++
	}
	res[(cnt - 1) % num_people] += candies
	return res
}

func findWords(words []string) []string {
	var char_pos map[string]int = make(map[string]int)
	char_pos["q"] = 1
	char_pos["w"] = 1
	char_pos["e"] = 1
	char_pos["r"] = 1
	char_pos["t"] = 1
	char_pos["y"] = 1
	char_pos["u"] = 1
	char_pos["i"] = 1
	char_pos["o"] = 1
	char_pos["p"] = 1

	char_pos["a"] = 2
	char_pos["s"] = 2
	char_pos["d"] = 2
	char_pos["f"] = 2
	char_pos["g"] = 2
	char_pos["h"] = 2
	char_pos["j"] = 2
	char_pos["k"] = 2
	char_pos["l"] = 2

	char_pos["z"] = 3
	char_pos["x"] = 3
	char_pos["c"] = 3
	char_pos["v"] = 3
	char_pos["b"] = 3
	char_pos["n"] = 3
	char_pos["m"] = 3

	var res []string

	for _,str := range words{
		var last_pos int = -1
		var equal bool = true
		s := strings.ToLower(str)
		for i := 0;i < len(s);i++{
			ele := string(s[i])
			if(last_pos == -1){
				last_pos = char_pos[ele]
			}else{
				if(last_pos != char_pos[ele]){
					equal = false
					break
				}
			}
		}
		if(equal){
			res = append(res, str)
		}
	}
	return res
}

//[[518,518],[71,971],[121,862],[967,607],[138,754],[513,337],[499,873],[337,387],[647,917],[76,417]]
//
func twoCitySchedCost(costs [][]int) int {
	var res int = 0
	var num int = len(costs)
	if num <= 1{
		return res
	}
	var a_cnt int = 0
	var b_cnt int = 0
	for i := 0;i < num;i++{
		for j := 1;j < num - i;j++{
			if math.Abs(float64(costs[j][0] - costs[j][1])) > math.Abs(float64(costs[j-1][0] - costs[j-1][1])){
				costs[j],costs[j-1] = costs[j-1],costs[j]
			}
		}
	}
	for _,ele := range costs{
		if ele[0] > ele[1] && b_cnt < num/2 {
			res += ele[1]
			b_cnt++
		} else{
			res += ele[0]
			a_cnt++
		}
	}
	return res
}

func lemonadeChange(bills []int) bool {
	var change map[int]int = make(map[int]int)
	for _,pay  := range bills{
		if pay == 5 {
			if _,ok := change[pay];ok{
				change[pay]++
			}else{
				change[pay] = 1
			}
		}else if pay == 10{
			if val,ok := change[5]; !ok || val <= 0{
				return false
			}
			change[5]--
			if _,ok := change[pay];ok{
				change[pay]++
			}else{
				change[pay] = 1
			}
		}else if pay == 20{
			if val,ok := change[5]; !ok || val <= 0{
				return false
			}
			if val,ok := change[10]; ok && val > 0{
				change[10]--
				change[5]--
			}else{
				if change[5] < 3{
					return false
				} else{
					change[5] -= 3
				}
			}
		}
	}
	return true
}

//Input: N = 4, trust = [[1,3],[1,4],[2,3],[2,4],[4,3]]
//Output: 3
//3
//[[[1,3],[2,3],[3,1]]
func findJudge(N int, trust [][]int) int {
	var len int = len(trust)
	var record []int = make([]int,N+1)
	var i int = 0
	for i < len{
		record[trust[i][0]] = -1
		if record[trust[i][1]] != -1{
			record[trust[i][1]]++
		}
		i++
	}
	for index,e := range record{
		if e == (N-1){
			return index
		}
	}
	return -1
}



type IntSlice []int
func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Swap(i, j int){ s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func findContentChildren(g []int, s []int) int {
	sort.Sort(IntSlice(g))
	sort.Sort(IntSlice(s))
	var len_g int = len(g)
	var len_s int = len(s)
	i := 0
	j := 0
	for j < len_s && i < len_g{
		if s[j] < g[i]{
			j++
		}else{
			i++
			j++
		}
	}
	return i
}

func is_binary_string(s string,len int) bool{
	start := 0
	end := len-1
	mid := len/2
	for pre_val := s[start];start < mid;start++{
		if s[start] != pre_val{
			return false
		}
	}
	for post_val := s[end]; end >= mid;end--{
		if s[end] != post_val{
			return false
		}
	}
	return true
}

//func countBinarySubstrings(s string) int {
//	var res int = 0
//	var len int = len(s)
//	for i := 2;i <= len;i = i + 2{
//		start := 0
//		end := i
//		for ;end <= len;{
//			if s[start] != s[end-1] {
//				if is_binary_string(s[start:end],end-start) {
//					res++
//				}
//			}
//			start++
//			end++
//		}
//	}
//	return res
//}

func is_point_in_rectangle(px int,py int,rec []int) bool{
	return px > rec[0] && px < rec[2] && py > rec[1] && py < rec[3]
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return rec1[0]<rec2[2] && rec2[0] < rec1[2] && rec1[3] > rec2[1] && rec1[1] < rec2[3];
}

type KthLargest struct {
	k int
	nums []int
	largest []int
}


func ConstructorK(k int, nums []int) KthLargest {
	var ele KthLargest = KthLargest{
		k:k,
		nums:nums,
	}
	sort.Sort(IntSlice(ele.nums))
	ele.largest = make([]int, k, k)
	copy(ele.largest,ele.nums[0:k])
	return ele
}

func (this *KthLargest) Add(val int) int {
	this.nums = append(this.nums, val)
	this.largest = append(this.largest, val)
	sort.Sort(IntSlice(this.largest))
	this.largest = this.largest[0:this.k]
	return this.nums[this.k - 1]
}


func bfs2(grid [][]int,i int,j int,depth int) int{
	var rows int = len(grid)
	var columns int = len(grid[0])
	if i < 0 || i >= rows || j < 0 || j >= columns || grid[i][j] != 1{
		return 0
	}
	grid[i][j] = 2
	depth++
	max_depth := depth
	res :=  bfs2(grid,i-1,j,depth)
	if res > max_depth{
		max_depth = res
	}
	res = bfs2(grid,i+1,j,depth)
	if res > max_depth{
		max_depth = res
	}
	res = bfs2(grid,i,j-1,depth)
	if res > max_depth{
		max_depth = res
	}
	res = bfs2(grid,i,j+1,depth)
	if res > max_depth{
		max_depth = res
	}
	return max_depth
}

//[[2],[1],[1],[1],[2],[1[1]]
func orangesRotting2(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	max_res := 0
	for i := 0;i < rows;i++{
		for j :=0;j < columns;j++{
			if grid[i][j] == 2{
				res := bfs2(grid,i-1,j,0)
				if res > max_res{
					max_res = res
				}
				res = bfs2(grid,i+1,j,0)
				if res > max_res{
					max_res = res
				}
				res = bfs2(grid,i,j-1,0)
				if res > max_res{
					max_res = res
				}
				res = bfs2(grid,i,j+1,0)
				if res > max_res{
					max_res = res
				}
			}
		}
	}
	for i := 0;i < rows;i++{
		for j :=0;j < columns;j++{
			if grid[i][j] == 1{
				return -1
			}
		}
	}
	return max_res
}
//[30,20,150,100,40]
func numPairsDivisibleBy60(time []int) int {
	res := 0
	dict := make(map[int]int)
	l := len(time)
	for i := 0;i < l;i++{
		abs_val := time[i]%60
		if cnt,ok := dict[abs_val];ok && cnt > 0{
			res += dict[abs_val]
		}
		needed := (600 - time[i])%60
		if cnt,ok := dict[needed];ok && cnt > 0{
			dict[needed]++
		}else{
			dict[needed] = 1
		}
	}
	return res
}

//744
func nextGreatestLetter(letters []byte, target byte) byte {
	var low int = 0
	var high int = len(letters) - 1
	if target < letters[low] || target > letters[high]{
		return letters[low]
	}
	var res byte = 0
	for low < high{
		var mid int = (low + high)/2
		if letters[mid] == target{
			if mid + 1 <= high{
				res = letters[mid+1]
			}
			break
		}else if letters[mid] < target{
			low = mid
			if letters[low] > target{
				res = letters[low]
				break
			}
		}else{
			high = mid
			if letters[high] < target{
				res = letters[mid]
				break
			}
		}
	}
	return res
}

func get_value(n int) int {
	if n == 0{
		return 1
	}else if n == 1 {
		return n
	}else{
		return n * get_value(n - 1)
	}
}

func gen_last_value(n, m int) int{
	first := get_value(n)
	second := get_value(m)
	third := get_value((n - m))
	return first / (second * third)
}

func numEquivDominoPairs(dominoes [][]int) int {
	var res int = 0
	var dict map[[2]int]int = make(map[[2]int]int)
	for _,ele := range dominoes{
		var max,min int
		if ele[0] > ele[1]{
			max = ele[0]
			min = ele[1]
		}else{
			max = ele[1]
			min = ele[0]
		}

		var key [2]int = [2]int{min,max}
		if _,ok := dict[key];ok{
			dict[key] += 1
		}else{
			dict[key] = 1
		}
	}
	for _,num := range dict{
		if num >= 2{
			res += gen_last_value(num,2)
		}
	}
	return res
}

//763 eager way
func partitionLabels(S string) []int {
	var res []int
	var dict [26]int
	for pos,e := range S{
		dict[e - 'a'] = pos
	}
	l := len(S)
	i := 0
	start := 0
	last := 0
	for i < l {
		if dict[S[i] - 'a'] > last{
			last = dict[S[i] - 'a']
		}
		if last == i{
			if last == i{
				res = append(res, last - start + 1)
				start = i + 1
			}
		}
		i++
	}
	return res
}

//1051
func heightChecker(heights []int) int {
	var cnt int = 0

	return cnt
}

func check_balance(s string) bool{
	var L_cnt int = 0
	var R_cnt int = 0
	for i := 0 ;i < len(s);i++{
		if s[i] == 'L'{
			L_cnt++
		}else{
			R_cnt++
		}
	}
	return L_cnt == R_cnt
}

//1221
//"RLRRLLRLRL"
func balancedStringSplit(s string) int {
	var cnt int = 0
	begin := 0
	cur := 2
	for ;begin < len(s) && cur <= len(s);{
		sub := s[begin:cur]
		if check_balance(sub) {
			begin = cur
			cur = begin + 1
			cnt++
		}else{
				cur++
		}
	}
	return cnt
}



//861
func matrixScore(A [][]int) int {
	//reverse line
	for i := 0;i < len(A); i++{
		if A[i][0] != 1{
			for j := 0;j < len(A[i]) ; j++{
				if A[i][j] == 0{
					A[i][j] = 1
				}else{
					A[i][j] = 0
				}
			}
		}
	}
	//reverse column
	for i := 0; i < len(A[0]); i++{
		var zero_cnt int = 0
		for j := 0 ; j < len(A);j++{
			if A[j][i] == 0{
				zero_cnt++
			}
		}
		if zero_cnt > len(A)/2{
			for j := 0 ; j < len(A);j++{
				if A[j][i] == 0{
					A[j][i] = 1
				}else{
					A[j][i] = 0
				}
			}
		}
	}
	//calculate
	var res int = 0
	for i := 0;i < len(A); i++{
		var val int = 0
		for j := len(A[i]) - 1;j >= 0 ; j-- {
			val += A[i][j] << uint32(len(A[i]) - 1 - j)
		}
		res += val
	}
	return res
}

//1222
func check_point_inline(start []int,end []int,third []int) bool{

	if (third[0] - start[0])*(end[1] - start[1]) == (end[0] - start[0]) * (third[1] - start[1]) &&
	   math.Min(float64(start[0]),float64(end[0])) < float64(third[0]) && float64(third[0]) < math.Max(float64(start[0]),float64(end[0])) &&
		math.Min(float64(start[1]),float64(end[1])) < float64(third[1]) && float64(third[1]) < math.Max(float64(start[1]),float64(end[1])){
		return true
	}else{
		return false;
	}
}

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	var res [][] int
	var samelines [][]int
	var samecolumns [][]int
	var sametilts [][]int
	for i := 0; i < len(queens);i++{
		if queens[i][0] == king[0]{
			samelines = append(samelines, queens[i])
		}else if queens[i][1] == king[1]{
			samecolumns = append(samecolumns,queens[i])
		}else if math.Abs(float64(queens[i][0] - king[0])) == math.Abs(float64(queens[i][1] - king[1])){
			sametilts = append(sametilts,queens[i])
		}
	}
	min_left_distance := -1
	min_right_distance := 8
	var left_pos []int
	var right_pos []int
	for i := 0;i < len(samelines);i++{
		if samelines[i][1] < king[1]{
			if samelines[i][1] > min_left_distance{
				min_left_distance = samelines[i][1]
				left_pos = samelines[i]
			}
		}else{
			if samelines[i][1] < min_right_distance{
				min_right_distance = samelines[i][1]
				right_pos = samelines[i]
			}
		}
	}
	if len(left_pos) > 0{
		res = append(res, left_pos)
	}
	if len(right_pos) > 0{
		res = append(res,right_pos)
	}
	min_upper_pos := -1
	min_bottom_pos := 8
	var upper_pos []int
	var bottom_pos []int
	for i := 0;i < len(samecolumns) ;i++ {
		if samecolumns[i][0] < king[0]{
			if samecolumns[i][0] > min_upper_pos{
				min_upper_pos = samecolumns[i][0]
				upper_pos = samecolumns[i]
			}
		}else{
			if samecolumns[i][0] < min_bottom_pos{
				min_bottom_pos = samecolumns[i][0]
				bottom_pos = samecolumns[i]
			}
		}
	}
	if len(upper_pos) > 0{
		res = append(res, upper_pos)
	}
	if len(bottom_pos) > 0{
		res = append(res,bottom_pos)
	}
	var min_tilt_collection [][]int //save nearest queens
	var first bool = true
	for i := 0;i < len(sametilts);i++{
		if first{
			first = false
			min_tilt_collection = append(min_tilt_collection, sametilts[i])
			continue
		}
		var add_to_result bool = true
		for j := 0;j < len(min_tilt_collection);{
			if check_point_inline(king,sametilts[i],min_tilt_collection[j]) {
				add_to_result = false
				break
			}

			if check_point_inline(king,min_tilt_collection[j],sametilts[i]) {
				fmt.Println("tilt ", sametilts[i], " closer than ", min_tilt_collection[j])
				min_tilt_collection = append(min_tilt_collection[:j], min_tilt_collection[j+1:]...)
			}else{
				j++
			}
		}
		if add_to_result{
			fmt.Println( sametilts[i], "  is valid")
			min_tilt_collection = append(min_tilt_collection,sametilts[i])
		}
	}
	res = append(res,min_tilt_collection...)
	return res
}

//811
//Input:
//["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"]
//Output:
//["901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1 intel.mail.com","951 com"]
func subdomainVisits(cpdomains []string) []string {
	var res []string
	var records map[string]int = make(map[string]int);
	for _,ss := range cpdomains{
		item := strings.Split(ss," ")
		cnt,ok := strconv.Atoi(item[0])
		if ok != nil{
			return res
		}
		domain := item[1]
		for len(domain) > 0{
			if _,ok := records[domain];ok{
				records[domain] += cnt
			}else{
				records[domain] = cnt
			}
			pos := strings.Index(domain,".")
			if pos == -1{
				break
			}
			domain = domain[pos+1:]
		}
	}
	for domain,cnt := range records{
		s := strconv.Itoa(cnt) + " " + domain
		res = append(res,s)
	}
	return res
}

//419
//var total int = 0
//func dfs_419(board [][]byte,x int,y int,last_is_ship bool) bool{
//	if x < 0 || x > len(board)-1 || y < 0 || y > len(board[0]) - 1{
//		return false
//	}
//	if board[x][y] == 'X'{
//
//	}
//
//	dfs_419(board,x + 1,y,board[x][y] == 'X')
//	dfs_419(board,x ,y + 1,board[x][y] == 'X')
//}
func point_is_ship(board [][]byte,x int,y int) bool{
	if x < 0 || x > len(board)-1 || y < 0 || y > len(board[0]) - 1{
		return false
	}
	return board[x][y] == 'X'
}

func countBattleships(board [][]byte) int {
	var res int
	for i := 0; i < len(board);i++{
		for j := 0; j < len(board[0]);j++{
			if board[i][j] == 'X' && !point_is_ship(board,i+1,j) && !point_is_ship(board,i,j+1){
				res++
			}
		}
	}
	return res
}

//1104

func pathInZigZagTree(label int) []int {
	var res []int
	res = append(res, label)
	var layer_num int = math.Ilogb(float64(label))

	for layer_num > 1{
		var parent int = label/2
		layer_min := int(math.Pow(2,float64(layer_num - 1)))
		layer_max := int(math.Pow(2,float64(layer_num)) - 1)
		parent = layer_max - (parent - layer_min)
		label = parent
		layer_num--
		res = append(res, parent)
	}
	res = append(res, 1)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

//1255
// Input: words = ["dog","cat","dad","good"],
// letters = ["a","a","c","d","d","d","g","o","o"],
// score = [1,0,9,5,0,0,3,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0]
//Output: 23
//Explanation:
//Score  a=1, c=9, d=5, g=3, o=2
//Given letters, we can form the words "dad" (5+1+5) and "good" (3+2+2+5) with a score of 23.
//Words "dad" and "dog" only get a score of 21.
func maxScoreWords(words []string, letters []byte, score []int) int {
	var res int

	return res
}

//1252
//n = 2, m = 3, indices = [[0,1],[1,1]]
func oddCells(n int, m int, indices [][]int) int {
	var res int = 0
	var matrix [][]int = make([][]int,n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, m)
	}
	for i:= 0; i < len(indices);i++{
		row := indices[i][0]
		for c:= 0; c < m;c++{
			matrix[row][c]++
		}
		column := indices[i][1]
		for r := 0;r < n;r++{
			matrix[r][column]++
		}
	}
	for i := 0; i < n;i++{
		for j := 0;j < m;j++{
			if matrix[i][j]%2 == 1{
				res++
			}
		}
	}
	fmt.Println(matrix)
	return res
}

//951
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if nil == root1 && nil == root2{
		return true
	}
	if (nil == root1 || nil == root2) || (root1.Val != root2.Val){
		return false
	}
	return (flipEquiv(root1.Left,root2.Left) && flipEquiv(root1.Right,root2.Right)) || (flipEquiv(root1.Left,root2.Right) && flipEquiv(root1.Right,root2.Left))
}

//1078
func findOcurrences(text string, first string, second string) []string {
	var res []string
	var words []string = strings.Split(text," ")
	if len(words) <= 1{
		return res
	}
	for slow := 0; slow < len(words) - 2;slow++{
		if words[slow] == first{
			if words[slow+1] == second{
				res = append(res, words[slow+2])
			}
		}
	}
	return res
}

//986
func intervalIntersection(A [][]int, B [][]int) [][]int {
	var res [][]int
	index_a,index_b := 0,0
	for index_a < len(A) && index_b < len(B){
		if A[index_a][1] < B[index_b][0]{
			index_a++
		}else if B[index_b][1] < A[index_a][0]{
			index_b++
		}else{
			max_start := 0
			if A[index_a][0] > B[index_b][0]{
				max_start = A[index_a][0]
			}else{
				max_start = B[index_b][0]
			}
			min_end := 0
			if A[index_a][1] > B[index_b][1]{
				min_end = B[index_b][1]
				index_b++
			}else{
				min_end = A[index_a][1]
				index_a++
			}
			res = append(res, []int{max_start,min_end})
		}
	}
	return res
}

// 791
// S = "cba"
// T = "abcd"
// Output: "cbad"
func customSortString(S string, T string) string {
	var records map[int32]int = make(map[int32]int)
	for i,val := range S{
		records[val] = i
	}

	for i := 1;i < len(T) - 1;i++{
		for j := 0; j < len(T) - i;j++{
			var first_score,second_score int = -1,-1
			if _,ok := records[int32(T[j])];ok{
				first_score = records[int32(T[j])]
			}
			if _,ok := records[int32(T[j+1])];ok{
				second_score = records[int32(T[j+1])]
			}
			if first_score > second_score{
				c := []byte(T)  // 将字符串 s 转换为 []byte 类型
				c[j] = T[j+1]
				c[j+1] = T[j]
				T = string(c)
			}
		}
	}
	return T
}

//1043
func max_int_slice(nums []int) int{
	var max int = 0
	for _,val := range nums{
		if val > max{
			max = val
		}
	}
	return max
}

func max_sum(A []int,begin int,K int,records map[int]int) int{
	var max int = 0
	if begin >= len(A) {
		return max
	}
	for i := 1;i <= K;i++{
		if (begin + i)  > len(A) {
			break
		}
		var sum int = 0
		if val,ok := records[begin+i];ok{
			sum = val
		}else{
			sum = max_sum(A,begin + i,K,records)
			records[begin + i] = sum
		}
		cur_sum := max_int_slice(A[begin:begin+i]) * (i) + sum
		if cur_sum > max{
			max = cur_sum
		}
	}
	return max
}

func maxSumAfterPartitioning(A []int, K int) int {
	var records map[int]int = make(map[int]int,len(A))
	return max_sum(A,0,K,records)
}

//134
// [1,2,3,4,5]
// [3,4,5,1,2]

//	{2,3,4}
//	{3,4,3}
func canCompleteCircuit(gas []int, cost []int) int {
	var start_point int = -1
	var l = len(gas)
	for begin := 0; begin <= len(gas);begin++ {
		trace := begin + 1
		cur_gas := 0
		for begin !=  trace % l {
			cur_gas += gas[(trace-1) % l]
			if cur_gas < cost[(trace-1) % l]{
				break
			}else{
				cur_gas -= cost[(trace-1) % l]
			}
			trace++
		}
		cur_gas += gas[(trace-1) % l]
		if begin == trace % l && cur_gas >= cost[(trace-1) % l]{
			return begin
		}
	}
	return start_point
}

//841
// [[1],[2],[3],[]]
//  [[1,3,2],[2,3],[2,1,3,1],[]]
//  [[1,3],[3,0,1],[2],[0]]
func canVisitAllRooms(rooms [][]int) bool {
	var room_num int = len(rooms)
	var visited map[int]bool = make(map[int]bool,room_num)
	for i := 0; i < room_num;i++{
		visited[i] = false
	}
	var target int = (room_num - 1) * room_num / 2
	var cur_num int = 0
	l := list.New()
	l.PushBack(0)
	for l.Len() != 0{
		ele := l.Back()
		val := ele.Value.(int)
		for i := 0; i < len(rooms[val]);i++{
			if !visited[rooms[val][i]]{
				cur_num += rooms[val][i]
				visited[rooms[val][i]] = true
				l.PushBack(rooms[val][i])
			}
		}
		l.Remove(ele)
		if cur_num >= target{
			return true
		}
	}
	return false
}

//1219
func max_int(a,b int)int{
	if a > b {
		return a
	}else{
		return b
	}
}

func dfs_gold(grid [][]int,x int,y int,record [][]int)int{
	if x >= len(grid) || y >= len(grid[0]) || x < 0 || y < 0{
		return 0
	}
	if grid[x][y] == 0 || record[x][y] != 0{
		return 0
	}
	record[x][y] = 1
	var dup_record1,dup_record2,dup_record3,dup_record4 [][]int = make([][]int,len(grid)),make([][]int,len(grid)),make([][]int,len(grid)),make([][]int,len(grid))
	for i := 0;i < len(grid);i++ {
		dup_record1[i] = make([]int,len(grid[i]))
		dup_record2[i] = make([]int,len(grid[i]))
		dup_record3[i] = make([]int,len(grid[i]))
		dup_record4[i] = make([]int,len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			dup_record1[i][j] = record[i][j]
			dup_record2[i][j] = record[i][j]
			dup_record3[i][j] = record[i][j]
			dup_record4[i][j] = record[i][j]

		}
	}
	var res int = 0
	res = max_int(grid[x][y] + dfs_gold(grid,x - 1,y,dup_record1),grid[x][y] + dfs_gold(grid,x,y-1,dup_record2))
	res = max_int(res,grid[x][y] + dfs_gold(grid,x + 1,y,dup_record3))
	res = max_int(res,grid[x][y] + dfs_gold(grid,x,y + 1,dup_record4))
	return res
}

func getMaximumGold(grid [][]int) int {
	var max int = 0
	for i := 0;i < len(grid);i++{
		for j := 0;j < len(grid[i]);j++{
			if grid[i][j] != 0{
				var record [][]int = make([][]int,len(grid))
				for row := 0; row < len(grid);row++{
					record[row] = make([]int,len(grid[i]))
				}
				max = max_int(max,dfs_gold(grid,i,j,record))
			}
		}
	}
	return max
}

//973
//func compare_distance(x1 int,y1 int,x2 int,y2 int) bool{
//	return (x1 * x1 + y1 * y1) - (x2 * x2 + y2 * y2) >= 0
//}
func compare_distance(x1 int,y1 int,x2 int,y2 int) bool{
	res :=  (x1 * x1 + y1 * y1) - (x2 * x2 + y2 * y2)
	return res > 0
}

//最小的k个元素用大顶堆
func init_mink(points [][]int,i int,N int){
	temp_x := points[i][0]
	temp_y := points[i][1]
	for i < N {
		left_child := 2*i + 1
		right_child := 2*i + 2
		if left_child >= N {
			break
		}
		if right_child >= N {
			if compare_distance(temp_x,temp_y,points[left_child][0],points[left_child][1]){
				break
			}else{
				points[i][0] = points[left_child][0]
				points[i][1] = points[left_child][1]
				i = left_child
				break
			}
		}else{
			var max_point_index int
			if compare_distance(points[right_child][0],points[right_child][1],points[left_child][0],points[left_child][1]) {
				max_point_index = right_child
			}else{
				max_point_index = left_child
			}
			if compare_distance(temp_x,temp_y,points[max_point_index][0],points[max_point_index][1]){
				break
			}else{
				points[i][0] = points[max_point_index][0]
				points[i][1] = points[max_point_index][1]
				i = max_point_index
			}
		}
	}
	points[i][0] = temp_x
	points[i][1] = temp_y
}

func kClosest(points [][]int, K int) [][]int {
	if len(points) == K{
		return points
	}
	//堆化前K个数
	for i := K/2 - 1;i >= 0;i-- {
		init_mink(points, i,K)
	}
	//从K+1个数开始，和堆顶元素进行比较。如果元素小于堆顶，则替换堆顶元素，并调整堆
	for i := K;i < len(points);i++{
		if compare_distance(points[i][0],points[i][1],points[0][0],points[0][1]){
			continue
		}
		temp_x := points[i][0]
		temp_y := points[i][1]
		points[i][0] = points[0][0]
		points[i][1] = points[0][1]
		points[0][0] = temp_x
		points[0][1] = temp_y
		init_mink(points,0,K)
	}
	return points[0:K]
}

func kClosest2(points [][]int, K int) [][]int{
	if len(points) == K{
		return points
	}
	for i := 0;i < K;i++{
		for j := len(points) - 1;j > i; j--{
			if compare_distance(points[j-1][0],points[j-1][1],points[j][0],points[j][1]){
				points[j],points[j-1] = points[j-1],points[j]
			}
		}
	}
	return points[0:K]
}

//1026
func pre_visit(node *TreeNode,min int,max int,max_diff int)int{
	if nil == node{
		return 0
	}
	diff := int(math.Max(math.Abs(float64(node.Val - min)),math.Abs(float64(node.Val - max))))
	max_diff = int(math.Max(float64(max_diff),float64(diff)))
	min = int(math.Min(float64(node.Val),float64(min)))
	max = int(math.Max(float64(node.Val),float64(max)))
	res := int(math.Max(float64(pre_visit(node.Left,min,max,max_diff)),float64(pre_visit(node.Right,min,max,max_diff))))
	res = int(math.Max(float64(res),float64(max_diff)))
	return res
}

func maxAncestorDiff(root *TreeNode) int {
	if nil == root{
		return 0
	}
	return pre_visit(root,root.Val,root.Val,0)
}

//1277
// Input: matrix =
// [
//  [0,1,1,1],
//  [1,1,1,1],
//  [0,1,1,1]
// ]
// Output: 15
func min_int_number(nums ...int)int{
	var min int = math.MaxInt32
	for _,n := range nums{
		if n < min{
			min = n
		}
	}
	return min
}

func max_int_number(nums ...int)int{
	var max int = math.MinInt32
	for _,n := range nums{
		if n > max{
			max = n
		}
	}
	return max
}

func countSquares(matrix [][]int) int {
	var cnt int = 0
	for i := 1;i < len(matrix);i++{
		for j := 1;j < len(matrix[0]);j++{
			if matrix[i][j] == 0{
				continue
			}
			matrix[i][j] = 1 + min_int_number(matrix[i-1][j],matrix[i][j-1],matrix[i-1][j-1])
		}
	}
	for i := 0;i < len(matrix);i++{
		for j := 0;j < len(matrix[0]);j++{
			cnt += matrix[i][j]
		}
	}
	return cnt
}

//221
func min_byte_number(nums ...byte)byte{
	var min byte = 127
	for _,n := range nums{
		if n < min{
			min = n
		}
	}
	return min
}

func maximalSquare(matrix [][]byte) int {
	for i := 0;i < len(matrix);i++{
		for j := 0;j < len(matrix[0]);j++{
			matrix[i][j] -= '0'
		}
	}

	for i := 1;i < len(matrix);i++{
		for j := 1;j < len(matrix[0]);j++{
			if matrix[i][j] == 0{
				continue
			}
			matrix[i][j] = 1 + min_byte_number(matrix[i-1][j],matrix[i][j-1],matrix[i-1][j-1])
		}
	}
	var res int = 0
	for i := 0;i < len(matrix);i++{
		for j := 0;j < len(matrix[0]);j++{
			if int(matrix[i][j]) > res{
				res = int(matrix[i][j])
			}
		}
	}
	return res * res
}

//14
func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0{
		return ""
	}else if l == 1{
		return strs[0]
	}
	var res string = strs[0]
	for i := 1;i < l;i++{
		for j:= 0;j < len(strs[i]) && j < len(res);j++{
			if strs[i][j] != res[j]{
				if j == 0{
					return ""
				}
				res = res[0:j]
				break
			}
		}
	}
	return res
}

//1143
//Input: text1 = "abcde", text2 = "ace"
//Output: 3
func longestCommonSubsequence(text1 string, text2 string) int {
	return 0
}

//1254
func closedIsland(grid [][]int) int {
	return 0
}

//931
func dp_minpath(A [][]int,i int,j int,memo [][]int)int{
	if i >= len(A) || i < 0 || j >= len(A[0]) || j < 0 {
		return 0
	}
	if memo[i][j] != 0{
		return memo[i][j]
	}
	var res int = math.MaxInt32
	res = int(math.Min(float64(A[i][j] + dp_minpath(A,i+1,j,memo)),float64(res)))
	if (j + 1) < len(A[0]) {
		res = int(math.Min(float64(A[i][j]+dp_minpath(A, i+1, j+1,memo)), float64(res)))
	}
	if (j - 1) >= 0 {
		res = int(math.Min(float64(A[i][j]+dp_minpath(A, i+1, j-1,memo)), float64(res)))
	}
	memo[i][j] = res
	return res
}

func minFallingPathSum(A [][]int) int {
	var res int
	var memo [][]int = make([][]int,len(A))
	for i := 0; i < len(A);i++{
		memo[i] = make([]int,len(A[0]))
	}
	for j := 0; j < len(A[0]);j++{
		res = int(math.Min(float64(dp_minpath(A,0,j,memo)),float64(math.MaxInt32)))
	}
	for _,v := range memo[0]{
		if v < res{
			res = v
		}
	}
	return res
}

//1260
// [[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]] K = 4
func shiftGrid(grid [][]int, k int) [][]int {
	var row_num int = len(grid)
	var col_num int = len(grid[0])
	var total = row_num * col_num
	var res []int = make([]int,total)
	for i := 0;i < row_num;i++{
		for j := 0;j < col_num;j++{
			res[(i * col_num + j + k) % total] = grid[i][j]
		}
	}
	var grid2 [][]int = make([][]int,len(grid))
	for i := 0;i < row_num;i++{
		grid2[i] = make([]int,len(grid[0]))
		for j := 0;j < col_num;j++{
			grid2[i][j] = res[i * col_num + j]
		}
	}
	return grid2
}

//1072
func maxEqualRowsAfterFlips(matrix [][]int) int {
	var res int = 0
	var record map[string]int = make(map[string]int)
	for row := 0; row < len(matrix);row++{
		var s string = ""
		var reverse_s string = ""
		for column := 0; column < len(matrix[0]);column++{
			if matrix[row][column] == 1{
				reverse_s += "0"
			}else{
				reverse_s += "1"
			}
			s += strconv.Itoa(matrix[row][column])
		}
		if _,ok := record[reverse_s];ok{
			record[reverse_s]++
		}else{
			record[reverse_s] = 1
		}
		if _,ok := record[s];ok{
			record[s]++
		}else{
			record[s] = 1
		}
	}
	for _,val := range record{
		if val > res{
			res = val
		}
	}
	return res
}

//515
func preorder_search_max(node *TreeNode,record *[]int,level int){
	if nil == node{
		return
	}
	if level == len(*record){
		*record = append(*record, node.Val)
	}else{
		if node.Val > (*record)[level]{
			(*record)[level] = node.Val
		}
	}
	preorder_search_max(node.Left,record,level + 1)
	preorder_search_max(node.Right,record,level + 1)
}

func largestValues(root *TreeNode) []int {
	var res []int
	if nil == root{
		return res
	}
	preorder_search_max(root,&res,0)
	return res
}

//1249
// "())()((("
func reverse(s string) string {
	s1 := []rune(s)
	for i := 0; i < len(s1)/2; i++ {
		tmp := s1[i]
		s1[i] = s1[len(s1)-1-i]
		s1[len(s1)-1-i] = tmp
	}
	return string(s1)
}

func minRemoveToMakeValid(s string) string {
	var res string
	var start_tag int = 0
	for i := 0; i < len(s);i++{
		if s[i] == '('{
			start_tag++
		}else if s[i] == ')'{
				if start_tag <= 0{
					continue
				}else{
					start_tag--
				}
		}
		res += string(s[i])
	}
	var res2 string
	i := len(res) - 1
	for ; i >= 0 && start_tag > 0;i--{
		if res[i] == '('{
			start_tag--
			continue
		}
		res2 += string(res[i])
	}
	res3 := res[0:i+1] + reverse(res2)
	return res3
}

//1074
//Input: matrix = [[1,-1],[-1,1]], target = 0
//Output: 5
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	var res int = 0
	for row := 0;row < len(matrix);row++{
		for col := 0;col < len(matrix[0]);col++{

		}
	}
	return res
}

//260
//Input:  [1,2,1,3,2,5]
//Output: [3,5]
func singleNumber(nums []int) []int {
	var val int = 0
	for i := 0;i < len(nums);i++{
		val = val ^ nums[i]
	}
	var tag int = 1
	for i := 1;i < 32;i++{
		if val | tag == val{
			break
		}else{
			tag = tag << 1
		}
	}
	fmt.Println(tag)
	var group [][]int = make([][]int,2)
	for i := 0;i < len(nums);i++{
		if nums[i] | tag == nums[i]{
			group[0] = append(group[0],nums[i])
		}else{
			group[1] = append(group[1],nums[i])
		}
	}
	var first_num,second_num int = 0,0
	for i := 0;i < len(group[0]);i++{
		first_num ^= group[0][i]
	}
	for i := 0;i < len(group[1]);i++{
		second_num ^= group[1][i]
	}
	return []int{first_num,second_num}
}

//647
func try_palindromic(s string,start int,end int) int{
	var cnt int = 0
	for start >= 0 && end < len(s){
		if s[start] == s[end]{
			cnt++
		}else{
			break
		}
		start--
		end++
	}
	return cnt
}

//647
func countSubstrings(s string) int {
	var res int = 0
	for start := 0;start < len(s);start++{
		res += try_palindromic(s,start,start)
		fmt.Println(res)
		res += try_palindromic(s,start,start+1)
		fmt.Println(res)
	}
	return res
}

//413
func is_same_diff(A []int,left int,right int) (diff int,ret bool){
	if left < 0 || right >= len(A){
		diff = 0
		ret = false
		return
	}
	diff = A[left+1] - A[left]
	ret = true
    for i := left + 1;i < right;i++{
    	if (A[i+1] - A[i]) != diff{
    		ret = false
    		break
		}
	}
	return
}

func numberOfArithmeticSlices(A []int) int {
	var l int = len(A)
	if l < 3{
		return 0
	}
	var record map[int]int = make(map[int]int)
	for start,end := 0, 2;start < l && end < l;{
		_,ok := is_same_diff(A,start,end)
		if ok{
			end++
			if _,ok := record[start];ok{
				record[start]++
			}else{
				record[start] = 1
			}
		}else{
			start = end - 1
			end = start + 2
		}
	}
	res := 0
	for _,val := range record{
		for val > 0{
			res += val
			val--
		}
	}
	return res
}

//1190
func reverseParentheses(s string) string {
	stack := NewStack()
	for i := 0;i < len(s);i++{
		if s[i] == ')'{
			var sub string = ""
			for !stack.Empty(){
				v := stack.Pop().(byte)
				if v != '('{
					sub += string(v)
				}else{
					for j := 0;j < len(sub);j++{
						stack.Push(sub[j])
					}
					break
				}
			}
		}else{
			stack.Push(s[i])
		}
	}
	var res string
	for !stack.Empty(){
		res += string(stack.Pop().(byte))
	}
	return reverse(res)
}

//865
func level_visit_865(node *TreeNode,level int)(max_depth int,res *TreeNode){
	if nil == node{
		return level,node
	}
	level++
	left_max_depth,leftnode := level_visit_865(node.Left,level)
	right_max_depth,rightnode := level_visit_865(node.Right,level)
	if left_max_depth > right_max_depth {
		if left_max_depth >level {
			return left_max_depth,leftnode
		}
		return level,node
	}else if left_max_depth < right_max_depth{
		if right_max_depth > level{
			return right_max_depth,rightnode
		}
		return level,node
	}else{
		return left_max_depth,node
	}
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	if nil == root{
		return root
	}
	_,res :=  level_visit_865(root,0)
	return res
}

//1281
func subtractProductAndSum(n int) int {
	var sum int = 0
	var times int = 1
	for n > 0{
		val := n % 10
		sum += val
		times *= val
		n = n/10
	}
	return times - sum
}

//1266
func abs_int(n int)int{
	if n < 0{
		return -n
	}
	return n
}

func minTimeToVisitAllPoints(points [][]int) int {
	var res int = 0
	for pos := 1; pos < len(points);pos++{
		res += max_int(abs_int(points[pos][0] - points[pos-1][0]),abs_int(points[pos][1] - points[pos-1][1]))
	}
	return res
}

//1094
//func carPooling(trips [][]int, capacity int) bool {
//
//}

//395
func longestSubstring(s string, k int) int{
	l := len(s)
	var less_than_k map[uint8]int = make(map[uint8]int)
	for i := 0;i < l;i++{
		if _,ok := less_than_k[s[i]];ok{
			less_than_k[s[i]]++
		}else{
			less_than_k[s[i]] = 1
		}
	}
	for c,n := range less_than_k{
		if n >= k{
			delete(less_than_k,c)
		}
	}
	if len(less_than_k) == 0{
		return l
	}
	if len(less_than_k) == l{
		return 0
	}
	var trace []int = make([]int,l)
	for i := 0;i < l;i++{
		if _,ok := less_than_k[s[i]];ok{
			trace[i] = -1
		}
	}
	var max_len = 0
	start := 0
	for start < l {
		for start < l{
			if trace[start] != - 1{
				break
			}
			start++
		}
		if start == l{
			break
		}
		end := start
		for end < l{
			if trace[end] == -1{
				break
			}
			end++
		}
		res := longestSubstring(s[start:end],k)
		if res > max_len{
			max_len = res
		}
		start++
	}
	return max_len
}

//func longestSubstring(s string, k int) int {
//	var record [26]int
//	for _,v := range s{
//		record[v - 'a']++
//	}
//	var less_than_k map[byte]int = make(map[byte]int)
//	var i byte = 0
//	for ;i < 26;i++{
//		if record[i] != 0 && record[i] < k{
//			less_than_k[i] = record[i]
//		}
//	}
//	begin,end := 0,len(s) - 1
//	for begin < len(s) && begin < end{
//		if len(less_than_k) == 0{
//			break
//		}
//		_,okbegin := less_than_k[s[begin] - 'a']
//		_,okend := less_than_k[s[end] - 'a']
//		if !okbegin && !okend{
//			begin++
//			end--
//			continue
//		}
//		if _,ok := less_than_k[s[begin] - 'a'];ok{
//			less_than_k[s[begin] - 'a']--
//			if less_than_k[s[begin] - 'a'] == 0{
//				delete(less_than_k,s[begin] - 'a')
//			}
//			begin++
//		}
//		if begin < end{
//			if _,ok := less_than_k[s[end] - 'a'];ok{
//				less_than_k[s[end] - 'a']--
//				if less_than_k[s[end] - 'a'] == 0{
//					delete(less_than_k,s[end] - 'a')
//				}
//				end--
//			}
//		}
//		if len(less_than_k) == 0{
//			break
//		}
//	}
//	if len(less_than_k) > 0{
//		return 0
//	}
//	sub := s[begin:end+1]
//	var rest map[int32]int = make(map[int32]int)
//	for _,v := range sub{
//		if _,ok := rest[v];ok{
//			rest[v]++
//		}else{
//			rest[v] = 1
//		}
//	}
//	for _,v := range rest{
//		if v < k{
//			return 0
//		}
//	}
//	return end - begin + 1
//}

//326
func isPowerOfThree(n int) bool {
	if n > 1{
		for n > 1{
			if n % 3 != 0{
				return false
			}
			n = n / 3
		}
	}
	return n == 1
}

//179
func compare_pre_digit(a int,b int)bool{
	var s1 string = strconv.Itoa(a)
	var s2 string = strconv.Itoa(b)
	return strings.Compare(s1 + s2,s2 + s1) <= 0
}

func largestNumber(nums []int) string {
	for i := 0;i < len(nums);i++{
		for j := 1;j < len(nums) - i;j++{
			if compare_pre_digit(nums[j-1],nums[j]){
				nums[j-1],nums[j] = nums[j],nums[j-1]
			}
		}
	}
	var res string
	for _,v := range nums{
		res += strconv.Itoa(v)
	}
	res = strings.TrimLeft(res,"0")
	if res == ""{
		return "0"
	}
	return res
}

//125
func isPalindrome(s string) bool {
	ls := strings.ToLower(s)
	begin,end := 0,len(ls) - 1
	for begin < end{
		for begin < end{
			if (ls[begin] >= 'a' && ls[begin] <= 'z') || (ls[begin] >= '0' && ls[begin] <= '9'){
				break
			}else{
				begin++
			}
		}
		for end > begin{
			if (ls[end] >= 'a' && ls[end] <= 'z') || (ls[end] >= '0' && ls[end] <= '9'){
				break
			}else{
				end--
			}
		}
		if begin < end{
			if ls[begin] != ls[end]{
				return false
			}
			begin++
			end--
		}
	}
	return true
}

//122
func dp_profit(prices []int,start int,has_stock bool)int{
	if start >= len(prices){
		return 0
	}
	if has_stock{
		return max_int(dp_profit(prices,start + 1,true),dp_profit(prices,start + 1,false) + prices[start])
	}else{
		return max_int(dp_profit(prices,start + 1,true) - prices[start],dp_profit(prices,start + 1,false))
	}
}

func maxProfit(prices []int) int {
	return dp_profit(prices,0,false)
}

func maxProfit2(prices []int)int{
	var sum int = 0
	for i := 1;i < len(prices);i++{
		if prices[i] > prices[i-1]{
			sum += prices[i] - prices[i-1]
		}
	}
	return sum
}

//39
//Symbol       Value
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
func romanToInt(s string) int {
	var table map[string]int = make(map[string]int)
	table["I"] = 1
	table["V"] = 5
	table["X"] = 10
	table["L"] = 50
	table["C"] = 100
	table["D"] = 500
	table["M"] = 1000
	var next map[string]string = make(map[string]string)
	next["I"] = "V"
	next["V"] = "X"
	next["X"] = "L"
	next["L"] = "C"
	next["C"] = "D"
	next["D"] = "M"
	next["M"] = ""

	var total int = 0
	for i := len(s) - 1;i >= 0;i--{
		v := table[string(s[i])]
		upper := math.MaxInt32
		if string(s[i]) != "M"{
			upper = table[next[string(s[i])]]
		}
		if total >= upper{
			total -= v
		}else{
			total += v
		}
	}
	return total
}

//287
func findDuplicate(nums []int) int {
	min := 0
	max := len(nums)-1
	for min < max{
		mid := min + (max - min)/2
		cnt := 0
		for _,v := range nums{
			if v <= mid{
				cnt++
			}
		}
		if cnt <= mid{
			min = mid + 1
		}else{
			max = mid
		}
	}
	return min
}

//Given input matrix =
//[
//  [ 5, 1, 9,11],
//  [ 2, 4, 8,10],
//  [13, 3, 6, 7],
//  [15,14,12,16]
//],
//
//rotate the input matrix in-place such that it becomes:
//[
//  [15,13, 2, 5],
//  [14, 3, 4, 1],
//  [12, 6, 8, 9],
//  [16, 7,10,11]
//]
func rotate(matrix [][]int)  {
	//swap by lines
	rows := len(matrix)
	for row := 0;row < rows/2;row++{
		matrix[row],matrix[rows - 1 - row] = matrix[rows - 1 - row],matrix[row]
	}
	//swap by diagonal line
	for i := 1;i < rows;i++{
		for j := 0;j < i;j++{
			matrix[i][j],matrix[j][i] = matrix[j][i],matrix[i][j]
		}
	}
}

//36
func isValidSudoku(board [][]byte) bool {
	var row_map [9][]byte
	var column_map [9][]byte
	var threeplusthree_map [9][]byte
	for i := 0;i < len(board);i++{
		for j := 0;j < len(board[0]);j++{
			if board[i][j] != '.'{
				for _,val := range row_map[i]{
					if val == board[i][j]{
						return false
					}
				}
				row_map[i] = append(row_map[i], board[i][j])
				for _,val := range column_map[j]{
					if val == board[i][j]{
						return false
					}
				}
				column_map[j] = append(column_map[j], board[i][j])
				index := i/3 * 3  + j / 3
				for _,val := range threeplusthree_map[index]{
					if val == board[i][j]{
						return false
					}
				}
				threeplusthree_map[index] = append(threeplusthree_map[index],board[i][j])
			}
		}
	}
	return true
}

//289
//Input:
//[
//  [0,1,0],
//  [0,0,1],
//  [1,1,1],
//  [0,0,0]
//]
//Output:
//[
//  [0,0,0],
//  [1,0,1],
//  [0,1,1],
//  [0,1,0]
//]
func get_status(board [][]int,i,j int)int{
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]){
		return 0
	}
	return board[i][j]
}

func gameOfLife(board [][]int)  {
	var neighbour_cnt [][]int = make([][]int,len(board))
	for i := 0;i < len(board);i++{
		neighbour_cnt[i] = make([]int,len(board[0]))
		for j := 0;j < len(board[0]);j++{
			neighbour_cnt[i][j] = get_status(board,i - 1,j - 1) + get_status(board,i - 1,j) + get_status(board,i - 1,j + 1) + get_status(board,i,j - 1) + get_status(board,i,j + 1)+ get_status(board,i + 1,j - 1) + get_status(board,i + 1,j) + get_status(board,i + 1,j + 1)
		}
	}
	for i := 0;i < len(board);i++{
		for j := 0;j < len(board[0]);j++{
			if board[i][j] == 1{
				if neighbour_cnt[i][j] < 2 || neighbour_cnt[i][j] > 3{
					board[i][j] = 0
				}
			}else{
				if neighbour_cnt[i][j] == 3{
					board[i][j] = 1
				}
			}
		}
	}
}

//64
//Input:
//[
//  [1,3,1],
//  [1,5,1],
//  [4,2,1]
//]
//Output: 7
//Explanation: Because the path 1→3→1→1→1 minimizes the sum.
func dp_minpathsum(grid [][]int,row int,col int,target_row int,target_col int,visited [][]int)(int,bool){
	if row < 0|| row > target_row || col < 0 || col > target_col {
		return 0,false
	}
	if visited[row][col] != 0{
		return visited[row][col],true
	}
	if row == target_row && col == target_col{
		visited[row][col] = grid[row][col]
		return grid[row][col],true
	}
	var sum int = math.MaxInt32
	var res bool = false
	if val,ok := dp_minpathsum(grid,row + 1,col,target_row,target_col,visited);ok{
		sum = min_int_number(grid[row][col] + val,sum)
		res = ok
	}
	if val,ok := dp_minpathsum(grid,row,col + 1,target_row,target_col,visited);ok{
		sum = min_int_number(grid[row][col] + val,sum)
		res = ok
	}
	if res{
		visited[row][col] = sum
	}
	return sum,res
}

func minPathSum(grid [][]int) int {
	var visited [][]int = make([][]int,len(grid))
	for i := 0;i < len(grid);i++{
		visited[i] = make([]int,len(grid[0]))
	}
	target_row := len(grid) - 1
	target_col := len(grid[0]) - 1
	val,_ :=  dp_minpathsum(grid,0,0,target_row,target_col,visited)
	return val
}

//1295
func findNumbers(nums []int) int {
	var res int = 0
	for i := 0;i < len(nums);i++{
		if len(strconv.Itoa(nums[i])) % 2 == 0{
			res++
		}
	}
	return res
}

//1275
func is_line(steps [][]int)bool{
	if len(steps) < 3{
		return false
	}
	var row_map map[int]int = make(map[int]int)
	var col_map map[int]int = make(map[int]int)
	var tilt_line_points int = 0
	var reverse_tilt_line_points int = 0
	for i := 0;i < len(steps);i++{
		if _,ok := row_map[steps[i][0]];ok{
			row_map[steps[i][0]] += steps[i][1] + 1
		}else{
			row_map[steps[i][0]] = steps[i][1] + 1
		}
		if _,ok := col_map[steps[i][1]];ok{
			col_map[steps[i][1]] += steps[i][0] + 1
		}else{
			col_map[steps[i][1]] = steps[i][0] + 1
		}
		if steps[i][0] == steps[i][1]{
			tilt_line_points++
		}
		if (steps[i][0] == 0 && steps[i][1] == 2) || (steps[i][0] == 1 && steps[i][1] == 1)||(steps[i][0] == 2 && steps[i][1] == 0)  {
			reverse_tilt_line_points++
		}
	}
	for _,val := range row_map{
		if val == 6{
			return true
		}
	}
	for _,val := range col_map{
		if val == 6{
			return true
		}
	}
	if tilt_line_points == 3 || reverse_tilt_line_points == 3{
		return true
	}
	return false
}

func tictactoe(moves [][]int) string {
	var a_steps [][]int
	var b_steps [][]int
	for i := 0;i < len(moves);i++{
		if i % 2 == 0{
			a_steps = append(a_steps,moves[i])
		}else{
			b_steps = append(b_steps,moves[i])
		}
	}
	a_win := is_line(a_steps)
	b_win := is_line(b_steps)
	if !a_win && !b_win{
		if len(moves) < 9{
			return "Pending"
		}else{
			return "Draw"
		}
	}else if a_win{
		return "A"
	}else {
		return "B"
	}
}


//131
//Input: "aab"
//Output:
//[
//  ["aa","b"],
//  ["a","a","b"]
//]
func recursive_string(s string,begin int,record *[][]string,previous []string) {
	if begin < 0 || begin >= len(s) {
		return
	}
	for i,j := begin,begin + 1;i < j && j < len(s);j++{
		if isPalindrome(s[i:j]) {
			if  isPalindrome(s[j:]) {
				var res []string
				res = append(res,previous...)
				res = append(res,s[i:j])
				res = append(res,s[j:])
				*record = append(*record,res)
				recursive_string(s,j,record,append(previous,s[i:j]))
			}else{
				recursive_string(s,j,record,append(previous,s[i:j]))
			}
		}
	}
}

func partition(s string) [][]string {
	var res [][]string
	recursive_string(s,0,&res,nil)
	if isPalindrome(s){
		res = append(res, []string{s})
	}
	return res
}

//17
func dfs_letter(collection []string,cur_pos int,cur_str string,target int,records *[]string){
	if cur_pos >= target{
		if len(cur_str) >= target{
			*records = append(*records, cur_str)
		}
		return
	}

	for i := 0;i < len(collection[cur_pos]);i++{
		cur_str += string(collection[cur_pos][i])
		dfs_letter(collection,cur_pos+1,cur_str,target,records)
		cur_str = cur_str[:len(cur_str)-1]
	}
}


func letterCombinations(digits string) []string {
	if len(digits) == 0{
		return []string{}
	}
	var records map[int]string = make(map[int]string)
	records[2] = "abc"
	records[3] = "def"
	records[4] = "ghi"
	records[5] = "jkl"
	records[6] = "mno"
	records[7] = "pqrs"
	records[8] = "tuv"
	records[9] = "wxyz"

	var collection []string
	for _,v := range digits{
		num,err := strconv.Atoi(string(v))
		if err == nil{
			collection = append(collection,records[num])
		}
	}
	var res []string
	dfs_letter(collection,0,"",len(collection),&res)
	return res
}

//200
//Input:
//11000
//11000
//00100
//00011
//
//Output: 3
func dfs_numIslands(grid [][]byte,row int,col int){
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] != '1'{
		return
	}
	grid[row][col] = '0'
	dfs_numIslands(grid,row - 1,col)
	dfs_numIslands(grid,row + 1,col)
	dfs_numIslands(grid,row,col - 1)
	dfs_numIslands(grid,row,col + 1)
}

func numIslands(grid [][]byte) int {
	var res int = 0
	for i := 0;i < len(grid);i++{
		for j := 0;j < len(grid[0]);j++{
			if grid[i][j] == '1'{
				dfs_numIslands(grid,i,j)
				res++
			}
		}
	}
	return res
}

//300
//Input: [10,9,2,5,3,7,101,18]
//Output: 4
//Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
func lengthOfLIS(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	var dp []int = make([]int,len(nums)+1)
	dp[0] = 1
	for i := 1;i < len(nums);i++{
		max := 1
		for j := 0;j < i;j++{
			if nums[i] > nums[j]{
				if (dp[j] + 1) > max{
					max = dp[j] + 1
				}
			}
		}
		dp[i] = max
	}
	var res int = 1
	for _,val := range dp{
		if val > res{
			res = val
		}
	}
	return res
}

//1299
//Input: arr = [17,18,5,4,6,1]
//Output: [18,6,6,6,1,-1]
func replaceElements(arr []int) []int {
	if len(arr) == 0{
		return arr
	}
	var biggest int = arr[len(arr) - 1]
	arr[len(arr) - 1] = -1
	for i := len(arr) - 2;i >= 0;i--{
		if arr[i] > biggest{
			tmp := arr[i]
			arr[i] = biggest
			biggest = tmp
		}else{
			arr[i] = biggest
		}
	}
	return arr
}

//128
//Input: [100, 4, 200, 1, 3, 2]    [9,1,4,7,3,-1,0,5,8,-1,6]
//Output: 4
//Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
func dfs_longestConsecutive(record map[int]bool,num int,increase bool)int{
	if _,ok := record[num];ok{
		if !record[num]{
			return 0
		}
		if increase{
			return 1 + dfs_longestConsecutive(record,num + 1,increase)
		}else{
			return 1 + dfs_longestConsecutive(record,num - 1,increase)
		}
	}else{
		return 0
	}
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0{
		return 0
	}

	var record map[int]bool = make(map[int]bool)
	for i := 0;i < len(nums);i++{
		if _,ok := record[nums[i]];ok{
			continue
		}
		record[nums[i]] = true
	}
	var max int = 1
	for num,_ := range record{
		if record[num]{
			res := 1 + dfs_longestConsecutive(record,num + 1,true) + + dfs_longestConsecutive(record,num - 1,false)
			record[num] = false
			if res > max{
				max = res
			}
		}
	}
	return max
}

//73
//[
//  [0,1,2,0],
//  [3,4,5,2],
//  [1,3,1,5]
//]
//Output:
//[
//  [0,0,0,0],
//  [0,4,5,0],
//  [0,3,1,0]
//]

func setZeroes(matrix [][]int)  {
	if len(matrix) == 0{
		return
	}
	if len(matrix[0]) == 0{
		return
	}
	var first_row_is_zero bool = false
	var first_col_is_zero bool = false
	for i := 0;i < len(matrix);i++{
		if matrix[i][0] == 0{
			first_col_is_zero = true
			break
		}
	}
	for i := 0;i < len(matrix[0]);i++{
		if matrix[0][i] == 0{
			first_row_is_zero = true
			break
		}
	}
	for i := 1;i < len(matrix);i++{
		for j := 1;j < len(matrix[0]);j++{
			if matrix[i][j] == 0{
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1;i < len(matrix);i++{
		for j := 1;j < len(matrix[0]);j++{
			if matrix[0][j] == 0 || matrix[i][0] == 0{
				matrix[i][j] = 0
			}
		}
	}
	if first_row_is_zero{
		for i := 0;i < len(matrix[0]);i++{
			matrix[0][i] = 0
		}
	}
	if first_col_is_zero{
		for i := 0;i < len(matrix);i++{
			matrix[i][0] = 0
		}
	}
}

//55
//Input: [2,3,1,1,4]
//Output: true
func dfs_canJump(nums []int,start int,dp []int)bool{
	if start >= len(nums) - 1 {
		return true
	}
	max_steps := nums[start]
	if max_steps == 0{
		return false
	}
	if dp[start] == 1{
		return false
	}
	for i := 1;i <= max_steps;i++{
		if dp[start+i] == 1{
			continue
		}
		res := dfs_canJump(nums,start + i,dp)
		if res{
			return true
		}
		dp[start + i] = 1
	}
	return false
}

func canJump(nums []int) bool {
	var dp []int = make([]int,len(nums))
	return dfs_canJump(nums,0,dp)
}

//1306
//Input: arr = [4,2,3,0,3,1,2], start = 5
//Output: true
//All possible ways to reach at index 3 with value 0 are:
//index 5 -> index 4 -> index 1 -> index 3
//index 5 -> index 6 -> index 4 -> index 1 -> index 3
func dfs_canReach(arr []int,start int,dp []int)bool{
	if start < 0 || start >= len(arr){
		return false
	}
	if dp[start] == 1{
		return false
	}
	if arr[start] == 0{
		return true
	}
	dp[start] = 1
	forward_res := dfs_canReach(arr,start + arr[start],dp)
	backward_res := dfs_canReach(arr,start - arr[start],dp)
	return forward_res || backward_res
}

func canReach(arr []int, start int) bool {
	var dp []int = make([]int,len(arr))
	return dfs_canReach(arr,start,dp)
}

//45
//Input: [2,3,1,1,4]
//Output: 2
//Explanation: The minimum number of jumps to reach the last index is 2.
//    Jump 1 step from index 0 to 1, then 3 steps to the last index.
func jump(nums []int) int {
	var dp []int = make([]int,len(nums))
	dp[len(nums) - 1] = 0
	for i := len(nums) - 2;i >=0;i--{
		min_step := len(nums)
		for l := 1;l <= nums[i] && (i + l) <= len(nums) - 1;l++{
			min_step = min_int_number(1 + dp[i + l],min_step)
			//min_step = min_int_number(1 + dp[len(nums) - 1 - i - j],min_step)
		}
		dp[i] = min_step
	}
	return dp[0]
}

//[2,3,1,1,4]
//bfs solution
func jump2(nums []int) int{
	l := len(nums)
	var cur_fastest_pos int = nums[0]
	var end int = 0
	var steps int = 0
	for i := 0;i < l - 1;i++{
		cur_fastest_pos = max_int(cur_fastest_pos,i + nums[i])
		if cur_fastest_pos >= l - 1{
			steps++
			break
		}
		if i == end{//beyond end edge
			end = cur_fastest_pos
			steps++
		}
	}
	return steps
}

//79
//board =
//[
//  ['A','B','C','E'],
//  ['S','F','C','S'],
//  ['A','D','E','E']
//]
//
//Given word = "ABCCED", return true.
//Given word = "SEE", return true.
//Given word = "ABCB", return false.
// 1: up 2: right 3:down 4:left
type DIRECTION int32
const (
	 UP_FORBIDDEN DIRECTION = 1
	 RIGHT_FORBIDDEN DIRECTION = 2
	 DOWN_FORBIDDEN DIRECTION = 3
	 LEFT_FORBIDDEN DIRECTION = 4
)
func dfs_exist(board [][]byte,word string,cur_pos int,row int,col int,forbidden_orientation DIRECTION,dp [][]bool)bool{
	if cur_pos >= len(word){
		return true
	}
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || dp[row][col]{
		return false
	}

	if word[cur_pos] == board[row][col]{
		cur_pos++
		dp[row][col] = true
	}

	//if dfs_exist(board,word,cur_pos,row - 1,col,dp) {
	//	return true
	//}
	//if dfs_exist(board,word,cur_pos,row + 1,col,dp) {
	//	return true
	//}
	//if dfs_exist(board,word,cur_pos,row,col - 1,dp){
	//	return true
	//}
	//if dfs_exist(board,word,cur_pos,row,col  + 1,dp){
	//	return true
	//}

	if word[cur_pos] == board[row][col]{
		cur_pos--
		dp[row][col] = false
	}

	return false
}

func exist(board [][]byte, word string) bool {
	var record [][]bool = make([][]bool,len(board))
	for i := 0; i < len(board);i++{
		record[i] = make([]bool,len(board[0]))
	}
	return dfs_exist(board,word,0,0,0,UP_FORBIDDEN,record)
}

//239
type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	// 由于是最大堆，所以使用大于号
	return h[i] > h[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 弹出最后一个元素
func (h *MaxHeap) Pop() interface{}{
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

//[]int{1,3,1,2,0,5}
func maxSlidingWindow(nums []int, k int) []int {
	var l int = len(nums)
	var res []int = make([]int,l - k + 1)
	h := make(MaxHeap,0)
	for i := 0;i < k;i++{
		h.Push(nums[i])
	}
	heap.Init(&h)
	res[0] = h[0]
	for i := 1;i <= (l - k);i++{
		for j := 0;j < k;j++{
			if nums[i-1] == h[j]{
				h = append(h[:j],h[j+1:]...)
				break
			}
		}
		heap.Push(&h,nums[i + k - 1])
		heap.Init(&h)
		res[i] = h[0]
	}
	return res
}

type Position [][]int

func (pos Position) Len()int{
	return len(pos)
}

func (pos Position) Less(i,j int)bool{
	if pos[i][0] == pos[j][0]{
		return pos[i][1] >= pos[j][1]
	}
	return pos[i][0] <= pos[j][0]
}

func (pos Position) Swap(i,j int){
	pos[i],pos[j] = pos[j],pos[i]
}

//56
//Input: [[1,3],[2,6],[8,10],[15,18]]
//Output: [[1,6],[8,10],[15,18]]
//Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
func merge(intervals [][]int) [][]int {
	var l int = len(intervals)
	if l <= 1{
		return intervals
	}
	sort.Sort(Position(intervals))
	var res [][]int
	res = append(res,intervals[0])
	for i := 1;i < len(intervals);i++{
		if res[len(res)-1][1] >= intervals[i][0]{
			if res[len(res)-1][1] <= intervals[i][1]{
				res[len(res)-1][1] = intervals[i][1]
			}
		} else{
			res = append(res,intervals[i])
		}
	}
	return res
}

//47
//[1,1,2,3,3,4,5,6]
func dfs_permuteUnique(cur_pos int,nums []int,res *[][]int,record map[string]bool){
	if cur_pos == len(nums){
		var s string
		for _,n := range nums{
			s += strconv.Itoa(n)
		}
		if _,ok := record[s];ok{
			return
		}else{
			record[s] = true
		}
		var c []int = make([]int,len(nums))
		copy(c,nums)
		*res = append(*res,c)
		return
	}
	for i := cur_pos;i < len(nums);i++{
		if i > cur_pos{
			if nums[i] == nums[i-1] || nums[i] == nums[cur_pos]{
				continue
			}
		}
		nums[cur_pos],nums[i] = nums[i],nums[cur_pos]
		dfs_permuteUnique(cur_pos + 1,nums,res,record)
		nums[cur_pos],nums[i] = nums[i],nums[cur_pos]
	}
}

func permuteUnique(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	var record map[string]bool = make(map[string]bool)
	dfs_permuteUnique(0,nums,&res,record)
	return res
}

//130
//Example:
//X X X X
//X O O X
//X X O X
//X O X X
//After running your function, the board should be:
//
//X X X X
//X X X X
//X X X X
//X O X X
func fill_x(board [][]byte,r int,c int,rows int,columns int){
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return
	}
	if board[r][c] == 'X'{
		return
	}
	board[r][c] = 'X'
	fill_x(board,r - 1,c,rows,columns)
	fill_x(board,r + 1,c,rows,columns)
	fill_x(board,r,c - 1,rows,columns)
	fill_x(board,r,c + 1,rows,columns)
}

func can_be_replace(board [][]byte,r int,c int,rows int,columns int,visited map[string]bool)bool {
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return false
	}
	if board[r][c] == 'X'{
		return true
	}
	if r == 0 || r == (rows - 1) || c == 0 || c == (columns - 1){
		return false
	}
	k := strconv.Itoa(r) + "," + strconv.Itoa(c)
	if _,ok := visited[k];ok{
		return true
	}
	visited[k] = true
	return can_be_replace(board,r - 1,c,rows,columns,visited) && can_be_replace(board,r + 1,c,rows,columns,visited) &&
		can_be_replace(board,r,c - 1,rows,columns,visited) && can_be_replace(board,r,c + 1,rows,columns,visited)
}

func solve(board [][]byte)  {
	rows := len(board)
	if rows == 0{
		return
	}
	columns := len(board[0])
	var record map[string]bool = make(map[string]bool)
	for i := 1;i < rows;i++{
		for j := 1;j < columns;j++{
			if board[i][j] != 'X'{
				k := strconv.Itoa(i) + "," + strconv.Itoa(j)
				if _,ok := record[k];ok{
					continue
				}
				var visited map[string]bool = make(map[string]bool)
				ret := can_be_replace(board,i,j,rows,columns,visited)
				if ret{
					for k,v := range visited{
						record[k] = v
					}
					fill_x(board,i,j,rows,columns)
				}
			}
		}
	}
}

//127
//Input:
//beginWord = "hit",
//endWord = "cog",
//wordList = ["hot","dot","dog","lot","log","cog"]
//
//Output: 5
func is_similar(first string,second string)bool{
	var difference int = 0
	for i,_ := range first{
		if second[i] != first[i]{
			difference++
		}
		if difference >= 2{
			return false
		}
	}
	return difference == 1
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var find bool = false
	for _,s := range wordList{
		if endWord == s{
			find = true
			break;
		}
	}
	if !find{
		return 0
	}
	var trace map[string]bool = make(map[string]bool) //store unvisited words
	for _,word := range wordList{
		trace[word] = true
	}
	var depth list.List
	depth.PushBack(beginWord)
	var steps int = 1
	for depth.Len() > 0{
		cur_step_size := depth.Len()
		for i := 0;i < cur_step_size;i++{
			ele := depth.Front()
			cur_word := ele.Value.(string)
			if cur_word == endWord{
				return steps
			}
			depth.Remove(ele)
			delete(trace,cur_word)
			for j := 0;j < len(cur_word);j ++{
				for k := 0;k < 26;k++{
					search_word := cur_word[0:j] + string(k + 'a') + cur_word[j+1:]
					if _,ok := trace[search_word];ok{
						depth.PushBack(search_word)
					}
				}
			}
		}
		steps++
	}
	return 0
}

//994
//Input: [[2,1,1],[1,1,0],[0,1,1]]
//Output: 4
func fill_near_orgin(grid [][]int,visited,r,c,rows,cols int){
	if r - 1 >= 0 && grid[r-1][c] == 1{
		grid[r-1][c] = 2
	}
}

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	columns := len(grid[0])
	var steps int = 0
	var unvisited map[string]bool = make(map[string]bool) //store unvisit good orange position
	var depth [][]int //store current step bad orange
	for row := 0;row < rows;row++{
		for col := 0;col < columns;col++{
			if grid[row][col] == 1{
				k := strconv.Itoa(row)+","+strconv.Itoa(col)
				unvisited[k] = true
			}
			if grid[row][col] == 2{
				depth = append(depth,[]int{row,col})
			}
		}
	}
	if len(unvisited) == 0{// If no good orange exist,return 0
		return 0
	}
	for len(depth) > 0{
		l := len(depth)
		var cur_depth [][]int = make([][]int,len(depth))
		copy(cur_depth,depth)
		depth = make([][]int,0)//clear last level bad orange
		for i := 0;i < l;i++{
			r := cur_depth[i][0]
			c := cur_depth[i][1]
			//check 4 direction's orange
			if r - 1 >= 0 && grid[r-1][c] == 1{
				grid[r-1][c] = 2
				k := strconv.Itoa(r-1)+","+strconv.Itoa(c)
				delete(unvisited,k)
				depth = append(depth, []int{r-1,c})
			}
			if r + 1 < rows && grid[r+1][c] == 1{
				grid[r+1][c] = 2
				k := strconv.Itoa(r+1)+","+strconv.Itoa(c)
				delete(unvisited,k)
				depth = append(depth, []int{r+1,c})
			}
			if c - 1 >= 0 && grid[r][c-1] == 1{
				grid[r][c-1] = 2
				k := strconv.Itoa(r)+","+strconv.Itoa(c-1)
				delete(unvisited,k)
				depth = append(depth, []int{r,c-1})
			}
			if c + 1 < columns && grid[r][c+1] == 1{
				grid[r][c+1] = 2
				k := strconv.Itoa(r)+","+strconv.Itoa(c+1)
				delete(unvisited,k)
				depth = append(depth, []int{r,c+1})
			}
		}
		steps++
		if len(unvisited) == 0{// when there is no good orange,return steps
			return steps
		}
	}
	if len(unvisited) > 0{
		return -1
	}
	return steps
}

//686
func repeatedStringMatch(A string, B string) int {
	for i,_ := range B{
		if !strings.Contains(A,B[i:i+1]){
			return -1
		}
	}
	la := len(A)
	lb := len(B)
	cur_lena := la
	var buffer bytes.Buffer
	i := 1
	var find bool = false
	for ; i <= lb; i++ {
		buffer.WriteString(A)
		cur_lena += la
		if cur_lena < lb{
			continue
		}
		dupA := buffer.String()
		if strings.Contains(dupA,B){
			find = true
			break
		}
	}
	if !find{
		return -1
	}
	return i
}

//322
//Input: coins = [1, 2, 5], amount = 11
//Output: 3
//Explanation: 11 = 5 + 5 + 1
//Top to bottom
//func dp_coinChange(coins []int,amount int,rest int,cur_index int,memo map[int]int)int{
//	if cur_index >= len(coins) || rest < 0 {
//		return math.MaxInt32
//	}
//	if rest == 0{
//		return 1
//	}
//	//if _,ok := memo[rest];ok{
//	//	return memo[rest]
//	//}
//	steps := math.MaxInt32
//	//choose current coin and continue
//	l1 := 1 + dp_coinChange(coins,amount,rest - coins[cur_index],cur_index,memo)
//	if l1 < steps{
//		steps = l1
//	}
//	//choose current coin and move to next
//	l2  := 1 + dp_coinChange(coins,amount,rest - coins[cur_index],cur_index + 1,memo)
//	if l2 < steps{
//		steps = l2
//	}
//	//skip current coin and move to next
//	l3 := dp_coinChange(coins,amount,rest,cur_index + 1,memo)
//	if l3 < steps{
//		steps = l3
//	}
//	if steps != math.MaxInt32 {
//		if val,ok := memo[rest];ok{
//			if steps < val{
//				memo[rest] = steps
//			}
//		}else{
//			memo[rest] = steps
//		}
//	}
//	return steps
//}
//
//func coinChange(coins []int, amount int) int {
//	sort.Ints(coins)
//	for i := 0;i < len(coins)/2;i++{
//		coins[i],coins[len(coins)-i-1] = coins[len(coins)-i-1],coins[i]
//	}
//	var memo map[int]int = make(map[int]int)
//	res :=  dp_coinChange(coins,amount,amount,0,memo)
//	return res
//}

//func coinChange(coins []int, amount int) int {
//	var memo []int = make([]int,amount + 1)//memo[i] means the min step from 0 to i
//	for i,_ := range memo{
//		memo[i] = math.MaxInt32
//	}
//	memo[0] = 0
//	// var steps int = 0
//	for i := 1;i <= amount;i++{
//		min_steps := math.MaxInt32
//		for _,c := range coins{
//			if (i - c) >= 0{
//				if memo[i-c] < min_steps{
//					min_steps = memo[i-c]
//				}
//			}
//		}
//		if min_steps != math.MaxInt32{
//			memo[i] = 1 + min_steps
//		}
//	}
//	if memo[amount] == math.MaxInt32{
//		return - 1
//	}
//	return memo[amount]
//}

func coinChange(coins []int, amount int) int {
	var dp []int = make([]int,amount + 1)
	for i := 0;i < len(dp);i++{
		dp[i] = -1
	}
	dp[0] = 0
	var min_coins int = math.MaxInt32
	for i := 1;i <= amount;i++{
		min_coins = math.MaxInt32
		for _,c := range coins{
			if c > i{
				continue
			}
			if dp[i - c] != - 1{
				min_coins = min_int(1 + dp[i-c],min_coins)
			}
		}
		if min_coins != math.MaxInt32{
			dp[i] = min_coins
		}
	}
	return dp[amount]
}

func dfs_coinChange(coins []int, amount int,record map[int]int)int{
	if amount == 0{
		return 0
	}
	if _,ok := record[amount];ok{
		return record[amount]
	}
	steps := math.MaxInt32
	for _,c := range coins{
		if c > amount{
			continue
		}
		res := dfs_coinChange(coins,amount - c,record)
		if res != -1{
			steps = min_int(res + 1,steps)
		}
	}
	if steps == math.MaxInt32{
		return -1
	}else{
		record[amount] = steps
		return steps
	}
}

func coinChange2(coins []int, amount int)int{
	if amount == 0{
		return 0
	}

	var coin_record map[int]int = make(map[int]int)
	return dfs_coinChange(coins,amount,coin_record)
}

//443
//Input:
//["a","b","b","b","b","b","b","b","b","b","b","b","b"]
//Output:
//Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].
func compress(chars []byte) int {
	var end int = len(chars) - 1
	var begin int = end - 1
	var input_pos int = end
	for begin >= 0{
		if chars[begin] == chars[end] && begin != 0{
			begin--
		}else{
			gap := end - begin
			if begin == 0{
				gap++
			}
			if gap >= 2{
				var c byte = chars[end]
				var input []byte
				for gap > 0{
					input = append(input, byte(gap % 10))
					gap = gap / 10
				}
				for _,v := range input{
					chars[input_pos] = v + '0'
					input_pos--
				}
				chars[input_pos] = c
				input_pos--
			}else{
				chars[input_pos] = chars[end]
				input_pos--
			}
			end = begin
			begin--
		}
	}
	copy(chars,chars[input_pos + 1:])
	return len(chars)
}

//1337
//Input: mat =
//[[1,1,0,0,0],
//[1,1,1,1,0],
//[1,0,0,0,0],
//[1,1,0,0,0],
//[1,1,1,1,1]],
//k = 3
//Output: [2,0,3]
func kWeakestRows(mat [][]int, k int) []int {
	var rows = len(mat)
	if rows == 0{
		return []int{}
	}
	var columns = len(mat[0])
	var record map[int][]int = make(map[int][]int)
	for i := 0;i < rows;i++{
		j := 0
		for ;j < columns;j++ {
			if mat[i][j] == 0 {
				break
			}
		}
		if _,ok := record[j];ok{
			record[j] = append(record[j], i)
		}else{
			record[j] = []int{i}
		}
	}
	var keys []int
	for l := range record {
		keys = append(keys, l)
	}
	sort.Ints(keys)
	var res []int
	var i int = 0
	for _,num := range keys{
		for _,row := range record[num]{
			res = append(res,row)
			i++
			if i == k{
				return res
			}
		}
	}
	return res
}

//394
//s = "3[a]2[bc]", return "aaabcbc".
//s = "3[a2[c]]", return "accaccacc".
//s = "2[abc]3[cd]ef", return "abcabccdcdcdef".
func decodeString(s string) string {
	var word_record []string
	var cnt_record []int
	var visit int = 0
	var word_num int
	var cur_string string
	for visit < len(s){
		if s[visit] == '['{
			cnt_record = append(cnt_record, word_num)
			word_record = append(word_record,cur_string)
			cur_string = ""
			word_num = 0
		}else if s[visit]==']'{
			pre := word_record[len(word_record) - 1]
			cnt := cnt_record[len(cnt_record) - 1]
			for i := 0;i < cnt;i++{
				pre += cur_string
			}
			cur_string = pre
			cnt_record = cnt_record[:len(word_record) - 1]
			word_record = word_record[:len(word_record) - 1]
		}else{
			if s[visit] < '0' || s[visit] > '9'{
				cur_string += string(s[visit])
			}else{
				c,_ := strconv.Atoi(string(s[visit]))
				word_num = word_num * 10 + c
			}
		}
		visit++
	}
	return cur_string
}

func dfs_decodeString2(s string,pos *int)string{
	var word string
	var num int = 0
	for ;*pos < len(s);*pos++{
		if s[*pos] == '['{
			*pos++
			var tmp string = dfs_decodeString2(s,pos)
			for j := 0;j < num;j++{
				word += tmp
			}
			num = 0
		}else if s[*pos] == ']'{
			return word
		}else if s[*pos] >= '0' && s[*pos] <= '9'{
			c,_ := strconv.Atoi(string(s[*pos]))
			num = num * 10 + c
		}else{
			word += string(s[*pos])
		}
	}
	return word
}

func decodeString2(s string) string {
	var pos int = 0
	return dfs_decodeString2(s,&pos)
}

//152
//Input: [-2,0,-1]
//Output: 0
//[2,3,-2,4,-2,4,6,-9,3]
func dp_maxProduct(nums []int,begin int,memo map[int][]int)(min int,max int){
	if begin < 0{
		return 1,1
	}
	if nums[begin] == 0{
		return 0,0
	}
	if _,ok := memo[begin];ok{
		return memo[begin][0],memo[begin][1]
	}
	s,b := dp_maxProduct(nums,begin - 1,memo)
	big := max_int(max_int(b * nums[begin],s * nums[begin]),nums[begin])
	small := min_int(min_int(b * nums[begin],s * nums[begin]),nums[begin])
	memo[begin] = make([]int,2)
	memo[begin][0] = small
	memo[begin][1] = big
	return small,big
}

func maxProduct(nums []int) int {
	var res int = math.MinInt32
	var record map[int][]int = make(map[int][]int)
	for i := 0;i < len(nums) ;i++{
		_,big :=  dp_maxProduct(nums,i,record)
		if big > res{
			res = big
		}
	}
	return res
}

func maxProduct2(nums []int) int{
	l := len(nums)
	if l == 0{
		return 0
	}
	if l == 1{
		return nums[0]
	}
	var dp_min []int = make([]int,l)
	var dp_max []int = make([]int,l)
	dp_min[0] = nums[0]
	dp_max[0] = nums[0]
	var max int = nums[0]
	for i := 1;i < l;i++{
		product1 := nums[i] * dp_min[i-1]
		product2 := nums[i] * dp_max[i-1]
		dp_min[i] = min_int_number(nums[i],product1,product2)
		dp_max[i] = max_int_number(nums[i],product1,product2)
		if dp_max[i] > max{
			max = dp_max[i]
		}
	}
	return max
}

//560
//Input:nums = [1,1,1], k = 2
//Output: 2
func subarraySum(nums []int, k int) int {
	var res int = 0
	var record map[int]int = make(map[int]int)
	var sum int = 0
	for _,n := range nums{
		sum += n
		if sum == k{
			res += 1
		}
		if _,ok := record[sum - k];ok{
			res += record[sum - k]
		}
		if _,ok := record[sum];ok{
			record[sum]++
		} else{
			record[sum] = 1
		}
	}
	return res
}

//309
//Input: [1,2,3,0,2]
//Output: 3
//Explanation: transactions = [buy, sell, cooldown, buy, sell]
func dp_maxProfit(prices []int,cur_pos int,has_stock bool,is_cooldown bool,nostock_memo map[int]int,hasstock_memo map[int]int)int{
	if cur_pos >= len(prices){
		return 0
	}

	var max int = math.MinInt32
	if is_cooldown{
		max = dp_maxProfit(prices,cur_pos + 1,has_stock,false,nostock_memo,hasstock_memo)
	}else{
		if has_stock{
			if _,ok := hasstock_memo[cur_pos];ok{
				return hasstock_memo[cur_pos]
			}
			max = max_int(prices[cur_pos] + dp_maxProfit(prices,cur_pos + 1,false,true,nostock_memo,hasstock_memo),
				dp_maxProfit(prices,cur_pos + 1,true,false,nostock_memo,hasstock_memo))
			hasstock_memo[cur_pos] = max
		}else{
			if _,ok := nostock_memo[cur_pos];ok{
				return nostock_memo[cur_pos]
			}
			max = max_int(dp_maxProfit(prices,cur_pos + 1,true,false,nostock_memo,hasstock_memo) - prices[cur_pos],
				dp_maxProfit(prices,cur_pos + 1,false,false,nostock_memo,hasstock_memo))
			nostock_memo[cur_pos] = max
		}
	}
	return max
}

func maxProfit3(prices []int) int {
	var nostock_record map[int]int = make(map[int]int)
	var hasstock_record map[int]int = make(map[int]int)
	return dp_maxProfit(prices,0,false,false,nostock_record,hasstock_record)
}

//1314
//Input: nums = [1,2,3,4]
//Output: [2,4,4,4]
//Explanation: The first pair [1,2] means we have freq = 1 and val = 2 so we generate the array [2].
//The second pair [3,4] means we have freq = 3 and val = 4 so we generate [4,4,4].
func decompressRLElist(nums []int) []int {
	l := len(nums)
	var total_length int = 0
	for i := 0;i < l;i += 2{
		total_length += nums[i]
	}
	var res []int = make([]int,total_length)
	var cur_pos int = 0
	for i := 0;i < l;i += 2{
		cnt := nums[i]
		val := nums[i + 1]
		for j := 0;j < cnt;j++{
			res[cur_pos] = val
			cur_pos++
		}
	}
	return res
}

//416
func dp_canPartition(nums []int,cur_sum int,cur_pos int,target int,memo map[int]bool)int{
	if cur_pos >= len(nums) || cur_sum > target{
		return 0
	}
	if _,ok := memo[cur_sum];ok{
		return 0
	}
	if cur_sum == target{
		return 1
	}
	cnt := dp_canPartition(nums,cur_sum + nums[cur_pos],cur_pos + 1,target,memo) + dp_canPartition(nums,cur_sum,cur_pos + 1,target,memo)
	if cnt == 0{
		memo[cur_sum] = false
	}
	return cnt
}

func canPartition(nums []int) bool {
	var sum int = 0
	for _,v := range nums{
		sum += v
	}
	if sum % 2 == 1{
		return false
	}
	var record map[int]bool = make(map[int]bool)
	return dp_canPartition(nums,0,0,sum / 2,record) > 0
}

//148
//Input: 4->2->1->3
//Output: 1->2->3->4
//Input: -1->5->3->4->0
//Output: -1->0->3->4->5
func merge_sortList(l1 *ListNode,l2 *ListNode) *ListNode{
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	var head *ListNode
	if l1.Val < l2.Val{
		head = l1
		l1 = l1.Next
	}else{
		head = l2
		l2 = l2.Next
	}
	var visit *ListNode = head
	for l1 != nil && l2 != nil{
		if l1.Val < l2.Val{
			visit.Next = l1
			visit = l1
			l1 = l1.Next
		}else{
			visit.Next = l2
			visit = l2
			l2 = l2.Next
		}
	}
	if l1 == nil{
		visit.Next = l2
	}else if l2 == nil{
		visit.Next = l1
	}
	return head
}
//-1->5->3->4->0
func sortList(head *ListNode) *ListNode {
	if nil == head || nil == head.Next{
		return head
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next
		if fast != nil{
			fast = fast.Next
		}
	}
	head2 := slow.Next
	slow.Next = nil
	newhead1 := sortList(head)
	newhead2 := sortList(head2)
	return merge_sortList(newhead1,newhead2)
}

//139
//Input: s = "applepenapple", wordDict = ["apple", "pen"]
//Output: true
//Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
//Note that you are allowed to reuse a dictionary word.
//Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//Output: false
func dfs_wordBreak(s string,wordDict []string,cur_pos int,memo map[int]bool)bool{
	if cur_pos >= len(s){
		return true
	}
	if _,ok := memo[cur_pos];ok{
		return false
	}
	sub := s[cur_pos:]
	for _,w := range wordDict{
		if strings.HasPrefix(sub,w){
			if dfs_wordBreak(s,wordDict,cur_pos + len(w),memo){
				return true
			}
		}
	}
	memo[cur_pos] = false
	return false
}

func wordBreak(s string, wordDict []string) bool {
	var record map[int]bool = make(map[int]bool)
	return dfs_wordBreak(s,wordDict,0,record)
}

//146
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
type DataNode struct{
	key int
	value int
}

type LRUCache struct {
	data []DataNode
	capacity int
	cur_size int
}

func Constructor146(capacity int) LRUCache {
	var obj LRUCache
	obj.data = make([]DataNode,capacity)
	obj.capacity = capacity
	obj.cur_size = 0
	return obj
}

func (this *LRUCache) Get(key int) int {
	for i := this.cur_size - 1; i >= 0;i--{
		if this.data[i].key == key{
			tmp := this.data[i]
			for j := i;j < this.cur_size - 1;j++{
				this.data[j] = this.data[j + 1]
			}
			this.data[this.cur_size - 1] = tmp
			return tmp.value
		}
	}
	return -1
}

func (this *LRUCache) Put(key int, value int)  {
	for i := this.cur_size - 1; i >= 0;i--{
		if this.data[i].key == key{
			this.data[i].value = value
			var tmp DataNode
			tmp.key = this.data[i].key
			tmp.value = value
			for j := i;j < this.cur_size - 1;j++{
				this.data[j] = this.data[j + 1]
			}
			this.data[this.cur_size - 1] = tmp
			return
		}
	}

	if this.cur_size == this.capacity{
		this.data = append(this.data[:0],this.data[1:]...)
		this.data = append(this.data,DataNode{key,value})
	}else{
		this.cur_size++
		this.data[this.cur_size - 1] = DataNode{key,value}
	}
}

//1356

type sortBit []int

func (s sortBit) Less(i, j int) bool {
	var bit_cnt1 int = 0
	var bit_cnt2 int = 0
	var move int = 1
	for s[i] >= move{
		if (s[i] | move) == s[i]{
			bit_cnt1++
		}
		move = move << 1
	}
	move = 1
	for s[j] >= move{
		if (s[j] | move) == s[j]{
			bit_cnt2++
		}
		move = move << 1
	}
	if bit_cnt1 == bit_cnt2{
		return s[i] < s[j]
	}
	return bit_cnt1 < bit_cnt2
}

func (s sortBit) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortBit) Len() int {
	return len(s)
}

func sortByBits(arr []int) []int {
	sort.Sort(sortBit(arr))
	return arr
}

//334
//Return true if there exists i, j, k
//such that arr[i] < arr[j] < arr[k] given 0 ≤ i < j < k ≤ n-1 else return false.
//Note: Your algorithm should run in O(n) time complexity and O(1) space complexity.
func increasingTriplet(nums []int) bool {
	l := len(nums)
	if l <= 2{
		return false
	}
	var min int = nums[0]
	var mid int = math.MaxInt32
	for i := 1;i < len(nums);i++ {
		if nums[i] > mid{
			return true
		}
		if nums[i] > min{
			mid = nums[i]
		}else{
			min = nums[i]
		}
	}
	return false;
}

//678
func checkValidString(s string) bool {
	var left []int
	var star []int
	for i := 0;i < len(s);i++{
		if s[i] == '('{
			left = append(left, i)
		}else if s[i] == '*'{
			star = append(star,i)
		}else{
			if len(left) == 0 && len(star) == 0{
				return false
			}
			if len(left) > 0{
				left = left[:len(left) - 1]
			}else if len(star) > 0{
				star = star[:len(star) - 1]
			}
		}
	}
	for len(left) > 0 && len(star) > 0{
		l := left[len(left) - 1]
		s := star[len(star) - 1]
		left = left[:len(left) - 1]
		star = star[:len(star) - 1]
		if l > s{
			return false
		}
	}
 	if len(left) > 0{
		return false
	}
	return true
}

//1370
//Input: s = "aaaabbbbcccc"
//Output: "abccbaabccba"
func sortString(s string) string {
	var record [26]int
	for _,c := range s{
		record[c-'a']++
	}
	var head_to_tail bool = true
	var res string
	head,tail := 0,25;
	for head <= tail {
		if record[head] == 0{
			head++
			continue
		}
		if record[tail] == 0{
			tail--
			continue
		}
		i := head
		j := tail
		for i <= j{
			if head_to_tail{
				if record[i] == 0{
					i++
					continue
				}
			}else{
				if record[j] == 0{
					j--
					continue
				}
			}
			if head_to_tail{
				res += string(i + 'a')
				record[i]--
				i++
			}else{
				res += string(j + 'a')
				record[j]--
				j--
			}
		}
		head_to_tail = !head_to_tail
	}
	return res
}

//1374
func generateTheString(n int) string {
	var res string
	if n % 2 == 1{
		for i := 0;i < n;i++{
			res += "a"
		}
		return res
	}else{
		for i := 0;i < n - 1;i++{
			res += "a"
		}
		res += "b"
		return res
	}
}

//23
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func merge_sort(list1 *ListNode,list2 *ListNode)*ListNode{
	if list1 == nil{
		return list2
	}
	if list2 == nil{
		return list1
	}
	var head *ListNode = nil
	var visit *ListNode = nil
	for list1 != nil && list2 != nil{
		if list1.Val < list2.Val{
			if head == nil{
				head = list1
				visit = list1
			}else{
				visit.Next = list1
				visit = list1
			}
			list1 = list1.Next
		}else{
			if head == nil{
				head = list2
				visit = list2
			}else{
				visit.Next = list2
				visit = list2
			}
			list2 = list2.Next
		}
	}
	for list1 != nil{
		visit.Next = list1
		visit = list1
		list1 = list1.Next
	}
	for list2 != nil{
		visit.Next = list2
		visit = list2
		list2 = list2.Next
	}
	return head
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0{
		return nil
	}
	var prev_lists []*ListNode = make([]*ListNode,len(lists))
	copy(prev_lists,lists)
	for len(prev_lists) > 1{
		var lists_merge []*ListNode
		l := len(prev_lists)
		for i := 0;i < l - 1;i += 2{
			new_list := merge_sort(prev_lists[i],prev_lists[i+1])
			lists_merge = append(lists_merge, new_list)
		}
		if l % 2 == 1{
			lists_merge = append(lists_merge,prev_lists[l - 1])
		}
		prev_lists = lists_merge
	}
	return prev_lists[0]
}

//1282
//Input: groupSizes = [3,3,3,3,3,1,3]
//Output: [[5],[0,1,2],[3,4,6]]
//Explanation:
//Other possible solutions are [[2,1,6],[5],[0,4,3]] and [[5],[0,6,2],[4,3,1]].
func groupThePeople(groupSizes []int) [][]int {
	//save the array index by size
	var record map[int][] int = make(map[int][]int)
	for i,s := range groupSizes{
		if _,ok := record[s];ok{
			record[s] = append(record[s], i)
		}else{
			record[s] = []int{i}
		}
	}
	var res [][]int
	for length,nums := range record{
		var collection []int
		for i := 0;i < len(nums);i++{
			collection = append(collection, nums[i])
			if i % length == length - 1{
				res = append(res, collection)
				collection = []int{}
			}
		}
	}
	return res
}

//1329
func diagonalSort(mat [][]int) [][]int {
	rows := len(mat)
	if rows <= 1{
		return mat
	}
	columns := len(mat[0])
	if columns <= 1{
		return mat
	}
	for r := rows - 1;r >= 0;r--{
		for i := 0;i < r;i++{
			for j := 0;j < columns - 1;j++{
				if mat[i][j] > mat[i + 1][j + 1]{
					mat[i][j], mat[i + 1][j + 1] = mat[i + 1][j + 1],mat[i][j]
				}
			}
		}
	}
	return mat
}

//
func luckyNumbers (matrix [][]int) []int {
	var rows int = len(matrix)
	var columns int = len(matrix[0])
	var res []int
	for i := 0;i < rows;i++{
		var min_num_column int = 0
		var cur_row_min int = matrix[i][0]
		for j := 0;j < columns;j++{
			if matrix[i][j] < cur_row_min{
				cur_row_min = matrix[i][j]
				min_num_column = j
			}
		}

		var pass bool = true
		for j := 0;j < rows;j++{
			if j == i{
				continue
			}
			if matrix[j][min_num_column] > matrix[i][min_num_column]{
				pass = false
				break
			}
		}
		if pass{
			res = append(res, matrix[i][min_num_column])
		}
	}
	return res
}

//
type CustomStack struct {
	data []int
	max_size int
	cur_size int
}


func Constructor(maxSize int) CustomStack {
	var obj CustomStack
	obj.max_size = maxSize
	return obj
}

func (this *CustomStack) Push(x int)  {
	if this.cur_size < this.max_size{
		this.data = append(this.data,x)
		this.cur_size++
	}
}


func (this *CustomStack) Pop() int {
	if this.cur_size > 0{
		res := this.data[this.cur_size - 1]
		this.cur_size--
		this.data = this.data[:this.cur_size]
		return res
	}
	return -1
}


func (this *CustomStack) Increment(k int, val int)  {
	for i := 0;i < k && i < this.cur_size;i++{
		this.data[i] += val
	}
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */

//72
func max_rectange(arr [][]int)int{
	rows := len(arr)
	columns := len(arr[0])
	var dp [][]int = make([][]int,rows)
	for i := 0; i < rows;i++{
		dp[i] = make([]int,columns)
	}
	for i := 0;i < rows;i++{
		dp[i][0] = arr[i][0]
	}
	for j := 0;j < columns;j++{
		dp[0][j] = arr[j][0]
	}
	var max int = 0
	for i := 1;i < rows;i++{
		for j := 1;j < columns;j++{
			if arr[i][j] == 0{
				continue
			}
			if dp[i - 1][j - 1] != 0 && dp[i - 1][j] != 0 && dp[i][j - 1] != 0{
				dp[i][j] = min_int_number(dp[i-1][j-1],dp[i - 1][j],dp[i][j - 1]) + 1
			}else{
				dp[i][j] = 1
			}
			if dp[i][j] > max{
				max = dp[i][j]
			}
		}
	}
	return max
}

//3
func lengthOfLongestSubstring(s string) int {
	var start int = 0
	var end int = 0
	var max int = 0
	var cur_len int = 0
	var record map[uint8]int = make(map[uint8]int)
	for end < len(s) && start < len(s){
		if _,ok := record[s[end]];ok{
			start = max_int(record[s[end]] + 1,start)
			record[s[end]] = end
			if cur_len > max{
				max = cur_len
			}
			cur_len = end - start + 1
			end++
		}else{
			cur_len++
			record[s[end]] = end
			end++
		}
	}
	if (end - start) > max{
		max = (end - start)
	}
	return max
}

//324
//Input: nums = [1, 5, 1, 1, 6, 4]
//Output: One possible answer is [1, 4, 1, 5, 1, 6].
func qsort_wiggleSort(nums []int,low int,high int,mid int){
	if low >= high{
		return
	}

	l := low
	h := high
	tag := nums[l]
	for l < h{
		for l < h && nums[h] > tag{
			h--
		}
		if l < h{
			nums[l] = nums[h]
			l++
		}
		for l < h && nums[l] < tag{
			l++
		}
		if l < h{
			nums[h] = nums[l]
			h--
		}
	}
	nums[l] = tag
	if l == mid{
		return
	}
	if l < mid{
		qsort_wiggleSort(nums,l+1,high,mid)
	}else{
		qsort_wiggleSort(nums,low,l-1,mid)
	}
}

//969
//Input: [3,2,4,1]
//Output: [4,2,4,3]
//Explanation:
//We perform 4 pancake flips, with k values 4, 2, 4, and 3.
//Starting state: A = [3, 2, 4, 1]
//After 1st flip (k=4): A = [1, 4, 2, 3]
//After 2nd flip (k=2): A = [4, 1, 2, 3]
//After 3rd flip (k=4): A = [3, 2, 1, 4]
//After 4th flip (k=3): A = [1, 2, 3, 4], which is sorted.
func reverse_pancakeSort(A []int,begin int,end int){
	for begin < end{
		A[begin],A[end] = A[end],A[begin]
		begin++
		end--
	}
}

func pancakeSort(A []int) []int {
	l := len(A)
	var res []int
	for i := l - 1;i >= 0;i--{
		for j := 0;j <= i;j++{
			if A[j] == i + 1{
				if j != 0{
					res = append(res,j + 1)
					reverse_pancakeSort(A,0,j)
				}
				if i != 0{
					reverse_pancakeSort(A,0,i)
					res = append(res,i + 1)
				}
				break
			}
		}
	}
	return res
}

//1343
//Input: arr = [2,2,2,2,5,5,5,8], k = 3, threshold = 4
//Output: 3
//Explanation: Sub-arrays [2,5,5],[5,5,5] and [5,5,8] have averages 4, 5 and 6 respectively.
//All other sub-arrays of size 3 have averages less than 4 (the threshold).
func numOfSubarrays(arr []int, k int, threshold int) int {
	l := len(arr)
	if l < k {
		return 0
	}
	var target int = threshold * k
	var sum int = 0
	for i := 0;i < k;i++{
		sum += arr[i]
	}
	var res int = 0
	begin := 0
	end := begin + k - 1
	for end < l{
		if sum >= target{
			res++
		}
		end++
		if end < l{
			sum -= arr[begin]
			sum += arr[end]
		}
		begin++
	}
	return res
}