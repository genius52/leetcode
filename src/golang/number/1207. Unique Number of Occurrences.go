package number

func uniqueOccurrences(arr []int) bool {
	var l int = len(arr)
	var cnt [2001]int
	for i := 0;i < l;i++{
		cnt[arr[i] + 1000]++
	}
	var frequency [2001]bool
	for i := 0;i <= 2000;i++{
		if cnt[i] != 0{
			if frequency[cnt[i]]{
				return false
			}else{
				frequency[cnt[i]] = true
			}
		}
	}
	return true
}