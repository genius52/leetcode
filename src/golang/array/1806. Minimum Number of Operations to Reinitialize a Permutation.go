package array

//如果 i % 2 == 0 ，那么 arr[i] = perm[i / 2]
//如果 i % 2 == 1 ，那么 arr[i] = perm[n / 2 + (i - 1) / 2]
func ReinitializePermutation(n int) int {
	var steps int = 0
	idx := 1
	for true {
		var cur int = 0
		if idx|1 == idx {
			cur = n/2 + (idx-1)/2
		} else {
			cur = idx / 2
		}
		steps++
		if cur == 1 {
			break
		} else {

			idx = cur
		}
	}
	return steps
}
