package array

//Input: trips = [[2,1,5],[3,3,7]], capacity = 4
//Output: false
func carPooling(trips [][]int, capacity int) bool {
	var record []int = make([]int,1001)
	for _,t := range trips{
		record[t[1]] += t[0]
		record[t[2]] -= t[0]
	}
	var passengers int = 0
	for i := 0;i <= 1000;i++{
		passengers += record[i]
		if passengers > capacity{
			return false
		}
	}
	return true
}