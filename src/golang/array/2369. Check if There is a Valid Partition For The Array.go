package array

func dfs_validPartition(nums []int,l int,idx int,visited []bool)bool{
	if idx >= l{
		return true
	}
	if idx == l - 1{
		return false
	}
	if idx == l - 2{
		return nums[idx] == nums[idx + 1]
	}
	if visited[idx]{
		return false
	}
	visited[idx] = true
	if nums[idx] == nums[idx + 1]{
		if nums[idx] == nums[idx + 2]{
			if dfs_validPartition(nums,l,idx + 3,visited){
				return true
			}
		}
		return dfs_validPartition(nums,l,idx + 2,visited)
	}else if (nums[idx] + 1) == nums[idx + 1] && (nums[idx + 1] + 1) == nums[idx + 2]{
		return dfs_validPartition(nums,l,idx + 3,visited)
	}
	return false
}

func validPartition(nums []int) bool {
	var l int = len(nums)
	var visited []bool = make([]bool,l)
	return dfs_validPartition(nums,l,0,visited)
}