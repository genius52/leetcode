package string_issue

//往 s1 的字母集合中添加一个字母。
//从 s1 的字母集合中删去一个字母。
//将 s1 中的一个字母替换成另外任意一个字母（也可以替换为这个字母本身）。
func get_parent(groups []int,node int)int{
	if groups[node] != node{
		groups[node] = get_parent(groups,groups[node])
	}
	return groups[node]
}

func GroupStrings(words []string) []int {
	var l int = len(words)
	var groups []int = make([]int,l)
	for i := 0;i < l;i++{
		groups[i] = i
	}
	var nums []int = make([]int,l)//value,idx
	var record map[int]int = make(map[int]int)
	for i := 0;i < l;i++ {
		var n int = 0
		for _, c := range words[i] {
			n |= 1 << (c - 'a')
		}
		nums[i] = n
		record[n] = i
	}
	for i := 0;i < l;i++{
		var n int = nums[i]
		for j := 0;j < 26;j++{
			if (n | (1 << j)) == n{
				//从 s1 的字母集合中删去一个字母。
				var del int = n ^ (1 << j)
				if idx2,ok2 := record[del];ok2{
					var pre_group2 int = get_parent(groups,idx2)
					var cur_group int = get_parent(groups,i)
					if cur_group != pre_group2{
						if cur_group < pre_group2{//组号小的优先
							groups[pre_group2] = cur_group
						}else{
							groups[cur_group] = pre_group2
						}
					}
				}
				//先去掉，再补一个
				//var n2 int = n ^ (1 << j)
				for k := 0;k < 26;k++{
					var replace int = del | (1 << k)
					if replace != del{
						if idx2,ok2 := record[replace];ok2{
							var pre_group2 int = get_parent(groups,idx2)
							var cur_group int = get_parent(groups,i)
							if cur_group != pre_group2{
								if cur_group < pre_group2{//组号小的优先
									groups[pre_group2] = cur_group
								}else{
									groups[cur_group] = pre_group2
								}
							}
						}
					}
				}
			}else{
				//往 s1 的字母集合中添加一个字母。
				var add int = n | (1 << j)
				if idx3,ok3 := record[add];ok3{
					var pre_group3 int = get_parent(groups,idx3)
					var cur_group int = get_parent(groups,i)
					if cur_group != pre_group3{
						if cur_group < pre_group3{//组号小的优先
							groups[pre_group3] = cur_group
						}else{
							groups[cur_group] = pre_group3
						}
					}
				}
			}
		}
	}

	var res []int = make([]int,2)
	var cnt []int = make([]int,l)
	for i := 0;i < l;i++{
		cnt[get_parent(groups,i)]++
	}
	var max_cnt int = 0
	for i := 0;i < l;i++{
		if cnt[i] != 0{
			res[0]++
		}
		if cnt[i] > max_cnt{
			max_cnt = cnt[i]
		}
	}
	res[1] = max_cnt
	return res
}