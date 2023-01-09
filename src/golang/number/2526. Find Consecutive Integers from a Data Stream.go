package number

type DataStream struct {
	value int
	k int
	cnt int
}

func Constructor2526(value int, k int) DataStream {
	var obj DataStream
	obj.value = value
	obj.k = k
	obj.cnt = 0
	return obj
}

func (this *DataStream) Consec(num int) bool {
	if num == this.value{
		if this.cnt < this.k{
			this.cnt++
		}
		return this.cnt == this.k
	}else{
		this.cnt = 0
		return false
	}
}