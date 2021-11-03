package array

func MaxSum(nums1 []int, nums2 []int) int{
	var l1 int = len(nums1)
	var l2 int = len(nums2)
	var visit1 int = 0
	var visit2 int = 0
	var sum1 int = 0
	var sum2 int = 0
	var sum int = 0
	for visit1 < l1 && visit2 < l2{
		if nums1[visit1] < nums2[visit2]{
			sum1 += nums1[visit1]
			visit1++
		}else if nums1[visit1] > nums2[visit2]{
			sum2 += nums2[visit2]
			visit2++
		}else{
			sum += nums1[visit1] + max_int(sum1,sum2)
			sum1 = 0
			sum2 = 0
			visit1++
			visit2++
		}
	}
	for visit1 < l1{
		sum1 += nums1[visit1]
		visit1++
	}
	for visit2 < l2{
		sum2 += nums2[visit2]
		visit2++
	}
	sum += max_int(sum1,sum2)
	return sum % (1e9 + 7)
}
//TLE
//func MaxSum(nums1 []int, nums2 []int) int {
//	var l1 int = len(nums1)
//	var l2 int = len(nums2)
//	var record1 map[int]int = make(map[int]int)//value:index
//	var record2 map[int]int = make(map[int]int)//value:index
//	var prefix1 []int = make([]int,l1 + 1)
//	var prefix2 []int = make([]int,l2 + 1)
//	for i := 0;i < l1;i++{
//		record1[nums1[i]] = i
//		prefix1[i + 1] = prefix1[i] + nums1[i]
//	}
//	for i := 0;i < l2;i++{
//		record2[nums2[i]] = i
//		prefix2[i + 1] = prefix2[i] + nums2[i]
//	}
//	var same_value_index [][2]int
//	for v1,i1 := range record1{
//		for v2,i2 := range record2{
//			if v1 == v2{
//				same_value_index = append(same_value_index,[2]int{i1,i2})
//			}
//		}
//	}
//	sort.Slice(same_value_index, func(i, j int) bool {
//		return same_value_index[i][0] < same_value_index[j][0]
//	})
//	if len(same_value_index) == 0{
//		return max_int(prefix1[l1],prefix2[l2])
//	}
//	var MOD int = 1e9 + 7
//	var l int = len(same_value_index)
//	var dp [][2]int = make([][2]int,l)//dp[i][0]:the biggest sum from start to nums1[same_value_index[i][0]]
//	for i := 0;i < l;i++{
//		idx1 := same_value_index[i][0]
//		idx2 := same_value_index[i][1]
//		if i == 0{
//			dp[i][0] = max_int(prefix1[idx1 + 1],prefix2[idx2 + 1])
//			dp[i][1] = max_int(prefix1[idx1 + 1],prefix2[idx2 + 1])
//		}else{
//			same1 := dp[i - 1][0] + prefix1[idx1 + 1] - prefix1[same_value_index[i - 1][0] + 1]
//			jump_from2 := dp[i - 1][1] + prefix2[idx2 + 1] - prefix2[same_value_index[i - 1][1] + 1]
//			dp[i][0] = max_int(same1,jump_from2)
//			same2 := dp[i - 1][1] + prefix2[idx2 + 1] - prefix2[same_value_index[i - 1][1] + 1]
//			jump_from1 := dp[i - 1][0] + prefix1[idx1 + 1] - prefix1[same_value_index[i - 1][0] + 1]
//			dp[i][1] = max_int(same2, jump_from1)
//		}
//	}
//	tail1 := prefix1[l1] - prefix1[same_value_index[l - 1][0] + 1]
//	tail2 := prefix2[l2] - prefix2[same_value_index[l - 1][1] + 1]
//	return max_int((dp[l - 1][0] + tail1), (dp[l - 1][1] + tail2))  % MOD
//}