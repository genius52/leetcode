package array

func quick_sort(nums []int,low int,high int){
	if low >= high{
		return
	}
	var l int = low
	var h int = high
	var tag = nums[low]
	for l < h{
		for h > l && nums[h] > tag{
			h--
		}
		if l < h{
			nums[l] = nums[h]
			l++
		}
		for l < h && nums[l] < tag{
			l++
		}
		if l < h{
			nums[h] = nums[l]
			h--
		}
	}
	nums[l] = tag
	quick_sort(nums,low,l-1)
	quick_sort(nums,l+1,high)
}

func SortArray(nums []int) []int {
	quick_sort(nums,0,len(nums)-1)
	return nums
}