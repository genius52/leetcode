package number

//Input: a = 2, b = 6, c = 5
//Output: 3
func minFlips(a int, b int, c int) int {
	var mask int = 1
	var res int = 0
	var limit int = max_int_number(a,b,c)
	for mask <= limit{
		if (c | mask) == c{ // 1
			if (a | mask) != a && (b | mask) != b{
				res++
			}
		}else{//0
			if (a | mask) == a {//1
				res++
			}
			if (b | mask) == b{
				res++
			}
		}
		mask <<= 1
	}
	return res
}