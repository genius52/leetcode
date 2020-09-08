package list_queue

import "container/list"

type MyCircularDeque struct {
	data *list.List
	capacity int
	cur_size int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor641(k int) MyCircularDeque {
	var obj MyCircularDeque
	obj.data = list.New()
	obj.capacity = k
	obj.cur_size = 0
	return obj
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.cur_size >= this.capacity{
		return false
	}
	this.data.PushFront(value)
	this.cur_size++
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.cur_size >= this.capacity{
		return false
	}
	this.data.PushBack(value)
	this.cur_size++
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.cur_size == 0{
		return false
	}
	front := this.data.Front()
	this.data.Remove(front)
	this.cur_size--
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.cur_size == 0{
		return false
	}
	tail := this.data.Back()
	this.data.Remove(tail)
	this.cur_size--
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.cur_size == 0{
		return -1
	}
	ele := this.data.Front().Value.(int)
	return ele
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.cur_size == 0{
		return -1
	}
	ele := this.data.Back().Value.(int)
	return ele
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.cur_size == 0
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.cur_size == this.capacity
}

