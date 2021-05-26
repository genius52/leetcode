package number

//2、3、5、7、11、13、17、19、23、29、31
func is_prime(n int)bool{
	var cnt int = 0
	for n != 0{
		if (n | 1) == n{
			cnt++
		}
		n = n >> 1
	}
	return cnt == 2 || cnt == 3 || cnt == 5|| cnt == 7|| cnt == 11|| cnt == 13|| cnt == 17|| cnt == 19 ||
	 		cnt == 23|| cnt == 29|| cnt == 31
}

func countPrimeSetBits(L int, R int) int {
	var res int = 0
	for i := L;i <= R;i++{
		if is_prime(i){
			res++
		}
	}
	return res
}