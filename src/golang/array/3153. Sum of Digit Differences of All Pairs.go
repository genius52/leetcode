package array

import "strconv"

func SumDigitDifferences(nums []int) int64 {
	s := strconv.Itoa(nums[0])
	var l2 int = len(s)
	var bit_cnt [][10]int = make([][10]int, l2)
	var res int64 = 0
	for i := 0; i < len(nums); i++ {
		var offset int = 0
		for nums[i] > 0 {
			m := nums[i] % 10
			bit_cnt[offset][m]++
			nums[i] /= 10
			offset++
		}
	}
	for i := 0; i < l2; i++ {
		var data []int
		for j := 0; j < 10; j++ {
			if bit_cnt[i][j] != 0 {
				data = append(data, bit_cnt[i][j])
			}
		}
		if len(data) > 1 {
			for m := 0; m < len(data); m++ {
				for n := m + 1; n < len(data); n++ {
					res += int64(data[m]) * int64(data[n])
				}
			}
		}
	}
	return res
}
