package diagram

func dfs_maximumRequests(requests [][]int,l int,n int,pos int,in_out []int,cur_cnt int,max_cnt *int){
	//if pos >= l{
		if check_maximumRequests(in_out,n){
			*max_cnt = max_int(*max_cnt,cur_cnt)
		}
	//}
	for i := pos;i < l;i++{
		//Do not choose current
		//dfs_maximumRequests(requests,l,n,i + 1,in_out,cur_cnt,max_cnt)
		in_out[requests[i][0]]--
		in_out[requests[i][1]]++
		cur_cnt++
		dfs_maximumRequests(requests,l,n,i + 1,in_out,cur_cnt,max_cnt)
		cur_cnt--
		in_out[requests[i][0]]++
		in_out[requests[i][1]]--
	}
}

func check_maximumRequests(in_out []int,l int)bool{
	for i := 0;i < l;i++{
		if in_out[i] != 0{
			return false
		}
	}
	return true
}

func MaximumRequests(n int, requests [][]int) int {
	var l int = len(requests)
	var in_out []int = make([]int,n)
	var res int = 0
	dfs_maximumRequests(requests,l,n,0,in_out,0,&res)
	return res
}