package list_queue


//146
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
type DataNode struct{
	key int
	value int
}

type LRUCache struct {
	data []DataNode
	capacity int
	cur_size int
}

func Constructor146(capacity int) LRUCache {
	var obj LRUCache
	obj.data = make([]DataNode,capacity)
	obj.capacity = capacity
	obj.cur_size = 0
	return obj
}

func (this *LRUCache) Get(key int) int {
	for i := this.cur_size - 1; i >= 0;i--{
		if this.data[i].key == key{
			tmp := this.data[i]
			for j := i;j < this.cur_size - 1;j++{
				this.data[j] = this.data[j + 1]
			}
			this.data[this.cur_size - 1] = tmp
			return tmp.value
		}
	}
	return -1
}

func (this *LRUCache) Put(key int, value int)  {
	for i := this.cur_size - 1; i >= 0;i--{
		if this.data[i].key == key{
			this.data[i].value = value
			var tmp DataNode
			tmp.key = this.data[i].key
			tmp.value = value
			for j := i;j < this.cur_size - 1;j++{
				this.data[j] = this.data[j + 1]
			}
			this.data[this.cur_size - 1] = tmp
			return
		}
	}

	if this.cur_size == this.capacity{
		this.data = append(this.data[:0],this.data[1:]...)
		this.data = append(this.data,DataNode{key,value})
	}else{
		this.cur_size++
		this.data[this.cur_size - 1] = DataNode{key,value}
	}
}