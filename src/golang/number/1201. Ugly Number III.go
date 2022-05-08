package number

func gcd64(a int64,b int64)int64{
	if b == 0 {
		return a
	}
	return gcd64(b, a%b)
}

func nthUglyNumber(n int, a int, b int, c int) int {
	var ab int64 = int64(a) * int64(b) / int64(gcd(a,b))
	var bc int64 = int64(b) * int64(c) / int64(gcd(b,c))
	var ac int64 = int64(a) * int64(c) / int64(gcd(a,c))
	var abc int64 = int64(a) * int64(bc) / int64(gcd64(int64(a),bc))
	var left int64 = 1
	var right int64 = 1e9
	for left < right{
		mid := (left + right)/2
		var cnt int64 = mid / int64(a) + mid/ int64(b) + mid / int64(c) - mid / ab - mid / bc - mid/ac +  mid / abc
		if cnt < int64(n){
			left = mid + 1
		}else{
			right = mid
		}
	}
	return int(left)
}