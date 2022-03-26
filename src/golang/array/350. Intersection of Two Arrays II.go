package array

func intersect(nums1 []int, nums2 []int) []int {
	var record1 map[int]int = make(map[int]int)
	var record2 map[int]int = make(map[int]int)
	for _,n := range nums1{
		record1[n]++
	}
	for _,n := range nums2{
		record2[n]++
	}
	var res []int
	if len(record1) < len(record2){
		for n,cnt1 := range record1{
			if cnt2,ok := record2[n];ok{
				for i := 0;i < min_int(cnt1,cnt2);i++{
					res = append(res,n)
				}
			}
		}
	}else{
		for n,cnt2 := range record2{
			if cnt1,ok := record1[n];ok{
				for i := 0;i < min_int(cnt1,cnt2);i++{
					res = append(res,n)
				}
			}
		}
	}
	return res
}