package tree


func minimumScore(nums []int, edges [][]int) int {
	var l int = len(nums)
	var edge_cnt int = len(edges)
	var graph [][]int = make([][]int,l)
	for _,edge := range edges{
		graph[edge[0]] = append(graph[edge[0]],edge[1])
		graph[edge[1]] = append(graph[edge[1]],edge[0])
	}
	var sum int = 0
	for i := 0;i < l;i++{
		sum ^= nums[i]
	}
	var res int = 2147483647
	for i := 0;i < edge_cnt;i++{
		for j := i + 1;j < edge_cnt;j++{

		}
	}
	return res
}