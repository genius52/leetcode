package array

//Input: [100, 4, 200, 1, 3, 2]    [9,1,4,7,3,-1,0,5,8,-1,6]
//Output: 4
//Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
func check_longestConsecutive(n int,record map[int]bool,step int)int{
	var res int = 0
	for true{
		if _,ok := record[n];ok{
			delete(record,n)
			n += step
			res++
		}else{
			break
		}
	}
	return res
}

func LongestConsecutive(nums []int) int{
	var l int = len(nums)
	if l <= 1 {
		return l
	}
	var record map[int]bool = make(map[int]bool)
	for i := 0; i < l; i++ {
		record[nums[i]] = true
	}
	var res int = 1
	for i := 0;i < l;i++{
		if _,ok := record[nums[i]];ok{
			cur := 1 + check_longestConsecutive(nums[i] - 1,record,-1) + check_longestConsecutive(nums[i] + 1,record,1)
			if cur > res{
				res = cur
			}
			delete(record,nums[i])
		}
	}
	return res
}


//func dfs_longestConsecutive(record map[int]bool,num int,increase bool)int{
//	if _,ok := record[num];ok{
//		if !record[num]{
//			return 0
//		}
//		if increase{
//			return 1 + dfs_longestConsecutive(record,num + 1,increase)
//		}else{
//			return 1 + dfs_longestConsecutive(record,num - 1,increase)
//		}
//	}else{
//		return 0
//	}
//}
//
//func longestConsecutive(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	var record map[int]bool = make(map[int]bool)
//	for i := 0; i < len(nums); i++ {
//		record[nums[i]] = true
//	}
//	var max int = 1
//	for num, _ := range record {
//		if record[num] {
//			res := 1 + dfs_longestConsecutive(record, num+1, true) + dfs_longestConsecutive(record, num-1, false)
//			record[num] = false
//			if res > max {
//				max = res
//			}
//		}
//	}
//	return max
//}

//func LongestConsecutive2(nums []int) int {
//	sort.Ints(nums)
//	var l int = len(nums)
//	var res int = 0
//	var left int = 0
//	for left < l{
//		var right int = left + 1
//		var dup_cnt int = 0
//		for right < l {
//			if nums[right - 1] == nums[right]{
//				right++
//				dup_cnt++
//			}else if (nums[right - 1] + 1) == nums[right]{
//				right++
//			}else{
//				break
//			}
//		}
//		cur_len := right - left - dup_cnt
//		if cur_len > res{
//			res = cur_len
//		}
//		left = right
//	}
//	return res
//}