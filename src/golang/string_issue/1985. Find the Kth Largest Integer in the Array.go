package string_issue

import (
	"sort"
)

//type mystring []string
//
//func (s mystring) Less(i, j int) bool {
//	if len(s[i]) < len(s[j]){
//		return true
//	}else if len(s[i]) > len(s[j]){
//		return false
//	}
//	l1 := len(s[i])
//	visit := 0
//	for visit < l1{
//		if s[i][visit] < s[j][visit]{
//			return true
//		}else if s[i][visit] > s[j][visit]{
//			return false
//		}else{
//			visit++
//		}
//	}
//	return true
//}
//
//func (s mystring) Swap(i, j int) {
//	s[i], s[j] = s[j], s[i]
//}
//
//func (s mystring) Len() int {
//	return len(s)
//}
//
//func kthLargestNumber(nums []string, k int) string {
//	var l int = len(nums)
//	sort.Sort(mystring(nums))
//	return nums[l - k]
//}

func kthLargestNumber(nums []string, k int) string{
	sort.Slice(nums, func(i, j int) bool {
		var l1 int = len(nums[i])
		var l2 int = len(nums[j])
		if l1 != l2{
			return l1 < l2
		}
		for k := 0;k < l1;k++{
			if nums[i][k] != nums[j][k]{
				return nums[i][k] < nums[j][k]
			}
		}
		return true
	})
	var l int = len(nums)
	return nums[l - k]
}