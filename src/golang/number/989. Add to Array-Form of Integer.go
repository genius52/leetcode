package number

func addToArrayForm(num []int, k int) []int {
	var l int = len(num)
	var idx int = 0
	for k > 0{
		mod := k % 10
		if idx >= l{
			num = append([]int{mod},num...)
		}else{
			num[l - 1 - idx] += mod
		}
		idx++
		k /= 10
	}
	l = len(num)
	var upgrade bool = false
	for i := l - 1;i >= 0;i--{
		if upgrade{
			num[i] += 1
		}
		if num[i] > 9{
			upgrade = true
			num[i] -= 10
		}else{
			upgrade = false
		}
	}
	if upgrade{
		num = append([]int{1},num...)
	}
	return num
}