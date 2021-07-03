package diagram

//Input
//["RLEIterator", "next", "next", "next", "next"]
//[[[3, 8, 0, 9, 2, 5]], [2], [1], [1], [2]]
//Output
//[null, 8, 8, 5, -1]
type RLEIterator struct {
	data []int
	cur int
	length int
}

func Constructor900(encoding []int) RLEIterator {
	var l int = len(encoding)
	var obj RLEIterator
	obj.data = make([]int,l)
	copy(obj.data,encoding)
	obj.length = l
	return obj
}

func (this *RLEIterator) Next(n int) int {
	for this.cur < this.length{
		if this.data[this.cur] >= n{
			this.data[this.cur] -= n
			var res int = this.data[this.cur + 1]
			if this.data[this.cur] == 0{
				this.cur += 2
			}
			return res
		}else{
			n -= this.data[this.cur]
			this.cur += 2
		}
	}
	return -1
}

//type RLEIterator struct {
//	data []int
//	count []int
//	total int
//	used int
//	pointer int
//}
//
//func Constructor900(A []int) RLEIterator {
//	var obj RLEIterator
//	var l int = len(A)
//	for i := 0;i < l;i++{
//		if i | 1 == i{
//			obj.data = append(obj.data,A[i])
//		}else{
//			obj.total += A[i]
//			obj.count = append(obj.count,A[i])
//		}
//	}
//	return obj
//}
//
//func (this *RLEIterator) Next(n int) int {
//	if this.used + n >= this.total{
//		this.used = this.total
//		return -1
//	}
//	this.used += n
//	var res int = -1
//	for n > 0{
//		res = this.data[this.pointer]
//		if this.count[this.pointer] < n{
//			n -= this.count[this.pointer]
//			this.count[this.pointer] = 0
//			this.pointer++
//		}else{
//			this.count[this.pointer] -= n
//			n = 0
//		}
//		for this.pointer < len(this.count) && this.count[this.pointer] == 0{
//			this.pointer++
//		}
//	}
//	return res
//}

