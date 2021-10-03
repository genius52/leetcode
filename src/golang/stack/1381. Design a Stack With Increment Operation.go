package stack

type CustomStack struct {
	data []int
	max_size int
	cur_size int
}

func Constructor(maxSize int) CustomStack {
	var obj CustomStack
	obj.max_size = maxSize
	return obj
}

func (this *CustomStack) Push(x int)  {
	if this.cur_size < this.max_size{
		this.data = append(this.data,x)
		this.cur_size++
	}
}

func (this *CustomStack) Pop() int {
	if this.cur_size > 0{
		res := this.data[this.cur_size - 1]
		this.cur_size--
		this.data = this.data[:this.cur_size]
		return res
	}
	return -1
}

func (this *CustomStack) Increment(k int, val int)  {
	for i := 0;i < k && i < this.cur_size;i++{
		this.data[i] += val
	}
}