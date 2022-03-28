package array

func findDifference(nums1 []int, nums2 []int) [][]int {
	var record1 map[int]bool = make(map[int]bool)
	var record2 map[int]bool = make(map[int]bool)
	for _,n := range nums1{
		record1[n] = true
	}
	for _,n := range nums2{
		record2[n] = true
	}
	var res [][]int = make([][]int,2)
	for n1,_ := range record1{
		if _,ok := record2[n1];!ok{
			res[0] = append(res[0],n1)
		}
	}
	for n2,_ := range record2{
		if _,ok := record1[n2];!ok{
			res[1] = append(res[1],n2)
		}
	}
	return res
}