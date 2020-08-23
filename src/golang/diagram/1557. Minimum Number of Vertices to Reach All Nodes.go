package diagram

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	var points []bool = make([]bool,n)
	var indegree map[int]int = make(map[int]int)
	for _,edge := range edges{
		indegree[edge[1]]++
		points[edge[1]] = true
	}
	var res []int
	for index,p := range points{
		if !p{
			res = append(res,index)
		}
	}
	return res
}