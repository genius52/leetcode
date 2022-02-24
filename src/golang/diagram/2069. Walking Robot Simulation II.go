package diagram

type Robot struct {
	columns int
	rows int
	r int
	c int
	dir int
}

func Constructor2069(width int, height int) Robot {
	var obj Robot
	obj.rows = height
	obj.columns = width
	return obj
}

func (this *Robot) Step(num int)  {
	var dirs [][]int = [][]int{{0,1},{1,0},{0,-1},{-1,0}}
	if num > (this.rows + this.columns) * 2{
		num = num % ((this.rows + this.columns - 2) * 2)
		if num == 0{
			num += (this.rows + this.columns - 2) * 2
		}
	}
	for num > 0{
		this.r = this.r + dirs[this.dir][0]
		this.c = this.c + dirs[this.dir][1]
		if this.r < 0{
			this.r = 0
			this.dir = (this.dir + 1) % 4
		}else if this.r >= this.rows{
			this.r = this.rows - 1
			this.dir = (this.dir + 1) % 4
		}else if this.c < 0{
			this.c = 0
			this.dir = (this.dir + 1) % 4
		}else if this.c >= this.columns{
			this.c = this.columns - 1
			this.dir = (this.dir + 1) % 4
		}else{
			num--
		}
	}
}

func (this *Robot) GetPos() []int {
	return []int{this.c,this.r}
}

func (this *Robot) GetDir() string {
	var res string
	switch this.dir {
	case 0:
		res = "East"
	case 1:
		res = "North"
	case 2:
		res = "West"
	case 3:
		res = "South"
	}
	return res
}