package diagram

func get_root(groups []int,i int)int{
	if groups[i] == i{
		return i
	}
	return get_root(groups,groups[i])
}

func RemoveStones(stones [][]int) int {
	var l int = len(stones)
	var groups []int = make([]int,l)
	for i := 0;i < l;i++{
		groups[i] = i
	}
	for i := 0;i < l;i++{
		for j := i + 1;j < l;j++{
			if stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1]{
				//找到祖先节点所在的组号，合并2个组
				groups[get_root(groups,i)] = get_root(groups,j)
			}
		}
	}
	var res int = 0
	for i := 0;i < l;i++{
		if groups[i] == i{
			res++
		}
	}
	return l - res
}