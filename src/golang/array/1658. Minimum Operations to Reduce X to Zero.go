package array

import "math"

func MinOperations1658(nums []int, x int) int{
	var l int = len(nums)
	var total int = 0
	for i := 0;i < l;i++{
		total += nums[i]
	}
	var target int = total - x
	if target < 0{
		return -1
	}
	i := 0
	j := 0
	var sum int = 0
	var steps1 int = math.MaxInt32
	for i < l && j < l{
		for j < l && sum < target{
			sum += nums[j]
			j++
		}
		if sum == target{
			cur_step := l - (j - i)
			if cur_step < steps1{
				steps1 = cur_step
			}
		}
		sum -= nums[i]
		i++
	}
	var steps2 int = math.MaxInt32
	i = l - 1
	j = l - 1
	sum = 0
	for i >=0 && j >= 0{
		for i >= 0 && sum < target{
			sum += nums[i]
			i--
		}
		if sum == target{
			cur_step := l - (j - i)
			if cur_step < steps2{
				steps2 = cur_step
			}
			break
		}
		sum -= nums[j]
		j--
	}
	if steps1 == math.MaxInt32 && steps2 == math.MaxInt32{
		return -1
	}
	if steps1 == math.MaxInt32{
		return steps2
	}
	if steps2 == math.MaxInt32{
		return steps1
	}
	if steps1 < steps2{
		return steps1
	}else{
		return steps2
	}
}

//func dp_minOperations(nums []int,head int,tail int,choose_head bool,x int)int{
//	if x == 0{
//		return 0
//	}
//	if x < 0{
//		return -1
//	}
//	if head > tail{
//		return -1
//	}
//	var res1 int = -1
//	var res2 int = -1
//	if choose_head{
//		res1 = dp_minOperations(nums,head + 1,tail,choose_head,x - nums[head])
//		res2 = dp_minOperations(nums,head + 1,tail,!choose_head,x - nums[head])
//
//	}else{
//		res1 = dp_minOperations(nums,head,tail - 1,choose_head,x - nums[tail])
//		res2 = dp_minOperations(nums,head,tail - 1,!choose_head,x - nums[tail])
//	}
//	if res1 == -1 && res2 == -1{
//		return -1
//	}
//	if res1 == -1{
//		return res2 + 1
//	}
//	if res2 == -1{
//		return res1 + 1
//	}
//	if res1 < res2{
//		return res1 + 1
//	}else{
//		return res2 + 1
//	}
//}
//
//func MinOperations1658(nums []int, x int) int {
//	var l int = len(nums)
//	res1 := dp_minOperations(nums,0,l - 1,true,x)
//	res2 := dp_minOperations(nums,0,l - 1,false,x)
//	if res1 == -1 && res2 == -1{
//		return -1
//	}
//	if res1 == -1{
//		return res2
//	}
//	if res2 == -1{
//		return res1
//	}
//	if res1 < res2{
//		return res1
//	}else{
//		return res2
//	}
//}