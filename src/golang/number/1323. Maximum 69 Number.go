package number

import "math"

func Maximum69Number (num int) int{
	if num == 0{
		return num
	}
	a := int(math.Log10(float64(num)))
	interger := int(math.Pow10(a))
	div := num / interger
	if div == 6{
		rest := num - 6 * interger
		return 9 * interger + rest
	}else{
		rest := num - 9 * interger
		return 9 * interger + Maximum69Number(rest)
	}
}

//Input: num = 9669
//Output: 9969
//func maximum69Number (num int) int {
//	var rest int= 0
//	if 1000 < num && num < 10000{
//		if (num / 1000) == 6{
//			return num + 3000
//		}
//		rest = num % 1000
//	}else if 100 < num && num < 1000{
//		if (num /100) == 6 {
//			return num + 300
//		}
//		rest = num % 100
//	}else if 10 < num && num < 100{
//		if (num / 10 ) == 6{
//			return num + 30
//		}
//		rest = num % 10
//	}else{
//		if num == 6{
//			return 9
//		}
//		return num
//	}
//	return num - rest + maximum69Number(rest)
//}