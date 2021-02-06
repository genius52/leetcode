package number

import "strconv"

//43
//Input: num1 = "123", num2 = "456"
//Output: "56088"
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0"{
		return "0"
	}
	var long_num string
	var short_num string
	l1 := len(num1)
	l2 := len(num2)
	if l1 >=l2{
		long_num = num1
		short_num = num2
	}else{
		long_num = num2
		short_num = num1
	}
	var tmp [][]int = make([][]int,len(short_num))
	for i := 0;i < len(short_num);i++ {
		tmp[i] = make([]int,len(long_num))
		for j := 0;j < len(long_num);j++ {
			v1,err1 := strconv.Atoi(string(short_num[len(short_num) - 1 - i]))
			v2,err2 := strconv.Atoi(string(long_num[len(long_num) - 1 - j]))
			if nil == err1 && nil == err2{
				tmp[i][j] = v1 * v2
			}
		}
	}
	var bit_sum []int = make([]int,len(short_num) + len(long_num))
	for i := 0;i <(len(long_num) + len(short_num));i++{//i表示 个、十、百、千、万
		for j := 0;j < len(short_num) && (i - j) >=0 ;j++{//j表示
			if i - j >= len(long_num){
				continue
			}
			bit_sum[i] += tmp[j][i - j]
		}
	}
	var res string
	var upgrade int = 0
	if bit_sum[0] >= 10{
		upgrade = bit_sum[0] / 10
		res += strconv.Itoa(bit_sum[0] % 10)
	}else{
		res += strconv.Itoa(bit_sum[0])

	}

	for i := 1;i < len(bit_sum);i++{
		if i == (len(bit_sum)-1) && bit_sum[i] == 0{
			break
		}
		val := bit_sum[i] + upgrade
		res = strconv.Itoa(val % 10) + res
		upgrade = val / 10
	}
	if upgrade > 0{
		res = strconv.Itoa(upgrade) + res
	}
	return res
}
