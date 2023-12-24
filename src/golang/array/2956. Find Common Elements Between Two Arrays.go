package array

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	var record1 map[int]bool = make(map[int]bool)
	var record2 map[int]bool = make(map[int]bool)
	for _, n := range nums1 {
		record1[n] = true
	}
	for _, n := range nums2 {
		record2[n] = true
	}
	var res []int = make([]int, 2)
	for _, n := range nums1 {
		if _, ok := record2[n]; ok {
			res[0]++
		}
	}
	for _, n := range nums2 {
		if _, ok := record1[n]; ok {
			res[1]++
		}
	}
	return res
}
