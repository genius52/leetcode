package diagram

//Input: bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
//Output: [10,55,45,25,25]
func CorpFlightBookings(bookings [][]int, n int) []int {
	var res []int = make([]int,n)
	var l int = len(bookings)
	for i := 0;i < l;i++{
		for j := bookings[i][0];j <= bookings[i][1];j++{
			res[j - 1] += bookings[i][2]
		}
	}
	return res
}