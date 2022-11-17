package string_issue



func MaxPalindromes(s string, k int) int{
	var l int = len(s)
	if k == 1{
		return l
	}
	var dp [][]bool = make([][]bool,l)
	for i := 0;i < l;i++{
		dp[i] = make([]bool,l)
	}
	for i := 0;i < l;i++{
		var left int = i
		var right int = i
		for left >= 0 && right < l{
			if s[left] != s[right]{
				break
			}
			dp[left][right] = true
			left--
			right++
		}
	}
	for i := 0;i < l;i++{
		var left int = i
		var right int = i + 1
		for left >= 0 && right < l{
			if s[left] != s[right]{
				break
			}
			dp[left][right] = true
			left--
			right++
		}
	}
	var cnt []int = make([]int,l + 1)
	for i := 1;i < l;i++{
		cnt[i] = cnt[i - 1]
		if i - k + 1 >= 0 && dp[i - k + 1][i]{
			if i - k >= 0{
				cnt[i] = max_int(cnt[i],1 + cnt[i - k])
			}else{
				cnt[i] = max_int(cnt[i],1)
			}
		}
		if i - k >= 0 && dp[i - k][i]{
			if i - k - 1 >= 0{
				cnt[i] = max_int(cnt[i],1 + cnt[i - k - 1])
			}else{
				cnt[i] = max_int(cnt[i],1)
			}
		}
	}
	return cnt[l - 1]
}

//len,end_idx
//func check_palindromic(s string,left int,right int,target_len int)(int,int){
//	l := len(s)
//	for left >= 0 && right < l && right - left + 1 < target_len{
//		if s[left] != s[right]{
//			return 0,-1
//		}
//		left--
//		right++
//	}
//	if s[left] == s[right]{
//		return right - left + 1,right
//	}
//	return 0,-1
//}

//func MaxPalindromes(s string, k int) int {
//	var l int = len(s)
//	if k == 1{
//		return l
//	}
//	var cnt []int = make([]int,l)//dp[i], s[:i]最多的回文串个数
//	var dp []int = make([]int,l)//dp[i]: 以i结尾的字符串 s[i - 2 * k:i]为回文串
//	for i := k/2;i + k/2 < l;i++{
//		ret1,end_idx1 := check_palindromic(s,i,i,k)
//		if ret1 > 0{
//			if dp[end_idx1] != 0{
//				dp[end_idx1] = min_int(dp[end_idx1],ret1)
//			}else{
//				dp[end_idx1] = ret1
//			}
//		}
//		ret2,end_idx2 := check_palindromic(s,i,i + 1,k)
//		if ret2 > 0{
//			if dp[end_idx2] != 0{
//				dp[end_idx2] = min_int(dp[end_idx2],ret2)
//			}else{
//				dp[end_idx2] = ret2
//			}
//		}
//	}
//	for i := k - 1;i < l;i++{
//		if dp[i] > 0{
//			if i - k < 0{
//				cnt[i] = 1
//			}else{
//				if i - dp[i] >= 0{
//					cnt[i] = 1 + cnt[i - dp[i]]
//				}else{
//					cnt[i] = 1
//				}
//			}
//		}else{
//			cnt[i] = cnt[i - 1]
//		}
//	}
//	return cnt[l - 1]
//}