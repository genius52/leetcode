package number

func calculateTax(brackets [][]int, income int) float64 {
	var l int = len(brackets)
	var res float64 = 0
	var pre int = 0
	for i := 0;i < l;i++{
		if income <= brackets[i][0]{
			res += float64(income - pre) * float64(brackets[i][1])/100
			break
		}else{
			res += float64(brackets[i][0] - pre) * float64(brackets[i][1])/100
		}
		pre = brackets[i][0]
	}
	return res
}