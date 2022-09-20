package array

func SmallestSubarrays(nums []int) []int {
	var l int = len(nums)
	var res []int = make([]int,l)
	var bit_idx [32]int //第i位为1的最左的索引位置
	for i := 0;i < 32;i++{
		bit_idx[i] = -1
	}
	for i := l - 1;i >= 0;i--{
		var n int = nums[i]
		var offset int = 0
		for n > 0{
			if n | 1 == n{
				bit_idx[offset] = i
			}
			n >>= 1
			offset++
		}
		var right int = -1
		for j := 0;j < 32;j++{
			right = max_int(right,bit_idx[j])
		}
		if right == -1{
			res[i] = 1
		}else{
			res[i] = right - i + 1
		}
	}
	return res
}