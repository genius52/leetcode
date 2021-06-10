package number

import (
	"container/list"
	"math"
)

//当车得到指令 "A" 时, 将会做出以下操作： position += speed, speed *= 2。
//
//当车得到指令 "R" 时, 将会做出以下操作：如果当前速度是正数，则将车速调整为 speed = -1 ；
//否则将车速调整为 speed = 1。  (当前所处位置不变。)
type Speed_dir struct{
	position int
	speed int
	forward bool
}

func racecar(target int) int {
	var q list.List
	var start Speed_dir
	start.position = 0
	start.speed = 1
	start.forward = true
	q.PushBack(start)
	var visited map[int]map[int]bool = make(map[int]map[int]bool)//position -> speed
	visited[0] = make(map[int]bool)
	visited[0][1] = true
	var steps int = 0
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var cur Speed_dir = q.Front().Value.(Speed_dir)
			q.Remove(q.Front())
			if cur.position == target{
				return steps
			}
			var A Speed_dir
			A.position = cur.position + cur.speed
			A.speed = cur.speed * 2
			if int(math.Abs(float64(A.position - target))) < target{
				if _,ok1 := visited[A.position];ok1{
					if _,ok2 := visited[A.position][A.speed];!ok2{
						q.PushBack(A)
						visited[A.position][A.speed] = true
					}
				}else{
					q.PushBack(A)
					visited[A.position] = make(map[int]bool)
					visited[A.position][A.speed] = true
				}
			}

			var R Speed_dir
			R.position = cur.position
			if cur.speed > 0{
				R.speed = -1
			}else{
				R.speed = 1
			}
			if int(math.Abs(float64(R.position - target))) < target {
				if _, ok1 := visited[R.position]; ok1 {
					if _, ok2 := visited[R.position][R.speed]; !ok2 {
						q.PushBack(R)
						visited[R.position][R.speed] = true
					}
				} else {
					q.PushBack(R)
					visited[R.position][R.speed] = true
				}
			}
		}
		steps++
	}
	return steps
}