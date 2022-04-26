package diagram

func countLatticePoints(circles [][]int) int {
	var visited map[int]bool = make(map[int]bool)
	var res int = 0
	var l int = len(circles)
	for i := 0;i < l;i++{
		x := circles[i][0]
		y := circles[i][1]
		r := circles[i][2]
		for m := x - r;m <= x + r;m++{
			for n := y - r;n <= y + r;n++{
				dis := (m - x) * (m - x) + (n - y) * (n - y)
				if dis <= (r * r){
					key := m * 1000 + n
					if !visited[key]{
						visited[key] = true
						res++
					}
				}
			}
		}
	}
	return res
}