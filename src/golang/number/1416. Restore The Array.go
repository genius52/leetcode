package number

import "strconv"

//输入：s = "1317", k = 2000
//输出：8
//解释：可行的数组方案为 [1317]，[131,7]，[13,17]，[1,317]，[13,1,7]，[1,31,7]，[1,3,17]，[1,3,1,7]
func bigger_than_limit(s1 string,limit string)bool{
	var l1 int = len(s1)
	var l2 int = len(limit)
	if l1 < l2{
		return false
	}
	if l1 > l2{
		return true
	}
	for i := 0;i < l1;i++{
		if s1[i] < limit[i]{
			return false
		}else if s1[i] > limit[i]{
			return true
		}
	}
	return false
}

func dfs_numberOfArrays(s string,l int,pos int,limit string,memo *[]int)int{
	if pos >= l{
		return 1
	}
	if s[pos] == '0'{
		return 0
	}
	if (*memo)[pos] != -1{
		return (*memo)[pos]
	}
	var res int = 0
	for i := pos + 1;i <= l;i++{//"i" is the postion of end
		if bigger_than_limit(s[pos:i],limit){
			break
		}
		res += dfs_numberOfArrays(s,l,i,limit,memo)
		res = res % (1e9 + 7)
	}
	(*memo)[pos] = res
	return res
}

func NumberOfArrays(s string, k int) int {
	limit := strconv.Itoa(k)
	var l int = len(s)
	var memo []int = make([]int,l)
	for i := 0;i < l;i++{
		memo[i] = -1
	}
	return dfs_numberOfArrays(s,l,0,limit,&memo)
}