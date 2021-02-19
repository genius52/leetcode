package array

//Input: n = 4, k = 2
//Output:
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//]
func dfs_combine(n int,cur_num int,cnt int,prev *[]int,res *[][]int){
	if cnt == 0{
		var data []int = make([]int,len(*prev))
		copy(data,*prev)
		*res = append(*res,data)
		return
	}
	if cur_num > n{
		return
	}
	//choose current number
	*prev = append(*prev,cur_num)
	dfs_combine(n,cur_num + 1,cnt - 1, prev,res)
	//skip current num
	var l int = len(*prev)
	if l > 0{
		*prev = (*prev)[:l-1]
	}
	dfs_combine(n,cur_num + 1,cnt,prev,res)
}

func Combine(n int, k int) [][]int {
	var res [][]int
	var choose []int
	dfs_combine(n,1,k,&choose,&res)
	return res
}