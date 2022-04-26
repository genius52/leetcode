package number

import "math/rand"

type Solution struct {
	nums_ []int
}

func Constructor384(nums []int) Solution {
	var obj Solution
	obj.nums_ = nums
	return obj
}

func (this *Solution) Reset() []int {
	return this.nums_
}

func (this *Solution) Shuffle() []int {
	var l int = len(this.nums_)
	var res []int = make([]int,l)
	copy(res,this.nums_)
	for i := 0;i < l;i++{
		swap_idx := rand.Intn(l - i) % (l - i) + i //从[i,l - 1]中选取一个
		res[i],res[swap_idx] = res[swap_idx],res[i]
	}
	return res
}