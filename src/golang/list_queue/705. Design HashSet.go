package list_queue

type MyHashSet struct {
	data [][]int
}

func Constructor705() MyHashSet {
	var obj MyHashSet
	obj.data = make([][]int,1000)
	return obj
}

func (this *MyHashSet) Add(key int)  {
	idx := key % 1000
	var find_idx int = -1
	for i,k := range this.data[idx]{
		if k == key{
			find_idx = i
			break
		}
	}
	if find_idx == -1{
		this.data[idx] = append(this.data[idx],key)
	}
}

func (this *MyHashSet) Remove(key int)  {
	idx := key % 1000
	var find_idx int = -1
	for i,k := range this.data[idx]{
		if k == key{
			find_idx = i
			break
		}
	}
	if find_idx != -1{
		this.data[idx] = append(this.data[idx][:find_idx],this.data[idx][find_idx + 1:]...)
	}
}

func (this *MyHashSet) Contains(key int) bool {
	idx := key % 1000
	for _,k := range this.data[idx]{
		if k == key{
			return true
		}
	}
	return false
}
