package number

//Input: arr = [1,2,3,4,5,10,6,7,8,9], k = 5
//Output: true
//Explanation: Pairs are (1,9),(2,8),(3,7),(4,6) and (5,10).
func CanArrange(arr []int, k int) bool {
	var l int = len(arr)
	var record map[int]int = make(map[int]int)
	for i := 0;i < l;i++{
		record[(arr[i] % k + k) % k]++
	}
	for val,cnt := range record{
		if val == 0{
			if cnt | 1 == cnt{//count is odd
				return false
			}
		}else{
			need := k - val
			if need == val{
				if cnt | 1 == cnt{//count is odd
					return false
				}
			}else{
				if cnt2,ok := record[need];!ok{
					return false
				}else{
					if cnt != cnt2{
						return false
					}
				}
			}
		}
	}
	return true
}