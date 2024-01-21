package array

func BeautifulIndices(s string, a string, b string, k int) []int {
	var l_s int = len(s)
	var l_a int = len(a)
	var l_b int = len(b)
	var record_a []int
	for left := 0; left < l_s && left+l_a <= l_s; left++ {
		right_a := left + l_a
		if s[left:right_a] == a {
			record_a = append(record_a, left)
		}
	}
	var record_b []int
	for left := 0; left < l_s && left+l_b <= l_s; left++ {
		right_b := left + l_b
		if s[left:right_b] == b {
			record_b = append(record_b, left)
		}
	}
	var res []int
	var idx_b int = 0
	for i := 0; i < len(record_a); i++ {
		var find bool = false
		for idx_b < len(record_b) && record_b[idx_b]-record_a[i] <= k {
			if abs_int(record_b[idx_b]-record_a[i]) <= k {
				find = true
				break
			}
			idx_b++
		}
		if find {
			res = append(res, record_a[i])
		}
	}
	return res
}
