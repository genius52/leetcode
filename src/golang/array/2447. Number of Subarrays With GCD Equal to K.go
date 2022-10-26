package array

func gcd(a int,b int)int{
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func SubarrayGCD(nums []int, k int) int{
	var l int = len(nums)
	var res int = 0
	for i := 0;i < l;i++{
		if nums[i] < k || gcd(nums[i],k) < k{
			continue
		}
		if nums[i] == k{
			res++
		}
		var common int = nums[i]
		for j := i + 1;j < l;j++{
			if nums[j] < k{
				break
			}
			common = gcd(common,nums[j])
			if common == k{
				res++
			}
			if common < k{
				break
			}
		}
	}
	return res
}

//func SubarrayGCD(nums []int, k int) int {
//	var l int = len(nums)
//	var record []int = make([]int,l)
//	for i := 0;i < l;i++{
//		record[i] = gcd(nums[i],k)
//	}
//	var res int = 0
//	var bad_idx int = -1
//	var k_idx_cnt int = 0
//	for i := 0;i < l;i++{
//		if record[i] == 1{
//			bad_idx = i
//			k_idx_cnt = 0
//		}else{
//			if nums[i] == k{
//				k_idx_cnt++
//				res += i - bad_idx
//			}else{
//				res += k_idx_cnt
//			}
//		}
//	}
//	return res
//}