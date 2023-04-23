package number

func maxDivScore(nums []int, divisors []int) int {
	//sort.Ints(nums)
	//sort.Ints(divisors)
	var l1 int = len(nums)
	var l2 int = len(divisors)
	var max_cnt int = 0
	var res int = divisors[0]
	for i := 0; i < l2; i++ {
		if divisors[i] == 1 {
			return 1
		}
		if i > 0 && divisors[i] == divisors[i-1] {
			continue
		}
		var cur_cnt int = 0
		for j := 0; j < l1; j++ {
			//if divisors[i] > nums[j] {
			//	break
			//}
			if (nums[j] % divisors[i]) == 0 {
				cur_cnt++
			}
		}
		if cur_cnt > max_cnt {
			res = divisors[i]
			max_cnt = cur_cnt
		}
		if cur_cnt == max_cnt {
			if divisors[i] < res {
				res = divisors[i]
			}
		}
	}
	return res
}
