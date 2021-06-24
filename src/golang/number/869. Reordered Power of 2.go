package number

func ReorderedPowerOf2(N int) bool{
	var n int = N
	var l1 int
	var cnt [10]int
	for n > 0{
		cnt[n % 10]++
		n /= 10
		l1++
	}
	var pow_num int = 1
	var l2 int = 1
	for l2 < l1{
		pow_num <<= 1
		var n2 int = pow_num
		var cur_l int = 0
		for n2 > 0{
			n2 /= 10
			cur_l++
		}
		l2 = cur_l
	}
	var record []int
	for l2 == l1{
		record = append(record,pow_num)
		pow_num <<= 1
		var n2 int = pow_num
		var cur_l int = 0
		for n2 > 0{
			n2 /= 10
			cur_l++
		}
		l2 = cur_l
	}
	for _,num := range record{
		var cur_cnt [10]int
		for num > 0{
			cur_cnt[num%10]++
			num /= 10
		}
		var equal bool = true
		for i := 0;i < 10;i++{
			if cur_cnt[i] != cnt[i]{
				equal = false
				break
			}
		}
		if equal{
			return true
		}
	}
	return false
}

//func ReorderedPowerOf2(N int) bool {
//	var n int = 1
//	var record map[string]bool = make(map[string]bool)
//	for n <= 2147483647{
//		var arr []int
//		var cur_num int = n
//		for cur_num > 0{
//			arr = append(arr,cur_num % 10)
//			cur_num = cur_num / 10
//		}
//		sort.Ints(arr)
//		var s string
//		for _,b:= range arr{
//			s += strconv.Itoa(b)
//		}
//		record[s] = true
//		n = n << 1
//	}
//	var arr2 []int
//	for N > 0{
//		arr2 = append(arr2,N % 10)
//		N = N / 10
//	}
//	sort.Ints(arr2)
//	var s2 string
//	for _,b:= range arr2{
//		s2 += strconv.Itoa(b)
//	}
//	if _,ok := record[s2];ok{
//		return true
//	}
//	return false
//}