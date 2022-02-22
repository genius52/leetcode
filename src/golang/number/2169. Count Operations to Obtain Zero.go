package number

func countOperations(num1 int, num2 int) int{
	var res int = 0
	for num1 != 0 && num2 != 0{
		if num1 < num2{
			num1,num2 = num2,num1
		}
		res += num1 / num2
		num1 %= num2
	}
	return res
}

//func countOperations(num1 int, num2 int) int {
//	var res int = 0
//	for num1 != 0 && num2 != 0{
//		if num1 >= num2{
//			num1 -= num2
//		}else{
//			num2 -= num1
//		}
//	}
//	return res
//}