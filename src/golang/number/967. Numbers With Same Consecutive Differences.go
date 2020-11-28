package number

//Input: n = 3, k = 7
//Output: [181,292,707,818,929]
func dfs_numsSameConsecDiff(n int,nums []int,k int,res *[]int){
	if n == 0{
		var num int = 0
		for _,n := range nums{
			num *= 10
			num += n
		}
		*res = append(*res,num)
		return
	}
	var l int = len(nums)
	var add_val int = nums[l - 1] + k
	if add_val <= 9{
		var nums2 []int = make([]int,l)
		copy(nums2,nums)
		nums2 = append(nums2,add_val)
		dfs_numsSameConsecDiff(n - 1,nums2,k,res)
	}
	var minus_val int = nums[l - 1] - k
	if minus_val != add_val && nums[l - 1] - k >= 0{
		var nums3 []int = make([]int,l)
		copy(nums3,nums)
		nums3 = append(nums3,nums[l - 1] - k)
		dfs_numsSameConsecDiff(n - 1,nums3,k,res)
	}
}

func NumsSameConsecDiff(n int, k int) []int{
	var res []int
	var visited [10]bool
	for i := 1;i + k <= 9;i++{
		dfs_numsSameConsecDiff(n - 1,[]int{i},k,&res)
		visited[i] = true
	}
	for i := 9;i - k >= 0;i--{
		if i != 0 && !visited[i]{
			dfs_numsSameConsecDiff(n - 1,[]int{i},k,&res)
		}
	}
	return res
}

//func NumsSameConsecDiff(n int, k int) []int {
//	var record map[int]bool = make(map[int]bool)
//	for i := 1;i + k <= 9;i++{
//		var num int = i
//		var last_num int = i
//		for j := 1;j < n;j++{
//			if last_num == i{
//				last_num += k
//			}else{
//				last_num -= k
//			}
//			num *= 10
//			num += last_num
//		}
//		record[num] = true
//	}
//	for i := 9;i - k >= 0;i--{
//		if i == 0{
//			continue
//		}
//		var num int = i
//		var last_num int = i
//		for j := 1;j < n;j++{
//			if last_num == i{
//				last_num -= k
//			}else{
//				last_num += k
//			}
//			num *= 10
//			num += last_num
//		}
//		record[num] = true
//	}
//	var res []int
//	for n,_ := range record{
//		res = append(res,n)
//	}
//	return res
//}