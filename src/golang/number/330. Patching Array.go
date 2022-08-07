package number

//Input: nums = [1,5,10], n = 20
//Output: 2


//Input: nums = [1,5,10], n = 20
//Output: 2
func MinPatches(nums []int, n int) int{
	var l int = len(nums)
	var res int = 0
	var need int = 1
	var index int = 0
	for need <= n{
		if index >= l{//增加数字need
			need = need * 2
			res++
			continue
		}
		if nums[index] > need{ //增加数字need
			need = need * 2
			res++
		}else{//[1，need + nums[index] - 1] 都被覆盖到了
			need = need + nums[index]
			index++
		}
	}
	return res
}

//func MinPatches(nums []int, n int) int {
//	var l int = len(nums)
//	var res int = 0
//	var index int = 0
//	var need int = 1
//	//at begining we can generate 0,so we need 1
//	for need <= n{
//		if index >= l{
//			need = need * 2
//			res++
//			continue
//		}
//		if nums[index] <= need{
//			need = need + nums[index]
//			index++
//		}else{
//			need = need * 2
//			res++
//		}
//	}
//	return res
//}

//func combine_minPatches(nums []int,l int,pos int,cur_sum int,high int,record map[int]bool){
//	if cur_sum > high{
//		return
//	}
//	if pos >= l{
//		record[cur_sum] = true
//		return
//	}
//	//select current
//	combine_minPatches(nums,l,pos + 1,cur_sum + nums[pos],high,record)
//	//skip current
//	combine_minPatches(nums,l,pos + 1,cur_sum,high,record)
//	//restart
//	combine_minPatches(nums,l,pos + 1,nums[pos],high,record)
//}

//func MinPatches(nums []int, n int) int {
//	var l int = len(nums)
//	var record map[int]bool = make(map[int]bool)
//	combine_minPatches(nums,l,0,0,n,record)
//	var fail_nums []int
//	for i := 1;i <= n;i++{
//		if !record[i]{
//			fail_nums = append(fail_nums,i)
//		}
//	}
//	var l2 int = len(fail_nums)
//	if l2 == 0{
//		return 0
//	}
//	sort.Ints(fail_nums)
//	var cnt int = 0
//	for len(fail_nums) > 0{
//		for num,_ := range record{
//			record[num + fail_nums[0]] = true
//		}
//		fail_nums = []int{}
//		for i := 1;i <= n;i++{
//			if !record[i]{
//				fail_nums = append(fail_nums,i)
//			}
//		}
//		cnt++
//	}
//	return cnt
//}