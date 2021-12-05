package number

import "sort"

func findEvenNumbers(digits []int) []int {
	var l int = len(digits)
	var record map[int]bool = make(map[int]bool)
	for i := 0; i < l; i++ {
		if digits[i] == 0 {
			continue
		}
		for j := 0; j < l; j++ {
			if j == i {
				continue
			}
			for k := 0; k < l; k++ {
				if k == i || k == j || digits[k]|1 == digits[k] {
					continue
				}
				n := 100*digits[i] + 10*digits[j] + digits[k]
				if n != 0 {
					record[n] = true
				}
			}
		}
	}
	var res []int
	for n, _ := range record {
		res = append(res, n)
	}
	sort.Ints(res)
	return res
}
