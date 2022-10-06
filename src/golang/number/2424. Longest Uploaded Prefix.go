package number

type LUPrefix struct {
	used []bool
	cur int
	limit int
}

func Constructor2424(n int) LUPrefix {
	var obj LUPrefix
	obj.used = make([]bool,n + 1)
	obj.limit = n
	return obj
}


func (this *LUPrefix) Upload(video int)  {
	this.used[video] = true
	for this.cur < this.limit && this.used[this.cur + 1]{
		this.cur++
	}
}

func (this *LUPrefix) Longest() int {
	return this.cur
}