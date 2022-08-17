package number

import "strconv"

//n = 34567

func CountSpecialNumbers(n int) int {
	var s string = strconv.Itoa(n)
	var l int = len(s)
	if l == 1{
		return n
	}
	var pre []int = make([]int,l)
	pre[0] = 9
	var choice int = 9
	for i := 1;i < l;i++{
		pre[i] = pre[i - 1] * choice
		choice--
	}
	var res int = 0
	for i := 0;i < l - 1;i++{
		res += pre[i]
	}
	var visited [10]bool
	for i := 0;i < l;i++{
		var num int = int(s[i] - '0')
		var can_use1 int = 0
		for j := 0;j < num;j++{
			if !visited[j]{
				can_use1++
			}
		}
		if i == 0{
			can_use1--
		}
		var can_use2 int = 0
		for j := 0;j <= 9;j++{
			if !visited[j]{
				can_use2++
			}
		}
		can_use2--
		var total int = can_use1
		for j := i + 1;j < l;j++{
			total *= can_use2
			can_use2--
		}
		res += total
		if visited[num]{
			break
		}
		visited[num] = true
	}
	var visited2 [10]bool
	var special bool = false
	for n > 0{
		mod := n % 10
		n /= 10
		if visited2[mod]{
			special = true
			break
		}
		visited2[mod] = true
	}
	if special{
		return res
	}else{
		return res + 1
	}
}