package array

//Input: nums = [2,2,2,1,2,2,1,2,2,2], k = 2
//Output: 16
func NumberOfSubarrays(nums []int, k int) int{
	var l int = len(nums)
	var res int = 0
	var record map[int]int = make(map[int]int)//key:odd cnt;  value:subarray cnt
	// for k = 1, [1], --------- the answer is 1: [1] + [ ]
	//Therefore, k = 1, [2,2,1], ---------- the answer 3: [2,2] + [1], [2] + [1], [ ] + [1]
	//[ ] can be count as one.
	record[0] = 1
	var odd_cnt int = 0
	for i := 0;i < l;i++{
		if (nums[i] | 1) == nums[i]{
			odd_cnt++
		}
		record[odd_cnt]++
		if odd_cnt >= k{
			if val,ok := record[odd_cnt - k];ok{
				res += val
			}
		}
	}
	return res
}

//func findfirst_numberOfSubarrays(nums []int,l int,start int)int{
//	var visit int = start
//	for visit < l{
//		if nums[visit] | 1 == nums[visit]{
//			break
//		}
//		visit++
//	}
//	return visit
//}
//
//func findlast_numberOfSubarrays(nums []int,l int,start int,k int)(bool,int,int){
//	var visit int = start
//	var odd_cnt int = 0
//	var last_odd int = start
//	var find_last bool = false
//	for visit < l{
//		if nums[visit] | 1 == nums[visit]{
//			odd_cnt++
//		}
//		if odd_cnt == k && !find_last{
//			find_last = true
//			last_odd = visit
//		}
//		if odd_cnt == k + 1{
//			break
//		}
//		visit++
//	}
//	if !find_last{
//		return false,l,l
//	}
//	return true,last_odd,visit - 1
//}
//
//func NumberOfSubarrays(nums []int, k int) int {
//	var l int = len(nums)
//	//var odd_cnt int = 0
//	var res int = 0
//	var visit int = 0
//	for visit < l{
//		odd_first := findfirst_numberOfSubarrays(nums,l,visit)
//		ret,odd_last,odd_end := findlast_numberOfSubarrays(nums,l,odd_first,k)
//		if !ret{
//			break
//		}
//		res += (odd_first - visit + 1) * (odd_end - odd_last + 1)
//		visit = odd_first + 1
//	}
//	return res
//}