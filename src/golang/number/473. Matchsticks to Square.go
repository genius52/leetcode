package number

import "sort"

//Input: [1,1,2,2,2]
//Output: true
func dfs_makesquare(nums []int,pos int,sum []int,target int)bool{
	if pos == len(nums) - 1 {
		if sum[0] == target && sum[0] == sum[1] && sum[0] == sum[2]  {
			return true
		} else{
			return false
		}
	}
	for i := 0;i < 4;i++{
		if (sum[i] + nums[pos])> target{
			continue
		}
		sum[i] += nums[pos]
		if dfs_makesquare(nums,pos + 1,sum,target){
			return true
		}
		sum[i] -= nums[pos]
	}
	return false
}

func Makesquare(nums []int) bool {
	var l int = len(nums)
	if l < 4{
		return false
	}
	var total int = 0
	for _,n := range nums{
		total += n
	}
	var target int = total / 4
	if total % 4 != 0{
		return false
	}
	sort.Ints(nums)
	var sum []int = make([]int,4)
	return dfs_makesquare(nums,0,sum,target)
}
