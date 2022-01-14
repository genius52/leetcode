package number

import (
	"container/list"
	"math"
)

//Input: x = 5, target = 501
//Output: 8
//Explanation: 5 * 5 * 5 * 5 - 5 * 5 * 5 + 5 / 5.
//The expression contains 8 operations.
func LeastOpsExpressTarget(x int, target int) int {
	var visited map[float64]bool = make(map[float64]bool)
	var q list.List
	q.PushBack(float64(x))
	var steps int = 0
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var n float64 = q.Front().Value.(float64)
			q.Remove(q.Front())
			if math.Abs(n - float64(target)) <= 1e-6{
				return steps
			}
			add := n + float64(x)
			if _,ok := visited[add];!ok{
				q.PushBack(add)
				visited[add] = true
			}
			minus := math.Abs(n - float64(x))
			if _,ok := visited[minus];!ok{
				q.PushBack(minus)
				visited[minus] = true
			}
			times := n * float64(x)
			if _,ok := visited[times];!ok{
				q.PushBack(times)
				visited[times] = true
			}
			divide1 :=  n / float64(x)
			if _,ok := visited[divide1];!ok{
				q.PushBack(divide1)
				visited[divide1] = true
			}
			if n != 0{
				divide2 :=  float64(x)/n
				if _,ok := visited[divide2];!ok{
					q.PushBack(divide2)
					visited[divide2] = true
				}
			}

		}
		steps++
	}
	return steps
}