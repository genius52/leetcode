package array

import (
	"strconv"
)

//Input: tasks = [1,2,3], sessionTime = 3
//Output: 2
func dfs_minSessions(tasks *[]int,l int,pos int,choose *[]int,sessionTime int,memo map[string]int,min_len *int)int{
	cur_len := len(*choose)
	if pos >= l{
		return cur_len
	}
	if cur_len >= *min_len{
		return cur_len
	}
	var keys string
	for i := 0;i < cur_len;i++{
		keys += "," + strconv.Itoa((*choose)[i])
	}
	keys += "|" + strconv.Itoa(pos)
	if val,ok := memo[keys];ok{
		return val
	}
	if cur_len == 0{
		*choose = append(*choose,(*tasks)[pos])
		return dfs_minSessions(tasks,l,pos + 1,choose,sessionTime,memo,min_len)
	}
	//add to new session
	*choose = append(*choose,(*tasks)[pos])
	var res int = dfs_minSessions(tasks,l,pos + 1,choose,sessionTime,memo,min_len)
	*choose = (*choose)[:len(*choose) - 1]
	for i := 0;i < cur_len;i++{
		if (*choose)[i] + (*tasks)[pos] > sessionTime{
			continue
		}
		(*choose)[i] += (*tasks)[pos]
		cur := dfs_minSessions(tasks,l,pos + 1,choose,sessionTime,memo,min_len)
		if cur < res{
			res = cur
		}
		(*choose)[i] -= (*tasks)[pos]
	}
	if res < *min_len{
		*min_len = res
	}
	memo[keys] = res
	return res
}

func MinSessions(tasks []int, sessionTime int) int {
	var choose []int
	var memo map[string]int = make(map[string]int)
	var min_len int = len(tasks)
	return dfs_minSessions(&tasks,len(tasks),0,&choose,sessionTime,memo,&min_len)
}