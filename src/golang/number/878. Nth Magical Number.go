package number

//Input: n = 5, a = 2, b = 4
//Output: 10
//
//Input: n = 3, a = 6, b = 4
//Output: 8
func gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	} else {
		return y
	}
}

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

func NthMagicalNumber(n int, a int, b int) int {
	if a % b == 0 || b % a == 0{
		return n * min_int(a,b) % 1000000007
	}
	m := lcm(a,b)
	var low int64 = 2
	var high int64 = int64(n * min_int(a, b))
	//var res int = 2
	for low < high{
		mid := low + (high - low)/2
		cnt := mid / int64(a) + mid / int64(b) - mid /int64(m)
		if cnt == int64(n){
			high = mid - 1
			return max_int_number(int(mid) / a * a,int(mid) / b * b,int(mid) / m * m) % 1000000007
		}
		if cnt < int64(n){
			low = mid + 1
		}else{
			high = mid - 1
		}
	}
	return int(low) % 1000000007
}