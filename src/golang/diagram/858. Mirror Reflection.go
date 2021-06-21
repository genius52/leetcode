package diagram

//Input: p = 2, q = 1
//Output: 2
//Explanation: The ray meets receptor 2 the first time it gets reflected back to the left wall.
func mirrorReflection(p int, q int) int {
	if q == 0{
		return 0
	}
	var k int = 1
	for true{
		rounds := p * k / q
		mod := p * k % q
		if mod == 0{
			if rounds % 2 == 0{//even round
				return 2
			}else{
				if (k | 1) == k{ //odd round and do not cross current square
					return 1
				}else{
					return 0
				}
			}
		}
		k++
	}
	return -1
}