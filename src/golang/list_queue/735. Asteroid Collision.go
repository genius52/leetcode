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
			if asteroids[i] > 0{
				q.PushBack(asteroids[i])
			}else{
				for q.Len() > 0 && q.Back().Value.(int) > 0 && q.Back().Value.(int) < -asteroids[i]{
					q.Remove(q.Back())
				}
				if q.Len() == 0 || q.Back().Value.(int) < 0{
					q.PushBack(asteroids[i])
				}else if q.Back().Value.(int) == -asteroids[i]{
					q.Remove(q.Back())
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