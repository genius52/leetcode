package array

func dfs_minOperations(nums1 []int, nums2 []int, l int, idx int, max_num1 int, max_num2 int, swap bool, memo [][2]int) int {
	if idx < 0 {
		return 0
	}
	if (nums1[idx] > max_num1 && nums1[idx] > max_num2) || (nums2[idx] > max_num1 && nums2[idx] > max_num2) {
		return -1
	}
	if swap {
		if memo[idx][1] != -2 {
			return memo[idx][1]
		}
	} else {
		if memo[idx][0] != -2 {
			return memo[idx][0]
		}
	}

	var res int = 1<<31 - 1
	//NOT Swap
	if nums1[idx] <= max_num1 && nums2[idx] <= max_num2 {
		ret := dfs_minOperations(nums1, nums2, l, idx-1, max_num1, max_num2, swap, memo)
		if ret != -1 {
			res = min_int(res, ret)
		}
	}
	//Swap
	if nums1[idx] <= max_num2 && nums2[idx] <= max_num1 {
		ret := dfs_minOperations(nums1, nums2, l, idx-1, max_num1, max_num2, swap, memo)
		if ret != -1 {
			res = min_int(res, 1+ret)
		}
	}
	if res == 1<<31-1 {
		return -1
	}
	if swap {
		memo[idx][1] = res
	} else {
		memo[idx][0] = res
	}
	return res
}

func MinOperations2934(nums1 []int, nums2 []int) int {
	var l int = len(nums1)
	var memo [][2]int = make([][2]int, l)
	for i := 0; i < l; i++ {
		memo[i][0] = -2
		memo[i][1] = -2
	}
	ret1 := dfs_minOperations(nums1, nums2, l, l-2, nums1[l-1], nums2[l-1], false, memo)
	ret2 := dfs_minOperations(nums1, nums2, l, l-2, nums2[l-1], nums1[l-1], true, memo)
	if ret2 != -1 {
		ret2++
	}
	return min_int(ret1, ret2)
}

func minOperations2934(nums1 []int, nums2 []int) int {
	var l int = len(nums1)
	var not_swap_last int = 0
	var swap_last int = 1
	var max_val int = max_int(nums1[l-1], nums2[l-1])
	var min_val int = min_int(nums1[l-1], nums2[l-1])
	for i := 0; i < l-1; i++ {
		if nums1[i] > max_val || nums2[i] > max_val {
			return -1
		}
		if min_int(nums1[i], nums2[i]) > min_val {
			return -1
		}
		if (nums1[i] > nums1[l-1]) || (nums2[i] > nums2[l-1]) {
			not_swap_last++
		}
		if (nums1[i] > nums2[l-1]) || (nums2[i] > nums1[l-1]) {
			swap_last++
		}
	}
	return min_int(swap_last, not_swap_last)
}
