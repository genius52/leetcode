package string_issue

//Input: s1 = "xxyyxyxyxx", s2 = "xyyxyxxxyx"
//Output: 4
//xx,yy cost xs1/2
//xy,yx cost
//xy,xx
//
func MinimumSwap(s1 string, s2 string) int {
	var l int = len(s1)
	var xs1 int = 0
	var ys1 int = 0
	var xs2 int = 0
	var ys2 int = 0
	for i := 0;i < l;i++{
		if s1[i] == s2[i]{
			continue
		}
		if s1[i] == 'x'{
			xs1++
		}else{
			ys1++
		}
		if s2[i] == 'x'{
			xs2++
		}else{
			ys2++
		}
	}
	if (xs1 + xs2) %2 != 0 || (ys1 + ys2) % 2 != 0{
		return -1
	}
	return xs1 / 2 + ys1 / 2 + xs1 % 2 + ys1 % 2
}