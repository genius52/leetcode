package array

import "sort"

//输入：dist = [3,2,4], speed = [5,3,2]
//输出：1
//解释：
//第 0 分钟开始时，怪物的距离是 [3,2,4]，你消灭了第一个怪物。
//第 1 分钟开始时，怪物的距离是 [X,0,2]，你输掉了游戏。
//你只能消灭 1 个怪物。
func EliminateMaximum(dist []int, speed []int) int {
	var l int = len(dist)
	var record []float32 = make([]float32, l)
	for i := 0; i < l; i++ {
		record[i] = float32(dist[i]) / float32(speed[i])
	}
	sort.Slice(record, func(i, j int) bool {
		return record[i] < record[j]
	})
	var res int = 0
	var time int = 0
	for i := 0; i < l; i++ {
		if record[i] <= float32(time) {
			return res
		}
		time++
		res++
	}
	return res
}
