package diagram

func ValidSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	if (p1[0] == p2[0] && p1[1] == p2[1]) || (p1[0] == p3[0] && p1[1] == p3[1]) || (p1[0] == p4[0] && p1[1] == p4[1])||
		(p2[0] == p3[0] && p2[1] == p3[1])|| (p2[0] == p4[0] && p2[1] == p4[1]) || (p3[0] == p4[0] && p3[1] == p4[1]) {
		return false
	}
	var record map[int]int = make(map[int]int)
	dis12 := (p1[0] - p2[0]) * (p1[0] - p2[0]) + (p1[1] - p2[1]) *(p1[1] - p2[1])
	dis13 := (p1[0] - p3[0]) * (p1[0] - p3[0]) + (p1[1] - p3[1]) *(p1[1] - p3[1])
	dis14 := (p1[0] - p4[0]) * (p1[0] - p4[0]) + (p1[1] - p4[1]) *(p1[1] - p4[1])
	record[dis12]++
	record[dis13]++
	record[dis14]++
	dis23 := (p2[0] - p3[0]) * (p2[0] - p3[0]) + (p2[1] - p3[1]) *(p2[1] - p3[1])
	dis24 := (p2[0] - p4[0]) * (p2[0] - p4[0]) + (p2[1] - p4[1]) *(p2[1] - p4[1])
	if _,ok := record[dis23];!ok{
		return false
	}
	if _,ok := record[dis24];!ok{
		return false
	}
	record[dis23]++
	record[dis24]++
	dis34 := (p3[0] - p4[0]) * (p3[0] - p4[0]) + (p3[1] - p4[1]) *(p3[1] - p4[1])
	if _,ok := record[dis34];!ok{
		return false
	}
	record[dis34]++
	return len(record) <= 2
}