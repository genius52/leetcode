package number

//Input: arr = [2,3,4,7,11], k = 5
//Output: 9
func FindKthPositive(arr []int, k int) int{
	var l int = len(arr)
	lack_cnt := arr[l - 1] - l
	if k > lack_cnt{
		return arr[l - 1] + k - lack_cnt
	}
	need_number := 1
	for i := 0;i < l;i++{
		if arr[i] != need_number{
			diff := arr[i] - need_number
			if diff >= k{
				return need_number + k - 1
			}
			k -= diff
			need_number = arr[i] + 1
		}else{
			need_number++
		}
	}
	return -1
}

func FindKthPositive2(arr []int, k int) int {
	var l int = len(arr)
	var pos int =
	var number int = 1
	for pos < l{
		for arr[pos] != number{
			k--
			if(k == 0){
				return number
			}
			number++
		}
		number++
		pos++
	}
	return arr[l - 1] + k
}