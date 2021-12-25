package string_issue

import "bytes"

//Input: s = "a1b2c3d4e"
//Output: "abbdcfdhe"
//Explanation: The digits are replaced as follows:
//- s[1] -> shift('a',1) = 'b'
//- s[3] -> shift('b',2) = 'd'
//- s[5] -> shift('c',3) = 'f'
//- s[7] -> shift('d',4) = 'h'
func replaceDigits(s string) string {
	var res string
	var l int = len(s)
	for i := 1; i < l; i += 2 {
		var c uint8 = s[i-1]
		dis := s[i] - '0'
		res += string(c) + string('a'+(c-'a'+uint8(dis))%26)
	}
	if (l | 1) == l {
		res += string(s[l-1])
	}
	return res
}

func replaceDigits2(s string) string {
	var data bytes.Buffer
	var l int = len(s)
	i := 1
	for ; i < l; i += 2 {
		var c uint8 = s[i-1]
		dis := s[i] - '0'
		data.WriteByte(c)
		data.WriteByte('a' + (c-'a'+uint8(dis))%26)
	}
	if l|1 == l {
		data.WriteByte(s[l-1])
	}
	return data.String()
}
