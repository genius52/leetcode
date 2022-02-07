package number

func countKDifference2(nums []int, k int) int{
	var cnt [101]int
	for _,n := range nums{
		cnt[n]++
	}
	var res int = 0
	for i := 1 + k;i <= 100;i++{
		res += cnt[i] * cnt[i - k]
	}
	return res
}

func countKDifference(nums []int, k int) int {
	var l int = len(nums)
	var res int = 0
	var record map[int]int = make(map[int]int)
	for i := 0;i < l;i++{
		neg := nums[i] - k
		pos := nums[i] + k
		if cnt,ok := record[neg];ok{
			res += cnt
		}
		if cnt,ok := record[pos];ok{
			res += cnt
		}
		record[nums[i]]++
	}
	return res
}