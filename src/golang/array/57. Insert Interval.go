package array

func Insert(intervals [][]int, newInterval []int) [][]int {
	var l int = len(intervals)
	if l == 0{
		return [][]int{newInterval}
	}
	var res [][]int
	var visit int = 0
	for visit < l{
		if intervals[visit][1] < newInterval[0] {
			res = append(res,intervals[visit])
			visit++
		}else{
			break
		}
	}
	if visit == l{
		res = append(res,newInterval)
		return res
	}
	visit2 := visit
	for visit2 < l{
		if newInterval[1] < intervals[visit2][0]{
			break
		}
		visit2++
	}
	var new_begin int = 0
	if intervals[visit][0] < newInterval[0]{
		new_begin = intervals[visit][0]
	}else{
		new_begin = newInterval[0]
	}
	var new_end int = 0
	if visit2 > 0{
		if intervals[visit2 - 1][1] > newInterval[1]{
			new_end = intervals[visit2 - 1][1]
		}else{
			new_end = newInterval[1]
		}
	}else{
		new_end = newInterval[1]
	}

	res = append(res,[]int{new_begin,new_end})
	for visit2 < l{
		res = append(res,intervals[visit2])
		visit2++
	}
	return res
}

func insert(intervals [][]int, newInterval []int) [][]int{
	var l int = len(intervals)
	var res [][]int
	var i int = 0
	for i < l{
		if newInterval[0] <= intervals[i][1]{
			break
		}
		res = append(res,intervals[i])
		i++
	}
	if i == l{
		res = append(res,newInterval)
		return res
	}
	var left int = min_int(newInterval[0],intervals[i][0])
	var right int = newInterval[1]
	for i < l && newInterval[1] >= intervals[i][0]{
		right = max_int(right,intervals[i][1])
		i++
	}
	res = append(res,[]int{left,right})
	for i < l {
		res = append(res,intervals[i])
		i++
	}
	return res
}