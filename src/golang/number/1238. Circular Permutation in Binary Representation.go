package number

func CircularPermutation(n int, start int) []int{
	var record []int = []int{0}
	var start_idx int = 0
	for i := 1;i <= n;i++{
		pre_len := len(record)
		mask := 1 << (i - 1)
		for j := pre_len - 1;j >= 0;j--{
			next_num := record[j] | mask
			if next_num == start{
				start_idx = len(record)
			}
			record = append(record,next_num)
		}
	}
	res := record[start_idx:]
	res = append(res,record[:start_idx]...)
	return res
}
