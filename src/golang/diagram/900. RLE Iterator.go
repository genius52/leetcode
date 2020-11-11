package diagram

type RLEIterator struct {
	data []int
	count []int
	total int
	used int
	pointer int
}

func Constructor900(A []int) RLEIterator {
	var obj RLEIterator
	var l int = len(A)
	for i := 0;i < l;i++{
		if i | 1 == i{
			obj.data = append(obj.data,A[i])
		}else{
			obj.total += A[i]
			obj.count = append(obj.count,A[i])
		}
	}
	return obj
}

func (this *RLEIterator) Next(n int) int {
	if this.used + n >= this.total{
		this.used = this.total
		return -1
	}
	this.used += n
	var res int = -1
	for n > 0{
		res = this.data[this.pointer]
		if this.count[this.pointer] < n{
			n -= this.count[this.pointer]
			this.count[this.pointer] = 0
			this.pointer++
		}else{
			this.count[this.pointer] -= n
			n = 0
		}
		for this.pointer < len(this.count) && this.count[this.pointer] == 0{
			this.pointer++
		}
	}
	return res
}

