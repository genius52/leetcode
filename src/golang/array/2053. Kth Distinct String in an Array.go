package array

func kthDistinct(arr []string, k int) string {
	var l int = len(arr)
	var record map[string]bool = make(map[string]bool)
	for i := 0;i < l;i++{
		if _,ok := record[arr[i]];ok{
			record[arr[i]] = false
		}else{
			record[arr[i]] = true
		}
	}
	for i := 0;i < l;i++{
		if record[arr[i]]{
			k--
			if k == 0{
				return arr[i]
			}
		}
	}
	return ""
}