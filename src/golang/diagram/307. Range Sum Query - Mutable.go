package diagram

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */

type NumArray struct {
	data []int
}

func Constructor307(nums []int) NumArray {
	var obj NumArray
	obj.data = nums
	return obj
}

func (this *NumArray) Update(i int, val int)  {
	this.data[i] = val
}

func (this *NumArray) SumRange(i int, j int) int {
	return 0
}