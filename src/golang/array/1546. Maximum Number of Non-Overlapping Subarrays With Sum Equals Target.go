package array

func MaxNonOverlapping(nums []int, target int) int {
	var l int = len(nums)
	//var dp []int = make([]int,l + 1)//dp[i] = max count from index [0 to i)
	var record map[int]int = make(map[int]int)
	var res int = 0
	var sum int = 0
	for i := 0;i < l;i++{
		sum += nums[i]
		if cnt,ok := record[target - sum];ok{
			res += cnt
		}
		record[sum]++
	}
	return res;
}