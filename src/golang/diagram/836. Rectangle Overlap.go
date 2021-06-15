package diagram

//Input: rec1 = [0,0,2,2], rec2 = [1,1,3,3]
func IsRectangleOverlap(rec1 []int, rec2 []int) bool {
	left1 := rec1[0]
	bottom1 := rec1[1]
	right1 := rec1[2]
	top1 := rec1[3]

	left2 := rec2[0]
	bottom2 := rec2[1]
	right2 := rec2[2]
	top2 := rec2[3]

	if rec1[0] == rec1[2] || rec1[1] == rec1[3] || rec2[0] == rec2[2] || rec2[1] == rec2[3] {
		return false // it is a line
	}

	if bottom1 >= top2 || bottom2 >= top1 || left1 >= right2 || left2 >= right1{
		return false
	}
	return true
}