package number

//Input: k = 3, n = 9
//Output: [[1,2,6], [1,3,5], [2,3,4]]
//func dfs_combinationSum3(cur_res []int,cur_num int,cnt_limit int,sum int,target_sum int,visited [9]int,record *[][]int){
//	if len(cur_res) == cnt_limit{
//		if sum == target_sum{
//			c := make([]int,cnt_limit)
//			copy(c,cur_res)
//			*record = append(*record,c)
//		}
//		return
//	}
//	if sum > target_sum || cur_num > 9{
//		return
//	}
//	if visited[cur_num - 1] == 0{
//		for i := cur_num;i <= 9;i++{
//			if sum + i > target_sum{
//				break
//			}
//			cur_res = append(cur_res,i)
//			visited[i - 1] = 1
//			dfs_combinationSum3(cur_res,i + 1,cnt_limit,sum + i,target_sum,visited,record)
//			cur_res = cur_res[:len(cur_res) - 1]
//			visited[i - 1] = 0
//		}
//	}else{
//		dfs_combinationSum3(cur_res,cur_num + 1,cnt_limit,sum,target_sum,visited,record)
//	}
//}
//
//func CombinationSum3(k int, n int) [][]int {
//	var visited [9]int = [9]int{0}
//	var record [][]int
//	dfs_combinationSum3([]int{},1,k,0,n,visited,&record)
//	return record
//}

func dfs_combinationSum3(cur_num int,nums []int,k int,target int,res *[][]int){
	if target == 0 && k == 0{
		var nums2 []int = make([]int,len(nums))
		copy(nums2,nums)
		*res = append(*res,nums2)
		return
	}
	if cur_num > 9 || k < 0 || target < 0{
		return
	}
	//add current number
	nums = append(nums,cur_num)
	dfs_combinationSum3(cur_num + 1,nums,k - 1,target - cur_num,res)
	nums = nums[:len(nums) - 1]
	//skip current number
	dfs_combinationSum3(cur_num + 1,nums,k,target,res)
}

func CombinationSum3(k int, n int) [][]int {
	var res [][]int
	dfs_combinationSum3(1,[]int{},k,n,&res)
	return res
}