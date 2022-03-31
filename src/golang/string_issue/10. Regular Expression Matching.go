package string_issue

//Input:
//s = "ab"
//p = ".*"
//Output: true
//Input:
//s = "aab"
//p = "c*a*b"
//Output: true
func dp_isMatch(s string,p string)bool{
	if len(s) == 0{
		if len(p) == 0{
			return true
		}
	}
	if len(p) == 0{
		return false
	}
	var l1 int = len(s)
	var l2 int = len(p)
	if l2 > 1{
		if p[1] == '*'{
			if p[0] == '.'{
				for i := 0;i <= l1;i++{
					ret := dp_isMatch(s[i:],p[2:])
					if ret{
						return true
					}
				}
				return false
			}else{
				for i := 0;i <= l1;i++{
					if i != l1 && s[i] == p[0]{
						ret := dp_isMatch(s[i:],p[2:])
						if ret{
							return true
						}
					}else{
						ret := dp_isMatch(s[i:],p[2:])
						if ret{
							return true
						}
						break
					}
				}
				return dp_isMatch(s,p[2:])
			}
		}else{
			if p[0] == '.'{
				if l1 >= 1{
					return dp_isMatch(s[1:],p[1:])
				}
				return false
			}else{
				if l1 >= 1{
					if p[0] == s[0]{
						return dp_isMatch(s[1:],p[1:])
					}
				}
				return false
			}
		}
	}else{
		if p[0] == '.'{
			return l1 == l2
		}else{
			if l1 >= 1{
				return p[0] == s[0] && l1 == l2
			}
			return false
		}
	}
	return false
}

func IsMatch(s string, p string) bool {
	return dp_isMatch(s,p)
}