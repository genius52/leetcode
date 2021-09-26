package number

import "sort"

//1356
type sortBit []int

func (s sortBit) Less(i, j int) bool {
	var bit_cnt1 int = 0
	var bit_cnt2 int = 0
	var move int = 1
	for s[i] >= move{
		if (s[i] | move) == s[i]{
			bit_cnt1++
		}
		move = move << 1
	}
	move = 1
	for s[j] >= move{
		if (s[j] | move) == s[j]{
			bit_cnt2++
		}
		move = move << 1
	}
	if bit_cnt1 == bit_cnt2{
		return s[i] < s[j]
	}
	return bit_cnt1 < bit_cnt2
}

func (s sortBit) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortBit) Len() int {
	return len(s)
}

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		var bit1 int = 0
		var bit2 int = 0
		num1 := arr[i]
		num2 := arr[j]
		for num1 > 0{
			if (num1 | 1) == num1{
				bit1++
			}
			num1 >>= 1
		}
		for num2 > 0{
			if (num2 | 1) == num2{
				bit2++
			}
			num2 >>= 1
		}
		if bit1 == bit2{
			return arr[i] < arr[j]
		}
		return bit1 < bit2
	})
	return arr
}