package array

func dfs_canJump(nums []int,start int,visited []bool)bool{
	if start >= len(nums) - 1 {
		return true
	}
	max_steps := nums[start]
	if max_steps == 0{
		return false
	}
	if visited[start]{
		return false
	}
	for i := 1;i <= max_steps;i++{
		if visited[start+i] {
			continue
		}
		res := dfs_canJump(nums,start + i,visited)
		if res{
			return true
		}
		visited[start + i] = true
	}
	return false
}

func canJump(nums []int) bool {
	var visited []bool = make([]bool,len(nums))
	return dfs_canJump(nums,0,visited)
}

//55
//Input: [2,3,1,1,4]
//Output: true
//Input: nums = [3,2,1,0,4]
//Output: false
func CanJump(nums []int) bool{
	var l int = len(nums)
	if l <= 1{
		return true
	}
	var longest_pos int = nums[0]
	var visit int = 1
	for longest_pos > 0{
		if longest_pos >= l - 1{
			return true
		}
		var next_visit = longest_pos + 1
		var cur_longest int = longest_pos
		for visit <= longest_pos{
			cur_longest = max_int(cur_longest,visit + nums[visit])
			if cur_longest >= l - 1{
				return true
			}
			visit++
		}
		if cur_longest <= longest_pos{
			return false
		}
		longest_pos = cur_longest
		visit = next_visit
	}
	return false
}