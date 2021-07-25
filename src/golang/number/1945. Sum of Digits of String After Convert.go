package number


//输入：s = "leetcode", k = 2
//输出：6
//解释：操作如下：
//- 转化："leetcode" ➝ "(12)(5)(5)(20)(3)(15)(4)(5)" ➝ "12552031545" ➝ 12552031545
//- 转换 #1：12552031545 ➝ 1 + 2 + 5 + 5 + 2 + 0 + 3 + 1 + 5 + 4 + 5 ➝ 33
//- 转换 #2：33 ➝ 3 + 3 ➝ 6
//因此，结果整数为 6 。
func getLucky(s string, k int) int {
	var data []int
	for _,c := range s{
		var n int = int(c - 'a' + 1)
		for n > 0{
			mod := n % 10
			n = n / 10
			data = append(data,mod)
		}
	}
	var res int = 0
	for k > 0{
		var l int = len(data)
		var total int = 0
		for i := 0;i < l;i++{
			total += data[i]
		}
		res = total
		var cur []int
		for total > 0{
			mod := total % 10
			total = total /10
			cur = append(cur,mod)
		}
		data = cur
		k--
	}
	return res
}