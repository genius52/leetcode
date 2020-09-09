package number

import (
	"math"
	"sort"
)

//Input: [[1,2], [2,3], [3,4]]
//Output: 2
//Explanation: The longest chain is [1,2] -> [3,4]

type sortPair [][]int

func (s sortPair) Less(i, j int) bool {
	if s[i][1] < s[j][1]{
		return true
	}else if s[i][1] > s[j][1]{
		return false
	}else{
		return s[i][0] < s[j][0]
	}
}

func (s sortPair) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortPair) Len() int {
	return len(s)
}

func FindLongestChain(pairs [][]int) int {
	var l int = len(pairs)
	sort.Sort(sortPair(pairs))
	var res int = 0
	var last_num int = math.MinInt32
	for i := 0;i < l;i++{
		if pairs[i][0] > last_num{
			res++
			last_num = pairs[i][1]
		}
	}
	return res
}

func FindLongestChain2(pairs [][]int) int {
	sort.Sort(sortPair(pairs))
	var l int = len(pairs)
	var dp []int = make([]int,l)//the longest length on index i
	for i := 0;i < l;i++{
		dp[i] = 1
	}
	for i := 0;i < l;i++{
		for j := 0;j < i;j++{
			if pairs[i][0] > pairs[j][1]{
				dp[i] = int(math.Max(float64(1 + dp[j]),float64(dp[i])))
			}else{
				dp[i] = int(math.Max(float64(dp[j]),float64(dp[i])))
			}
		}
	}
	return dp[l-1]
}