package array

func dfs_canChoose(groups [][]int,index int,nums []int,pos int)bool{
	var group_len int = len(groups)
	var nums_len int = len(nums)
	if index >= group_len{
		return true
	}
	var sub_len = len(groups[index])
	for i := 0;i < sub_len;i++{
		if pos + i >= nums_len{
			return false
		}
		if groups[index][i] != nums[pos + i]{
			return false
		}
	}
	return dfs_canChoose(groups,index + 1,nums,pos + sub_len)
}

func CanChoose(groups [][]int, nums []int) bool {
	var l int = len(nums)
	var res bool = false
	for i := 0;i < l;i++{
		res = dfs_canChoose(groups,0,nums,i)
		if res{
			break
		}
	}
	return res
}