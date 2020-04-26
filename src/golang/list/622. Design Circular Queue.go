package list

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
//MyCircularQueue circularQueue = new MyCircularQueue(3); // set the size to be 3
//circularQueue.enQueue(1);  // return true
//circularQueue.enQueue(2);  // return true
//circularQueue.enQueue(3);  // return true
//circularQueue.enQueue(4);  // return false, the queue is full
//circularQueue.Rear();  // return 3
//circularQueue.isFull();  // return true
//circularQueue.deQueue();  // return true
//circularQueue.enQueue(4);  // return true
//circularQueue.Rear();  // return 4
type MyCircularQueue struct {
	data []int
	end int
	cur_pos int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	var obj MyCircularQueue
	obj.data = make([]int,k)
	obj.end = k - 1
	obj.cur_pos = -1
	return obj
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.cur_pos >= this.end{
		return false
	}
	this.cur_pos++
	this.data[this.cur_pos] = value
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.cur_pos == -1{
		return false
	}
	tmp := this.data[1:]
	copy(this.data,tmp)
	this.cur_pos--
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.cur_pos == -1{
		return - 1
	}
	return this.data[0]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.cur_pos == -1{
		return - 1
	}
	return this.data[this.cur_pos]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.cur_pos < 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.cur_pos == this.end
}
