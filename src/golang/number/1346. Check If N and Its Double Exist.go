package number

func checkIfExist(arr []int) bool {
	var record map[int]bool = make(map[int]bool)
	for _,n := range arr{
		if n % 2 == 0{
			small := n / 2
			if _,ok := record[small];ok{
				return true
			}
		}
		big := n * 2
		if _,ok := record[big];ok{
			return true
		}
		record[n] = true
	}
	return false
}