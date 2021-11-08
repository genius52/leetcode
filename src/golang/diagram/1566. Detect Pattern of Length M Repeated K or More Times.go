package diagram

//Input: arr = [1,2,1,2,1,1,1,3], m = 2, k = 2
//Output: true
func ContainsPattern(arr []int, m int, k int) bool{
	var l int = len(arr)
	var cnt int = 0
	var same int = 0
	for i := 0;i < l - m;i++{
		if arr[i] != arr[i + m]{
			cnt = 0
			same = 0
		}else{
			same++
			if same == m{
				same = 0
				cnt++
			}
		}
		if cnt == k - 1{
			return true
		}
	}
	return false
}

//func ContainsPattern(arr []int, m int, k int) bool {
//	var l int = len(arr)
//	for i := 0;i < l;i++{
//		var repeats int = 0
//		var last_s string
//		for j := i;j < l - m + 1;j += m{
//			var s string
//			for p := 0;p < m;p++{
//				s += strconv.Itoa(arr[j + p])
//			}
//			if len(last_s) == 0{
//				last_s = s
//				repeats++
//			}else{
//				if last_s == s{
//					repeats++
//				}else{
//					last_s = s
//					repeats = 1
//				}
//			}
//			if repeats >= k{
//				return true
//			}
//		}
//	}
//	return false
//}