package string_issue

import "strings"

//00 -> 10
//10 -> 01
//Input: binary = "000110"
//Output: "111011"
//Explanation: A valid transformation sequence can be:
//"000110" -> "000101"
//"000101" -> "100101"
//"100101" -> "110101"
//"110101" -> "110011"
//"110011" -> "111011"
func maximumBinaryString(binary string) string {
	var l int = len(binary)
	var data []string = make([]string,l)
	var idx int = 0
	for idx < l{
		if binary[idx] == '1'{
			data[idx] = "1"
			idx++
		}else{
			break
		}
	}
	if idx == l{
		return strings.Join(data,"")
	}
	var idx2 int = idx
	var one_cnt int = 0
	for idx2 < l{
		if binary[idx2] == '1'{
			one_cnt++
		}
		idx2++
	}
	for idx < l - 1 - one_cnt{
		data[idx] = "1"
		idx++
	}
	data[idx] = "0"
	idx++
	for one_cnt > 0{
		data[idx] = "1"
		one_cnt--
		idx++
	}
	return strings.Join(data,"")
}