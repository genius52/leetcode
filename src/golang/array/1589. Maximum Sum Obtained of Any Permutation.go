package array

import "sort"

//Input: nums = [1,2,3,4,5,10], requests = [[0,2],[1,3],[1,1]]
//Output: 47
//Explanation: A permutation with the max total sum is [4,10,5,3,2,1] with request sums [19,18,10].
func MaxSumRangeQuery(nums []int, requests [][]int) int{
	var l int = len(nums)
	var record []int = make([]int,l)
	var len_request int = len(requests)
	for i := 0;i < len_request;i++{
		record[requests[i][0]]++
		if requests[i][1] < l - 1{
			record[requests[i][1] + 1]--
		}
	}
	for i := 1;i < l;i++{
		record[i] += record[i - 1]
	}
	sort.Ints(record)
	sort.Ints(nums)
	var res int = 0
	for i := l - 1;i >= 0;i--{
		if record[i] == 0{
			break
		}
		res += record[i] * nums[i]
		res %= 1000000007
	}
	return res
}

//func MaxSumRangeQuery(nums []int, requests [][]int) int {
//	sort.Ints(nums)
//	var l int = len(nums)
//	var record []int = make([]int,l)
//	var len_request int = len(requests)
//	for i := 0;i < len_request;i++{
//		for begin := requests[i][0];begin <= requests[i][1];begin++{
//			record[begin]++
//		}
//	}
//	var q []int
//	for i := 0;i < l;i++{
//		if record[i] == 0{
//			continue
//		}
//		q = append(q,record[i])
//	}
//	sort.Ints(q)
//	var res int = 0
//	var pos int = l - 1
//	var q_len int = len(q)
//	for i := q_len - 1;i >= 0;i--{
//		res += q[i] * nums[pos]
//		res %= 1000000007
//		pos--
//	}
//	return res
//}