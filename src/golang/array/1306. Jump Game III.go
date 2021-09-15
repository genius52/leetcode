package array

import "container/list"

//Input: arr = [4,2,3,0,3,1,2], start = 5
//Output: true
//All possible ways to reach at index 3 with value 0 are:
//index 5 -> index 4 -> index 1 -> index 3
//index 5 -> index 6 -> index 4 -> index 1 -> index 3
func dfs_canReach(arr []int,l int,pos int,visited *[]bool)bool{
	if pos < 0 || pos >= l{
		return false
	}
	if (*visited)[pos]{
		return false
	}
	if arr[pos] == 0{
		return true
	}
	(*visited)[pos] = true
	forward := dfs_canReach(arr,l,pos + arr[pos],visited)
	backward := dfs_canReach(arr,l,pos - arr[pos],visited)
	return forward || backward
}

func CanReach(arr []int, start int) bool {
	var l int = len(arr)
	var visited []bool = make([]bool,l)
	return dfs_canReach(arr,l,start,&visited)
}

func canReach(arr []int, start int) bool{
	var l int = len(arr)
	var q list.List
	q.PushBack(start)
	var visited []bool = make([]bool,l)
	visited[start] = true
	for q.Len() > 0{
		var cur_len int = q.Len()
		for i := 0;i < cur_len;i++{
			pos := q.Front().Value.(int)
			if arr[pos] == 0{
				return true
			}
			q.Remove(q.Front())
			forward := pos + arr[pos]
			if forward >= 0 && forward < l{
				if !visited[forward]{
					visited[forward] = true
					q.PushBack(forward)
				}
			}
			backward := pos - arr[pos]
			if backward >= 0 && backward < l{
				if !visited[backward]{
					visited[backward] = true
					q.PushBack(backward)
				}
			}
		}
	}
	return false
}