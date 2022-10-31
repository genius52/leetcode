package number

func averageValue(nums []int) int {
	var sum int = 0
	var cnt int = 0
	for _,n := range nums{
		if n != 0 && n % 3 == 0{
			n2 := n / 3
			if (n2 | 1) != n2{
				sum += n
				cnt++
			}
		}
	}
	if cnt == 0{
		return 0
	}
	return sum / cnt
}