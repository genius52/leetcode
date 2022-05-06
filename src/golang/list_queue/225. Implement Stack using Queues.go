package list_queue

import "container/list"

type MyStack struct {
	q1 list.List
	q2 list.List
}

func Constructor225() MyStack {
	var obj MyStack
	return obj
}

func (this *MyStack) Push(x int)  {
	this.q1.PushBack(x)
}

func (this *MyStack) Pop() int {
	for this.q1.Len() > 1{
		this.q2.PushBack(this.q1.Front().Value.(int))
		this.q1.Remove(this.q1.Front())
	}
	res := this.q1.Front().Value.(int)
	this.q1.Remove(this.q1.Front())
	for this.q2.Len() > 0{
		this.q1.PushBack(this.q2.Front().Value.(int))
		this.q2.Remove(this.q2.Front())
	}
	return res
}

func (this *MyStack) Top() int {
	for this.q1.Len() > 1{
		this.q2.PushBack(this.q1.Front().Value.(int))
		this.q1.Remove(this.q1.Front())
	}
	res := this.q1.Front().Value.(int)
	for this.q2.Len() > 0{
		this.q1.PushBack(this.q2.Front().Value.(int))
		this.q2.Remove(this.q2.Front())
	}
	this.q1.PushBack(res)
	this.q1.Remove(this.q1.Front())
	return res
}

func (this *MyStack) Empty() bool {
	return this.q1.Len() == 0
}