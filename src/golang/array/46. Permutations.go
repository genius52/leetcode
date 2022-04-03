package array

func dfs_permute(nums []int,l int,idx int,res *[][]int){
	if idx >= l{
		var data []int = make([]int,l)
		copy(data,nums)
		*res = append(*res,data)
	}
	for i := idx;i < l;i++{
		nums[i],nums[idx] = nums[idx],nums[i]
		dfs_permute(nums,l,idx + 1,res)
		nums[i],nums[idx] = nums[idx],nums[i]
	}
}

func Permute(nums []int) [][]int {
	var res [][]int
	dfs_permute(nums,len(nums),0,&res)
	return res
}