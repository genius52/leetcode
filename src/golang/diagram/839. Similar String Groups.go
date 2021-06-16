package diagram

//DFS solution
func checksimilar_numSimilarGroups(s1 string,s2 string)bool{
	var l1 int = len(s1)
	var l2 int = len(s2)
	if l1 != l2{
		return false
	}
	var diff_cnt int = 0
	var index1 int = -1
	var index2 int = -1
	for i := 0;i < l1;i++{
		if s1[i] != s2[i]{
			diff_cnt++
			if index1 == -1{
				index1 = i
			}else {
				index2 = i
			}
		}
		if diff_cnt > 2{
			return false
		}
	}
	if diff_cnt == 0{
		return true
	}
	if diff_cnt != 2{
		return false
	}
	if s1[index1] == s2[index2] && s1[index2] == s2[index1]{
		return true
	}
	return false
}

func dfs_numSimilarGroups(strs []string,last string,visited map[string]bool){
	if _,ok := visited[last];ok{
		return
	}
	visited[last] = true
	for _,s := range strs{
		if checksimilar_numSimilarGroups(s,last){
			dfs_numSimilarGroups(strs,s,visited)
		}
	}
}

func numSimilarGroups(strs []string) int {
	var l int = len(strs)
	var visited map[string]bool = make(map[string]bool)
	var groups int = 0
	for i := 0;i < l;i++{
		if _,ok := visited[strs[i]];ok{
			continue
		}
		groups++
		dfs_numSimilarGroups(strs,strs[i],visited)
	}
	return groups
}

//Union find solution
func find_root(groups []int,i int)int {
	if groups[i] == i{
		return i
	}
	parent := find_root(groups,groups[i])
	//groups[i] = parent
	return parent
}

func numSimilarGroups2(strs []string) int {
	var l int = len(strs)
	var groups []int = make([]int,l)
	for i := 0;i < l;i++{
		groups[i] = i
	}
	for i := 1;i < l;i++{
		for j := 0;j < i;j++{
			if checksimilar_numSimilarGroups(strs[i],strs[j]){
				groups[find_root(groups,j)] = i//如果j和i相似，把j所在的root指向i，i作为新的root
			}
		}
	}
	var res int = 0
	for i := 0;i < l;i++{
		if groups[i] == i{
			res++
		}
	}
	return res
}