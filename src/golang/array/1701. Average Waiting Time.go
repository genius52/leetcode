package array

//Input: customers = [[1,2],[2,5],[4,3]]
//Output: 5.00000
func AverageWaitingTime(customers [][]int) float64 {
	var l int = len(customers)
	var total int = customers[0][1]
	var pre_finish int = customers[0][0] + customers[0][1]
	for i := 1;i < l;i++{
		if pre_finish <= customers[i][0]{
			total += customers[i][1]
			pre_finish = customers[i][0] + customers[i][1]
		}else{
			total += pre_finish - customers[i][0] + customers[i][1]
			pre_finish = pre_finish  + customers[i][1]
		}
	}
	return float64(total)/float64(l)
}