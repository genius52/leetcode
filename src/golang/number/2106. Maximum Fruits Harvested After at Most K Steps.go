package number

func MaxTotalFruits(fruits [][]int, startPos int, k int) int {
	var l int = len(fruits)
	last_pos := max_int(fruits[l - 1][0],startPos + k)
	var prefix []int = make([]int,last_pos + 2)
	for i := 0;i < l;i++{
		prefix[fruits[i][0] + 1] = fruits[i][1]
	}
	for i := 1;i <= last_pos + 1;i++{
		if prefix[i] == 0{
			prefix[i] = prefix[i - 1]
		}else{
			prefix[i] += prefix[i - 1]
		}
	}
	var begin_pos int = 0
	if startPos - k > 0{
		begin_pos = startPos - k
	}
	var res int = 0
	for begin_pos <= startPos{
		var right int = startPos + (k - (startPos - begin_pos))/2
		cur := prefix[right + 1] - prefix[begin_pos]
		if cur > res{
			res = cur
		}
		begin_pos++
	}
	var end_pos int = startPos + k
	for end_pos >= startPos{
		var left int = startPos - (k - (end_pos - startPos))/2
		if left < 0{
			left = 0
		}
		cur := prefix[end_pos + 1] - prefix[left]
		if cur > res{
			res = cur
		}
		end_pos--
	}
	return res
}