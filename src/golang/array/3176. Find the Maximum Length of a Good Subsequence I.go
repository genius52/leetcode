package array

func dp_maximumLength(nums []int, l int, idx int, k int, cnt int, last int, memo [][][]int) int {
	if cnt > k {
		return -(1 << 31)
	}
	if idx == l {
		return 0
	}
	if memo[idx][last][cnt] != -(1 << 31) {
		return memo[idx][last][cnt]
	}

	var res int = 0
	if nums[idx] == last {
		res = 1 + dp_maximumLength(nums, l, idx+1, k, cnt, last, memo) //add current
	} else {
		ret1 := 0 //add current
		if last == 0 {
			ret1 = 1 + dp_maximumLength(nums, l, idx+1, k, cnt, nums[idx], memo)
		} else {
			ret1 = 1 + dp_maximumLength(nums, l, idx+1, k, cnt+1, nums[idx], memo)
		}
		ret2 := dp_maximumLength(nums, l, idx+1, k, cnt, last, memo) //skip current
		res = max_int(ret1, ret2)
	}
	memo[idx][last][cnt] = res
	return res
}

func MaximumLength(nums []int, k int) int {
	var l int = len(nums)
	//var memo [501][501][26]int
	var memo [][][]int = make([][][]int, 501)
	for i := 0; i <= 500; i++ {
		memo[i] = make([][]int, 501)
		for j := 0; j <= 500; j++ {
			memo[i][j] = make([]int, 26)
			for m := 0; m < 26; m++ {
				memo[i][j][m] = -(1 << 31)
			}
		}
	}
	return dp_maximumLength(nums, l, 0, k, 0, 0, memo)
}
