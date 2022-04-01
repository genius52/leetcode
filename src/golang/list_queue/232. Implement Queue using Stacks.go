package list_queue

type MyQueue struct {
	data []int
}

func Constructor232() MyQueue {
	var obj MyQueue
	return obj
}

func (this *MyQueue) Push(x int)  {
	this.data = append(this.data,x)
}

func (this *MyQueue) Pop() int {
	var l int = len(this.data)
	res := this.data[0]
	this.data = this.data[1:l]
	return res
}

func (this *MyQueue) Peek() int {
	return this.data[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.data) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */