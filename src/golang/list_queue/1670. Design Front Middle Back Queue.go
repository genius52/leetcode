package list_queue

type FrontMiddleBackQueue struct {
	data []int
}

func Constructor1670() FrontMiddleBackQueue {
	var obj FrontMiddleBackQueue
	return obj
}

func (this *FrontMiddleBackQueue) PushFront(val int)  {
	this.data = append([]int{val},this.data...)
}

func (this *FrontMiddleBackQueue) PushMiddle(val int)  {
	var l int = len(this.data)
	var pos int = l/2
	rear := append([]int{},this.data[pos:]...)
	this.data = append(this.data[:pos],val)
	this.data = append(this.data,rear...)

	// rear:=append([]string{},ss[index:]...)
	//    ss=append(ss[0:index],"inserted")
	//    ss=append(ss,rear...)
}

func (this *FrontMiddleBackQueue) PushBack(val int)  {
	this.data = append(this.data,val)
}

func (this *FrontMiddleBackQueue) PopFront() int {
	var l int = len(this.data)
	if l == 0{
		return -1
	}
	res := this.data[0]
	this.data = append(this.data[:0],this.data[1:]...)
	return res
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	var l int = len(this.data)
	if l == 0{
		return -1
	}
	var pos int = (l + 1)/2 - 1
	ret := this.data[pos]
	this.data = append(this.data[:pos],this.data[pos + 1:]...)
	return ret
}

func (this *FrontMiddleBackQueue) PopBack() int {
	var l int = len(this.data)
	if l == 0{
		return -1
	}
	res := this.data[l - 1]
	this.data = append(this.data[:l - 1],this.data[l:]...)
	return res
}