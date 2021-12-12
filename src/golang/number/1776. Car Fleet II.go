package number

import "container/list"

//输入：cars = [[1,2],[2,1],[4,3],[7,2]]
//输出：[1.00000,-1.00000,3.00000,-1.00000]
//解释：经过恰好 1 秒以后，第一辆车会与第二辆车相遇，并形成一个 1 m/s 的车队。经过恰好 3 秒以后，
//第三辆车会与第四辆车相遇，并形成一个 2 m/s 的车队。

func GetCollisionTimes(cars [][]int) []float64 {
	var l int = len(cars)
	var res []float64 = make([]float64, l)
	var q list.List
	for i := l - 1; i >= 0; i-- {
		res[i] = -1
		for q.Len() > 0 {
			last_idx := q.Back().Value.(int)
			if cars[i][1] > cars[last_idx][1] {
				//这辆车撞车前，我能不能撞上他？
				cur_time := float64(cars[last_idx][0]-cars[i][0]) / float64(cars[i][1]-cars[last_idx][1])
				if res[last_idx] == -1 || cur_time < res[last_idx] {
					res[i] = cur_time
					break
				} else {
					q.Remove(q.Back()) //撞他之前的车
				}
			} else {
				q.Remove(q.Back())
			}
		}
		q.PushBack(i)
	}
	return res
}
