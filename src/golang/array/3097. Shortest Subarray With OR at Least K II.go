package array

func MinimumSubarrayLength3097(nums []int, k int) int {
	if k == 0 {
		return 1
	}
	var l int = len(nums)
	var res int = 1<<31 - 1
	var left int = 0
	var right int = 0
	var sum int = 0
	var bit_cnt [32]int
	for left < l && right < l {
		for right < l && sum < k {
			sum |= nums[right]
			var n int = nums[right]
			var offset int = 0
			for n > 0 {
				if n|1 == n {
					bit_cnt[offset]++
				}
				offset++
				n >>= 1
			}
			right++
		}
		//if right == l {
		//	if sum >= k {
		//		cur_len := right - left
		//		if cur_len < res {
		//			res = cur_len
		//		}
		//	}
		//	break
		//}
		for left < right && sum >= k {
			cur_len := right - left
			if cur_len < res {
				res = cur_len
			}
			var n int = nums[left]
			var offset int = 0
			for n > 0 {
				if n|1 == n {
					bit_cnt[offset]--
					if bit_cnt[offset] == 0 {
						sum ^= 1 << offset
					}
				}
				offset++
				n >>= 1
			}
			left++
		}
	}
	if res == 1<<31-1 {
		return -1
	}
	return res
}
