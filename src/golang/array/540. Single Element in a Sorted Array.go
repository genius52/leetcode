package array

func SingleNonDuplicate(nums []int) int {
	var l int = len(nums)
	var left int = 0
	var right int = l - 1
	for left < right{
		mid := (left + right)/2
		if nums[mid] != nums[mid + 1] && nums[mid] != nums[mid - 1]{
			return nums[mid]
		}
		if nums[mid] == nums[mid + 1]{
			mid_end := right - mid + 1
			if (mid_end | 1) == mid_end { //mid to end is odd
				left = mid + 2
			}else{//mid to end is even
				right = mid - 1
			}
		}else if nums[mid] == nums[mid - 1]{
			begin_mid := mid - left + 1
			if (begin_mid | 1) == begin_mid {  //total length is odd,begin to mid is odd
				right = mid - 2
			}else{//total length is even,begin to mid is odd
				left = mid + 1
			}
		}
	}
	return nums[left]
}