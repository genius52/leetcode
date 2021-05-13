package array

func dfs_canPartitionKSubsets(nums []int,target int,total_set int,cur_set int,cur_sum int,pos int,visited []bool)bool{
	if cur_sum > target{
		return false
	}
	if total_set == cur_set + 1{
		return true
	}
	if cur_sum == target{
		return dfs_canPartitionKSubsets(nums,target,total_set,cur_set + 1,0,0,visited)
	}
	if pos >= len(nums){
		return false
	}
	for i := pos;i < len(nums);i++{
		if !visited[i] && nums[i] + cur_sum <= target{
			visited[i] = true
			res := dfs_canPartitionKSubsets(nums,target,total_set,cur_set,cur_sum + nums[i],i + 1,visited)
			if res{
				return true
			}
			visited[i] = false
		}
		if cur_sum == 0 && cur_set != total_set{
			break
		}
	}
	return false
}

func CanPartitionKSubsets(nums []int, k int) bool {
	var total int = 0
	for _,n := range nums{
		total += n
	}
	var l int = len(nums)
	target := total / k
	if total % k != 0 || nums[l - 1] > target{
		return false
	}
	var visited []bool = make([]bool,l)
	return dfs_canPartitionKSubsets(nums,target,k,0,0,0,visited)
}