package diagram

//Input: points = [[1,3],[3,3],[5,3],[2,2]], queries = [[2,3,1],[4,3,1],[1,1,2]]
//Output: [3,2,2]
//Explanation: The points and circles are shown above.
//queries[0] is the green circle, queries[1] is the red circle, and queries[2] is the blue circle.
func cal_distance(x1,y1,x2,y2 int)int{
	return (x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2)
}

func CountPoints(points [][]int, queries [][]int) []int {
	var l int = len(queries)
	var res []int = make([]int,l)
	for i := 0;i < l;i++{
		var cnt int = 0
		for _,p := range points{
			dis := cal_distance(queries[i][0],queries[i][1],p[0],p[1])
			if dis <= (queries[i][2] * queries[i][2]){
				cnt++
			}
		}
		res[i] = cnt
	}
	return res
}