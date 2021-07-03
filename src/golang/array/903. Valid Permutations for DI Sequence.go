package array

//Input: s = "DID"
//Output: 5
//Explanation:
//The 5 valid permutations of (0, 1, 2, 3) are:
//(1, 0, 3, 2)
//(2, 0, 3, 1)
//(2, 1, 3, 0)
//(3, 0, 2, 1)
//(3, 1, 2, 0)

// dp from top to bottom (TLE)
func dfs_numPermsDISequence(s string,l int,cur_pos int,data []int,visited []bool)int{
	if cur_pos >= l{
		for i := 0;i <= l;i++{
			if !visited[i]{
				if s[cur_pos - 1] == 'D'{
					if data[cur_pos - 1] < i{
						return 0
					}else{
						return 1
					}
				}else{
					if data[cur_pos - 1] < i{
						return 1
					}else{
						return 0
					}
				}
			}
		}
	}
	var res int = 0
	if s[cur_pos] == 'D'{
		for i := l;i >= 1;i--{
			if visited[i]{
				continue
			}
			if cur_pos == 0{
				visited[i] = true
				data[cur_pos] = i
				res += dfs_numPermsDISequence(s,l,cur_pos + 1,data,visited)
				visited[i] = false
			}else{
				if s[cur_pos - 1] == 'D'{
					if i > data[cur_pos - 1]{
						continue
					}
				}else{
					if i < data[cur_pos - 1]{
						continue
					}
				}
				visited[i] = true
				data[cur_pos] = i
				res += dfs_numPermsDISequence(s,l,cur_pos + 1,data,visited)
				visited[i] = false
			}
		}
	}else if s[cur_pos] == 'I'{
		for i := 0;i < l;i++{
			if visited[i]{
				continue
			}
			if cur_pos == 0{
				visited[i] = true
				data[cur_pos] = i
				res += dfs_numPermsDISequence(s,l,cur_pos + 1,data,visited)
				visited[i] = false
			}else{
				if s[cur_pos - 1] == 'D'{
					if i > data[cur_pos - 1]{
						continue
					}
				}else{
					if i < data[cur_pos - 1]{
						continue
					}
				}
				visited[i] = true
				data[cur_pos] = i
				res += dfs_numPermsDISequence(s,l,cur_pos + 1,data,visited)
				visited[i] = false
			}
		}
	}
	return res
}

func NumPermsDISequence(s string) int {
	var l int = len(s)
	var data []int = make([]int,l + 1)
	var visited []bool = make([]bool,l + 1)
	var res int = dfs_numPermsDISequence(s,l,0,data,visited)
	return res
}