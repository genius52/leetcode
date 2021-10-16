package number

import "strconv"

//func gcd(a int,b int)int{
//	if b == 0 {
//		return a
//	}
//	return gcd(b, a%b)
//}

func simplifiedFractions(n int) []string {
	var res []string
	for i := 2;i <= n;i++{
		for j := 1 ;j < i;j++{
			if gcd(i,j) == 1{
				res = append(res,strconv.Itoa(j) + "/" + strconv.Itoa(i))
			}
		}
	}
	return res
}