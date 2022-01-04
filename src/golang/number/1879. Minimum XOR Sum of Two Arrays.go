package number

//输入：nums1 = [1,0,3], nums2 = [5,3,4]
//输出：8
//解释：将 nums2 重新排列得到 [5,4,3] 。
//异或值之和为 (1 XOR 5) + (0 XOR 4) + (3 XOR 3) = 4 + 4 + 0 = 8 。
func dp_minimumXORSum(nums1 []int, nums2 []int, l int, idx1 int, status int, memo map[int]int) int {
	if idx1 >= l {
		return 0
	}
	if _, ok := memo[status]; ok {
		return memo[status]
	}
	var res int = 2147483647
	for j := 0; j < l; j++ {
		if (status | (1 << j)) == status {
			continue
		}
		cur := nums1[idx1] ^ nums2[j]
		cur += dp_minimumXORSum(nums1, nums2, l, idx1+1, status|1<<j, memo)
		res = min_int(res, cur)
	}
	memo[status] = res
	return res
}

func MinimumXORSum(nums1 []int, nums2 []int) int {
	var memo map[int]int = make(map[int]int)
	return dp_minimumXORSum(nums1, nums2, len(nums1), 0, 0, memo)
}
