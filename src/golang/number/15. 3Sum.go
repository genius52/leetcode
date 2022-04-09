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

func threeSum(nums []int) [][]int{
	sort.Ints(nums)
	var res [][]int
	var l int = len(nums)
	for i := 0;i < l - 2;i++{
		if nums[i] > 0{
			break
		}
		if i != 0 && nums[i] == nums[i - 1]{
			continue
		}
		var left int = i + 1
		var right int = l - 1
		if nums[i] + nums[left] > 0{
			break
		}
		for left < right{
			if left != i + 1 && left < right && nums[left] == nums[left - 1]{
				left++
				continue
			}
			if right != l - 1 && left < right && nums[right] == nums[right + 1]{
				right++
				continue
			}
			if left >= right{
				break
			}
			if nums[i] + nums[left] > 0 || nums[i] + nums[right] * 2 < 0{
				break
			}
			if nums[i] + nums[left] + nums[right] == 0{
				res = append(res,[]int{nums[i],nums[left],nums[right]})
				left++
				right--
			}else if nums[i] + nums[left] + nums[right] < 0{
				left++
			}else if nums[i] + nums[left] + nums[right] > 0{
				right--
			}
		}
	}
	return res
}