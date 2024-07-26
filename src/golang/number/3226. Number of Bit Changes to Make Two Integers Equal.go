package number

func minChanges3226(n int, k int) int {
	var res int = 0
	var offset int = 0
	for n != k {
		if k|(1<<offset) == k { //k的offset位是1
			if n|(1<<offset) != n {
				return -1
			}
		} else { //k的offset位是0
			if n|(1<<offset) == n {
				k |= 1 << offset
				res++
			}
		}
		offset++
	}
	return res
}
