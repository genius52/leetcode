package string_issue

//"10101"
func dfs_numWays(s string,l int,pos int,cur_one_cnt int,one_target int,divide_cnt int)int{
	if pos >= l{
		return 0
	}
	if divide_cnt == 2{
		return 1
	}

	if s[pos] == '1'{
		cur_one_cnt++
		if cur_one_cnt > one_target{
			return 0
		}
		if cur_one_cnt == one_target{
			stop_here := dfs_numWays(s,l,pos + 1,0,one_target,divide_cnt + 1)
			visit_next := dfs_numWays(s,l,pos + 1,cur_one_cnt,one_target,divide_cnt)
			return stop_here + visit_next
		}else{
			return dfs_numWays(s,l,pos + 1,cur_one_cnt + 1,one_target,divide_cnt)
		}
	}else{
		if cur_one_cnt == one_target{
			stop_here := dfs_numWays(s,l,pos + 1,0,one_target,divide_cnt + 1)
			visit_next := dfs_numWays(s,l,pos + 1,cur_one_cnt,one_target,divide_cnt)
			return stop_here + visit_next
		}else{
			return dfs_numWays(s,l,pos + 1,cur_one_cnt,one_target,divide_cnt)
		}
	}
}

func NumWays(s string) int {
	var one_cnt int = 0
	var l int = len(s)
	for i := 0;i < l;i++{
		if s[i] == '1'{
			one_cnt++
		}
	}
	if one_cnt % 3 != 0{
		return 0
	}
	var target int = one_cnt / 3
	return dfs_numWays(s,l,0,0,target,0)
}

func mathZuhe(n int, m int) int {
	divide := (jieCheng(n-m) * jieCheng(m)) % 1000000007
	total := jieCheng(n)  /divide

	total %= 1000000007
	return total
}

//阶乘
func jieCheng(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
		result %= 1000000007
	}
	return result
}

func NumWays2(s string) int {
	var one_cnt int = 0
	var l int = len(s)
	for i := 0;i < l;i++{
		if s[i] == '1'{
			one_cnt++
		}
	}
	if one_cnt % 3 != 0{
		return 0
	}
	if one_cnt == 0{
		return (l-2) * (l-1)/2 % 1000000007 //mid section has 1 + 2 + 3... + n-2 options
	}
	var target int = one_cnt / 3
	var res int = 0
	var begin int = 0
	var end int = l - 1
	var begin_one_cnt int = 0
	var end_one_cnt int = 0
	for begin_one_cnt < target {
		if s[begin] == '1'{
			begin_one_cnt++
		}
		begin++
	}
	begin--
	for end_one_cnt < target{
		if s[end] == '1'{
			end_one_cnt++
		}
		end--
	}
	end++
	var move_begin int = begin + 1
	for s[move_begin] != '1'{
		move_begin++
	}
	move_begin--
	var move_end int = end - 1
	for s[move_end] != '1'{
		move_end--
	}
	move_end++
	res += (move_begin - begin + 1) * (end - move_end + 1) % 1000000007
	return res
}