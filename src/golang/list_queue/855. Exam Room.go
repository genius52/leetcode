package list_queue

import "math"

type ExamRoom struct {
	seats []int
	total int
}

func Constructor855(N int) ExamRoom {
	var obj ExamRoom
	obj.total = N
	//obj.seats = make([]int,N)
	return obj
}

func (this *ExamRoom) Seat() int {
	var l int = len(this.seats)
	if l == 0{
		this.seats = append(this.seats,0)
	}else if l == 1{
		this.seats = append(this.seats,this.total - 1)
	}else{
		//var gap int = int(math.Abs(float64(this.seats[0]) - float64(this.seats[1])))
		for i := 2;i < l;i++{
			gap1 := int(math.Abs(float64(this.seats[i]) - float64(this.seats[i - 1])))
			gap2 := int(math.Abs(float64(this.seats[i]) - float64(this.seats[i - 2])))
			if gap1 >= gap2{

			}
		}
	}
	return 0
}

func (this *ExamRoom) Leave(p int)  {
	//this.seats[p] = false
}