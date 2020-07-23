package array

func GardenNoAdj(N int, paths [][]int) []int {
	var res []int = make([]int,N)
	var graph map[int][]int = make(map[int][]int)
	for _,v := range paths{
		graph[v[0]] = append(graph[v[0]],v[1])
		graph[v[1]] = append(graph[v[1]],v[0])
	}
	for start_point,v := range graph{
		if(res[start_point - 1] != 0){
			continue
		}
		var used_color []bool = make([]bool,4)
		for _,neighbours := range v{
			if(res[neighbours - 1] != 0){
				used_color[res[neighbours - 1] - 1] = true
			}
		}
		for c,used := range used_color{
			if(!used){
				res[start_point - 1] = c + 1
				break
			}
		}
	}
	for n,color := range res{
		if(color == 0) {
			res[n] = 1
		}
	}
	return res
}
