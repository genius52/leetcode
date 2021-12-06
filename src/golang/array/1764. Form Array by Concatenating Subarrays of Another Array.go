package array

//Input: groups = [[1,-1,-1],[3,-2,0]], nums = [1,-1,0,1,-1,-1,3,-2,0]
//Output: true
//Explanation: You can choose the 0th subarray as [1,-1,0,1,-1,-1,3,-2,0] and the 1st one as [1,-1,0,1,-1,-1,3,-2,0].
func dfs_canChoose(groups [][]int, l1 int, idx1 int, nums []int, l2 int, idx2 int) bool {
	if idx1 >= l1 {
		return true
	}
	if idx2 >= l2 {
		return false
	}
	var i int = idx2
	var j int = 0
	var group_len int = len(groups[idx1])
	for i < l2 && j < group_len {
		if groups[idx1][j] == nums[i] {
			i++
			j++
		} else {
			break
		}
	}
	var res bool = false
	if j == group_len {
		res = dfs_canChoose(groups, l1, idx1+1, nums, l2, i)
	}
	if res {
		return true
	}
	return dfs_canChoose(groups, l1, idx1, nums, l2, idx2+1)
}

func CanChoose(groups [][]int, nums []int) bool {
	return dfs_canChoose(groups, len(groups), 0, nums, len(nums), 0)
}
