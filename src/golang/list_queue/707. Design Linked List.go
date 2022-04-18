package list_queue

type MyLinkedList struct {
	data []int
}


func Constructor707() MyLinkedList {
	var obj MyLinkedList
	return obj
}


func (this *MyLinkedList) Get(index int) int {
	if index >= len(this.data){
		return -1
	}
	return this.data[index]
}


func (this *MyLinkedList) AddAtHead(val int)  {
	this.data = append([]int{val},this.data...)
}


func (this *MyLinkedList) AddAtTail(val int)  {
	this.data = append(this.data,val)
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index > len(this.data){
		return
	}
	if index == len(this.data){
		this.data = append(this.data,val)
		return
	}
	if index < 0{
		this.data = append([]int{val},this.data...)
		return
	}
	this.data = append(this.data[:index], append([]int{val}, this.data[index:]...)...)
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index >= len(this.data){
		return
	}
	this.data = append(this.data[:index],this.data[index + 1:]...)
}