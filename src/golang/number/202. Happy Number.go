package number

//Input: n = 19
//Output: true
//Explanation:
//12 + 92 = 82
//82 + 22 = 68
//62 + 82 = 100
//12 + 02 + 02 = 1
func isHappy(n int) bool {
	var record map[int]bool = make(map[int]bool)
	for true{
		var cur int = n
		var sum int = 0
		for cur != 0{
			mod := cur % 10
			sum += mod * mod
			cur = cur/10
		}
		if sum == 1{
			return true
		}
		if _,ok := record[sum];ok{
			return false
		}else{
			record[sum] = true
		}
		n = sum
	}
	return false
}