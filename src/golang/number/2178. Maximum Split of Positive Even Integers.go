package number

func MaximumEvenSplit(finalSum int64) []int64 {
	var res []int64
	if finalSum | 1 == finalSum{
		return res
	}
	var i int64 = 2
	for ;i <= finalSum;i += 2{
		//n*a1+n(n-1)d/2
		cnt := i/2
		sum := cnt * 2 + cnt * (cnt - 1)
		rest := finalSum - sum
		if rest < sum{
			var j int64 = 2
			for ;j <= finalSum;j += 2{
				res = append(res,j)
				finalSum -= j
			}
			res[len(res) - 1] += finalSum
			return res
		}
	}
	return res
}