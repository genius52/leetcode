package number

//func gcd(a int, b int) int {
//	if b == 0 {
//		return a
//	}
//	return gcd(b, a%b)
//}

func CountDifferentSubsequenceGCDs(nums []int) int {
	var l int = len(nums)
	var record [200005]int
	var cnt int = 0
	var max_val int = 0
	for i := 0; i < l; i++ {
		record[nums[i]]++
		if record[nums[i]] == 1 {
			cnt++
		}
		if nums[i] > max_val {
			max_val = nums[i]
		}
	}
	for i := 1; i <= max_val; i++ {
		if record[i] > 0 {
			continue
		}
		var cur_gcd int = 0
		//check if num i is gcd of  more than one numbers
		for j := i; j <= max_val; j += i {
			if record[j] == 0 {
				continue
			}
			cur_gcd = gcd(j, cur_gcd)
			if cur_gcd == i {
				break
			}
		}
		if cur_gcd == i {
			cnt++
			//record[i]++
		}
	}
	return cnt
}
