package number

func totalMoney(n int) int{
	var days int = 0
	var cur int = 1
	var total int = 0
	for days < n{
		if days != 0 && days % 7 == 0{
			cur -= 6
		}
		total += cur
		days++
		cur++
	}
	return total
}