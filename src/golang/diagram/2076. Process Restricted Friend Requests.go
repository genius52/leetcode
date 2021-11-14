package diagram

//func get_parent(groups []int,node int)int{
//	if groups[node] != node{
//		groups[node] = get_parent(groups,groups[node])
//	}
//	return groups[node]
//}

func FriendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	var groups []int = make([]int,n)//朋友在一个组里
	for i := 0;i < n;i++{
		groups[i] = i
	}
	var l int = len(requests)
	var res []bool = make([]bool,l)
	for i := 0;i < l;i++{
		group1 := get_parent(groups,requests[i][0])
		group2 := get_parent(groups,requests[i][1])
		if group1 == group2{
			res[i] = true

		}else{
			var valid bool = true
			for _,restric := range restrictions{
				group3 := get_parent(groups,restric[0])
				group4 := get_parent(groups,restric[1])
				//group3 和 group4 不可能是朋友,如果group1 和 group2和他们相同，一定不能成为朋友
				if (group1 == group3 && group2 == group4) || (group1 == group4 && group2 == group3){
					valid = false
					break
				}
			}
			if valid{
				res[i] = true
				groups[group1] = group2
			}else{
				res[i] = false
			}
		}
	}
	return res
}