package number

import (
	"strconv"
	"strings"
)

func MaxValue(n string, x int) string {
	if n == "0"{
		return strconv.Itoa(x)
	}
	var is_neg bool = false
	if n[0] == '-' {
		is_neg = true
	}
	var s_len int = len(n)
	var data_len int = s_len
	if is_neg{
		data_len--
	}
	var data []string = make([]string,data_len)
	var i int = 0
	if is_neg{
		i++
	}
	var index int = 0
	for ;i < s_len;i++{
		data[index] = string(n[i])
		index++
	}
	var l int = len(data)
	var visit int = 0
	if !is_neg{
		for visit < l{
			cur,_ := strconv.Atoi(data[visit])
			if cur < x{
				rear := append([]string{},data[visit:]...)
				data = append(data[:visit],strconv.Itoa(x))
				data = append(data,rear...)
				break
			}else{
				visit++
			}
		}
	}else{
		for visit < l{
			cur,_ := strconv.Atoi(data[visit])
			if cur > x{
				rear := append([]string{},data[visit:]...)
				data = append(data[:visit],strconv.Itoa(x))
				data = append(data,rear...)
				break
			}else{
				visit++
			}
		}
	}

	if visit == l{
		data = append(data,strconv.Itoa(x))
	}
	var res string = strings.Join(data,"")
	if is_neg{
		res = "-" + res
	}
	return res
}