package number

func min_int64(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func MinCost2735(nums []int, x int) int64 {
	var l int = len(nums)
	var pre [][]int64 = make([][]int64, l)
	for i := 0; i < l; i++ {
		pre[i] = make([]int64, l)
		pre[i][0] = int64(nums[i])
	}
	for i := 0; i < l; i++ {
		for j := 1; j < l; j++ {
			pre[i][j] = min_int64(pre[i][j-1], int64(nums[(i+l-j)%l]))
		}
	}
	var min_cost []int64 = make([]int64, l) //min_cost[i]: 右移i次的代价
	for i := 0; i < l; i++ {
		min_cost[i] += int64(nums[i])
	}
	var res int64 = 1<<62 - 1
	for i := 0; i < l; i++ {
		var cur int64 = 0
		for j := 0; j < l; j++ {
			cur += pre[j][i]
		}
		res = min_int64(res, cur+int64(x*i))
	}
	return res
}

//Wrong solution
//func MinCost2735(nums []int, x int) int64 {
//	var l int = len(nums)
//	var min_cost []int64 = make([]int64, l)
//	var max_offset int = 0
//	for i := 0; i < l; i++ {
//		min_cost[i] = int64(nums[i])
//		//var min_cost_offset int = -1
//		offset := 1
//		for offset < l {
//			idx := (i + l - offset) % l
//			var cur int64 = int64(nums[idx] + x*offset)
//			if cur < min_cost[i] {
//				min_cost[i] = int64(nums[idx])
//				if offset > max_offset {
//					max_offset = offset
//				}
//			}
//			offset++
//		}
//	}
//	var res int64 = 0
//	for i := 0; i < l; i++ {
//		res += min_cost[i]
//	}
//	res += int64(max_offset * x)
//	return res
//}
