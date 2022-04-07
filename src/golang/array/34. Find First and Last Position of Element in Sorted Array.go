package array

func SearchRange(nums []int, target int) []int {
	var l int = len(nums)
	var res []int = make([]int,2)
	res[0] = -1
	res[1] = -1
	if l == 0{
		return res
	}
	var left int = 0
	var right int = l - 1
	for left < right{
		mid := (left + right) /2
		if nums[mid] == target{
			break
		}else if nums[mid] < target{
			left = mid + 1
		}else{
			right = mid - 1
		}
	}
	mid := (left + right)/2
	if nums[mid] != target{
		return res
	}
	var start int = left
	var end int = mid
	for start < end{
		mid2 := (start + end)/2
		if nums[mid2] != target{
			start = mid2 + 1
		}else{
			end = mid2
		}
	}
	res[0] = start
	start = mid
	end = right
	for start < end{
		mid2 := (start + end) /2
		if nums[mid2] == target{
			start = mid2
			if start + 1 == end{
				if nums[end] == target{
					start = end
				}else{
					break
				}
			}
		}else{
			end = mid2 - 1
		}
	}
	res[1] = start
	return res
}