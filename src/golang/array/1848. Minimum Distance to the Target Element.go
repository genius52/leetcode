package array

func getMinDistance(nums []int, target int, start int) int {
	var l int = len(nums)
	var min_diff int = 2147483647
	for i := 0; i < l; i++ {
		if nums[i] == target {
			cur := abs_int(i - start)
			if cur < min_diff {
				min_diff = cur
				if min_diff == 0 {
					break
				}
			}
		}
	}
	return min_diff
}

//func getMinDistance(nums []int, target int, start int) int {
//	var target_index []int
//	var l int = len(nums)
//	for i := 0;i < l;i++{
//		if nums[i] == target{
//			target_index = append(target_index,i)
//		}
//	}
//	var min_diff int = 2147483647
//	for _,v := range target_index{
//		min_diff = min_int(min_diff,int(math.Abs(float64(v - start))))
//	}
//	return min_diff
//}
