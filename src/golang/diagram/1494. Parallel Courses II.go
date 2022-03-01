package diagram

import "container/list"

//输入：n = 5, dependencies = [[2,1],[3,1],[4,1],[1,5]], k = 2
//输出：4
//解释：上图展示了题目输入的图。一个最优方案是：第一学期上课程 2 和 3，第二学期上课程 4 ，第三学期上课程 1 ，第四学期上课程 5 。

func minNumberOfSemesters(n int, relations [][]int, k int) int {
	var in []int = make([]int,n + 1)
	var out []int = make([]int,n + 1)
	var graph [][]bool = make([][]bool,n + 1)
	for i := 0;i <= n;i++{
		graph[i] = make([]bool,n + 1)
	}
	for _,edge := range relations{
		graph[edge[0]][edge[1]] = true
		in[edge[1]]++
		out[edge[0]]++
	}
	var q list.List
	//将入度为0，并且深度最大的节点优先加入队列
	for i := 1;i <= n;i++{
		if in[i] == 0{
			q.PushBack(i)
		}
	}
	var steps int = 0
	for q.Len() > 0{

	}
	return steps
}