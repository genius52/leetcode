package string_issue

//Input:
//s = "ab"
//p = ".*"
//Output: true
//Input:
//s = "aab"
//p = "c*a*b"
//Output: true
func dp_isMatch(s string,p string,idx_s int,idx_p int,last_p uint8)bool{
	if idx_s >= len(s){
		return true
	}
	if idx_p >= len(p){
		return false
	}
	return false
}

func isMatch(s string, p string) bool {
	return dp_isMatch(s,p,0,0,0)
}