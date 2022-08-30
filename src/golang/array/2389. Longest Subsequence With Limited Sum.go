package array

import "sort"

func answerQueries(nums []int, queries []int) []int {
	var l1 int = len(nums)
	var l2 int = len(queries)
	var res []int = make([]int,l2)
	sort.Ints(nums)
	var prefix_sum []int = make([]int,l1 + 1)
	for i := 0;i < l1;i++{
		prefix_sum[i + 1] = prefix_sum[i] + nums[i]
	}
	for i := 0;i < l2;i++{
		//for j := 0;j <= l1;j++{
		//	if prefix_sum[j] > queries[i]{
		//		break
		//	}
		//	res[i] = j
		//}
		find_idx := sort.Search(l1 + 1, func(j int) bool {
			return prefix_sum[j] > queries[i]
		})
		res[i] = find_idx - 1
	}
	return res
}