package number

func MostFrequentEven(nums []int) int {
	var record map[int]int = make(map[int]int)
	for _,n := range nums{
		if n | 1 != n{
			record[n]++
		}
	}
	var small_num int = 1 << 31 - 1
	var max_freq int = 0
	for n,cnt := range record{
		if cnt > max_freq{
			small_num = n
			max_freq = cnt
		}else if cnt == max_freq{
			if n < small_num{
				small_num = n
			}
		}
	}
	if small_num == 1 << 31 - 1{
		return -1
	}
	return small_num
}