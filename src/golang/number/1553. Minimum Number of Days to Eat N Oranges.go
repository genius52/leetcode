package number

import "container/list"

func dp_minDays(n int)int{
	return 0
}

//吃掉一个橘子。
//如果剩余橘子数 n 能被 2 整除，那么你可以吃掉 n/2 个橘子。
//如果剩余橘子数 n 能被 3 整除，那么你可以吃掉 2*(n/3) 个橘子。
func MinDays(n int) int {
	var visited map[int]bool = make(map[int]bool)
	visited[n] = true
	var steps int = 0
	for true{
		var next_visited map[int]bool = make(map[int]bool)
		for val,_ := range visited{
			if val == 0{
				return steps
			}
			if val % 3 == 0{
				next3 := val - 2*(val/3)
				next_visited[next3] = true
			}
			if val | 1 != val{
				next2 := val / 2
				next_visited[next2] = true
			}
			next1 := val - 1
			next_visited[next1] = true
		}
		visited = next_visited
		steps++
	}
	return steps
}

func MinDays2(n int) int {
	var q list.List
	q.PushBack(n)
	var visited map[int]bool = make(map[int]bool)
	var steps int = 0
	for q.Len() > 0{
		var cur_len int = q.Len()
		for i := 0;i < cur_len;i++{
			cur := q.Front().Value.(int)
			q.Remove(q.Front())
			if cur == 0{
				return steps
			}
			if _,ok := visited[cur];ok{
				continue
			}
			visited[cur] = true
			if cur % 3 == 0{
				next3 := cur - 2*(cur/3)
				q.PushBack(next3)
			}
			if cur | 1 != cur{
				next2 := cur / 2
				q.PushBack(next2)
			}
			next1 := cur - 1
			q.PushBack(next1)
		}
		steps++
	}
	return steps
}