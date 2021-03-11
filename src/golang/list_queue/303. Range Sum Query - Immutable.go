package list_queue

//obj := Constructor(nums);
//param_1 := obj.SumRange(i,j);
type NumArray struct {
	sum []int
}

func Constructor3(nums []int) NumArray {
	var obj NumArray
	var l int = len(nums)
	obj.sum = make([]int,l + 1)
	obj.sum[0] = 0
	for i := 1;i <= l;i++{
		obj.sum[i] = nums[i - 1] + obj.sum[i - 1]
	}
	return obj
}

func (this *NumArray) SumRange(i int, j int) int {
	if i < 0 || j >= len(this.sum){
		return 0
	}
	return this.sum[j + 1] - this.sum[i]
}