package array

type SubrectangleQueries struct {
	rect_ [][]int
	history_ [][5]int
}


func Constructor_1476(rectangle [][]int) SubrectangleQueries {
	var obj SubrectangleQueries
	obj.rect_ = rectangle
	return obj
}


func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {
	this.history_ = append(this.history_,[5]int{row1,col1,row2,col2,newValue})
}


func (this *SubrectangleQueries) GetValue(row int, col int) int {
	var l int = len(this.history_)
	for i := l - 1;i >= 0;i--{
		if row >= this.history_[i][0] && row <= this.history_[i][2] &&
			col >= this.history_[i][1] && col <= this.history_[i][3]{
			return this.history_[i][4]
		}
	}
	return this.rect_[row][col]
}