package number

import "container/list"

func minimumOperations(nums []int, start int, goal int) int {
	var l int = len(nums)
	var q list.List
	q.PushBack(start)
	var steps int = 0
	var visited map[int]bool = make(map[int]bool)
	for q.Len() > 0{
		var cur_len int = q.Len()
		for i := 0;i < cur_len;i++{
			var cur int = q.Front().Value.(int)
			q.Remove(q.Front())
			if cur == goal{
				return steps
			}
			if cur < 0 || cur > 1000{
				continue
			}
			for j := 0;j < l;j++{
				next1 := cur + nums[j]
				if _,ok := visited[next1];!ok{
					visited[next1] = true
					q.PushBack(next1)
				}
				next2 := cur - nums[j]
				if _,ok := visited[next2];!ok{
					visited[next2] = true
					q.PushBack(next2)
				}
				next3 := cur ^ nums[j]
				if _,ok := visited[next3];!ok{
					visited[next3] = true
					q.PushBack(next3)
				}
			}
		}
		steps++
	}
	return -1
}