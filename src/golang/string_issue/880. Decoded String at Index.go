package string_issue


func DecodeAtIndex(S string, K int) string{
	var cur_len int = 0
	var l int = len(S)
	var i int = 0
	for ;i < l;i++{
		if S[i] >= '0' && S[i] <= '9'{
			cur_len *= int(S[i] - '0')
		}else{
			cur_len++
		}
		if cur_len >= K{
			break
		}
	}
	for j := i;j >= 0;j--{
		///abcdabcdabcdabcdabcdabcd ,cur_len = 24,cnt = 6
		//k = 15
		if S[j] >= '0' && S[j] <= '9'{
			cnt := int(S[j] - '0')
			cur_len = cur_len / cnt
			K = K % cur_len //这两个位置字符必定相同
			//abcd，len = 4,k = 3
		}else{
			if K % cur_len == 0{//???
				return string(S[j])
			}
			cur_len--
		}
	}
	return ""
}