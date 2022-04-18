package number

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	var res int64 = 0
	for i := 0;cost1 * i <= total;i++{
		rest := total - cost1 * i
		if rest == 0{
			res++
		}else if rest > 0{
			if rest /cost2 > 0{
				res += int64(rest /cost2) + 1
			}else{
				res++
			}
		}
	}
	return res
}