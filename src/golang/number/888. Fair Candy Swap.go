package number

func FairCandySwap(aliceSizes []int, bobSizes []int) []int {
	var sum1 int = 0
	var sum2 int = 0
	var cnt2 map[int]bool = make(map[int]bool)
	for _,n := range aliceSizes{
		sum1 += n
	}
	for _,n := range bobSizes{
		sum2 += n
		cnt2[n] = true
	}
	var res []int = make([]int,2)
	diff := (sum1 - sum2)/2
	for _,n := range aliceSizes{
		var find int = 0
		find = n - diff
		if _,ok := cnt2[find];ok{
			res[0] = n
			res[1] = find
			break
		}
	}
	return res
}