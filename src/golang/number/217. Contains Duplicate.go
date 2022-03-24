package number

import "sort"

func containsDuplicate(nums []int) bool{
	sort.Ints(nums)
	var l int = len(nums)
	for i := 1;i < l;i++{
		if nums[i] == nums[i - 1]{
			return true
		}
	}
	return false
}

//func containsDuplicate(nums []int) bool {
//	var record map[int]bool = make(map[int]bool)
//	for _,n := range nums{
//		if _,ok := record[n];ok{
//			return true
//		}
//		record[n] = true
//	}
//	return false
//}