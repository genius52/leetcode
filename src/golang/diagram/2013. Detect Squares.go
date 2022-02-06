package diagram

type DetectSquares struct {
	cnt [1001][1001]int
	points [][]int
}


func Constructor2013() DetectSquares {
	var obj DetectSquares
	//obj.record = make(map[int]int)
	return obj
}


func (this *DetectSquares) Add(point []int)  {
	this.cnt[point[0]][point[1]]++
	this.points = append(this.points,point)
}


func (this *DetectSquares) Count(point []int) int {
	var res int = 0
	x := point[0]
	y := point[1]
	for _,p := range this.points{
		if p[0] != x && p[1] != y && abs_int(x - p[0]) == abs_int(y - p[1]){
			res += this.cnt[x][p[1]] * this.cnt[p[0]][y]
		}
	}
	return res
}