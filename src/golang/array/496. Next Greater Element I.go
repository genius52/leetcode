package array

import "container/list"

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	var l1 int = len(nums1)
	var l2 int = len(nums2)
	var q list.List
	var record map[int]int = make(map[int]int)
	for i := 0;i < l2;i++{
		for q.Len() > 0 && nums2[i] > q.Back().Value.(int){
			record[q.Back().Value.(int)] = nums2[i]
			q.Remove(q.Back())
		}
		q.PushBack(nums2[i])
	}
	var res []int = make([]int,l1)
	for i := 0;i < l1;i++{
		if val,ok := record[nums1[i]];ok{
			res[i] = val
		}else{
			res[i] = -1
		}
	}
	return res
}