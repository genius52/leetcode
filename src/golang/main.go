package main

import (
	"fmt"
	"math"
	"oj/array"
	"oj/diagram"
	"oj/list_queue"
	"oj/number"
	"oj/other"
	"oj/stack"
	"oj/string_issue"
	"oj/tree"
	"strings"
)

func main() {
	{
		start := []int{6, 0, 3}
		d := 2
		res := array.MaxPossibleScore(start, d)
		fmt.Println(res)
	}
	{
		queries := [][]int{{1, 2}, {3, 4}, {2, 3}, {-3, 0}}
		k := 2
		res := number.ResultsArray(queries, k)
		fmt.Println(res)
	}
	{
		nums := []int{1, 2, 3, 4, 3, 2, 5}
		k := 3
		res := array.ResultsArray2(nums, k)
		fmt.Println(res)
	}
	{
		nums := []int{1, 2, 3, 4, 3, 2, 5}
		k := 3
		res := array.ResultsArray(nums, k)
		fmt.Println(res)
	}
	{
		s := "10101"
		k := 1
		res := string_issue.CountKConstraintSubstrings(s, k)
		fmt.Println(res)
	}
	{
		n := 5
		queries := [][]int{{2, 4}, {0, 2}, {0, 4}}
		res := diagram.ShortestDistanceAfterQueries3244(n, queries)
		fmt.Println(res)
	}
	{
		edges := [][]int{{2, 0}, {4, 2}, {1, 2}, {3, 1}, {5, 1}}
		res := tree.CountGoodNodes(edges)
		fmt.Println(res)
	}
	{
		n := 5
		queries := [][]int{{2, 4}, {0, 2}, {0, 4}}
		res := diagram.ShortestDistanceAfterQueries(n, queries)
		fmt.Println(res)
	}
	{
		m := 6
		n := 3
		horizontalCut := []int{2, 3, 2, 3, 1}
		verticalCut := []int{1, 2}
		res := number.MinimumCost3218(m, n, horizontalCut, verticalCut)
		fmt.Println(res)
	}
	{
		var l1 list_queue.ListNode
		l1.Val = 1
		var l2 list_queue.ListNode
		l2.Val = 2
		var l3 list_queue.ListNode
		l3.Val = 3
		var l4 list_queue.ListNode
		l4.Val = 4
		var l5 list_queue.ListNode
		l5.Val = 5
		l1.Next = &l2
		l2.Next = &l3
		l3.Next = &l4
		l4.Next = &l5
		nums := []int{1, 2, 3, 5}
		res := list_queue.ModifiedList(nums, &l1)
		fmt.Println(res)
	}
	{
		colors := []int{0, 1, 0, 1, 0}
		k := 3
		res := array.NumberOfAlternatingGroups2(colors, k)
		fmt.Println(res)
	}
	{
		var grid [][]byte = [][]byte{{'.', '.'}, {'Y', 'X'}}
		res := array.NumberOfSubmatrices(grid)
		fmt.Println(res)
	}
	{
		nums := []int{1, 2, 3, 4, 5}
		k := 2
		res := array.MaximumLength3202(nums, k)
		fmt.Println(res)
	}
	{
		nums := []int{1, 2, 3, 4}
		res := array.MaximumLength3201(nums)
		fmt.Println(res)
	}
	{
		red := 1
		blue := 1
		res := number.MaxHeightOfTriangle(red, blue)
		fmt.Println(res)
	}
	{
		//grid := [][]int{{1, 0, 1}, {1, 1, 1}}
		grid := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 1}, {1, 1, 0}}
		res := diagram.MinimumSum(grid)
		fmt.Println(res)
	}
	{
		nums := []int{1, -2, 3, 4}
		res := array.MaximumTotalCost(nums)
		fmt.Println(res)
	}
	{
		power := []int{1, 1, 3, 4}
		res := number.MaximumTotalDamage(power)
		fmt.Println(res)
	}
	{
		nums := []int{65, 25, 26, 70, 40, 46, 37}
		k := 5
		res := array.MaximumLength(nums, k)
		fmt.Println(res)
	}
	{
		rewardValues := []int{6, 13, 9, 19}
		res := number.MaxTotalReward(rewardValues)
		fmt.Println(res)
	}
	{
		n := 4
		k := 5
		res := number.ValueAfterKSeconds(n, k)
		fmt.Println(res)
	}
	{
		skills := []int{4, 2, 6, 3, 9}
		k := 2
		res := number.FindWinningPlayer(skills, k)
		fmt.Println(res)
	}
	{
		s := "aaba*"
		res := string_issue.ClearStars(s)
		fmt.Println(res)
	}
	{
		days := 10
		meetings := [][]int{{5, 7}, {1, 3}, {9, 10}}
		res := number.CountDays(days, meetings)
		fmt.Println(res)
	}
	{
		nums1 := []int{1, 2, 4, 12}
		nums2 := []int{2, 4}
		k := 3
		res := number.NumberOfPairs(nums1, nums2, k)
		fmt.Println(res)
	}
	{
		word := "aaaaaaaaaaaaaabb"
		res := string_issue.CompressedString(word)
		fmt.Println(res)
	}
	{
		limit := 1
		queries := [][]int{{0, 1}, {0, 4}, {0, 4}, {0, 1}, {1, 2}}
		res := number.QueryResults(limit, queries)
		fmt.Println(res)
	}
	{
		nums := []int{13, 23, 12, 33, 45, 25, 78, 58, 90, 43}
		res := array.SumDigitDifferences(nums)
		fmt.Println(res)
	}
	{
		nums := []int{4, 3, 1, 6}
		queries := [][]int{{0, 2}, {2, 3}}
		res := array.IsArraySpecial2(nums, queries)
		fmt.Println(res)
	}
	{
		s := "fabccddg"
		res := string_issue.MinimumSubstringsInPartition(s)
		fmt.Println(res)
	}
	{
		n := 3
		x := 4
		res := array.MinEnd(n, x)
		fmt.Println(res)
	}
	{
		nums1 := []int{4, 20, 16, 12, 8}
		nums2 := []int{14, 18, 10}
		res := array.MinimumAddedInteger(nums1, nums2)
		fmt.Println(res)
	}
	{
		zero := 199
		one := 199
		limit := 20
		res := array.NumberOfStableArrays(zero, one, limit)
		fmt.Println(res)
	}
	{
		nums := []int{98, 52}
		k := 82
		res := number.MinOperationsToMakeMedianK(nums, k)
		fmt.Println(res)
	}
	{
		s := "xaxcd"
		k := 4
		res := string_issue.GetSmallestString3106(s, k)
		fmt.Println(res)
	}
	{
		nums := []int{1, 2, 32, 21}
		k := 55
		res := array.MinimumSubarrayLength3097(nums, k)
		fmt.Println(res)
	}
	{
		nums := []int{2, 1, 8}
		k := 10
		res := array.MinimumSubarrayLength(nums, k)
		fmt.Println(res)
	}
	{
		k := 2
		res := number.MinOperations3091(k)
		fmt.Println(res)
	}
	{
		s := "bcbbbcba"
		res := string_issue.MaximumLengthSubstring(s)
		fmt.Println(res)
	}
	{
		word := "inn"
		k := 3
		res := string_issue.MinimumDeletions3085(word, k)
		fmt.Println(res)
	}
	{
		s := "abcdefghijklmnopqrstuvwxy??"
		res := string_issue.MinimizeStringValue(s)
		fmt.Println(res)
	}
	{
		nums := []int{2, 2, 0}
		changeIndices := []int{2, 2, 2, 2, 3, 2, 2, 1}
		res := number.EarliestSecondToMarkIndices(nums, changeIndices)
		fmt.Println(res)
	}
	{
		arr := []string{"abc", "bcd", "abcd"}
		res := string_issue.ShortestSubstrings(arr)
		fmt.Println(res)
	}
	{
		edges := [][]int{{0, 6, 3}, {6, 5, 3}, {0, 3, 1}, {3, 2, 7}, {3, 1, 6}, {3, 4, 2}}
		signalSpeed := 3
		res := tree.CountPairsOfConnectableServers(edges, signalSpeed)
		fmt.Println(res)
	}
	{
		grid := [][]int{{7, 6, 3}, {6, 6, 1}}
		k := 18
		res := array.CountSubmatrices(grid, k)
		fmt.Println(res)
	}
	{
		grid := [][]int{{1, 2, 2}, {1, 1, 0}, {0, 1, 0}}
		res := array.MinimumOperationsToWriteY(grid)
		fmt.Println(res)
	}
	{
		nums := []int{1, 2, 14, 15}
		res := array.ResultArray(nums)
		fmt.Println(res)
	}
	{
		arr1 := []int{1, 10, 100}
		arr2 := []int{1000}
		res := array.LongestCommonPrefix(arr1, arr2)
		fmt.Println(res)
	}
	{
		nums := []int{1, 9, 7, 3, 2, 7, 4, 12, 2, 6}
		res := array.MaxOperations3040(nums)
		fmt.Println(res)
	}
	{
		var res int64 = -1 << 63
		res = math.MinInt64
		fmt.Println(res)
	}
	{
		points := [][]int{{1, 5}, {2, 0}, {5, 5}}
		res := diagram.NumberOfPairs(points)
		fmt.Println(res)
	}
	{
		n := 3
		m := 2
		res := number.FowerGame(n, m)
		fmt.Println(res)
	}
	{
		nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
		res := number.MaximumLength(nums)
		fmt.Println(res)
	}
	{
		s := "aAbBcC"
		res := string_issue.CountKeyChanges(s)
		fmt.Println(res)
	}
	{
		word := "abzaqsqcyrbzsrvamylmyxdjl"
		res := number.MinimumPushes2(word)
		fmt.Println(res)
	}
	{
		nums := []int{20, 16}
		res := array.CanSortArray(nums)
		fmt.Println(res)
	}
	{
		nums := []int{1, 10000000}
		res := array.MaximumGap(nums)
		fmt.Println(res)
	}
	{
		s := "isawsquirrelnearmysquirrelhouseohmy"
		a := "my"
		b := "squirrel"
		k := 15
		res := array.BeautifulIndices(s, a, b, k)
		fmt.Println(res)
	}
	{
		nums1 := []int{10, 8, 7, 9}
		nums2 := []int{7, 9, 9, 5}
		res := number.MaximumSetSize(nums1, nums2)
		fmt.Println(res)
	}
	{
		a := 4
		b := 3
		c := 3
		d := 4
		e := 5
		f := 2
		res := number.MinMovesToCaptureTheQueen(a, b, c, d, e, f)
		fmt.Println(res)
	}
	{
		x := 26
		y := 1
		res := number.MinimumOperationsToMakeEqual(x, y)
		fmt.Println(res)
	}
	{
		nums := []int{46, 8, 2, 4, 1, 4, 10, 2, 4, 10, 2, 5, 7, 3, 1}
		res := array.MissingInteger(nums)
		fmt.Println(res)
	}
	{
		s := "aada"
		res := string_issue.MaximumLength2(s)
		fmt.Println(res)
	}
	{
		s := "cccerrrecdcdccedecdcccddeeeddcdcddedccdceeedccecde"
		res := string_issue.MaximumLength(s)
		fmt.Println(res)
	}
	{
		nums := []int{1, 2, 3, 4}
		res := array.IncremovableSubarrayCount(nums)
		fmt.Println(res)
	}
	{
		nums := []int{6, 10, 5, 12, 7, 11, 6, 6, 12, 12, 11, 7}
		k := 2
		res := array.DivideArray2966(nums, k)
		fmt.Println(res)
	}
	{
		grid := [][]int{{1, 3}, {2, 2}}
		res := array.FindMissingAndRepeatedValues(grid)
		fmt.Println(res)
	}
	{
		nums := []int{1, 3, 2, 3, 3}
		k := 2
		res := array.CountSubarrays2962(nums, k)
		fmt.Println(res)
	}
	{
		nums := []int{1, 2, 3, 1, 2, 3, 1, 2}
		k := 2
		res := array.MaxSubarrayLength(nums, k)
		fmt.Println(res)
	}
	{
		word := "zyxyxyz"
		res := string_issue.RemoveAlmostEqualCharacters(word)
		fmt.Println(res)
	}
	{
		coins := []int{1, 1, 1}
		target := 20
		res := number.MinimumAddedCoins(coins, target)
		fmt.Println(res)
	}
	{
		//n = 2, m = 3, hBars = [2,3], vBars = [2,3,4]
		n := 4
		m := 40
		hBars := []int{2, 3, 4, 5}
		vBars := []int{36, 41, 6, 34, 33}
		res := array.MaximizeSquareHoleArea(n, m, hBars, vBars)
		fmt.Println(res)
	}
	{
		nums := []int{4, 3, 23, 84, 34, 88, 44, 44, 18, 15}
		limit := 3
		res := array.LexicographicallySmallestArray(nums, limit)
		fmt.Println(res)
	}
	{
		s := "baeyh"
		k := 2
		res := string_issue.BeautifulSubstrings(s, k)
		fmt.Println(res)
	}
	{
		prices := []int{3, 1, 2}
		res := array.MinimumCoins(prices)
		fmt.Println(res)
	}
	{
		//53449611838892
		//b =
		//712958946092406
		var a int64 = 53449611838892
		var b int64 = 712958946092406
		var n int = 6
		res := number.MaximumXorProduct(a, b, n)
		fmt.Println(res)
	}
	{
		nums1 := []int{2, 3, 4, 5, 9}
		nums2 := []int{8, 8, 4, 4, 4}
		res := array.MinOperations2934(nums1, nums2)
		fmt.Println(res)
	}
	{
		access_times := [][]string{{"akuhmu", "0454"}, {"aywtqh", "0523"}, {"akuhmu", "0518"}, {"ihhkc", "0439"}, {"ihhkc", "0508"}, {"akuhmu", "0529"}, {"aywtqh", "0530"}, {"aywtqh", "0419"}}
		res := number.FindHighAccessEmployees(access_times)
		fmt.Println(res)
	}
	{
		n := 5
		res := number.StringCount(n)
		fmt.Println(res)
	}
	{
		n := 5
		limit := 2
		res := number.DistributeCandies2929(n, limit)
		fmt.Println(res)
	}
	{
		edges := [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {4, 5}}
		values := []int{5, 2, 5, 2, 1, 1}
		res := tree.MaximumScoreAfterOperations(edges, values)
		fmt.Println(res)
	}
	{
		nums1 := []int{0, 7, 28, 17, 18}
		nums2 := []int{1, 2, 6, 26, 1, 0, 27, 3, 0, 30}
		res := number.MinSum(nums1, nums2)
		fmt.Println(res)
	}
	{
		s := "1001"
		res := string_issue.MinChanges(s)
		fmt.Println(res)
	}
	{
		nums := []int{1, 2, 1}
		res := array.SumCounts(nums)
		fmt.Println(res)
	}
	{
		grid := [][]int{{1, 2}, {3, 4}}
		res := array.ConstructProductMatrix(grid)
		fmt.Println(res)
	}
	{
		nums := []int{5, 1, 4, 1}
		indexDifference := 2
		valueDifference := 4
		res := array.FindIndices2905(nums, indexDifference, valueDifference)
		fmt.Println(res)
	}
	{
		s := "1111111011111"
		k := 12
		res := string_issue.ShortestBeautifulSubstring(s, k)
		fmt.Println(res)
	}
	{
		s1 := "1100011000"
		s2 := "0101001010"
		//s1 := "1100"
		//s2 := "0101"
		x := 2
		res := string_issue.MinOperations2896(s1, s2, x)
		fmt.Println(res)
	}
	{
		processorTime := []int{121, 99}
		tasks := []int{287, 315, 293, 260, 333, 362, 69, 233}
		res := number.MinProcessingTime(processorTime, tasks)
		fmt.Println(res)
	}
	{
		nums := []int{22, 21, 29, 22}
		res := array.MaxSubarrays(nums)
		fmt.Println(res)
	}
	{
		nums := []int{1, 2, 3}
		target := 5
		res := array.MinSizeSubarray(nums, target)
		fmt.Println(res)
	}
	{
		nums := []int{5, 7, 1, 3}
		res := array.MaxSubarrays(nums)
		fmt.Println(res)
	}
	{
		maxHeights := []int{6, 5, 3, 9, 2, 7}
		res := array.MaximumSumOfHeights2866(maxHeights)
		fmt.Println(res)
	}
	{
		n := 3
		k := 2
		budget := 15
		composition := [][]int{{1, 1, 1}, {1, 1, 10}}
		stock := []int{0, 0, 0}
		cost := []int{1, 2, 3}
		res := number.MaxNumberOfAlloys(n, k, budget, composition, stock, cost)
		fmt.Println(res)
	}
	{
		nums := []int{3, 1, 9, 6}
		modulo := 3
		k := 0
		res := array.CountInterestingSubarrays(nums, modulo, k)
		fmt.Println(res)
	}
	{
		//num := "5079348434341304786187138148710190664"
		num := "25"
		res := number.MinimumOperations2844(num)
		fmt.Println(res)
	}
	{
		n := 5
		offers := [][]int{{0, 0, 1}, {0, 2, 2}, {1, 3, 2}}
		res := number.MaximizeTheProfit(n, offers)
		fmt.Println(res)
	}
	{
		n := 5
		k := 6
		res := number.MinimumSum2829(n, k)
		fmt.Println(res)
	}
	{
		nums := []int{2, 1, 3, 2, 1}
		res := array.MinimumOperations2826(nums)
		fmt.Println(res)
	}
	{
		str1 := "abc"
		str2 := "ad"
		res := string_issue.CanMakeSubsequence(str1, str2)
		fmt.Println(res)
	}
	{
		nums := []int{6, -1, 7, 4, 2, 3}
		target := 8
		res := array.CountPairs2824(nums, target)
		fmt.Println(res)
	}
	{
		nums := []int{2, 2, 1}
		m := 4
		res := array.CanSplitArray(nums, m)
		fmt.Println(res)
	}
	{
		s := "string"
		res := string_issue.FinalString(s)
		fmt.Println(res)
	}
	{
		nums := []int{11, 4, 10}
		res := array.MinimumSeconds(nums)
		fmt.Println(res)
	}
	{
		a := "cab"
		b := "a"
		c := "b"
		res := string_issue.MinimumString(a, b, c)
		fmt.Println(res)
	}
	{
		nums := []int{1, 3, 1, 2, 2}
		res := array.CountCompleteSubarrays(nums)
		fmt.Println(res)
	}
	{
		parent := []int{-1, 0, 0, 1, 1, 2}
		s := "acaabc"
		res := tree.CountPalindromePaths(parent, s)
		fmt.Println(res)
	}
	{
		n := 114
		x := 1
		res := number.NumberOfWays2787(n, x)
		fmt.Println(res)
	}
	{
		nums := []int{2, 3, 6, 1, 9, 2}
		x := 5
		res := array.MaxScore2(nums, x)
		fmt.Println(res)
	}
	{
		nums := []int{2, 2, 3, 1, 1, 0}
		k := 3
		res := array.CheckArray(nums, k)
		fmt.Println(res)
	}
	{
		//nums1 := []int{3, 19, 13, 19}
		//nums2 := []int{20, 18, 7, 14}
		nums1 := []int{11, 7, 7, 9}
		nums2 := []int{19, 19, 1, 7}
		res := array.MaxNonDecreasingLength(nums1, nums2)
		fmt.Println(res)
	}
	{
		nums := []int{1, 3, 6, 4, 1, 2}
		target := 3
		res := array.MaximumJumps(nums, target)
		fmt.Println(res)
	}
	{
		s := "1011"
		res := string_issue.MinimumBeautifulSubstrings(s)
		fmt.Println(res)
	}
	{
		num1 := 135
		num2 := 26
		res := number.MakeTheIntegerZero(num1, num2)
		fmt.Println(res)
	}
	{
		//4
		//[[4,3],[2,16],[1,21],[3,22],[1,13],[3,10],[2,1],[1,12],[4,13],[2,18]]
		//8
		//[14,28,29]
		n := 4
		logs := [][]int{{4, 3}, {2, 16}, {1, 21}, {3, 22}, {1, 13}, {3, 10}, {2, 1}, {1, 12}, {4, 13}, {2, 18}}
		x := 8
		queries := []int{14, 28, 29}
		res := number.CountServers(n, logs, x, queries)
		fmt.Println(res)
	}
	{
		//words := []string{"aad", "ebeb", "jgd", "ghdhc", "jcdc", "ffigd", "b", "jae", "j", "bbg", "gaii", "iha", "fj", "h", "aajf", "ccaf", "digbc", "efdi", "eb", "ff"}
		words := []string{"aa", "ab", "bc"}
		res := string_issue.MinimizeConcatenatedLength(words)
		fmt.Println(res)
	}
	{
		x := 2
		y := 5
		z := 1
		res := string_issue.LongestString(x, y, z)
		fmt.Println(res)
	}
	{
		words := []string{"ku", "dd", "gu", "uk"}
		res := string_issue.MaximumNumberOfStringPairs(words)
		fmt.Println(res)
	}
	{
		nums := []int{2, 3, 6}
		res := array.SpecialPerm(nums)
		fmt.Println(res)
	}
	{
		nums := []int{20, 1, 15, 6, 7}
		x := 5
		res := number.MinCost2735(nums, x)
		fmt.Println(res)
	}
	{
		s := "acbbc"
		res := string_issue.SmallestString(s)
		fmt.Println(res)
	}
	{
		nums := []int{3, 30, 24}
		res := number.FindNonMinOrMax(nums)
		fmt.Println(res)
	}
	{
		nums := []int{-2, 0, 2}
		s := "RLL"
		d := 3
		res := number.SumDistance(nums, s, d)
		fmt.Println(res)
	}
	{
		s := "0011"
		res := string_issue.MinimumCost(s)
		fmt.Println(res)
	}
	{
		nums := []int{9, 3, 6}
		res := array.MaxStrength(nums)
		fmt.Println(res)
	}
	{
		s := "leetscode"
		dictionary := []string{"leet", "code", "leetcode"}
		res := string_issue.MinExtraChar(s, dictionary)
		fmt.Println(res)
	}
	{
		n := 37
		res := number.PunishmentNumber(n)
		fmt.Println(res)
	}
	{
		s := "ABFCACDB"
		res := string_issue.MinLength(s)
		fmt.Println(res)
	}
	{
		n := 6
		edges := [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}}
		res := diagram.CountCompleteComponents(n, edges)
		fmt.Println(res)
	}
	{
		grid := [][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}
		res := array.MaxMoves(grid)
		fmt.Println(res)
	}
	{
		derived := []int{1, 1, 0}
		res := number.DoesValidArrayExist(derived)
		fmt.Println(res)
	}
	{
		n := 7
		cost := []int{1, 5, 2, 2, 3, 3, 1}
		res := tree.MinIncrements(n, cost)
		fmt.Println(res)
	}
	{
		n := 4
		var edges [][]int = [][]int{{0, 1}, {1, 2}, {1, 3}}
		var price []int = []int{2, 2, 10, 6}
		var trips [][]int = [][]int{{0, 3}, {2, 1}, {2, 3}}
		res := tree.MinimumTotalPrice(n, edges, price, trips)
		fmt.Println(res)
	}
	//{
	//	start := []int{1, 1}
	//	target := []int{4, 5}
	//	specialRoads := [][]int{{1, 2, 3, 3, 2}, {3, 4, 4, 5, 1}}
	//	res := diagram.MinimumCost(start, target, specialRoads)
	//	fmt.Println(res)
	//}
	{
		A := []int{1, 3, 2, 4}
		B := []int{3, 1, 2, 4}
		res := array.FindThePrefixCommonArray(A, B)
		fmt.Println(res)
	}
	{
		nums := []int{2, 3, 7, 5, 10}
		res := array.FindPrefixScore(nums)
		fmt.Println(res)
	}
	{
		coins := []int{1, 0, 0, 0, 0, 1}
		edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}}
		res := tree.CollectTheCoins(coins, edges)
		fmt.Println(res)
	}
	{
		nums := []int{4, 2, 1, 2}
		p := 1
		res := array.MinimizeMax(nums, p)
		fmt.Println(res)
	}
	{
		nums := []int{1, 3, 1, 1, 2}
		res := array.Distance(nums)
		fmt.Println(res)
	}
	{
		arr := []int{1, 4, 1, 3}
		k := 3
		res := array.MakeSubKSumEqual(arr, k)
		fmt.Println(res)
	}
	{
		nums := []int{3, 1, 6, 8}
		queries := []int{1, 5}
		res := array.MinOperations2602(nums, queries)
		fmt.Println(res)
	}
	{
		nums := []int{4, 9, 6, 10}
		res := array.PrimeSubOperation(nums)
		fmt.Println(res)
	}
	{
		nums := []int{1, -10, 7, 13, 6, 8}
		value := 5
		res := number.FindSmallestInteger(nums, value)
		fmt.Println(res)
	}
	{
		nums := []int{4, 5, 6}
		k := 2
		res := number.BeautifulSubsets(nums, k)
		fmt.Println(res)
	}
	{
		nums := []int{2, 1, 3, 4, 5, 2}
		res := array.FindScore(nums)
		fmt.Println(res)
	}
	{
		nums := []int{1, 2, 3, 4}
		res := array.MaximizeGreatness(nums)
		fmt.Println(res)
	}
	{
		money := 17
		children := 2
		res := number.DistMoney(money, children)
		fmt.Println(res)
	}
	{
		nums := []int{4, 3, 1, 2, 4}
		res := array.BeautifulSubarrays(nums)
		fmt.Println(res)
	}
	{
		nums := []int{2, -1, 0, 1, -3, 3, -3}
		res := array.MaxScore(nums)
		fmt.Println(res)
	}
	{
		n := 6
		edges := [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {3, 5}}
		price := []int{9, 8, 7, 6, 10, 5}
		res := diagram.MaxOutput(n, edges, price)
		fmt.Println(res)
	}
	{
		num := 4325
		res := number.SplitNum(num)
		fmt.Println(res)
	}
	{
		word := "998244353"
		m := 3
		res := string_issue.DivisibilityArray(word, m)
		fmt.Println(res)
	}
	{
		nums := []int{42, 83, 48, 10, 24, 55, 9, 100, 10, 17, 17, 99, 51, 32, 16, 98, 99, 31, 28, 68, 71, 14, 64, 29, 15, 40}
		res := array.MaxNumOfMarkedIndices(nums)
		fmt.Println(res)
	}
	{
		nums := []int{3, 4, 4, 5}
		res := number.SquareFreeSubsets(nums)
		fmt.Println(res)
	}
	{
		n := 27
		res := number.MinOperations2571(n)
		fmt.Println(res)
	}
	{
		nums := []int{2, 1, 3, 4, 5, 6}
		res := number.MinImpossibleOR(nums)
		fmt.Println(res)
	}
	{
		num := 11891
		res := number.MinMaxDifference(num)
		fmt.Println(res)
	}
	{
		s := "abacaba"
		t := "bzaa"
		res := string_issue.MinimumScore(s, t)
		fmt.Println(res)
	}
	{
		basket1 := []int{84, 80, 43, 8, 80, 88, 43, 14, 100, 88}
		basket2 := []int{32, 32, 42, 68, 68, 100, 42, 84, 14, 8}
		res := array.MinCost2561(basket1, basket2)
		fmt.Println(res)
	}
	{
		//[1,1,0,0,0,0]
		//[1,1,1,1,1,1]
		//[1,1,1,1,1,1]
		//[1,1,1,1,0,1]
		//[0,0,0,1,1,1]
		//[0,0,0,1,0,1]
		//[0,0,0,1,0,1]
		//[0,0,0,1,0,1]
		//[0,0,0,1,0,1]
		//[0,0,0,1,1,1]]
		grid := [][]int{{1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 0, 1}, {0, 0, 0, 1, 1, 1}, {0, 0, 0, 1, 0, 1}, {0, 0, 0, 1, 0, 1}, {0, 0, 0, 1, 0, 1}, {0, 0, 0, 1, 0, 1}, {0, 0, 0, 1, 1, 1}}
		//grid := [][]int{{1,1,1},{1,0,1},{1,1,1}}
		res := array.IsPossibleToCutPath(grid)
		fmt.Println(res)
	}
	{
		prizePositions := []int{1, 1, 2, 2, 3, 3, 5}
		k := 2
		res := array.MaximizeWin(prizePositions, k)
		fmt.Println(res)
	}
	{
		nums := []int{2, 7, 9, 3, 1}
		k := 2
		res := array.MinCapability(nums, k)
		fmt.Println(res)
	}
	{
		nums := []int{1, 2, 1, 2}
		k := 2
		res := array.MinCost2547(nums, k)
		fmt.Println(res)
	}
	{
		nums := []int{3, 1, 4, 3, 2, 2, 4}
		k := 2
		res := array.CountGood(nums, k)
		fmt.Println(res)
	}
	{
		n := 2
		queries := [][]int{{0, 0, 1, 1}}
		res := array.RangeAddQueries(n, queries)
		fmt.Println(res)
	}
	{
		word1 := "aa"
		word2 := "ab"
		res := string_issue.IsItPossible(word1, word2)
		fmt.Println(res)
	}
	{
		s := "165462"
		k := 60
		res := string_issue.MinimumPartition(s, k)
		fmt.Println(res)
	}
	{
		nums := []int{2, 4, 3, 7, 10, 6}
		res := number.DistinctPrimeFactors(nums)
		fmt.Println(res)
	}
	{
		price := []int{1, 3, 1}
		k := 2
		res := number.MaximumTastiness(price, k)
		fmt.Println(res)
	}
	{
		s := "aabaaaacaabc"
		k := 2
		res := string_issue.TakeCharacters(s, k)
		fmt.Println(res)
	}
	{
		divisor1 := 2
		divisor2 := 7
		uniqueCnt1 := 1
		uniqueCnt2 := 3
		res := array.MinimizeSet(divisor1, divisor2, uniqueCnt1, uniqueCnt2)
		fmt.Println(res)
	}
	{
		forts := []int{1, 0, 0, -1, 0, 0, 0, 0, 1}
		res := array.CaptureForts(forts)
		fmt.Println(res)
	}
	{
		//n := 11
		//edges := [][]int{{5,9},{8,1},{2,3},{7,10},{3,6},{6,7},{7,8},{5,1},{5,7},{10,11},{3,7},{6,11},{8,11},{3,4},{8,9},{9,1},{2,10},{9,11},{5,11},{2,5},{8,10},{2,7},{4,1},{3,10},{6,1},{4,9},{4,6},{4,5},{2,4},{2,11},{5,8},{6,9},{4,10},{3,11},{4,7},{3,5},{7,1},{2,9},{6,10},{10,1},{5,6},{3,9},{2,6},{7,9},{4,11},{4,8},{6,8},{3,8},{9,10},{5,10},{2,8},{7,11}}
		//n := 4
		//edges := [][]int{{1,2},{3,4}}
		n := 21
		edges := [][]int{{2, 19}, {16, 17}, {8, 14}, {2, 16}, {12, 20}, {12, 14}, {16, 18}, {15, 16}, {10, 21}, {3, 5}, {13, 18}, {17, 20}, {14, 17}, {9, 12}, {5, 15}, {5, 6}, {3, 7}, {2, 21}, {10, 13}, {8, 16}, {7, 18}, {4, 6}, {9, 1}, {13, 21}, {18, 20}, {7, 14}, {4, 19}, {5, 8}, {3, 11}, {11, 1}, {7, 12}, {4, 7}, {3, 16}, {13, 17}, {17, 19}, {9, 13}, {7, 19}, {10, 16}, {4, 13}, {4, 5}, {2, 15}, {12, 19}, {11, 16}, {2, 9}, {11, 17}, {17, 1}, {16, 21}, {4, 10}, {10, 14}, {14, 16}, {4, 1}, {13, 20}, {5, 20}, {4, 14}, {4, 21}, {10, 20}, {2, 14}, {8, 15}, {4, 8}, {6, 19}, {15, 1}, {19, 1}, {8, 19}, {15, 21}, {3, 12}, {11, 18}, {9, 17}, {18, 19}, {7, 21}, {3, 21}, {16, 19}, {11, 15}, {5, 1}, {8, 17}, {3, 15}, {8, 1}, {10, 19}, {3, 8}, {6, 16}, {2, 8}, {5, 18}, {11, 13}, {11, 20}, {14, 21}, {6, 20}, {4, 20}, {12, 13}, {5, 12}, {10, 11}, {9, 15}, {3, 19}, {9, 20}, {14, 18}, {21, 1}, {13, 19}, {8, 21}, {2, 13}, {3, 10}, {9, 18}, {19, 21}, {6, 7}, {3, 18}, {2, 18}, {6, 14}, {3, 17}, {5, 21}, {14, 20}, {8, 9}, {16, 1}, {3, 4}, {13, 1}, {5, 9}, {4, 15}, {17, 21}, {20, 21}, {2, 17}, {13, 14}, {11, 14}, {9, 16}, {10, 18}, {6, 15}, {6, 12}, {3, 13}, {5, 11}, {6, 1}, {12, 17}, {8, 10}, {5, 10}, {8, 18}, {4, 12}, {10, 1}, {6, 13}, {4, 18}, {7, 20}, {7, 16}, {2, 6}, {12, 21}, {4, 17}, {15, 18}, {13, 16}, {15, 20}, {7, 10}, {6, 10}, {2, 20}, {7, 15}, {18, 1}, {12, 1}, {3, 20}, {7, 1}, {14, 15}, {4, 9}, {11, 19}, {7, 9}, {5, 17}, {18, 21}, {6, 21}, {8, 11}, {6, 17}, {3, 14}, {7, 11}, {5, 7}, {7, 13}, {6, 8}, {6, 9}, {10, 12}, {5, 16}, {2, 4}, {17, 18}, {9, 11}, {12, 16}, {3, 6}, {12, 18}, {3, 9}, {11, 12}, {14, 19}, {10, 15}, {5, 13}, {8, 13}, {15, 17}, {2, 10}, {11, 21}, {20, 1}, {6, 18}, {2, 12}, {19, 20}, {6, 11}, {8, 12}, {2, 3}, {12, 15}, {2, 11}, {9, 10}, {7, 17}, {9, 19}, {13, 15}, {7, 8}, {4, 11}, {2, 5}, {5, 19}, {16, 20}, {15, 19}, {9, 14}, {14, 1}, {10, 17}, {9, 21}, {2, 7}, {8, 20}, {5, 14}, {4, 16}}
		res := diagram.IsPossible(n, edges)
		fmt.Println(res)
	}
	{
		var n int = 15
		res := number.SmallestValue(n)
		fmt.Println(res)
	}
	{
		//[2,5,1,4,3,6]
		//1
		//nums := []int{2,5,1,4,3,6}
		//k := 1
		nums := []int{3, 2, 1, 4, 5}
		k := 4
		res := array.CountSubarrays2488(nums, k)
		fmt.Println(res)
	}
	{
		n := 4
		roads := [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}
		res := diagram.MinScore(n, roads)
		fmt.Println(res)
	}
	{
		s := "0000000"
		res := string_issue.CountPalindromes(s)
		fmt.Println(res)
	}
	{
		//[5,2,13,3,8]
		var l1 list_queue.ListNode
		l1.Val = 1
		var l2 list_queue.ListNode
		l2.Val = 1
		var l3 list_queue.ListNode
		l3.Val = 1
		var l4 list_queue.ListNode
		l4.Val = 1
		var l5 list_queue.ListNode
		l5.Val = 8
		l1.Next = &l2
		l2.Next = &l3
		l3.Next = &l4
		//l4.Next = &l5
		res := list_queue.RemoveNodes(&l1)
		fmt.Println(res.Val)
	}
	{
		customers := "YYNY"
		res := number.BestClosingTime(customers)
		fmt.Println(res)
	}
	{
		s := "23542185131"
		k := 3
		minLength := 3
		res := number.BeautifulPartitions(s, k, minLength)
		fmt.Println(res)
	}
	{
		//	TreeAncestor treeAncestor = new TreeAncestor(7, [-1, 0, 0, 1, 1, 2, 2]);
		//treeAncestor.getKthAncestor(3, 1); // returns 1 which is the parent of 3
		//treeAncestor.getKthAncestor(5, 2); // returns 0 which is the grandparent of 5
		//treeAncestor.getKthAncestor(6, 3);
		n := 7
		parent := []int{-1, 0, 0, 1, 1, 2, 2}
		obj := tree.Constructor1483(n, parent)
		res := obj.GetKthAncestor(3, 1)
		res = obj.GetKthAncestor(5, 2)
		res = obj.GetKthAncestor(6, 3)
		fmt.Println(res)
	}
	{
		s := "dwdwdnns"
		k := 3
		res := string_issue.MaxPalindromes(s, k)
		fmt.Println(res)
	}
	{
		pizza := []string{"A..", "AAA", "..."}
		k := 3
		res := array.Ways(pizza, k)
		fmt.Println(res)
	}
	{
		nums1 := []int{4, 7, 6, 3}
		nums2 := []int{3, 4, 6, 7}
		var data_idx map[int]int = make(map[int]int)
		var l int = len(nums1)
		for i := 0; i < l; i++ {
			data_idx[nums1[i]] = i
		}
		var res int = 0
		for i := 0; i < l; i++ {
			if nums1[i] == nums2[i] {
				continue
			}
			nums1[i], nums1[data_idx[nums2[i]]] = nums2[i], nums1[i]
			res++
		}
		fmt.Println(res)
	}
	{
		nums := []int{3, 6, 2, 7, 1}
		k := 6
		res := array.SubarrayLCM(nums, k)
		fmt.Println(res)
	}
	{
		//[[0,1},{1,2},{2,3]]
		//3
		//[-5644,-6018,1188,-8502]
		edges := [][]int{{0, 1}, {1, 2}, {2, 3}}
		bob := 3
		amount := []int{-5644, -6018, 1188, -8502}
		res := diagram.MostProfitablePath(edges, bob, amount)
		fmt.Println(res)
	}
	{
		low := 2
		high := 3
		zero := 1
		one := 2
		res := number.CountGoodStrings(low, high, zero, one)
		fmt.Println(res)
	}
	{
		nums := []int{15, 77, 97, 26, 53, 41, 26, 13, 12, 18, 17, 42, 13, 33, 34, 70, 48, 65, 62, 78, 33, 60, 96, 53, 23, 14, 60, 70, 57, 41, 29, 12, 79, 65, 52, 30}
		res := array.DistinctAverages(nums)
		fmt.Println(res)
	}
	{
		seats := [][]byte{{'#', '.', '#', '#', '.', '#'},
			{'.', '#', '#', '#', '#', '.'},
			{'#', '.', '#', '#', '.', '#'}}
		//seats := [][]byte{{'.', '.', '.', '.', '#', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'},
		//	{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '#', '.'},
		//	{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '#', '.', '.', '.', '.', '.'},
		//	{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '#', '.', '.', '#', '.'}}
		res := array.MaxStudents(seats)
		fmt.Println(res)
	}
	{
		//[9,11,99,101]
		//[[10,1},{7,1},{14,1},{100,1},{96,1},{103,1]]
		//[0,4,6]
		//factory =
		//[[2,2},{6,2]]
		robot := []int{0, 4, 6}
		factory := [][]int{{2, 2}, {6, 2}}
		res := array.MinimumTotalDistance(robot, factory)
		fmt.Println(res)
	}
	{
		//[1,1,1,7,8,9]
		//3
		//[11,11,7,2,9,4,17,1]
		//1
		nums := []int{11, 11, 7, 2, 9, 4, 17, 1}
		k := 1
		res := array.MaximumSubarraySum(nums, k)
		fmt.Println(res)
	}
	{
		var n int64 = 94598
		target := 1
		res := number.MakeIntegerBeautiful(n, target)
		fmt.Println(res)
	}
	{
		words := []string{"mll", "edd", "jii", "tss", "fee", "dcc", "nmm", "abb", "utt", "zyy", "xww", "tss", "wvv", "xww", "utt"}
		res := string_issue.OddString(words)
		fmt.Println(res)
	}
	{
		nums := []int{2, 3, 1, 5, 4}
		res := array.MaxValueAfterReverse(nums)
		fmt.Println(res)
	}
	{
		//[1,5,3,6,7]
		//[1,6,3,3]
		arr1 := []int{1, 5, 3, 6, 7}
		arr2 := []int{1, 6, 3, 3}
		res := array.MakeArrayIncreasing(arr1, arr2)
		fmt.Println(res)
	}
	{
		//[1,2,5]
		//[4,1,3]
		nums := []int{1, 2, 5}
		target := []int{4, 1, 3}
		res := array.MakeSimilar(nums, target)
		fmt.Println(res)
	}
	{
		//[15,11,5]
		//1
		nums := []int{19, 9, 9}
		k := 1
		res := array.SubarrayGCD(nums, k)
		fmt.Println(res)
	}
	{
		nums := []int{1, 3, 5, 2, 7, 5}
		minK := 1
		maxK := 5
		res := array.CountSubarrays2444(nums, minK, maxK)
		fmt.Println(res)
	}
	{
		//[1,2,1,1,1]
		//[[0,1},{1,3},{3,4},{4,2]]
		nums := []int{6, 2, 2, 2, 6}
		edges := [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}
		res := diagram.ComponentValue(nums, edges)
		fmt.Println(res)
	}
	{
		num := 99
		res := number.SumOfNumberAndReverse(num)
		fmt.Println(res)
	}
	{
		nums := []int{3, 7, 1, 6}
		res := array.MinimizeArrayValue(nums)
		fmt.Println(res)
	}
	{
		//n := 15
		//queries := [][]int{{0,1},{2,2},{0,3}}
		n := 2
		queries := [][]int{{0, 0}}
		res := number.ProductQueries(n, queries)
		fmt.Println(res)
	}
	{
		var l1 list_queue.ListNode
		l1.Val = 1
		var l2 list_queue.ListNode
		l2.Val = 2
		var l3 list_queue.ListNode
		l3.Val = 3
		var l4 list_queue.ListNode
		l4.Val = 4
		l1.Next = &l2
		l2.Next = &l3
		l3.Next = &l4
		res := list_queue.DeleteMiddle(&l1)
		fmt.Println(res)
	}
	{
		s := "aaaaab"
		res := string_issue.LastSubstring(s)
		fmt.Println(res)
	}
	//{
	//	// 制造一个协程泄露事件
	//	for i := 0; i < 5; i++ {
	//		go func() {
	//			select {
	//			default:
	//				return
	//			} // 不阻塞
	//		}()
	//		go func() {
	//			select {} // 泄漏了 1000 个协程
	//		}()
	//	}
	//	mux := http.NewServeMux()
	//	mux.HandleFunc("/debug/pprof/", pprof.Index)
	//	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	//	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	//	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	//	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	//
	//	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	//		w.Write([]byte("hello"))
	//	})
	//	http.ListenAndServe(":6066", mux)
	//}
	{
		str1 := "abac"
		str2 := "cab"
		res := string_issue.ShortestCommonSupersequence2(str1, str2)
		fmt.Println(res)
	}
	{
		grid := [][]int{{1, 5, 3, 7, 3, 2, 3, 5}}
		k := 29
		res := diagram.NumberOfPaths(grid, k)
		fmt.Println(res)
	}
	{
		s := "aaaaa"
		res := string_issue.DeleteString(s)
		fmt.Println(res)
	}
	{
		s := "bac"
		res := string_issue.RobotWithString(s)
		fmt.Println(res)
	}
	{
		n := 10
		logs := [][]int{{0, 3}, {2, 5}, {0, 9}, {1, 15}}
		res := array.HardestWorker(n, logs)
		fmt.Println(res)
	}
	{
		num1 := 25
		num2 := 72
		res := number.MinimizeXor(num1, num2)
		fmt.Println(res)
	}
	{
		N := 37
		L := 50
		K := 8
		res := number.NumMusicPlaylists(N, L, K)
		fmt.Println(res)
	}
	{
		var dominoes string = ".L.R...LR..L.."
		res := diagram.PushDominoes(dominoes)
		fmt.Println(res)
	}
	{
		nums := []int{878724, 201541, 179099, 98437, 35765, 327555, 475851, 598885, 849470, 943442}
		k := 4
		res := array.GoodIndices(nums, k)
		fmt.Println(res)
	}
	{
		nums := []int{96317, 96317, 96317, 96317, 96317, 96317, 96317, 96317, 96317, 279979}
		res := array.LongestSubarray2419(nums)
		fmt.Println(res)
	}
	{
		n := 12
		res := number.ConcatenatedBinary(n)
		fmt.Println(res)
	}
	{
		transactions := [][]int{{2, 1}, {5, 0}, {4, 2}}
		res := number.MinimumMoney(transactions)
		fmt.Println(res)
	}
	{
		var t1 tree.TreeNode
		t1.Val = 7
		var t2 tree.TreeNode
		t2.Val = 13
		var t3 tree.TreeNode
		t3.Val = 11
		t1.Left = &t2
		t1.Right = &t3
		res := tree.ReverseOddLevels(&t1)
		fmt.Println(res)
	}
	{
		nums := []int{1, 0, 0}
		res := array.SmallestSubarrays(nums)
		fmt.Println(res)
	}
	{
		arriveAlice := "09-01"
		leaveAlice := "10-19"
		arriveBob := "06-19"
		leaveBob := "10-20"
		res := number.CountDaysTogether(arriveAlice, leaveAlice, arriveBob, leaveBob)
		fmt.Println(res)
	}
	{
		var i int32 = 1<<31 - 1
		fmt.Println(i)
	}
	{
		grid := [][]int{{1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}}
		res := tree.Construct(grid)
		fmt.Println(res)
	}
	{
		var a int = 1 | 0b00111000
		fmt.Println(a)
		data := []int{230, 136, 145}
		res := number.ValidUtf8(data)
		fmt.Println(res)
	}
	{
		nums := []int{0, 1, 2, 2, 4, 4, 1}
		res := number.MostFrequentEven(nums)
		fmt.Println(res)
	}
	{
		nums := []int{1, 3, 8, 48, 10}
		res := array.LongestNiceSubarray(nums)
		fmt.Println(res)
	}
	{
		startPos := 989
		endPos := 1000
		k := 99
		res := number.NumberOfWays2400(startPos, endPos, k)
		fmt.Println(res)
	}
	{
		//mat := [][]int{{1,0,0,1,0,0,1,1,1,1,0,1},{1,1,1,1,0,0,1,1,1,0,0,1},{0,1,1,0,0,0,1,1,0,1,1,1},{0,0,0,1,0,0,1,0,1,0,1,1},{0,1,1,1,0,0,0,1,1,1,1,0},{1,1,0,1,0,1,1,1,1,0,1,1},{1,1,0,0,0,1,0,0,0,0,1,1},{1,1,1,0,1,0,0,1,1,0,0,1},{1,0,1,1,1,0,1,0,0,0,1,0},{0,0,0,1,1,0,1,1,1,1,1,1},{1,1,0,1,0,0,1,1,1,0,1,1},{0,0,1,0,0,1,1,1,1,0,1,1}}
		//cols := 11
		mat := [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}}
		cols := 2
		res := array.MaximumRows(mat, cols)
		fmt.Println(res)
	}
	{
		n := 9
		res := number.IsStrictlyPalindromic(n)
		fmt.Println(res)
	}
	{
		k := 3
		rowConditions := [][]int{{1, 2}, {3, 2}}
		colConditions := [][]int{{2, 1}, {3, 2}}
		res := diagram.BuildMatrix(k, rowConditions, colConditions)
		fmt.Println(res)
	}
	{
		//intervals = [[1,3},{1,4},{2,5},{3,5]]
		//var intervals [][]int = [][]int{{2, 10}, {3, 7}, {3, 15}, {4, 11}, {6, 12}, {6, 16}, {7, 8}, {7, 11}, {7, 15}, {11, 12}}
		//var intervals [][]int = [][]int{{1,3},{1,4},{2,5},{3,5}}
		var intervals [][]int = [][]int{{1, 3}, {3, 7}, {5, 7}, {7, 8}}
		res := array.IntersectionSizeTwo(intervals)
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
		t2.Left = &t3
		t3.Left = &t4
		t4.Left = &t5
		start := 2
		res := tree.AmountOfTime(&t1, start)
		fmt.Println(res)
	}
	{
		num := "0000"
		res := number.LargestPalindromic(num)
		fmt.Println(res)
	}
	{
		//1
		//1
		//[1,1,1,1]
		//[1,1,1,50]
		//100
		//100
		//[1,2,3,4,5,6,7,8,9]
		//[1,2,3,1,2,3,1,2,10]
		initialEnergy := 100
		initialExperience := 100
		energy := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		experience := []int{1, 2, 3, 1, 2, 3, 1, 2, 10}
		res := number.MinNumberOfHours(initialEnergy, initialExperience, energy, experience)
		fmt.Println(res)
	}
	{
		s := "dztz"
		shifts := [][]int{{0, 0, 0}, {1, 1, 1}}
		res := string_issue.ShiftingLetters3(s, shifts)
		fmt.Println(res)
	}
	{
		s := "11100"
		res := string_issue.SecondsToRemoveOccurrences(s)
		fmt.Println(res)
	}
	{
		blocks := "WBBWWBBWBW"
		k := 7
		res := string_issue.MinimumRecolors(blocks, k)
		fmt.Println(res)
	}
	{
		var stickers []string = []string{"with", "example", "science"}
		var target string = "thehat"
		res := string_issue.MinStickers(stickers, target)
		fmt.Println(res)
	}
	{
		n := 99
		res := number.CountSpecialNumbers(n)
		fmt.Println(res)
	}
	{
		grid := [][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}
		res := array.LargestLocal(grid)
		fmt.Println(res)
	}
	{
		n := 13
		k := 2
		res := number.FindKthNumber(n, k)
		fmt.Println(res)
	}
	{
		nums := []int{7, 2, 5, 10, 8}
		m := 2
		res := array.SplitArray(nums, m)
		fmt.Println(res)
	}
	{
		s := "acfgbd"
		k := 2
		res := string_issue.LongestIdealString(s, k)
		fmt.Println(res)
	}
	{
		nums := []int{3, 9, 3}
		res := array.MinimumReplacement(nums)
		fmt.Println(res)
	}
	{
		edges := []int{-1, 4, -1, 2, 0, 4}
		res := diagram.LongestCycle(edges)
		fmt.Println(res)
	}
	{
		var nums []int = []int{1, 3}
		var n int = 6
		res := number.MinPatches(nums, n)
		fmt.Println(res)
	}
	{
		s := "catsanddog"
		wordDict := []string{"cat", "cats", "and", "sand", "dog"}
		res := string_issue.WordBreak140(s, wordDict)
		fmt.Println(res)
	}
	{
		//[1,2,3,1,10,11,22,12,32,22,44,31,23,45,64]
		//7
		nums := []int{1, 2, 3, 1, 10, 11, 22, 12, 32, 22, 44, 31, 23, 45, 64}
		k := 7
		res := number.CountExcellentPairs(nums, k)
		fmt.Println(res)
	}
	{
		//["NumberContainers","change","find","change","find","find","find"]
		//[[},{1,10},{10},{1,20},{10},{20},{30]]
		obj := number.Constructor2349()
		obj.Change(1, 10)
		res := obj.Find(10)
		obj.Change(1, 20)
		res = obj.Find(20)
		fmt.Println(res)
	}
	{
		nums := []string{"64333639502", "65953866768", "17845691654", "87148775908", "58954177897", "70439926174", "48059986638", "47548857440", "18418180516", "06364956881", "01866627626", "36824890579", "14672385151", "71207752868"}
		queries := [][]int{{9, 4}, {6, 1}, {3, 8}, {12, 9}, {11, 4}, {4, 9}, {2, 7}, {10, 3}, {13, 1}, {13, 1}, {6, 1}, {5, 10}}
		res := number.SmallestTrimmedNumbers(nums, queries)
		fmt.Println(res)
	}
	{
		var routes [][]int = [][]int{{1, 2, 7}, {3, 6, 7}}
		var source int = 1
		var target int = 6
		res := diagram.NumBusesToDestination(routes, source, target)
		fmt.Println(res)
	}
	{
		var nums []int = []int{14, 9, 8, 4, 3, 2}
		res := array.CanPartition2(nums)
		fmt.Println(res)
	}
	{
		numCourses := 2
		prerequisites := [][]int{{1, 0}}
		res := diagram.FindOrder(numCourses, prerequisites)
		fmt.Println(res)
	}
	{
		numRows := 5
		res := array.Generate(numRows)
		fmt.Println(res)
	}
	{
		nums := []int{1, 3, 4, 3, 1}
		threshold := 6
		res := array.ValidSubarraySize(nums, threshold)
		fmt.Println(res)
	}
	{
		m := 2
		n := 2
		N := 2
		i := 0
		j := 0
		res := diagram.FindPaths(m, n, N, i, j)
		fmt.Println(res)
	}
	{
		grid := [][]int{{1, 1}, {3, 4}}
		res := array.CountPaths(grid)
		fmt.Println(res)
	}
	{
		var t1 tree.TreeNode
		//var t2 tree.TreeNode
		//var t3 tree.TreeNode
		//var t4 tree.TreeNode
		//var t5 tree.TreeNode
		//t1.Left = &t2
		//t1.Right = &t3
		//t2.Left = &t4
		//t2.Right = &t5
		res := tree.DiameterOfBinaryTree(&t1)
		fmt.Println(res)
	}
	{
		var l1 list_queue.ListNode
		l1.Val = 1
		var l2 list_queue.ListNode
		l2.Val = 2
		var l3 list_queue.ListNode
		l3.Val = 3
		var l4 list_queue.ListNode
		l4.Val = 4
		var l5 list_queue.ListNode
		l5.Val = 5
		l1.Next = &l2
		l2.Next = &l3
		l3.Next = &l4
		l4.Next = &l5
		res := list_queue.OddEvenList(&l1)
		fmt.Println(res)
	}
	{
		strs := []string{"ab", "a"}
		res := string_issue.LongestCommonPrefix(strs)
		fmt.Println(res)
	}
	{
		n := 6
		delay := 2
		forget := 4
		res := number.PeopleAwareOfSecret2(n, delay, forget)
		fmt.Println(res)
	}
	{
		key := "the quick brown fox jumps over the lazy dog"
		message := "vkbs bs t suepuv"
		res := string_issue.DecodeMessage(key, message)
		fmt.Println(res)
	}
	{
		//"a##c"
		//"#a#c"
		//"bxj##tw"
		//"bxo#j##tw"
		//"bxj##tw"
		//"bxj###tw"
		s := "bxj##tw"
		t := "bxj###tw"
		res := string_issue.BackspaceCompare(s, t)
		fmt.Println(res)
	}
	{
		n := 3
		res := number.DistinctSequences(n)
		fmt.Println(res)
	}
	{
		nums := []int{3, 2, 4, 6}
		res := array.MaximumXOR(nums)
		fmt.Println(res)
	}
	{
		n := 3
		edges := [][]int{{0, 1}, {0, 2}, {1, 2}}
		res := diagram.CountPairs(n, edges)
		fmt.Println(res)
	}
	{
		ideas := []string{"coffee", "donuts", "time", "toffee"}
		res := string_issue.DistinctNames(ideas)
		fmt.Println(res)
	}
	{
		m := 3
		n := 5
		prices := [][]int{{1, 4, 2}, {2, 2, 7}, {2, 1, 3}}
		res := diagram.SellingWood(m, n, prices)
		fmt.Println(res)
	}
	{
		s := "0011110011"
		k := 522399436
		res := string_issue.LongestSubsequence(s, k)
		fmt.Println(res)
	}
	{
		num := 10
		k := 8
		res := number.MinimumNumbers(num, k)
		fmt.Println(res)
	}
	{
		nums := []int{2, 1, 4, 3, 5}
		var k int64 = 10
		res := array.CountSubarrays(nums, k)
		fmt.Println(res)
	}
	{
		cookies := []int{8, 15, 10, 20, 8}
		k := 2
		res := array.DistributeCookies(cookies, k)
		fmt.Println(res)
	}
	{
		grid := [][]int{{5, 3}, {4, 0}, {2, 1}}
		moveCost := [][]int{{9, 8}, {1, 5}, {10, 12}, {18, 6}, {2, 4}, {14, 3}}
		res := diagram.MinPathCost(grid, moveCost)
		fmt.Println(res)
	}
	{
		//"fool3e7bar"
		//"leet"
		//[["e","3"},{"t","7"},{"t","8"]]
		s := "fool3e7bar"
		sub := "leet"
		mappings := [][]byte{{'e', '3'}, {'t', '7'}, {'t', '8'}}
		res := string_issue.MatchReplacement(s, sub, mappings)
		fmt.Println(res)
	}
	{
		var nums []int = []int{4, 2, 4, 5, 6}
		res := array.MaximumUniqueSubarray(nums)
		fmt.Println(res)
	}
	{
		//[3,2,20,1,1,3]
		//10
		var nums []int = []int{3, 2, 20, 1, 1, 3}
		var x int = 10
		res := array.MinOperations1658(nums, x)
		fmt.Println(res)
	}
	{
		var n int = 4
		res := diagram.SolveNQueens(n)
		fmt.Println(res)
	}
	{
		//nums := []int{5,3,4,4,7,3,6,11,8,5,11}
		nums := []int{4, 5, 7, 7, 13}
		res := array.TotalSteps(nums)
		fmt.Println(res)
	}
	{
		sentence := "there are $1 $2 and 5$ candies in the shop"
		discount := 50
		res := string_issue.DiscountPrices(sentence, discount)
		fmt.Println(res)
	}
	{
		messages := []string{"How is leetcode for everyone", "Leetcode is useful for practice"}
		senders := []string{"Bob", "Charlie"}
		res := array.LargestWordCount(messages, senders)
		fmt.Println(res)
	}
	{
		s := "aaa"
		res := string_issue.AppealSum(s)
		fmt.Println(res)
	}
	{
		grid := [][]int{{0, 2, 0, 0, 1}, {0, 2, 0, 2, 2}, {0, 2, 0, 0, 0}, {0, 0, 2, 2, 0}, {0, 0, 0, 0, 0}}
		//grid := [][]int{{0,2,0,0,0,0,0},{0,0,0,2,2,1,0},{0,2,0,0,1,2,0},{0,0,2,2,2,0,2},{0,0,0,0,0,0,0}}
		res := diagram.MaximumMinutes(grid)
		fmt.Println(res)
	}
	{
		bottom := 2
		top := 9
		special := []int{4, 6}
		res := number.MaxConsecutive(bottom, top, special)
		fmt.Println(res)
	}
	{
		tiles := [][]int{{1, 5}, {10, 11}, {12, 18}, {20, 25}, {30, 32}}
		carpetLen := 10
		res := array.MaximumWhiteTiles(tiles, carpetLen)
		fmt.Println(res)
	}
	{
		pressedKeys := "22233"
		res := number.CountTexts(pressedKeys)
		fmt.Println(res)
	}
	{
		deadends := []string{"0201", "0101", "0102", "1212", "2002"}
		target := "0202"
		res := diagram.OpenLock(deadends, target)
		fmt.Println(res)
	}
	{
		graph := [][]int{{1, 2, 3}, {0}, {0}, {0}}
		res := diagram.ShortestPathLength(graph)
		fmt.Println(res)
	}
	{
		var x int = 3
		var y int = 5
		var z int = 4
		res := number.CanMeasureWater(x, y, z)
		fmt.Println(res)
	}
	{
		//["TimeMap","set","set","get","get","get","get","get"]
		//[[},{"love","high",10},{"love","low",20},{"love",5},{"love",10},{"love",15},{"love",20},{"love",25]]
		obj := list_queue.Constructor981()
		obj.Set("love", "high", 10)
		obj.Set("love", "low", 20)
		res := obj.Get("love", 5)
		res = obj.Get("love", 10)
		res = obj.Get("love", 15)
		res = obj.Get("love", 20)
		res = obj.Get("love", 25)
		fmt.Println(res)
	}
	{
		var amount int = 5
		var coins []int = []int{1, 2, 5}
		res := number.Change(amount, coins)
		fmt.Println(res)
	}
	{
		nums := []int{1, 4, 0, -1, -2, -3, -1, -2}
		res := array.Find132pattern(nums)
		fmt.Println(res)
	}
	{
		nums := []int{3, 5, 6, 7}
		target := 9
		res := array.NumSubseq(nums, target)
		fmt.Println(res)
	}
	{
		//["MyStack","push","push","top","pop","empty"]
		//[[},{1},{2},{},{},{]]
		var obj list_queue.MyStack = list_queue.Constructor225()
		obj.Push(1)
		obj.Push(2)
		res := obj.Top()
		res = obj.Pop()
		res2 := obj.Empty()
		fmt.Println(res, res2)
	}
	{
		nums1 := []int{1, 7, 5}
		nums2 := []int{2, 3, 5}
		res := array.MinAbsoluteSumDiff(nums1, nums2)
		fmt.Println(res)
	}
	{
		nums := []int{4, 2, 0}
		res := array.MinimumAverageDifference(nums)
		fmt.Println(res)
	}
	{
		n := 6
		connections := [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}
		res := diagram.MinReorder2(n, connections)
		fmt.Println(res)
	}
	{
		//var graph [][]int = [][]int{{1,2},{2,3},{5},{0},{5},{},{}}
		var graph [][]int = [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}
		res := diagram.EventualSafeNodes2(graph)
		fmt.Println(res)
	}
	{
		//var n int = 5
		//var connections [][]int = [][]int{{0, 1}, {0, 2}, {3, 4}, {2, 3}}
		var n int = 4
		var connections [][]int = [][]int{{0, 1}, {0, 2}, {1, 2}}
		res := diagram.MakeConnected2(n, connections)
		fmt.Println(res)
	}
	{
		//var data []int = []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}
		//var target int = 8
		var data []int = []int{1, 1, 2, 2, 2, 2}
		var target int = 5
		res := array.ThreeSumMulti2(data, target)
		fmt.Println(res)
	}
	{
		points := [][]int{{9, -25}, {-4, 1}, {-1, 5}, {-7, 7}}
		res := diagram.MaxPoints(points)
		fmt.Println(res)
	}
	{
		//var s = "cba"
		//var pairs [][int = [][]int{{0, 1}, {1, 2}}
		// s = "dcab", pairs = [[0,3},{1,2]]
		var s = "dcab"
		var pairs [][]int = [][]int{{0, 3}, {1, 2}}
		res := diagram.SmallestStringWithSwaps2(s, pairs)
		fmt.Println(res)
	}
	{
		//flowers := [][]int{{28,37},{23,33},{39,39},{49,50},{41,45},{14,47}}
		//persons := []int{44,41}
		//flowers := [][]int{{1,10},{3,3}}
		//persons := []int{3,3,2}
		flowers := [][]int{{50, 50}, {19, 27}, {40, 46}, {42, 48}, {22, 46}, {41, 50}, {11, 36}, {14, 29}}
		persons := []int{38}
		res := number.FullBloomFlowers(flowers, persons)
		fmt.Println(res)
	}
	{
		rectangles := [][]int{{1, 2}, {2, 3}, {2, 5}}
		points := [][]int{{2, 1}, {1, 4}}
		res := diagram.CountRectangles(rectangles, points)
		fmt.Println(res)
	}
	{
		left := 0
		right := 0
		res := number.RangeBitwiseAnd(left, right)
		fmt.Println(res)
	}
	{
		coins := []int{1, 2, 5}
		amount := 11
		res := number.CoinChange(coins, amount)
		fmt.Println(res)
	}
	{
		n := 8
		res := number.IntegerBreak(n)
		fmt.Println(res)
	}
	{
		//nums1 = [4,1,2], nums2 = [1,3,4,2].
		nums1 := []int{4, 1, 2}
		nums2 := []int{1, 3, 4, 2}
		res := array.NextGreaterElement(nums1, nums2)
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
		//t2.Left = &t4
		//t2.Right = &t5
		res := tree.PathSum2(&t1, 3)
		fmt.Println(res)
	}
	{
		obj := list_queue.Constructor705()
		obj.Add(1)
		obj.Add(2)
		res := obj.Contains(1)
		res = obj.Contains(3)
		obj.Add(2)
		res = obj.Contains(2)
		obj.Remove(2)
		res = obj.Contains(2)
		fmt.Println(res)
	}
	{
		parent := []int{-1, 0, 1}
		s := "aab"
		res := tree.LongestPath(parent, s)
		fmt.Println(res)
	}
	{
		n := 5
		k := 2
		res := list_queue.FindTheWinner(n, k)
		fmt.Println(res)
	}
	{
		grid := [][]int{{23, 17, 15, 3, 20}, {8, 1, 20, 27, 11}, {9, 4, 6, 2, 21}, {40, 9, 1, 10, 6}, {22, 7, 4, 5, 3}}
		res := diagram.MaxTrailingZeros(grid)
		fmt.Println(res)
	}
	{
		obj := number.Constructor2242()
		obj.Deposit([]int{0, 10, 0, 3, 0})
		res := obj.Withdraw(500)
		obj.Deposit([]int{0, 1, 0, 1, 1})
		res = obj.Withdraw(600)
		res = obj.Withdraw(550)
		fmt.Println(res)
	}
	{
		var t1 tree.TreeNode
		t1.Val = 1
		var t2 tree.TreeNode
		t2.Val = 2
		var t3 tree.TreeNode
		t3.Val = 3
		t1.Left = &t2
		t1.Right = &t3
		res := tree.IncreasingBST(&t1)
		fmt.Println(res)
	}
	{
		candidates := []int{2, 3, 6, 7}
		target := 7
		res := array.CombinationSum(candidates, target)
		fmt.Println(res)
	}
	{
		var t1 tree.TreeNode
		t1.Val = 1
		var t2 tree.TreeNode
		t2.Val = 2
		var t3 tree.TreeNode
		t3.Val = 3
		t2.Left = &t1
		t2.Right = &t3
		res := tree.ConvertBST(&t2)
		fmt.Println(res)
	}
	{
		var nums []int = []int{1, 2, 2, 3}
		res := array.SubsetsWithDup(nums)
		fmt.Println(res)
	}
	{
		nums := []int{1, 2, 3}
		res := array.Subsets(nums)
		print(res)
	}
	{
		//"9990"
		//"999000"
		res := number.Multiply("2", "3")
		fmt.Println(res)
	}
	{
		grid := [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}
		res := array.CountNegatives(grid)
		fmt.Println(res)
	}
	{
		var grid [][]int = [][]int{{0, 1}, {1, 0}}
		res := diagram.ShortestPathBinaryMatrix(grid)
		fmt.Println(res)
	}
	{
		isConnected := [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}
		res := diagram.FindCircleNum(isConnected)
		fmt.Println(res)
	}
	{
		target := 7
		nums := []int{2, 3, 1, 2, 4, 3}
		res := array.MinSubArrayLen3(target, nums)
		fmt.Println(res)
	}
	{
		//[57,44,92,28,66,60,37,33,52,38,29,76,8,75,22]
		//18
		var nums []int = []int{57, 44, 92, 28, 66, 60, 37, 33, 52, 38, 29, 76, 8, 75, 22}
		var k int = 18
		res := number.NumSubarrayProductLessThanK(nums, k)
		fmt.Println(res)
	}
	{
		flowers := []int{1, 3, 1, 1}
		var newFlowers int64 = 7
		target := 6
		full := 12
		partial := 1
		res := number.MaximumBeauty(flowers, newFlowers, target, full, partial)
		fmt.Println(res)
	}
	{
		expression := "247+38"
		res := number.MinimizeResult(expression)
		fmt.Println(res)
	}
	{
		var matrix [][]int = [][]int{{1, 4}, {2, 5}}
		var target int = 2
		res := array.SearchMatrix(matrix, target)
		fmt.Println(res)
	}
	{
		var nums []int = []int{9, -9, 4, 12, 12, 0, -14, -7, 10, -1, 11, -10, -3, 2, -9, 0, 8, -9, -5, -1, 10, 5, 13, -5, -9, -12, 9, -3, 10, 10, -10, 4, 8, 1, -7, -2, -14, -6, 6, 11, 8, -6, 9, 13, 11, 7, -10, -4, 14, 0, 3, 1, -10, -7, 3, -12, -3, -11, 0, -8, -15, 5, 3, 8, 13, 11, 13, -15, -9, 4, 3, 6, 5, -11, -4, -6, 4, 1, 5, -5, -13, -7, 11, -8, 2, -1, -12, 14, 3, 3, 13, -5, -14, -7, 11, 14, -11, 9, 6, -13, -9, -13, 1, 11, -9, 12, -10, 2, -1, 3, 12, -7, 3, 0, 0, 12, 6, 3, 3, -13, 14, 1, -3}
		res := number.ThreeSum(nums)
		fmt.Println(res)
	}
	{
		//[5,7,7,8,8,10]
		//8
		nums := []int{1, 2, 3}
		target := 2
		res := array.SearchRange(nums, target)
		fmt.Println(res)
	}
	{
		//[4,5,6,7,8,1,2,3]
		//8
		nums := []int{4, 5, 6, 7, 8, 1, 2, 3}
		target := 8
		//nums := []int{4,5,6,7,0,1,2}
		//target := 0
		//nums := []int{1,3}
		//target := 1
		res := array.Search33(nums, target)
		fmt.Println(res)
	}
	{
		var n uint32 = 43261596
		res := number.ReverseBits(n)
		fmt.Println(res)
	}
	{
		//keys := []byte{'a','b','c','d'}
		//values := []string{"ei","zf","ei","am"}
		//dictionary := []string{"abcd","acbd","adbc","badc","dacb","cadb","cbda","abad"}

		keys := []byte{'b'}
		values := []string{"ca"}
		dictionary := []string{"aaa", "cacbc", "bbaba", "bb"}
		obj := string_issue.Constructor2227(keys, values, dictionary)
		//["abcd"},{"eizfeiam"]
		res1 := obj.Encrypt("bbb")
		fmt.Println(res1)
		res2 := obj.Decrypt("cacaca")
		fmt.Println(res2)
	}
	{
		floor := []int{1, 2, 4, 3}
		var res int = array.MaxArea(floor)
		fmt.Println(res)
	}
	{
		candies := []int{1}
		var k int64 = 1
		res := number.MaximumCandies(candies, k)
		fmt.Println(res)
	}
	{
		var triangle [][]int = [][]int{{-1}, {-2, -3}}
		res := number.MinimumTotal3(triangle)
		fmt.Println(res)
	}
	{
		nums := []int{1, 2, 3}
		res := array.Permute(nums)
		fmt.Println(res)
	}
	{
		//var n int = 21
		//var k int = 17
		//var maxPts int = 10
		//6707
		//6047
		//9446
		var n int = 6707
		var k int = 6047
		var maxPts int = 9446
		res := number.New21Game(n, k, maxPts)
		fmt.Println(res)
	}
	{
		var t1 tree.Node
		t1.Val = 1
		var t2 tree.Node
		t2.Val = 2
		var t3 tree.Node
		t3.Val = 3
		var t4 tree.Node
		t4.Val = 4
		var t5 tree.Node
		t5.Val = 5
		var t6 tree.Node
		t6.Val = 6
		var t7 tree.Node
		t7.Val = 7
		t1.Left = &t2
		t1.Right = &t3
		t2.Left = &t4
		t2.Right = &t5
		t3.Left = &t6
		t3.Right = &t7
		res := tree.Connect(&t1)
		fmt.Println(res)
	}
	{
		//"mississippi"
		//"mis*is*ip*."

		//"mississippi"
		//"mis*is*p*."

		//"aaa"
		//"ab*a"

		//"aab"
		//"c*a*b"

		//"a"
		//".*..a*"
		s := "a"
		p := ".*..a*"
		res := string_issue.IsMatch(s, p)
		fmt.Println(res)
	}
	{
		piles := [][]int{{1, 100, 3}, {7, 8, 9}}
		k := 2
		res := number.MaxValueOfCoins(piles, k)
		fmt.Println(res)
	}
	{
		s := "leetcode"
		res := string_issue.FirstUniqChar(s)
		fmt.Println(res)
	}
	{
		//"ab"
		//"eidboaoo"
		var s1 string = "ab"
		var s2 string = "eidboaoo"
		res := string_issue.CheckInclusion(s1, s2)
		fmt.Println(res)
	}
	{
		input := "abcabcabc"
		res := string_issue.LengthOfLongestSubstring(input)
		fmt.Println(res)
	}
	{
		queries := []int{1, 2, 3, 4, 5, 90}
		intLength := 3
		res := number.KthPalindrome(queries, intLength)
		fmt.Println(res)
	}
	{
		nums := []int{1, 2, 3, 4, 5, 6, 7}
		k := 3
		array.Rotate189(nums, k)
		fmt.Println(nums)
	}
	{
		floor := "10110101"
		numCarpets := 2
		carpetLen := 2
		res := string_issue.MinimumWhiteTiles(floor, numCarpets, carpetLen)
		fmt.Println(res)
	}
	{
		var nums []int = []int{4}
		var target int = 2
		res := array.SearchInsert(nums, target)
		fmt.Println(res)
	}
	{
		nums := []int{5}
		target := 5
		res := array.Search(nums, target)
		fmt.Println(res)
	}
	{
		s := "tssaaaa"
		res := string_issue.MinMovesToMakePalindrome(s)
		fmt.Println(res)
	}
	{
		//[[99,7]]
		//85
		//95
		//tires = [[2,3},{3,4]], changeTime = 5, numLaps = 4
		//tires := [][]int{{99,7}}
		//changeTime := 85
		//numLaps := 95
		//[[1,10},{2,2},{3,4]]
		//6
		//5
		tires := [][]int{{1, 10}, {2, 2}, {3, 4}}
		changeTime := 6
		numLaps := 5
		res := number.MinimumFinishTime(tires, changeTime, numLaps)
		fmt.Println(res)
	}
	{
		numArrows := 9
		aliceArrows := []int{1, 1, 0, 1, 0, 0, 2, 1, 0, 1, 2, 0}
		res := number.MaximumBobPoints(numArrows, aliceArrows)
		fmt.Println(res)
	}
	{
		directions := "RLRSLL"
		res := diagram.CountCollisions(directions)
		fmt.Println(res)
	}
	{
		nums := []int{6, 6, 5, 5, 4, 1}
		res := array.CountHillValley(nums)
		fmt.Println(res)
	}
	{
		text := "b"
		pattern := "bc"
		res := string_issue.MaximumSubsequenceCount(text, pattern)
		fmt.Println(res)
	}
	{
		s := "1100101"
		res := string_issue.MinimumTime(s)
		fmt.Println(res)
	}
	{
		n := 2
		artifacts := [][]int{{0, 0, 0, 0}, {0, 1, 1, 1}}
		dig := [][]int{{0, 0}, {0, 1}}
		res := array.DigArtifacts(n, artifacts, dig)
		fmt.Println(res)
	}
	{
		nums := []int{2, 2, 2, 2, 2}
		key := 2
		k := 2
		res := array.FindKDistantIndices(nums, key, k)
		fmt.Println(res)
	}
	{
		words := []string{"ghnv", "uip", "tenv", "hvepx", "e", "ktc", "byjdt", "ulm", "cae", "ea"}
		res := string_issue.GroupStrings(words)
		fmt.Println(res)
	}
	{
		corridor := "SSPPSPS"
		res := number.NumberOfWays(corridor)
		fmt.Println(res)
	}
	{
		//[[0,0,1,1},{0,0,1,1]]
		//2
		//1
		grid := [][]int{{0, 0, 1, 1}, {0, 0, 1, 1}}
		stampHeight := 2
		stampWidth := 1
		res := array.PossibleToStamp(grid, stampHeight, stampWidth)
		fmt.Println(res)
	}
	{
		descriptions := [][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}}
		res := tree.CreateBinaryTree(descriptions)
		fmt.Println(res)
	}
	{
		fruits := [][]int{{0, 9}, {4, 1}, {5, 7}, {6, 2}, {7, 4}, {10, 9}}
		startPos := 5
		k := 4
		res := number.MaxTotalFruits(fruits, startPos, k)
		fmt.Println(res)
	}
	{
		nums := []int{5, 6}
		k := 6
		res := number.MinimalKSum(nums, k)
		fmt.Println(res)
	}
	{
		s := "K1:L2"
		res := string_issue.CellsInRange(s)
		fmt.Println(res)
	}
	{
		n := 8
		edgeList := [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}}
		res := diagram.GetAncestors(n, edgeList)
		fmt.Println(res)
	}
	{
		//[9,8,7,6,5,4,3,2,1,0]
		//[0,1,2,3,4,5,6,7,8,9]
		mapping := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
		nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		res := number.SortJumbled(mapping, nums)
		fmt.Println(res)
	}
	{
		bombs := [][]int{{1, 2, 3}, {2, 3, 1}, {3, 4, 2}, {4, 5, 3}, {5, 6, 4}}
		res := diagram.MaximumDetonation(bombs)
		fmt.Println(res)
	}
	{
		//[[1,2},{1,3},{2,1]]
		pairs := [][]int{{5, 13}, {10, 6}, {11, 3}, {15, 19}, {16, 19}, {1, 10}, {19, 11}, {4, 16}, {19, 9}, {5, 11},
			{5, 6}, {13, 5}, {13, 9}, {9, 15}, {11, 16}, {6, 9}, {9, 13}, {3, 1}, {16, 5}, {6, 5}}
		res := diagram.ValidArrangement(pairs)
		fmt.Println(res)
	}
	{
		nums := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
		k := 3
		res := array.GetAverages(nums, k)
		fmt.Println(res)
	}
	{
		grid := [][]int{{1, 1, 1, 1, 0}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {0, 1, 0, 0, 1}}
		res := diagram.CountPyramids(grid)
		fmt.Println(res)
	}
	{
		time := []int{1}
		totalTrips := 1
		res := number.MinimumTime(time, totalTrips)
		fmt.Println(res)
	}
	{
		items := [][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}
		queries := []int{1, 2, 3, 4, 5, 6}
		res := array.MaximumBeauty(items, queries)
		fmt.Println(res)
	}
	{
		tasks := []int{10, 15, 30}
		workers := []int{0, 10, 10, 10, 10}
		pills := 3
		strength := 10
		//tasks := []int{3,2,1}
		//workers := []int{0,3,3}
		//pills := 1
		//strength := 1
		res := number.MaxTaskAssign(tasks, workers, pills, strength)
		fmt.Println(res)
	}
	{
		values := []int{0, 32, 10, 43}
		edges := [][]int{{0, 1, 10}, {1, 2, 15}, {0, 3, 10}}
		maxTime := 49
		res := diagram.MaximalPathQuality(values, edges, maxTime)
		fmt.Println(res)
	}
	{
		obj := diagram.Constructor2069(6, 3)

		obj.Step(28)
		//obj.Step(19)
		//obj.Step(8)
		//obj.Step(36)
		pos := obj.GetPos()
		dir := obj.GetDir()

		obj.Step(4)
		pos = obj.GetPos()
		dir = obj.GetDir()
		fmt.Println(pos)
		fmt.Println(dir)
	}
	{
		//"aabkljajfjzxcvlfjlsajfaljflbvcxfjsljflsajfldsfsdfzvcxzabab"
		//3
		s := "aabklazxcvllsaallbvcxsllsaldssdzvcxzabab"
		repeatLimit := 3
		res := string_issue.RepeatLimitedString(s, repeatLimit)
		fmt.Println(res)
	}
	{
		var finalSum int64 = 28
		res := number.MaximumEvenSplit(finalSum)
		fmt.Println(res)
	}
	{
		nums := []int{3, 1, 2, 2, 2, 1, 3}
		var k int = 2
		res := array.CountPairs(nums, k)
		fmt.Println(res)
	}
	{
		beans := []int{2, 10, 3, 2}
		res := array.MinimumRemoval(beans)
		fmt.Println(res)
	}
	{
		edges := [][]int{{0, 1}, {1, 2}}
		patience := []int{0, 2, 1}
		res := diagram.NetworkBecomesIdle(edges, patience)
		fmt.Println(res)
	}
	{
		//s = "leetcode", k = 4, letter = "e", repetition = 2
		s := "facfffkfnffoppfffzfz"
		k := 9
		var letter byte = 'f'
		repetition := 9
		res := string_issue.SmallestSubsequence2(s, k, letter, repetition)
		fmt.Println(res)
	}
	{
		stones := []int{2}
		res := number.StoneGameIX(stones)
		fmt.Println(res)
	}
	{
		nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 30827, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		k := 0
		res := array.WaysToPartition(nums, k)
		fmt.Println(res)
	}
	{
		//"1+2*3+4" [13,21,11,15]
		//s = "6+0*1", answers = [12,9,6,4,8,6]
		//s = "3+5*2", answers = [13,0,10,13,13,16,16]
		s := "4+1+3*5+1*3+1+1*1+6*1+0*3+5*4+4"
		answers := []int{13, 21, 11, 15}
		res := number.ScoreOfStudents(s, answers)
		fmt.Println(res)
	}
	{
		var num int64 = -7605
		res := number.SmallestNumber(num)
		fmt.Println(res)
	}
	{
		var nums []int = []int{4, 1, 2, 3}
		res := array.SortEvenOdd(nums)
		fmt.Println(res)
	}
	{
		nums := []int{1, 10, 100, 1000}
		res := array.MinOperations2009(nums)
		fmt.Println(res)
	}
	{
		nextVisit := []int{0, 0, 2}
		res := number.FirstDayBeenInAllRooms(nextVisit)
		fmt.Println(res)
	}
	{
		//[2,3,3,4,4,4,5,6,7,10]
		//12
		tasks := []int{2, 3, 3, 4, 4, 4, 5, 6, 7, 10}
		sessionTime := 12
		res := array.MinSessions(tasks, sessionTime)
		fmt.Println(res)
	}
	{
		nums := []int{0, 1, 2, 2}
		res := array.CountSpecialSubsequences(nums)
		fmt.Println(res)
	}
	{
		statements := [][]int{{2, 1, 2}, {1, 2, 2}, {2, 0, 2}}
		res := number.MaximumGood(statements)
		fmt.Println(res)
	}
	{
		grid := [][]int{{1, 2, 0, 1}, {1, 3, 3, 1}, {0, 2, 5, 1}}
		pricing := []int{2, 3}
		start := []int{2, 3}
		k := 2
		res := array.HighestRankedKItems(grid, pricing, start, k)
		fmt.Println(res)
	}
	{
		cost := []int{6}
		res := number.MinimumCost(cost)
		fmt.Println(res)
	}
	{
		heights := []int{5, 1, 2, 3, 10}
		res := array.CanSeePersonsCount(heights)
		fmt.Println(res)
	}
	{
		segments := [][]int{{4, 16, 12}, {9, 10, 15}, {18, 19, 13}, {3, 13, 20}, {12, 16, 3}, {2, 10, 10}, {3, 11, 4}, {13, 16, 6}}
		res := array.SplitPainting(segments)
		fmt.Println(res)
	}
	{
		m := 5
		n := 5
		res := array.ColorTheGrid(m, n)
		fmt.Println(res)
	}
	{
		n := 2
		batteries := []int{3, 3, 3}
		res := number.MaxRunTime(n, batteries)
		fmt.Println(res)
	}
	{
		s := "aabca"
		res := string_issue.CountPalindromicSubsequence(s)
		fmt.Println(res)
	}
	{
		var points [][]int = [][]int{{1, 2, 3}, {1, 5, 1}, {3, 1, 1}}
		res := array.MaxPoints(points)
		fmt.Println(res)
	}
	{
		n := 5
		//paths := [][]int{{0,1,2,3,4}, {2,3,4}, {4,0,1,2,3}}
		//paths := [][]int{{0,1,2,3,4}, {4,3,2,1,0}}
		paths := [][]int{{0, 1, 0, 1, 0, 1, 0, 1, 0}, {0, 1, 3, 0, 1, 4, 0, 1, 0}}
		res := diagram.LongestCommonSubpath(n, paths)
		fmt.Println(res)
	}
	{
		var n int64 = 100
		res := number.CountGoodNumbers(n)
		fmt.Println(res)
	}
	{
		//[1,1,2,3]
		//[1,1,1,1]
		dist := []int{1, 1, 2, 3}
		speed := []int{1, 1, 1, 1}
		res := array.EliminateMaximum(dist, speed)
		fmt.Println(res)
	}
	{
		plantTime := []int{1, 4, 3}
		growTime := []int{2, 3, 1}
		res := number.EarliestFullBloom(plantTime, growTime)
		fmt.Println(res)
	}
	{
		//["mox","bj","rsy","jqsh"]
		//["trk","vjb","jkr"]
		//["lhum","bim","cyjhm","h","a","ljxd","run","zrl","dmf","y"]
		//["da","lzkr","uzc","jdxlz","yt","emdf"]
		startWords := []string{"g", "vf", "ylpuk", "nyf", "gdj", "j", "fyqzg", "sizec"}
		targetWords := []string{"r", "am", "jg", "umhjo", "fov", "lujy", "b", "uz", "y"}
		res := string_issue.WordCount(startWords, targetWords)
		fmt.Println(res)
	}
	{
		nums := []int{1, 1, 1, 0, 1}
		res := array.MinSwaps(nums)
		fmt.Println(res)
	}
	{
		words := []string{"ll", "lb", "bb", "bx", "xx", "lx", "xx", "lx", "ll", "xb", "bx", "lb", "bb", "lb", "bl", "bb", "bx", "xl", "lb", "xx"}
		res := string_issue.LongestPalindrome2131(words)
		fmt.Println(res)
	}
	{
		var l1 list_queue.ListNode
		l1.Val = 5
		var l2 list_queue.ListNode
		l2.Val = 4
		var l3 list_queue.ListNode
		l3.Val = 2
		var l4 list_queue.ListNode
		l4.Val = 1
		l1.Next = &l2
		l2.Next = &l3
		l3.Next = &l4
		res := list_queue.PairSum(&l1)
		fmt.Println(res)
	}
	{
		nums := []int{1, 2, 10, 5, 7}
		res := array.CanBeIncreasing(nums)
		fmt.Println(res)
	}
	{
		n := 11
		firstPlayer := 2
		secondPlayer := 4
		//n := 5
		//firstPlayer := 1
		//secondPlayer := 5
		res := number.EarliestAndLatest(n, firstPlayer, secondPlayer)
		fmt.Println(res)
	}
	{
		packages := []int{2, 3, 5}
		boxes := [][]int{{4, 8}, {2, 8}}
		res := number.MinWastedSpace(packages, boxes)
		fmt.Println(res)
	}
	{
		nums := []int{1, 1, 2, 2, 3}
		res := array.ReductionOperations(nums)
		fmt.Println(res)
	}
	{
		dist := []int{7, 3, 5, 5}
		speed := 2
		hoursBefore := 10
		res := number.MinSkips(dist, speed, hoursBefore)
		fmt.Println(res)
	}
	{
		n := "13"
		x := 2
		res := number.MaxValue(n, x)
		fmt.Println(res)
	}
	{
		nums1 := []int{1, 0, 3}
		nums2 := []int{5, 3, 4}
		res := number.MinimumXORSum(nums1, nums2)
		fmt.Println(res)
	}
	{
		n := 5
		k := 3
		res := number.RearrangeSticks(n, k)
		fmt.Println(res)
	}
	{
		nums := []int{1, 3}
		res := array.SubsetXORSum(nums)
		fmt.Println(res)
	}
	{
		nums := []int{2, 5, 9}
		res := number.SumOfFlooredPairs(nums)
		fmt.Println(res)
	}
	{
		//"rrrde"
		//[[4,2},{3,4},{0,3},{1,0},{2,1]]
		//"iiiiii"
		//[[0,1},{1,2},{2,3},{3,4},{4,5]]
		//"abaca"
		//[[0,1},{0,2},{2,3},{3,4]]

		//"ab"
		//[[1,0},{1,1]]
		colors := "ab"
		edges := [][]int{{1, 0}, {1, 1}}
		res := diagram.LargestPathValue(colors, edges)
		fmt.Println(res)
	}
	{
		s := "())(()(()(())()())(())((())(()())((())))))(((((((())(()))))("
		locked := "100011110110011011010111100111011101111110000101001101001111"
		res := string_issue.CanBeValid(s, locked)
		fmt.Println(res)
	}
	{
		recipes := []string{"bread", "sandwich"}
		ingredients := [][]string{{"yeast", "flour"}, {"bread", "meat"}}
		supplies := []string{"yeast", "flour", "meat"}
		res := diagram.FindAllRecipes(recipes, ingredients, supplies)
		fmt.Println(res)
	}
	{
		word := "aeiou"
		res := string_issue.LongestBeautifulSubstring(word)
		fmt.Println(res)
	}
	{
		grid := [][]int{{3, 4, 5, 1, 3}, {3, 3, 4, 2, 3}, {20, 30, 200, 40, 10}, {1, 5, 5, 4, 1}, {4, 3, 2, 2, 5}}
		res := array.GetBiggestThree(grid)
		fmt.Println(res)
	}
	{
		obstacles := []int{0, 1, 2, 3, 0}
		res := diagram.MinSideJumps2(obstacles)
		fmt.Println(res)
	}
	{
		//[12,6,12,6,14,2,13,17,3,8,11,7,4,11,18,8,8,3]
		//1
		//[2,2,2,2,2,1,1,4,4,3,3,3,3,3]
		//1
		arr := []int{2, 2, 2, 2, 2, 1, 1, 4, 4, 3, 3, 3, 3, 3}
		k := 1
		res := array.KIncreasing(arr, k)
		fmt.Println(res)
	}
	{
		nums := []int{6, 9}
		res := number.CountDifferentSubsequenceGCDs(nums)
		fmt.Println(res)
	}
	{
		security := []int{5, 3, 3, 3, 5, 6, 2}
		time := 2
		res := number.GoodDaysToRobBank(security, time)
		fmt.Println(res)
	}
	{
		//"Game is ON"
		//"Game are ON"
		//"URhnaPlQqSx h"
		//"URhnaPlQqSx RpASX h"
		sentence1 := "Game is ON"
		sentence2 := "Game are ON"
		//sentence1 := "URhnaPlQqSx h"
		//sentence2 := "URhnaPlQqSx RpASX h"
		res := string_issue.AreSentencesSimilar(sentence1, sentence2)
		fmt.Println(res)
	}
	{
		n := 2
		res := array.ReinitializePermutation(n)
		fmt.Println(res)
	}
	{
		//9
		//5
		//24
		n := 9
		index := 5
		maxSum := 24
		res := array.MaxValue(n, index, maxSum)
		fmt.Println(res)
	}
	{
		nums := []int{1, 2, 3, 4, 5, 6}
		res := number.MaxScore1799(nums)
		fmt.Println(res)
	}
	{
		nums := []int{1, 4, 3, 7, 4, 5}
		k := 3
		res := array.MaximumScore(nums, k)
		fmt.Println(res)
	}
	{
		coins := []int{1, 3}
		res := number.GetMaximumConsecutive(coins)
		fmt.Println(res)
	}
	{
		var n int = 12
		res := number.CheckPowersOfThree(n)
		fmt.Println(res)
	}
	{
		cars := [][]int{{1, 2}, {2, 1}, {4, 3}, {7, 2}}
		res := number.GetCollisionTimes(cars)
		fmt.Println(res)
	}
	{
		word1 := "cacb"
		word2 := "cbba"
		res := string_issue.LongestPalindrome1771(word1, word2)
		fmt.Println(res)
	}
	{
		//[2,3]
		//[4,5,100]
		//18
		baseCosts := []int{2, 3}
		toppingCosts := []int{4, 5, 100}
		target := 18
		res := number.ClosestCost(baseCosts, toppingCosts, target)
		fmt.Println(res)
	}
	{
		nums := []int{5, 6, 10, 2, 3, 6, 15}
		edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}}
		res := tree.GetCoprimes(nums, edges)
		fmt.Println(res)
	}
	{
		ch := make(chan int, 10)
		for i := 0; i < 10; i++ {
			select {
			case x := <-ch:
				fmt.Println(x)
			case ch <- i:
				fmt.Println("--", i)
			}
		}
	}
	{
		boxes := "110"
		res := number.MinOperations2(boxes)
		fmt.Println(res)
	}
	{
		var s string = "YazaAay"
		res := string_issue.LongestNiceSubstring(s)
		fmt.Println(res)
	}
	{
		nums := []int{7, -9, 15, -2}
		goal := -5
		res := array.MinAbsDifference(nums, goal)
		fmt.Println(res)
	}
	{
		s := "abcbdd"
		res := string_issue.CheckPartitioning(s)
		fmt.Println(res)
	}
	{
		var s string = "aabbbaa"
		res := string_issue.MinimumLength(s)
		fmt.Println(res)
	}
	{
		var nums []int = []int{2, -5, 1, -4, 3, -2}
		res := number.MaxAbsoluteSum(nums)
		fmt.Println(res)
	}
	{
		var groups [][]int = [][]int{{1, -1, -1}, {3, -2, 0}}
		var nums []int = []int{1, -1, 0, 1, -1, -1, 3, -2, 0, 1}
		//var groups [][]int = [][]int{{6636698, 4693069, -2178984, -2253405, -2732131, 8550889, -2324014, -2561263}, {-8973571, -9146179, 7704305, -1063430, -6569826}, {2791091}, {-9691134, 651171, 9835817, 4163881, 4944714, 8166788, -9025553, -9250995, 1599781}}
		//var nums []int = []int{8550889, -2178984, 6636698, 4693069, -2178984, -2253405, -2732131, 8550889, -2324014, -2561263, -2324014, 8550889, -8973571, -9146179, 7704305, -1063430, -6569826, 2791091, -9691134, 651171, 9835817, 4163881, 4944714, 8166788, -9025553, -9250995, 1599781}
		res := array.CanChoose(groups, nums)
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
		var t6 tree.TreeNode
		t6.Val = 6
		t5.Left = &t1
		t5.Right = &t2
		t1.Left = &t3
		t2.Left = &t6
		t2.Right = &t4
		start := 3
		dest := 6
		res := tree.GetDirections(&t5, start, dest)
		fmt.Println(res)
	}
	{
		nums := []int{7, 17}
		maxOperations := 2
		res := number.MinimumSize(nums, maxOperations)
		fmt.Println(res)
	}
	{
		//[7,11,5,3,8]
		//[[2,2,6},{4,2,4},{2,13,1000000000]]
		candiesCount := []int{7, 11, 5, 3, 8}
		queries := [][]int{{2, 2, 6}, {4, 2, 4}, {2, 13, 1000000000}}
		res := number.CanEat(candiesCount, queries)
		fmt.Println(res)
	}
	{
		a := "dee"
		b := "a"
		res := string_issue.MinCharacters(a, b)
		fmt.Println(res)
	}
	{
		encoded := []int{6, 5, 4, 6}
		res := number.Decode2(encoded)
		fmt.Println(res)
	}
	{
		n := 2
		languages := [][]int{{1}, {2}, {1, 2}}
		friendships := [][]int{{1, 2}, {1, 3}, {2, 3}}
		res := diagram.MinimumTeachings(n, languages, friendships)
		fmt.Println(res)
	}
	{
		jobs := []int{1, 2, 4, 7, 8}
		k := 2
		res := number.MinimumTimeRequired(jobs, k)
		fmt.Println(res)
	}
	{
		input := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
		res := array.MaximalRectangle(input)
		fmt.Println(res)
	}
	{
		//[5,1,2,4,3]
		//[1,5,4,2,3]
		//[[0,4},{4,2},{1,3},{1,4]]
		var source []int = []int{5, 1, 2, 4, 3}
		var target []int = []int{1, 5, 4, 2, 3}
		var allowedSwaps [][]int = [][]int{{0, 4}, {4, 2}, {1, 3}, {1, 4}}
		res := diagram.MinimumHammingDistance(source, target, allowedSwaps)
		fmt.Println(res)
	}
	{
		target := []int{6, 4, 8, 1, 3, 2}
		arr := []int{4, 7, 6, 2, 3, 8, 6, 1}
		res := array.MinOperations1713(target, arr)
		fmt.Println(res)
	}
	{
		input := []int{0, 1, 0, 3, 2, 3}
		res := array.LengthOfLIS(input)
		fmt.Println(res)
	}
	{
		nums := []int{8892, 2631, 7212, 1188, 6580, 1690, 5950, 7425, 8787, 4361, 9849, 4063, 9496, 9140, 9986, 1058, 2734, 6961, 8855, 2567, 7683, 4770, 40, 850, 72, 2285, 9328, 6794, 8632, 9163, 3928, 6962, 6545, 6920, 926, 8885, 1570, 4454, 6876, 7447, 8264, 3123, 2980, 7276, 470, 8736, 3153, 3924, 3129, 7136, 1739, 1354, 661, 1309, 6231, 9890, 58, 4623, 3555, 3100, 3437}
		res := array.WaysToSplit(nums)
		fmt.Println(res)
	}
	{
		//[5,2,4,6,6,3]
		//[[12,4},{8,1},{6,3]]
		//nums := []int{0,1,2,3,4}
		//queries := [][]int{{1,3},{5,6},{3,1}}
		nums := []int{5, 2, 4, 6, 6, 3}
		queries := [][]int{{12, 4}, {8, 1}, {6, 3}}
		//nums := []int{874267,1416960,0,24133703,3996339,213319831,140874167,11289135,0,0,0,203530386,13542138,641645,0,33097768,497897924,25853758,733984105,1169259}
		//queries := [][]int{{902696869,17211}}
		res := number.MaximizeXor(nums, queries)
		fmt.Println(res)
	}
	{
		nums := []int{3, 10, 5, 25, 2, 8}
		res := number.FindMaximumXOR(nums)
		fmt.Println(res)
	}
	{
		//[2,2]
		//[1,4]
		//[1,0,21,18]
		//[9,4,1,2,13,18,14,23]
		startPos := []int{2, 2}
		homePos := []int{1, 4}
		rowCosts := []int{1, 0, 21, 18}
		colCosts := []int{9, 4, 1, 2, 13, 18, 14, 23}
		res := array.MinCost(startPos, homePos, rowCosts, colCosts)
		fmt.Println(res)
	}
	{
		n := 3
		edgeList := [][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}}
		queries := [][]int{{0, 1, 2}, {0, 2, 5}}
		res := diagram.DistanceLimitedPathsExist(n, edgeList, queries)
		fmt.Println(res)
	}
	{
		var nums []int = []int{10, -5, -2, 4, 0, 3}
		var k int = 3
		res := array.MaxResult(nums, k)
		fmt.Println(res)
	}
	{
		cuboids := [][]int{{7, 11, 17}, {7, 17, 11}, {11, 7, 17}, {11, 17, 7}, {17, 7, 11}, {17, 11, 7}}
		res := array.MaxHeight(cuboids)
		fmt.Println(res)
	}
	{
		number := "1234567"
		res := string_issue.ReformatNumber(number)
		fmt.Println(res)
	}
	{
		stones := []int{7, 90, 5, 1, 100, 10, 10, 2}
		res := array.StoneGameVII(stones)
		fmt.Println(res)
	}
	{
		//[9,8,3,8]
		//[10,6,9,5]
		aliceValues := []int{9, 8, 3, 8}
		bobValues := []int{10, 6, 9, 5}
		res := number.StoneGameVI(aliceValues, bobValues)
		fmt.Println(res)
	}
	{
		//[7,3,16,15,1,13,1,2,14,5,3,10,6,2,7,15]
		//8
		nums := []int{7, 3, 16, 15, 1, 13, 1, 2, 14, 5, 3, 10, 6, 2, 7, 15}
		k := 8
		res := array.MinimumIncompatibility(nums, k)
		fmt.Println(res)
	}
	{
		//[28,50,76,80,64,30,32,84,53,8]
		//84
		//[1,2,4,3]
		nums := []int{1, 2, 4, 3}
		limit := 4
		res := array.MinMoves(nums, limit)
		fmt.Println(res)
	}
	{
		//[71,18,52,29,55,73,24,42,66,8,80,2]
		//3
		nums := []int{71, 18, 52, 29, 55, 73, 24, 42, 66, 8, 80, 2}
		k := 3
		res := array.MostCompetitive(nums, k)
		fmt.Println(res)
	}
	{
		//[100,92,89,77,74,66,64,66,64]
		nums := []int{100, 92, 89, 77, 74, 66, 64, 66, 64}
		res := array.MinimumMountainRemovals(nums)
		fmt.Println(res)
	}
	{
		tasks := [][]int{{1, 3}, {2, 4}, {10, 11}, {10, 12}, {8, 9}}
		res := array.MinimumEffort(tasks)
		fmt.Println(res)
	}
	{
		sequence := "bbaabaabaabaaabbbaaabaababbabbaabbabbbabbabbabbbbabaababbbaaaababbbabbaababababbabbbaabbaabbbb"
		word := "bba"
		res := string_issue.MaxRepeating(sequence, word)
		fmt.Println(res)
	}
	{
		var nums []int = []int{2, 1, 6, 4}
		res := array.WaysToMakeFair(nums)
		fmt.Println(res)
	}
	{
		var n int = 5
		var k int = 73
		res := string_issue.GetSmallestString(n, k)
		fmt.Print(res)
	}
	{
		nums := []int{1, 1, 2, 2}
		quantity := []int{2, 2}
		res := number.CanDistribute(nums, quantity)
		fmt.Println(res)
	}
	{
		k := 3
		n := 7
		res := number.KMirror(k, n)
		fmt.Println(res)
	}
	{
		s := "aababbab"
		res := string_issue.MinimumDeletions(s)
		fmt.Println(res)
	}
	{
		inventory := []int{2, 8, 4, 10, 6}
		orders := 20
		//inventory := []int{1000000000}
		//orders := 1000000000
		res := array.MaxProfit1648(inventory, orders)
		fmt.Println(res)
	}
	{
		s := "bbcebab"
		res := string_issue.MinDeletions(s)
		fmt.Println(res)
	}
	{
		words := []string{"abab", "baba", "abba", "baab"}
		target := "abba"
		res := string_issue.NumWays1639(words, target)
		fmt.Println(res)
	}
	{
		//var heights [][]int = [][]int{{8,3,2,5,2,10,7,1,8,9},{1,4,9,1,10,2,4,10,3,5},
		//	{4,10,10,3,6,1,3,9,8,8},{4,4,6,10,10,10,2,10,8,8},{9,10,2,4,1,2,2,6,5,7},
		//	{2,9,2,6,1,4,7,6,10,9},{8,8,2,10,8,2,3,9,5,3},{2,10,9,3,5,1,7,4,5,6},
		//	{2,3,9,2,5,10,2,7,1,8},{9,10,4,10,7,4,9,3,1,6}}
		var heights [][]int = [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}
		res := diagram.MinimumEffortPath(heights)
		fmt.Println(res)
	}
	{
		s := "43987654"
		a := 7
		b := 3
		res := string_issue.FindLexSmallestString(s, a, b)
		fmt.Println(res)
	}
	{
		//[[1,2,5},{2,1,7},{3,1,9]]
		//2
		//[[30,34,31},{10,44,24},{14,28,23},{50,38,1},{40,13,6},{16,27,9},{2,22,23},{1,6,41},{34,22,40},{40,10,11]]
		//20
		towers := [][]int{{30, 34, 31}, {10, 44, 24}, {14, 28, 23}, {50, 38, 1}, {40, 13, 6}, {16, 27, 9}, {2, 22, 23}, {1, 6, 41}, {34, 22, 40}, {40, 10, 11}}
		radius := 20
		res := diagram.BestCoordinate(towers, radius)
		fmt.Println(res)
	}
	{
		n := 4
		edges := [][]int{{1, 2}, {2, 3}, {2, 4}}
		res := diagram.CountSubgraphsForEachDiameter(n, edges)
		fmt.Println(res)
	}
	{
		//"pvhmupgqeltozftlmfjjde"
		//"yjgpzbezspnnpszebzmhvp"
		a := "x"
		b := "y"
		res := string_issue.CheckPalindromeFormation(a, b)
		fmt.Println(res)
	}
	{
		//["daniel","daniel","daniel","luis","luis","luis","luis"]
		//["10:00","10:40","11:00","09:00","11:00","13:00","15:00"]
		keyName := []string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"}
		keyTime := []string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"}
		res := number.AlertNames(keyName, keyTime)
		fmt.Println(res)
	}
	{
		//输入：n = 3, restrictions = [[0,1]], requests = [[0,2},{2,1]]
		//输出：[true,false]
		n := 5
		restrictions := [][]int{{0, 1}, {1, 2}, {2, 3}}
		requests := [][]int{{0, 4}, {1, 2}, {3, 1}, {3, 4}}
		res := diagram.FriendRequests(n, restrictions, requests)
		fmt.Println(res)
	}
	{
		n := 5
		requests := [][]int{{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4}}
		//n := 3
		//requests := [][]int{{0,0},{1,2},{2,1}}
		//n := 3
		//requests := [][]int{{1,1}}
		res := diagram.MaximumRequests(n, requests)
		fmt.Println(res)
	}
	{
		customers := []int{10, 10, 6, 4, 7}
		boardingCost := 3
		runningCost := 8
		res := number.MinOperationsMaxProfit(customers, boardingCost, runningCost)
		fmt.Println(res)
	}
	{
		//cost := [][]int{{1, 3, 5}, {4, 1, 1}, {1, 5, 3}}
		cost := [][]int{{15, 96}, {36, 2}}
		res := diagram.ConnectTwoGroups(cost)
		fmt.Println(res)
	}
	{
		//var nums []int = []int{12,7,16,4,4,13,13,8}
		//var p int = 61
		var nums []int = []int{3, 1, 4, 2}
		var p int = 6
		res := number.MinSubarray(nums, p)
		fmt.Println(res)
	}
	{
		var nums []int = []int{1, 2, 3, 4, 5}
		var requests [][]int = [][]int{{1, 3}, {0, 1}}
		res := array.MaxSumRangeQuery(nums, requests)
		fmt.Println(res)
	}
	{
		points := [][]int{{0, 0}, {1, 1}, {1, 0}, {-1, 1}}
		res := diagram.MinCostConnectPoints(points)
		fmt.Println(res)
	}
	{
		locations := []int{1, 2, 3}
		start := 0
		finish := 2
		fuel := 40
		res := diagram.CountRoutes(locations, start, finish, fuel)
		fmt.Println(res)
	}
	{
		grid := [][]int{{1, 0, 0}, {1, 1, 0}, {1, 0, 0}}
		res := diagram.MinDays(grid)
		fmt.Println(res)
	}
	{
		//var nums []int = []int{5,-20,-20,-39,-5,0,0,0,36,-32,0,-7,-10,-7,21,20,-12,-34,26,2}
		var nums []int = []int{1, 2, 3, 5, -6, 4, 0, 10}
		res := array.GetMaxLen(nums)
		fmt.Println(res)
	}
	{
		stoneValue := []int{6, 2, 3, 4, 5, 5}
		//stoneValue := []int{1,2,3}
		res := array.StoneGameV(stoneValue)
		fmt.Println(res)
	}
	{
		//var arr []int = []int{6,3,5,5,5,5,1,5,6,2,5,1,2,5,3,5,1,3,5,5,6,4,1,2}
		//var m int = 1
		//var k int = 5
		var arr []int = []int{1, 2, 1, 2, 1, 1, 1, 3}
		m := 2
		k := 2
		res := diagram.ContainsPattern(arr, m, k)
		fmt.Println(res)
	}
	{
		//[3,5,1,2,4]
		//1
		arr := []int{3, 2, 5, 6, 10, 8, 9, 4, 1, 7}
		m := 3
		res := number.FindLatestStep(arr, m)
		fmt.Println(res)
	}
	{
		n := 1
		quantities := []int{100000}
		res := number.MinimizedMaximum(n, quantities)
		fmt.Println(res)
	}
	{
		var n int = 4
		var rounds []int = []int{1, 3, 1, 2}
		res := diagram.MostVisited(n, rounds)
		fmt.Println(res)
	}
	{
		n := 7
		cuts := []int{1, 3, 4, 5}
		res := number.MinCost2(n, cuts)
		fmt.Println(res)
	}
	{
		//[["a","a","c","a"},{"a","b","d","a"},{"a","b","b","a"},{"a","a","a","a"]]
		//grid := [][]byte{{'a','a','a','a'},{'a','b','b','a'},{'a','b','b','a'},{'a','a','a','a'}}
		grid := [][]byte{{'a', 'a', 'a', 'a'}, {'a', 'b', 'd', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'a', 'a', 'a'}}
		res := array.ContainsCycle(grid)
		fmt.Println(res)
	}
	{
		var n int = 3
		res := array.MinOperations(n)
		fmt.Println(res)
	}
	{
		//var n int = 20
		//var k int = 1048575
		var n int = 3
		var k int = 6
		res := string_issue.FindKthBit(n, k)
		fmt.Println(res)
	}
	{
		s := "3242415"
		res := string_issue.LongestAwesome(s)
		fmt.Println(res)
	}
	{
		//"dbnbhyilnxgmsunpjtdpsgcneaulepufbydspqpofimcffagbmgeratthxblehwpobhwzjoilfyagqpgvqwocjzqxlbtxdtytjonjuxbrwkphaoordcfmmellnxfujeslejzyubfeyhztqzuyxdntzwuggnwlfsrvfmhxpgmdqdoczmurjsuqwhiavcmgoqketlrrmqgmutymrnxmwdpdtgrgdpndbfvpohpecqxslzodypbkillxtqpfjpatoybfwxqlcxxvhsnidqwcdiiupumrvzrvticzpertybvsuqgmefqblxhegvogydhusbstqurrzgaevdftwmgdsogjjgydxeubtxecissauovmefjyionhgvsbljuzzxkgntxodtgnhijcvfkrwwqrbtsxdidudbozmusicextuuydysjwejlzdtckphwoichauncmegkalmaywzdfjebuofrwntrmlmrnaqfttyymbuyrvsdskfynxlnvcyymwhmblnvdvvvhmozlxyujnqczamkrwmietxajcobaaqoymwrjxylusqlnuquduzsbgailaorvqknlxrpwvcvndovjmjtghliguixskpdkqpvorqosspmjnljotetrfezoffdisjxytaibzczcyiiudjkdubkgeypuveqjpnptekzqzpprcqaeimenfichpkhfqcuqwygdkalbftvdlpuihaxhwqcdsnjfboicnouozcykliknybgbsckedgyxevnhgriiqwnukiqgkgycwvjofswmvfgwzucnwdappdbteniuonjqyemwzrynbbgsaovuawcmufrzyrmhipbuyktvghkrqpuxxdvsuqxbrordpjfnfqkfwitywyuokbdesaxgvdlpnfabsuajouygnkuvddkyzeokpxhdbdxtsrwfofllqzraqayvigkgczgtdufsdxcjumlcvlkzzemyhdausghtayxomqmqnpznwplmmexrqj...
		//"lkftipnsvhcfgvagwdoxyuzlanrooubeeufvkydtxukpbmpmhqrthmyvgevkofhnktmcmnbrvdsmarhdlqworlklmuooetmkagpkwpmjzydefnbphedzfpezisveqohgnivgppmogxqqqtrjcjdcfqauxhvlmifhuhxycabcluswfwmuksqiyewnxbslsszymnyxolkeglbpaljpmbwuhndeflmsysxtobbcczuasaqedgpwzjhxwpmwxlaeaxckbzjqipfxmbbjzszpnjcstbdfovuuotsbmtthjvssndargbzgzneunmlqpnkcncmsrkzagwengoodyemvmsmtkgaxkythdaiwejdrsesycksexfnurfpgndukgdnidhaplwsbskqlpfpeljwkxqrowsjefynxhavmrwooarqgtywaqkjptenexzoctmjgbmyfkrmlepopnjcqaihsxsjfuuzuxvhnkmxofscvpkdxwryymprkdwaoycoyhmvhqnjazwddrqwsbgowgkliobynbabidkobkypdulcmeywdznhezmcapmlcrruggydjlkzecwmyhtmormjafigtnaufiuyfrlxslsneorscrsckpwzkgqmmpanbfifqxaujmjkrjwuiytwlkdftwqlqtppvsgyymthetwocqilhwndwhrlwpmmlvngjijwebxfzschcpakiqryjqrdcdxkudmujlxxajmmdscycyruznzlbkzvfywfelajzaaoffcrffvomqdbaiaowqtgpykhxgyjffksxwyytxibvphltrhoxywqyprulnyfzplwnleadhcuyfwjszpwyfaikurghyyvexihcyjzvoivjtwmpkhahtnjjlbkhtaomyhjunzxnanylmsvdrozrnjpccwldnexfqluvrbenqatgoiqotmyikmhvyjpnihsyxnvnvibnhgdvwukbzbaogpafoipbspbkxcygdwcxomdpdtgddua...
		//615804
		s := "iqssxdlb"
		t := "dyuqrwyr"
		k := 40
		res := string_issue.CanConvertString(s, t, k)
		fmt.Println(res)
	}
	{
		var arr []int = []int{4, 7, 11}
		var k int = 5
		res := number.FindKthPositive(arr, k)
		fmt.Println(res)
	}
	{
		nums1 := []int{4, 5, 8, 10, 23, 45}
		nums2 := []int{4, 8, 23, 44, 45, 99}
		res := array.MaxSum(nums1, nums2)
		fmt.Println(res)
	}
	{
		s := "aaabcccabccbcadacbbcbcdabcdcd"
		k := 5
		res := string_issue.GetLengthOfOptimalCompression(s, k)
		fmt.Println(res)
	}
	{
		var arr []int = []int{2, 1, 4}
		res := array.NumOfSubarrays2(arr)
		fmt.Println(res)
	}
	{
		arr := []int{9, 12, 3, 7, 15}
		target := 5
		res := number.ClosestToTarget(arr, target)
		fmt.Println(res)
	}
	{
		n := 7
		edges := [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}
		labels := "abaedcd"
		res := tree.CountSubTrees(n, edges, labels)
		fmt.Println(res)
	}
	{
		n := 8
		res := number.WinnerSquareGame(n)
		fmt.Println(res)
	}
	{
		nums := []int{6, 6, 0, 1, 1, 4, 6}
		res := array.MinDifference1509(nums)
		fmt.Println(res)
	}
	{
		//[-1,1,-2,2,-3,3,-4,4]
		//3
		arr := []int{-1, 1, -2, 2, -3, 3, -4, 4}
		k := 3
		res := number.CanArrange(arr, k)
		fmt.Println(res)
	}
	{
		//nums := []int{0,1,1,1,0,1,1,0,1}
		nums := []int{0, 0, 0}
		res := array.LongestSubarray(nums)
		fmt.Println(res)
	}
	{
		rains := []int{1, 0, 2, 0, 3, 0, 2, 0, 0, 0, 1, 2, 3}
		res := array.AvoidFlood(rains)
		fmt.Println(res)
	}
	{
		names := []string{"kaido", "kaido(1)", "kaido", "kaido(1)"}
		res := string_issue.GetFolderNames(names)
		fmt.Println(res)
	}
	{
		//houses = [1,4,8,10,20], k = 3
		houses := []int{1, 4, 8, 10, 20}
		k := 3
		res := array.MinDistance(houses, k)
		fmt.Println(res)
	}
	{
		parents := []int{-1, 0, 3, 0, 3, 1}
		res := tree.CountHighestScoreNodes(parents)
		fmt.Println(res)
	}
	{
		houses := []int{0, 2, 1, 2, 0}
		cost := [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}}
		m := 5
		n := 2
		target := 3
		res := diagram.MinCost1473(houses, cost, m, n, target)
		fmt.Println(res)
	}
	{
		arr := []int{7, 3, 4, 7}
		target := 7
		res := array.MinSumOfLengths(arr, target)
		fmt.Println(res)
	}
	{
		prices := []int{8, 4, 6, 2, 3}
		res := array.FinalPrices(prices)
		fmt.Println(res)
	}
	{
		grid := [][]int{{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1}}
		res := array.CherryPickup2(grid)
		fmt.Println(res)
	}
	{
		var n int = 3
		var prerequisites [][]int = [][]int{{1, 0}, {2, 0}}
		var queries [][]int = [][]int{{0, 1}, {2, 0}}
		res := diagram.CheckIfPrerequisite(n, prerequisites, queries)
		fmt.Println(res)
	}
	{
		//[-5,-1,-2]
		//[3,3,5,5]
		nums1 := []int{-5, -1, -2}
		nums2 := []int{3, 3, 5, 5}
		res := array.MaxDotProduct2(nums1, nums2)
		fmt.Println(res)
	}
	{
		sentence := "i love eating burger"
		searchWord := "burg"
		res := string_issue.IsPrefixOfWord(sentence, searchWord)
		fmt.Println(res)
	}
	{
		cost := []int{4, 3, 2, 5, 6, 7, 2, 5, 5}
		target := 5000
		res := number.LargestNumber2(cost, target)
		fmt.Println(res)
	}
	{
		s := "1 box has 3 blue 4 red 6 green and 12 yellow marbles"
		res := string_issue.AreNumbersAscending(s)
		fmt.Println(res)
	}
	{
		arr := []int{1, 3, 5, 7, 9}
		res := array.CountTriplets(arr)
		fmt.Println(res)
	}
	{
		hats := [][]int{{1, 2, 3}, {2, 3, 5, 6}, {1, 3, 7, 9}, {1, 8, 9}, {2, 5, 7}}
		res := diagram.NumberWays(hats)
		fmt.Println(res)
	}
	{
		nums := []int{1, 0, 0, 0, 1, 0, 0, 1}
		k := 2
		res := array.KLengthApart(nums, k)
		fmt.Println(res)
	}
	{
		//"vqglqavmsnmktsxwxcpxhuujuanxueuymzifycytalizwnvrjeoipfoqbiqdxsnclcvoafqwfwcmuwitjgqghkicc"
		//"wqvloqrxbfjuxwriltxhmrmfpzitkwhitwhvatmknyhzigcuxfsosxetioqfeyewoljymhdwgwvjcdhmkpdfbbzta"
		s1 := "vqglqavmsnmktsxwxcpxhuujuanxueuymzifycytalizwnvrjeoipfoqbiqdxsnclcvoafqwfwcmuwitjgqghkicc"
		s2 := "wqvloqrxbfjuxwriltxhmrmfpzitkwhitwhvatmknyhzigcuxfsosxetioqfeyewoljymhdwgwvjcdhmkpdfbbzta"
		res := string_issue.CheckIfCanBreak3(s1, s2)
		fmt.Println(res)
	}
	{
		num := 123456
		res := number.MaxDiff(num)
		fmt.Println(res)
	}
	{
		nums := []int{10, 2, -10, 5, 20}
		k := 2
		res := array.ConstrainedSubsetSum(nums, k)
		fmt.Println(res)
	}
	{
		input := [][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11, 13, 14}, {12, 13}}
		res := array.FindDiagonalOrder(input)
		fmt.Println(res)
	}
	{
		input := []int{1, 1000, 1}
		k := 1
		res := number.MaxScore2(input, k)
		fmt.Println(res)
	}
	{
		input := "0100"
		res := number.MaxScore(input)
		fmt.Println(res)
	}
	{
		s := "1234567890"
		k := 90
		res := number.NumberOfArrays(s, k)
		fmt.Println(res)
	}
	{
		words := []string{"mass", "as", "hero", "superhero"}
		res := string_issue.StringMatching(words)
		fmt.Println(res)
	}
	{
		grid := [][]int{{1, 5}, {2, 3}}
		x := 1
		res := number.MinOperations(grid, x)
		fmt.Println(res)
	}
	{
		values := []int{-1, -2, -3}
		res := number.StoneGameIII(values)
		fmt.Println(res)
	}
	{
		s := "1101"
		res := number.NumSteps(s)
		fmt.Println(res)
	}
	{
		satisfaction := []int{-2, 5, -1, 0, 3, -3}
		res := number.MaxSatisfaction(satisfaction)
		fmt.Println(res)
	}
	{
		n := 13
		res := number.CountLargestGroup(n)
		fmt.Println(res)
	}
	{
		grid := [][]int{{4, 1, 3}, {6, 1, 2}}
		res := diagram.HasValidPath(grid)
		fmt.Println(res)
	}
	{
		nums := []int{0, 1, 2, 3, 4}
		index := []int{0, 1, 2, 2, 1}
		res := array.CreateTargetArray(nums, index)
		fmt.Println(res)
	}
	{
		//[4,2,2,5,4,5,4,5,3,3,6,1,2,4,2,1,6,5,4,2,3,4,2,3,3,5,4,1,4,4,5,3,6,1,5,2,3,3,6,1,6,4,1,3]
		//2
		//53
		rolls := []int{4, 2, 2, 5, 4, 5, 4, 5, 3, 3, 6, 1, 2, 4, 2, 1, 6, 5, 4, 2, 3, 4, 2, 3, 3, 5, 4, 1, 4, 4, 5, 3, 6, 1, 5, 2, 3, 3, 6, 1, 6, 4, 1, 3}
		mean := 2
		n := 53
		res := number.MissingRolls(rolls, mean, n)
		fmt.Println(res)
	}
	{
		answerKey := "FFFTTFTTFT"
		k := 3
		res := string_issue.MaxConsecutiveAnswers(answerKey, k)
		fmt.Println(res)
	}
	{
		nums := []string{"777", "7", "77", "77"}
		target := "7777"
		res := string_issue.NumOfPairs(nums, target)
		fmt.Println(res)
	}
	{
		//10
		//[[2,3,6},{8,9,8},{5,9,7},{8,9,1},{2,9,2},{9,10,6},{7,10,10},{6,7,9},{4,9,7},{2,3,1]]
		n := 10
		rides := [][]int{{2, 3, 6}, {8, 9, 8}, {5, 9, 7}, {8, 9, 1}, {2, 9, 2}, {9, 10, 6}, {7, 10, 10}, {6, 7, 9}, {4, 9, 7}, {2, 3, 1}}
		res := diagram.MaxTaxiEarnings(n, rides)
		fmt.Println(res)
	}
	{
		changed := []int{1, 3, 4, 2, 6, 8, 1, 2}
		res := array.FindOriginalArray(changed)
		fmt.Println(res)
	}
	{
		//r2 := 64 | 240
		//fmt.Println(r2)
		//n := 3
		//reservedSeats := [][]int{{1,2},{1,3},{1,8},{2,6},{3,1},{3,10}}
		//2
		//[[2,1},{1,8},{2,6]]
		n := 2
		reservedSeats := [][]int{{2, 1}, {1, 8}, {2, 6}}
		res := array.MaxNumberOfFamilies(n, reservedSeats)
		fmt.Println(res)
	}
	{
		n := 7
		edges := [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}
		t := 2
		target := 4
		res := diagram.FrogPosition(n, edges, t, target)
		fmt.Println(res)
	}
	{
		//11
		//4
		//[5,9,6,10,-1,8,9,1,9,3,4]
		//[0,213,0,253,686,170,975,0,261,309,337]
		//n = 4, headID = 2, manager = [3,3,-1,2], informTime = [0,0,162,914]
		//n = 7, headID = 6, manager = [1,2,3,4,5,6,-1], informTime = [0,6,5,4,3,2,1]
		//n = 15, headID = 0, manager = [-1,0,0,1,1,2,2,3,3,4,4,5,5,6,6], informTime = [1,1,1,1,1,1,1,0,0,0,0,0,0,0,0]
		var n int = 11
		var headID int = 4
		var manager []int = []int{5, 9, 6, 10, -1, 8, 9, 1, 9, 3, 4}
		var informTime []int = []int{0, 213, 0, 253, 686, 170, 975, 0, 261, 309, 337}
		res := diagram.NumOfMinutes(n, headID, manager, informTime)
		fmt.Println(res)
	}
	{
		//[5,4,8,3,null,6,3]
		//var t1 tree.TreeNode
		//t1.Val = 4
		//var t2 tree.TreeNode
		//t2.Val = 8
		//var t3 tree.TreeNode
		//t3.Val = 6
		//var t4 tree.TreeNode
		//t4.Val = 1
		//var t5 tree.TreeNode
		//t5.Val = 9
		//var t6 tree.TreeNode
		//t6.Val = -5
		//var t7 tree.TreeNode
		//t7.Val = 4
		//var t8 tree.TreeNode
		//t8.Val = -3
		//var t9 tree.TreeNode
		//t9.Val = 10
		//t1.Left = &t2
		//t2.Left = &t3
		//t2.Right = &t4
		//t3.Left = &t5
		//t4.Left = &t6
		//t4.Right = &t7
		//t6.Right = &t8
		//t7.Right = &t9

		var t1 tree.TreeNode
		t1.Val = 4
		var t2 tree.TreeNode
		t2.Val = -3
		var t3 tree.TreeNode
		t3.Val = 6
		var t4 tree.TreeNode
		t4.Val = -5
		var t5 tree.TreeNode
		t5.Val = -2
		var t6 tree.TreeNode
		t6.Val = 5
		var t7 tree.TreeNode
		t7.Val = 9
		t1.Left = &t2
		t1.Right = &t3
		t2.Left = &t4
		t2.Right = &t5
		t3.Left = &t6
		t3.Right = &t7
		res := tree.MaxSumBST(&t1)
		fmt.Println(res)
	}
	{
		s := "leetcode"
		res := string_issue.SortString(s)
		fmt.Println(res)
	}
	{
		votes := []string{"BCA", "CAB", "CBA", "ABC", "ACB", "BAC"}
		res := number.RankTeams(votes)
		fmt.Println(res)
	}
	{
		//ret := number.Cmp_largestMultipleOfThree("112","111")
		//fmt.Println(ret)
		digits := []int{5, 2, 3}
		res := number.LargestMultipleOfThree(digits)
		fmt.Println(res)
	}
	{
		n := 3
		res := number.CountOrders(n)
		fmt.Println(res)
	}
	{
		//4
		//[3,-1,1,-1]
		//[-1,-1,0,-1]
		n := 4
		left := []int{3, -1, 1, -1}
		right := []int{-1, -1, 0, -1}
		res := tree.ValidateBinaryTreeNodes(n, left, right)
		fmt.Println(res)
	}
	{
		board := [][]byte{{'#', ' ', '#'}, {' ', ' ', '#'}, {'#', 'c', ' '}}
		word := "abc"
		res := diagram.PlaceWordInCrossword(board, word)
		fmt.Println(res)
	}
	{
		var s string = "aabcaa"
		res := string_issue.NumberOfSubstrings(s)
		fmt.Println(res)
	}
	{
		arr := []int{7, 7, 2, 1, 7, 7, 7, 3, 4, 1}
		res := array.MinJumps(arr)
		fmt.Println(res)
	}
	{
		arr := []int{13, 9, 7, 10, 6, 12}
		d := 3
		res := array.MaxJumps(arr, d)
		fmt.Println(res)
	}
	{
		input := [][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}
		k := 3
		res := array.KWeakestRows(input, k)
		fmt.Println(res)
	}
	{
		jobDifficulty := []int{11, 111, 22, 222, 33, 333, 44, 444}
		d := 6
		res := number.MinDifficulty(jobDifficulty, d)
		fmt.Println(res)
	}
	{
		restaurants := [][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}}
		veganFriendly := 1
		maxPrice := 50
		maxDistance := 10
		res := number.FilterRestaurants(restaurants, veganFriendly, maxPrice, maxDistance)
		fmt.Println(res)
	}
	{
		arr := []int{37, 37, 12, 28, 9, 100, 56, 80, 5, 12}
		res := array.ArrayRankTransform(arr)
		fmt.Println(res)
	}
	{
		mat := [][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}
		res := array.DiagonalSort(mat)
		fmt.Println(res)
	}
	{
		palindrome := "aba"
		res := string_issue.BreakPalindrome(palindrome)
		fmt.Println(res)
	}
	{
		num := 6
		res := number.Maximum69Number(num)
		fmt.Println(res)
	}
	{
		word := "YEAR"
		res := diagram.MinimumDistance(word)
		fmt.Println(res)
	}
	{
		text := "abcabcabc"
		res := string_issue.DistinctEchoSubstrings(text)
		fmt.Println(res)
	}
	{
		s := "noob"
		res := string_issue.MinInsertions1312(s)
		fmt.Println(res)
	}
	{
		watchedVideos := [][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}}
		friends := [][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}}
		id := 0
		level := 1
		res := diagram.WatchedVideosByFriends(watchedVideos, friends, id, level)
		fmt.Println(res)
	}
	{
		s := "1326#"
		res := string_issue.FreqAlphabets(s)
		fmt.Println(res)
	}
	{
		arr := []int{3, 0, 2, 1, 2}
		start := 2
		res := array.CanReach(arr, start)
		fmt.Println(res)
	}
	{
		board := []string{"E23", "2X2", "12S"}
		res := diagram.PathsWithMaxScore2(board)
		fmt.Println(res)
	}
	{
		//status := []int{1,0,1,0}
		//candies := []int{7,5,4,100}
		//keys := [][]int{{},{},{1},{}}
		//containedBoxes := [][]int{{1,2},{3},{},{}}
		//initialBoxes := []int{0}

		//status := []int{1,0,0,0,0,0}
		//candies := []int{1,1,1,1,1,1}
		//keys := [][]int{{1,2,3,4,5},{},{},{},{},{}}
		//containedBoxes := [][]int{{1,2,3,4,5},{},{},{},{},{}}
		//initialBoxes := []int{0}

		status := []int{1, 0, 0, 0}
		candies := []int{1, 2, 3, 4}
		keys := [][]int{{1, 2}, {3}, {}, {}}
		containedBoxes := [][]int{{2}, {3}, {1}, {}}
		initialBoxes := []int{0}
		//[1,0,0,0]
		//[1,2,3,4]
		//[[1,2},{3},{},{]]
		//[[2},{3},{1},{]]
		//[0]
		res := number.MaxCandies(status, candies, keys, containedBoxes, initialBoxes)
		fmt.Println(res)
	}
	{
		//[1547,83230,57084,93444,70879]
		//71237
		var arr []int = []int{1547, 83230, 57084, 93444, 70879}
		var target int = 71237
		res := number.FindBestValue(arr, target)
		fmt.Println(res)
	}
	{
		grid := [][]int{{0, 1, 1}, {1, 1, 1}, {1, 0, 0}}
		k := 2
		res := diagram.ShortestPath(grid, k)
		fmt.Println(res)
	}
	{
		grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		res := array.MinFallingPathSum2(grid)
		fmt.Println(res)
	}
	{
		s := "faagaabbclagedtaabbcwnejzpaabbcuarkgwgoaabbcefwra"
		k := 27
		res := string_issue.PalindromePartition(s, k)
		fmt.Println(res)
	}
	{
		steps := 300
		arrLen := 9999
		res := number.NumWays(steps, arrLen)
		fmt.Println(res)
	}
	{
		var moves [][]int = [][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}}
		res := diagram.Tictactoe(moves)
		fmt.Println(res)
	}
	{
		var nums []int = []int{1, 2, 3, 4, 4}
		res := array.MaxSumDivThree(nums)
		fmt.Println(res)
	}
	{
		words := []string{"dog", "cat", "dad", "good"}
		letters := []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'}
		score := []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		res := number.MaxScoreWords(words, letters, score)
		fmt.Println(res)
	}
	{
		input := [][]int{{0, 1}, {1, 1}}
		res := array.OddCells(2, 3, input)
		fmt.Println(res)
	}
	{
		//[1,1,2,1,1]
		//3
		var nums []int = []int{1, 1, 2, 1, 1}
		var k int = 3
		res := array.NumberOfSubarrays(nums, k)
		fmt.Println(res)
	}
	{
		s := "())()((("
		res := string_issue.MinRemoveToMakeValid(s)
		fmt.Println(res)
	}
	{
		n := 3
		start := 2
		res := number.CircularPermutation(n, start)
		fmt.Println(res)
	}
	{
		//var edges [][]int = [][]int{{2,1},{3,1},{4,2},{1,4}}
		var edges [][]int = [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}}
		res := diagram.FindRedundantDirectedConnection(edges)
		fmt.Println(res)
	}
	{
		coordinates := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}
		res := diagram.CheckStraightLine(coordinates)
		fmt.Println(res)
	}
	{
		n := 3
		rollMax := []int{1, 1, 1, 2, 2, 3}
		res := number.DieSimulator2(n, rollMax)
		fmt.Println(res)
	}
	{
		//input := [][]int{{0,0,0,22,0,24},{34,23,18,0,23,2},{11,39,20,12,0,0},{39,8,0,2,0,1},{19,32,26,20,20,30},{0,38,26,0,29,31}}
		//input := [][]int{{1,0,7},{2,0,6},{3,4,5},{0,3,0},{9,0,20}}
		grid := [][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}
		res := diagram.GetMaximumGold(grid)
		fmt.Println(res)
	}
	{
		grid := [][]int{{0, 0, 0, 0, 0, 1},
			{1, 1, 0, 0, 1, 0},
			{0, 0, 0, 0, 1, 1},
			{0, 0, 1, 0, 1, 0},
			{0, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 0, 0}}
		res := diagram.MinimumMoves(grid)
		fmt.Println(res)
	}
	{
		s := "pbbcggttciiippooaais"
		k := 2
		res := string_issue.RemoveDuplicates2(s, k)
		fmt.Println(res)
	}
	{
		n := 4
		connections := [][]int{{0, 1}, {1, 2}, {2, 3}}
		res := diagram.CriticalConnections(n, connections)
		fmt.Println(res)
	}
	{
		//[-5,4,-4,-3,5,-3]
		//3
		arr := []int{-9, 13, 4, -16, -12, -16, 3, -7, 5, -16, 16, 8, -1, -13, 15, 3}
		k := 6
		res := array.KConcatenationMaxSum2(arr, k)
		fmt.Println(res)
	}
	{
		s := "(ed(et(oc))el)"
		res := string_issue.ReverseParentheses(s)
		fmt.Println(res)
	}
	{
		distance := []int{1, 2, 3, 4}
		start := 0
		destination := 3
		res := diagram.DistanceBetweenBusStops(distance, start, destination)
		fmt.Println(res)
	}
	{
		//[1,3,2,-3,-2,5,5,-5,1]
		var l1 list_queue.ListNode
		l1.Val = 1
		var l2 list_queue.ListNode
		l2.Val = 3
		var l3 list_queue.ListNode
		l3.Val = 2
		var l4 list_queue.ListNode
		l4.Val = -3
		var l5 list_queue.ListNode
		l5.Val = -2
		var l6 list_queue.ListNode
		l6.Val = 5
		var l7 list_queue.ListNode
		l7.Val = 5
		var l8 list_queue.ListNode
		l8.Val = -5
		var l9 list_queue.ListNode
		l9.Val = 1

		l1.Next = &l2
		l2.Next = &l3
		l3.Next = &l4
		l4.Next = &l5
		l5.Next = &l6
		l6.Next = &l7
		l7.Next = &l8
		l8.Next = &l9

		res := list_queue.RemoveZeroSumSublists2(&l1)
		fmt.Println(res.Val)
	}
	{
		var grid [][]int = [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
		res := diagram.MaxDistance2(grid)
		fmt.Println(res)
	}
	{
		piles := []int{2, 7, 9, 4, 4}
		res := number.StoneGameII(piles)
		fmt.Println(res)
	}
	{
		target := "leetcode"
		res := diagram.AlphabetBoardPath2(target)
		fmt.Println(res)
	}
	{
		//6
		//[[4,1},{3,5},{5,2},{1,4},{4,2},{0,0},{2,0},{1,1]]
		//[[5,5},{5,0},{4,4},{0,3},{1,0]]
		//var n int = 6
		//var red_edges [][]int = [][]int{{4,1},{3,5},{5,2},{1,4},{4,2},{0,0},{2,0},{1,1}}
		//var blue_edges [][]int = [][]int{{5,5},{5,0},{4,4},{0,3},{1,0}}
		//n = 3, red_edges = [[0,1},{1,2]], blue_edges = []
		var n int = 3
		var red_edges [][]int = [][]int{{0, 1}, {1, 2}}
		var blue_edges [][]int = [][]int{}
		res := diagram.ShortestAlternatingPaths(n, red_edges, blue_edges)
		fmt.Println(res)
	}
	{
		word := "fd"
		res := string_issue.MinTimeToType(word)
		fmt.Println(res)
	}
	{
		req_skills := []string{"algorithms", "math", "java", "reactjs", "csharp", "aws"}
		people := [][]string{{"algorithms", "math", "java"}, {"algorithms", "math", "reactjs"},
			{"java", "csharp", "aws"}, {"reactjs", "csharp"}, {"csharp", "math"}, {"aws", "java"}}
		res := diagram.SmallestSufficientTeam(req_skills, people)
		fmt.Println(res)
	}
	{
		var hours []int = []int{9, 9, 6, 0, 6, 6, 9}
		//var hours []int = []int{6,6,9,0,9,9,6}
		res := array.LongestWPI(hours)
		fmt.Println(res)
	}
	{
		arr1 := []int{28, 6, 22, 8, 44, 17}
		arr2 := []int{22, 28, 8, 6}
		res := array.RelativeSortArray2(arr1, arr2)
		fmt.Println(res)
	}
	{
		var bookings [][]int = [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
		var n int = 5
		res := diagram.CorpFlightBookings2(bookings, n)
		fmt.Println(res)
	}
	{
		expression := "|(&(t,f,t),!(t))"
		res := string_issue.ParseBoolExpr(expression)
		fmt.Println(res)
	}
	{
		count := []int{0, 1, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		res := number.SampleStats(count)
		fmt.Println(res)
	}
	{
		s := "cbaacabcaaccaacababa"
		res := string_issue.SmallestSubsequence(s)
		fmt.Println(res)
	}
	{
		matrix := [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}
		target := 0
		res := array.NumSubmatrixSumTarget(matrix, target)
		fmt.Println(res)
	}
	{
		row := 2
		col := 6
		cells := [][]int{{1, 4}, {1, 3}, {2, 1}, {2, 5}, {2, 2}, {1, 5}, {2, 4}, {1, 2}, {1, 6}, {2, 3}, {2, 6}, {1, 1}}
		res := diagram.LatestDayToCross(row, col, cells)
		fmt.Println(res)
	}
	{
		//"ABABCCABAB"
		//"ABAB"
		str1 := "ABCABC"
		str2 := "ABC"
		res := string_issue.GcdOfStrings(str1, str2)
		fmt.Println(res)
	}
	{
		arr := []int{1, 9, 4, 6, 7}
		res := array.PrevPermOpt1(arr)
		fmt.Println(res)
	}
	{
		s := "aa"
		res := string_issue.LongestDupSubstring(s)
		fmt.Println(res)
	}
	{
		s := "011010"
		minJump := 2
		maxJump := 3
		res := string_issue.CanReach(s, minJump, maxJump)
		fmt.Println(res)
	}
	{
		nums := []int{1, 2, 4}
		k := 5
		res := number.MaxFrequency(nums, k)
		fmt.Println(res)
	}
	{
		nums := []int{10, 20, 15, 30, 20}
		k := 2
		res := array.MinSpaceWastedKResizing(nums, k)
		fmt.Println(res)
	}
	{
		blocked := [][]int{{1, 1}}
		source := []int{0, 0}
		target := []int{99999, 99999}
		res := diagram.IsEscapePossible(blocked, source, target)
		fmt.Println(res)
	}
	{
		A := []int{69, 22, 21, 27, 26, 62, 69, 81, 55, 85, 95, 40, 91, 33, 72, 88, 86}
		res := number.MinScoreTriangulation(A)
		fmt.Println(res)
	}
	{
		clips := [][]int{{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9}}
		time := 9
		res := array.VideoStitching(clips, time)
		fmt.Println(res)
	}
	{
		N := 3
		res := number.BaseNeg2(N)
		res2 := -5 / -2
		fmt.Println(res2)
		fmt.Println(res)
	}
	{
		S := "110101011011000011011111000000"
		N := 15
		res := string_issue.QueryString(S, N)
		fmt.Println(res)
	}
	{
		var K int = 1243
		res := number.SmallestRepunitDivByK(K)
		fmt.Println(res)
	}
	{
		preorder := []int{4, 2}
		res := tree.BstFromPreorder(preorder)
		fmt.Println(res)
	}
	{
		//n := 5
		//lamps := [][]int{{0,0},{4,4}}
		//queries := [][]int{{1,1},{1,0}}
		n := 6
		lamps := [][]int{{2, 5}, {4, 2}, {0, 3}, {0, 5}, {1, 4}, {4, 2}, {3, 3}, {1, 0}}
		queries := [][]int{{4, 3}, {3, 1}, {5, 3}, {0, 5}, {4, 4}, {3, 3}}
		res := diagram.GridIllumination(n, lamps, queries)
		fmt.Println(res)
	}
	{
		var neededApples int64 = 1000000000
		res := number.MinimumPerimeter(neededApples)
		fmt.Println(res)
	}
	{
		n := 4
		res := number.Clumsy(n)
		fmt.Println(res)
	}
	{
		//var A []int = []int{0,0,1,1,1}
		//var K int = 2
		//var A []int = []int{1,1,1,0,0,0,1,1,1,1,0}
		//var K int = 2
		var A []int = []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
		var K int = 3
		res := array.LongestOnes2(A, K)
		fmt.Println(res)
	}
	{
		words := []string{"bella", "label", "roller"}
		res := string_issue.CommonChars(words)
		fmt.Println(res)
	}
	{
		nums := []int{1, 1, 8, 1, 8}
		res := array.NumSquarefulPerms(nums)
		fmt.Println(res)
	}
	{
		nums := []int{0, 0, 0, 1, 0, 1, 1, 0}
		k := 1
		res := array.MinKBitFlips(nums, k)
		fmt.Println(res)
	}
	{
		var input [][]int = [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
		res := diagram.OrangesRotting(input)
		fmt.Println(res)
	}
	{
		//[1,2,1,3,4]
		//3
		nums := []int{1, 2, 1, 2, 3}
		k := 2
		res := array.SubarraysWithKDistinct(nums, k)
		fmt.Println(res)
	}
	{
		A := []int{8, -10, 10, -7, 4, -2}
		queries := [][]int{{6, 4}, {-7, 0}, {-3, 5}, {6, 1}, {-8, 1}, {-10, 2}}
		res := number.SumEvenAfterQueries(A, queries)
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
		t5.Left = &t1
		t5.Right = &t2
		t2.Left = &t4
		t2.Right = &t3
		voyage := []int{5, 2, 3, 4, 1}
		res := tree.FlipMatchVoyage(&t5, voyage)
		fmt.Println(res)
	}
	{
		input := []int{3, 2, 4, 1}
		res := array.PancakeSort2(input)
		fmt.Println(res)
	}
	{
		wordlist := []string{"KiTe", "kite", "hare", "Hare"}
		queries := []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}
		res := string_issue.Spellchecker(wordlist, queries)
		fmt.Println(res)
	}
	//{
	//	x := 5
	//	target := 501
	//	res := number.LeastOpsExpressTarget(x,target)
	//	fmt.Println(res)
	//}
	{
		points := [][]int{{0, 1}, {2, 1}, {1, 1}, {1, 0}, {2, 0}}
		res := diagram.MinAreaFreeRect(points)
		fmt.Println(res)
	}
	{
		strs := []string{"a"}
		res := string_issue.MinDeletionSize3(strs)
		fmt.Println(res)
	}
	{
		rods := []int{1, 2, 3, 6, 7, 8, 9, 10, 12, 14, 15, 44, 22, 33, 16, 13, 76, 33, 24, 42}
		res := number.TallestBillboard(rods)
		fmt.Println(res)
	}
	{
		rungs := []int{4, 8, 12, 16}
		dist := 3
		res := array.AddRungs(rungs, dist)
		fmt.Println(res)
	}
	{
		//cells := []int{0,1,1,1,0,0,0,0}
		//n := 99
		cells := []int{1, 0, 0, 1, 0, 0, 1, 0}
		n := 1000000000
		//cells := []int{1,1,0,1,1,0,1,1}
		//n := 6
		res := array.PrisonAfterNDays(cells, n)
		fmt.Println(res)
	}
	{
		strs := []string{"ca", "bb", "ac"}
		res := string_issue.MinDeletionSize(strs)
		fmt.Println(res)
	}
	{
		words := []string{"hello", "leetcode"}
		order := "hlabcdefgijkmnopqrstuvwxyz"
		res := string_issue.IsAlienSorted(words, order)
		fmt.Println(res)
	}
	{
		nums := []int{2, 3, 6, 7, 4, 12, 21, 39}
		//nums := []int{20,50,9,63}
		res := diagram.LargestComponentSize(nums)
		fmt.Println(res)
	}
	{
		var deck []int = []int{17, 13, 11, 2, 3, 5, 7}
		res := array.DeckRevealedIncreasing(deck)
		fmt.Println(res)
	}
	{
		var stones [][]int = [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}
		res := diagram.RemoveStones(stones)
		fmt.Println(res)
	}
	{
		var pushed []int = []int{1, 2, 3, 4, 5}
		var popped []int = []int{4, 5, 3, 2, 1}
		res := stack.ValidateStackSequences2(pushed, popped)
		fmt.Println(res)
	}
	{
		var A []int = []int{1, 1, 1, 2, 2, 2}
		res := array.MinIncrementForUnique2(A)
		fmt.Println(res)
	}
	{
		s := "aba"
		res := string_issue.DistinctSubseqII(s)
		fmt.Println(res)
	}
	{
		arr := []int{9, 7, 3}
		res := array.ValidMountainArray(arr)
		fmt.Println(res)
	}
	{
		var n int = 3
		res := diagram.KnightDialer(n)
		fmt.Println(res)
	}
	{
		var A [][]int = [][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}
		res := diagram.ShortestBridge(A)
		fmt.Println(res)
	}
	//{
	//	num := "25??"
	//	res := number.SumGame(num)
	//	fmt.Println(res)
	//}
	{
		//maze := [][]byte{{'+','+','.','+'},{'.','.','.','+'},{'+','+','+','.'}}
		maze := [][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}}
		entrance := []int{1, 0}
		res := diagram.NearestExit(maze, entrance)
		fmt.Println(res)
	}
	{
		n := 10
		res := number.CountTriples(n)
		fmt.Println(res)
	}
	{
		//input := [][]int{{-19,-1,-96,48,-94,36,16,55,-42,37,-59,6,-32,96,95,-58,13,-34,94,85},{17,44,36,-29,84,80,-34,50,-99,64,13,91,-27,25,-36,57,20,98,-100,-72},{-92,-75,86,90,-4,90,64,56,50,-63,10,-15,90,-66,-66,32,-69,-78,1,60},{21,51,-47,-43,-14,99,44,90,8,11,99,-62,57,59,69,50,-69,32,85,13},{-28,90,12,-18,23,61,-55,-97,6,89,36,26,26,-1,46,-50,79,-45,89,86},{-85,-10,49,-10,2,62,41,92,-67,85,86,27,89,-50,77,55,22,-82,-94,-98},{-50,53,-23,55,25,-22,76,-93,-7,66,-75,42,-35,-96,-5,4,-92,13,-31,-100},{-62,-78,8,-92,86,69,90,-37,81,97,53,-45,34,19,-19,-39,-88,-75,-74,-4},{29,53,-91,65,-92,11,49,26,90,-31,17,-84,12,63,-60,-48,40,-49,-48,88},{100,-69,80,11,-93,17,28,-94,52,64,-86,30,-9,-53,-8,-68,-33,31,-5,11},{9,64,-31,63,-84,-15,-30,-10,67,2,98,73,-77,-37,-96,47,-97,78,-62,-17},{-88,-38,-22,-90,54,42,-29,67,-85,-90,-29,81,52,35,13,61,-18,-94,61,-62},{-23,-29,-76,-30,-65,23,31,-98,-9,11,75,-1,-84,-90,73,58,72,-48,30,-81},{66,-33,91,-6,-94,82,25,-43,-93,-25,-69,10,-71,-65,85,28,-52,76,25,90},{-3,78,36,-92,-52,-44,-66,-53,-55,76,-7,76,-73,13,-98,86,-99,-22,61,100},{-97,65,2,-93,56,-78,22,56,35,-24,-95,-13,83,-34,-51,-73,2,7,-86,-19},{32,94,-14,-13,-6,-55,-21,29,-21,16,67,100,77,-26,-96,22,-5,-53,-92,-36},{60,93,-79,76,-91,43,-95,-16,74,-21,85,43,21,-68,-32,-18,18,100,-43,1},{87,-31,26,53,26,51,-61,92,-65,17,-41,27,-42,-14,37,-46,46,-31,-74,23},{-67,-14,-20,-85,42,36,56,9,11,-66,-59,-55,5,64,-29,77,47,44,-33,-77}}
		input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		res := array.MinFallingPathSum(input)
		fmt.Println(res)
	}
	{
		var graph [][]int = [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
		var initial []int = []int{0, 1}
		res := diagram.MinMalwareSpread2(graph, initial)
		fmt.Println(res)
	}
	{
		arr := []int{0, 1, 0, 1, 1}
		res := number.ThreeEqualParts(arr)
		fmt.Println(res)
	}
	{
		var s string = "11011"
		res := string_issue.MinFlipsMonoIncr(s)
		fmt.Println(res)
	}
	{
		//"alex"
		//"aaleexeex"
		name := "alex"
		typed := "aaleexeex"
		res := string_issue.IsLongPressedName(name, typed)
		fmt.Println(res)
	}
	{
		//[[1,0,0},{0,1,0},{0,0,1]]
		// [0,2]
		//graph := [][]int{{1,0,0},{0,1,1},{0,1,1}}
		//initial := []int{0,1}

		//graph := [][]int{{1,1,0},{1,1,0},{0,0,1}}
		//initial := []int{0,1,2}

		graph := [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
		initial := []int{1, 2}
		res := diagram.MinMalwareSpread(graph, initial)
		fmt.Println(res)
	}
	{
		s := "Test1ng-Leet=code-Q!"
		res := string_issue.ReverseOnlyLetters(s)
		fmt.Println(res)
	}
	{
		var A []int = []int{1, 1, 1, 0, 6, 12}
		res := array.PartitionDisjoint(A)
		fmt.Println(res)
	}
	{
		var nums []int = []int{5, 1, 1, 2, 0, 0}
		res := array.SortArray(nums)
		fmt.Println(res)
	}
	{
		//var board [][]int = [][]int{
		//		{-1,-1,-1,-1,-1,-1},
		//		{-1,-1,-1,-1,-1,-1},
		//		{-1,-1,-1,-1,-1,-1},
		//		{-1,35,-1,-1,13,-1},
		//		{-1,-1,-1,-1,-1,-1},
		//		{-1,15,-1,-1,-1,-1}}
		//var board [][]int = [][]int{{-1,-1,-1},{-1,9,8},{-1,8,9}}
		//var board [][]int = [][]int{{1,1,-1},{1,1,1},{-1,1,1}}
		//var board [][]int = [][]int{{-1,1,2,-1},{2,13,15,-1},{-1,10,-1,-1},{-1,6,2,8}}
		//var board [][]int = [][]int{{-1,-1,-1,46,47,-1,-1,-1},{51,-1,-1,63,-1,31,21,-1},{-1,-1,26,-1,-1,38,-1,-1},{-1,-1,11,-1,14,23,56,57},
		//	{11,-1,-1,-1,49,36,-1,48},{-1,-1,-1,33,56,-1,57,21},{-1,-1,-1,-1,-1,-1,2,-1},{-1,-1,-1,8,3,-1,6,56}}
		var board [][]int = [][]int{{-1, -1, -1, 30, -1, 144, 124, 135, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1, 167, 93, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1, -1, 77, -1, -1, 90, -1, -1}, {-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 122, -1}, {-1, -1, -1, 23, -1, -1, -1, -1, -1, 155, -1, -1, -1}, {-1, -1, 140, 29, -1, -1, -1, -1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, 115, -1, -1, -1, 109, -1, -1, 5},
			{-1, 57, -1, 99, 121, -1, -1, 132, -1, -1, -1, -1, -1}, {-1, 48, -1, -1, -1, 68, -1, -1, -1, -1, 31, -1, -1}, {-1, 163, 147, -1, 77, -1, -1, 114, -1, -1, 80, -1, -1}, {-1, -1, -1, -1, -1, 57, 28, -1, -1, 129, -1, -1, -1}, {-1, -1, -1, -1, -1, -1, -1, -1, -1, 87, -1, -1, -1}}
		res := array.SnakesAndLadders(board)
		fmt.Println(res)
	}
	{
		var A []int = []int{3, 1, 2, 4}
		res := array.SumSubarrayMins(A)
		fmt.Println(res)
	}
	{
		var nums []int = []int{3, 4, 1, 2}
		res := array.SortArrayByParity(nums)
		fmt.Println(res)
	}
	{
		s := "DID"
		res := array.NumPermsDISequence(s)
		fmt.Println(res)
	}
	{
		var fruits []int = []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
		res := array.TotalFruit(fruits)
		fmt.Println(res)
	}
	{
		//["1","2","3","4","6","7","9"]
		//333
		//["3","4","5","6"]
		//66
		var digits []string = []string{"7"}
		var n int = 7
		res := number.AtMostNGivenDigitSet(digits, n)
		fmt.Println(res)
	}
	{
		s := "baaca"
		k := 1
		res := string_issue.OrderlyQueue(s, k)
		fmt.Println(res)
	}
	{
		var N int = 7
		res := tree.AllPossibleFBT(N)
		fmt.Println(len(res))
	}
	{
		words := []string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}
		res := string_issue.NumSpecialEquivGroups(words)
		fmt.Println(res)
	}
	{
		var nums []int = []int{96, 87, 191, 197, 40, 101, 108, 35, 169, 50, 168, 182, 95, 80, 144, 43, 18, 60, 174, 13, 77, 173, 38, 46, 80, 117, 13, 19, 11, 6, 13, 118, 39, 80, 171, 36, 86, 156, 165, 190, 53, 49, 160, 192, 57, 42, 97, 35, 124, 200, 84, 70, 145, 180, 54, 141, 159, 42, 66, 66, 25, 95, 24, 136, 140, 159, 71, 131, 5, 140, 115, 76, 151, 137, 63, 47, 69, 164, 60, 172, 153, 183, 6, 70, 40, 168, 133, 45, 116, 188, 20, 52, 70, 156, 44, 27, 124, 59, 42, 172}
		res := array.SumSubseqWidths(nums)
		fmt.Println(res)
	}
	{
		words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
		pattern := "abb"
		res := string_issue.FindAndReplacePattern(words, pattern)
		fmt.Println(res)
	}
	{
		k := 2
		n := 9
		res := number.SuperEggDrop(k, n)
		fmt.Println(res)
	}
	{
		var aliceSizes []int = []int{1, 1}
		var bobSizes []int = []int{2, 2}
		res := number.FairCandySwap(aliceSizes, bobSizes)
		fmt.Println(res)
	}
	{
		//var N int = 10
		//var dislikes [][]int= [][]int{{4,7},{4,8},{5,6},{1,6},{3,7},{2,5},{5,8},{1,2},{4,9},{6,10},{8,10},{3,6},{2,10},{9,10},{3,9},{2,3},{1,9},{4,6},{5,7},{3,8},{1,8},{1,7},{2,4}}
		var N int = 4
		var dislikes [][]int = [][]int{{1, 2}, {1, 3}, {2, 4}}
		res := diagram.PossibleBipartition(N, dislikes)
		fmt.Println(res)
	}
	{
		var S string = "leet2code3"
		var K int = 10
		res := string_issue.DecodeAtIndex(S, K)
		fmt.Println(res)
	}
	{
		var word string = "aba"
		res := string_issue.WonderfulSubstrings(word)
		fmt.Println(res)
	}
	{
		var grid [][]int = [][]int{{40, 10}, {30, 20}}
		var k int = 1
		res := array.RotateGrid(grid, k)
		fmt.Println(res)
	}
	{
		//64
		//0
		//[80, 40]
		//[88, 88]
		n := 5
		minProfit := 3
		var group []int = []int{2, 2}
		var profit []int = []int{2, 3}
		res := diagram.ProfitableSchemes(n, minProfit, group, profit)
		fmt.Println(res)
	}
	{
		input := []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}
		res := array.LongestConsecutive(input)
		fmt.Println(res)
	}
	{
		//1000000000
		//40000
		//29999
		n := 1000000000
		a := 40000
		b := 29999
		res := number.NthMagicalNumber(n, a, b)
		fmt.Println(res)
	}
	{
		var piles []int = []int{5, 3, 4, 5}
		res := number.StoneGame(piles)
		fmt.Println(res)
	}
	{
		var N int = 22
		res := number.BinaryGap(N)
		fmt.Println(res)
	}
	{
		var grid []string = []string{"Dd#b@", ".fE.e", "##.B.", "#.cA.", "aF.#C"}
		res := diagram.ShortestPathAllKeys(grid)
		fmt.Println(res)
	}
	{
		var nums []int = []int{8, 1, 1, 8}
		var queries [][]int = [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 3}}
		res := array.MinDifference(nums, queries)
		fmt.Println(res)
	}
	{
		var grid1 [][]int = [][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}}
		var grid2 [][]int = [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}}
		res := diagram.CountSubIslands(grid1, grid2)
		fmt.Println(res)
	}
	{
		var startTime string = "22:45"
		var finishTime string = "23:44"
		res := number.NumberOfRounds(startTime, finishTime)
		fmt.Println(res)
	}
	{
		num := "1246"
		res := number.LargestOddNumber(num)
		fmt.Println(res)
	}
	{
		//"aabcaabcabcbaabcc"
		//"bcabacabcabcabcaa"
		var s1 string = "aabcaabcabcbaabcc"
		var s2 string = "bcabacabcabcabcaa"
		res := string_issue.KSimilarity(s1, s2)
		fmt.Println(res)
	}
	{
		var arr []int = []int{3, 5, 3, 2, 0}
		res := array.PeakIndexInMountainArray(arr)
		fmt.Println(res)
	}
	{
		var rectangles [][]int = [][]int{{39, 31, 70, 52}, {2, 8, 86, 31}, {28, 5, 78, 80}, {45, 96, 62, 100}, {17, 64, 72, 67}, {20, 41, 66, 68}}
		res := diagram.RectangleArea(rectangles)
		fmt.Println(res)
	}
	{
		var rooms [][]int = [][]int{{1, 3, 2}, {2, 3}, {2, 1, 3, 1}, {}}
		res := diagram.CanVisitAllRooms(rooms)
		fmt.Println(res)
	}
	{
		var grid [][]int = [][]int{{1, 1, 1, 4}, {4, 6, 9, 2}}
		res := array.LargestMagicSquare(grid)
		fmt.Println(res)
	}
	{
		var ranges [][]int = [][]int{{2, 4}}
		var left int = 1
		var right int = 3
		res := array.IsCovered(ranges, left, right)
		fmt.Println(res)
	}
	{
		var rec1 []int = []int{-1, 0, 1, 1}
		var rec2 []int = []int{0, -1, 0, 1}
		res := diagram.IsRectangleOverlap(rec1, rec2)
		fmt.Println(res)
	}
	{
		var img1 [][]int = [][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}}
		var img2 [][]int = [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}}
		res := array.LargestOverlap(img1, img2)
		fmt.Println(res)
	}
	{
		//var s string = "abzzcdqq"
		//var indexes []int = []int{0, 2}
		//var sources []string = []string{"ab","ec"}
		//var targets []string = []string{"eee", "ffff"}

		var s string = "abcd"
		var indexes []int = []int{0, 2}
		var sources []string = []string{"a", "cd"}
		var targets []string = []string{"eee", "ffff"}
		res := string_issue.FindReplaceString(s, indexes, sources, targets)
		fmt.Println(res)
	}
	{
		var image [][]int = [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}
		res := array.FlipAndInvertImage(image)
		fmt.Println(res)
	}
	{
		var triplets [][]int = [][]int{{7, 15, 15}, {11, 8, 3}, {5, 3, 4}, {12, 9, 9}, {5, 12, 10}, {7, 15, 10},
			{7, 6, 4}, {3, 9, 8}, {2, 13, 1}, {14, 2, 3}}
		var target []int = []int{14, 6, 4}
		res := array.MergeTriplets(triplets, target)
		fmt.Println(res)
	}
	{
		//s = "abcacb", p = "ab", removable = [3,1,0]
		//var s = "cdijywkabpzxph"
		//var p = "iwpph"
		//var removable []int = []int{11,0,1,8,6}
		var s = "abcbddddd"
		var p = "abcd"
		var removable []int = []int{3, 2, 1, 4, 5, 6}
		res := string_issue.MaximumRemovals(s, p, removable)
		fmt.Println(res)
	}
	{
		var words []string = []string{"abc", "aabc", "bc"}
		res := string_issue.MakeEqual(words)
		fmt.Println(res)
	}
	{
		var n int = 5
		res := number.ConsecutiveNumbersSum(n)
		fmt.Println(res)
	}
	{
		var s string = "ABC"
		res := string_issue.UniqueLetterString(s)
		fmt.Println(res)
	}
	{
		var grid [][]int = [][]int{{1, 1}, {1, 0}}
		res := diagram.LargestIsland(grid)
		fmt.Println(res)
	}
	{
		var paragraph string = "Bob hit a ball, the hit BALL flew far after it was hit."
		var banned []string = []string{"hit"}
		res := string_issue.MostCommonWord(paragraph, banned)
		fmt.Println(res)
	}
	{
		input := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
		res := string_issue.SubdomainVisits(input)
		fmt.Println(res)
	}
	{
		//"heeellooo"
		//["hello", "hi", "helo"]
		var S string = "heeellooo"
		var words []string = []string{"hello", "hi", "helo"}
		res := string_issue.ExpressiveWords(S, words)
		fmt.Println(res)
	}
	{
		var N int = 60
		res := number.SoupServings(N)
		fmt.Println(res)
	}
	{
		var nums []int = []int{1, 2, 3, 4, 5, 6, 7, 8}
		res := array.SplitArraySameAverage(nums)
		fmt.Println(res)
	}
	{
		var s string = "01001001101"
		res := string_issue.MinFlips(s)
		fmt.Println(res)
	}
	{
		//[[1,0,0,0},{1,1,0,0]]
		//[[1,1},{1,0]]
		var grid [][]int = [][]int{{1, 0, 0, 0}, {1, 1, 0, 0}}
		var hits [][]int = [][]int{{1, 1}, {1, 0}}
		res := diagram.HitBricks(grid, hits)
		fmt.Println(res)
	}
	{
		//var mat [][]int = [][]int{{0,0,0},{0,1,0},{1,1,1}}
		//var target [][]int = [][]int{{1,1,1},{0,1,0},{0,0,0}}
		var mat [][]int = [][]int{{1, 1}, {0, 1}}
		var target [][]int = [][]int{{1, 1}, {1, 0}}
		res := array.FindRotation(mat, target)
		fmt.Println(res)
	}
	{
		var A []int = []int{1, 3, 0, 2, 4}
		res := array.BestRotation(A)
		fmt.Println(res)
	}
	{
		var s string = "abcde"
		var goal string = "abced"
		res := string_issue.RotateString(s, goal)
		fmt.Println(res)
	}
	{
		var nums []int = []int{2, 1, 4, 3}
		var left int = 3
		var right int = 6
		res := array.NumSubarrayBoundedMax(nums, left, right)
		fmt.Println(res)
	}
	{
		var board []string = []string{"XOX", "OXO", "XOX"}
		res := diagram.ValidTicTacToe(board)
		fmt.Println(res)
	}
	{
		var S string = "cbsdf"
		var T string = "abckw"
		res := string_issue.CustomSortString(S, T)
		fmt.Println(res)
	}
	{
		var N int = 4
		res := diagram.NumTilings(N)
		fmt.Println(res)
	}
	{
		//var n int = 5
		//var edges [][]int = [][]int{{0,1,100},{0,2,100},{0,3,10},{1,2,100},{1,4,10},{2,1,10},{2,3,100},{2,4,100},{3,2,10},{3,4,100}}
		//var src int = 0
		//var dst int = 4
		//var k int = 3

		var n int = 3
		var edges [][]int = [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
		var src int = 0
		var dst int = 2
		var k int = 1
		res := diagram.FindCheapestPrice(n, edges, src, dst, k)
		fmt.Println(res)
	}
	{
		//[1,0,48,null,null,12,49]
		var t1 tree.TreeNode
		t1.Val = 1
		var t2 tree.TreeNode
		t2.Val = 0
		var t3 tree.TreeNode
		t3.Val = 48
		var t4 tree.TreeNode
		t4.Val = 12
		var t5 tree.TreeNode
		t5.Val = 49
		t1.Left = &t2
		t1.Right = &t3
		t3.Left = &t4
		t3.Right = &t5
		res := tree.MinDiffInBST(&t1)
		fmt.Println(res)
	}
	{
		firstWord := "acb"
		secondWord := "cba"
		targetWord := "cdb"
		res := string_issue.IsSumEqual(firstWord, secondWord, targetWord)
		fmt.Println(res)
	}
	{
		var board [][]int = [][]int{{1, 2, 3}, {4, 0, 5}}
		res := diagram.SlidingPuzzle(board)
		fmt.Println(res)
	}
	{
		var arr []int = []int{2, 1, 3, 4, 4}
		res := array.MaxChunksToSorted2(arr)
		fmt.Println(res)
	}
	{
		var arr []int = []int{1, 2, 3, 4, 5, 0}
		res := array.MaxChunksToSorted(arr)
		fmt.Println(res)
	}
	{
		input := "ababcbacadefegdehijhklij"
		res := string_issue.PartitionLabels(input)
		fmt.Println(res)
	}
	{
		var bottom string = "CCC"
		var allowed []string = []string{"CBB", "ACB", "ABD", "CDB", "BDC", "CBC", "DBA", "DBB", "CAB", "BCB", "BCC", "BAA", "CCD", "BDD", "DDD", "CCA", "CAA", "CCC", "CCB"}
		res := diagram.PyramidTransition(bottom, allowed)
		fmt.Println(res)
	}
	{
		var target int = 5
		res := number.ReachNumber(target)
		fmt.Println(res)
	}
	{
		var cost []int = []int{10, 15, 20}
		res := array.MinCostClimbingStairs2(cost)
		fmt.Println(res)
	}
	{
		var input []byte = []byte{'c', 'f', 'j'}
		res := string_issue.NextGreatestLetter(input, 'c')
		fmt.Println(res)
	}
	{
		var times [][]int = [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
		var N int = 4
		var K int = 2
		res := diagram.NetworkDelayTime(times, N, K)
		fmt.Println(res)
	}
	{
		var dist []int = []int{1, 3, 2}
		var hour float64 = 2.1
		res := number.MinSpeedOnTime(dist, hour)
		fmt.Println(res)
	}
	{

		var grid [][]int = [][]int{{1, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 0},
			{0, 0, 0, 1, 0, 0, 1}, {1, 0, 0, 1, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 0}, {0, 0, 0, 1, 1, 1, 1}}
		res := array.CherryPickup(grid)
		fmt.Println(res)
	}
	{
		input := []int{73, 74, 75, 71, 69, 72, 76, 73}
		res := list_queue.DailyTemperatures(input)
		fmt.Println(res)
	}
	{
		var N int = 984
		res := number.MonotoneIncreasingDigits(N)
		fmt.Println(res)
	}
	{
		var asteroids []int = []int{1, -2, -2, -2}
		res := list_queue.AsteroidCollision(asteroids)
		fmt.Println(res)
	}
	{
		var left int = 1
		var right int = 22
		res := number.SelfDividingNumbers(left, right)
		fmt.Println(res)
	}
	{
		var nums []int = []int{1, 7, 3, 6, 5, 6}
		res := number.PivotIndex(nums)
		fmt.Println(res)
	}
	{
		var source []string = []string{"a/*/b//*c", "blank", "d//*e/*/f"}
		res := string_issue.RemoveComments(source)
		fmt.Println(res)
	}
	{
		var A []int = []int{9, 1, 2, 8, 1, 2, 3, 4, 5, 9}
		var B []int = []int{0, 1, 2, 1, 4, 1, 2, 3, 4, 5, 0}
		res := array.FindLength(A, B)
		fmt.Println(res)
	}
	{
		var box [][]byte = [][]byte{{'#', '#', '*', '.', '*', '.'},
			{'#', '#', '#', '*', '.', '.'},
			{'#', '#', '#', '.', '#', '.'}}
		res := array.RotateTheBox(box)
		fmt.Println(res)
	}
	{
		var nums []int = []int{2, 5, 4, 2, 4, 5, 3, 1, 2, 4}
		res := number.MaxSumMinProduct(nums)
		fmt.Println(res)
	}
	{
		var memory1 int = 2
		var memory2 int = 2
		res := number.MemLeak(memory1, memory2)
		fmt.Println(res)
	}
	{
		var cards []int = []int{4, 1, 8, 7}
		exp := number.JudgePoint24ex(cards)
		fmt.Println(exp)
	}
	{
		var nums []int = []int{4, 3, 2, 3, 5, 2, 1}
		var k int = 4
		res := array.CanPartitionKSubsets(nums, k)
		fmt.Println(res)
	}
	{
		var s string = "00110011"
		res := string_issue.CountBinarySubstrings(s)
		fmt.Println(res)
	}
	{
		var nums []int = []int{1, 2, 1, 2, 6, 7, 5, 1}
		var k int = 2
		res := array.MaxSumOfThreeSubarrays2(nums, k)
		fmt.Println(res)
	}
	{
		var equation string = "x=x+2"
		res := number.SolveEquation(equation)
		fmt.Println(res)
	}
	{
		var m int = 1
		var n int = 3
		var k int = 2
		res := number.FindKthNumber2(m, n, k)
		fmt.Println(res)
	}
	{
		var t1 tree.TreeNode
		t1.Val = 5
		var t2 tree.TreeNode
		t2.Val = 8
		var t3 tree.TreeNode
		t3.Val = 5
		t1.Left = &t2
		t1.Right = &t3
		res := tree.FindSecondMinimumValue(&t1)
		fmt.Println(res)
	}
	//{
	//	var t1 tree.TreeNode
	//	t1.Val = 1
	//	var t2 tree.TreeNode
	//	t2.Val = 2
	//	var t3 tree.TreeNode
	//	t3.Val = 3
	//	var t4 tree.TreeNode
	//	t4.Val = 4
	//	t3.Left = &t1
	//	t3.Right = &t4
	//	t1.Right = &t2
	//	var low int = 1
	//	var high int = 2
	//	res := tree.TrimBST(&t3,low,high)
	//	fmt.Println(res)
	//}
	{
		var nums []int = []int{4, 2, 1}
		res := array.CheckPossibility(nums)
		fmt.Println(res)
	}
	{
		//var arr []int = []int{1,7,9,11,19,25}
		//var k int = 4
		//var x int = 10
		var arr []int = []int{1, 2, 3, 4, 5}
		var k int = 4
		var x int = 3
		res := array.FindClosestElements2(arr, k, x)
		fmt.Println(res)
	}
	{
		var t1 tree.TreeNode
		var t2 tree.TreeNode
		var t3 tree.TreeNode
		var t4 tree.TreeNode
		t1.Val = 1
		t2.Val = 2
		t3.Val = 3
		t4.Val = 4
		t1.Left = &t2
		//t1.Right = &t3
		//t2.Right = &t4
		res := tree.PrintTree(&t1)
		fmt.Println(res)
	}
	{
		var s string = "1234"
		res := string_issue.SplitString(s)
		fmt.Println(res)
	}
	{
		var s string = "1*1"
		res := number.NumDecodings(s)
		fmt.Println(res)
	}
	{
		var nums []int = []int{8, 7, 3, 5, 3, 6, 1, 4}
		res := number.FindErrorNums(nums)
		fmt.Println(res)
	}
	{
		var nums []int = []int{1, 12, -5, -6, 50, 3}
		var k int = 4
		res := array.FindMaxAverage(nums, k)
		fmt.Println(res)
	}
	{
		var n int = 1
		var logs []string = []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"}
		res := list_queue.ExclusiveTime(n, logs)
		fmt.Println(res)
	}
	{
		other.Context_timout_test()
	}
	{
		var nums []int = []int{2, 3, 4, 7}
		var maximumBit int = 3
		res := number.GetMaximumXor(nums, maximumBit)
		fmt.Println(res)
	}
	{
		var points [][]int = [][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}
		var queries [][]int = [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}}
		res := diagram.CountPoints(points, queries)
		fmt.Println(res)
	}
	{
		var paths []string = []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)",
			"root/c/d 4.txt(efgh)", "root 4.txt(efgh)"}
		res := string_issue.FindDuplicate(paths)
		fmt.Println(res)
	}
	{
		var flowerbed []int = []int{0, 0, 1, 0, 1}
		var n int = 1
		res := array.CanPlaceFlowers(flowerbed, n)
		fmt.Println(res)
	}
	{
		var n int = 5
		res := number.FindIntegers(n)
		fmt.Println(res)
	}
	{
		var nums [][]int = [][]int{{1, 2, 3, 4}}
		var r int = 2
		var c int = 2
		res := array.MatrixReshape(nums, r, c)
		fmt.Println(res)
	}
	{
		input := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		target := 0
		res := array.SubarraySum(input, target)
		fmt.Println(res)
	}
	{
		var n int = 2147483476
		res := number.NextGreaterElement2(n)
		fmt.Println(res)
	}
	{
		var n int = 1
		res := string_issue.CheckRecord2(n)
		fmt.Println(res)
	}
	{
		//var boxes []int = []int{1,3,2,2,2,3,4,3,1,3,3,2,2,1,1,4,6,6,7,7,7,7,7,3,3,3,2,1,3,2,2,2,3,4,3,1,3,3,2,2,1,1,4,6,6,7,7,7,7,7,3,3,3,2,1,3,2,2,2,3,4,3,1,3,3,2,2,1,1,4,6,6,7,7,7,7,7,3,3,3,2}
		//res := array.RemoveBoxes(boxes)
		//fmt.Println(res)
	}
	{
		var matrix [][]int = [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
		res := array.UpdateMatrix(matrix)
		fmt.Println(res)
	}
	{
		var s string = "abcdefg"
		var k int = 2
		res := string_issue.ReverseStr(s, k)
		fmt.Println(res)
	}
	{
		var nums []int = []int{1, 1, 2}
		res := array.SingleNonDuplicate(nums)
		fmt.Println(res)
	}
	{
		var nums []int = []int{23, 2, 4, 6, 6}
		var k int = 7
		res := array.CheckSubarraySum(nums, k)
		fmt.Println(res)
	}
	{
		input := []int{1, 2, 1}
		target := 0
		res := number.FindTargetSumWays(input, target)
		fmt.Println(res)
	}
	{
		var num int = 123
		res := number.FindComplement(num)
		fmt.Println(res)
	}
	{
		var houses []int = []int{1, 5}
		var heaters []int = []int{10}
		res := number.FindRadius2(houses, heaters)
		fmt.Println(res)
	}
	{
		//var words []string = []string{"cat","dog","catdog"}
		var words []string = []string{"catsdogcats", "cat", "cats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}
		res := string_issue.FindAllConcatenatedWordsInADict(words)
		fmt.Println(res)
	}
	{
		var nums []int = []int{4, 3, 2, 7, 8, 2, 3, 1}
		res := number.FindDisappearedNumbers(nums)
		fmt.Println(res)
	}
	{
		input := [][]int{{0, 0}, {1, 0}, {2, 0}}
		res := diagram.NumberOfBoomerangs(input)
		fmt.Println(res)
	}
	{
		var A []int = []int{2, 2, 3, 4}
		res := number.NumberOfArithmeticSlices(A)
		fmt.Println(res)
	}
	{
		var arr []int = []int{4, 3, 2, 7, 8, 2, 3, 1}
		var res = array.FindDuplicates(arr)
		fmt.Println(res)
	}
	{
		var n int = 1500
		var k int = 600
		res := number.FindKthNumber(n, k)
		fmt.Println(res)
	}
	{
		//Input: head = [1,2,null,3]
		//Output: [1,3,2]
		//Explanation:
		//
		//The input multilevel linked list is as follows:
		//
		//  1---2---NULL
		//  |
		//  3---NULL
		var n1 list_queue.Node
		n1.Val = 1
		var n2 list_queue.Node
		n2.Val = 2
		var n3 list_queue.Node
		n3.Val = 3
		n1.Child = &n2
		n2.Child = &n3
		//n2.Prev = &n1
		res := list_queue.Flatten(&n1)
		fmt.Println(res)
	}
	{
		var s string = "AAAA"
		var k int = 0
		res := string_issue.CharacterReplacement(s, k)
		fmt.Println(res)
	}
	{
		var word string = "035985750011523523129774573439111590559325"
		res := string_issue.NumDifferentIntegers(word)
		fmt.Println(res)
	}
	{
		var num1 string = "99"
		var num2 string = "1"
		res := number.AddStrings(num1, num2)
		fmt.Println(res)
	}
	{
		res := array.NumberOfArithmeticSlices([]int{1, 2, 3, 4, 8, 9, 10})
		fmt.Println(res)
	}
	{
		var num int = -26
		res := number.ToHex(num)
		fmt.Println(res)
	}
	{
		var stones []int = []int{0, 1, 2, 3, 4, 8, 9, 11}
		res := array.CanCross2(stones)
		fmt.Println(res)
	}
	{
		var n int = 12
		res := number.FindNthDigit2(n)
		fmt.Println(res)
	}

	{
		var n int = 7
		res := number.ReadBinaryWatch2(n)
		fmt.Println(res)
	}
	{
		input := "ababacb"
		k := 3
		res := string_issue.LongestSubstring(input, k)
		fmt.Println(res)
	}
	{
		var s string = "3[a2[c]]"
		res := string_issue.DecodeString2(s)
		fmt.Println(res)
	}
	{
		var nums []int = []int{1, 2, 3}
		res := number.LargestDivisibleSubset(nums)
		fmt.Println(res)
	}
	{
		var envelopes [][]int = [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
		res := number.MaxEnvelopes(envelopes)
		fmt.Println(res)
	}
	{
		var s string = "aA"
		res := string_issue.ReverseVowels(s)
		fmt.Println(res)
	}
	{
		var n int = 8
		res := number.IsPowerOfFour(n)
		fmt.Println(res)
	}
	{
		var num int = 5
		res := number.CountBits(num)
		fmt.Println(res)
	}
	{
		var words []string = []string{"abcd", "dcba", "lls", "s", "sssll"}
		//var words []string = []string{"a","abc","aba",""}
		res := string_issue.PalindromePairs(words)
		fmt.Println(res)
	}
	{
		nums := []int{0, 4, 2, 1, 0, -1, -3}
		res := array.IncreasingTriplet2(nums)
		fmt.Println(res)
	}
	{
		var nums []int = []int{2147483647, -2147483648, -1, 0}
		var lower int = -1
		var upper int = 0
		res := number.CountRangeSum(nums, lower, upper)
		fmt.Println(res)
	}
	{
		var n int = 3
		res := number.BulbSwitch(n)
		fmt.Println(res)
	}
	{
		var words []string = []string{"a", "aa", "aaa", "aaaa"}
		res := array.MaxProduct(words)
		fmt.Println(res)
	}
	{
		input := []int{3, 1, 5, 8}
		res := number.MaxCoins312(input)
		fmt.Println(res)
	}
	{
		//var n int = 6
		//var edges [][]int = [][]int{{3,0},{3,1},{3,2},{3,4},{5,4}}
		var n int = 2
		var edges [][]int = [][]int{{0, 1}}
		res := tree.FindMinHeightTrees(n, edges)
		fmt.Println(res)
	}
	{
		//[8,5,9,5,1,6,9]
		//[2,6,4,3,8,4,1,0,7,2,9,2,8]
		//20
		//[5,7,7,0,1,6,7,2,2,4,6,8,9,2,0,9,8,7,6,3,9,4,8,8,4,5,3,3,7,4,3,2,8,9,8,4,0,2,0,2,2,0,4,2,2,8,6,7,1,0,8,7,5,4,6,4,1,7,4,4,3,7,5,8,8,0,3,1,3,4,6,0,6,9,6,6,4,2,1,9,3,7,4,4,4,2,1,9,5,2,1,7,6,0,1,3,5,3,7,7]
		//[8,3,7,8,6,9,1,5,5,0,5,2,8,7,8,3,3,7,9,2]
		//100
		var nums1 []int = []int{5, 7, 7, 0, 1, 6, 7, 2, 2, 4, 6, 8, 9, 2, 0, 9, 8, 7, 6, 3, 9, 4, 8, 8, 4, 5, 3, 3, 7, 4, 3, 2, 8, 9, 8, 4, 0, 2, 0, 2, 2, 0, 4, 2, 2, 8, 6, 7, 1, 0, 8, 7, 5, 4, 6, 4, 1, 7, 4, 4, 3, 7, 5, 8, 8, 0, 3, 1, 3, 4, 6, 0, 6, 9, 6, 6, 4, 2, 1, 9, 3, 7, 4, 4, 4, 2, 1, 9, 5, 2, 1, 7, 6, 0, 1, 3, 5, 3, 7, 7}
		var nums2 []int = []int{8, 3, 7, 8, 6, 9, 1, 5, 5, 0, 5, 2, 8, 7, 8, 3, 3, 7, 9, 2}
		var k int = 100
		res := number.MaxNumber(nums1, nums2, k)
		fmt.Println(res)
	}
	{
		s := "()())()"
		res := string_issue.RemoveInvalidParentheses(s)
		fmt.Println(res)
	}
	{
		secret := "1123"
		guess := "0111"
		res := number.GetHint(secret, guess)
		fmt.Println(res)
	}
	{
		var pattern string = "abba"
		var s string = "dog dog dog dog"
		res := string_issue.WordPattern(pattern, s)
		fmt.Println(res)
	}
	{
		//[14,16,12,1,16,17,6,8,5,19,16,13,16,3,11,16,4,16,9,7]
		var nums []int = []int{14, 16, 12, 1, 16, 17, 6, 8, 5, 19, 16, 13, 16, 3, 11, 16, 4, 16, 9, 7}
		//var nums []int = []int{1,3,4,2,2}
		res := number.FindDuplicate(nums)
		fmt.Println(res)
	}
	{
		var nums []int = []int{1}
		array.MoveZeroes2(nums)
		fmt.Println(nums[0])
	}
	{
		var n int = 12
		res := number.NumSquares(n)
		fmt.Println(res)
	}
	{
		res := number.NthUglyNumber2(10)
		fmt.Print(res)
	}
	{
		var n int = 5
		res := number.IsUgly(n)
		fmt.Println(res)
	}
	{
		var n int = 38
		res := number.AddDigits(n)
		fmt.Println(res)
	}
	{
		var s string = "aabcb"
		res := string_issue.BeautySum(s)
		fmt.Println(res)
	}
	{
		var nums []int = []int{1, 2, 3, 4}
		res := number.ProductExceptSelf(nums)
		fmt.Println(res)
	}
	{
		var n int = 1
		res := number.IsPowerOfTwo(n)
		fmt.Println(res)
	}
	{
		input := " 3+5 / 2 - 3"
		res := number.Calculate227(input)
		fmt.Println(res)
	}
	{
		var n int = 10
		res := number.GetMoneyAmount(n)
		fmt.Println(res)
	}
	{
		//"3456237490"
		//9191
		var num string = "3456237490"
		var target int = 9191
		res := number.AddOperators(num, target)
		fmt.Println(res)
	}
	{
		courses := 2
		input := [][]int{{1, 0}}
		res := diagram.CanFinish(courses, input)
		fmt.Println(res)
	}
	{
		var s string = " 2-1 + 2  "
		res := number.Calculate224(s)
		fmt.Println(res)
	}
	{
		var s string = "abbaa"
		var t string = "cddcd"
		res := string_issue.IsIsomorphic(s, t)
		fmt.Println(res)
	}
	{
		var s string = "aagd"
		res := string_issue.ShortestPalindrome(s)
		fmt.Println(res)
	}
	{
		var k int = 2
		var prices []int = []int{2, 4, 1}
		res := array.MaxProfit(k, prices)
		fmt.Println(res)
	}
	{
		ss := strings.TrimLeft("00010", "0")
		fmt.Println(ss)
		res := number.LargestNumber([]int{3, 32, 34, 5, 9})
		fmt.Println(res)
	}
	{
		var nums []int = []int{-1, -1, 2147483647}
		res := number.MajorityElement1(nums)
		fmt.Println(res)
	}
	{
		var n int = 26
		res := number.ConvertToTitle(n)
		fmt.Println(res)
	}
	{
		input := []int{2, 3, -2, 4, -2, 4, 6, -9, 3}
		res := number.MaxProduct2(input)
		fmt.Println(res)
	}
	{
		input := []string{"4", "13", "5", "/", "+"}
		res := number.EvalRPN2(input)
		fmt.Println(res)
	}
	{
		//-1->5->3->4->0
		var l1 list_queue.ListNode
		l1.Val = 4
		var l2 list_queue.ListNode
		l2.Val = 2
		var l3 list_queue.ListNode
		l3.Val = 1
		var l4 list_queue.ListNode
		l4.Val = 3
		var l5 list_queue.ListNode
		l5.Val = 0
		l1.Next = &l2
		l2.Next = &l3
		l3.Next = &l4
		l4.Next = &l5
		res := list_queue.InsertionSortList2(&l1)
		fmt.Println(res.Val)
	}
	{
		var l1 list_queue.ListNode
		l1.Val = 1
		var l2 list_queue.ListNode
		l2.Val = 2
		var l3 list_queue.ListNode
		l3.Val = 3
		var l4 list_queue.ListNode
		l4.Val = 4
		var l5 list_queue.ListNode
		l5.Val = 5
		l1.Next = &l2
		l2.Next = &l3
		l3.Next = &l4
		l4.Next = &l5
		list_queue.ReorderList2(&l1)
	}
	{
		//[1,3,null,null,2]
		var t1 tree.TreeNode
		t1.Val = 4
		var t2 tree.TreeNode
		t2.Val = 2
		var t3 tree.TreeNode
		t3.Val = 1
		var t4 tree.TreeNode
		t4.Val = 3
		t1.Left = &t2
		t2.Right = &t3
		t3.Right = &t4
		tree.RecoverTree(&t1)
		fmt.Println(t1)
	}
	{
		var s string = "baba"
		res := string_issue.MinCut(s)
		fmt.Println(res)
	}
	{
		//"red"
		//"tax"
		//["ted","tex","red","tax","tad","den","rex","pee"]
		var beginWord string = "red"
		var endWord string = "tax"
		var wordList []string = []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"}
		res := string_issue.FindLadders(beginWord, endWord, wordList)
		fmt.Println(res)
	}
	{
		var prices []int = []int{3, 3, 5, 0, 0, 3, 1, 4}
		res := number.MaxProfit(prices)
		fmt.Println(res)
	}
	{
		var prices []int = []int{7, 1, 5, 3, 6, 4}
		res := number.MaxProfit1(prices)
		fmt.Println(res)
	}
	{
		var s string = "adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc"
		var t string = "bcddceeeebecbc"
		res := string_issue.NumDistinct(s, t)
		fmt.Println(res)
	}
	{
		var isWater [][]int = [][]int{{0, 0}, {1, 1}, {1, 0}}
		res := array.HighestPeak(isWater)
		fmt.Println(res)
	}
	{
		var l1 list_queue.ListNode
		l1.Val = 1
		var l2 list_queue.ListNode
		l2.Val = 2
		var l3 list_queue.ListNode
		l3.Val = 3
		var l4 list_queue.ListNode
		l4.Val = 4
		var l5 list_queue.ListNode
		l5.Val = 5
		l1.Next = &l2
		l2.Next = &l3
		l3.Next = &l4
		l4.Next = &l5
		var left int = 2
		var right int = 4
		res := list_queue.ReverseBetween(&l1, left, right)
		fmt.Println(res)
	}
	{
		var s string = "10"
		res := string_issue.NumDecodings(s)
		fmt.Println(res)
	}
	{
		var l1 list_queue.ListNode
		l1.Val = 1
		var l2 list_queue.ListNode
		l2.Val = 4
		var l3 list_queue.ListNode
		l3.Val = 3
		var l4 list_queue.ListNode
		l4.Val = 2
		var l5 list_queue.ListNode
		l5.Val = 5
		var l6 list_queue.ListNode
		l6.Val = 2
		l1.Next = &l2
		l2.Next = &l3
		l3.Next = &l4
		l4.Next = &l5
		l5.Next = &l6
		res := list_queue.Partition(&l1, 3)
		fmt.Println(res.Val)
	}
	{
		input := "23"
		res := diagram.LetterCombinations(input)
		fmt.Println(res)
	}
	{
		//head = [1,2,3,3,4,4,5]
		var l1 list_queue.ListNode
		l1.Val = 1
		var l2 list_queue.ListNode
		l2.Val = 2
		var l3 list_queue.ListNode
		l3.Val = 3
		var l4 list_queue.ListNode
		l4.Val = 3
		var l5 list_queue.ListNode
		l5.Val = 4
		var l6 list_queue.ListNode
		l6.Val = 4
		var l7 list_queue.ListNode
		l7.Val = 5
		var l8 list_queue.ListNode
		l8.Val = 5
		l1.Next = &l2
		l2.Next = &l3
		l3.Next = &l4
		l4.Next = &l5
		l5.Next = &l6
		l6.Next = &l7
		l7.Next = &l8
		res := list_queue.DeleteDuplicates2(&l1)
		fmt.Println(res.Val)
	}
	{
		//var nums []int = []int{0,0,1,1,1,2,2,3,3,4,4}
		//var nums []int = []int{0,1,2,3}
		//res := array.RemoveDuplicates(nums)
		//fmt.Println(res)
	}
	{
		board := [][]byte{{'A'}}
		word := "A"
		res := array.Exist(board, word)
		fmt.Println(res)
	}
	{
		var n int = 4
		var k int = 2
		res := array.Combine(n, k)
		fmt.Println(res)
	}
	{
		s := "a"
		t := "aa"
		res := string_issue.MinWindow(s, t)
		fmt.Println(res)
	}
	{
		input := [][]int{{1, 2, 3}, {4, 5, 6}}
		res := array.MinPathSum2(input)
		fmt.Println(res)
	}
	{
		var obstacleGrid [][]int = [][]int{{0, 1}, {0, 0}}
		res := array.UniquePathsWithObstacles(obstacleGrid)
		fmt.Println(res)
	}
	{
		var n int = 3
		res := array.GenerateMatrix(n)
		fmt.Println(res)
	}
	{
		var s string = "zzzzz"
		res := string_issue.CountHomogenous(s)
		fmt.Println(res)
	}
	{
		var s string = "11111"
		res := string_issue.MinOperations(s)
		fmt.Println(res)
	}
	{
		input := []int{3, 2, 1, 0, 4}
		res := array.CanJump(input)
		fmt.Println(res)
	}
	{
		var matrix [][]int = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
		res := array.SpiralOrder(matrix)
		fmt.Println(res)
	}
	{
		//"uuurruuuruuuuuuuuruuuuu"
		//"urrrurrrrrrrruurrrurrrurrrrruu"
		//"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
		//"bbbbbbbbbbbbbbbbbbbb"
		//"uuurr"
		//"urrru"
		var word1 string = "uuurr"
		var word2 string = "urrru"
		res := string_issue.LargestMerge(word1, word2)
		fmt.Println(res)
	}
	{
		var nums []int = []int{2, 1, 3, 4}
		res := array.Check(nums)
		fmt.Println(res)
	}
	{
		var nums []int = []int{1, 2, 3, 2}
		res := number.SumOfUnique(nums)
		fmt.Println(res)
	}
	{
		//var haystack string = "ababaabab"
		//var needle string = "abaabab"
		var haystack string = "hello"
		var needle string = "ll"
		res := string_issue.StrStr(haystack, needle)
		fmt.Println(res)
	}
	{
		//[1,2,4,8,16,32,64,128]
		//82
		var nums []int = []int{1, 2, 4, 8, 16, 32, 64, 128}
		var target int = 82
		res := number.ThreeSumClosest(nums, target)
		fmt.Println(res)
	}
	{
		//[1998]
		//1999
		//2000
		//2000
		var forbidden []int = []int{1998}
		var a int = 1999
		var b int = 2000
		var x int = 2000
		res := array.MinimumJumps(forbidden, a, b, x)
		fmt.Println(res)
	}
	{
		//A = [1,3,5,4], B = [1,2,3,7]
		//[0,7,8,10,10,11,12,13,19,18]
		//[4,4,5,7,11,14,15,16,17,20]
		var A []int = []int{0, 7, 8, 10, 10, 11, 12, 13, 19, 18}
		var B []int = []int{4, 4, 5, 7, 11, 14, 15, 16, 17, 20}
		res := array.MinSwap(A, B)
		fmt.Println(res)
	}
	{
		var maxChoosableInteger int = 10
		var desiredTotal int = 40
		res := number.CanIWin(maxChoosableInteger, desiredTotal)
		fmt.Println(res)
	}
	{
		var A []int = []int{4, 3, 2, 6}
		res := array.MaxRotateFunction(A)
		fmt.Println(res)
	}
	{
		var s1 string = "sea"
		var s2 string = "eat"
		res := string_issue.MinDistance(s1, s2)
		fmt.Println(res)
	}
	{
		var s string = "eleetminicoworoep"
		res := string_issue.FindTheLongestSubstring(s)
		fmt.Println(res)
	}
	{
		//head = [1,2,3,4,5], k = 2
		var l1 list_queue.ListNode
		l1.Val = 1
		var l2 list_queue.ListNode
		l2.Val = 2
		var l3 list_queue.ListNode
		l3.Val = 3
		var l4 list_queue.ListNode
		l4.Val = 4
		var l5 list_queue.ListNode
		l5.Val = 5
		l1.Next = &l2
		l2.Next = &l3
		l3.Next = &l4
		l4.Next = &l5
		var m int = 1
		res := list_queue.SwapNodes(&l1, m)
		fmt.Println(res.Val)
	}
	{
		var n int = 2
		res := array.ConstructDistancedSequence(n)
		fmt.Println(res)
	}
	{
		var s string = "cdbcbbaaabab"
		var x int = 4
		var y int = 5
		res := string_issue.MaximumGain(s, x, y)
		fmt.Println(res)
	}
	{
		var deliciousness []int = []int{1048576, 1048576}
		res := number.CountPairs(deliciousness)
		fmt.Println(res)
	}
	{
		var grid [][]int = [][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}
		res := array.FindBall(grid)
		fmt.Println(res)
	}
	{
		var nums []int = []int{-3, -1, -2}
		res := array.MaxSubArray(nums)
		fmt.Println(res)
	}
	{
		var nums []int = []int{1, 12, 2, 6, 3, 4}
		res := number.TupleSameProduct(nums)
		fmt.Println(res)
	}
	{
		var s string = "a"
		var t string = "a"
		res := string_issue.CountSubstrings(s, t)
		fmt.Println(res)
	}
	{
		//4
		//[[1,3,2},{2,3,0},{1,0,3},{1,0,2]]
		//[[2,1},{3,0]]
		var n int = 4
		var preferences [][]int = [][]int{{1, 3, 2}, {2, 3, 0}, {1, 0, 3}, {1, 0, 2}}
		var pairs [][]int = [][]int{{2, 1}, {3, 0}}
		res := array.UnhappyFriends(n, preferences, pairs)
		fmt.Println(res)
	}
	{
		var nums1 []int = []int{2, 3, 6, 3, 5, 1, 1, 3, 4, 6, 6}
		var nums2 []int = []int{3, 3, 7, 7, 5, 2, 1, 2, 7, 6}
		res := number.NumTriplets(nums1, nums2)
		fmt.Println(res)
	}
	{
		var arr []int = []int{1, 2, 3, 10, 4, 2, 3, 5}
		res := array.FindLengthOfShortestSubarray(arr)
		fmt.Println(res)
	}
	{
		var s string = "10101"
		res := string_issue.NumWays2(s)
		fmt.Println(res)
	}
	{
		//[269826447,974181916,225871443,189215924,664652743,592895362,754562271,335067223,996014894,510353008,48640772,228945137]
		//3
		var position []int = []int{1, 2, 3, 4, 5}
		var m int = 5
		res := array.MaxDistance(position, m)
		fmt.Println(res)
	}
	{
		var s string = "((()))())()()()))))"
		res := string_issue.MinInsertions(s)
		fmt.Println(res)
	}
	{
		var n int = 3
		var edges [][]int = [][]int{{0, 1}}
		var succProb []float64 = []float64{0.5}
		var start int = 0
		var end int = 2
		res := diagram.MaxProbability(n, edges, succProb, start, end)
		fmt.Println(res)
	}
	{
		var s string = "111111"
		res := string_issue.NumSub(s)
		fmt.Println(res)
	}
	{
		var n int = 9
		var left []int = []int{5}
		var right []int = []int{4}
		res := array.GetLastMoment(n, left, right)
		fmt.Println(res)
	}
	{
		var n int = 1
		var k int = 4
		res := string_issue.GetHappyString(n, k)
		fmt.Println(res)
	}
	{
		var radius int = 1
		var x_center int = 0
		var y_center int = 0
		var x1 int = 1
		var y1 int = -1
		var x2 int = 3
		var y2 int = 1
		res := diagram.CheckOverlap(radius, x_center, y_center, x1, y1, x2, y2)
		fmt.Println(res)
	}
	{
		var light []int = []int{1}
		res := array.NumTimesAllBlue(light)
		fmt.Println(res)
	}
	{
		//[1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]
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
		var t6 tree.TreeNode
		t6.Val = 6
		var t7 tree.TreeNode
		t7.Val = 7
		var t8 tree.TreeNode
		t8.Val = 8
		t1.Right = &t2
		t2.Left = &t3
		t2.Right = &t4
		t4.Left = &t5
		t4.Right = &t6
		t5.Right = &t7
		t7.Right = &t8
		res := tree.LongestZigZag(&t1)
		fmt.Println(res)
	}
	{
		var n int = 4
		//var edges [][]int = [][]int{{0,1,2},{0,4,8},{1,2,3},{1,4,2},{2,3,1},{3,4,1}}
		var edges [][]int = [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}
		var distanceThreshold int = 4
		res := diagram.FindTheCity(n, edges, distanceThreshold)
		fmt.Println(res)
	}
	{
		var s string = "CONTEST IS COMING"
		res := string_issue.PrintVertically(s)
		fmt.Println(res)
	}
	{
		var s string = "abeeadcdbceedccdabcd"
		var maxLetters int = 4
		var minSize int = 2
		var maxSize int = 3
		res := string_issue.MaxFreq(s, maxLetters, minSize, maxSize)
		fmt.Println(res)
	}
	{
		var mat [][]int = [][]int{{18, 70}, {61, 1}, {25, 85}, {14, 40}, {11, 96}, {97, 96}, {63, 45}}
		var threshold int = 40184
		res := diagram.MaxSideLength(mat, threshold)
		fmt.Println(res)
	}
	{
		//10
		//1000000000
		var low int = 10
		var high int = 1000000000
		res := number.SequentialDigits(low, high)
		fmt.Println(res)
	}
	{
		var characters string = "abc"
		var combinationLength int = 2
		res := diagram.Constructor1286(characters, combinationLength)
		fmt.Println(res)
	}
	{
		var products []string = []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
		var searchWord string = "mouse"
		res := string_issue.SuggestedProducts(products, searchWord)
		fmt.Println(res)
	}
	{
		var boxTypes [][]int = [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}
		var truckSize int = 10
		res := array.MaximumUnits(boxTypes, truckSize)
		fmt.Println(res)
	}
	{
		var grid [][]int = [][]int{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0},
			{1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}}
		res := diagram.ClosedIsland(grid)
		fmt.Println(res)
	}
	{
		var upper int = 2
		var lower int = 3
		var colsum []int = []int{2, 2, 1, 1}
		res := array.ReconstructMatrix(upper, lower, colsum)
		fmt.Println(res)
	}
	{
		var s1 string = "xxyyxyxyxx"
		var s2 string = "xyyxyxxxyx"
		res := string_issue.MinimumSwap(s1, s2)
		fmt.Println(res)
	}
	{
		var arr []string = []string{"abcdefghijklmnopqrstuvwxyz"}
		res := string_issue.MaxLength(arr)
		fmt.Println(res)
	}
	{
		var s string = "QQWWWWWWQQQQQWWWWWWW"
		res := string_issue.BalancedString(s)
		fmt.Println(res)
	}
	{
		var folder []string = []string{"/a/b/c", "/a", "/a/b/d"}
		res := string_issue.RemoveSubfolders(folder)
		fmt.Println(res)
	}
	{
		var d int = 1
		var f int = 6
		var target int = 3
		res := number.NumRollsToTarget2(d, f, target)
		fmt.Println(res)
	}
	{
		var arr []int = []int{4, 12, 10, 0, -2, 7, -8, 9, -9, -12, -12, 8, 8}
		var difference int = 0
		res := array.LongestSubsequence(arr, difference)
		fmt.Println(res)
	}
	{
		//"krpgjbjjznpzdfy"
		//"nxargkbydxmsgby"
		//14
		//"krrgw"
		//"zjxss"
		//19
		var s string = "krpgjbjjznpzdfy"
		var t string = "nxargkbydxmsgby"
		var maxCost int = 14
		res := string_issue.EqualSubstring(s, t, maxCost)
		fmt.Println(res)
	}
	{
		var customers [][]int = [][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}}
		res := array.AverageWaitingTime(customers)
		fmt.Println(res)
	}
	{
		var students []int = []int{1, 1, 1, 0, 0, 1}
		var sandwiches []int = []int{1, 0, 0, 0, 1, 1}
		res := array.CountStudents2(students, sandwiches)
		fmt.Println(res)
	}
	{
		var arr []int = []int{-5, 4, -4, -3, 5, -3}
		var k int = 3
		res := array.KConcatenationMaxSum(arr, k)
		fmt.Println(res)
	}
	{
		var arr []int = []int{-1, -3}
		res := array.MaximumSum(arr)
		fmt.Println(res)
	}
	{
		var s string = "abcda"
		var queries [][]int = [][]int{{3, 3, 0}, {1, 2, 0}, {0, 3, 1}, {0, 3, 2}, {0, 4, 1}}
		res := string_issue.CanMakePaliQueries(s, queries)
		fmt.Println(res)
	}
	{
		var text string = "aaaaa"
		res := string_issue.MaxRepOpt1(text)
		fmt.Println(res)
	}
	{
		var number string = "123456"
		res := string_issue.ReformatNumber(number)
		fmt.Println(res)
	}
	{
		var nums []int = []int{1, 2, 3, 4}
		res := array.MovesToMakeZigzag(nums)
		fmt.Println(res)
	}
	{
		var text1 string = "abc"
		var text2 string = "abcabc"
		res := string_issue.LongestCommonSubsequence(text1, text2)
		fmt.Println(res)
	}
	{
		var grid [][]int = [][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {1, 0, 0, 1}, {1, 1, 1, 1}, {1, 0, 1, 1}, {1, 1, 0, 1}}
		//var grid [][]int = [][]int{{1,1,1},{1,0,1},{1,1,1}}
		res := array.Largest1BorderedSquare(grid)
		fmt.Println(res)
	}
	{
		var books [][]int = [][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}
		var shelf_width int = 4
		res := diagram.MinHeightShelves(books, shelf_width)
		fmt.Println(res)
	}
	{
		var nums []int = []int{2, 3, 5}
		res := array.GetSumAbsoluteDifferences(nums)
		fmt.Println(res)
	}
	{
		var customers []int = []int{1}
		var grumpy []int = []int{0}
		var X int = 1
		res := array.MaxSatisfied(customers, grumpy, X)
		fmt.Println(res)
	}
	{
		var words []string = []string{"a", "ab", "ac", "bd", "abc", "abd", "abdd"}
		res := string_issue.LongestStrChain(words)
		fmt.Println(res)
	}
	{
		var instructions string = "GG"
		res := diagram.IsRobotBounded(instructions)
		fmt.Println(res)
	}
	{
		//[1,3,7,1,7,5]
		//[1,9,2,5,1]
		var A []int = []int{1, 3, 7, 1, 7, 5}
		var B []int = []int{1, 9, 2, 5, 1}
		res := diagram.MaxUncrossedLines2(A, B)
		fmt.Println(res)
	}
	{
		var A []int = []int{3, 8, 1, 3, 2, 1, 8, 9, 0}
		var L int = 3
		var M int = 2
		res := array.MaxSumTwoNoOverlap(A, L, M)
		fmt.Println(res)
	}
	{
		var A []int = []int{0, 0, 0, 0}
		res := array.LongestArithSeqLength(A)
		fmt.Println(res)
	}
	{
		var queries []string = []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}
		var pattern string = "FoBa"
		res := string_issue.CamelMatch(queries, pattern)
		fmt.Println(res)
	}
	{
		//[2,7,4,3,5]
		var l0 list_queue.ListNode
		l0.Val = 2
		var l1 list_queue.ListNode
		l1.Val = 7
		var l2 list_queue.ListNode
		l2.Val = 4
		var l3 list_queue.ListNode
		l3.Val = 3
		var l4 list_queue.ListNode
		l4.Val = 5
		l0.Next = &l1
		l1.Next = &l2
		l2.Next = &l3
		l3.Next = &l4
		res := list_queue.NextLargerNodes(&l0)
		fmt.Println(res)
	}
	{
		var A []int = []int{1, 2}
		res := array.MaxScoreSightseeingPair(A)
		fmt.Println(res)
	}
	{
		var weights []int = []int{1, 2, 3, 1, 1}
		var D int = 4
		res := array.ShipWithinDays(weights, D)
		fmt.Println(res)
	}
	{
		//[1,2,1,1,1,2,2,2]
		//[2,1,2,2,2,2,2,2]
		//A = [3,5,1,2,3], B = [3,6,3,3,4]
		var A []int = []int{1, 2, 1, 1, 1, 2, 2, 2}
		var B []int = []int{2, 1, 2, 2, 2, 2, 2, 2}
		res := array.MinDominoRotations(A, B)
		fmt.Println(res)
	}
	{
		//var X int = 2133
		//var Y int = 12321313
		//res := number.BrokenCalc(X,Y)
		//fmt.Println(res)
	}
	{
		var equations []string = []string{"c==c", "f!=a", "f==b", "b==c"}
		res := diagram.EquationsPossible(equations)
		fmt.Println(res)
	}
	{
		var A []int = []int{4, 5, 0, -2, -3, 1}
		var K int = 5
		res := array.SubarraysDivByK(A, K)
		fmt.Println(res)
	}
	{
		var n int = 2
		var k int = 0
		res := number.NumsSameConsecDiff(n, k)
		fmt.Println(res)
	}
	{
		var A []int = []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}
		res := array.MaxWidthRamp(A)
		fmt.Println(res)
	}
	{
		var tokens []int = []int{100, 200, 200, 400, 300, 400, 500}
		var P int = 200
		res := array.BagOfTokensScore(tokens, P)
		fmt.Println(res)
	}
	{
		var points [][]int = [][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {4, 1}, {4, 3}}
		res := diagram.MinAreaRect(points)
		fmt.Println(res)
	}
	{
		var A []int = []int{1, 0, 1, 0, 1}
		var S int = 2
		res := array.NumSubarraysWithSum2(A, S)
		fmt.Println(res)
	}
	{
		//["CBTInserter","insert","insert","get_root"]
		//[[[1,2,3,4,5,6]},{7},{8},{]]
		//var t1 tree.TreeNode
		//t1.Val = 1
		//var t2 tree.TreeNode
		//t2.Val = 2
		//var t3 tree.TreeNode
		//t3.Val = 3
		//var t4 tree.TreeNode
		//t4.Val = 4
		//var t5 tree.TreeNode
		//t5.Val = 5
		//var t6 tree.TreeNode
		//t6.Val = 6
		//t1.Left = &t2
		//t1.Right = &t3
		//t2.Left = &t4
		//t2.Right = &t5
		//t3.Left = &t6
		//obj := tree.Constructor919(&t1)
		//obj.Insert(7)
		//obj.Insert(8)
		var t1 tree.TreeNode
		t1.Val = 1
		obj := tree.Constructor919(&t1)
		obj.Insert(2)
		obj.Insert(3)
		obj.Insert(4)
		r := obj.Get_root()
		fmt.Println(r.Val)
	}
	{
		var word1 string = "llvlllvllqvvllvvqqvqqtqelvvvlvvvjlvvllvevllvlelvtqvlvtlvlvvlqvevlvllvmlctvvvfvqlvlqlvvvlllvvlvvqvvvllvqljvltillcjelllvvevlvvcevvlvvvqvqqlvlttvqfvvvqlvelvvvlvjvqllvlvvvvqqlvlllvlvvlevvqevvvvvvvvvtvqvvqlllvvtvvvleelqlqvlvvqcqvlallvqvvlvqylelvllvqlleqqjvvevvvlvltvvvqlvvvlvlqllvlvvvevvvjvvvvqvqqlkltvltetevvelqeqlnlqvllvvllvfvvltlvvvvlqlvlvlqvlvlqvlvvvelvqlvnfveeevlelllvvllvvvvvlvvlvvvvtqqllvvvvllelvlvevqvqvvvlllllqvqlqvevqvqtvvlvvvvjvlqvleqvvqlevvvlvltqjvvectqlvllveqtevltvllvtvqlvvvvllvvvvlelvtvvvllvlvlvvvvljvevvjevlvvvvllvvevelvvlvvjvelllvvvlellvevqvvllvevvqlvvvvtqvevlllvevevvvvlnvldllvlvvvvqlevllvlvvvlllvtvvlvltqvlevvvvvvtqvvellvllqvvlqvvlvelvvvrvvvveltvevvllvvlvvqvlevvvlqvvlqltelevvqllvvvvqtlivlvfvtvellvvvlvlqvlllyvvqlltlqevvqvvlvvlvtgllvlvvxvlvvlevvfvvlvelqvqevljqvltvvlvllllqqlevvvvevvllllllvvvllqllllvllveelvqtjlvllkvjvvqvvllvvqvqlvvevlvlvllvvvlvvvqvevvqlvlrqvelvceqqqvvqlvllvlllliqvnvqlexlllvlvvvvlvvlvltvvvvvlvlvvllotvqvgeveqllvvvvgvlevqlqllqvvflveleelvqlvlvvqvlvvtttvveqqvqvlvvqlvlvvlvllvellvllevevtellvqlvvelvvqvlvevlvyvtvqvvvqlvevevqyllllvqvllvlevqqlevvlqllvveevvvlvlvvvvqvvvvgvvvvllvlvqvvevvdlqeevvevvtlvllvlvteveqllvvlvqeclvvlvlvvelvtvvvlvqvlqvlvvllvevvqqtlvtflqlevllzqvvvvlvlvlvlvvllvjvjvlvvvvqlqvllvvvlvqvlelelqltvvvltlvvveqvvvflltlvllevelqvlvlvvlqvvlvqlvqvlqvlvvqlvtvlvvvvvlvvlflqvtvveqgvvvelvqlllvvlvevlvlvvtvvvlvlvqvvvlvvvlvlvlvvvtvvvlqvqvlvqvejvvlvvtqlvelvvqvvvllqjqqvvlqlvlllqvlclvvlvvqlllljlvvvllqvelvfvplvjeqelvvlvvvlevvqevklelltlevlqvvepllvvlvllveelvjlllvlvevvtlqlvvgveelvtvlvlvvvtvevevyqllvlvevllvvvvvvtvvtvvvvtlvqvvvvvevvelctvvlvlvlvvvvvvylqvvavvvvlvlnettvlvvlvvlletlvqvlcqlvvvvvvlvlvvvvvgqvellltevlvvfvvqjztvlvvlvlvlvqvllqvvvvvjvyvlvlflvvvvveevvvllqvlevvvvvvlqlvvellvvvvlqevvfvvletevvvqlqoqvqvvlllvvvvvlqlelevvvvvvvvvvvvllvlyqqeevllvlvvtvlvqvvvlvvvelvvqvvllevlvvvlvlvvvlljtqlvvvelvytvvqtiqvvvlllelvvlvlvvvlvngvlqvejlvqvvlvvllvevvvvvlvvevvvqvljlqvvtmlllllllevgvevqvlqvqvvelqlvtlvvvlvetlevlvllvqellljvvlvvvvlvvqelvvlllvvvtqllvpvvlvvlvlvvzvtlvqlvevfvlqvllvllqqelvtvlvlvllllvltvvqqqvfvvellvevqlvevvivvvvvvevltveivvvqvqtlvlvlqlvrqvevlvvlvvvfvevvvqvllvqqvetlvlvljvlllvlqvqllvellqvllvlqcvvvvljelqqvvplqvvlevlvllljvlvlelvqlvvltvlqvtllvclvvvvvvvvevlllvvvvllnvvvtijlvlvvvvqlvtelqvqlvlqvvllvlvqvvvlzpvlvqlqvveqlllvvlvqlvvvoqeeellvlllvvvqfvlqvelqveevqqvvvglvtvllvllvqleltvfqvvgfllvqvlqlvvlqvvlvvlqvlyvlllqvjvvlvqvvblvlvqvllqlvllllvvqvtlqvqvejvqlvllvvvvvvfvlvlvlvvvllttvvllvqlvvlvlllkevqveqlvjllvqvevllvlvvvlvlavvvelvelqllqevevvlvvllqllvveveevvvvlvvvlgvvllteltvvtvgvvltvevvlevtlvvvvvvlelllquqlelvlvvvlvtvqlvvllvvllvvlvlvvvltevvlqvvvveveclvqevlevvevvlvvvljlllvvllllvqqqvvlletqlcevvvlvlvlllvellqvqvvzlvfvqlvlqqqveqllvvltvvelvqvvlvqlllvqlvvvvtllqjtllltlvvvvvtllvlvvqvllvlqvvvtvevqtlvlelvlqvvlvpqeqvqvqevlvllvlvqlveevvefllvlllllvleqevllvvllvvlvqveqqvvllvvalvlvqevvlvqvlvevzqvlqvvvvlqlvlqvevlqvvvvlllvvvvvvvvlvevrljlvvlvvltllqvlvvelvvqfqvvqyelvlvlvvvvlvvvvvvllvvlevvvvvqqlqelvllvvvqvvlvejllvvltqlqvlvqqqvvvvevveljivlvjvevlqvlgveqqvtvlqvvvelevvvvllvylvvcvvllvivqllftvvqlvlvqlvvlllqevvvllllvvvvvvfqvvillvftvlevqlqelvvvvvvlevvvfevlvvvqvllvlvvtllvqvvvvvvlvvfvvlltvleqvevlvllvvfvllvqvvqelllvelevvvvlvvvplollltlvclevvttveelllltvvegtvvqtvlvqvvelevvkevvtelvvevevtlvqvlqqneevlvqvlllklnlvvvellexvvlvjlvvvvvldljvvvvtvevlvqllevqllllllvfltlqqleqtevvvvevvltlltvlvvvgllllveldvlvvlvvglvvqlvllvvvvlvetvvlvvevvqvflvqvecvvevvvlvqqlvvlvlvlevlgellvvelvyvevtvvvvqvvvvvleqvvllvluvllefqqeqvejvlvvvvllvvvqlvvqtvvvvvvlvvvtvvvlevlvjlvvlevlqlltlvllvlvvvlvvllllvvtvvvllqqljvlkvjelqvlvllqvclllvdletlvlvvlqovvqlvvvltvvqvvlvvevvtvvvvllvqlleqqvvllvllllxglvlvqvlvvvlvvqvqvqvvveelqlvlqqllvqevfzvvqqvvlkvqqvllqvqvvvvvvllqqllvllvlqllllvzvvqvvaqflllqvvllvvlqllellelvvqvlvlllllellvvvlfevvlftqvllvvllvvljlvvlvveqqvlvqevtvlevvlvtvqqlvvvvltvvvvvvlvlllvvvvevlvvlvllqvvtlvviqvtllltvvvlvvvevvvlvlvvvvvvvvlvvluveqlllvvlleeqtelvtvllvvvvlvqlqtetlvvelvjevvvllvplvqvfjvtlivvvlfvtetvlvvlevvlqlvlvelvqellllvvvveqevevqvvvvllvvtvlvlqlvlllvlevevqlvqvcpqvqllvevvelvlvvlvcvjvlvvvlvlqvvlvvvvvlqezqlvelvellqvllgolivvvelvvvvvvvcvvlvvvvvvqvlvvvvvvvvvyljlveqvelvlqqlvtvlvqplvltvlvllqqltvvltvvqvvvvvvevvtlvevvllvllevvlvvvllleveqvvvvvvlvvvqvvlvvlmvvvlvvlvlvlvlvvvlllvvvvlllllcvllvvvlqvvqjlklvlvvvvqqlqqllvoevvqvlvvfvtvvlfvvvlllvlvvlvvqllvllvevlclvvtlvlvvlvllelvvvelevvvvlvvllvylvveldvvvjvvevtvfelvlevqvlhfvvlvlqlqvvvvqvlvvvivvvyvqvvvlvvttvvvlvvvllleelzvvvvvlltvvqlvvelvelqtvvqvvlllvllveqvlvvvqllevevvvqvvqlqvtgvvlvllvlvtvvllvvzevevlvvlvvlvjllvvvvtvvlvlvvvvtqlvlllevqvvlvyvvltvtvvlvvqvllqvvvtvvvvlvqvvlvlvlvlqclvllyvevejlvvvjlvlvlvvvlqteqlvlqlqlvevlevllqqzvvvlvlqvttvljvvtvlqlvqllvvvlvvqlqlgvvvvvvvvvvvvllljlvlllvvlvlfflvvvvvlllltlvelvqvvlqvlvtqlvevvvllvlegvqlvvvvlvlvvdllvtveteveltvqtvevvevlqlavtvlvvtvqljvlvslvllpvvlvvlvvlqeveevllvtvevtlqvvvvvvlvelvqlvlvvtlvltvvlqvtvvvlllvvvlvvivvvvevvlvvlqevvvvvevlqlqvvqvvlleevvvvvvvqvjlvfvvvlvvevvqvvvvlqvvllqvvtvllvvlvvlvlelvvqqllvvvevqrqvveqllvellvvvqvvvlvltlctlvvllvllqvqllvqvvvvvevlvvevvvktlvlvvgvvvzvqlvvvevevvlvlvelltlelvllveveellqvvlqtqvqltvvevvqlqvvllqtvlqvvelenejtvvevvvlvqvlvtleiejvvlvlyvqvvvvlvelvkqvvvtvvvllvlqlvqvtlvvlvvvlqvlllvvvcvevvvlvevlvevvvvlvvlttvyvqqvlltqvflvvvvvvlvlvvvvtlllvlvvzqvvveviqvellvvvevvlvvelqvvlqlvlvvvlglvvtvlevleevqtqllvvlvvvvvlvlvlvlqlvfllvlvlnvlqvqvqvlvvltvvlfqlvcvvelldtqvelvvllelqnvlllllvvvovvfclcfvlqllvlvvevllelvllvlvvtvlvfevvltvevlelvllltvvvlvvvfelvvqvvvvlvvqzvvllqvevqeqvvveevvvvvvvvlelvqvlkvveqvellvvlvlllfvveltvvlvvvvqeclqlevvlvevevtqvldevvvyvljlvvvlvqtvvvvvvlveqvvvqvvvvlveqllqltvvvvvtvqllvlvvvvlvvvvellllvvlllvelvlllvlqllqvvtvvvefvvvgejllllvvqeqelvlvvvvtqvlvqjvvlllvvqqvjvldvlvvvlvvvvlqlevqllqclqevlecvveevvvlvvvvvvvvvzvvvlvqvvqlvllcelvlvllvlvvlvlqlflvvvvlvevxvtevlllvvvqvlvvlvvvlqvqvvvvvvfllvvvvcqlllvlvvvllvqletlteotvvtlvtllvqlvvlllcgqqlvlvtqqltvtevvvvvvvlvlleqlvlvovvvvlvtvvqvvtlvvvqqclvllevlvvltqqqvlvvlvelvvttvlqvvvvlnvvtevlvtveqlvqlvtvlvelqvvlvveqqeqqvvvvvtllvvevvvlvvvqevvvvvvvvqvlvvllvvllqqlqltlqllqvvllevqlvllvvfevlllvvgvvvvvlvvlvvvqeleevvvvlvvevvvlvlevelvtvvlcevlttlvvvvvvelqvqqvqvllvvlvvqvvlvvvlllvqvvqvvuqevvtjjvvvevlqvvvvvvvllveljveflqlljqvtvtvvvvlqqevlvvvvvvlvvvvvcvvjvvvvqevqvqvvgljvqvqvjvcvvlevvvqqovvlvlvlllvlvvqvltvlvllvteelqlvlevvvvvlqllklvellelqvvjqjvvlfqlvvvvvqetvlvvvvvtvvvvvllvvlvqtlvvqlvtvvvqvvvvllgevvvevvvlvlvcveyvvjqvqvlvlllitdelvvlvvevllvlelvvvvvvvleqvevlvyvvvvvvvtlvqvvlllevvvvlvwllvlvllvllvlvevvlvvellvqlvellvvvjlfvglqvvvlvvjllqqqqvvvllllvqvvvqqelevvlvllvcvvldvlevlvvvnlllvvvvvgqelvvvqvvfvvteefvvvvlvvevvvvglvlfvveqtvveenelvevqvvtvlevvemlvvtvvvvlvlvlfvvlqlvtltlqvlvfqvvtllvvvllqvvvvqllvqvqvlvvlvltqeevqvqllvvvqvevvveellnvlvvvlyvqvlvqvvvlclvvvtltllqvvqvvvdqvllvvllsqvtlvvqellqvvlvtlvvqlqvvvlvqlqvvveelvvllvvvegvvvvvvvvaqvlvlvleeltvlvqvlvvevvlvtvvqevvtevllvlvvvvllvveqvlvlvlllvllvvlvtvlvqlvvtvvlvvlvvvvqvqlvqvlveetvvvvvevlletvvellqlvvqvvvllvvvevlvlvvlvevvlvvvqevlltvlllftvvvqvvlvcvvvvllvqlvvvvlqvllqlvflvqvvvqvqttlevltvvvlqqvvqvvvvvvvyvvveflvvqlvlvvlfqlelqvvlqlelvvtvelveevvlllevevvvlllvilvvlvvvlvlvlllvqvvvvvvilvlvqvvvvlltvillvveqlvvvllvvvqllvvzevevqvvvvevvvvlvlqvlqvlvltqlvllvvvvvvvqllllevlvvelnvvlevlvvqlvqvvllvqvvallvqvvllvvevvjllvvvvllevlvvlvlvvvvvllllvvvlqqlelvlqqvvqqlvqeqvvlvvtqveveqjllellyvlqllvvvvlvvtlvvvvlvqtlqvlvfvllvqlvvvqvtvclvovllevlqvqlolvvlllqvvjellltlvvrqqvlllellllvvvqevvevqvtvlqlvqvvvlvvelelvtxveevelvtlvveeivvevevelllqtvelvvlllvqevvvvlllvlvevevvllvvvvrvvteevtvljqqvvelvlvvvvvvvvqlvvvqxvvflclevlqvelvvvqlellvlllveevlvlvvqlvnevevtlqllqlllvvevvvvvqvqvllvvvvvlqvvtvvvvtvrlvvlvelllvlvvvvvtvvlvvvtvlqvvevvelltvvtylvqlvqvqvllqvvtvqvellvqvlllvvllvvlvqvvevivqvqvvvvvvlvlvllvteqvzqlvllveillqqyeevtgqvvlvvlvvlvittllqtvtvvvvevlvllqevltvtelqveyvelqvqlevevvlllqvvvtlvlvlllllqvqlveqvvvvelvvvvvvtyeqllevvqltvqvlqvwgvvvlvelqvlvlcvvvvlllqdvlvlvlevqvvqfvrlvtvvvlvlvvlevvvvvqlvlvvvlvvqvlvvvvvqvvivlvvevvvlqlqlvqvlleqvvltvvvllvlllvvvvevvqvllllvjqqvevlvvlvivlltlvelvvlvvqlvlvlllrvvvevqvvivqvevevv"
		var word2 string = "wwswzppwwpnppnowpwpowpwpzmwzwwwpmzwwwwpmpppnpzpwwwtpwwspwjwwppwzpwpwpmpwnpwmnwwmwmwwpppwzmwzwpwppnwwpnnmcpowpwwwpwpwwmpwwppppwmnpwpmwpnpwmwppwwwnwwwwsmnwpwwwfwmpmwwnpwwwwppppwmwnwpwwwppgwpwwnpwmjpwwwwpmwpowwmmwwwwwwmjpwpwzpwpwwppzpppwwowpwwvmwwwzwpwpmnppwpwmwppwrpzwpwwptwwmpwwwpwnjpntmmpwwwzwzwwwmwvnzwpppwwwzzpwpwwppwmwznowpwwpmpmppwwwpwwzzzzwjpmmwppvwmzwwwmnmnmpmwpmmzmwmwmznapzpwwwpnwwwwwpwjwwwwpppwpwnwmpwwpwmpzwwwzwwwzpwwwpwwwpppwwwwppwwwpwawwzpwwmwwpzpwpwpzppwmmpwwpprwppwwppnwwnwwwywppppwwwwzwmwwwwwppmplwppwppmwwwpzppwwznwppwmpwpmjppwwmpwwwwwwmpwmmmwwppwpzzzwwwwwwwpzpwpwmppppwznnnpwwwwzwwzzwmppwwnawwpzzppurwwwwmnwpwppwwpwpwwzwwwpppwwwppwwppwwwwwwpwpppmmwwpwppwwmmpmwzpwamwmwampmwzmmpwwpzpwwmwppmwmpwpwwnwzmnwowpzpwmzwwppmwwnwpzpwwwpmppmwmwwwwwwwpwmapmmzmzwwpwwpmwwwwwmwwwwwpwmpzmwwplprppwbwwppwpppwwwwzzwpmmnppwpmwzwnzawwnpwwwwppmzmwwwpwpwpwwnpwwwwwppwpppppwwpwnapzwwnwppwmwnwpwwwwwwwwpwwppmpzwwwwwpppppzwpwwwpwwppwwmnwwpwwpmpwmowwpwwxwwwwpwwawzywzzwwzwwwwzpmpwwwwwzmwmpwwpwwmwzpprpppppppwwpkwzwppnwrwmwpwpwpwwwppwwppwzppwwwwzewnwmmwwwpzmpymwwoznzwwzpmwpwwwpzwwwppwmpmwprwmnzwwjwwpzwwanwwaamwwwmpmjwwpwwmwomwwpwpwwpzrnwpzwmwwwwxwwwmwwwwwpmpnwzwwmmwwwwoywwmyppwwwppmpmzwwmmmpwnwzmzppzzpwpwwwwwpzwnwwwmzzwwzwpwwnmwnppwwppwwzwznpwwmmwpwpmppwmzpwmppwmwppnwppppzmwmwwwpwzpwwzwpznpwswpzpppappwzpzwpmwpkwppwzwwtwmwpzvmpwmvpmpwmapjrwpwpnwwzzzpwpwwmnpwjwwppwpzwwwawpwwwwwwwapnzapppwpzpepppwwpwmwpppwwwwwwpwmwwmwwwwpwwpzpwpnmppppmwzwwwppmmwzzwnwwwwzzppzzwwwwwpowpnwzwzwwwwwpwpppwmwwwzprmpwwwmpwmpwzpwwwwwwwwwwpwnzwpzpwwrmanpwpwpwppwwwzwwwpwmmwwppwmwwpympppmwpwpzpwpwwppwmwwwzwwmwmppwwpwmwpwwwpwwpwwnmpzwpwppppzzpwnwwpwwmgwnwwwppwwzpnpwwwpmmwpwpzrzwwwpppwpnwmtjwnwpwmwwwwwpmwwmpzzawwpwmzwwpppwppzmppwwrwwppmmwwwpwwwwwzppwtzzpmpzwzmwwyppwwwzwwpppwwpwpzwwwmpppnzwwwpzzpwwwpwpwwwjpwwpwzwtwpwwwwmpwpppwpwwwwlpwpwpwwpwrwwwzwwwwmmpppwwwwkwnpzppmwwpwwwmpmmwmwwpwpwwwpwmwwpwwpmwppppnpwpzwpwmzwmpwpwpzpwwmawmpwpwwppwzspwwnmwpwpwwpnwwwhppwjwwpwawwwwpzwpmwwzpwpwwnwwwzpwpwwmepwwppwwwwkpwwwwwppwwwynpwpwmwwpmwmppwrpwmwppppmwwpwazwwwpwwnpppwwwwzwppmmwwwwpmwzpmwpwowwzmwwwwwwwwpwppwwwpwwwpwwmpwzwzzpmpppwrwwpwnwwwwppwpwwwpzwpmmwwwbowwwpwwpppwpwwwpwwpwwwwwnzwmwmpprwppwpazwpmwwzppwpwmzpmpzmwwpwpnpwwwwzzmzwznwmwznpzwzzwwwppwwpmwwzpmwwpwwpmzwwwwwpnfwawmwwwpwznzpzmwwpwwwznmwwwwmwwwkppwpwzwpwpwwpwwpwwwnwpwwmwplwwpwppwmnpwpppwmpwwwwpmpwwmppwrmpwwpppwppwzwpnqwwpwpppwwnpwwwppwnwprmpwnmwmwppzpnpwwppwzwmwwwwwmswwnwnwwwwrmwnsmwwzwpmpnnwwwpmpwmwzpplzvwwwwwwpzwwpwwnwmpwwwwnwapwwwwwwwwmppwmwwwwwmprwwpzwwmzwmwwnwnwwpwwwnnpwwppwpwwwgrwpwozwpwwpwppwzpmwpwwwwswpwzpnpwwppppwwpwwwwwjwspmwowpzpwwwwnmwnwtwpwwwppzzmpppwzoawwpzmpppmmmmmwwpzmpwwwwwwmwwmzwmwowwnnwzwpmwwwkwmzwzwmwwpmnwwwpzmwnwwwwwmwmwwwpmwzppwpwzwzzwwpppywwpwwwwpwwwpzwzwwnwwwwmwzwzpwzwzpwwwztampnwpwmwwmwppwwwpwzwwwppppwzpwzwzpwnwwwwpwwampwwzpzpwwwpmwwwppzpzwzwmmwwwszpppwwjwppwpmnwwzmwpzwzwwwwmwwwwwpwzwzpwppmwwwwwpnwwpwpwwmwmmwpwmmzwwzwwrwwpwawpewppwmmwwpwwwszewkfwjwpwpppwpwmwpwwwzpwpmwtwppwwazwpwmwmonwmpwwtpwwpowwzwzwpppwnwzwwzkwwwzpwnmmvwnoprwpmwwpwwpwpwwwpwpwjwwwwpwwtzpzwpwwzpzwwwwnwrnwwrpwwwnpwwwmpwmpwmpwmowwpppwzmmwwwwmwpmpwmpzwwpwpwnpnpwppwpwnnwwwpwwqwpwwvppwwpwmwzmwwpppzwwmpppwpwpzwpwpwmpwmzmpzmpwwpmppwwwzpwrwwwwwpwmpwmpwwpwmapwmwpfmnmpwwmwzwmwwzapwwwmnmpwwpwpgnwopwmzmpwmwpwppzmwwwppwwmwmpwpppwwwpmpwwwwwapapptwwwwzpzwwzwwpwwwwppppwpwwwzppmbpnwmwpwwwwpwmwmmpppzwpzwwwwwwtwtnzzwpwpmpwowwypwpmpnpzkwwpzzmwznppzwpppwjppnpmwwpwpwpwwwnzmpwwpwjwmmnpawppwpzwwwprpmwrwppppwwwwmpmwpwpwwpwwppnawppnnppwpwpwnzwppppwpwppprwwwmpwwwwpjpwwzmwppnwppwwppwpwppwwwwpwwpwwrwwwrzmwpwpmwwzzwpwwsppwpmmwppwpznppwwzcpppmpwwwnwwwwwzpzwpwptpwzpmwwwwwpwwwppwwwppwwwpzzzwwzwppwwprpwwwwwamwpwtppmyzwspwwwwzzpwprwwwpwwwwwwpppmpwwmwzmwmpmwppzppkwmpwpwnnpwmzppwwpwwwpawzpwnpwwmwznmzpwpwpwwpwtwpwwmwwnmpwwppzwmwmwwzpwwppwwwpnzzpwwoprwbnpwmwmzwwwpnmwpmwpwpppmwzwmwwwwmpwjpmwzzlmnwmpwwwpmnwmwwpwppmpwwpzwwpwzpwppwjpwppmwwepwmwwtpwwwwwpwwwwpmppwnwawwrpnmppwwwwzwwwwwppwwwwwwmzzppwwwwwnwgpwppwwpmmwkpwzwwgnppwzwpwwpwwwwswlpppwnwwwwwwpwwnmwmwnzppmwpwpezwpwmnppnnpwwzwmzpnpwwppzwwwrmzwpwpwjwpppwmwnapwzapwpwzhwwmwwzwzwwppnpwwnmwppwpmwwpwwmwwwpwpppwwmwmppmpppwpwppwwwpwmwwwpmwwwppwpmwpswppwpzpwpwwpppwwabwzpnwpnnwmwwwppwwwwgwzowwmwwwpspmwwwwmnwpzwwwwppbpwpwwpwwzwwwpwppmwwmwzapwwzpnwpwwmwpwzwownpmwzpwppzpwpwwewwppwwpwwpmpwwwwpwwzwpwzppwpwwwpzpwwwwpwwmmpwmpppwpwwwwpwzpwwpmnmwmlpwpwwmwpwpnzwmwmpwwrwwpwzrwwpwpwpznwwzwwmpmmamwnpwwmpwmpwwnzppwnmpmwzwwwpznpwpwwpwwwzmpwnwnmpwzwwwppwpzwpnrmppspwwpwmppzpptmwmwpwwpnwpppawmwwwzwnpmwmpwzwwzawjwwpwiwpnwmpwwwwwwwwwwmmwnwmwppwmwwpznwpwmznpzpzpmpwkwmpwwwppwzwpwpnwwpwwmzmwwpwwwwppewwzzwpwpwwpwmwwpmppzwwwwpmpmwwwwpwwmwwwwwwpmmwwwnwwpwzwzwrknwapzwbzmwwawmpokzwwnwwwzpzwwpnwpwwppwwppwzwpwwwwwwmnzwpmzwszwwzwswwpwwmpwmpwpnwozwwmzmwmwwmwwwzwpppmpwwmwywwprwwpwwwwwwwppzmwwwpuwmpwwppwwmwzpwrwwmwwpwwfpnzwwwopwwmwpzmwmpmzrzwmwwwpmnwpwwwwwppmbwwwmpzwwwzwwwwwwpwwwwzppwnwwmmpppwppwppwpwwpwtmwpwwzppppwrmwppwwewpwmpwpwszpwpzpppzpwrwmwpwawazwpwwpwmwwpwwpwprpwawwwwwnrmwpwpwwzzwppmwpmwzwwmnwwwwmmpwwwjwemawmwwwwwwlwppwwpnmwwnwwwmpwpwwpuwpwzwwwwwpnwwppznmwwpznnwwwawwwmwmnozwwwpwzwwpwzmwspwpwmmwpwmppwmwwwpwwmwpppppwwwpwppzwwnwwnpwpwpwmzwmwwwzwwwwwwzwmzwppwwwpwwmwspppwpzwzwmnwpzwmmppwnwpnwwwpwppxppwwpmwkwpppwmpwwwzzwwwpwwpwpppwwpppmpmwpppppznwwpwwwpwzwwppwwpwpznwwpwewwprnwwfwpmwwwpppppmpwnpwwwmmppwpppwwwzpzppwppmpwpmpowppwwwawpnrwwwzwmpwwzpnpppppwwwznwwwpppzwwwpwlmwzpwnwpwwwwwwpwwpwppwzzwpzwwwwamalwwwrwwppwzmwwwwmwpwwwwmnwjwppmpqnpzwpwwmwwpwpwnppppkpptowpmpwwmwpwppmwnzrpwwwwrwpawzwwmwwwwwpwmwwmppwpwpwwwppwwrwnpnwwwpwwwywpzwzwpwwwwwpmpmjmwwnqwpwwzppwnpwwpgfwpmzpmwpwwpppzpmppwwpppzmwmwzwmpwpwzwppwwwrwjwwpnwwwwpzwpmtpwmnnppwpnpwwwpwwppwwwpppwppwwwwwwwpwwmwwwzpwpwmzwwzwwzpwbwwnplpnzzmnwspmwwwpmwtzwpwzmwwmwwmpppwwwwpmwzmwpwwzpzppzwszwpwwwmwwwpwnpapwpwpmwppmwmpwwwwpapwwawwmtwzpywwzmwpopwpwpwpzmppmppppwmpmwwwzwppepopwmwawwpwwwwppwwewzpwzwpwwwwwozwmwpwpspwwpwpawwpwpwpwznwwpwwzpmpwpnwppwpwwppwwwwpwwpppwmwwwwpwpwpppwppwwwwpmwppwwpwzppwpmmwwzmwwppnwazzwpmtwwpnpwwpwmwwzzpmwpwwwwpmmwzvmnwwwwpwppppwmpwzwpwpwwwnwwzwnzpmwwwzwdwppwpmwzgwppzpzwmaswmpwzmpmzwwzrzmmwpwwwwwzwwppwwwwpapppwpppmnpwzppwwppwwwppzpzywmwmwwwzzwwmwmpnwppzwmwwwwpwpmzpppwwwzpwwzpwwwkppwppwwzwnnwwmwzpwppxmwpwwpwpwppwwwwwnpwwwwwwwnowwzpwwwwwwnpwwpwpzwwwwnwwwwppmappwpwopwwnpmpzwpwwpmmwwwwzwwwzmwwmpwprwwwwwpmwpppwpwpwpwpppwmpmmwzpwpwwwmwppwwzpwwwwpwmppwwnpmzwwwpwwwwmmmwwawwwnwwmpwawmppzwzzwwpwwppwzmnzpmpzzwwpwmnwwwwppmwwzpznpwppwwpzwwppwwwwnwmpwwpmmpwawpwwwpwrwwwpmpwwwwwwzpmwpwpwwwpppmpwzpwppnpmfwjrwpzwwpppnwwwmppppwwpwznpwpwwwwrzpwpwpmpwpzmmwppwwpwpgppzwmwpwpmwnwpwmwmwwwrwpppwwwwwzppwpawwrwwpmnzppmwwgwpppwwmomwmwwpwnppmwpwwwwpwwpwpwwrwwzwpmpwwpwkapwpwmzpppwmwwwwwzppwwpzxwwmwpwwwpwnswwpppwpwwpppwwwmwwzwzppwwwwmwwzwpzwwwwmwpwwwwwwpwmwapwwwpwmnpwwzzwmwpzzwpwpmwwanmwwnpwzwwmwwwwzpwwpwzwzpwppwpmppwnwwwwwpppwmpnpzzmpmwwwnmjxwpwwjpzwwwpppwwwwprwwwwppzwwwzpwmppwmwpwwwlppwpwwpppzwwwmmmmwpwppwwmmwwmwzwpwpwzwpwwpypawwzptpwmpwwppmpzwwmapwwwwwwwzwwrpswnwwzwwppwwpppzpwngnmmwmwwwwmpptpzwmwwpppwwwwwpwwwwwpwbpwxmmwwwgmmppmwpmwwwnwnpwwpwwmppzwwwwmawwjwpwwspwnppmwwwnwwpmzwpwwwpwrppmwpmwypnzpapzwwmwwppowpwwppzmpwwwmnwwwwwwwppwwwwwmwwwwpwwmwrppwwopwwwppmwwwwwpppmwwwwwwwwwwpwmwnpwwwwwppzwwmwwwpwwmwnwwommwpwwwpwwwwwwmwppnwmwwpwzppwawpwpwwzmwwwwzppppwnrwwpppwwpppppwwzwmwwwzpwwwwwmpwwwpwpwrmppwmazwnsmwwpppowwwwwjpzwmpwwwwwwpmzwwwzmppwwwpppwpwopwwmwwwwnmwzppwmnzppwwpppwwrmzwzwwwwzpwmpwpppwzwmpwpwwmppzwppwywpazwpwwmwwpjpwpwmwmpowmwwzwpppmwpwwwppmzppwppwwwwwmnpwwppzwwpwpwwppwwnzmwwwwwzzwpzpppppppwwwpwppwzpwwpnpppppwpwwwnznwppwppmpwwwppwppwpwwwzwwmppmwwpwwpwpnpwpmnwmwzpppwpzpmmpzwwpwwmmwwwwprppwwwpwwpwwpwwwzpwwwwppwwwwzwpzzwmpzpwpwwppmpwwmpwwpnwmwwwzpnpppwzmpwpwpppwzpwwpmwppmmzwpwwpwpmwwrwvnpwwpwwmpzprapppwpwwpwpppwwwzpwpwnwwwbzpwmpmppwwppwpmwpwnmzpwnwwwnwapwwpmwwzpwuwzwwwmwmzwnjpnwwmwwpwpppmawwwzpwpwwmzpwppwpwpmwzmwwwppwwmwwwpwzmnppwmwpnwwwomppw"
		res := string_issue.CloseStrings(word1, word2)
		fmt.Println(res)
	}
	{
		var s string = "aabbbbaabababbbbaaaaaabbababaaabaabaabbbabbbbabbabbababaabaababbbbaaaaabbabbabaaaabbbabaaaabbaaabbbaabbaaaaabaa"
		res := string_issue.MinimumDeletions(s)
		fmt.Println(res)
	}
	{
		var code []int = []int{2, 4, 9, 3}
		var k int = -2
		res := array.Decrypt(code, k)
		fmt.Println(res)
	}
	{
		var A []string = []string{"amazon", "apple", "facebook", "google", "leetcode"}
		var B []string = []string{"e", "oo"}
		//["facebook","leetcode"]
		res := array.WordSubsets(A, B)
		fmt.Println(res)
	}
	{
		var A []int = []int{1, 3, 6}
		var K int = 3
		res := array.SmallestRangeII(A, K)
		fmt.Println(res)
	}
	{
		obj := array.Constructor901()
		//[100},{80},{60},{70},{60},{75},{85]
		res := obj.Next(100)
		res = obj.Next(80)
		res = obj.Next(60)
		res = obj.Next(70)
		res = obj.Next(68)
		res = obj.Next(75)
		res = obj.Next(85)
		fmt.Println(res)
	}
	{
		var input []int = []int{3, 8, 0, 9, 2, 5}
		obj := diagram.Constructor900(input)
		res := obj.Next(2)
		res = obj.Next(1)
		res = obj.Next(1)
		res = obj.Next(1)
		res = obj.Next(1)
		fmt.Println(res)
	}
	{
		var A []int = []int{1}
		res := array.SubarrayBitwiseORs(A)
		fmt.Println(res)
	}
	{
		var R int = 5
		var C int = 6
		var r0 int = 1
		var c0 int = 4
		res := array.SpiralMatrixIII(R, C, r0, c0)
		fmt.Println(res)
	}
	{
		var people []int = []int{3, 5, 3, 4, 1, 3, 2, 5, 4, 7, 8, 9, 6, 5, 4, 3, 3, 2, 7, 8, 5}
		var limit int = 10
		res := number.NumRescueBoats(people, limit)
		fmt.Println(res)
	}
	{
		var piles []int = []int{30, 11, 23, 4, 20}
		var H int = 6
		res := number.MinEatingSpeed(piles, H)
		fmt.Println(res)
	}
	{
		var n int = 2 << 29
		res := number.ReorderedPowerOf2(n)
		fmt.Println(res)
	}
	{
		var S string = "(()(()))"
		res := string_issue.ScoreOfParentheses(S)
		fmt.Println(res)
	}
	{
		var S string = "zd"
		var shifts []int = []int{0, 1}
		res := string_issue.ShiftingLetters(S, shifts)
		fmt.Println(res)
	}
	{
		var arr []int = []int{2, 3, 2}
		res := array.LongestMountain(arr)
		fmt.Println(res)
	}
	{
		var s string = "17522"
		res := number.SplitIntoFibonacci(s)
		fmt.Println(res)
	}
	//{
	//	var heights []int= []int{4,2,7,6,9,14,12}
	//	var bricks int = 5
	//	var ladders int = 1
	//	res := array.FurthestBuilding(heights,bricks,ladders)
	//	fmt.Println(res)
	//}
	{
		var n int = 3
		res := number.CountVowelStrings(n)
		fmt.Println(res)
	}
	{
		var arr []int = []int{1, 3, 5, 7}
		var pieces [][]int = [][]int{{2, 4, 6, 8}}
		res := array.CanFormArray(arr, pieces)
		fmt.Println(res)
	}
	{
		var nums []int = []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}
		res := array.FrequencySort(nums)
		fmt.Println(res)
	}
	{
		//var difficulty []int = []int{2,4,6,8,10}
		//var profit []int = []int{10,20,30,40,50}
		//var worker []int = []int{4,5,6,7}
		var difficulty []int = []int{66, 1, 28, 73, 53, 35, 45, 60, 100, 44, 59, 94, 27, 88, 7, 18, 83, 18, 72, 63}
		var profit []int = []int{66, 20, 84, 81, 56, 40, 37, 82, 53, 45, 43, 96, 67, 27, 12, 54, 98, 19, 47, 77}
		var worker []int = []int{61, 33, 68, 38, 63, 45, 1, 10, 53, 23, 66, 70, 14, 51, 94, 18, 28, 78, 100, 16}
		res := array.MaxProfitAssignment(difficulty, profit, worker)
		fmt.Println(res)
	}
	{
		//var intervals [][]int = [][]int{{1,2},{3,5},{6,7},{8,10},{12,16}}
		//var newInterval []int = []int{16,18}
		var intervals [][]int = [][]int{{1, 5}}
		var newInterval []int = []int{0, 0}
		res := array.Insert(intervals, newInterval)
		fmt.Println(res)
	}
	{
		var A []int = []int{1, 2, 4}
		res := tree.NumFactoredBinaryTrees(A)
		fmt.Println(res)
	}
	{
		//var A []int = []int{9,1,2,3,9}
		//var K int = 3
		var A []int = []int{4, 1, 7, 5, 6, 2, 3}
		var K int = 4
		res := array.LargestSumOfAverages(A, K)
		fmt.Println(res)
	}
	{
		var poured int = 10
		var query_row int = 4
		var query_glass int = 3
		res := number.ChampagneTower(poured, query_row, query_glass)
		fmt.Println(res)
	}
	{
		var nums []int = []int{4, 6, 5, 9, 3, 7}
		var l []int = []int{0, 0, 2}
		var r []int = []int{2, 3, 5}
		res := array.CheckArithmeticSubarrays(nums, l, r)
		fmt.Println(res)
	}
	{
		var S string = "abcde"
		var words []string = []string{"a", "bb", "acd", "ace"}
		res := string_issue.NumMatchingSubseq(S, words)
		fmt.Println(res)
	}
	{
		var n int = 999
		var k int = 34
		res := diagram.NumberOfSets(n, k)
		fmt.Println(res)
	}
	{
		//var graph [][]int = [][]int{{1},{0,3},{3},{1,2}}
		var graph [][]int = [][]int{{}, {2, 4, 6}, {1, 4, 8, 9}, {7, 8}, {1, 2, 8, 9}, {6, 9}, {1, 5, 7, 8, 9}, {3, 6, 9}, {2, 3, 4, 6, 9}, {2, 4, 5, 6, 7, 8}}
		//var graph [][]int = [][]int{{},{2,3},{1},{1}}
		//var graph [][]int = [][]int{{1},{0},{4},{4},{2,3}}
		res := diagram.IsBipartite(graph)
		fmt.Println(res)
	}
	{
		var N int = 4
		var K int = 6
		res := number.KthGrammar(N, K)
		fmt.Println(res)
	}
	{
		var start string = "XXXRXXLXXXXXXXXRXXXR"
		var end string = "XXXXRLXXXXXXXXXXXXRR"
		res := string_issue.CanTransform(start, end)
		fmt.Println(res)
	}
	{
		var S string = "aaabbc"
		res := string_issue.ReorganizeString(S)
		fmt.Println(res)
	}
	{
		var A []int = []int{1, 2, 0}
		res := array.IsIdealPermutation(A)
		fmt.Println(res)
	}
	{
		var mat [][]int = [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}
		res := array.NumSubmat(mat)
		fmt.Println(res)
	}
	{
		var N int = 5
		var mines [][]int = [][]int{{0, 0}, {0, 3}, {1, 1}, {1, 4}, {2, 3}, {3, 0}, {4, 2}}
		res := diagram.OrderOfLargestPlusSign(N, mines)
		fmt.Println(res)
	}
	{
		var t1 tree.TreeNode
		t1.Val = 1
		var t2 tree.TreeNode
		t2.Val = 10
		var t3 tree.TreeNode
		t3.Val = 4

		var t4 tree.TreeNode
		t4.Val = 3
		var t5 tree.TreeNode
		t5.Val = 7
		var t6 tree.TreeNode
		t6.Val = 9
		t1.Left = &t2
		t1.Right = &t3
		t2.Left = &t4
		t3.Left = &t5
		t3.Right = &t6
		res := tree.IsEvenOddTree(&t1)
		fmt.Println(res)
	}
	{
		var rowSum []int = []int{3, 8}
		var colSum []int = []int{4, 7}
		res := array.RestoreMatrix(rowSum, colSum)
		fmt.Println(res)
	}
	{
		var nums []int = []int{3, 4, 2, 324, 432, 233, 2332, 2234, 7766, 5544, 3334, 3, 9, 88, 19}
		res := array.DeleteAndEarn2(nums)
		fmt.Println(res)
	}
	{
		var grid [][]int = [][]int{{1, -1, 0, -3, 4, 3, -3, 3, -1, 3, 0, 0, -4, 2}, {2, -2, -3, -4, 0, -2, -3, 3, 1, 4, 1, -3, -1, -4},
			{-4, 4, -4, -4, 2, -4, 3, 0, -2, -4, 3, 4, -1, 0}, {-3, 3, -4, -4, 3, 4, 4, 1, -1, -1, 0, 3, 4, 1},
			{1, 3, -4, 2, 2, -3, 1, -3, -4, -4, -1, -4, -4, 4}, {1, 1, -1, 1, -1, -1, 3, -4, -1, 2, -2, 3, -4, 0},
			{1, 0, 3, 3, 1, 4, 1, 1, -4, -1, -3, 4, -4, 4}, {4, 3, 2, 3, 0, -1, 2, -4, 1, 0, 0, 1, 3, 4},
			{-4, 4, -4, -4, 2, -2, 2, -1, 0, -2, 2, 4, -2, -1}, {-2, 3, 4, -4, 3, 3, -2, -1, 0, -3, 4, -2, -1, -4},
			{4, 3, 3, 3, -3, 1, 2, -4, -1, 4, -3, -3, 2, 0}, {3, 3, 0, 1, -4, -4, -3, 3, -2, -4, 2, 4, -3, 3},
			{-3, 0, 1, 3, 0, 0, 0, -4, -1, 4, -1, -3, 1, 1}, {-1, 4, 0, -3, 1, -3, -1, 2, 1, -3, -1, -4, 4, 1}}
		res := array.MaxProductPath(grid)
		fmt.Println(res)
	}
	{
		s := "aa"
		res := string_issue.MaxUniqueSplit(s)
		fmt.Println(res)
	}
	{
		//root = [1, 2, 3], k = 5
		var l1, l2, l3 list_queue.ListNode
		l1.Val = 1
		l2.Val = 2
		l3.Val = 3
		l1.Next = &l2
		l2.Next = &l3
		var k int = 5
		res := list_queue.SplitListToParts(&l1, k)
		fmt.Println(res)
	}
	{
		//[1,3,2,8,4,56,35,13,56,22,26,41,9]
		//2
		var prices []int = []int{1, 3, 2, 8, 4, 56, 35, 13, 56, 22, 26, 41, 9}
		var fee int = 2
		res := array.MaxProfit4(prices, fee)
		fmt.Println(res)
	}
	{
		var text string = "a"
		res := string_issue.ReorderSpaces(text)
		fmt.Println(res)
	}
	{
		var s = "aabaa"
		var cost []int = []int{1, 2, 3, 4, 1}
		res := diagram.MinCost(s, cost)
		fmt.Println(res)
	}
	{
		var arr []int = []int{1, 4, 2, 5, 3}
		res := array.SumOddLengthSubarrays(arr)
		fmt.Println(res)
	}
	{
		//["MagicDictionary", "buildDict", "search", "search", "search", "search"]
		//[[], [["hello","hallo","leetcode"]], ["hello"], ["hhllo"], ["hell"], ["leetcoded"]]
		obj := string_issue.Constructor()
		var dictionary []string = []string{"hello", "hallo", "leetcode"}
		obj.BuildDict(dictionary)
		res := obj.Search("hhllo")
		fmt.Println(res)
	}
	{
		var nums []int = []int{1, 2, 3, 1, 2, 3, 1, 2, 3}
		res := array.FindNumberOfLIS(nums)
		fmt.Println(res)
	}
	{
		var nums []int = []int{2, 2, 2, 2, 2}
		res := array.FindNumberOfLIS(nums)
		fmt.Println(res)
	}
	{
		var num int = 1993
		res := number.MaximumSwap(num)
		fmt.Println(res)
	}
	{
		var mat [][]int = [][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}
		res := array.NumSpecial(mat)
		fmt.Println(res)
	}
	{
		var n int = 6
		res := number.MinSteps(n)
		fmt.Println(res)
	}
	{
		//var dictionary []string = []string{"a", "aa", "aaa", "aaaa"}
		//var sentence string = "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"
		var dictionary []string = []string{"cat", "bat", "rat"}
		var sentence string = "the cattle was rattled by the battery"
		res := string_issue.ReplaceWords(dictionary, sentence)
		fmt.Println(res)
	}
	{
		var pairs [][]int = [][]int{{-10, -8}, {8, 9}, {-5, 0}, {6, 10}, {-6, -4}, {1, 7}, {9, 10}, {-4, 7}}
		res := number.FindLongestChain2(pairs)
		fmt.Println(res)
	}
	{
		var price []int = []int{2, 3, 4}
		var special [][]int = [][]int{{1, 1, 0, 4}, {2, 2, 1, 9}}
		var needs []int = []int{1, 2, 1}
		res := array.ShoppingOffers(price, special, needs)
		fmt.Println(res)
	}
	{
		//       4
		//     /   \
		//    2     6
		//   / \   /
		//  3   1 5
		//[4,2,6,3,1,5]
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
		var t6 tree.TreeNode
		t6.Val = 6
		t4.Left = &t2
		t4.Right = &t6
		t2.Left = &t3
		t2.Right = &t1
		t6.Left = &t5
		var n int = 1
		var d int = 2
		res := tree.AddOneRow(&t4, n, d)
		fmt.Println(res)
	}
	{
		var nums []int = []int{2, 2, 3, 4}
		res := array.TriangleNumber(nums)
		fmt.Println(res)
	}
	{
		var s string = "??yw?ipkj?"
		res := string_issue.ModifyString(s)
		fmt.Println(res)
	}
	{
		p1 := []int{0, 0}
		p2 := []int{-1, 0}
		p3 := []int{1, 0}
		p4 := []int{0, 1}
		res := diagram.ValidSquare(p1, p2, p3, p4)
		fmt.Println(res)
	}
	{
		var board [][]byte = [][]byte{{'B', '1', 'E', '1', 'B'},
			{'B', '1', 'M', '1', 'B'},
			{'B', '1', '1', '1', 'B'},
			{'B', 'B', 'B', 'B', 'B'}}
		var click []int = []int{1, 2}
		res := diagram.UpdateBoard(board, click)
		fmt.Println(res)
	}
	{
		var N int = 12
		res := array.CountArrangement(N)
		fmt.Println(res)
	}
	{
		var s string = "abpcplea"
		var d []string = []string{"b", "c", "a"}
		res := string_issue.FindLongestWord(s, d)
		fmt.Println(res)
	}
	{
		rows := 1
		columns := 2
		s := array.Constructor519(rows, columns)
		res := s.Flip()
		fmt.Println(res)
		res = s.Flip()
		fmt.Println(res)
		s.Reset()
		res = s.Flip()
		fmt.Println(res)
	}
	{
		var nums []int = []int{1, 0, 1, 99, 101}
		res := array.MinOperations2(nums)
		fmt.Println(res)
	}
	{
		var i int = 1234
		res := number.ThousandSeparator(i)
		fmt.Println(res)
	}
	{
		var s string = "bbbab"
		res := string_issue.LongestPalindromeSubseq(s)
		fmt.Println(res)
	}
	{
		var nums []int = []int{3, 2, 1}
		res := array.NextGreaterElements(nums)
		fmt.Println(res)
	}
	{
		//var nums []int = []int{3606449,6,5,9,452429,7,9580316,9857582,8514433,9,6,6614512,753594,5474165,4,2697293,8,7,1}
		//res := number.PredictTheWinner(nums)
		//fmt.Println(res)
	}
	{
		var input []int = []int{1, 1, 1, 2, 2, 3, 3, 3}
		res := number.Makesquare(input)
		fmt.Println(res)
	}
	{
		var s string = "123:123.123.123.123"
		res := string_issue.ValidIPAddress(s)
		fmt.Println(res)
	}
	{
		var s string = "zabcopabc"
		res := string_issue.FindSubstringInWraproundString(s)
		fmt.Println(res)
	}
	{
		//var nums []int = []int{0,0,0}
		//var target int = 0
		//res := array.MaxNonOverlapping(nums,target)
		//fmt.Println(res)
	}
	{
		var nums []int = []int{-1, -2, -3, -4, -5}
		res := array.CircularArrayLoop(nums)
		fmt.Println(res)
	}
	{
		//var s string = "sSA"
		//res := string_issue.MakeGood(s)
		//fmt.Println(res)
	}
	{
		var s string = "xsi"
		res := string_issue.OriginalDigits(s)
		fmt.Println(res)
	}
	{
		var nums []int = []int{}
		tree.Constructor307(nums)
		//var obj = tree.Constructor307(nums)
		//res := obj.SumRange(1,3)
		//obj.Update(6,8)
		//res = obj.SumRange(4,6)
		//fmt.Println(res)
	}
	{
		var citations []int = []int{0, 1, 3, 5, 6}
		res := array.HIndex2(citations)
		fmt.Println(res)
	}
	{
		A := -3
		B := 0
		C := 0
		D := 4
		E := 0
		F := -1
		G := 9
		H := 2
		res := diagram.ComputeArea(A, B, C, D, E, F, G, H)
		fmt.Println(res)
	}
	{
		version1 := "0.1"
		version2 := "1.1"
		res := string_issue.CompareVersion(version1, version2)
		fmt.Println(res)
	}
	{
		var arr []int = []int{1, 25, 35, 42, 68, 70}
		var k int = 2
		res := array.GetWinner2(arr, k)
		fmt.Println(res)
	}
	{
		var s string = "20000000000000000000"
		res := number.MyAtoi(s)
		fmt.Println(res)
	}
	{
		var path string = "/"
		res := string_issue.SimplifyPath(path)
		fmt.Println(res)
	}
	{
		var s1 string = "delete"
		var s2 string = "leet"
		res := string_issue.LCS_minimumDeleteSum(s1, s2)
		fmt.Println(res)
	}
	{
		s := "aacaba"
		res := string_issue.NumSplits(s)
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
		t1.Left = &t2
		t1.Right = &t3
		t2.Right = &t4
		var disatance int = 3
		res := tree.CountPairs(&t1, disatance)
		fmt.Println(res)
	}
	{
		var target string = "10111"
		res := number.MinFlips(target)
		fmt.Println(res)
	}
	{
		var s = "codeleet"
		var indices []int = []int{4, 5, 6, 7, 0, 2, 1, 3}
		res := string_issue.RestoreString(s, indices)
		fmt.Println(res)
	}
	{
		var dividend int = -2147483648
		var divisor int = -1
		res := number.Divide(dividend, divisor)
		fmt.Println(res)
	}
	{
		var year int = 2019
		var month int = 8
		var day int = 31
		res := number.DayOfTheWeek(day, month, year)
		fmt.Println(res)
	}
	{
		date1 := "2020-01-15"
		date2 := "2019-12-31"
		res := number.DaysBetweenDates(date1, date2)
		fmt.Println(res)
	}
	{
		var chips []int = []int{2, 2, 2, 3, 3}
		res := number.MinCostToMoveChips(chips)
		fmt.Println(res)
	}
	{
		var n int = 100
		res := number.NumPrimeArrangements(n)
		fmt.Println(res)
	}
	{
		var queries []string = []string{"bbb", "cc"}
		words := []string{"a", "aa", "aaa", "aaaa"}
		res := string_issue.NumSmallerByFrequency(queries, words)
		fmt.Println(res)
	}
	{
		var heights []int = []int{1, 2, 3, 4, 5}
		res := array.HeightChecker2(heights)
		fmt.Println(res)
	}
	{
		var N int = 4
		var paths [][]int = [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 3}, {2, 4}}
		res := array.GardenNoAdj(N, paths)
		fmt.Println(res)
	}
	{
		var input [][]int = [][]int{{1, 1}, {2, 2}, {3, 3}}
		res := number.IsBoomerang(input)
		fmt.Println(res)
	}
	{
		var a int = 3
		var b int = 5
		var c int = 1
		res := number.NumMovesStones(a, b, c)
		fmt.Println(res)
	}
	{
		var N int = 3
		res := number.DivisorGame(N)
		fmt.Println(res)
	}
	{
		var x int = 2
		var y int = 3
		var bound int = 10
		res := number.PowerfulIntegers(x, y, bound)
		fmt.Println(res)
	}
	{
		input := []int{0, 2, 0, 0}
		res := number.LargestTimeFromDigits(input)
		fmt.Println(res)
	}
	{
		input := "9,3,4,#,#,1,#,#,2,#,6,#,#"
		res := tree.IsValidSerialization(input)
		fmt.Println(res)
	}
	{
		n := 12
		primes := []int{2, 7, 13, 19}
		res := number.NthSuperUglyNumber2(n, primes)
		fmt.Println(res)
	}
	{
		input := "0235813"
		res := number.IsAdditiveNumber(input)
		fmt.Println(res)
	}
	{
		input := [][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}}
		obj := array.Constructor304(input)
		res := obj.SumRegion(2, 1, 4, 3)
		fmt.Print(res)
	}
	{
		input := []int{0, 0, 0, 0}
		target := 0
		res := number.FourSum(input, target)
		fmt.Println(res)
	}
	{
		input := []int{3, 2, 3}
		res := number.MajorityElement(input)
		fmt.Println(res)
	}
	{
		input := []int{0, 2, 3, 4, 6, 8, 9}
		res := array.SummaryRanges(input)
		fmt.Println(res)
	}
	{
		input := []int{2, 2}
		k := 3
		t := 0
		res := array.ContainsNearbyAlmostDuplicate(input, k, t)
		fmt.Println(res)
	}
	{
		k := 22
		res := number.FindMinFibonacciNumbers(k)
		fmt.Println(res)
	}
	{
		//["MyCircularQueue","enQueue","Rear","Rear","deQueue","enQueue","Rear",
		var obj list_queue.MyCircularQueue = list_queue.Constructor(8)

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
		k := 3
		n := 9
		res := number.CombinationSum3(k, n)
		fmt.Println(res)
	}
	{
		input := []int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240}
		res := array.Rob(input)
		fmt.Println(res)
	}
	{
		var obj tree.Trie = tree.Constructor208()

		obj.Insert("apple")
		res := obj.Search("apple")  // returns true
		res = obj.Search("app")     // returns false
		res = obj.StartsWith("app") // returns true
		obj.Insert("app")
		res = obj.Search("app") // returns true
		fmt.Println(res)
	}
	{
		obj := diagram.Constructor()
		obj.AddWord("and")
		obj.AddWord("abd")
		obj.AddWord("ab")
		res := obj.Search(".and")
		res = obj.Search(".bcd")
		fmt.Println(res)
	}
	{
		input := []int{9, 8, 7, 4, 3, 2, 1, 6}
		array.Recursive_qsort(input, 0, len(input)-1)
		fmt.Println(input)
	}
	{
		nums := []int{1, 2, 3, 4, 5}
		s := 11
		res := array.MinSubArrayLen2(s, nums)
		fmt.Println(res)
	}
	{
		input := "  hello   world!  "
		res := string_issue.ReverseWords(input)
		fmt.Println(res)
	}
	{
		//[[2,4},{1,3},{2,4},{1,3]]
		var node1 diagram.Node
		node1.Val = 1
		var node2 diagram.Node
		node2.Val = 2
		var node3 diagram.Node
		node3.Val = 3
		var node4 diagram.Node
		node4.Val = 4
		node1.Neighbors = append(node1.Neighbors, &node2, &node4)
		node2.Neighbors = append(node1.Neighbors, &node1, &node3)
		node3.Neighbors = append(node2.Neighbors, &node2, &node4)
		node4.Neighbors = append(node2.Neighbors, &node1, &node3)
		res := diagram.CloneGraph(&node1)
		fmt.Println(res.Val)
	}
	{
		input := []int{0, 1, 0, 1, 0, 1, 99}
		res := number.SingleNumber(input)
		fmt.Println(res)
	}
	{
		var list1 list_queue.ListNode
		list1.Val = 1
		var list2 list_queue.ListNode
		list2.Val = 2
		var list3 list_queue.ListNode
		list3.Val = 3
		var list4 list_queue.ListNode
		list4.Val = 4
		var list5 list_queue.ListNode
		list5.Val = 5
		//list1.Next = &list2
		//list2.Next = &list3
		//list3.Next = &list4
		//list4.Next = &list5
		list_queue.ReorderList(&list1)
		fmt.Println(list1.Val)
	}
	{
		input := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
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
		res := number.AngleClock(4, 50)
		fmt.Println(res)
	}
	{
		input := [][]int{{1, 6}, {4, 6}, {4, 8}}
		res := array.RemoveCoveredIntervals(input)
		fmt.Println(res)
	}
	{
		input := []int{2, 4, 7, 8, 9, 10, 14, 15, 18, 23, 32, 50}
		//res := number.LenLongestFibSubseq(input)
		res := number.LenLongestFibSubseq2(input)
		fmt.Println(res)
	}
	{
		res := number.NumOfWays(5000)
		fmt.Println(res)
	}
	{
		input := 3
		res := number.IntToRoman(input)
		fmt.Println(res)
	}
	{
		input := 10
		res := tree.GenerateTrees(input)
		fmt.Println(res)
	}
	{
		input := []int{2, 5, 7, 9}
		target := 16
		res := number.TwoSum2(input, target)
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
	//	input := []string_issue{ "//", "/ "}
	//	res := regionsBySlashes(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{4,6,7,7}
	//	res := findSubsequences(input)
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
	//	input := "cbcccbbccbbacbcaccbacabacacbcccbbccbbacb"
	//	res := longestPrefix(input)
	//	fmt.Println(res)
	//}
	//{
	//
	//	board := [][]byte{{'a'}}
	//	words := []string_issue{"a"}
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
	//	input := []int{4,2,0,3,2,5}
	//	res := largestRectangleArea(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := []int{2,0,1}
	//	res := countSmaller(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := []string_issue{"2", "1", "+", "3", "*"}
	//	res := evalRPN(input)
	//	fmt.Println(res)
	//}
	//{
	//	input := "ac"
	//	res := longestPalindrome(input)
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
	//	input := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	//	res := matrixBlockSum(input,1)
	//	fmt.Println(res)
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
	//	//[[2},{2,1},{1,1},{2,3},{4,1},{1},{2]]
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
	//	//wordDict := []string_issue{"cats", "dog", "sand", "and", "cat"}
	//	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	//	wordDict := []string_issue{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}
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
	//	wordList := []string_issue{"ricky","grind","cubic","panic","lover","farce","gofer","sales","flint","omens","lipid","briny","cloth","anted","slime","oaten","harsh","touts","stoop","cabal","lazed","elton","skunk","nicer","pesky","kusch","bused","kinda","tunis","enjoy","aches","prowl","babar","rooms","burst","slush","pines","urine","pinky","bayed","mania","light","flare","wares","women","verne","moron","shine","bluer","zeros","bleak","brief","tamra","vasts","jamie","lairs","penal","worst","yowls","pills","taros","addle","alyce","creep","saber","floyd","cures","soggy","vexed","vilma","cabby","verde","euler","cling","wanna","jenny","donor","stole","sakha","blake","sanes","riffs","forge","horus","sered","piked","prosy","wases","glove","onset","spake","benin","talks","sites","biers","wendy","dante","allan","haven","nears","shaka","sloth","perky","spear","spend","clint","dears","sadly","units","vista","hinds","marat","natal","least","bough","pales","boole","ditch","greys","slunk","bitch","belts","sense","skits","monty","yawns","music","hails","alien","gibes","lille","spacy","argot","wasps","drubs","poops","bella","clone","beast","emend","iring","start","darla","bells","cults","dhaka","sniff","seers","bantu","pages","fever","tacky","hoses","strop","climb","pairs","later","grant","raven","stael","drips","lucid","awing","dines","balms","della","galen","toned","snips","shady","chili","fears","nurse","joint","plump","micky","lions","jamal","queer","ruins","frats","spoof","semen","pulps","oldie","coors","rhone","papal","seals","spans","scaly","sieve","klaus","drums","tided","needs","rider","lures","treks","hares","liner","hokey","boots","primp","laval","limes","putts","fonda","damon","pikes","hobbs","specs","greet","ketch","braid","purer","tsars","berne","tarts","clean","grate","trips","chefs","timex","vicky","pares","price","every","beret","vices","jodie","fanny","mails","built","bossy","farms","pubic","gongs","magma","quads","shell","jocks","woods","waded","parka","jells","worse","diner","risks","bliss","bryan","terse","crier","incur","murky","gamed","edges","keens","bread","raced","vetch","glint","zions","porno","sizes","mends","ached","allie","bands","plank","forth","fuels","rhyme","wimpy","peels","foggy","wings","frill","edgar","slave","lotus","point","hints","germs","clung","limed","loafs","realm","myron","loopy","plush","volts","bimbo","smash","windy","sours","choke","karin","boast","whirr","tiber","dimes","basel","cutes","pinto","troll","thumb","decor","craft","tared","split","josue","tramp","screw","label","lenny","apses","slept","sikhs","child","bouts","cites","swipe","lurks","seeds","fists","hoard","steed","reams","spoil","diego","peale","bevel","flags","mazes","quart","snipe","latch","lards","acted","falls","busby","holed","mummy","wrong","wipes","carlo","leers","wails","night","pasty","eater","flunk","vedas","curse","tyros","mirth","jacky","butte","wired","fixes","tares","vague","roved","stove","swoon","scour","coked","marge","cants","comic","corns","zilch","typos","lives","truer","comma","gaily","teals","witty","hyper","croat","sways","tills","hones","dowel","llano","clefs","fores","cinch","brock","vichy","bleed","nuder","hoyle","slams","macro","arabs","tauts","eager","croak","scoop","crime","lurch","weals","fates","clipt","teens","bulls","domed","ghana","culls","frame","hanky","jared","swain","truss","drank","lobby","lumps","pansy","whews","saris","trite","weeps","dozes","jeans","flood","chimu","foxes","gelds","sects","scoff","poses","mares","famed","peers","hells","laked","zests","wring","steal","snoot","yodel","scamp","ellis","bandy","marry","jives","vises","blurb","relay","patch","haley","cubit","heine","place","touch","grain","gerry","badly","hooke","fuchs","savor","apron","judge","loren","britt","smith","tammy","altar","duels","huber","baton","dived","apace","sedan","basts","clark","mired","perch","hulks","jolly","welts","quack","spore","alums","shave","singe","lanny","dread","profs","skeet","flout","darin","newed","steer","taine","salvo","mites","rules","crash","thorn","olive","saves","yawed","pique","salon","ovens","dusty","janie","elise","carve","winds","abash","cheep","strap","fared","discs","poxed","hoots","catch","combo","maize","repay","mario","snuff","delve","cored","bards","sudan","shuns","yukon","jowls","wayne","torus","gales","creek","prove","needy","wisps","terri","ranks","books","dicky","tapes","aping","padre","roads","nines","seats","flats","rains","moira","basic","loves","pulls","tough","gills","codes","chest","teeny","jolts","woody","flame","asked","dulls","hotly","glare","mucky","spite","flake","vines","lindy","butts","froth","beeps","sills","bunny","flied","shaun","mawed","velds","voled","doily","patel","snake","thigh","adler","calks","desks","janus","spunk","baled","match","strip","hosed","nippy","wrest","whams","calfs","sleet","wives","boars","chain","table","duked","riped","edens","galas","huffs","biddy","claps","aleut","yucks","bangs","quids","glenn","evert","drunk","lusts","senna","slate","manet","roted","sleep","loxes","fluky","fence","clamp","doted","broad","sager","spark","belch","mandy","deana","beyer","hoist","leafy","levee","libel","tonic","aloes","steam","skews","tides","stall","rifts","saxon","mavis","asama","might","dotes","tangs","wroth","kited","salad","liens","clink","glows","balky","taffy","sided","sworn","oasis","tenth","blurt","tower","often","walsh","sonny","andes","slump","scans","boded","chive","finer","ponce","prune","sloes","dined","chums","dingo","harte","ahead","event","freer","heart","fetch","sated","soapy","skins","royal","cuter","loire","minot","aisle","horny","slued","panel","eight","snoop","pries","clive","pored","wrist","piped","daren","cells","parks","slugs","cubed","highs","booze","weary","stain","hoped","finny","weeds","fetid","racer","tasks","right","saint","shahs","basis","refer","chart","seize","lulls","slant","belay","clots","jinny","tours","modes","gloat","dunks","flute","conch","marts","aglow","gayer","lazes","dicks","chime","bears","sharp","hatch","forms","terry","gouda","thins","janet","tonya","axons","sewed","danny","rowdy","dolts","hurry","opine","fifty","noisy","spiky","humid","verna","poles","jayne","pecos","hooky","haney","shams","snots","sally","ruder","tempe","plunk","shaft","scows","essie","dated","fleet","spate","bunin","hikes","sodas","filly","thyme","fiefs","perks","chary","kiths","lidia","lefty","wolff","withe","three","crawl","wotan","brown","japed","tolls","taken","threw","crave","clash","layer","tends","notes","fudge","musky","bawdy","aline","matts","shirr","balks","stash","wicks","crepe","foods","fares","rotes","party","petty","press","dolly","mangy","leeks","silly","leant","nooks","chapt","loose","caged","wages","grist","alert","sheri","moody","tamps","hefts","souls","rubes","rolex","skulk","veeps","nonce","state","level","whirl","bight","grits","reset","faked","spiny","mixes","hunks","major","missy","arius","damns","fitly","caped","mucus","trace","surat","lloyd","furry","colin","texts","livia","reply","twill","ships","peons","shear","norms","jumbo","bring","masks","zippy","brine","dorks","roded","sinks","river","wolfs","strew","myths","pulpy","prank","veins","flues","minus","phone","banns","spell","burro","brags","boyle","lambs","sides","knees","clews","aired","skirt","heavy","dimer","bombs","scums","hayes","chaps","snugs","dusky","loxed","ellen","while","swank","track","minim","wiled","hazed","roofs","cantu","sorry","roach","loser","brass","stint","jerks","dirks","emory","campy","poise","sexed","gamer","catty","comte","bilbo","fasts","ledge","drier","idles","doors","waged","rizal","pured","weirs","crisp","tasty","sored","palmy","parts","ethel","unify","crows","crest","udder","delis","punks","dowse","totes","emile","coded","shops","poppa","pours","gushy","tiffs","shads","birds","coils","areas","boons","hulls","alter","lobes","pleat","depth","fires","pones","serra","sweat","kline","malay","ruled","calve","tired","drabs","tubed","wryer","slung","union","sonya","aided","hewed","dicey","grids","nixed","whits","mills","buffs","yucky","drops","ready","yuppy","tweet","napes","cadre","teach","rasps","dowdy","hoary","canto","posed","dumbo","kooks","reese","snaky","binge","byron","phony","safer","friar","novel","scale","huron","adorn","carla","fauna","myers","hobby","purse","flesh","smock","along","boils","pails","times","panza","lodge","clubs","colby","great","thing","peaks","diana","vance","whets","bergs","sling","spade","soaks","beach","traps","aspen","romps","boxed","fakir","weave","nerds","swazi","dotty","curls","diver","jonas","waite","verbs","yeast","lapel","barth","soars","hooks","taxed","slews","gouge","slags","chang","chafe","saved","josie","syncs","fonds","anion","actor","seems","pyrex","isiah","glued","groin","goren","waxes","tonia","whine","scads","knelt","teaks","satan","tromp","spats","merry","wordy","stake","gland","canal","donna","lends","filed","sacks","shied","moors","paths","older","pooch","balsa","riced","facet","decaf","attic","elder","akron","chomp","chump","picky","money","sheer","bolls","crabs","dorms","water","veers","tease","dummy","dumbs","lethe","halls","rifer","demon","fucks","whips","plops","fuses","focal","taces","snout","edict","flush","burps","dawes","lorry","spews","sprat","click","deann","sited","aunts","quips","godly","pupil","nanny","funks","shoon","aimed","stacy","helms","mints","banks","pinch","local","twine","pacts","deers","halos","slink","preys","potty","ruffs","pusan","suits","finks","slash","prods","dense","edsel","heeds","palls","slats","snits","mower","rares","ailed","rouge","ellie","gated","lyons","duded","links","oaths","letha","kicks","firms","gravy","month","kongo","mused","ducal","toted","vocal","disks","spied","studs","macao","erick","coupe","starr","reaps","decoy","rayon","nicks","breed","cosby","haunt","typed","plain","trays","muled","saith","drano","cower","snows","buses","jewry","argus","doers","flays","swish","resin","boobs","sicks","spies","bails","wowed","mabel","check","vapid","bacon","wilda","ollie","loony","irked","fraud","doles","facts","lists","gazed","furls","sunks","stows","wilde","brick","bowed","guise","suing","gates","niter","heros","hyped","clomp","never","lolls","rangy","paddy","chant","casts","terns","tunas","poker","scary","maims","saran","devon","tripe","lingo","paler","coped","bride","voted","dodge","gross","curds","sames","those","tithe","steep","flaks","close","swops","stare","notch","prays","roles","crush","feuds","nudge","baned","brake","plans","weepy","dazed","jenna","weiss","tomes","stews","whist","gibed","death","clank","cover","peeks","quick","abler","daddy","calls","scald","lilia","flask","cheer","grabs","megan","canes","jules","blots","mossy","begun","freak","caved","hello","hades","theed","wards","darcy","malta","peter","whorl","break","downs","odder","hoofs","kiddo","macho","fords","liked","flees","swing","elect","hoods","pluck","brook","astir","bland","sward","modal","flown","ahmad","waled","craps","cools","roods","hided","plath","kings","grips","gives","gnats","tabby","gauls","think","bully","fogey","sawed","lints","pushy","banes","drake","trail","moral","daley","balds","chugs","geeky","darts","soddy","haves","opens","rends","buggy","moles","freud","gored","shock","angus","puree","raves","johns","armed","packs","minis","reich","slots","totem","clown","popes","brute","hedge","latin","stoke","blend","pease","rubik","greer","hindi","betsy","flows","funky","kelli","humps","chewy","welds","scowl","yells","cough","sasha","sheaf","jokes","coast","words","irate","hales","camry","spits","burma","rhine","bends","spill","stubs","power","voles","learn","knoll","style","twila","drove","dacca","sheen","papas","shale","jones","duped","tunny","mouse","floss","corks","skims","swaps","inned","boxer","synch","skies","strep","bucks","belau","lower","flaky","quill","aural","rufus","floes","pokes","sends","sates","dally","boyer","hurts","foyer","gowns","torch","luria","fangs","moats","heinz","bolts","filet","firth","begot","argue","youth","chimp","frogs","kraft","smite","loges","loons","spine","domes","pokey","timur","noddy","doggy","wades","lanes","hence","louts","turks","lurid","goths","moist","bated","giles","stood","winos","shins","potts","brant","vised","alice","rosie","dents","babes","softy","decay","meats","tanya","rusks","pasts","karat","nuked","gorge","kinks","skull","noyce","aimee","watch","cleat","stuck","china","testy","doses","safes","stage","bayes","twins","limps","denis","chars","flaps","paces","abase","grays","deans","maria","asset","smuts","serbs","whigs","vases","robyn","girls","pents","alike","nodal","molly","swigs","swill","slums","rajah","bleep","beget","thanh","finns","clock","wafts","wafer","spicy","sorer","reach","beats","baker","crown","drugs","daisy","mocks","scots","fests","newer","agate","drift","marta","chino","flirt","homed","bribe","scram","bulks","servo","vesta","divas","preps","naval","tally","shove","ragas","blown","droll","tryst","lucky","leech","lines","sires","pyxed","taper","trump","payee","midge","paris","bored","loads","shuts","lived","swath","snare","boned","scars","aeons","grime","writs","paige","rungs","blent","signs","davis","dials","daubs","rainy","fawns","wrier","golds","wrath","ducks","allow","hosea","spike","meals","haber","muses","timed","broom","burks","louis","gangs","pools","vales","altai","elope","plied","slain","chasm","entry","slide","bawls","title","sings","grief","viola","doyle","peach","davit","bench","devil","latex","miles","pasha","tokes","coves","wheel","tried","verdi","wanda","sivan","prior","fryer","plots","kicky","porch","shill","coats","borne","brink","pawed","erwin","tense","stirs","wends","waxen","carts","smear","rival","scare","phase","bragg","crane","hocks","conan","bests","dares","molls","roots","dunes","slips","waked","fours","bolds","slosh","yemen","poole","solid","ports","fades","legal","cedes","green","curie","seedy","riper","poled","glade","hosts","tools","razes","tarry","muddy","shims","sword","thine","lasts","bloat","soled","tardy","foots","skiff","volta","murks","croci","gooks","gamey","pyxes","poems","kayla","larva","slaps","abuse","pings","plows","geese","minks","derby","super","inked","manic","leaks","flops","lajos","fuzes","swabs","twigs","gummy","pyres","shrew","islet","doled","wooly","lefts","hunts","toast","faith","macaw","sonia","leafs","colas","conks","altos","wiped","scene","boors","patsy","meany","chung","wakes","clear","ropes","tahoe","zones","crate","tombs","nouns","garth","puked","chats","hanks","baked","binds","fully","soaps","newel","yarns","puers","carps","spelt","lully","towed","scabs","prime","blest","patty","silky","abner","temps","lakes","tests","alias","mines","chips","funds","caret","splat","perry","turds","junks","cramp","saned","peary","snarl","fired","stung","nancy","bulge","styli","seams","hived","feast","triad","jaded","elvin","canny","birth","routs","rimed","pusey","laces","taste","basie","malls","shout","prier","prone","finis","claus","loops","heron","frump","spare","menus","ariel","crams","bloom","foxed","moons","mince","mixed","piers","deres","tempt","dryer","atone","heats","dario","hawed","swims","sheet","tasha","dings","clare","aging","daffy","wried","foals","lunar","havel","irony","ronny","naves","selma","gurus","crust","percy","murat","mauro","cowed","clang","biker","harms","barry","thump","crude","ulnae","thong","pager","oases","mered","locke","merle","soave","petal","poser","store","winch","wedge","inlet","nerdy","utter","filth","spray","drape","pukes","ewers","kinds","dates","meier","tammi","spoor","curly","chill","loped","gooey","boles","genet","boost","beets","heath","feeds","growl","livid","midst","rinds","fresh","waxed","yearn","keeps","rimes","naked","flick","plies","deeps","dirty","hefty","messy","hairy","walks","leper","sykes","nerve","rover","jived","brisk","lenin","viper","chuck","sinus","luger","ricks","hying","rusty","kathy","herds","wider","getty","roman","sandy","pends","fezes","trios","bites","pants","bless","diced","earth","shack","hinge","melds","jonah","chose","liver","salts","ratty","ashed","wacky","yokes","wanly","bruce","vowel","black","grail","lungs","arise","gluts","gluey","navel","coyer","ramps","miter","aldan","booth","musty","rills","darns","tined","straw","kerri","hared","lucks","metes","penny","radon","palms","deeds","earls","shard","pried","tampa","blank","gybes","vicki","drool","groom","curer","cubes","riggs","lanky","tuber","caves","acing","golly","hodge","beard","ginny","jibed","fumes","astor","quito","cargo","randi","gawky","zings","blind","dhoti","sneak","fatah","fixer","lapps","cline","grimm","fakes","maine","erika","dealt","mitch","olden","joist","gents","likes","shelf","silts","goats","leads","marin","spire","louie","evans","amuse","belly","nails","snead","model","whats","shari","quote","tacks","nutty","lames","caste","hexes","cooks","miner","shawn","anise","drama","trike","prate","ayers","loans","botch","vests","cilia","ridge","thugs","outed","jails","moped","plead","tunes","nosed","wills","lager","lacks","cried","wince","berle","flaws","boise","tibet","bided","shred","cocky","brice","delta","congo","holly","hicks","wraps","cocks","aisha","heard","cured","sades","horsy","umped","trice","dorky","curve","ferry","haler","ninth","pasta","jason","honer","kevin","males","fowls","awake","pores","meter","skate","drink","pussy","soups","bases","noyes","torts","bogus","still","soupy","dance","worry","eldon","stern","menes","dolls","dumpy","gaunt","grove","coops","mules","berry","sower","roams","brawl","greed","stags","blurs","swift","treed","taney","shame","easel","moves","leger","ville","order","spock","nifty","brian","elias","idler","serve","ashen","bizet","gilts","spook","eaten","pumas","cotes","broke","toxin","groan","laths","joins","spots","hated","tokay","elite","rawer","fiats","cards","sassy","milks","roost","glean","lutes","chins","drown","marks","pined","grace","fifth","lodes","rusts","terms","maxes","savvy","choir","savoy","spoon","halve","chord","hulas","sarah","celia","deems","ninny","wines","boggy","birch","raved","wales","beams","vibes","riots","warty","nigel","askew","faxes","sedge","sheol","pucks","cynic","relax","boers","whims","bents","candy","luann","slogs","bonny","barns","iambs","fused","duffy","guilt","bruin","pawls","penis","poppy","owing","tribe","tuner","moray","timid","ceded","geeks","kites","curio","puffy","perot","caddy","peeve","cause","dills","gavel","manse","joker","lynch","crank","golda","waits","wises","hasty","paves","grown","reedy","crypt","tonne","jerky","axing","swept","posse","rings","staff","tansy","pared","glaze","grebe","gonna","shark","jumps","vials","unset","hires","tying","lured","motes","linen","locks","mamas","nasty","mamie","clout","nader","velma","abate","tight","dales","serer","rives","bales","loamy","warps","plato","hooch","togae","damps","ofter","plumb","fifes","filmy","wiper","chess","lousy","sails","brahe","ounce","flits","hindu","manly","beaux","mimed","liken","forts","jambs","peeps","lelia","brews","handy","lusty","brads","marne","pesos","earle","arson","scout","showy","chile","sumps","hiked","crook","herbs","silks","alamo","mores","dunce","blaze","stank","haste","howls","trots","creon","lisle","pause","hates","mulch","mined","moder","devin","types","cindy","beech","tuned","mowed","pitts","chaos","colds","bidet","tines","sighs","slimy","brain","belle","leery","morse","ruben","prows","frown","disco","regal","oaken","sheds","hives","corny","baser","fated","throe","revel","bores","waved","shits","elvia","ferns","maids","color","coifs","cohan","draft","hmong","alton","stine","cluck","nodes","emily","brave","blair","blued","dress","bunts","holst","clogs","rally","knack","demos","brady","blues","flash","goofy","blocs","diane","colic","smile","yules","foamy","splay","bilge","faker","foils","condo","knell","crack","gallo","purls","auras","cakes","doves","joust","aides","lades","muggy","tanks","middy","tarps","slack","capet","frays","donny","venal","yeats","misty","denim","glass","nudes","seeps","gibbs","blows","bobbi","shane","yards","pimps","clued","quiet","witch","boxes","prawn","kerry","torah","kinko","dingy","emote","honor","jelly","grins","trope","vined","bagel","arden","rapid","paged","loved","agape","mural","budge","ticks","suers","wendi","slice","salve","robin","bleat","batik","myles","teddy","flatt","puppy","gelid","largo","attar","polls","glide","serum","fundy","sucks","shalt","sewer","wreak","dames","fonts","toxic","hines","wormy","grass","louse","bowls","crass","benny","moire","margo","golfs","smart","roxie","wight","reign","dairy","clops","paled","oddly","sappy","flair","shown","bulgy","benet","larch","curry","gulfs","fends","lunch","dukes","doris","spoke","coins","manna","conga","jinns","eases","dunno","tisha","swore","rhino","calms","irvin","clans","gully","liege","mains","besot","serge","being","welch","wombs","draco","lynda","forty","mumps","bloch","ogden","knits","fussy","alder","danes","loyal","valet","wooer","quire","liefs","shana","toyed","forks","gages","slims","cloys","yates","rails","sheep","nacho","divan","honks","stone","snack","added","basal","hasps","focus","alone","laxes","arose","lamed","wrapt","frail","clams","plait","hover","tacos","mooch","fault","teeth","marva","mucks","tread","waves","purim","boron","horde","smack","bongo","monte","swirl","deals","mikes","scold","muter","sties","lawns","fluke","jilts","meuse","fives","sulky","molds","snore","timmy","ditty","gasps","kills","carey","jawed","byers","tommy","homer","hexed","dumas","given","mewls","smelt","weird","speck","merck","keats","draws","trent","agave","wells","chews","blabs","roves","grieg","evens","alive","mulls","cared","garbo","fined","happy","trued","rodes","thurs","cadet","alvin","busch","moths","guild","staci","lever","widen","props","hussy","lamer","riley","bauer","chirp","rants","poxes","shyer","pelts","funny","slits","tinge","ramos","shift","caper","credo","renal","veils","covey","elmer","mated","tykes","wooed","briar","gears","foley","shoes","decry","hypes","dells","wilds","runts","wilts","white","easts","comer","sammy","lochs","favor","lance","dawns","bushy","muted","elsie","creel","pocks","tenet","cagey","rides","socks","ogled","soils","sofas","janna","exile","barks","frank","takes","zooms","hakes","sagan","scull","heaps","augur","pouch","blare","bulbs","wryly","homey","tubas","limbo","hardy","hoagy","minds","bared","gabby","bilks","float","limns","clasp","laura","range","brush","tummy","kilts","cooed","worms","leary","feats","robes","suite","veals","bosch","moans","dozen","rarer","slyer","cabin","craze","sweet","talon","treat","yanks","react","creed","eliza","sluts","cruet","hafts","noise","seder","flies","weeks","venus","backs","eider","uriel","vouch","robed","hacks","perth","shiny","stilt","torte","throb","merer","twits","reeds","shawl","clara","slurs","mixer","newts","fried","woolf","swoop","kaaba","oozed","mayer","caned","laius","lunge","chits","kenny","lifts","mafia","sowed","piled","stein","whack","colts","warms","cleft","girds","seeks","poets","angel","trade","parsi","tiers","rojas","vexes","bryce","moots","grunt","drain","lumpy","stabs","poohs","leapt","polly","cuffs","giddy","towns","dacha","quoth","provo","dilly","carly","mewed","tzars","crock","toked","speak","mayas","pssts","ocher","motel","vogue","camps","tharp","taunt","drone","taint","badge","scott","scats","bakes","antes","gruel","snort","capes","plate","folly","adobe","yours","papaw","hench","moods","clunk","chevy","tomas","narcs","vonda","wiles","prigs","chock","laser","viced","stiff","rouse","helps","knead","gazer","blade","tumid","avail","anger","egged","guide","goads","rabin","toddy","gulps","flank","brats","pedal","junky","marco","tinny","tires","flier","satin","darth","paley","gumbo","rared","muffs","rower","prude","frees","quays","homes","munch","beefs","leash","aston","colon","finch","bogey","leaps","tempo","posts","lined","gapes","locus","maori","nixes","liven","songs","opted","babel","wader","barer","farts","lisps","koran","lathe","trill","smirk","mamma","viler","scurf","ravel","brigs","cooky","sachs","fulls","goals","turfs","norse","hauls","cores","fairy","pluto","kneed","cheek","pangs","risen","czars","milne","cribs","genes","wefts","vents","sages","seres","owens","wiley","flume","haded","auger","tatty","onion","cater","wolfe","magic","bodes","gulls","gazes","dandy","snags","rowed","quell","spurn","shore","veldt","turns","slavs","coach","stalk","snuck","piles","orate","joyed","daily","crone","wager","solos","earns","stark","lauds","kasey","villa","gnaws","scent","wears","fains","laced","tamer","pipes","plant","lorie","rivet","tamed","cozen","theme","lifer","sunny","shags","flack","gassy","eased","jeeps","shire","fargo","timer","brash","behan","basin","volga","krone","swiss","docks","booed","ebert","gusty","delay","oared","grady","buick","curbs","crete","lucas","strum","besom","gorse","troth","donne","chink","faced","ahmed","texas","longs","aloud","bethe","cacao","hilda","eagle","karyn","harks","adder","verse","drays","cello","taped","snide","taxis","kinky","penes","wicca","sonja","aways","dyers","bolas","elfin","slope","lamps","hutch","lobed","baaed","masts","ashes","ionic","joyce","payed","brays","malts","dregs","leaky","runny","fecal","woven","hurls","jorge","henna","dolby","booty","brett","dykes","rural","fight","feels","flogs","brunt","preen","elvis","dopey","gripe","garry","gamma","fling","space","mange","storm","arron","hairs","rogue","repel","elgar","ruddy","cross","medan","loses","howdy","foams","piker","halts","jewel","avery","stool","cruel","cases","ruses","cathy","harem","flour","meted","faces","hobos","charm","jamar","cameo","crape","hooey","reefs","denny","mitts","sores","smoky","nopes","sooty","twirl","toads","vader","julep","licks","arias","wrote","north","bunks","heady","batch","snaps","claws","fouls","faded","beans","wimps","idled","pulse","goons","noose","vowed","ronda","rajas","roast","allah","punic","slows","hours","metal","slier","meaty","hanna","curvy","mussy","truth","troys","block","reels","print","miffs","busts","bytes","cream","otter","grads","siren","kilos","dross","batty","debts","sully","bares","baggy","hippy","berth","gorky","argon","wacko","harry","smoke","fails","perms","score","steps","unity","couch","kelly","rumps","fines","mouth","broth","knows","becky","quits","lauri","trust","grows","logos","apter","burrs","zincs","buyer","bayer","moose","overt","croon","ousts","lands","lithe","poach","jamel","waive","wiser","surly","works","paine","medal","glads","gybed","paint","lorre","meant","smugs","bryon","jinni","sever","viols","flubs","melts","heads","peals","aiken","named","teary","yalta","styes","heist","bongs","slops","pouts","grape","belie","cloak","rocks","scone","lydia","goofs","rents","drive","crony","orlon","narks","plays","blips","pence","march","alger","baste","acorn","billy","croce","boone","aaron","slobs","idyls","irwin","elves","stoat","doing","globe","verve","icons","trial","olsen","pecks","there","blame","tilde","milky","sells","tangy","wrack","fills","lofty","truce","quark","delia","stowe","marty","overs","putty","coral","swine","stats","swags","weans","spout","bulky","farsi","brest","gleam","beaks","coons","hater","peony","huffy","exert","clips","riven","payer","doped","salas","meyer","dryad","thuds","tilts","quilt","jetty","brood","gulch","corps","tunic","hubby","slang","wreck","purrs","punch","drags","chide","sulks","tints","huger","roped","dopes","booby","rosin","outer","gusto","tents","elude","brows","lease","ceres","laxer","worth","necks","races","corey","trait","stuns","soles","teems","scrip","privy","sight","minor","alisa","stray","spank","cress","nukes","rises","gusts","aurae","karma","icing","prose","biked","grand","grasp","skein","shaky","clump","rummy","stock","twain","zoned","offed","ghats","mover","randy","vault","craws","thees","salem","downy","sangs","chore","cited","grave","spinx","erica","raspy","dying","skips","clerk","paste","moved","rooks","intel","moses","avers","staid","yawls","blast","lyres","monks","gaits","floor","saner","waver","assam","infer","wands","bunch","dryly","weedy","honey","baths","leach","shorn","shows","dream","value","dooms","spiro","raped","shook","stead","moran","ditto","loots","tapir","looms","clove","stops","pinks","soppy","ripen","wench","shone","bauds","doric","leans","nadia","cries","camus","boozy","maris","fools","morns","bides","greek","gauss","roget","lamar","hazes","beefy","dupes","refed","felts","larry","guile","ables","wants","warns","toils","bathe","edger","paced","rinks","shoos","erich","whore","tiger","jumpy","lamas","stack","among","punts","scalp","alloy","solon","quite","comas","whole","parse","tries","reeve","tiled","deena","roomy","rodin","aster","twice","musts","globs","parch","drawn","filch","bonds","tells","droop","janis","holds","scant","lopes","based","keven","whiny","aspic","gains","franz","jerri","steel","rowel","vends","yelps","begin","logic","tress","sunni","going","barge","blood","burns","basks","waifs","bones","skill","hewer","burly","clime","eking","withs","capek","berta","cheap","films","scoot","tweed","sizer","wheat","acton","flung","ponds","tracy","fiver","berra","roger","mutes","burke","miked","valve","whisk","runes","parry","toots","japes","roars","rough","irons","romeo","cages","reeks","cigar","saiph","dully","hangs","chops","rolls","prick","acuff","spent","sulla","train","swell","frets","names","anita","crazy","sixth","blunt","fewer","large","brand","slick","spitz","rears","ogres","toffy","yolks","flock","gnawn","eries","blink","skier","feted","tones","snail","ether","barbs","noses","hears","upset","awash","cloud","trunk","degas","dungs","rated","shall","yeahs","coven","sands","susan","fable","gunny","began","serfs","balls","dinky","madge","prong","spilt","lilly","brawn","comet","spins","raids","dries","sorts","makes","mason","mayra","royce","stout","mealy","pagan","nasal","folds","libby","coups","photo","mosey","amens","speed","lords","board","fetal","lagos","scope","raked","bonus","mutts","willy","sport","bingo","thant","araby","bette","rebel","gases","small","humus","grosz","beset","slays","steve","scrap","blahs","south","pride","heels","tubes","beady","lacey","genus","mauls","vying","spice","sexes","ester","drams","today","comae","under","jests","direr","yoked","tempi","early","boats","jesus","warts","guppy","gilda","quota","token","edwin","ringo","gaped","lemon","hurst","manor","arrow","mists","prize","silas","blobs","diets","ervin","stony","buddy","bates","rabid","ducat","ewing","jaunt","beads","doyen","blush","thoth","tiles","piper","short","peron","alley","decks","shunt","whirs","cushy","roils","betty","plugs","woken","jibes","foray","merak","ruing","becks","whale","shoot","dwelt","spawn","fairs","dozed","celts","blond","tikes","sabin","feint","vamps","cokes","willa","slues","bills","force","curst","yokel","surer","miler","fices","arced","douse","hilly","lucio","tongs","togas","minty","sagas","pates","welsh","bruno","decal","elate","linux","gyros","pryor","mousy","pains","shake","spica","pupal","probe","mount","shirk","purus","kilns","rests","graze","hague","spuds","sweep","momma","burch","maces","samar","brace","riser","booms","build","camel","flyer","synge","sauna","tonga","tings","promo","hides","clair","elisa","bower","reins","diann","lubed","nulls","picks","laban","milch","buber","stomp","bosom","lying","haled","avert","wries","macon","skids","fumed","ogles","clods","antic","nosey","crimp","purge","mommy","cased","taxes","covet","clack","butch","panty","lents","machs","exude","tooth","adore","shuck","asses","after","terra","dices","aryan","regor","romes","stile","cairo","maura","flail","eaves","estes","sousa","visas","baron","civet","kitty","freed","ralph","tango","gawks","cheat","study","fancy","fiber","musks","souse","brims","claim","bikes","venue","sired","thymi","rivas","skimp","pleas","woman","gimpy","cawed","minos","pints","knock","poked","bowen","risky","towel","oinks","linus","heals","pears","codas","inner","pitch","harpy","niger","madly","bumpy","stair","files","nobel","celli","spars","jades","balmy","kooky","plums","trues","gloss","trims","daunt","tubby","dared","wadis","smell","darby","stink","drill","dover","ruler","laden","dikes","layla","fells","maker","joked","horns","these","baize","spahn","whens","edged","mushy","plume","tucks","spurs","husky","dried","bigot","pupas","drily","aware","hagar","newly","knots","pratt","feces","sabik","watts","cooke","riles","seamy","fleas","dusts","barfs","roans","pawns","vivid","kirks","tania","feral","tubae","horne","aries","brits","combs","chunk","stork","waned","texan","elide","glens","emery","autos","trams","dosed","cheri","baits","jacks","whose","fazed","matte","swans","maxed","write","spays","orion","traci","horse","stars","strut","goods","verge","scuff","award","dives","wires","burnt","dimly","sleds","mayan","biped","quirk","sofia","slabs","waste","robby","mayor","fatty","items","bowel","mires","swarm","route","swash","sooth","paved","steak","upend","sough","throw","perts","stave","carry","burgs","hilts","plane","toady","nadir","stick","foist","gnarl","spain","enter","sises","story","scarf","ryder","glums","nappy","sixes","honed","marcy","offer","kneel","leeds","lites","voter","vince","bursa","heave","roses","trees","argos","leann","grimy","zelma","crick","tract","flips","folks","smote","brier","moore","goose","baden","riled","looks","sober","tusks","house","acmes","lubes","chows","neath","vivas","defer","allay","casey","kmart","pests","proms","eying","cider","leave","shush","shots","karla","scorn","gifts","sneer","mercy","copes","faxed","spurt","monet","awoke","rocky","share","gores","drawl","tears","mooed","nones","wined","wrens","modem","beria","hovel","retch","mates","hands","stymy","peace","carat","coots","hotel","karen","hayed","mamet","cuing","paper","rages","suave","reuse","auden","costs","loner","rapes","hazel","rites","brent","pumps","dutch","puffs","noons","grams","teats","cease","honda","pricy","forgo","fleck","hired","silos","merge","rafts","halon","larks","deere","jello","cunts","sifts","boner","morin","mimes","bungs","marie","harts","snobs","sonic","hippo","comes","crops","mango","wrung","garbs","natty","cents","fitch","moldy","adams","sorta","coeds","gilds","kiddy","nervy","slurp","ramon","fuzed","hiker","winks","vanes","goody","hawks","crowd","bract","marla","limbs","solve","gloom","sloop","eaton","memos","tames","heirs","berms","wanes","faint","numbs","holes","grubs","rakes","waist","miser","stays","antis","marsh","skyed","payne","champ","jimmy","clues","fatal","shoed","freon","lopez","snowy","loins","stale","thank","reads","isles","grill","align","saxes","rubin","rigel","walls","beers","wispy","topic","alden","anton","ducts","david","duets","fries","oiled","waken","allot","swats","woozy","tuxes","inter","dunne","known","axles","graph","bumps","jerry","hitch","crews","lucia","banal","grope","valid","meres","thick","lofts","chaff","taker","glues","snubs","trawl","keels","liker","stand","harps","casks","nelly","debby","panes","dumps","norma","racks","scams","forte","dwell","dudes","hypos","sissy","swamp","faust","slake","maven","lowed","lilts","bobby","gorey","swear","nests","marci","palsy","siege","oozes","rates","stunt","herod","wilma","other","girts","conic","goner","peppy","class","sized","games","snell","newsy","amend","solis","duane","troop","linda","tails","woofs","scuds","shies","patti","stunk","acres","tevet","allen","carpi","meets","trend","salty","galls","crept","toner","panda","cohen","chase","james","bravo","styed","coals","oates","swami","staph","frisk","cares","cords","stems","razed","since","mopes","rices","junes","raged","liter","manes","rearm","naive","tyree","medic","laded","pearl","inset","graft","chair","votes","saver","cains","knobs","gamay","hunch","crags","olson","teams","surge","wests","boney","limos","ploys","algae","gaols","caked","molts","glops","tarot","wheal","cysts","husks","vaunt","beaus","fauns","jeers","mitty","stuff","shape","sears","buffy","maced","fazes","vegas","stamp","borer","gaged","shade","finds","frock","plods","skied","stump","ripes","chick","cones","fixed","coled","rodeo","basil","dazes","sting","surfs","mindy","creak","swung","cadge","franc","seven","sices","weest","unite","codex","trick","fusty","plaid","hills","truck","spiel","sleek","anons","pupae","chiba","hoops","trash","noted","boris","dough","shirt","cowls","seine","spool","miens","yummy","grade","proxy","hopes","girth","deter","dowry","aorta","paean","corms","giant","shank","where","means","years","vegan","derek","tales"}
	//	//beginWord := "qa"
	//	//endWord := "sq"
	//	//wordList := []string_issue{"si","go","se","cm","so","ph","mt","db","mb","sb","kr","ln","tm","le","av","sm","ar","ci","ca","br","ti","ba","to","ra","fa","yo","ow","sn","ya","cr","po","fe","ho","ma","re","or","rn","au","ur","rh","sr","tc","lt","lo","as","fr","nb","yb","if","pb","ge","th","pm","rb","sh","co","ga","li","ha","hz","no","bi","di","hi","qa","pi","os","uh","wm","an","me","mo","na","la","st","er","sc","ne","mn","mi","am","ex","pt","io","be","fm","ta","tb","ni","mr","pa","he","lr","sq","ye"}
	//	//beginWord := "hot"
	//	//endWord := "dog"
	//	//wordList := []string_issue{"hot","cog","dog","tot","hog","hop","pot","dot"}
	//	//beginWord := "hot"
	//	//endWord := "dog"
	//	//wordList := []string_issue{"hot","dog","dot"}
	//	res := ladderLength(beginWord,endWord,wordList)
	//	fmt.Println(res)
	//}
	//{
	//	input := [][]byte{{'O','X','X','O','X'},{'X','O','O','X','O'},{'X','O','X','O','X'},{'O','X','O','O','O'},{'X','X','O','X','O'}}
	//	solve(input)
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
	//	preorder := []int{3,9,20,15,7}
	//	inorder := []int{9,3,15,20,7}
	//	res := buildTree2(preorder,inorder)
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
	//	input := "aa"
	//	res := partition(input)
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
	//	res := pathInZigZagTree(14)
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
	//	var input [][]int = [][]int{{1,3},{2,3},{3,1}}
	//	res := findJudge(3,input)
	//	fmt.Println(res)
	//}
	//{
	//	var input []string_issue = []string_issue{"Hello","Alaska", "Dad", "Peace"}
	//	ret := findWords(input)
	//	fmt.Println(ret)
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
