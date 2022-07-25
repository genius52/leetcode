package number

type NumberContainers struct {
	idx_num map[int]int
	num_idx map[int]map[int]bool
}

func Constructor2349() NumberContainers {
	var obj NumberContainers
	obj.idx_num = make(map[int]int)
	obj.num_idx = make(map[int]map[int]bool)
	return obj
}

func (this *NumberContainers) Change(index int, number int)  {
	if pre_num,ok := this.idx_num[index];ok{
		delete(this.num_idx[pre_num],index)
		this.idx_num[index] = number
		if this.num_idx[number] == nil{
			this.num_idx[number] = make(map[int]bool)
		}
		this.num_idx[number][index] = true
	}else{
		this.idx_num[index] = number
		if this.num_idx[number] == nil{
			this.num_idx[number] = make(map[int]bool)
		}
		this.num_idx[number][index] = true
	}
}

func (this *NumberContainers) Find(number int) int {
	if _,ok := this.num_idx[number];ok{
		var res int = 2147483647
		for idx,_ := range this.num_idx[number]{
			if idx < res{
				res = idx
			}
		}
		if res == 2147483647{
			return -1
		}
		return res
	}else{
		return -1
	}
}