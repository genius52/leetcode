package array

func maximumSegmentSum(nums []int, removeQueries []int) []int64 {
	var l int = len(nums)
	var groups []int = make([]int,l + 1)
	for i := 0;i <= l;i++{
		groups[i] = i
	}
	var get_parent func (groups []int,node int)int
	get_parent = func (groups []int,node int)int{
		if groups[node] != node{
			groups[node] = get_parent(groups,groups[node])
		}
		return groups[node]
	}
	var sum []int64 = make([]int64,l + 1)//sum[i]: sum of  i'th group
	var res []int64 = make([]int64,l)
	for i := l - 1;i > 0;i--{
		idx := removeQueries[i]
		group := get_parent(groups,idx + 1)
		groups[idx] = group
		sum[group] += sum[idx] + int64(nums[idx])
		res[i - 1]  = max_int64(res[i],sum[group])
	}
	return res
}