package main

import (
	"../array"
	"../diagram"
	"../list"
	"../number"
	"../string"
	"../tree"
	"fmt"
)

func main(){
	{
		n := 12
		primes := []int{2,7,13,19}
		res := number.NthSuperUglyNumber2(n,primes)
		fmt.Println(res)
	}
	{
		input := "0235813"
		res := number.IsAdditiveNumber(input)
		fmt.Println(res)
	}
	{
		input := [][]int{{3,0,1,4,2},{5,6,3,2,1},{1,2,0,1,5},{4,1,0,1,7},{1,0,3,0,5}}
		obj := array.Constructor304(input)
		res := obj.SumRegion(2,1,4,3)
		fmt.Print(res)
	}
	{
		res := number.NthUglyNumber(1)
		fmt.Print(res)
	}
	{
		input := []int{0,0,0,0}
		target := 0
		res := number.FourSum(input,target)
		fmt.Println(res)
	}
	{
		input := [][]int{{1,2,3,4,5},{6,7},{8},{9,10,11},{12,13,14,15,16}}
		res := array.FindDiagonalOrder(input)
		fmt.Println(res)
	}
	{
		input := []int{1,79,80,1,1,1,200,1}
		k := 6
		res := number.MaxScore2(input,k)
		fmt.Println(res)
	}
	{
		input := "10"
		res := number.MaxScore(input)
		fmt.Println(res)
	}
	{
		input := []int{3,2,3}
		res := number.MajorityElement(input)
		fmt.Println(res)
	}
	{
		input := []int{0,2,3,4,6,8,9}
		res := array.SummaryRanges(input)
		fmt.Println(res)
	}
	{
		input := []int{2,2}
		k := 3
		t := 0
		res := array.ContainsNearbyAlmostDuplicate(input,k,t)
		fmt.Println(res)
	}
	{
		k := 22
		res := number.FindMinFibonacciNumbers(k)
		fmt.Println(res)
	}
	{
		//["MyCircularQueue","enQueue","Rear","Rear","deQueue","enQueue","Rear",
		var obj list.MyCircularQueue = list.Constructor(8)

		res := obj.EnQueue(3)
		res = obj.EnQueue(9)
		res = obj.EnQueue(5)
		res = obj.EnQueue(0)
		res = obj.DeQueue()
		res = obj.DeQueue()
		res = obj.IsEmpty()
		res = obj.IsEmpty()
		fmt.Println(res)
	}
	{
		m := 8
		n := 50
		N := 23
		i := 5
		j := 26
		res := diagram.FindPaths(m, n, N, i, j)
		fmt.Println(res)
	}
	{
		k := 3
		n := 9
		res := number.CombinationSum3(k,n)
		fmt.Println(res)
	}
	{
		input := []int{114,117,207,117,235,82,90,67,143,146,53,108,200,91,80,223,58,170,110,236,81,90,222,160,165,195,187,199,114,235,197,187,69,129,64,214,228,78,188,67,205,94,205,169,241,202,144,240}
		res := array.Rob(input)
		fmt.Println(res)
	}
	{
		var obj tree.Trie  = tree.Constructor208()

		obj.Insert("apple");
		res := obj.Search("apple");   // returns true
		res = obj.Search("app");     // returns false
		res = obj.StartsWith("app"); // returns true
		obj.Insert("app");
		res = obj.Search("app");     // returns true
		fmt.Println(res)
	}
	{
		obj := diagram.Constructor();
		obj.AddWord("and")
		obj.AddWord("abd")
		obj.AddWord("ab")
		res := obj.Search(".and")
		res = obj.Search(".bcd")
		fmt.Println(res)
	}
	{
		input := []int{9,8,7,4,3,2,1,6}
		array.Recursive_qsort(input,0,len(input) - 1)
		fmt.Println(input)
	}
	{
		nums := []int{1,2,3,4,5}
		s := 11
		res := array.MinSubArrayLen2(s,nums)
		fmt.Println(res)
	}
	{
		input := "  hello   world!  "
		res := string.ReverseWords(input)
		fmt.Println(res)
	}
	{
		//[[2,4],[1,3],[2,4],[1,3]]
		var node1 diagram.Node
		node1.Val = 1
		var node2 diagram.Node
		node2.Val = 2
		var node3 diagram.Node
		node3.Val = 3
		var node4 diagram.Node
		node4.Val = 4
		node1.Neighbors = append(node1.Neighbors,&node2,&node4)
		node2.Neighbors = append(node1.Neighbors,&node1,&node3)
		node3.Neighbors = append(node2.Neighbors,&node2,&node4)
		node4.Neighbors = append(node2.Neighbors,&node1,&node3)
		res := diagram.CloneGraph(&node1)
		fmt.Println(res.Val)
	}
	{
		//-1->5->3->4->0
		var l1 list.ListNode
		l1.Val = 4
		var l2 list.ListNode
		l2.Val = 2
		var l3 list.ListNode
		l3.Val = 1
		var l4 list.ListNode
		l4.Val = 3
		var l5 list.ListNode
		l5.Val = 0
		l1.Next = &l2
		l2.Next = &l3
		l3.Next = &l4
		//l4.Next = &l5
		res := list.InsertionSortList(&l1)
		fmt.Println(res.Val)
	}
	{
		input := []int{0,1,0,1,0,1,99}
		res := number.SingleNumber(input)
		fmt.Println(res)
	}
	{
		var list1 list.ListNode
		list1.Val = 1
		var list2 list.ListNode
		list2.Val = 2
		var list3 list.ListNode
		list3.Val = 3
		var list4 list.ListNode
		list4.Val = 4
		var list5 list.ListNode
		list5.Val = 5
		//list1.Next = &list2
		//list2.Next = &list3
		//list3.Next = &list4
		//list4.Next = &list5
		list.ReorderList(&list1)
		fmt.Println(list1.Val)
	}
	{
		input := [][]int{{2},{3,4},{6,5,7},{4,1,8,3}}
		//res := number.MinimumTotal(input)
		res := number.MinimumTotal2(input)
		fmt.Println(res)
	}
	{
		var t1 tree.TreeNode
		t1.Val = 1
		var t2 tree.TreeNode
		t2.Val = 2
		var t3 tree.TreeNode
		t3.Val = 3
		var t4 tree.TreeNode
		t4.Val = 4
		var t5 tree.TreeNode
		t5.Val = 5
		t1.Left = &t2
		t1.Right = &t3
		t2.Left = &t4
		t2.Right = &t5
		tree.Inorder_visit_norecursive(&t1)
		//tree.Preorder_visit_norecursive(&t1)
	}
	{
		res := number.AngleClock(4,50)
		fmt.Println(res)
	}
	{
		input := [][]int{{1,6},{4,6},{4,8}}
		res := array.RemoveCoveredIntervals(input)
		fmt.Println(res)
	}
	{
		input := []int{2,4,7,8,9,10,14,15,18,23,32,50}
		//res := number.LenLongestFibSubseq(input)
		res := number.LenLongestFibSubseq2(input)
		fmt.Println(res)
	}
	{
		res := number.NumOfWays(5000)
		fmt.Println(res)
	}
	{
		var t1 tree.TreeNode
		t1.Val = 1
		var t2 tree.TreeNode
		t2.Val = 2
		var t3 tree.TreeNode
		t3.Val = 3
		var t4 tree.TreeNode
		t4.Val = 4
		var t5 tree.TreeNode
		t5.Val = 5
		t1.Left = &t2
		t1.Right = &t3
		t2.Left = &t4
		t2.Right = &t5
		res := tree.PathSum2(nil,22)
		fmt.Println(res)
	}
	{
		input := 3
		res := number.IntToRoman(input)
		fmt.Println(res)
	}
	{
		//input := "/a//b////c/d//././/.."
		input := "/a/../../b/../c//.//"
		res := string.SimplifyPath(input)
		fmt.Println(res)
	}
	{
		input := 10
		res := tree.GenerateTrees(input)
		fmt.Println(res)
	}
	{
		input := []int{2,5,7,9}
		target := 16
		res := number.TwoSum2(input,target)
		fmt.Println(res)
	}
	//{
	//	s := "PAYPALISHIRING"
	//	numRows := 3
	//	res := convert(s,numRows)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{-2,-3,-1}
	//	res := maxSubarraySumCircular(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{-2,1,-3,4,-1,2,1,-5,4}
	//	res := maxSubArray(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{4,4,4,4}
	//	k := 4
	//	threshold := 1
	//	res := numOfSubarrays(input,k,threshold)
	//	fmt.Println(res)
	//}
	//{
	//	pre := []int{1,2,4,5,3,6,7}
	//	post := []int{4,5,2,6,7,3,1}
	//	res := constructFromPrePost(pre,post)
	//	fmt.Println(res)
	//}
	//{
	//	input := []string{ "//", "/ "}
	//	res := regionsBySlashes(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{4,6,7,7}
	//	res := findSubsequences(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{3,2,4,1}
	//	res := pancakeSort(input)
	//	fmt.Println(res)
	//}
	//{
	//	var t1 TreeNode
	//	t1.Val = 1
	//	var t2 TreeNode
	//	t2.Val = 2
	//	var t3 TreeNode
	//	t3.Val = 3
	//	var t4 TreeNode
	//	t4.Val = 4
	//	//var t5 TreeNode
	//	//t5.Val = 5
	//	t1.Left = &t2
	//	t1.Right = &t3
	//	t2.Left = &t4
	//	//t2.Right = &t5
	//	res := lcaDeepestLeaves(&t1)
	//	fmt.Println(res)
	//}
	//{
	//	input := [][]int{{5,6},{1,0}}
	//	res := minFlips(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := "1-2--3---4-5--6---7"
	//	res := recoverFromPreorder(input)
	//	fmt.Println(res)
	//}
	//{
	//	res := fractionToDecimal(2,4)
	//	fmt.Println(res)
	//}
	//{
	//	res := myPow(3,4)
	//	fmt.Println(res)
	//}
	//{
	//	//input := []int{1, 3, 2, 1, 3, 2}
	//	//wiggleSort(input)
	//	//fmt.Println("over")
	//}
	//{
	//	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	//	wordDict := []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}
	//	res := wordBreak2(s,wordDict)
	//	fmt.Println(res)
	//}
	//{
	//	input := "cbcccbbccbbacbcaccbacabacacbcccbbccbbacb"
	//	res := longestPrefix(input)
	//	fmt.Println(res)
	//}
	//{
	//
	//	board := [][]byte{{'a'}}
	//	words := []string{"a"}
	//	res := findWords2(board,words)
	//	fmt.Println(res)
	//}
	//{
	//	var input [][]int = [][]int{{1,0,0,0},{0,0,0,0},{0,0,0,0},{0,0,0,0},{0,0,0,2}}
	//	res := uniquePathsIII(input)
	//	fmt.Println(res)
	//}
	//{
	//	s := "ab"
	//	p := ".*"
	//	res := isMatch(s,p)
	//	fmt.Println(res)
	//}
	//{
	//	input := ")()())"
	//	res := longestValidParentheses(input)
	//	res = dp_longestValidParentheses(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{3,4,-1,1}
	//	res := firstMissingPositive(input)
	//	fmt.Println(res)
	//}
	//{
	//	var t1 TreeNode
	//	t1.Val = -2
	//	var t2 TreeNode
	//	t2.Val = 6
	//	var t3 TreeNode
	//	t3.Val = 0
	//	var t4 TreeNode
	//	t4.Val = -6
	//	t1.Left = &t2
	//	t2.Left = &t3
	//	t2.Right = &t4
	//	res := maxPathSum(&t1)
	//	fmt.Println(res)
	//}
	//{
	//	lo := 1
	//	hi := 1000
	//	k := 777
	//	res := getKth(lo,hi,k)
	//	fmt.Println(res)
	//}
	//{
	//	arr1 := []int{4,5,8}
	//	arr2 := []int{10,9,1,8}
	//	d := 2
	//	res := findTheDistanceValue(arr1,arr2,d)
	//	fmt.Println(res)
	//}
	//{
	//	s := "a"
	//	t := "aa"
	//	res := minWindow(s,t)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{4,2,0,3,2,5}
	//	res := largestRectangleArea(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{73, 74, 75, 71, 69, 72, 76, 73}
	//	res := dailyTemperatures(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{2,0,1}
	//	res := countSmaller(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := "ababacb"
	//	k := 3
	//	res := longestSubstring(input,k)
	//	fmt.Println(res)
	//}
	//{
	//	input := []string{"2", "1", "+", "3", "*"}
	//	res := evalRPN(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := "ac"
	//	res := longestPalindrome(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := "abcabcabc"
	//	res := lengthOfLongestSubstring(input)
	//	fmt.Println(res)
	//}
	//{
	//	//tree_nums :=0
	//	//var trees []int
	//	//fmt.Scan(&tree_nums)
	//	//for i := 0; i < tree_nums; i++ {
	//	//	x:=0
	//	//	fmt.Scan(&x)
	//	//	trees = append(trees,x)
	//	//}
	//	//var hours_limit int
	//	//fmt.Scan(&hours_limit)
	//	//
	//	//var total int
	//	//for _,n := range trees{
	//	//	total += n
	//	//}
	//	//var least_eat_speed int = int(math.Ceil(float64(total) / float64(hours_limit)))
	//	//var cost_times int = hours_limit + 1
	//	//for cost_times > hours_limit{
	//	//	cost_times = 0
	//	//	for i := 0;i < len(trees);i++{
	//	//		var peachs int = trees[i]
	//	//		for peachs > 0{
	//	//			cost_times++
	//	//			peachs -= least_eat_speed
	//	//		}
	//	//	}
	//	//	if cost_times <= hours_limit{
	//	//		break
	//	//	}
	//	//	least_eat_speed++
	//	//}
	//	//fmt.Println(least_eat_speed)
	//}
	//{
	//	input := [][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}
	//	res := maximalRectangle(input)
	//	fmt.Println(res)
	//}
	//{
	//	s := "()())()"
	//	res := removeInvalidParentheses(s)
	//	fmt.Println(res)
	//}
	//{
	//	//word1 := "horse"
	//	//word2 := "ros"
	//	word1 := "ab"
	//	word2 := "c"
	//	res := minDistance(word1,word2)
	//	fmt.Println(res)
	//}
	//{
	//	var t1 TreeNode
	//	t1.Val = 1
	//	var t2 TreeNode
	//	t2.Val = 2
	//	var t3 TreeNode
	//	t3.Val = 3
	//	var t4 TreeNode
	//	t4.Val = 4
	//	var t5 TreeNode
	//	t5.Val = 5
	//	t1.Right = &t2
	//	t2.Right = &t3
	//	t3.Right = &t4
	//	t4.Right = &t5
	//	res := balanceBST(&t1)
	//	fmt.Println(res.Val)
	//}
	//{
	//	var customStack CustomStack = Constructor(3) // Stack is Empty []
	//	customStack.Push(1)                          // stack becomes [1]
	//	customStack.Push(2)                          // stack becomes [1, 2]
	//	res := customStack.Pop()                            // return 2 --> Return top of the stack 2, stack becomes [1]
	//	customStack.Push(2)                         // stack becomes [1, 2]
	//	customStack.Push(3)                          // stack becomes [1, 2, 3]
	//	customStack.Push(4)                          // stack still [1, 2, 3], Don't add another elements as size is 4
	//	customStack.Increment(5, 100)               // stack becomes [101, 102, 103]
	//	customStack.Increment(2, 100)                // stack becomes [201, 202, 103]
	//	res = customStack.Pop()                            // return 103 --> Return top of the stack 103, stack becomes [201, 202]
	//	res = customStack.Pop()                            // return 202 --> Return top of the stack 102, stack becomes [201]
	//	res = customStack.Pop()                            // return 201 --> Return top of the stack 101, stack becomes []
	//	res = customStack.Pop()                            // return -1 --> Stack is empty return -1.
	//	fmt.Println(res)
	//}
	//{
	//	input := [][]int{{7,8},{1,2}}
	//	res := luckyNumbers(input)
	//	fmt.Println(res)
	//}
	//{
	//	nums := 5
	//	input := [][]int{{1,0},{2,0},{3,1},{3,2}}
	//	res := findOrder(nums,input)
	//	fmt.Println(res)
	//}
	//{
	//	courses := 2
	//	input := [][]int{{1,0}}
	//	res := canFinish(courses,input)
	//	fmt.Println(res)
	//}
	//{
	//	input := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	//	res := matrixBlockSum(input,1)
	//	fmt.Println(res)
	//}
	//{
	//	res := allPossibleFBT(7)
	//	fmt.Println(len(res))
	//}
	//{
	//	input := [][]int{{3,3,1,1},{2,2,1,2},{1,1,1,2}}
	//	res := diagonalSort(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{3,3,3,3,3,1,3}
	//	res := groupThePeople(input)
	//	fmt.Println(res)
	//}
	//{
	//	//[
	//	//	1->4->5,
	//	//	1->3->4,
	//	//	2->6
	//	//]
	//	var lists []*ListNode
	//	var l1 ListNode
	//	l1.Val = 1
	//	var l11 ListNode
	//	l11.Val = 4
	//	var l12 ListNode
	//	l12.Val = 5
	//	l1.Next = &l11
	//	l11.Next = &l12
	//	lists = append(lists, &l1)
	//
	//	var l2 ListNode
	//	l2.Val = 1
	//	var l21 ListNode
	//	l21.Val = 3
	//	var l22 ListNode
	//	l22.Val = 4
	//	l2.Next = &l21
	//	l21.Next = &l22
	//	lists = append(lists,&l2)
	//
	//	var l3 ListNode
	//	l3.Val = 2
	//	var l31 ListNode
	//	l31.Val = 6
	//	l3.Next = &l31
	//	lists = append(lists,&l3)
	//	res := mergeKLists(lists)
	//	fmt.Println(res)
	//}
	//{
	//	res := generateTheString(4)
	//	fmt.Println(res)
	//}
	//{
	//	s := ""
	//	res := sortString(s)
	//	fmt.Println(res)
	//}
	//{
	//	s := "(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())"
	//	res := checkValidString(s)
	//	fmt.Println(res)
	//}
	//{
	//	obj := Constructor933();
	//	res := obj.Ping(1);
	//	fmt.Println(res)
	//	res = obj.Ping(100);
	//	fmt.Println(res)
	//	res = obj.Ping(3001);
	//	fmt.Println(res)
	//	res = obj.Ping(3002);
	//	fmt.Println(res)
	//	res = obj.Ping(3007);
	//	fmt.Println(res)
	//	res = obj.Ping(3120);
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{1,6,5,4,3,2,1}
	//	res := findPeakElement(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := " 3+5 / 2 - 3"
	//	res := calculate(input)
	//	fmt.Println(res)
	//}
	//{
	//	nums := []int{30,1,20,21,15,19,2,-1,3}
	//	res := increasingTriplet(nums)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{3,1,5,8}
	//	res := maxCoins(input)
	//	fmt.Println(res)
	//}
	//{
	//	//input := []int{3,3,3,3,5,5,5,2,2,7}
	//	//input := []int{7,7,7,7,7,7}
	//	input := []int{1,2,3,4,5,6,7,8,9,10}
	//	res := minSetSize(input)
	//	fmt.Println(res)
	//}
	//{
	//	res := closestDivisors(208656121)
	//	fmt.Println(res)
	//}
	//{
	//	n := 4
	//	left := []int{1,-1,3,-1}
	//	right := []int{2,-1,-1,-1}
	//	res := validateBinaryTreeNodes(n,left,right)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{10,100,1000,10000}
	//	res := sortByBits(input)
	//	fmt.Println(res)
	//}
	//{
	//	num := "112"
	//	k := 1
	//	res := removeKdigits(num,k)
	//	fmt.Println(res)
	//}
	//{
	//	//["LRUCache","put","put","put","put","get","get"]
	//	//[[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
	//	obj := Constructor146(2);
	//	obj.Put(2, 1);
	//	obj.Put(2, 2);
	//	res := obj.Get(2);
	//	obj.Put(1, 1);
	//	obj.Put(4, 1);
	//	res = obj.Get(2);
	//
	//	//obj.Put(2,1)
	//	//obj.Put(1,1)
	//	//obj.Put(2,3)
	//	//obj.Put(4,1)
	//	//res := obj.Get(1);
	//	//res = obj.Get(2);
	//	fmt.Println(res)
	//}
	//{
	//	//s := "catsandog"
	//	//wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	//	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	//	wordDict := []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}
	//	res := wordBreak(s,wordDict)
	//	fmt.Println(res)
	//}
	//{
	//	//-1->5->3->4->0
	//	var l1 ListNode
	//	l1.Val = 1
	//	var l2 ListNode
	//	l2.Val = 5
	//	var l3 ListNode
	//	l3.Val = 3
	//	var l4 ListNode
	//	l4.Val = 4
	//	var l5 ListNode
	//	l5.Val = 0
	//	l1.Next = &l2
	//	l2.Next = &l3
	//	l3.Next = &l4
	//	l4.Next = &l5
	//	res := sortList(&l1)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{1,2,3,4}
	//	res := decompressRLElist(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{70,4,83,56,94,72,78,43,2,86,65,100,94,56,41,66,3,33,10,3,45,94,15,12,78,60,58,0,58,15,21,7,11,41,12,96,83,77,47,62,27,19,40,63,30,4,77,52,17,57,21,66,63,29,51,40,37,6,44,42,92,16,64,33,31,51,36,0,29,95,92,35,66,91,19,21,100,95,40,61,15,83,31,55,59,84,21,99,45,64,90,25,40,6,41,5,25,52,59,61,51,37,92,90,20,20,96,66,79,28,83,60,91,30,52,55,1,99,8,68,14,84,59,5,34,93,25,10,93,21,35,66,88,20,97,25,63,80,20,86,33,53,43,86,53,55,61,77,9,2,56,78,43,19,68,69,49,1,6,5,82,46,24,33,85,24,56,51,45,100,94,26,15,33,35,59,25,65,32,26,93,73,0,40,92,56,76,18,2,45,64,66,64,39,77,1,55,90,10,27,85,40,95,78,39,40,62,30,12,57,84,95,86,57,41,52,77,17,9,15,33,17,68,63,59,40,5,63,30,86,57,5,55,47,0,92,95,100,25,79,84,93,83,93,18,20,32,63,65,56,68,7,31,100,88,93,11,43,20,13,54,34,29,90,50,24,13,44,89,57,65,95,58,32,67,38,2,41,4,63,56,88,39,57,10,1,97,98,25,45,96,35,22,0,37,74,98,14,37,77,54,40,17,9,28,83,13,92,3,8,60,52,64,8,87,77,96,70,61,3,96,83,56,5,99,81,94,3,38,91,55,83,15,30,39,54,79,55,86,85,32,27,20,74,91,99,100,46,69,77,34,97,0,50,51,21,12,3,84,84,48,69,94,28,64,36,70,34,70,11,89,58,6,90,86,4,97,63,10,37,48,68,30,29,53,4,91,7,56,63,22,93,69,93,1,85,11,20,41,36,66,67,57,76,85,37,80,99,63,23,71,11,73,41,48,54,61,49,91,97,60,38,99,8,17,2,5,56,3,69,90,62,75,76,55,71,83,34,2,36,56,40,15,62,39,78,7,37,58,22,64,59,80,16,2,34,83,43,40,39,38,35,89,72,56,77,78,14,45,0,57,32,82,93,96,3,51,27,36,38,1,19,66,98,93,91,18,95,93,39,12,40,73,100,17,72,93,25,35,45,91,78,13,97,56,40,69,86,69,99,4,36,36,82,35,52,12,46,74,57,65,91,51,41,42,17,78,49,75,9,23,65,44,47,93,84,70,19,22,57,27,84,57,85,2,61,17,90,34,49,74,64,46,61,0,28,57,78,75,31,27,24,10,93,34,19,75,53,17,26,2,41,89,79,37,14,93,55,74,11,77,60,61,2,68,0,15,12,47,12,48,57,73,17,18,11,83,38,5,36,53,94,40,48,81,53,32,53,12,21,90,100,32,29,94,92,83,80,36,73,59,61,43,100,36,71,89,9,24,56,7,48,34,58,0,43,34,18,1,29,97,70,92,88,0,48,51,53,0,50,21,91,23,34,49,19,17,9,23,43,87,72,39,17,17,97,14,29,4,10,84,10,33,100,86,43,20,22,58,90,70,48,23,75,4,66,97,95,1,80,24,43,97,15,38,53,55,86,63,40,7,26,60,95,12,98,15,95,71,86,46,33,68,32,86,89,18,88,97,32,42,5,57,13,1,23,34,37,13,65,13,47,55,85,37,57,14,89,94,57,13,6,98,47,52,51,19,99,42,1,19,74,60,8,48,28,65,6,12,57,49,27,95,1,2,10,25,49,68,57,32,99,24,19,25,32,89,88,73,96,57,14,65,34,8,82,9,94,91,19,53,61,70,54,4,66,26,8,63,62,9,20,42,17,52,97,51,53,19,48,76,40,80,6,1,89,52,70,38,95,62,24,88,64,42,61,6,50,91,87,69,13,58,43,98,19,94,65,56,72,20,72,92,85,58,46,67,2,23,88,58,25,88,18,92,46,15,18,37,9,90,2,38,0,16,86,44,69,71,70,30,38,17,69,69,80,73,79,56,17,95,12,37,43,5,5,6,42,16,44,22,62,37,86,8,51,73,46,44,15,98,54,22,47,28,11,75,52,49,38,84,55,3,69,100,54,66,6,23,98,22,99,21,74,75,33,67,8,80,90,23,46,93,69,85,46,87,76,93,38,77,37,72,35,3,82,11,67,46,53,29,60,33,12,62,23,27,72,35,63,68,14,35,27,98,94,65,3,13,48,83,27,84,86,49,31,63,40,12,34,79,61,47,29,33,52,100,85,38,24,1,16,62,89,36,74,9,49,62,89}
	//	res := maxProfit3(input)
	//	fmt.Println(res)
	//}
	//{
	//	//input := []byte{'A','A','A','A','A','A','B','C','D','E','F','G'}
	//	//intervel := 2
	//	//res := leastInterval(input,intervel)
	//	//fmt.Println(res)
	//}
	//{
	//	input := []int{0,0,0,0,0,0,0,0,0,0}
	//	target := 0
	//	res := subarraySum(input,target)
	//	fmt.Println(res)
	//}
	//{
	//	input := 3
	//	res := numberOfSteps(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{2,3,-2,4,-2,4,6,-9,3}
	//	res := maxProduct(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := "3[a]2[bc]"
	//	//res := decodeString(input)
	//	res := decodeString2(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := [][]int{{1,1,0,0,0},{1,1,1,1,0},{1,0,0,0,0},{1,1,0,0,0},{1,1,1,1,1}}
	//	k := 3
	//	res := kWeakestRows(input,k)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{1,1,1,1,1}
	//	target := 3
	//	res := findTargetSumWays(input,target)
	//	fmt.Println(res)
	//}
	//{
	//	//    1
	//	//   / \
	//	//  2   5
	//	// / \   \
	//	//3   4   6
	//	var t1 TreeNode
	//	t1.Val = 1
	//	var t2 TreeNode
	//	t2.Val = 2
	//	var t3 TreeNode
	//	t3.Val = 3
	//	var t4 TreeNode
	//	t4.Val = 4
	//	var t5 TreeNode
	//	t5.Val = 5
	//	var t6 TreeNode
	//	t6.Val = 6
	//	t1.Left = &t2
	//	t1.Right = &t5
	//	t2.Left = &t3
	//	t2.Right = &t4
	//	t5.Right = &t6
	//	flatten(&t1)
	//	fmt.Println(t1)
	//}
	//{
	//	inorder := []int{9,3,15,20,7}
	//	postorder := []int{9,15,7,20,3}
	//	res := buildTree(inorder,postorder)
	//	fmt.Println(res)
	//}
	//{
	//	//[-10,-3,0,5,9]
	//	var l1 ListNode
	//	l1.Val = -10
	//	var l2 ListNode
	//	l2.Val = -3
	//	var l3 ListNode
	//	l3.Val = 0
	//	var l4 ListNode
	//	l4.Val = 5
	//	var l5 ListNode
	//	l5.Val = 9
	//	l1.Next = &l2
	//	l2.Next = &l3
	//	l3.Next = &l4
	//	l4.Next = &l5
	//	res := sortedListToBST(&l1)
	//	fmt.Println(res)
	//}
	//{
	//	input := []byte{'a','a','b','b','c','c','c'}
	//	res := compress(input)
	//	fmt.Println(res)
	//	fmt.Println(input)
	//}
	//{
	//	input := [][]int{{0,0},{1,0},{2,0}}
	//	res := numberOfBoomerangs(input)
	//	fmt.Println(res)
	//}
	//{
	//	secret := "1123"
	//	guess := "0111"
	//	res := getHint(secret,guess)
	//	fmt.Println(res)
	//}
	//{
	//	//res := countAndSay(90)
	//	//fmt.Println(res)
	//}
	//{
	//	input := []int{37,12,28,56,12}
	//	res := arrayRankTransform(input)
	//	fmt.Println(res)
	//}
	//{
	//	//[186,419,83,408]
	//	//6249
	//	coins := []int{186,419,83,408}
	//	target := 6249
	//	//coins :=  []int{3,7,405,436}
	//	//target := 8839
	//	res := coinChange(coins,target)
	//	//res = coinChange2(coins,target)
	//	fmt.Println(res)
	//}
	//{
	//	res := getRow(5)
	//	fmt.Println(res)
	//}
	//{
	//	A := "aa"
	//	B := "aaaaa"
	//	res := repeatedStringMatch(A,B)
	//	fmt.Println(res)
	//}
	//{
	//	res := countBinarySubstrings("00110011")
	//	fmt.Println(res)
	//}
	//{
	//	var input [][]int = [][]int{{2,1,1},{1,1,0},{0,1,1}}
	//	res := orangesRotting(input)
	//	fmt.Println(res)
	//}
	//{
	//	res := tribonacci(25)
	//	fmt.Println(res)
	//}
	//{
	//	res := maximum69Number(996)
	//	fmt.Println(res)
	//}
	//{
	//	beginWord := "nanny"
	//	endWord :=  "aloud"
	//	wordList := []string{"ricky","grind","cubic","panic","lover","farce","gofer","sales","flint","omens","lipid","briny","cloth","anted","slime","oaten","harsh","touts","stoop","cabal","lazed","elton","skunk","nicer","pesky","kusch","bused","kinda","tunis","enjoy","aches","prowl","babar","rooms","burst","slush","pines","urine","pinky","bayed","mania","light","flare","wares","women","verne","moron","shine","bluer","zeros","bleak","brief","tamra","vasts","jamie","lairs","penal","worst","yowls","pills","taros","addle","alyce","creep","saber","floyd","cures","soggy","vexed","vilma","cabby","verde","euler","cling","wanna","jenny","donor","stole","sakha","blake","sanes","riffs","forge","horus","sered","piked","prosy","wases","glove","onset","spake","benin","talks","sites","biers","wendy","dante","allan","haven","nears","shaka","sloth","perky","spear","spend","clint","dears","sadly","units","vista","hinds","marat","natal","least","bough","pales","boole","ditch","greys","slunk","bitch","belts","sense","skits","monty","yawns","music","hails","alien","gibes","lille","spacy","argot","wasps","drubs","poops","bella","clone","beast","emend","iring","start","darla","bells","cults","dhaka","sniff","seers","bantu","pages","fever","tacky","hoses","strop","climb","pairs","later","grant","raven","stael","drips","lucid","awing","dines","balms","della","galen","toned","snips","shady","chili","fears","nurse","joint","plump","micky","lions","jamal","queer","ruins","frats","spoof","semen","pulps","oldie","coors","rhone","papal","seals","spans","scaly","sieve","klaus","drums","tided","needs","rider","lures","treks","hares","liner","hokey","boots","primp","laval","limes","putts","fonda","damon","pikes","hobbs","specs","greet","ketch","braid","purer","tsars","berne","tarts","clean","grate","trips","chefs","timex","vicky","pares","price","every","beret","vices","jodie","fanny","mails","built","bossy","farms","pubic","gongs","magma","quads","shell","jocks","woods","waded","parka","jells","worse","diner","risks","bliss","bryan","terse","crier","incur","murky","gamed","edges","keens","bread","raced","vetch","glint","zions","porno","sizes","mends","ached","allie","bands","plank","forth","fuels","rhyme","wimpy","peels","foggy","wings","frill","edgar","slave","lotus","point","hints","germs","clung","limed","loafs","realm","myron","loopy","plush","volts","bimbo","smash","windy","sours","choke","karin","boast","whirr","tiber","dimes","basel","cutes","pinto","troll","thumb","decor","craft","tared","split","josue","tramp","screw","label","lenny","apses","slept","sikhs","child","bouts","cites","swipe","lurks","seeds","fists","hoard","steed","reams","spoil","diego","peale","bevel","flags","mazes","quart","snipe","latch","lards","acted","falls","busby","holed","mummy","wrong","wipes","carlo","leers","wails","night","pasty","eater","flunk","vedas","curse","tyros","mirth","jacky","butte","wired","fixes","tares","vague","roved","stove","swoon","scour","coked","marge","cants","comic","corns","zilch","typos","lives","truer","comma","gaily","teals","witty","hyper","croat","sways","tills","hones","dowel","llano","clefs","fores","cinch","brock","vichy","bleed","nuder","hoyle","slams","macro","arabs","tauts","eager","croak","scoop","crime","lurch","weals","fates","clipt","teens","bulls","domed","ghana","culls","frame","hanky","jared","swain","truss","drank","lobby","lumps","pansy","whews","saris","trite","weeps","dozes","jeans","flood","chimu","foxes","gelds","sects","scoff","poses","mares","famed","peers","hells","laked","zests","wring","steal","snoot","yodel","scamp","ellis","bandy","marry","jives","vises","blurb","relay","patch","haley","cubit","heine","place","touch","grain","gerry","badly","hooke","fuchs","savor","apron","judge","loren","britt","smith","tammy","altar","duels","huber","baton","dived","apace","sedan","basts","clark","mired","perch","hulks","jolly","welts","quack","spore","alums","shave","singe","lanny","dread","profs","skeet","flout","darin","newed","steer","taine","salvo","mites","rules","crash","thorn","olive","saves","yawed","pique","salon","ovens","dusty","janie","elise","carve","winds","abash","cheep","strap","fared","discs","poxed","hoots","catch","combo","maize","repay","mario","snuff","delve","cored","bards","sudan","shuns","yukon","jowls","wayne","torus","gales","creek","prove","needy","wisps","terri","ranks","books","dicky","tapes","aping","padre","roads","nines","seats","flats","rains","moira","basic","loves","pulls","tough","gills","codes","chest","teeny","jolts","woody","flame","asked","dulls","hotly","glare","mucky","spite","flake","vines","lindy","butts","froth","beeps","sills","bunny","flied","shaun","mawed","velds","voled","doily","patel","snake","thigh","adler","calks","desks","janus","spunk","baled","match","strip","hosed","nippy","wrest","whams","calfs","sleet","wives","boars","chain","table","duked","riped","edens","galas","huffs","biddy","claps","aleut","yucks","bangs","quids","glenn","evert","drunk","lusts","senna","slate","manet","roted","sleep","loxes","fluky","fence","clamp","doted","broad","sager","spark","belch","mandy","deana","beyer","hoist","leafy","levee","libel","tonic","aloes","steam","skews","tides","stall","rifts","saxon","mavis","asama","might","dotes","tangs","wroth","kited","salad","liens","clink","glows","balky","taffy","sided","sworn","oasis","tenth","blurt","tower","often","walsh","sonny","andes","slump","scans","boded","chive","finer","ponce","prune","sloes","dined","chums","dingo","harte","ahead","event","freer","heart","fetch","sated","soapy","skins","royal","cuter","loire","minot","aisle","horny","slued","panel","eight","snoop","pries","clive","pored","wrist","piped","daren","cells","parks","slugs","cubed","highs","booze","weary","stain","hoped","finny","weeds","fetid","racer","tasks","right","saint","shahs","basis","refer","chart","seize","lulls","slant","belay","clots","jinny","tours","modes","gloat","dunks","flute","conch","marts","aglow","gayer","lazes","dicks","chime","bears","sharp","hatch","forms","terry","gouda","thins","janet","tonya","axons","sewed","danny","rowdy","dolts","hurry","opine","fifty","noisy","spiky","humid","verna","poles","jayne","pecos","hooky","haney","shams","snots","sally","ruder","tempe","plunk","shaft","scows","essie","dated","fleet","spate","bunin","hikes","sodas","filly","thyme","fiefs","perks","chary","kiths","lidia","lefty","wolff","withe","three","crawl","wotan","brown","japed","tolls","taken","threw","crave","clash","layer","tends","notes","fudge","musky","bawdy","aline","matts","shirr","balks","stash","wicks","crepe","foods","fares","rotes","party","petty","press","dolly","mangy","leeks","silly","leant","nooks","chapt","loose","caged","wages","grist","alert","sheri","moody","tamps","hefts","souls","rubes","rolex","skulk","veeps","nonce","state","level","whirl","bight","grits","reset","faked","spiny","mixes","hunks","major","missy","arius","damns","fitly","caped","mucus","trace","surat","lloyd","furry","colin","texts","livia","reply","twill","ships","peons","shear","norms","jumbo","bring","masks","zippy","brine","dorks","roded","sinks","river","wolfs","strew","myths","pulpy","prank","veins","flues","minus","phone","banns","spell","burro","brags","boyle","lambs","sides","knees","clews","aired","skirt","heavy","dimer","bombs","scums","hayes","chaps","snugs","dusky","loxed","ellen","while","swank","track","minim","wiled","hazed","roofs","cantu","sorry","roach","loser","brass","stint","jerks","dirks","emory","campy","poise","sexed","gamer","catty","comte","bilbo","fasts","ledge","drier","idles","doors","waged","rizal","pured","weirs","crisp","tasty","sored","palmy","parts","ethel","unify","crows","crest","udder","delis","punks","dowse","totes","emile","coded","shops","poppa","pours","gushy","tiffs","shads","birds","coils","areas","boons","hulls","alter","lobes","pleat","depth","fires","pones","serra","sweat","kline","malay","ruled","calve","tired","drabs","tubed","wryer","slung","union","sonya","aided","hewed","dicey","grids","nixed","whits","mills","buffs","yucky","drops","ready","yuppy","tweet","napes","cadre","teach","rasps","dowdy","hoary","canto","posed","dumbo","kooks","reese","snaky","binge","byron","phony","safer","friar","novel","scale","huron","adorn","carla","fauna","myers","hobby","purse","flesh","smock","along","boils","pails","times","panza","lodge","clubs","colby","great","thing","peaks","diana","vance","whets","bergs","sling","spade","soaks","beach","traps","aspen","romps","boxed","fakir","weave","nerds","swazi","dotty","curls","diver","jonas","waite","verbs","yeast","lapel","barth","soars","hooks","taxed","slews","gouge","slags","chang","chafe","saved","josie","syncs","fonds","anion","actor","seems","pyrex","isiah","glued","groin","goren","waxes","tonia","whine","scads","knelt","teaks","satan","tromp","spats","merry","wordy","stake","gland","canal","donna","lends","filed","sacks","shied","moors","paths","older","pooch","balsa","riced","facet","decaf","attic","elder","akron","chomp","chump","picky","money","sheer","bolls","crabs","dorms","water","veers","tease","dummy","dumbs","lethe","halls","rifer","demon","fucks","whips","plops","fuses","focal","taces","snout","edict","flush","burps","dawes","lorry","spews","sprat","click","deann","sited","aunts","quips","godly","pupil","nanny","funks","shoon","aimed","stacy","helms","mints","banks","pinch","local","twine","pacts","deers","halos","slink","preys","potty","ruffs","pusan","suits","finks","slash","prods","dense","edsel","heeds","palls","slats","snits","mower","rares","ailed","rouge","ellie","gated","lyons","duded","links","oaths","letha","kicks","firms","gravy","month","kongo","mused","ducal","toted","vocal","disks","spied","studs","macao","erick","coupe","starr","reaps","decoy","rayon","nicks","breed","cosby","haunt","typed","plain","trays","muled","saith","drano","cower","snows","buses","jewry","argus","doers","flays","swish","resin","boobs","sicks","spies","bails","wowed","mabel","check","vapid","bacon","wilda","ollie","loony","irked","fraud","doles","facts","lists","gazed","furls","sunks","stows","wilde","brick","bowed","guise","suing","gates","niter","heros","hyped","clomp","never","lolls","rangy","paddy","chant","casts","terns","tunas","poker","scary","maims","saran","devon","tripe","lingo","paler","coped","bride","voted","dodge","gross","curds","sames","those","tithe","steep","flaks","close","swops","stare","notch","prays","roles","crush","feuds","nudge","baned","brake","plans","weepy","dazed","jenna","weiss","tomes","stews","whist","gibed","death","clank","cover","peeks","quick","abler","daddy","calls","scald","lilia","flask","cheer","grabs","megan","canes","jules","blots","mossy","begun","freak","caved","hello","hades","theed","wards","darcy","malta","peter","whorl","break","downs","odder","hoofs","kiddo","macho","fords","liked","flees","swing","elect","hoods","pluck","brook","astir","bland","sward","modal","flown","ahmad","waled","craps","cools","roods","hided","plath","kings","grips","gives","gnats","tabby","gauls","think","bully","fogey","sawed","lints","pushy","banes","drake","trail","moral","daley","balds","chugs","geeky","darts","soddy","haves","opens","rends","buggy","moles","freud","gored","shock","angus","puree","raves","johns","armed","packs","minis","reich","slots","totem","clown","popes","brute","hedge","latin","stoke","blend","pease","rubik","greer","hindi","betsy","flows","funky","kelli","humps","chewy","welds","scowl","yells","cough","sasha","sheaf","jokes","coast","words","irate","hales","camry","spits","burma","rhine","bends","spill","stubs","power","voles","learn","knoll","style","twila","drove","dacca","sheen","papas","shale","jones","duped","tunny","mouse","floss","corks","skims","swaps","inned","boxer","synch","skies","strep","bucks","belau","lower","flaky","quill","aural","rufus","floes","pokes","sends","sates","dally","boyer","hurts","foyer","gowns","torch","luria","fangs","moats","heinz","bolts","filet","firth","begot","argue","youth","chimp","frogs","kraft","smite","loges","loons","spine","domes","pokey","timur","noddy","doggy","wades","lanes","hence","louts","turks","lurid","goths","moist","bated","giles","stood","winos","shins","potts","brant","vised","alice","rosie","dents","babes","softy","decay","meats","tanya","rusks","pasts","karat","nuked","gorge","kinks","skull","noyce","aimee","watch","cleat","stuck","china","testy","doses","safes","stage","bayes","twins","limps","denis","chars","flaps","paces","abase","grays","deans","maria","asset","smuts","serbs","whigs","vases","robyn","girls","pents","alike","nodal","molly","swigs","swill","slums","rajah","bleep","beget","thanh","finns","clock","wafts","wafer","spicy","sorer","reach","beats","baker","crown","drugs","daisy","mocks","scots","fests","newer","agate","drift","marta","chino","flirt","homed","bribe","scram","bulks","servo","vesta","divas","preps","naval","tally","shove","ragas","blown","droll","tryst","lucky","leech","lines","sires","pyxed","taper","trump","payee","midge","paris","bored","loads","shuts","lived","swath","snare","boned","scars","aeons","grime","writs","paige","rungs","blent","signs","davis","dials","daubs","rainy","fawns","wrier","golds","wrath","ducks","allow","hosea","spike","meals","haber","muses","timed","broom","burks","louis","gangs","pools","vales","altai","elope","plied","slain","chasm","entry","slide","bawls","title","sings","grief","viola","doyle","peach","davit","bench","devil","latex","miles","pasha","tokes","coves","wheel","tried","verdi","wanda","sivan","prior","fryer","plots","kicky","porch","shill","coats","borne","brink","pawed","erwin","tense","stirs","wends","waxen","carts","smear","rival","scare","phase","bragg","crane","hocks","conan","bests","dares","molls","roots","dunes","slips","waked","fours","bolds","slosh","yemen","poole","solid","ports","fades","legal","cedes","green","curie","seedy","riper","poled","glade","hosts","tools","razes","tarry","muddy","shims","sword","thine","lasts","bloat","soled","tardy","foots","skiff","volta","murks","croci","gooks","gamey","pyxes","poems","kayla","larva","slaps","abuse","pings","plows","geese","minks","derby","super","inked","manic","leaks","flops","lajos","fuzes","swabs","twigs","gummy","pyres","shrew","islet","doled","wooly","lefts","hunts","toast","faith","macaw","sonia","leafs","colas","conks","altos","wiped","scene","boors","patsy","meany","chung","wakes","clear","ropes","tahoe","zones","crate","tombs","nouns","garth","puked","chats","hanks","baked","binds","fully","soaps","newel","yarns","puers","carps","spelt","lully","towed","scabs","prime","blest","patty","silky","abner","temps","lakes","tests","alias","mines","chips","funds","caret","splat","perry","turds","junks","cramp","saned","peary","snarl","fired","stung","nancy","bulge","styli","seams","hived","feast","triad","jaded","elvin","canny","birth","routs","rimed","pusey","laces","taste","basie","malls","shout","prier","prone","finis","claus","loops","heron","frump","spare","menus","ariel","crams","bloom","foxed","moons","mince","mixed","piers","deres","tempt","dryer","atone","heats","dario","hawed","swims","sheet","tasha","dings","clare","aging","daffy","wried","foals","lunar","havel","irony","ronny","naves","selma","gurus","crust","percy","murat","mauro","cowed","clang","biker","harms","barry","thump","crude","ulnae","thong","pager","oases","mered","locke","merle","soave","petal","poser","store","winch","wedge","inlet","nerdy","utter","filth","spray","drape","pukes","ewers","kinds","dates","meier","tammi","spoor","curly","chill","loped","gooey","boles","genet","boost","beets","heath","feeds","growl","livid","midst","rinds","fresh","waxed","yearn","keeps","rimes","naked","flick","plies","deeps","dirty","hefty","messy","hairy","walks","leper","sykes","nerve","rover","jived","brisk","lenin","viper","chuck","sinus","luger","ricks","hying","rusty","kathy","herds","wider","getty","roman","sandy","pends","fezes","trios","bites","pants","bless","diced","earth","shack","hinge","melds","jonah","chose","liver","salts","ratty","ashed","wacky","yokes","wanly","bruce","vowel","black","grail","lungs","arise","gluts","gluey","navel","coyer","ramps","miter","aldan","booth","musty","rills","darns","tined","straw","kerri","hared","lucks","metes","penny","radon","palms","deeds","earls","shard","pried","tampa","blank","gybes","vicki","drool","groom","curer","cubes","riggs","lanky","tuber","caves","acing","golly","hodge","beard","ginny","jibed","fumes","astor","quito","cargo","randi","gawky","zings","blind","dhoti","sneak","fatah","fixer","lapps","cline","grimm","fakes","maine","erika","dealt","mitch","olden","joist","gents","likes","shelf","silts","goats","leads","marin","spire","louie","evans","amuse","belly","nails","snead","model","whats","shari","quote","tacks","nutty","lames","caste","hexes","cooks","miner","shawn","anise","drama","trike","prate","ayers","loans","botch","vests","cilia","ridge","thugs","outed","jails","moped","plead","tunes","nosed","wills","lager","lacks","cried","wince","berle","flaws","boise","tibet","bided","shred","cocky","brice","delta","congo","holly","hicks","wraps","cocks","aisha","heard","cured","sades","horsy","umped","trice","dorky","curve","ferry","haler","ninth","pasta","jason","honer","kevin","males","fowls","awake","pores","meter","skate","drink","pussy","soups","bases","noyes","torts","bogus","still","soupy","dance","worry","eldon","stern","menes","dolls","dumpy","gaunt","grove","coops","mules","berry","sower","roams","brawl","greed","stags","blurs","swift","treed","taney","shame","easel","moves","leger","ville","order","spock","nifty","brian","elias","idler","serve","ashen","bizet","gilts","spook","eaten","pumas","cotes","broke","toxin","groan","laths","joins","spots","hated","tokay","elite","rawer","fiats","cards","sassy","milks","roost","glean","lutes","chins","drown","marks","pined","grace","fifth","lodes","rusts","terms","maxes","savvy","choir","savoy","spoon","halve","chord","hulas","sarah","celia","deems","ninny","wines","boggy","birch","raved","wales","beams","vibes","riots","warty","nigel","askew","faxes","sedge","sheol","pucks","cynic","relax","boers","whims","bents","candy","luann","slogs","bonny","barns","iambs","fused","duffy","guilt","bruin","pawls","penis","poppy","owing","tribe","tuner","moray","timid","ceded","geeks","kites","curio","puffy","perot","caddy","peeve","cause","dills","gavel","manse","joker","lynch","crank","golda","waits","wises","hasty","paves","grown","reedy","crypt","tonne","jerky","axing","swept","posse","rings","staff","tansy","pared","glaze","grebe","gonna","shark","jumps","vials","unset","hires","tying","lured","motes","linen","locks","mamas","nasty","mamie","clout","nader","velma","abate","tight","dales","serer","rives","bales","loamy","warps","plato","hooch","togae","damps","ofter","plumb","fifes","filmy","wiper","chess","lousy","sails","brahe","ounce","flits","hindu","manly","beaux","mimed","liken","forts","jambs","peeps","lelia","brews","handy","lusty","brads","marne","pesos","earle","arson","scout","showy","chile","sumps","hiked","crook","herbs","silks","alamo","mores","dunce","blaze","stank","haste","howls","trots","creon","lisle","pause","hates","mulch","mined","moder","devin","types","cindy","beech","tuned","mowed","pitts","chaos","colds","bidet","tines","sighs","slimy","brain","belle","leery","morse","ruben","prows","frown","disco","regal","oaken","sheds","hives","corny","baser","fated","throe","revel","bores","waved","shits","elvia","ferns","maids","color","coifs","cohan","draft","hmong","alton","stine","cluck","nodes","emily","brave","blair","blued","dress","bunts","holst","clogs","rally","knack","demos","brady","blues","flash","goofy","blocs","diane","colic","smile","yules","foamy","splay","bilge","faker","foils","condo","knell","crack","gallo","purls","auras","cakes","doves","joust","aides","lades","muggy","tanks","middy","tarps","slack","capet","frays","donny","venal","yeats","misty","denim","glass","nudes","seeps","gibbs","blows","bobbi","shane","yards","pimps","clued","quiet","witch","boxes","prawn","kerry","torah","kinko","dingy","emote","honor","jelly","grins","trope","vined","bagel","arden","rapid","paged","loved","agape","mural","budge","ticks","suers","wendi","slice","salve","robin","bleat","batik","myles","teddy","flatt","puppy","gelid","largo","attar","polls","glide","serum","fundy","sucks","shalt","sewer","wreak","dames","fonts","toxic","hines","wormy","grass","louse","bowls","crass","benny","moire","margo","golfs","smart","roxie","wight","reign","dairy","clops","paled","oddly","sappy","flair","shown","bulgy","benet","larch","curry","gulfs","fends","lunch","dukes","doris","spoke","coins","manna","conga","jinns","eases","dunno","tisha","swore","rhino","calms","irvin","clans","gully","liege","mains","besot","serge","being","welch","wombs","draco","lynda","forty","mumps","bloch","ogden","knits","fussy","alder","danes","loyal","valet","wooer","quire","liefs","shana","toyed","forks","gages","slims","cloys","yates","rails","sheep","nacho","divan","honks","stone","snack","added","basal","hasps","focus","alone","laxes","arose","lamed","wrapt","frail","clams","plait","hover","tacos","mooch","fault","teeth","marva","mucks","tread","waves","purim","boron","horde","smack","bongo","monte","swirl","deals","mikes","scold","muter","sties","lawns","fluke","jilts","meuse","fives","sulky","molds","snore","timmy","ditty","gasps","kills","carey","jawed","byers","tommy","homer","hexed","dumas","given","mewls","smelt","weird","speck","merck","keats","draws","trent","agave","wells","chews","blabs","roves","grieg","evens","alive","mulls","cared","garbo","fined","happy","trued","rodes","thurs","cadet","alvin","busch","moths","guild","staci","lever","widen","props","hussy","lamer","riley","bauer","chirp","rants","poxes","shyer","pelts","funny","slits","tinge","ramos","shift","caper","credo","renal","veils","covey","elmer","mated","tykes","wooed","briar","gears","foley","shoes","decry","hypes","dells","wilds","runts","wilts","white","easts","comer","sammy","lochs","favor","lance","dawns","bushy","muted","elsie","creel","pocks","tenet","cagey","rides","socks","ogled","soils","sofas","janna","exile","barks","frank","takes","zooms","hakes","sagan","scull","heaps","augur","pouch","blare","bulbs","wryly","homey","tubas","limbo","hardy","hoagy","minds","bared","gabby","bilks","float","limns","clasp","laura","range","brush","tummy","kilts","cooed","worms","leary","feats","robes","suite","veals","bosch","moans","dozen","rarer","slyer","cabin","craze","sweet","talon","treat","yanks","react","creed","eliza","sluts","cruet","hafts","noise","seder","flies","weeks","venus","backs","eider","uriel","vouch","robed","hacks","perth","shiny","stilt","torte","throb","merer","twits","reeds","shawl","clara","slurs","mixer","newts","fried","woolf","swoop","kaaba","oozed","mayer","caned","laius","lunge","chits","kenny","lifts","mafia","sowed","piled","stein","whack","colts","warms","cleft","girds","seeks","poets","angel","trade","parsi","tiers","rojas","vexes","bryce","moots","grunt","drain","lumpy","stabs","poohs","leapt","polly","cuffs","giddy","towns","dacha","quoth","provo","dilly","carly","mewed","tzars","crock","toked","speak","mayas","pssts","ocher","motel","vogue","camps","tharp","taunt","drone","taint","badge","scott","scats","bakes","antes","gruel","snort","capes","plate","folly","adobe","yours","papaw","hench","moods","clunk","chevy","tomas","narcs","vonda","wiles","prigs","chock","laser","viced","stiff","rouse","helps","knead","gazer","blade","tumid","avail","anger","egged","guide","goads","rabin","toddy","gulps","flank","brats","pedal","junky","marco","tinny","tires","flier","satin","darth","paley","gumbo","rared","muffs","rower","prude","frees","quays","homes","munch","beefs","leash","aston","colon","finch","bogey","leaps","tempo","posts","lined","gapes","locus","maori","nixes","liven","songs","opted","babel","wader","barer","farts","lisps","koran","lathe","trill","smirk","mamma","viler","scurf","ravel","brigs","cooky","sachs","fulls","goals","turfs","norse","hauls","cores","fairy","pluto","kneed","cheek","pangs","risen","czars","milne","cribs","genes","wefts","vents","sages","seres","owens","wiley","flume","haded","auger","tatty","onion","cater","wolfe","magic","bodes","gulls","gazes","dandy","snags","rowed","quell","spurn","shore","veldt","turns","slavs","coach","stalk","snuck","piles","orate","joyed","daily","crone","wager","solos","earns","stark","lauds","kasey","villa","gnaws","scent","wears","fains","laced","tamer","pipes","plant","lorie","rivet","tamed","cozen","theme","lifer","sunny","shags","flack","gassy","eased","jeeps","shire","fargo","timer","brash","behan","basin","volga","krone","swiss","docks","booed","ebert","gusty","delay","oared","grady","buick","curbs","crete","lucas","strum","besom","gorse","troth","donne","chink","faced","ahmed","texas","longs","aloud","bethe","cacao","hilda","eagle","karyn","harks","adder","verse","drays","cello","taped","snide","taxis","kinky","penes","wicca","sonja","aways","dyers","bolas","elfin","slope","lamps","hutch","lobed","baaed","masts","ashes","ionic","joyce","payed","brays","malts","dregs","leaky","runny","fecal","woven","hurls","jorge","henna","dolby","booty","brett","dykes","rural","fight","feels","flogs","brunt","preen","elvis","dopey","gripe","garry","gamma","fling","space","mange","storm","arron","hairs","rogue","repel","elgar","ruddy","cross","medan","loses","howdy","foams","piker","halts","jewel","avery","stool","cruel","cases","ruses","cathy","harem","flour","meted","faces","hobos","charm","jamar","cameo","crape","hooey","reefs","denny","mitts","sores","smoky","nopes","sooty","twirl","toads","vader","julep","licks","arias","wrote","north","bunks","heady","batch","snaps","claws","fouls","faded","beans","wimps","idled","pulse","goons","noose","vowed","ronda","rajas","roast","allah","punic","slows","hours","metal","slier","meaty","hanna","curvy","mussy","truth","troys","block","reels","print","miffs","busts","bytes","cream","otter","grads","siren","kilos","dross","batty","debts","sully","bares","baggy","hippy","berth","gorky","argon","wacko","harry","smoke","fails","perms","score","steps","unity","couch","kelly","rumps","fines","mouth","broth","knows","becky","quits","lauri","trust","grows","logos","apter","burrs","zincs","buyer","bayer","moose","overt","croon","ousts","lands","lithe","poach","jamel","waive","wiser","surly","works","paine","medal","glads","gybed","paint","lorre","meant","smugs","bryon","jinni","sever","viols","flubs","melts","heads","peals","aiken","named","teary","yalta","styes","heist","bongs","slops","pouts","grape","belie","cloak","rocks","scone","lydia","goofs","rents","drive","crony","orlon","narks","plays","blips","pence","march","alger","baste","acorn","billy","croce","boone","aaron","slobs","idyls","irwin","elves","stoat","doing","globe","verve","icons","trial","olsen","pecks","there","blame","tilde","milky","sells","tangy","wrack","fills","lofty","truce","quark","delia","stowe","marty","overs","putty","coral","swine","stats","swags","weans","spout","bulky","farsi","brest","gleam","beaks","coons","hater","peony","huffy","exert","clips","riven","payer","doped","salas","meyer","dryad","thuds","tilts","quilt","jetty","brood","gulch","corps","tunic","hubby","slang","wreck","purrs","punch","drags","chide","sulks","tints","huger","roped","dopes","booby","rosin","outer","gusto","tents","elude","brows","lease","ceres","laxer","worth","necks","races","corey","trait","stuns","soles","teems","scrip","privy","sight","minor","alisa","stray","spank","cress","nukes","rises","gusts","aurae","karma","icing","prose","biked","grand","grasp","skein","shaky","clump","rummy","stock","twain","zoned","offed","ghats","mover","randy","vault","craws","thees","salem","downy","sangs","chore","cited","grave","spinx","erica","raspy","dying","skips","clerk","paste","moved","rooks","intel","moses","avers","staid","yawls","blast","lyres","monks","gaits","floor","saner","waver","assam","infer","wands","bunch","dryly","weedy","honey","baths","leach","shorn","shows","dream","value","dooms","spiro","raped","shook","stead","moran","ditto","loots","tapir","looms","clove","stops","pinks","soppy","ripen","wench","shone","bauds","doric","leans","nadia","cries","camus","boozy","maris","fools","morns","bides","greek","gauss","roget","lamar","hazes","beefy","dupes","refed","felts","larry","guile","ables","wants","warns","toils","bathe","edger","paced","rinks","shoos","erich","whore","tiger","jumpy","lamas","stack","among","punts","scalp","alloy","solon","quite","comas","whole","parse","tries","reeve","tiled","deena","roomy","rodin","aster","twice","musts","globs","parch","drawn","filch","bonds","tells","droop","janis","holds","scant","lopes","based","keven","whiny","aspic","gains","franz","jerri","steel","rowel","vends","yelps","begin","logic","tress","sunni","going","barge","blood","burns","basks","waifs","bones","skill","hewer","burly","clime","eking","withs","capek","berta","cheap","films","scoot","tweed","sizer","wheat","acton","flung","ponds","tracy","fiver","berra","roger","mutes","burke","miked","valve","whisk","runes","parry","toots","japes","roars","rough","irons","romeo","cages","reeks","cigar","saiph","dully","hangs","chops","rolls","prick","acuff","spent","sulla","train","swell","frets","names","anita","crazy","sixth","blunt","fewer","large","brand","slick","spitz","rears","ogres","toffy","yolks","flock","gnawn","eries","blink","skier","feted","tones","snail","ether","barbs","noses","hears","upset","awash","cloud","trunk","degas","dungs","rated","shall","yeahs","coven","sands","susan","fable","gunny","began","serfs","balls","dinky","madge","prong","spilt","lilly","brawn","comet","spins","raids","dries","sorts","makes","mason","mayra","royce","stout","mealy","pagan","nasal","folds","libby","coups","photo","mosey","amens","speed","lords","board","fetal","lagos","scope","raked","bonus","mutts","willy","sport","bingo","thant","araby","bette","rebel","gases","small","humus","grosz","beset","slays","steve","scrap","blahs","south","pride","heels","tubes","beady","lacey","genus","mauls","vying","spice","sexes","ester","drams","today","comae","under","jests","direr","yoked","tempi","early","boats","jesus","warts","guppy","gilda","quota","token","edwin","ringo","gaped","lemon","hurst","manor","arrow","mists","prize","silas","blobs","diets","ervin","stony","buddy","bates","rabid","ducat","ewing","jaunt","beads","doyen","blush","thoth","tiles","piper","short","peron","alley","decks","shunt","whirs","cushy","roils","betty","plugs","woken","jibes","foray","merak","ruing","becks","whale","shoot","dwelt","spawn","fairs","dozed","celts","blond","tikes","sabin","feint","vamps","cokes","willa","slues","bills","force","curst","yokel","surer","miler","fices","arced","douse","hilly","lucio","tongs","togas","minty","sagas","pates","welsh","bruno","decal","elate","linux","gyros","pryor","mousy","pains","shake","spica","pupal","probe","mount","shirk","purus","kilns","rests","graze","hague","spuds","sweep","momma","burch","maces","samar","brace","riser","booms","build","camel","flyer","synge","sauna","tonga","tings","promo","hides","clair","elisa","bower","reins","diann","lubed","nulls","picks","laban","milch","buber","stomp","bosom","lying","haled","avert","wries","macon","skids","fumed","ogles","clods","antic","nosey","crimp","purge","mommy","cased","taxes","covet","clack","butch","panty","lents","machs","exude","tooth","adore","shuck","asses","after","terra","dices","aryan","regor","romes","stile","cairo","maura","flail","eaves","estes","sousa","visas","baron","civet","kitty","freed","ralph","tango","gawks","cheat","study","fancy","fiber","musks","souse","brims","claim","bikes","venue","sired","thymi","rivas","skimp","pleas","woman","gimpy","cawed","minos","pints","knock","poked","bowen","risky","towel","oinks","linus","heals","pears","codas","inner","pitch","harpy","niger","madly","bumpy","stair","files","nobel","celli","spars","jades","balmy","kooky","plums","trues","gloss","trims","daunt","tubby","dared","wadis","smell","darby","stink","drill","dover","ruler","laden","dikes","layla","fells","maker","joked","horns","these","baize","spahn","whens","edged","mushy","plume","tucks","spurs","husky","dried","bigot","pupas","drily","aware","hagar","newly","knots","pratt","feces","sabik","watts","cooke","riles","seamy","fleas","dusts","barfs","roans","pawns","vivid","kirks","tania","feral","tubae","horne","aries","brits","combs","chunk","stork","waned","texan","elide","glens","emery","autos","trams","dosed","cheri","baits","jacks","whose","fazed","matte","swans","maxed","write","spays","orion","traci","horse","stars","strut","goods","verge","scuff","award","dives","wires","burnt","dimly","sleds","mayan","biped","quirk","sofia","slabs","waste","robby","mayor","fatty","items","bowel","mires","swarm","route","swash","sooth","paved","steak","upend","sough","throw","perts","stave","carry","burgs","hilts","plane","toady","nadir","stick","foist","gnarl","spain","enter","sises","story","scarf","ryder","glums","nappy","sixes","honed","marcy","offer","kneel","leeds","lites","voter","vince","bursa","heave","roses","trees","argos","leann","grimy","zelma","crick","tract","flips","folks","smote","brier","moore","goose","baden","riled","looks","sober","tusks","house","acmes","lubes","chows","neath","vivas","defer","allay","casey","kmart","pests","proms","eying","cider","leave","shush","shots","karla","scorn","gifts","sneer","mercy","copes","faxed","spurt","monet","awoke","rocky","share","gores","drawl","tears","mooed","nones","wined","wrens","modem","beria","hovel","retch","mates","hands","stymy","peace","carat","coots","hotel","karen","hayed","mamet","cuing","paper","rages","suave","reuse","auden","costs","loner","rapes","hazel","rites","brent","pumps","dutch","puffs","noons","grams","teats","cease","honda","pricy","forgo","fleck","hired","silos","merge","rafts","halon","larks","deere","jello","cunts","sifts","boner","morin","mimes","bungs","marie","harts","snobs","sonic","hippo","comes","crops","mango","wrung","garbs","natty","cents","fitch","moldy","adams","sorta","coeds","gilds","kiddy","nervy","slurp","ramon","fuzed","hiker","winks","vanes","goody","hawks","crowd","bract","marla","limbs","solve","gloom","sloop","eaton","memos","tames","heirs","berms","wanes","faint","numbs","holes","grubs","rakes","waist","miser","stays","antis","marsh","skyed","payne","champ","jimmy","clues","fatal","shoed","freon","lopez","snowy","loins","stale","thank","reads","isles","grill","align","saxes","rubin","rigel","walls","beers","wispy","topic","alden","anton","ducts","david","duets","fries","oiled","waken","allot","swats","woozy","tuxes","inter","dunne","known","axles","graph","bumps","jerry","hitch","crews","lucia","banal","grope","valid","meres","thick","lofts","chaff","taker","glues","snubs","trawl","keels","liker","stand","harps","casks","nelly","debby","panes","dumps","norma","racks","scams","forte","dwell","dudes","hypos","sissy","swamp","faust","slake","maven","lowed","lilts","bobby","gorey","swear","nests","marci","palsy","siege","oozes","rates","stunt","herod","wilma","other","girts","conic","goner","peppy","class","sized","games","snell","newsy","amend","solis","duane","troop","linda","tails","woofs","scuds","shies","patti","stunk","acres","tevet","allen","carpi","meets","trend","salty","galls","crept","toner","panda","cohen","chase","james","bravo","styed","coals","oates","swami","staph","frisk","cares","cords","stems","razed","since","mopes","rices","junes","raged","liter","manes","rearm","naive","tyree","medic","laded","pearl","inset","graft","chair","votes","saver","cains","knobs","gamay","hunch","crags","olson","teams","surge","wests","boney","limos","ploys","algae","gaols","caked","molts","glops","tarot","wheal","cysts","husks","vaunt","beaus","fauns","jeers","mitty","stuff","shape","sears","buffy","maced","fazes","vegas","stamp","borer","gaged","shade","finds","frock","plods","skied","stump","ripes","chick","cones","fixed","coled","rodeo","basil","dazes","sting","surfs","mindy","creak","swung","cadge","franc","seven","sices","weest","unite","codex","trick","fusty","plaid","hills","truck","spiel","sleek","anons","pupae","chiba","hoops","trash","noted","boris","dough","shirt","cowls","seine","spool","miens","yummy","grade","proxy","hopes","girth","deter","dowry","aorta","paean","corms","giant","shank","where","means","years","vegan","derek","tales"}
	//	//beginWord := "qa"
	//	//endWord := "sq"
	//	//wordList := []string{"si","go","se","cm","so","ph","mt","db","mb","sb","kr","ln","tm","le","av","sm","ar","ci","ca","br","ti","ba","to","ra","fa","yo","ow","sn","ya","cr","po","fe","ho","ma","re","or","rn","au","ur","rh","sr","tc","lt","lo","as","fr","nb","yb","if","pb","ge","th","pm","rb","sh","co","ga","li","ha","hz","no","bi","di","hi","qa","pi","os","uh","wm","an","me","mo","na","la","st","er","sc","ne","mn","mi","am","ex","pt","io","be","fm","ta","tb","ni","mr","pa","he","lr","sq","ye"}
	//	//beginWord := "hot"
	//	//endWord := "dog"
	//	//wordList := []string{"hot","cog","dog","tot","hog","hop","pot","dot"}
	//	//beginWord := "hot"
	//	//endWord := "dog"
	//	//wordList := []string{"hot","dog","dot"}
	//	res := ladderLength(beginWord,endWord,wordList)
	//	fmt.Println(res)
	//}
	//{
	//	input := [][]byte{{'O','X','X','O','X'},{'X','O','O','X','O'},{'X','O','X','O','X'},{'O','X','O','O','O'},{'X','X','O','X','O'}}
	//	solve(input)
	//}
	//{
	//	res := multiply("111","222")
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{1,1,2,5,2,3,4}
	//	res := permuteUnique(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := [][]int{{1,4},{1,5}}
	//	res := merge(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{1,3,1,2,0,5}
	//	res := maxSlidingWindow(input,3)
	//	fmt.Println(res)
	//}
	//{
	//	board := [][]byte{{'A','B','C','E'}, {'S','F','C','S'}, {'A','D','E','E'}}
	//	word := "SEE"
	//	res := exist(board,word)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{8,2,4,4,4,9,5,2,5,8,8,0,8,6,9,1,1,6,3,5,1,2,6,6,0,4,8,6,0,3,2,8,7,6,5,1,7,0,3,4,8,3,5,9,0,4,0,1,0,5,9,2,0,7,0,2,1,0,8,2,5,1,2,3,9,7,4,7,0,0,1,8,5,6,7,5,1,9,9,3,5,0,7,5}
	//	res := jump(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{3,0,2,1,2}
	//	res := canReach(input,2)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{25000,24999,24998,6,5,4,3,2,1,1,0,0}
	//	res := canJump(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{9,1,4,7,3,-1,0,5,8,-1,6}
	//	res := longestConsecutive(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{10,9,2,5,3,4}
	//	res := lengthOfLIS(input)
	//	fmt.Println(res)
	//}
	//{
	//	preorder := []int{3,9,20,15,7}
	//	inorder := []int{9,3,15,20,7}
	//	res := buildTree2(preorder,inorder)
	//	fmt.Println(res)
	//}
	//{
	//	res := numSquares(12)
	//	fmt.Println(res)
	//}
	//{
	//	res := numTrees(4)
	//	fmt.Println(res)
	//}
	//{
	//	input := [][]byte{{'1','1','0','0','0'},{'1','1','0','0','0'},{'0','0','1','0','0'},{'0','0','0','1','1'}}
	//	res := numIslands(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := "23"
	//	res := letterCombinations(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := "aa"
	//	res := partition(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := [][]int{{1,3,1},{1,5,1},{4,2,1}}
	//	res := minPathSum(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{2,3,6,7}
	//	target := 7
	//	res := combinationSum(input,target)
	//	fmt.Println(res)
	//}
	//{
	//	//input := [][]int{{0,1,0}, {0,0,1}, {1,1,1}, {0,0,0}}
	//	input := [][]int{{0,0}}
	//	gameOfLife(input)
	//}
	//{
	//	input := [][]int{{5,1,9,11},
	//		{2,4,8,10},
	//		{13,3,6,7},
	//		{15,14,12,16}}
	//	rotate(input)
	//	fmt.Println(input)
	//}
	//{
	//	input := []int{1,3,4,2,2}
	//	res := findDuplicate(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := "MCMXCIV"
	//	res := romanToInt(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := "A man, a plan, a canal: Panama"
	//	res := isPalindrome(input)
	//	fmt.Println(res)
	//}
	//{
	//	ss := strings.TrimLeft("00010","0")
	//	fmt.Println(ss)
	//	res := largestNumber([]int{3,32,34,5,9})
	//	fmt.Println(res)
	//}
	//{
	//	res := longestSubstring("aaabb", 3)
	//	fmt.Println(res)
	//}
	//{
	//	res := subtractProductAndSum(234)
	//	fmt.Println(res)
	//}
	//{
	//	res := reverseParentheses("(u(love)i)")
	//	fmt.Println(res)
	//}
	//{
	//	res := numberOfArithmeticSlices([]int{1,2,3,4,8,9,10})
	//	fmt.Println(res)
	//}
	//{
	//	//input := [][]int{{-19,-1,-96,48,-94,36,16,55,-42,37,-59,6,-32,96,95,-58,13,-34,94,85},{17,44,36,-29,84,80,-34,50,-99,64,13,91,-27,25,-36,57,20,98,-100,-72},{-92,-75,86,90,-4,90,64,56,50,-63,10,-15,90,-66,-66,32,-69,-78,1,60},{21,51,-47,-43,-14,99,44,90,8,11,99,-62,57,59,69,50,-69,32,85,13},{-28,90,12,-18,23,61,-55,-97,6,89,36,26,26,-1,46,-50,79,-45,89,86},{-85,-10,49,-10,2,62,41,92,-67,85,86,27,89,-50,77,55,22,-82,-94,-98},{-50,53,-23,55,25,-22,76,-93,-7,66,-75,42,-35,-96,-5,4,-92,13,-31,-100},{-62,-78,8,-92,86,69,90,-37,81,97,53,-45,34,19,-19,-39,-88,-75,-74,-4},{29,53,-91,65,-92,11,49,26,90,-31,17,-84,12,63,-60,-48,40,-49,-48,88},{100,-69,80,11,-93,17,28,-94,52,64,-86,30,-9,-53,-8,-68,-33,31,-5,11},{9,64,-31,63,-84,-15,-30,-10,67,2,98,73,-77,-37,-96,47,-97,78,-62,-17},{-88,-38,-22,-90,54,42,-29,67,-85,-90,-29,81,52,35,13,61,-18,-94,61,-62},{-23,-29,-76,-30,-65,23,31,-98,-9,11,75,-1,-84,-90,73,58,72,-48,30,-81},{66,-33,91,-6,-94,82,25,-43,-93,-25,-69,10,-71,-65,85,28,-52,76,25,90},{-3,78,36,-92,-52,-44,-66,-53,-55,76,-7,76,-73,13,-98,86,-99,-22,61,100},{-97,65,2,-93,56,-78,22,56,35,-24,-95,-13,83,-34,-51,-73,2,7,-86,-19},{32,94,-14,-13,-6,-55,-21,29,-21,16,67,100,77,-26,-96,22,-5,-53,-92,-36},{60,93,-79,76,-91,43,-95,-16,74,-21,85,43,21,-68,-32,-18,18,100,-43,1},{87,-31,26,53,26,51,-61,92,-65,17,-41,27,-42,-14,37,-46,46,-31,-74,23},{-67,-14,-20,-85,42,36,56,9,11,-66,-59,-55,5,64,-29,77,47,44,-33,-77}}
	//	input := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	//	res := minFallingPathSum(input)
	//	fmt.Println(res)
	//}
	//{
	//	//input := [][]byte{{"1","0","1","0","0"},{"1","0","1","1","1"},{"1","1","1","1","1"},{"1","0","0","1","0"}}
	//}
	//{
	//	input := [][]int{{0,1,1,1}, {1,1,1,1}, {0,1,1,1}}
	//	res := countSquares(input)
	//	fmt.Println(res)
	//}
	//{
	//	//input := [][]int{{3,3},{5,-1},{-2,4}}
	//	//input := [][]int{{1,3},{-2,2},{2,-2}}
	//	input := [][]int{{8,6},{9,7},{2,8},{5,7},{10,1},{4,1},{2,7},{9,8},{-2,4},{3,3},{5,-1},{-2,4}}
	//	res := kClosest(input,4)
	//	fmt.Println(res)
	//}
	//{
	//	input := [][]int{{0,0,0,22,0,24},{34,23,18,0,23,2},{11,39,20,12,0,0},{39,8,0,2,0,1},{19,32,26,20,20,30},{0,38,26,0,29,31}}
	//	//input := [][]int{{1,0,7},{2,0,6},{3,4,5},{0,3,0},{9,0,20}}
	//	//input := [][]int{{0,6,0},{5,8,7},{0,9,0}}
	//	res := getMaximumGold(input)
	//	fmt.Println(res)
	//}
	//{
	//	//input := [][]int{{1},{2},{3},{}}
	//	//input := [][]int{{1,3,2},{2,3},{2,1,3,1},{}}
	//	input := [][]int{{1,3},{3,0,1},{2},{0}}
	//	res := canVisitAllRooms(input)
	//	fmt.Println(res)
	//}
	//{
	//	//gas := []int{1,2,3,4,5}
	//	//cost := []int{3,4,5,1,2}
	//	gas := []int{2,3,4}
	//	cost := []int{3,4,3}
	//	res := canCompleteCircuit(gas,cost)
	//	fmt.Println(res)
	//}
	//{
	//	//pre := []int{1,2,4,5,3,6,7}
	//	//post := []int{4,5,2,6,7,3,1}
	//	pre := []int{1,2,3}
	//	post := []int{2,3,1}
	//	res := constructFromPrePost(pre,post)
	//	fmt.Println(res)
	//}
	//{
	//	res := mctFromLeafValues([]int{6,2,4})
	//	fmt.Println(res)
	//}
	//{
	//	res := maxSumAfterPartitioning([]int{1,15,7,9,2,5,10},3)
	//	fmt.Println(res)
	//}
	//{
	//	res := customSortString("cba","abcd")
	//	fmt.Println(res)
	//}
	//{
	//	input := [][]int{{0,1},{1,1}}
	//	res := oddCells(2,3,input)
	//	fmt.Println(res)
	//}
	//{
	//	res := pathInZigZagTree(14)
	//	fmt.Println(res)
	//}
	//{
	//	input := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	//	res := subdomainVisits(input)
	//	fmt.Println(res)
	//}
	//{
	//	queens := [][]int{{5,6},{7,7},{2,1},{0,7},{1,6},{5,1},{3,7},{0,3},{4,0},{1,2},{6,3},{5,0},{0,4},{2,2},{1,1},{6,4},{5,4},{0,0},{2,6},{4,5},{5,2},{1,4},{7,5},{2,3},{0,5},{4,2},{1,0},{2,7},{0,1},{4,6},{6,1},{0,6},{4,3},{1,7}}
	//	king := []int{3,4}
	//	res := queensAttacktheKing(queens,king)
	//	fmt.Println(res)
	//}
	//{
	//	var input [][]int = [][]int{{0,0,1,1},{1,0,1,0},{1,1,0,0}}
	//	matrixScore(input)
	//}
	//{
	//	input := "RLRRLLRLRL"
	//	res := balancedStringSplit(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := "ababcbacadefegdehijhklij"
	//	res := partitionLabels(input)
	//	fmt.Println(res)
	//}
	//{
	//	s1 := "ab"
	//	s2 := "eidboaoo"
	//	res := checkInclusion(s1,s2)
	//	fmt.Println(res)
	//}
	//{
	//	houses := []int{1,2,3,4,9,7,6,98}
	//	heaters := []int{30,65}
	//	res := findRadius(houses,heaters)
	//	fmt.Println(res)
	//}
	//{
	//	arr1 := []int{2,3,1,3,2,4,6,7,9,2,7,19,19}
	//	arr2 := []int{2,1,4,3,9,6}
	//	res := relativeSortArray(arr1,arr2)
	//	fmt.Println(res)
	//}
	//{
	//	var input []int = []int{1,1,1,1,2,2,2,2,2,2}
	//	res := hasGroupsSizeX(input)
	//	fmt.Println(res)
	//}
	//{
	//	var input []byte = []byte{'c', 'f', 'j'}
	//	res := nextGreatestLetter(input,'c')
	//	fmt.Println(res)
	//}
	//{
	//	var input []int = []int{30,20,150,100,40}
	//	res := numPairsDivisibleBy60(input)
	//	fmt.Println(res)
	//}
	//{
	//	nums := []int{3,3,2,6,7,8}
	//	largest := make([]int, 4, 4)
	//	copy(largest,nums[0:4])
	//	fmt.Println(largest)
	//}
	//{
	//	res := countBinarySubstrings(  "00110011")
	//	fmt.Println(res)
	//}
	//{
	//	var input [][]int = [][]int{{1,3},{2,3},{3,1}}
	//	res := findJudge(3,input)
	//	fmt.Println(res)
	//}
	//{
	//	var input []string = []string{"Hello","Alaska", "Dad", "Peace"}
	//	ret := findWords(input)
	//	fmt.Println(ret)
	//}
	//{
	//	floor := []int{2,3,4,5,18,17,6}
	//	var res int = maxArea(floor)
	//	fmt.Println(res)
	//}
	//{
	//	generate(5)
	//}
	//{
	//	var arr []int = []int{4,3,2,7,8,2,3,1}
	//	var res = findDuplicates(arr)
	//	fmt.Println(res)
	//}
	//
	//var arr []int = []int{1,2,2,3,3,3,3,4,4,5,6,6,7,7,8}
	//len := remove_duplicated_sorted_array(arr)
	//fmt.Println(len)
	//fmt.Println("start")
	//
	//
	//ll := tree_node{
	//	4,
	//	nil,
	//	nil,
	//}
	//lr := tree_node{
	//	8,
	//	nil,
	//	nil,
	//}
	//l :=tree_node{
	//	6,
	//	&ll,
	//	&lr,
	//}
	//
	//rr := tree_node{
	//	16,
	//	nil,
	//	nil,
	//}
	//rl := tree_node{
	//	12,
	//	nil,
	//	nil,
	//}
	//r := tree_node{
	//	14,
	//	&rl,
	//	&rr,
	//
	//}
	//root := tree_node{
	//	11,
	//	&l,
	//	&r,
	//}
	//inorder_tree(&root)
	//fmt.Println(root)
}