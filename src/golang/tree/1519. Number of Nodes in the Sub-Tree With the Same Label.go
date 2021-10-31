package tree

func dfs_countSubTrees(graph *[]map[int]bool,cur int,labels string,res *[]int)[26]int{
	var total [26]int
	for k,_ := range (*graph)[cur]{
		child := dfs_countSubTrees(graph,k,labels,res)
		for i := 0;i < 26;i++{
			total[i] += child[i]
		}
	}
	total[labels[cur] - 'a']++
	(*res)[cur] = total[labels[cur] - 'a']
	return total
}

func cut_edge(graph *[]map[int]bool,cur int,parent int){
	for k,_ := range (*graph)[cur]{
		if k != parent{
			delete((*graph)[k],cur)
			cut_edge(graph,k,cur)
		}
	}
}

func CountSubTrees(n int, edges [][]int, labels string) []int {
	var l int = len(edges)
	var graph []map[int]bool = make([]map[int]bool,n)
	for i := 0;i < l;i++{
		if graph[edges[i][0]] == nil{
			graph[edges[i][0]] = make(map[int]bool)
		}
		graph[edges[i][0]][edges[i][1]] = true
		if graph[edges[i][1]] == nil{
			graph[edges[i][1]] = make(map[int]bool)
		}
		graph[edges[i][1]][edges[i][0]] = true
	}
	//删除子节点指向自己的
	cut_edge(&graph,0,-1)
	var res []int = make([]int,n)
	dfs_countSubTrees(&graph,0,labels,&res)
	return res
}