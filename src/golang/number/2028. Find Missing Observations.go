package number

func MissingRolls(rolls []int, mean int, n int) []int {
	var m int = len(rolls)
	var total int = mean * (m + n)
	var rest int = total
	for i := 0;i < m;i++{
		rest -= rolls[i]
	}
	if rest <= 0 || rest > (6 * n) || rest < n{
		return []int{}
	}
	var res []int = make([]int,n)
	var idx int = 0
	num := rest / n
	over := rest % n
	for n > 0{
		if over > 0{
			res[idx] = num + 1
			over--
		}else{
			res[idx] = num
		}
		idx++
		n--
	}
	return res
}