package list_queue

import "sort"

type KthLargest struct {
	k int
	nums []int
}

func ConstructorK(k int, nums []int) KthLargest {
	var obj KthLargest = KthLargest{
		k:k,
		nums:nums,
	}
	sort.Ints(obj.nums)
	if len(obj.nums) > k{
		obj.nums = obj.nums[len(obj.nums) - k:]
	}
	return obj
}

func (this *KthLargest) Add(val int) int {
	if len(this.nums) < this.k{
		this.nums = append(this.nums,val)
		sort.Ints(this.nums)
		return this.nums[0]
	}
	if val < this.nums[0]{
		return this.nums[0]
	}
	this.nums = append(this.nums,val)
	sort.Ints(this.nums)
	this.nums = this.nums[1:]
	return this.nums[0]
}