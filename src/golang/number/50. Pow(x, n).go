package number


//20
//Input: 2.00000, -2
//Output: 0.25000
func myPow(x float64, n int) float64 {
	if n == 0{
		return 1
	}
	if n == 1{
		return x
	}
	if n < 0{
		return myPow(1 / x,-n)
	}
	if n == 2{
		return x * x
	}
	if n % 2 == 0{
		return myPow(myPow(x,n/2),2)
	}else{
		return x * myPow(myPow(x,n/2), 2)
	}
}