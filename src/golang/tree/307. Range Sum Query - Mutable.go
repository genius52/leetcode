package tree

import "math"

type NumArray struct {
	sum int
	begin,end int
	left,right *NumArray
}

func build_segment_tree(nums []int,begin int,end int) (ret *NumArray){
	var obj NumArray
	if(begin == end){
		obj.sum = nums[begin]
		obj.begin = begin
		obj.end = end
		obj.left = nil
		obj.right = nil
		return &obj
	}
	var mid int = begin + (end - begin)/2
	left := build_segment_tree(nums,begin,mid)
	right := build_segment_tree(nums,mid + 1,end)
	obj.sum = left.sum + right.sum
	obj.left = left
	obj.right = right
	obj.begin = begin
	obj.end = end
	return &obj
}

func Constructor307(nums []int) NumArray {
	var l int = len(nums)
	if(l == 0){
		var obj NumArray
		return obj
	}
	obj := build_segment_tree(nums,0,l - 1)
	return *obj
}

func (this *NumArray) Update2(i int, val int) int{
	if(this.begin == i && this.end == i){
		diff := val - this.sum
		this.sum = val
		return diff
	}
	var mid = this.begin + (this.end - this.begin) /2
	if(i >= this.begin && i <= mid){
		diff := this.left.Update2(i,val)
		this.sum += diff
		return diff
	}else{
		diff := this.right.Update2(i,val)
		this.sum += diff
		return diff
	}
}

func (this *NumArray) Update(i int, val int)  {
	this.Update2(i,val)
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == this.begin && j == this.end{
		return this.sum
	}else{
		var mid = this.begin + (this.end - this.begin) /2
		if(i > mid){
			return this.right.SumRange(i,int(math.Min(float64(this.end),float64(j))))
		}
		if(j <= mid){
			return this.left.SumRange(int(math.Max(float64(this.begin),float64(i))),j)
		}
		if(i <= mid && j > mid){
			l_sum := this.left.SumRange(int(math.Max(float64(this.begin),float64(i))),mid)
			r_sum := this.right.SumRange(mid + 1,int(math.Min(float64(this.end),float64(j))))
			return l_sum + r_sum
		}
		return 0
	}
}