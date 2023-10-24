package number

import "sort"

func MinProcessingTime(processorTime []int, tasks []int) int {
	var n int = len(processorTime)
	sort.Ints(processorTime)
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i] > tasks[j]
	})
	var res int = 0
	//最早空闲的处理最耗时的任务
	for i := 0; i < n; i++ {
		//for j := i * 4; j <= i*4+3; j++ {
		res = max_int(res, processorTime[i]+tasks[i*4+3])
		//}
	}
	return res
}
