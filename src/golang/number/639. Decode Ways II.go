package number

//Input: s = "2*"
//Output: 15
//Explanation: The encoded message can represent any of the encoded messages
//"21", "22", "23", "24", "25", "26", "27", "28", or "29".
//"21", "22", "23", "24", "25", and "26" have 2 ways of being decoded, but "27", "28", and "29" only have 1 way.
//Hence, there are a total of (6 * 2) + (3 * 1) = 12 + 3 = 15 ways to decode "2*".
//func dfs_numDecodings(s string,l int,pos int)int{
//	if pos >= l{
//		return 0
//	}
//	if s[pos] == '0'{
//		return 0
//	}
//	if pos == l - 1{
//		if s[pos] != '*'{
//			return 1
//		}else{
//			return 9
//		}
//	}
//	var res int = 0
//	if s[pos] >= '3' && s[pos] <= '9'{
//		res = dfs_numDecodings(s,l,pos + 1)
//	}else if s[pos] == '2'{
//		res += dfs_numDecodings(s,l,pos + 1)//single 2
//		if pos < l - 1{
//			if s[pos + 1] >= '0' && s[pos + 1] <= '6'{
//				if pos == l - 2{
//					res++
//				}else{
//					res += dfs_numDecodings(s,l,pos + 2)
//				}
//			}else if s[pos + 1] == '*'{
//				res += 6
//				res += dfs_numDecodings(s,l,pos + 2)
//			}
//		}
//	}else if s[pos] == '1'{
//		res += dfs_numDecodings(s,l,pos + 1)
//		if pos < l - 1{
//			if pos == l - 2{
//				if s[pos + 1] == '*'{
//					res += 9
//				}else{
//					res++
//				}
//			}else{
//				if s[pos + 1] == '*'{
//					res += 9
//				}
//				res += dfs_numDecodings(s,l,pos + 2)
//			}
//			//if s[pos + 1] == '*'{
//			//	res += 9
//			//}
//			//if pos == l - 2{
//			//	res++
//			//}else{
//			//	res += dfs_numDecodings(s,l,pos + 2)
//			//}
//		}
//	}else if s[pos] == '0'{
//		res = 0
//	}else if s[pos] == '*'{
//		res = 9 * dfs_numDecodings(s,l,pos + 1)//* = single number
//		if pos < l - 1{//number XX
//			if pos == l - 2{
//				if (s[pos + 1] >= '0' && s[pos + 1] <= '6'){
//					res += 2//1x 2x
//				}else{
//					res++//17 18 19
//				}
//			}else{
//				if (s[pos + 1] >= '0' && s[pos + 1] <= '6'){//*= 2X
//					res += dfs_numDecodings(s,l,pos + 2)//begin from new position
//				}
//				res += dfs_numDecodings(s,l,pos + 2)//* = 1X
//			}
//		}
//	}
//	res = res % 1000000007
//	return res
//}
//
//func NumDecodings(s string) int {
//	var l int = len(s)
//	return dfs_numDecodings(s,l,0)
//}

func NumDecodings(s string) int {
	var l int = len(s)

	for i := 0;i < l;i++{

	}
	return 0
}