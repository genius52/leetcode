package diagram

func dfs_validArrangement(graph map[int][]int,cur int,res *[][]int){
	for len(graph[cur]) > 0{
		var next int = graph[cur][len(graph[cur]) - 1]
		graph[cur] = graph[cur][:len(graph[cur]) - 1]
		//fmt.Println(strconv.Itoa(cur) + "," + strconv.Itoa(next))
		dfs_validArrangement(graph,next,res)
		*res = append(*res,[]int{cur,next})
	}
}

func ValidArrangement(pairs [][]int) [][]int {
	var outdegree map[int]int = make(map[int]int)
	var graph map[int][]int = make(map[int][]int)
	for _,p := range pairs{
		outdegree[p[0]]++
		outdegree[p[1]]--
		graph[p[0]] = append(graph[p[0]],p[1])
	}
	var start int = -1
	for k,v := range outdegree{
		if v == 1{
			start = k
		}
	}
	var res [][]int
	if start == -1{
		for k,_ := range outdegree{
			start = k
			break
		}
	}
	dfs_validArrangement(graph,start,&res)

	var l int = len(res)
	var res2 [][]int = make([][]int,l)
	for i := 0;i < l;i++{
		res2[i] = res[l - 1 - i]
	}
	return res2
}