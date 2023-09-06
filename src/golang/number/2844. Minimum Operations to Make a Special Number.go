package number

// "25","50","75","00"
func MinimumOperations2844(num string) int {
	var l int = len(num)
	if l == 1 {
		if num == "0" {
			return 0
		} else {
			return 1
		}
	}
	var res int = l
	for i := l - 1; i > 0; i-- {
		if num[i] != '0' && num[i] != '5' {
			continue
		}
		if num[i] == '0' {
			cur := l - 1
			if cur < res {
				res = cur
			}
		}
		for j := i - 1; j >= 0; j-- {
			if num[i] == '0' {
				if num[j] == '5' || num[j] == '0' {
					//(l - 1 - i) + (i - 1 - j)
					cur := (l - 1 - i) + (i - 1 - j)
					if cur < res {
						res = cur
					}
				}
			} else if num[i] == '5' {
				if num[j] == '2' || num[j] == '7' {
					cur := (l - 1 - i) + (i - 1 - j)
					if cur < res {
						res = cur
					}
				}
			}
		}
	}
	return res
}
