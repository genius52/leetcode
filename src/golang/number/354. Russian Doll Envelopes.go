package number

import "sort"

//Input: envelopes = [[5,4],[6,4],[6,7],[2,3]]
//Output: 3
//Explanation: The maximum number of envelopes you can Russian doll is 3 ([2,3] => [5,4] => [6,7]).

// LIS solution
func MaxEnvelopes(envelopes [][]int) int {
	var l int = len(envelopes)
	if l <= 1{
		return l
	}
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0]{
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	var dp []int //dp[i]: 长度为i + 1的width
	for i := 0;i < l;i++{
		find_idx := sort.SearchInts(dp,envelopes[i][1])
		if find_idx == len(dp){
			dp = append(dp,envelopes[i][1])
		}else{
			dp[find_idx] = envelopes[i][1]
		}
	}
	return len(dp)
}


//type sortWidth [][]int
//
//func (s sortWidth) Less(i, j int) bool {
//	if s[i][0] < s[j][0]{
//		return true
//	}else if s[i][0] > s[j][0]{
//		return false
//	}else{
//		return s[i][1] > s[j][1]
//	}
//}
//
//func (s sortWidth) Swap(i, j int) {
//	s[i], s[j] = s[j], s[i]
//}
//
//func (s sortWidth) Len() int {
//	return len(s)
//}
//
//func MaxEnvelopes(envelopes [][]int) int {
//	var l int = len(envelopes)
//	if l <= 1{
//		return l
//	}
//	sort.Sort(sortWidth(envelopes))
//	var dp []int = make([]int,l)
//	dp[0] = 1
//	var max_len int = 1
//	for i := 1;i < l;i++{
//		var cur_max_len int = 1
//		for j := i - 1;j >= 0;j--{
//			if envelopes[i][0] > envelopes[j][0]&& envelopes[i][1] > envelopes[j][1]{
//				var cur_len int = 1 + dp[j]
//				if cur_len > cur_max_len{
//					cur_max_len = cur_len
//				}
//			}
//		}
//		if max_len < cur_max_len{
//			max_len = cur_max_len
//		}
//		dp[i] = cur_max_len
//	}
//	return max_len
//}