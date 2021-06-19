package list_queue


//OJ fail
//Out of memory
//type ExamRoom struct {
//	size int
//	used int
//	status []bool
//}
//
//func Constructor855(N int) ExamRoom {
//	var obj ExamRoom
//	obj.size = N
//	obj.used = 0
//	obj.status = make([]bool,N)
//	return obj
//}
//
//func findClosest(seats []bool) int {
//	var l int = len(seats)
//	var left int = 0
//	for left < l && !seats[left]{
//		left++
//	}
//	var idx int = -1
//	var max_distance int = left
//	if left != 0{
//		idx = 0
//	}
//	var right int = l - 1
//	for right >= 0 && !seats[right]{
//		right--
//	}
//	if  (l - right - 1) > max_distance{
//		max_distance =  l - right - 1
//		idx = l - 1
//	}
//	var head int = left
//	for head < right {
//		for head < right && !seats[head]{
//			head++
//		}
//		var tail int = head + 1
//		for tail < right && !seats[tail]{
//			tail++
//		}
//		dis := (tail - head)/2
//		if dis > max_distance{
//			max_distance = dis
//			idx = head + dis
//		}
//		if dis == max_distance{
//			if idx == l - 1{
//				idx = head + dis
//			}
//		}
//		head = tail
//	}
//	return idx
//}
//
//func (this *ExamRoom) Seat() int {
//	if this.used == 0{
//		this.status[0] = true
//		this.used++
//		return 0
//	}
//	idx := findClosest(this.status)
//	this.status[idx] = true
//	this.used++
//	return idx
//}
//
//func (this *ExamRoom) Leave(p int)  {
//	if this.status[p]{
//		this.used--
//		this.status[p] = false
//	}
//}