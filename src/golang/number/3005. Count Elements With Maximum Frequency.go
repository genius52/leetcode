package number

func maxFrequencyElements(nums []int) int {
	var max_frequency int = 0
	var record map[int]int = make(map[int]int)
	for _, n := range nums {
		record[n]++
		if record[n] > max_frequency {
			max_frequency = record[n]
		}
	}
	var res int = 0
	for _, cnt := range record {
		if cnt == max_frequency {
			res += max_frequency
		}
	}
	return res
}
