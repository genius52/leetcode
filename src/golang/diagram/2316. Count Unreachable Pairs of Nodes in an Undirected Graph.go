package diagram

func CountPairs(n int, edges [][]int) int64 {
	var groups []int = make([]int,n)
	for i := 0;i < n;i++{
		groups[i] = i
	}
	type funcType func (g *[]int,node int)int
	var get_parent funcType
	get_parent = func (g *[]int,node int)int{
		if (*g)[node] != node{
			(*g)[node] = get_parent(g,(*g)[node])
		}
		return (*g)[node]
	}
	for _,edge := range edges{
		group1 := get_parent(&groups,edge[0])
		group2 := get_parent(&groups,edge[1])
		if group1 != group2{
			if group1 < group2{
				groups[group2] = group1
			}else{
				groups[group1] = group2
			}
		}
	}
	var record map[int]int = make(map[int]int)
	for i := 0;i < n;i++{
		number := get_parent(&groups,i)
		record[number]++
	}
	var res int64 = 0
	for _,cnt := range record{
		res += int64(cnt) * int64(n - cnt)
	}
	return res
}