package tree

func smallestMissingValueSubtree(parents []int, nums []int) []int {
	var l int = len(parents)
	var graph [][]int = make([][]int,l)
	for i := 0;i < l;i++{
		if parents[i] != -1{
			graph[parents[i]] = append(graph[parents[i]],i)
		}
	}
	var res []int = make([]int,l)
	var dfs func(int) map[int]bool
	dfs = func(i int)map[int]bool{
		var visited map[int]bool = make(map[int]bool)
		res[i] = 1
		for _,child := range graph[i]{
			ret := dfs(child)
			if res[child] > res[i]{
				res[i] = res[child]
			}
			if len(ret) > len(visited){
				visited,ret = ret,visited
			}
			for val,_ := range ret{
				visited[val] = true
			}
		}
		visited[nums[i]] = true
		for visited[res[i]]{
			res[i]++
		}
		return visited
	}
	dfs(0)
	return res
}