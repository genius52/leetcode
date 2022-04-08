package list_queue

const size int = 100

type MyHashMap struct {
	data [][][2]int
}

func Constructor706() MyHashMap {
	var obj MyHashMap
	obj.data = make([][][2]int,size)
	return obj
}

func (this *MyHashMap) Put(key int, value int)  {
	idx := key % size
	for i := 0;i < len(this.data[idx]);i++{
		if this.data[idx][i][0] == key{
			this.data[idx][i][1] = value
			return
		}
	}
	this.data[idx] = append(this.data[idx],[2]int{key,value})
}

func (this *MyHashMap) Get(key int) int {
	idx := key % size
	for i := 0;i < len(this.data[idx]);i++{
		if this.data[idx][i][0] == key{
			return this.data[idx][i][1]
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int)  {
	idx := key % size
	for i := 0;i < len(this.data[idx]);i++{
		if this.data[idx][i][0] == key{
			this.data[idx] = append(this.data[idx][:i],this.data[idx][i + 1:]...)
			break
		}
	}
}