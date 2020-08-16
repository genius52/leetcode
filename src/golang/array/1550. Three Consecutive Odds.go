package array

//Input: arr = [1,2,34,3,4,5,7,23,12]
//Output: true
//Explanation: [5,7,23] are three consecutive odds.
func threeConsecutiveOdds(arr []int) bool {
	var l int = len(arr)
	for i := 1;i < l - 1;{
		if arr[i - 1] % 2 == 1{
			if arr[i] % 2 == 1{
				if arr[i + 1] % 2 == 1{
					return true
				}else{
					i++
				}
			}else{
				i += 2
			}
		}else{
			if arr[i] % 2 == 1{
				if arr[i + 1] % 2 == 1{
					i++
				}else{
					i += 2
				}
			}else{
				i += 2
			}
		}
	}
	return false
}
