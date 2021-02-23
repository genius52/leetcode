package array


//128
//Input: [100, 4, 200, 1, 3, 2]    [9,1,4,7,3,-1,0,5,8,-1,6]
//Output: 4
//Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
func dfs_longestConsecutive(record map[int]bool,num int,increase bool)int{
	if _,ok := record[num];ok{
		if !record[num]{
			return 0
		}
		if increase{
			return 1 + dfs_longestConsecutive(record,num + 1,increase)
		}else{
			return 1 + dfs_longestConsecutive(record,num - 1,increase)
		}
	}else{
		return 0
	}
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var record map[int]bool = make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		record[nums[i]] = true
	}
	var max int = 1
	for num, _ := range record {
		if record[num] {
			res := 1 + dfs_longestConsecutive(record, num+1, true) + dfs_longestConsecutive(record, num-1, false)
			record[num] = false
			if res > max {
				max = res
			}
		}
	}
	return max
}