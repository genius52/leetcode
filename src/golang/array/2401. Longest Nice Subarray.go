package array

func LongestNiceSubarray(nums []int) int {
	var l int = len(nums)
	var one_cnt [32]int
	var res int = 0
	var left int = 0
	var right int = 0
	for left < l && right < l{
		for right < l{
			var n int = nums[right]
			var offset int = 0
			var over bool = false
			for n > 0{
				if n | 1 == n{
					one_cnt[offset]++
					if one_cnt[offset] > 1{
						over = true
					}
				}
				n >>= 1
				offset++
			}
			if over{
				break
			}
			right++
		}
		cur := right - left
		if cur > res{
			res = cur
		}
		for left < l{
			var n int = nums[left]
			var offset int = 0
			for n > 0{
				if n | 1 == n{
					one_cnt[offset]--
				}
				n >>= 1
				offset++
			}
			var meet bool = true
			for i := 0;i < 32;i++{
				if one_cnt[i] > 1{
					meet = false
					break
				}
			}
			left++
			if meet{
				break
			}
		}
		right++
	}
	return res
}