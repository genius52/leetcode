package array

import "sort"

func MakeSimilar(nums []int, target []int) int64 {
	var odd_target []int
	var even_target []int
	for _,n := range target{
		if n | 1 == n{
			odd_target = append(odd_target,n)
		}else{
			even_target = append(even_target,n)
		}
	}
	sort.Ints(odd_target)
	sort.Ints(even_target)
	sort.Ints(nums)
	var l int = len(nums)
	//var odd_len int = len(odd_target)
	//var even_len int = len(even_target)
	var res int64 = 0
	var available_increase_steps int = 0//当前可用的减少次数
	var available_decrease_steps int = 0
	//var even_increase_steps int = 0
	//var even_decrease_steps int = 0
	var odd_idx int = 0
	var even_idx int = 0
	for i := 0;i < l;i++{
		if nums[i] | 1 == nums[i]{
			if nums[i] > odd_target[odd_idx]{
				need_decrease_steps := (nums[i] - odd_target[odd_idx]) / 2 //当前需要减少多少次
				if  available_decrease_steps >= need_decrease_steps{
					available_decrease_steps -= need_decrease_steps
				}else{
					need_decrease_steps -= available_decrease_steps
					available_decrease_steps = 0
					available_increase_steps += need_decrease_steps
					res += int64(need_decrease_steps)
				}
			}else if nums[i] < odd_target[odd_idx]{
				need_increase_steps := (odd_target[odd_idx] - nums[i]) / 2 //当前需要增加多少次
				if available_increase_steps >= need_increase_steps{
					available_increase_steps -= need_increase_steps
				}else{
					need_increase_steps -= available_increase_steps
					available_increase_steps = 0
					available_decrease_steps += need_increase_steps
					res += int64(need_increase_steps)
				}
			}
			odd_idx++
			//a = append(a[:i], a[i+1:]...) //删除中间1个元素
			//var odd_idx int = 0
			//var steps int = 0
			//for odd_idx < odd_len && abs_int(odd_target[odd_idx] - nums[i]) % 2 == 1{
			//	steps++
			//	odd_idx++
			//}
			//res += int64(abs_int(odd_target[odd_idx] - nums[i]))/2
			//odd_target = append(odd_target[:odd_idx],odd_target[odd_idx + 1:]...)
		}else{
			if nums[i] > even_target[even_idx]{
				need_decrease_steps := (nums[i] - even_target[even_idx]) / 2 //当前需要减少多少次
				if  available_decrease_steps >= need_decrease_steps{
					available_decrease_steps -= need_decrease_steps
				}else{
					need_decrease_steps -= available_decrease_steps
					available_decrease_steps = 0
					available_increase_steps += need_decrease_steps
					res += int64(need_decrease_steps)
				}
			}else if nums[i] < even_target[even_idx]{
				need_increase_steps := (even_target[even_idx] - nums[i]) / 2 //当前需要增加多少次
				if available_increase_steps >= need_increase_steps{
					available_increase_steps -= need_increase_steps
				}else{
					need_increase_steps -= available_increase_steps
					available_increase_steps = 0
					available_decrease_steps += need_increase_steps
					res += int64(need_increase_steps)
				}
			}
			even_idx++
		}
	}
	return res
}