package array

func dp_subsets(nums []int,l int,idx int,pre []int,res *[][]int){
	if idx >= l{
		return
	}
	//add current
	var next []int
	if len(pre) > 0{
		next = make([]int,len(pre))
		copy(next,pre)
	}
	pre = append(pre,nums[idx])
	*res = append(*res,pre)
	dp_subsets(nums,l,idx + 1,pre,res)

	//skip current
	dp_subsets(nums,l,idx + 1,next,res)
}

func Subsets(nums []int) [][]int {
	var l int = len(nums)
	var res [][]int = [][]int{[]int{}}
	var pre []int
	dp_subsets(nums,l,0,pre,&res)
	return res
}