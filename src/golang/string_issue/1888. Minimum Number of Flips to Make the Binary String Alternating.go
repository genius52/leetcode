package string_issue

//Input: s = "111000"
//Output: 2
//Explanation: Use the first operation two times to make s = "100011".
//Then, use the second operation on the third and sixth elements to make s = "101010".
func MinFlips(s string) int {
	var l int = len(s)
	var zero_on_oddindex int = 0
	var zero_on_evenindex int = 0
	var one_on_oddindex int = 0
	var one_on_evenindex int = 0
	for i := 0;i < l;i++{
		if s[i] == '1'{
			if (i | 1) == i{
				one_on_oddindex++
			}else{
				one_on_evenindex++
			}
		}else{
			if (i | 1) == i{
				zero_on_oddindex++
			}else{
				zero_on_evenindex++
			}
		}
	}
	var res int = min_int(zero_on_oddindex + one_on_evenindex,zero_on_evenindex + one_on_oddindex)
	if (l | 1) != l{
		return res
	}
	for i := 0;i < l;i++{
		//move current char to the end
		if (i | 1) == i{ //odd index
			if s[i] == '0'{
				zero_on_oddindex--
				zero_on_evenindex++
			}else{
				one_on_oddindex--
				one_on_evenindex++
			}
		}else{//even index
			if s[i] == '0'{
				zero_on_evenindex--
				zero_on_oddindex++
			}else{
				one_on_evenindex--
				one_on_oddindex++
			}
		}
		res = min_int(res,zero_on_oddindex + one_on_evenindex)
		res = min_int(res,zero_on_evenindex + one_on_oddindex)
	}
	return res
}