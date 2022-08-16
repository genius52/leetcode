package string_issue

import (
	"strconv"
)

//输入：pattern = "IIIDIDDD"
//输出："123549876"
func dfs_smallestNumber(pattern string,l int,idx int,data *[]int,used [10]bool)(bool,[]int){
	if idx >= l{
		return true,*data
	}
	var next int = 0
	var limit int = 0
	if pattern[idx] == 'I'{
		next = (*data)[len(*data) - 1] + 1
		limit = 9
	}else{
		next = 1
		limit = (*data)[len(*data) - 1] - 1
	}
	for ;next <= limit;next++{
		if used[next]{
			continue
		}
		used[next] = true
		*data = append(*data,next)
		ret,r := dfs_smallestNumber(pattern,l,idx + 1,data,used)
		used[next] = false
		if !ret{
			*data = (*data)[:len(*data) - 1]
		}else{
			return ret,r
		}
	}
	return false,*data
}

func smallestNumber(pattern string) string {
	var l int = len(pattern)
	var res string
	for i := 1;i <= 9;i++{
		var data []int = []int{i}
		var used [10]bool
		used[i] = true
		ret,data := dfs_smallestNumber(pattern,l,0,&data,used)
		if ret{
			for j := 0;j < len(data);j++{
				res += strconv.Itoa(data[j])
			}
			break
		}
	}
	return res
}