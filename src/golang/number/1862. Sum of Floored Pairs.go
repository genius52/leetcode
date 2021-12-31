package number

func SumOfFlooredPairs(nums []int) int {
	var max_num int = 0
	var l int = len(nums)
	for i := 0; i < l; i++ {
		if nums[i] > max_num {
			max_num = nums[i]
		}
	}
	var record []int = make([]int, max_num+1)
	for i := 0; i < l; i++ {
		record[nums[i]]++
	}
	var prefix []int = make([]int, max_num+1)
	for i := 1; i <= max_num; i++ {
		prefix[i] = prefix[i-1] + record[i]
	}
	var res int = 0
	var MOD int = 1e9 + 7
	for i := 1; i <= max_num; i++ {
		cnt := prefix[i] - prefix[i-1]
		if cnt == 0 {
			continue
		}
		pre_cnt := prefix[i-1]
		for j := i*2 - 1; ; j += i {
			if j > max_num {
				j = max_num
			}
			cnt2 := prefix[j] - pre_cnt
			pre_cnt = prefix[j]
			res += j / i * cnt * cnt2
			res %= MOD
			if j == max_num {
				break
			}
		}
	}
	return res
}

//func SumOfFlooredPairs(nums []int) int {
//	var record map[int]int = make(map[int]int)
//	for _, n := range nums {
//		record[n]++
//	}
//	var res int = 0
//	var MOD int = 1e9 + 7
//	for k1, v1 := range record {
//		for k2, v2 := range record {
//			if k1 == k2 {
//				res += v1 * v2
//				res %= MOD
//			} else {
//				res += k1 / k2 * v1 * v2
//				res %= MOD
//			}
//		}
//	}
//	return res
//}
