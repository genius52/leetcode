package list_queue

import (
	"container/list"
)

func AsteroidCollision(asteroids []int) []int {
	var l int = len(asteroids)
	var q list.List
	for i := 0;i < l;i++{
		if q.Len() == 0{
			q.PushBack(asteroids[i])
		}else{
			var last int = q.Back().Value.(int)
			if asteroids[i] > 0 || (asteroids[i] < 0 && last < 0){
				q.PushBack(asteroids[i])
			}else if asteroids[i] < 0 {
				var insert bool = true
				for q.Len() > 0 && q.Back().Value.(int) > 0{
					if q.Back().Value.(int) < -asteroids[i]{
						q.Remove(q.Back())
					}else if q.Back().Value.(int) == -asteroids[i]{
						q.Remove(q.Back())
						insert = false
						break
					}else{
						insert = false
						break
					}
				}
				if insert{
					q.PushBack(asteroids[i])
				}
			}
		}
	}
	var q_len int = q.Len()
	var res []int = make([]int,q_len)
	var index int = 0
	for i := 0;i < q_len;i++{
		res[index] = q.Front().Value.(int)
		index++
		q.Remove(q.Front())
	}
	return res
}