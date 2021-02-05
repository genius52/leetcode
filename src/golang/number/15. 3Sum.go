package number

import (
	"sort"
	"strconv"
)

//Input: nums = [-1,0,1,2,-1,-4]
//Output: [[-1,-1,2],[-1,0,1]]
func ThreeSum(nums []int) [][]int {
	var res [][]int
	var l int = len(nums)
	var visited map[string]bool = make(map[string]bool)
	for i := 0;i < l - 2;i++{
		var target int = nums[i]
		var record map[int]bool = make(map[int]bool)
		for j := i + 1;j < l;j++{
			if _,ok := record[-target - nums[j]];ok{
				var data []int = []int{target,-target - nums[j],nums[j]}
				sort.Ints(data)
				k := strconv.Itoa(data[0]) + "|" + strconv.Itoa(data[1]) + "|" + strconv.Itoa(data[2])
				if _,ok := visited[k];!ok{
					res = append(res,data)
					visited[k] = true
				}
			}
			record[nums[j]] = true
		}
	}
	return res
}