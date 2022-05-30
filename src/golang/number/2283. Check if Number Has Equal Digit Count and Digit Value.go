package number

func digitCount(num string) bool {
	var l int = len(num)
	var val_cnt []int = make([]int,l)
	for i := 0;i < l;i++{
		n := int(num[i] - '0')
		if n >= l{
			return false
		}
		val_cnt[n]++
	}
	for i := 0;i < l;i++{
		cur := int(num[i] - '0')
		if cur != val_cnt[i]{
			return false
		}
	}
	return true
}