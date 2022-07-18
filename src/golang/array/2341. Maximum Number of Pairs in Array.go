package array

func numberOfPairs(nums []int) []int {
	var res []int = make([]int,2)
	var data [101]int
	for _,n := range nums{
		data[n]++
	}
	for i := 0;i <= 100;i++{
		if data[i] == 0{
			continue
		}
		if (data[i] | 1) == data[i]{
			res[1]++
		}
		res[0] += data[i]/2
	}
	return res
}