package diagram

import "sort"

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	type funcType func (g *[]int,node int)int
	var get_parent funcType
	get_parent = func (g *[]int,node int)int{
		if (*g)[node] != node{
			(*g)[node] = get_parent(g,(*g)[node])
		}
		return (*g)[node]
	}
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})
	var groups []int = make([]int,n)
	for i := 0;i < n;i++{
		groups[i] = i
	}
	groups[firstPerson] = 0
	var l int = len(meetings)
	k := 0
	for k < l{
		time := meetings[k][2]
		var record map[int]bool = make(map[int]bool)
		for k < l && meetings[k][2] == time{
			group1 := get_parent(&groups,meetings[k][0])
			group2 := get_parent(&groups,meetings[k][1])
			if group1 != group2{
				if group1 < group2{//组号小的优先
					groups[group2] = group1
				}else{
					groups[group1] = group2
				}
			}
			record[meetings[k][0]] = true
			record[meetings[k][1]] = true
			k++
		}
		for key,_ := range record{
			if get_parent(&groups,key) != 0{
				groups[key] = key
			}
		}
	}
	var res []int
	for i := 0;i < n;i++{
		if get_parent(&groups,i) == 0{
			res = append(res,i)
		}
	}
	return res
}