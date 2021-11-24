package array

func MinMoves(nums []int, limit int) int {
	var l int = len(nums)
	var record []int = make([]int,limit * 2 + 2)//record[i]: nums[i] + nums[l - 1 - i] = i时,move的次数
	for i := 0;i < l/2;i++{
		// [2, 2 * limit] 范围 + 2
		record[2] += 2
		record[limit * 2 + 1] -= 2

		record[min_int(nums[i],nums[l - 1 - i]) + 1] -= 1
		record[max_int(nums[i],nums[l - 1 - i]) + limit + 1] += 1

		record[nums[i] + nums[l - 1 - i]] -= 1
		record[nums[i] + nums[l - 1 - i] + 1] += 1
	}
	var res int = l * 2
	var cnt int = 0
	for i := 2;i <= limit * 2;i++{
		cnt += record[i]
		if cnt < res{
			res = cnt
		}
	}
	return res
}
//TLE
//func MinMoves(nums []int, limit int) int {
//	var l int = len(nums)
//	var record []int = make([]int,limit * 2 + 1)//record[i]: nums[i] + nums[l - 1 - i] = i时,move的次数
//	for i := 2;i <= limit * 2;i++{
//		record[i] = l
//	}
//	for i := 0;i < l/2;i++{
//		for j := min_int(nums[i],nums[l - 1 - i]) + 1;j <= max_int(nums[i],nums[l - 1 - i]) + limit;j++{
//			record[j] -= 1
//		}
//		record[nums[i] + nums[l - 1 - i]] -= 1
//	}
//	var res int = l * 2
//	for i := 2;i <= limit * 2;i++{
//		if record[i] < res{
//			res = record[i]
//		}
//	}
//	return res
//}