package diagram

import "math"

//Input: nums = [4,6,15,35]
//Output: 4

//func find_root(groups []int,i int)int {
//	if groups[i] == i{
//		return i
//	}
//	parent := find_root(groups,groups[i])
//	return parent
//}


func LargestComponentSize(nums []int) int{
	var l int = len(nums)
	var max_num int = 0
	for i := 0;i < l;i++{
		if nums[i] > max_num{
			max_num = nums[i]
		}
	}
	var groups []int = make([]int,max_num + 1)
	for i := 0;i <= max_num;i++{
		groups[i] = i
	}
	for i := 0;i < l;i++{
		//if (nums[i] % 2) == 0{
		//	groups[find_root(groups,nums[i])] = groups[find_root(groups,2)]
		//	groups[find_root(groups,nums[i])] = groups[find_root(groups,nums[i] / 2)]
		//}
		limit := int(math.Sqrt(float64(nums[i])))
		for j := 2;j <= limit;j++{
			if (nums[i] % j) == 0{
				mod := nums[i]/j
				//合并左边的组和右边的组，nums[0] = 6,groups[6] = 0,groups[2] = 0,groups[3] = 0
				groups[find_root(groups,nums[i])] = groups[find_root(groups,j)]
				groups[find_root(groups,nums[i])] = groups[find_root(groups,mod)]
			}
		}
	}
	var record map[int]int = make(map[int]int)
	var res int = 0
	for i := 0;i < l;i++{
		group_num := find_root(groups,nums[i])
		if _,ok := record[group_num];ok{
			record[group_num]++
		}else{
			record[group_num] = 1
		}
		if record[group_num] > res{
			res = record[group_num]
		}
	}
	return res
}

//TLE
//func gcd_largestComponentSize(a, b int) int {
//	if b == 0 {
//		return a
//	}
//	return gcd_largestComponentSize(b, a%b)
//}
//
//func find_root_largestComponentSize(groups []int,i int)int {
//	if groups[i] == i{
//		return i
//	}
//	parent := find_root_largestComponentSize(groups,groups[i])
//	return parent
//}
//
//func LargestComponentSize(nums []int) int {
//	var l int = len(nums)
//	var groups []int = make([]int,l)
//	for i := 0;i < l;i++{
//		groups[i] = i
//	}
//	for i := 1;i < l;i++{
//		for j := 0;j < i;j++{
//			m := gcd_largestComponentSize(nums[i],nums[j])
//			if m > 1{
//				groups[find_root_largestComponentSize(groups,j)] = i
//			}
//		}
//	}
//	var res int = 0
//	var record map[int]int = make(map[int]int)
//	for i := 0;i < l;i++{
//		group_num := find_root_largestComponentSize(groups,i)
//		if _,ok := record[group_num];ok{
//			record[group_num]++
//		}else{
//			record[group_num] = 1
//		}
//		if record[group_num] > res{
//			res = record[group_num]
//		}
//	}
//	return res
//}