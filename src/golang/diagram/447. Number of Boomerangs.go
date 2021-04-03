package diagram

func NumberOfBoomerangs(points [][]int) int {
	var res int = 0
	for i := 0;i < len(points);i++{
		target := points[i]
		var dis_record map[int]int = make(map[int]int)
		for j := 0;j < len(points);j++{
			if i == j{
				continue
			}
			dis := (target[0] - points[j][0]) * (target[0] - points[j][0]) + (target[1] - points[j][1]) * (target[1] - points[j][1])
			if _,ok := dis_record[dis];ok{
				dis_record[dis]++
			}else{
				dis_record[dis] = 1
			}
		}
		for _,v := range dis_record{
			if v <= 1{
				continue
			}
			res += v * (v - 1)
		}
	}
	return res
}