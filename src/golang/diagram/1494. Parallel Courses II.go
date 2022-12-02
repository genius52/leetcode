package diagram

//输入：n = 5, dependencies = [[2,1],[3,1],[4,1],[1,5]], k = 2
//输出：4
//解释：上图展示了题目输入的图。一个最优方案是：第一学期上课程 2 和 3，第二学期上课程 4 ，第三学期上课程 1 ，第四学期上课程 5 。

func dp_minNumberOfSemesters(graph [][]int,in []int,out []int,n int,k int,cur map[int]bool,status int,memo map[int]int)int{
	if status == 1 << n - 1{
		return 0
	}
	if v,ok := memo[status];ok{
		return v
	}
	var res int = 1 << 31 - 1
	for i := 1;i <= 1 << n - 1;i++{
		//有k个1,并且都在cur里
		var num int = i
		var offset int = 1
		var choose map[int]bool = make(map[int]bool)
		for num > 0{
			if (num | 1) == num{
				choose[offset] = true
			}
			num >>= 1
			offset++
		}
		if len(choose) > k ||len(choose) != len(cur){
			continue
		}
		var exist bool = true
		for key,_ := range choose{
			if _,ok := cur[key];!ok{
				exist = false
				break
			}
		}
		if !exist{
			continue
		}
		var next_status int = status
		//for key,_ := range choose{
		//
		//}
		ret := 1 + dp_minNumberOfSemesters(graph,in,out,n,k,cur,next_status,memo)
		res = min_int(res,ret)
	}
	memo[status] = res
	return res
}

func minNumberOfSemesters(n int, relations [][]int, k int) int {
	var in []int = make([]int,n + 1)
	var out []int = make([]int,n + 1)
	var graph [][]int = make([][]int,n + 1)
	for _,edge := range relations{
		graph[edge[0]] = append(graph[edge[0]],edge[1])
		graph[edge[1]] = append(graph[edge[1]],edge[0])
		in[edge[1]]++
		out[edge[0]]++
	}
	var cur map[int]bool
	for i := 1;i <= n;i++{
		if in[i] == 0{
			cur[i] = true
		}
	}
	var memo map[int]int = make(map[int]int)
	return dp_minNumberOfSemesters(graph,in,out,n,k,cur,0,memo)
}