package number

//输入：n = 3000
//输出：3133
//解释：
//3133 是一个数值平衡数，因为：
//- 数字 1 出现 1 次。
//- 数字 3 出现 3 次。
//这也是严格大于 3000 的最小数值平衡数。
func check_num(n int)bool{
	var num_cnt [10]int
	for n > 0{
		mod := n % 10
		num_cnt[mod]++
		n = n / 10
	}
	for i := 0;i < 10;i++{
		if num_cnt[i] == 0{
			continue
		}
		if num_cnt[i] != i{
			return false
		}
	}
	return true
}

func nextBeautifulNumber(n int) int {
	var res int = 0
	n++
	for true{
		if check_num(n){
			return n
		}
		n++
	}
	return res
}